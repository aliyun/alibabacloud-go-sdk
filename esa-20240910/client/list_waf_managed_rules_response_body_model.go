// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListWafManagedRulesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPageNumber(v int32) *ListWafManagedRulesResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *ListWafManagedRulesResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *ListWafManagedRulesResponseBody
	GetRequestId() *string
	SetRules(v []*ListWafManagedRulesResponseBodyRules) *ListWafManagedRulesResponseBody
	GetRules() []*ListWafManagedRulesResponseBodyRules
	SetTotalCount(v int64) *ListWafManagedRulesResponseBody
	GetTotalCount() *int64
}

type ListWafManagedRulesResponseBody struct {
	// Page number.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// Page size.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// Request ID.
	//
	// example:
	//
	// 36af3fcc-43d0-441c-86b1-428951dc8225
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// List of managed rules.
	Rules []*ListWafManagedRulesResponseBodyRules `json:"Rules,omitempty" xml:"Rules,omitempty" type:"Repeated"`
	// Total number of rules after filtering.
	//
	// example:
	//
	// 20
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListWafManagedRulesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListWafManagedRulesResponseBody) GoString() string {
	return s.String()
}

func (s *ListWafManagedRulesResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListWafManagedRulesResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListWafManagedRulesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListWafManagedRulesResponseBody) GetRules() []*ListWafManagedRulesResponseBodyRules {
	return s.Rules
}

func (s *ListWafManagedRulesResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListWafManagedRulesResponseBody) SetPageNumber(v int32) *ListWafManagedRulesResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListWafManagedRulesResponseBody) SetPageSize(v int32) *ListWafManagedRulesResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListWafManagedRulesResponseBody) SetRequestId(v string) *ListWafManagedRulesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListWafManagedRulesResponseBody) SetRules(v []*ListWafManagedRulesResponseBodyRules) *ListWafManagedRulesResponseBody {
	s.Rules = v
	return s
}

func (s *ListWafManagedRulesResponseBody) SetTotalCount(v int64) *ListWafManagedRulesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListWafManagedRulesResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListWafManagedRulesResponseBodyRules struct {
	// Protection action of the managed rule.
	//
	// example:
	//
	// deny
	Action *string `json:"Action,omitempty" xml:"Action,omitempty"`
	// ID of the managed rule.
	//
	// example:
	//
	// 100001
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// Name of the managed rule.
	//
	// example:
	//
	// SQL injection
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// Protection level of the managed rule.
	//
	// example:
	//
	// 1
	ProtectionLevel *int32 `json:"ProtectionLevel,omitempty" xml:"ProtectionLevel,omitempty"`
	// Protection status of the managed rule.
	//
	// example:
	//
	// on
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ListWafManagedRulesResponseBodyRules) String() string {
	return dara.Prettify(s)
}

func (s ListWafManagedRulesResponseBodyRules) GoString() string {
	return s.String()
}

func (s *ListWafManagedRulesResponseBodyRules) GetAction() *string {
	return s.Action
}

func (s *ListWafManagedRulesResponseBodyRules) GetId() *int64 {
	return s.Id
}

func (s *ListWafManagedRulesResponseBodyRules) GetName() *string {
	return s.Name
}

func (s *ListWafManagedRulesResponseBodyRules) GetProtectionLevel() *int32 {
	return s.ProtectionLevel
}

func (s *ListWafManagedRulesResponseBodyRules) GetStatus() *string {
	return s.Status
}

func (s *ListWafManagedRulesResponseBodyRules) SetAction(v string) *ListWafManagedRulesResponseBodyRules {
	s.Action = &v
	return s
}

func (s *ListWafManagedRulesResponseBodyRules) SetId(v int64) *ListWafManagedRulesResponseBodyRules {
	s.Id = &v
	return s
}

func (s *ListWafManagedRulesResponseBodyRules) SetName(v string) *ListWafManagedRulesResponseBodyRules {
	s.Name = &v
	return s
}

func (s *ListWafManagedRulesResponseBodyRules) SetProtectionLevel(v int32) *ListWafManagedRulesResponseBodyRules {
	s.ProtectionLevel = &v
	return s
}

func (s *ListWafManagedRulesResponseBodyRules) SetStatus(v string) *ListWafManagedRulesResponseBodyRules {
	s.Status = &v
	return s
}

func (s *ListWafManagedRulesResponseBodyRules) Validate() error {
	return dara.Validate(s)
}
