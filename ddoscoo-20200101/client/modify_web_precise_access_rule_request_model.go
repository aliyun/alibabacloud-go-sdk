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
	// The domain of the website.
	//
	// > You must configure a forwarding rule for the domain. You can call the [DescribeDomains](https://help.aliyun.com/document_detail/91724.html) operation to query all domains.
	//
	// This parameter is required.
	//
	// example:
	//
	// example.aliyundoc.com
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	// The validity period of the rule, in seconds. This parameter applies only when **action*	- is set to **block**. Requests that match the rule are blocked for the specified duration. If you do not specify this parameter, the rule is permanently effective.
	//
	// example:
	//
	// 600
	Expires *int32 `json:"Expires,omitempty" xml:"Expires,omitempty"`
	// The ID of the resource group to which the Anti-DDoS Pro instance belongs. If you do not specify this parameter, the instance belongs to the default resource group.
	//
	// example:
	//
	// rg-acfm2pz25js****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The configuration of the precise access control rule, specified as a JSON-formatted string. The string has the following structure:
	//
	// - **action**: String. Required. The action to perform on matching requests. Valid values:
	//
	//   - **accept**: Allows the request.
	//
	//   - **block**: Blocks the request.
	//
	//   - **challenge**: Initiates a challenge to verify the request.
	//
	//   - **watch**: Monitors the request.
	//
	// - **name**: String. Required. The name of the rule.
	//
	// - **condition**: Map. Required. The match conditions. The structure is as follows:
	//
	//   > If you specify multiple match conditions, they are evaluated with a logical AND.
	//
	//   - **field**: String. Required. The match field.
	//
	//   - **match_method**: String. Required. The match method.
	//
	//     > For information about the valid values for **field*	- and **match_method**, see the "Additional information about request parameters" section.
	//
	//   - **content**: String. The value to match. You can specify either this parameter or `content_list`.
	//
	//   - **content_list**: String. A list of values to match. You can specify either this parameter or `content`.
	//
	//   > For information about the valid values for **match_method*	- and **content/content_list**, see the "Additional information about request parameters" section.
	//
	// - **header_name**: String. Optional. The name of the request header. This parameter applies only when **field*	- is set to **header**.
	//
	// - **case_insensitive**: Boolean. Optional. The default value is false. Valid values:
	//
	//   - **false**: case-sensitive
	//
	//   - **true**: case-insensitive
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
