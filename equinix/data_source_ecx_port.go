package equinix

import (
	"context"
	"fmt"

	"github.com/equinix/ecx-go/v2"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

var ecxPortSchemaNames = map[string]string{
	"UUID":          "uuid",
	"Name":          "name",
	"Region":        "region",
	"IBX":           "ibx",
	"MetroCode":     "metro_code",
	"Priority":      "priority",
	"Encapsulation": "encapsulation",
	"Buyout":        "buyout",
	"Bandwidth":     "bandwidth",
	"Status":        "status",
}

func dataSourceECXPort() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceECXPortRead,
		Schema: map[string]*schema.Schema{
			ecxPortSchemaNames["UUID"]: {
				Type:     schema.TypeString,
				Computed: true,
			},
			ecxPortSchemaNames["Name"]: {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.StringIsNotEmpty,
			},
			ecxPortSchemaNames["Region"]: {
				Type:     schema.TypeString,
				Computed: true,
			},
			ecxPortSchemaNames["IBX"]: {
				Type:     schema.TypeString,
				Computed: true,
			},
			ecxPortSchemaNames["MetroCode"]: {
				Type:     schema.TypeString,
				Computed: true,
			},
			ecxPortSchemaNames["Priority"]: {
				Type:     schema.TypeString,
				Computed: true,
			},
			ecxPortSchemaNames["Encapsulation"]: {
				Type:     schema.TypeString,
				Computed: true,
			},
			ecxPortSchemaNames["Buyout"]: {
				Type:     schema.TypeBool,
				Computed: true,
			},
			ecxPortSchemaNames["Bandwidth"]: {
				Type:     schema.TypeString,
				Computed: true,
			},
			ecxPortSchemaNames["Status"]: {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func dataSourceECXPortRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	conf := m.(*Config)
	var diags diag.Diagnostics
	name := d.Get(ecxPortSchemaNames["Name"]).(string)
	ports, err := conf.ecx.GetUserPorts()
	if err != nil {
		return diag.FromErr(err)
	}
	var filteredPorts []ecx.Port
	for _, port := range ports {
		if ecx.StringValue(port.Name) == name {
			filteredPorts = append(filteredPorts, port)
		}
	}
	if len(filteredPorts) < 1 {
		return diag.Errorf("profile query returned no results, please change your search criteria")
	}
	if len(filteredPorts) > 1 {
		return diag.Errorf("query returned more than one result, please try more specific search criteria")
	}
	if err := updateECXPortResource(filteredPorts[0], d); err != nil {
		return diag.FromErr(err)
	}
	return diags
}

func updateECXPortResource(port ecx.Port, d *schema.ResourceData) error {
	d.SetId(ecx.StringValue(port.UUID))
	if err := d.Set(ecxPortSchemaNames["UUID"], port.UUID); err != nil {
		return fmt.Errorf("error reading UUID: %s", err)
	}
	if err := d.Set(ecxPortSchemaNames["Region"], port.Region); err != nil {
		return fmt.Errorf("error reading Region: %s", err)
	}
	if err := d.Set(ecxPortSchemaNames["IBX"], port.IBX); err != nil {
		return fmt.Errorf("error reading IBX: %s", err)
	}
	if err := d.Set(ecxPortSchemaNames["MetroCode"], port.MetroCode); err != nil {
		return fmt.Errorf("error reading MetroCode: %s", err)
	}
	if err := d.Set(ecxPortSchemaNames["Priority"], port.Priority); err != nil {
		return fmt.Errorf("error reading Priority: %s", err)
	}
	if err := d.Set(ecxPortSchemaNames["Encapsulation"], port.Encapsulation); err != nil {
		return fmt.Errorf("error reading Encapsulation: %s", err)
	}
	if err := d.Set(ecxPortSchemaNames["Buyout"], port.Buyout); err != nil {
		return fmt.Errorf("error reading Buyout: %s", err)
	}
	if err := d.Set(ecxPortSchemaNames["Bandwidth"], port.Bandwidth); err != nil {
		return fmt.Errorf("error reading Bandwidth: %s", err)
	}
	if err := d.Set(ecxPortSchemaNames["Status"], port.Status); err != nil {
		return fmt.Errorf("error reading Status: %s", err)
	}
	return nil
}
