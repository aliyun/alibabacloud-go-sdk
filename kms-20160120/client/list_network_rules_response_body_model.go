// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListNetworkRulesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetNetworkRules(v *ListNetworkRulesResponseBodyNetworkRules) *ListNetworkRulesResponseBody
	GetNetworkRules() *ListNetworkRulesResponseBodyNetworkRules
	SetPageNumber(v int32) *ListNetworkRulesResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *ListNetworkRulesResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *ListNetworkRulesResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *ListNetworkRulesResponseBody
	GetTotalCount() *int32
}

type ListNetworkRulesResponseBody struct {
	// A list of access control rules.
	NetworkRules *ListNetworkRulesResponseBodyNetworkRules `json:"NetworkRules,omitempty" xml:"NetworkRules,omitempty" type:"Struct"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// 3bf02f7a-015b-4f34-be0f-cc043fda2d33
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 1
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListNetworkRulesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListNetworkRulesResponseBody) GoString() string {
	return s.String()
}

func (s *ListNetworkRulesResponseBody) GetNetworkRules() *ListNetworkRulesResponseBodyNetworkRules {
	return s.NetworkRules
}

func (s *ListNetworkRulesResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListNetworkRulesResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListNetworkRulesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListNetworkRulesResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListNetworkRulesResponseBody) SetNetworkRules(v *ListNetworkRulesResponseBodyNetworkRules) *ListNetworkRulesResponseBody {
	s.NetworkRules = v
	return s
}

func (s *ListNetworkRulesResponseBody) SetPageNumber(v int32) *ListNetworkRulesResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListNetworkRulesResponseBody) SetPageSize(v int32) *ListNetworkRulesResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListNetworkRulesResponseBody) SetRequestId(v string) *ListNetworkRulesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListNetworkRulesResponseBody) SetTotalCount(v int32) *ListNetworkRulesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListNetworkRulesResponseBody) Validate() error {
	if s.NetworkRules != nil {
		if err := s.NetworkRules.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListNetworkRulesResponseBodyNetworkRules struct {
	NetworkRule []*ListNetworkRulesResponseBodyNetworkRulesNetworkRule `json:"NetworkRule,omitempty" xml:"NetworkRule,omitempty" type:"Repeated"`
}

func (s ListNetworkRulesResponseBodyNetworkRules) String() string {
	return dara.Prettify(s)
}

func (s ListNetworkRulesResponseBodyNetworkRules) GoString() string {
	return s.String()
}

func (s *ListNetworkRulesResponseBodyNetworkRules) GetNetworkRule() []*ListNetworkRulesResponseBodyNetworkRulesNetworkRule {
	return s.NetworkRule
}

func (s *ListNetworkRulesResponseBodyNetworkRules) SetNetworkRule(v []*ListNetworkRulesResponseBodyNetworkRulesNetworkRule) *ListNetworkRulesResponseBodyNetworkRules {
	s.NetworkRule = v
	return s
}

func (s *ListNetworkRulesResponseBodyNetworkRules) Validate() error {
	if s.NetworkRule != nil {
		for _, item := range s.NetworkRule {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListNetworkRulesResponseBodyNetworkRulesNetworkRule struct {
	// The name of the access control rule.
	//
	// example:
	//
	// networkrule_test
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The network type. The value is fixed as Private. Self-managed applications can access KMS instances only over a private virtual private cloud (VPC).
	//
	// example:
	//
	// Private
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListNetworkRulesResponseBodyNetworkRulesNetworkRule) String() string {
	return dara.Prettify(s)
}

func (s ListNetworkRulesResponseBodyNetworkRulesNetworkRule) GoString() string {
	return s.String()
}

func (s *ListNetworkRulesResponseBodyNetworkRulesNetworkRule) GetName() *string {
	return s.Name
}

func (s *ListNetworkRulesResponseBodyNetworkRulesNetworkRule) GetType() *string {
	return s.Type
}

func (s *ListNetworkRulesResponseBodyNetworkRulesNetworkRule) SetName(v string) *ListNetworkRulesResponseBodyNetworkRulesNetworkRule {
	s.Name = &v
	return s
}

func (s *ListNetworkRulesResponseBodyNetworkRulesNetworkRule) SetType(v string) *ListNetworkRulesResponseBodyNetworkRulesNetworkRule {
	s.Type = &v
	return s
}

func (s *ListNetworkRulesResponseBodyNetworkRulesNetworkRule) Validate() error {
	return dara.Validate(s)
}
