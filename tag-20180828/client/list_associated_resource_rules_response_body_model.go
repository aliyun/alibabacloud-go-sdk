// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAssociatedResourceRulesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetNextToken(v string) *ListAssociatedResourceRulesResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListAssociatedResourceRulesResponseBody
	GetRequestId() *string
	SetRules(v []*ListAssociatedResourceRulesResponseBodyRules) *ListAssociatedResourceRulesResponseBody
	GetRules() []*ListAssociatedResourceRulesResponseBodyRules
}

type ListAssociatedResourceRulesResponseBody struct {
	// Determine if there is a token for the next query based on `NextToken`. Values:
	//
	// - If `NextToken` is empty, it indicates there is no next query.
	//
	// - If `NextToken` has a value, that value is the token for the next query start.
	//
	// This parameter is required.
	//
	// example:
	//
	// caeba0bbb2be03f84eb48b699f0a****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// Request ID.
	//
	// example:
	//
	// 6E27F22C-EDA3-132E-A53F-77DE3BC2343D
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// List of associated resource tag rules.
	Rules []*ListAssociatedResourceRulesResponseBodyRules `json:"Rules,omitempty" xml:"Rules,omitempty" type:"Repeated"`
}

func (s ListAssociatedResourceRulesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListAssociatedResourceRulesResponseBody) GoString() string {
	return s.String()
}

func (s *ListAssociatedResourceRulesResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListAssociatedResourceRulesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListAssociatedResourceRulesResponseBody) GetRules() []*ListAssociatedResourceRulesResponseBodyRules {
	return s.Rules
}

func (s *ListAssociatedResourceRulesResponseBody) SetNextToken(v string) *ListAssociatedResourceRulesResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListAssociatedResourceRulesResponseBody) SetRequestId(v string) *ListAssociatedResourceRulesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListAssociatedResourceRulesResponseBody) SetRules(v []*ListAssociatedResourceRulesResponseBodyRules) *ListAssociatedResourceRulesResponseBody {
	s.Rules = v
	return s
}

func (s *ListAssociatedResourceRulesResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListAssociatedResourceRulesResponseBodyRules struct {
	ExistingStatus *string `json:"ExistingStatus,omitempty" xml:"ExistingStatus,omitempty"`
	// Setting name of the associated resource tag rule.
	//
	// example:
	//
	// rule:UpdateLoadBalancerZones-UpdateLoadBalancerAddressTypeConfig-TagAlb:Alb-LoadBalancer:Vpc-Eip
	SettingName *string `json:"SettingName,omitempty" xml:"SettingName,omitempty"`
	// Whether the associated resource tag rule is enabled. Values:
	//
	// - Enable: Enabled.
	//
	// - Disable: Disabled.
	//
	// example:
	//
	// Disable/Enable
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// List of tag keys for the associated resource tag rule.
	TagKeys []*string `json:"TagKeys,omitempty" xml:"TagKeys,omitempty" type:"Repeated"`
}

func (s ListAssociatedResourceRulesResponseBodyRules) String() string {
	return dara.Prettify(s)
}

func (s ListAssociatedResourceRulesResponseBodyRules) GoString() string {
	return s.String()
}

func (s *ListAssociatedResourceRulesResponseBodyRules) GetExistingStatus() *string {
	return s.ExistingStatus
}

func (s *ListAssociatedResourceRulesResponseBodyRules) GetSettingName() *string {
	return s.SettingName
}

func (s *ListAssociatedResourceRulesResponseBodyRules) GetStatus() *string {
	return s.Status
}

func (s *ListAssociatedResourceRulesResponseBodyRules) GetTagKeys() []*string {
	return s.TagKeys
}

func (s *ListAssociatedResourceRulesResponseBodyRules) SetExistingStatus(v string) *ListAssociatedResourceRulesResponseBodyRules {
	s.ExistingStatus = &v
	return s
}

func (s *ListAssociatedResourceRulesResponseBodyRules) SetSettingName(v string) *ListAssociatedResourceRulesResponseBodyRules {
	s.SettingName = &v
	return s
}

func (s *ListAssociatedResourceRulesResponseBodyRules) SetStatus(v string) *ListAssociatedResourceRulesResponseBodyRules {
	s.Status = &v
	return s
}

func (s *ListAssociatedResourceRulesResponseBodyRules) SetTagKeys(v []*string) *ListAssociatedResourceRulesResponseBodyRules {
	s.TagKeys = v
	return s
}

func (s *ListAssociatedResourceRulesResponseBodyRules) Validate() error {
	return dara.Validate(s)
}
