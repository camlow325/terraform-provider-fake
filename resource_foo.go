package main

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceFoo() *schema.Resource {
	return &schema.Resource{
		Create: resourceFooCreate,
		Read:   resourceFooRead,
		Update: resourceFooUpdate,
		Delete: resourceFooDelete,

		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"parameters": &schema.Schema{
				Type:     schema.TypeMap,
				Required: true,
			},
		},
	}
}

func resourceFooCreate(d *schema.ResourceData, m interface{}) error {
	name := d.Get("name").(string)
	d.SetId(name)
	return resourceFooRead(d, m)
}

func resourceFooRead(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceFooUpdate(d *schema.ResourceData, m interface{}) error {
	return resourceFooRead(d, m)
}

func resourceFooDelete(d *schema.ResourceData, m interface{}) error {
	return nil
}
