package provider

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func New() *schema.Provider {
	return &schema.Provider{
		ResourcesMap: map[string]*schema.Resource{
			"hello_resource": getResource(),
		},
	}
}

func getResource() *schema.Resource {
	return &schema.Resource{
		CreateContext: getCreateContext,
		ReadContext:   getReadContext,
		UpdateContext: getUpdateContext,
		DeleteContext: getDeleteContext,
		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"message": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

// This function is called when the resource is first created, i.e., during terraform apply.
func getCreateContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	name := d.Get("name").(string)
	if name == "" {
		return diag.Errorf("Name is missing")
	}
	id := fmt.Sprintf("id-%s", name)
	d.SetId(id)
	message := fmt.Sprintf("Hi %s, %s id created!", name, id)
	d.Set("message", message)
	return nil
}

// This function is called when Terraform needs to refresh the current state of the resource — for example, during terraform refresh or terraform plan.
func getReadContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	name := d.Get("name").(string)
	id := fmt.Sprintf("id-%s", name)
	var message string
	if id != d.Id() {
		d.SetId("")
		message = fmt.Sprintf("Hi %s, %s id created!", name, id)
	} else {
		message = fmt.Sprintf("Hi %s, %s id exist!", name, id)
	}
	d.Set("message", message)
	return nil
}

func getUpdateContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	if d.HasChange("name") {
		name := d.Get("name").(string)
		id := fmt.Sprintf("id-%s", name)
		d.SetId(id)
		message := fmt.Sprintf("Hi %s, %s id updated!", name, id)
		d.Set("message", message)
	}
	return nil
}

func getDeleteContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	message := fmt.Sprintf("Hi %s, %s id deleted!", d.Get("name").(string), d.Id())
	d.SetId("")
	d.Set("message", message)
	return nil
}
