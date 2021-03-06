package qingcloud

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/magicshui/qingcloud-go/loadbalancer"
)

func resourceQingcloudLoadbalancerPloicyRule() *schema.Resource {
	return &schema.Resource{
		Create: resourceQingcloudLoadbalancerPloicyRuleCreate,
		Read:   resourceQingcloudLoadbalancerPloicyRuleRead,
		Update: resourceQingcloudLoadbalancerPloicyRuleUpdate,
		Delete: resourceQingcloudLoadbalancerPloicyRuleDelete,
		Schema: map[string]*schema.Schema{
			"policy": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "要添加规则的转发策略ID",
			},
			"name": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "规则名称",
			},
			"type": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				Description: "规则匹配类型：按域名”domain” 还是按URL “url”	",
				ValidateFunc: withinArrayString("domain", "url"),
			},
			"val": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "匹配规则，支持正则表达式",
			},
		},
	}
}

func resourceQingcloudLoadbalancerPloicyRuleCreate(d *schema.ResourceData, meta interface{}) error {
	clt := meta.(*QingCloudClient).loadbalancer
	params := loadbalancer.AddLoadBalancerPolicyRulesRequest{}
	params.LoadbalancerPolicy.Set(d.Get("policy").(string))
	params.RulesNLoadbalancerPolicyRuleName.Add(d.Get("name").(string))
	params.RulesNRuleType.Add(d.Get("type").(string))
	params.RulesNVal.Add(d.Get("val").(string))
	resp, err := clt.AddLoadBalancerPolicyRules(params)
	if err != nil {
		return err
	}
	d.SetId(resp.LoadbalancerPoliciyRules[0])
	return nil
}

func resourceQingcloudLoadbalancerPloicyRuleRead(d *schema.ResourceData, meta interface{}) error {
	clt := meta.(*QingCloudClient).loadbalancer
	params := loadbalancer.DescribeLoadBalancerPolicyRulesRequest{}
	params.LoadbalancerPolicyRulesN.Add(d.Id())
	resp, err := clt.DescribeLoadBalancerPolicyRules(params)
	if err != nil {
		return err
	}
	lp := resp.LoadbalancerPoliciyRule[0]
	d.Set("type", lp.RuleType)
	d.Set("val", lp.Val)
	d.Set("policy", lp.LoadbalancerPolicyID)
	return nil
}

func resourceQingcloudLoadbalancerPloicyRuleUpdate(d *schema.ResourceData, meta interface{}) error {
	return nil
}

func resourceQingcloudLoadbalancerPloicyRuleDelete(d *schema.ResourceData, meta interface{}) error {
	clt := meta.(*QingCloudClient).loadbalancer
	params := loadbalancer.DeleteLoadBalancerPolicyRulesRequest{}
	params.LoadbalancerPolicyRulesN.Add(d.Id())
	_, err := clt.DeleteLoadBalancerPolicyRules(params)
	if err != nil {
		return err
	}
	d.SetId("")
	return nil
}
