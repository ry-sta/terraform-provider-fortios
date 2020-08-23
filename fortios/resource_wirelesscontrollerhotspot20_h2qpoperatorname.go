// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Yuffie Zhu (@yuffiezhu), Yue Wang (@yuew-ftnt)

// Description: Configure operator friendly name.

package fortios

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceWirelessControllerHotspot20H2QpOperatorName() *schema.Resource {
	return &schema.Resource{
		Create: resourceWirelessControllerHotspot20H2QpOperatorNameCreate,
		Read:   resourceWirelessControllerHotspot20H2QpOperatorNameRead,
		Update: resourceWirelessControllerHotspot20H2QpOperatorNameUpdate,
		Delete: resourceWirelessControllerHotspot20H2QpOperatorNameDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type: schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Optional: true,
				Computed: true,
			},
			"value_list": &schema.Schema{
				Type: schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"index": &schema.Schema{
							Type: schema.TypeInt,
							ValidateFunc: validation.IntBetween(1, 10),
							Optional: true,
							Computed: true,
						},
						"lang": &schema.Schema{
							Type: schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 3),
							Optional: true,
							Computed: true,
						},
						"value": &schema.Schema{
							Type: schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 252),
							Optional: true,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func resourceWirelessControllerHotspot20H2QpOperatorNameCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectWirelessControllerHotspot20H2QpOperatorName(d)
	if err != nil {
		return fmt.Errorf("Error creating WirelessControllerHotspot20H2QpOperatorName resource while getting object: %v", err)
	}

	o, err := c.CreateWirelessControllerHotspot20H2QpOperatorName(obj)

	if err != nil {
		return fmt.Errorf("Error creating WirelessControllerHotspot20H2QpOperatorName resource: %v", err)
	}

	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("WirelessControllerHotspot20H2QpOperatorName")
	}

	return resourceWirelessControllerHotspot20H2QpOperatorNameRead(d, m)
}

func resourceWirelessControllerHotspot20H2QpOperatorNameUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectWirelessControllerHotspot20H2QpOperatorName(d)
	if err != nil {
		return fmt.Errorf("Error updating WirelessControllerHotspot20H2QpOperatorName resource while getting object: %v", err)
	}

	o, err := c.UpdateWirelessControllerHotspot20H2QpOperatorName(obj, mkey)
	if err != nil {
		return fmt.Errorf("Error updating WirelessControllerHotspot20H2QpOperatorName resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("WirelessControllerHotspot20H2QpOperatorName")
	}

	return resourceWirelessControllerHotspot20H2QpOperatorNameRead(d, m)
}

func resourceWirelessControllerHotspot20H2QpOperatorNameDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	err := c.DeleteWirelessControllerHotspot20H2QpOperatorName(mkey)
	if err != nil {
		return fmt.Errorf("Error deleting WirelessControllerHotspot20H2QpOperatorName resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceWirelessControllerHotspot20H2QpOperatorNameRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	o, err := c.ReadWirelessControllerHotspot20H2QpOperatorName(mkey)
	if err != nil {
		return fmt.Errorf("Error reading WirelessControllerHotspot20H2QpOperatorName resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectWirelessControllerHotspot20H2QpOperatorName(d, o)
	if err != nil {
		return fmt.Errorf("Error reading WirelessControllerHotspot20H2QpOperatorName resource from API: %v", err)
	}
	return nil
}


func flattenWirelessControllerHotspot20H2QpOperatorNameName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerHotspot20H2QpOperatorNameValueList(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil
	}

	result := make([]map[string]interface{}, 0, len(l))

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})

		pre_append := "" // table

		pre_append = pre + "." + strconv.Itoa(con) + "." + "index"
		if _, ok := i["index"]; ok {
			tmp["index"] = flattenWirelessControllerHotspot20H2QpOperatorNameValueListIndex(i["index"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "lang"
		if _, ok := i["lang"]; ok {
			tmp["lang"] = flattenWirelessControllerHotspot20H2QpOperatorNameValueListLang(i["lang"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "value"
		if _, ok := i["value"]; ok {
			tmp["value"] = flattenWirelessControllerHotspot20H2QpOperatorNameValueListValue(i["value"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func flattenWirelessControllerHotspot20H2QpOperatorNameValueListIndex(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerHotspot20H2QpOperatorNameValueListLang(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerHotspot20H2QpOperatorNameValueListValue(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}


func refreshObjectWirelessControllerHotspot20H2QpOperatorName(d *schema.ResourceData, o map[string]interface{}) error {
	var err error


	if err = d.Set("name", flattenWirelessControllerHotspot20H2QpOperatorNameName(o["name"], d, "name")); err != nil {
		if !fortiAPIPatch(o["name"]) {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

    if isImportTable() {
        if err = d.Set("value_list", flattenWirelessControllerHotspot20H2QpOperatorNameValueList(o["value-list"], d, "value_list")); err != nil {
            if !fortiAPIPatch(o["value-list"]) {
                return fmt.Errorf("Error reading value_list: %v", err)
            }
        }
    } else {
        if _, ok := d.GetOk("value_list"); ok {
            if err = d.Set("value_list", flattenWirelessControllerHotspot20H2QpOperatorNameValueList(o["value-list"], d, "value_list")); err != nil {
                if !fortiAPIPatch(o["value-list"]) {
                    return fmt.Errorf("Error reading value_list: %v", err)
                }
            }
        }
    }


	return nil
}

func flattenWirelessControllerHotspot20H2QpOperatorNameFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}


func expandWirelessControllerHotspot20H2QpOperatorNameName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerHotspot20H2QpOperatorNameValueList(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	result := make([]map[string]interface{}, 0, len(l))

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})
		pre_append := ""  // table

		pre_append = pre + "." + strconv.Itoa(con) + "." + "index"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["index"], _ = expandWirelessControllerHotspot20H2QpOperatorNameValueListIndex(d, i["index"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "lang"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["lang"], _ = expandWirelessControllerHotspot20H2QpOperatorNameValueListLang(d, i["lang"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "value"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["value"], _ = expandWirelessControllerHotspot20H2QpOperatorNameValueListValue(d, i["value"], pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandWirelessControllerHotspot20H2QpOperatorNameValueListIndex(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerHotspot20H2QpOperatorNameValueListLang(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerHotspot20H2QpOperatorNameValueListValue(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}


func getObjectWirelessControllerHotspot20H2QpOperatorName(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})


	if v, ok := d.GetOk("name"); ok {
		t, err := expandWirelessControllerHotspot20H2QpOperatorNameName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("value_list"); ok {
		t, err := expandWirelessControllerHotspot20H2QpOperatorNameValueList(d, v, "value_list")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["value-list"] = t
		}
	}


	return &obj, nil
}
