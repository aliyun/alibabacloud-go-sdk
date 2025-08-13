// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListConfigRulesForTargetResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v []*ListConfigRulesForTargetResponseBodyData) *ListConfigRulesForTargetResponseBody
	GetData() []*ListConfigRulesForTargetResponseBodyData
	SetNextToken(v string) *ListConfigRulesForTargetResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListConfigRulesForTargetResponseBody
	GetRequestId() *string
}

type ListConfigRulesForTargetResponseBody struct {
	// The tag detection tasks.
	Data []*ListConfigRulesForTargetResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// Indicates whether the next query is required.
	//
	// 	- If the value of this parameter is empty (`"NextToken": ""`), all results are returned, and the next query is not required.
	//
	// 	- If the value of this parameter is not empty, the next query is required, and the value is the token used to start the next query.
	//
	// example:
	//
	// caeba0bbb2be03f84eb48b699f0a****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 7126AECD-D7AD-5073-8E88-DD2BD1FC139E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListConfigRulesForTargetResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListConfigRulesForTargetResponseBody) GoString() string {
	return s.String()
}

func (s *ListConfigRulesForTargetResponseBody) GetData() []*ListConfigRulesForTargetResponseBodyData {
	return s.Data
}

func (s *ListConfigRulesForTargetResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListConfigRulesForTargetResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListConfigRulesForTargetResponseBody) SetData(v []*ListConfigRulesForTargetResponseBodyData) *ListConfigRulesForTargetResponseBody {
	s.Data = v
	return s
}

func (s *ListConfigRulesForTargetResponseBody) SetNextToken(v string) *ListConfigRulesForTargetResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListConfigRulesForTargetResponseBody) SetRequestId(v string) *ListConfigRulesForTargetResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListConfigRulesForTargetResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListConfigRulesForTargetResponseBodyData struct {
	// The ID of the account group.
	//
	// You can use the ID to query the content of the related resource non-compliance report in Cloud Config.
	//
	// >  This parameter is returned only if you use the Tag Policy feature in multi-account mode.
	//
	// example:
	//
	// ca-efdc33dc9b37002d****
	AggregatorId *string `json:"AggregatorId,omitempty" xml:"AggregatorId,omitempty"`
	// The ID of the rule.
	//
	// example:
	//
	// cr-0lb4866180880069****
	ConfigRuleId *string `json:"ConfigRuleId,omitempty" xml:"ConfigRuleId,omitempty"`
	// The use scenario of the tag policy. Valid values:
	//
	// 	- tags: enables tags with specified tag values to be added to resources.
	//
	// 	- rg_inherit: enables resources in a resource group to automatically inherit tags from the resource group.
	//
	// example:
	//
	// tags
	PolicyType *string `json:"PolicyType,omitempty" xml:"PolicyType,omitempty"`
	// Indicates whether automatic remediation is enabled. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// false
	Remediation *bool `json:"Remediation,omitempty" xml:"Remediation,omitempty"`
	// The tag key.
	//
	// example:
	//
	// CostCenter
	TagKey *string `json:"TagKey,omitempty" xml:"TagKey,omitempty"`
	// The tag value for automatic remediation.
	//
	// example:
	//
	// Project
	TagValue *string `json:"TagValue,omitempty" xml:"TagValue,omitempty"`
	// The ID of the object.
	//
	// example:
	//
	// 134254031178****
	TargetId *string `json:"TargetId,omitempty" xml:"TargetId,omitempty"`
	// The type of the object. Valid values:
	//
	// 	- USER: the current logon account. This value is available if you use the Tag Policy feature in single-account mode.
	//
	// 	- ROOT: the Root folder in the resource directory. This value is available if you use the Tag Policy feature in multi-account mode.
	//
	// 	- FOLDER: a folder other than the Root folder in the resource directory. This value is available if you use the Tag Policy feature in multi-account mode.
	//
	// 	- ACCOUNT: a member in the resource directory. This value is available if you use the Tag Policy feature in multi-account mode.
	//
	// example:
	//
	// USER
	TargetType *string `json:"TargetType,omitempty" xml:"TargetType,omitempty"`
}

func (s ListConfigRulesForTargetResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListConfigRulesForTargetResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListConfigRulesForTargetResponseBodyData) GetAggregatorId() *string {
	return s.AggregatorId
}

func (s *ListConfigRulesForTargetResponseBodyData) GetConfigRuleId() *string {
	return s.ConfigRuleId
}

func (s *ListConfigRulesForTargetResponseBodyData) GetPolicyType() *string {
	return s.PolicyType
}

func (s *ListConfigRulesForTargetResponseBodyData) GetRemediation() *bool {
	return s.Remediation
}

func (s *ListConfigRulesForTargetResponseBodyData) GetTagKey() *string {
	return s.TagKey
}

func (s *ListConfigRulesForTargetResponseBodyData) GetTagValue() *string {
	return s.TagValue
}

func (s *ListConfigRulesForTargetResponseBodyData) GetTargetId() *string {
	return s.TargetId
}

func (s *ListConfigRulesForTargetResponseBodyData) GetTargetType() *string {
	return s.TargetType
}

func (s *ListConfigRulesForTargetResponseBodyData) SetAggregatorId(v string) *ListConfigRulesForTargetResponseBodyData {
	s.AggregatorId = &v
	return s
}

func (s *ListConfigRulesForTargetResponseBodyData) SetConfigRuleId(v string) *ListConfigRulesForTargetResponseBodyData {
	s.ConfigRuleId = &v
	return s
}

func (s *ListConfigRulesForTargetResponseBodyData) SetPolicyType(v string) *ListConfigRulesForTargetResponseBodyData {
	s.PolicyType = &v
	return s
}

func (s *ListConfigRulesForTargetResponseBodyData) SetRemediation(v bool) *ListConfigRulesForTargetResponseBodyData {
	s.Remediation = &v
	return s
}

func (s *ListConfigRulesForTargetResponseBodyData) SetTagKey(v string) *ListConfigRulesForTargetResponseBodyData {
	s.TagKey = &v
	return s
}

func (s *ListConfigRulesForTargetResponseBodyData) SetTagValue(v string) *ListConfigRulesForTargetResponseBodyData {
	s.TagValue = &v
	return s
}

func (s *ListConfigRulesForTargetResponseBodyData) SetTargetId(v string) *ListConfigRulesForTargetResponseBodyData {
	s.TargetId = &v
	return s
}

func (s *ListConfigRulesForTargetResponseBodyData) SetTargetType(v string) *ListConfigRulesForTargetResponseBodyData {
	s.TargetType = &v
	return s
}

func (s *ListConfigRulesForTargetResponseBodyData) Validate() error {
	return dara.Validate(s)
}
