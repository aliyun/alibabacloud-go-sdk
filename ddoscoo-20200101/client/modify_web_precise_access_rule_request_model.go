// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyWebPreciseAccessRuleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomain(v string) *ModifyWebPreciseAccessRuleRequest
	GetDomain() *string
	SetExpires(v int32) *ModifyWebPreciseAccessRuleRequest
	GetExpires() *int32
	SetResourceGroupId(v string) *ModifyWebPreciseAccessRuleRequest
	GetResourceGroupId() *string
	SetRules(v string) *ModifyWebPreciseAccessRuleRequest
	GetRules() *string
}

type ModifyWebPreciseAccessRuleRequest struct {
	// The domain name of the website.
	//
	// > A forwarding rule must be configured for the domain name. You can call the [DescribeDomains](https://help.aliyun.com/document_detail/91724.html) operation to query all domain names.
	//
	// This parameter is required.
	//
	// example:
	//
	// example.aliyundoc.com
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	// The validity period of the rule. Unit: seconds. This parameter takes effect only when **action*	- of a rule is **block**. Access requests that match the rule are blocked within the specified validity period of the rule. If you do not specify this parameter, this rule takes effect all the time.
	//
	// example:
	//
	// 600
	Expires *int32 `json:"Expires,omitempty" xml:"Expires,omitempty"`
	// The ID of the resource group to which the instance belongs in Resource Management. This parameter is empty by default, which indicates that the instance belongs to the default resource group.
	//
	// example:
	//
	// rg-acfm2pz25js****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The settings of the accurate access control rule. This parameter is a JSON string. The following list describes the fields in the value of the parameter:
	//
	// 	- **action**: the action that is performed if the rule is matched. This field is required and must be of the string type. Valid values:
	//
	//     	- **accept**: allows the requests that match the rule.
	//
	//     	- **block**: blocks the requests that match the rule.
	//
	//     	- **challenge**: implements a CAPTCHA for the requests that match the rule.
	//
	// 	- **name**: the name of the rule. This field is required and must be of the string type.
	//
	// 	- **condition**: the match conditions. This field is required and must be of the map type. A match condition contains the following parameters.
	//
	//     **
	//
	//     **Note**The AND logical operator is used to define the relationship among multiple match conditions.
	//
	//     	- **field**: the match field. This parameter is required and must be of the string type.
	//
	//     	- **match_method**: the logical relation. This parameter is required and must be of the string type.
	//
	//         **
	//
	//         **Note**For information about the mappings between the **field*	- and **match_method*	- parameters, see the Mappings between the field and match_method parameters table in this topic.
	//
	//     	- **content**: the match content. This parameter is required and must be of the string type.
	//
	// 	- **header_name**: the HTTP header. This parameter is optional and must be of the string type. This parameter takes effect only when **field*	- is **header**.
	//
	// This parameter is required.
	//
	// example:
	//
	// [{"action":"block","name":"testrule","condition":[{"field":"uri","match_method":"contain","content":"/test/123"}]}]
	Rules *string `json:"Rules,omitempty" xml:"Rules,omitempty"`
}

func (s ModifyWebPreciseAccessRuleRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyWebPreciseAccessRuleRequest) GoString() string {
	return s.String()
}

func (s *ModifyWebPreciseAccessRuleRequest) GetDomain() *string {
	return s.Domain
}

func (s *ModifyWebPreciseAccessRuleRequest) GetExpires() *int32 {
	return s.Expires
}

func (s *ModifyWebPreciseAccessRuleRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *ModifyWebPreciseAccessRuleRequest) GetRules() *string {
	return s.Rules
}

func (s *ModifyWebPreciseAccessRuleRequest) SetDomain(v string) *ModifyWebPreciseAccessRuleRequest {
	s.Domain = &v
	return s
}

func (s *ModifyWebPreciseAccessRuleRequest) SetExpires(v int32) *ModifyWebPreciseAccessRuleRequest {
	s.Expires = &v
	return s
}

func (s *ModifyWebPreciseAccessRuleRequest) SetResourceGroupId(v string) *ModifyWebPreciseAccessRuleRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *ModifyWebPreciseAccessRuleRequest) SetRules(v string) *ModifyWebPreciseAccessRuleRequest {
	s.Rules = &v
	return s
}

func (s *ModifyWebPreciseAccessRuleRequest) Validate() error {
	return dara.Validate(s)
}
