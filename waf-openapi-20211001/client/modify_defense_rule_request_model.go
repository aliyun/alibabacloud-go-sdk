// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDefenseRuleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDefenseScene(v string) *ModifyDefenseRuleRequest
	GetDefenseScene() *string
	SetDefenseType(v string) *ModifyDefenseRuleRequest
	GetDefenseType() *string
	SetInstanceId(v string) *ModifyDefenseRuleRequest
	GetInstanceId() *string
	SetRegionId(v string) *ModifyDefenseRuleRequest
	GetRegionId() *string
	SetResource(v string) *ModifyDefenseRuleRequest
	GetResource() *string
	SetResourceManagerResourceGroupId(v string) *ModifyDefenseRuleRequest
	GetResourceManagerResourceGroupId() *string
	SetRules(v string) *ModifyDefenseRuleRequest
	GetRules() *string
	SetTemplateId(v int64) *ModifyDefenseRuleRequest
	GetTemplateId() *int64
}

type ModifyDefenseRuleRequest struct {
	// The protection scenario to modify. For more information, see the **DefenseScene*	- parameter in [CreateDefenseRule](https://help.aliyun.com/document_detail/461421.html).
	//
	// example:
	//
	// waf_group
	DefenseScene *string `json:"DefenseScene,omitempty" xml:"DefenseScene,omitempty"`
	// The type of the protection rule.
	//
	// example:
	//
	// template
	DefenseType *string `json:"DefenseType,omitempty" xml:"DefenseType,omitempty"`
	// Instance ID of the WAF instance.
	//
	// > You can call the [DescribeInstance](https://help.aliyun.com/document_detail/433756.html) operation to query instance ID of the current WAF instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// waf_cdnsdf3****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The region where the WAF instance is deployed. Valid values:
	//
	// - **cn-hangzhou**: the Chinese mainland.
	//
	// - **ap-southeast-1**: outside the Chinese mainland.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The protected object associated with the rule to modify.
	//
	// > This parameter is required only when **DefenseType*	- is set to **resource**.
	//
	// example:
	//
	// rencs***-waf
	Resource *string `json:"Resource,omitempty" xml:"Resource,omitempty"`
	// The ID of the Alibaba Cloud resource group.
	//
	// example:
	//
	// rg-acfm***q
	ResourceManagerResourceGroupId *string `json:"ResourceManagerResourceGroupId,omitempty" xml:"ResourceManagerResourceGroupId,omitempty"`
	// The details of the protection rule. The value is a string that is converted from a JSON object constructed by a series of parameters. When you configure this parameter, you must specify the rule ID and the protection rule configuration to modify. The following parameters are included:
	//
	// - **id**: Long | Required | The rule ID.
	//
	// - Protection rule configuration: Same as the **Rules*	- parameter of the **CreateDefenseRule*	- operation. For more information, see the **protection rule parameter description*	- in [CreateDefenseRule](https://help.aliyun.com/document_detail/461421.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// [
	//
	//       {
	//
	//             "id": 2344,
	//
	//             "policyId": 1012,
	//
	//             "action": "block"
	//
	//       }
	//
	// ]
	Rules *string `json:"Rules,omitempty" xml:"Rules,omitempty"`
	// The ID of the protection rule template.
	//
	// > This parameter is required only when **DefenseType*	- is set to **template**.
	//
	// example:
	//
	// 5325
	TemplateId *int64 `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
}

func (s ModifyDefenseRuleRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyDefenseRuleRequest) GoString() string {
	return s.String()
}

func (s *ModifyDefenseRuleRequest) GetDefenseScene() *string {
	return s.DefenseScene
}

func (s *ModifyDefenseRuleRequest) GetDefenseType() *string {
	return s.DefenseType
}

func (s *ModifyDefenseRuleRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ModifyDefenseRuleRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifyDefenseRuleRequest) GetResource() *string {
	return s.Resource
}

func (s *ModifyDefenseRuleRequest) GetResourceManagerResourceGroupId() *string {
	return s.ResourceManagerResourceGroupId
}

func (s *ModifyDefenseRuleRequest) GetRules() *string {
	return s.Rules
}

func (s *ModifyDefenseRuleRequest) GetTemplateId() *int64 {
	return s.TemplateId
}

func (s *ModifyDefenseRuleRequest) SetDefenseScene(v string) *ModifyDefenseRuleRequest {
	s.DefenseScene = &v
	return s
}

func (s *ModifyDefenseRuleRequest) SetDefenseType(v string) *ModifyDefenseRuleRequest {
	s.DefenseType = &v
	return s
}

func (s *ModifyDefenseRuleRequest) SetInstanceId(v string) *ModifyDefenseRuleRequest {
	s.InstanceId = &v
	return s
}

func (s *ModifyDefenseRuleRequest) SetRegionId(v string) *ModifyDefenseRuleRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyDefenseRuleRequest) SetResource(v string) *ModifyDefenseRuleRequest {
	s.Resource = &v
	return s
}

func (s *ModifyDefenseRuleRequest) SetResourceManagerResourceGroupId(v string) *ModifyDefenseRuleRequest {
	s.ResourceManagerResourceGroupId = &v
	return s
}

func (s *ModifyDefenseRuleRequest) SetRules(v string) *ModifyDefenseRuleRequest {
	s.Rules = &v
	return s
}

func (s *ModifyDefenseRuleRequest) SetTemplateId(v int64) *ModifyDefenseRuleRequest {
	s.TemplateId = &v
	return s
}

func (s *ModifyDefenseRuleRequest) Validate() error {
	return dara.Validate(s)
}
