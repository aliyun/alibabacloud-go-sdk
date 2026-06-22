// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSasContainerWebDefenseRuleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCriteria(v string) *ListSasContainerWebDefenseRuleRequest
	GetCriteria() *string
	SetCurrentPage(v int32) *ListSasContainerWebDefenseRuleRequest
	GetCurrentPage() *int32
	SetLogicalExp(v string) *ListSasContainerWebDefenseRuleRequest
	GetLogicalExp() *string
	SetPageSize(v int32) *ListSasContainerWebDefenseRuleRequest
	GetPageSize() *int32
}

type ListSasContainerWebDefenseRuleRequest struct {
	// The conditions for searching assets. This parameter is in JSON format and contains the following fields:
	//
	// - **name**: The search item.
	//
	// - **value**: The value of the search item.
	//
	// - **logicalExp**: The logical relationship among multiple search item values. Valid values:
	//
	//     - **OR**: The search item values are evaluated by using the OR operator.
	//
	//     - **AND**: The search item values are evaluated by using the AND operator.
	//
	// example:
	//
	// [{\\"name\\":\\"ruleName\\",\\"value\\":\\"test-1818\\",\\"logicalExp\\":\\"AND\\"}]
	Criteria *string `json:"Criteria,omitempty" xml:"Criteria,omitempty"`
	// The page number of the current page in a paged query. Minimum value: 1. Default value: 1.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The logical relationship among multiple search conditions. Valid values:
	//
	// - **OR**: The search conditions are evaluated by using the OR operator.
	//
	// - **AND**: The search conditions are evaluated by using the AND operator.
	//
	// example:
	//
	// OR
	LogicalExp *string `json:"LogicalExp,omitempty" xml:"LogicalExp,omitempty"`
	// The number of entries per page in a paged query. Default value: **20**, which indicates that 20 entries are displayed per page.
	//
	// example:
	//
	// 50
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListSasContainerWebDefenseRuleRequest) String() string {
	return dara.Prettify(s)
}

func (s ListSasContainerWebDefenseRuleRequest) GoString() string {
	return s.String()
}

func (s *ListSasContainerWebDefenseRuleRequest) GetCriteria() *string {
	return s.Criteria
}

func (s *ListSasContainerWebDefenseRuleRequest) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *ListSasContainerWebDefenseRuleRequest) GetLogicalExp() *string {
	return s.LogicalExp
}

func (s *ListSasContainerWebDefenseRuleRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListSasContainerWebDefenseRuleRequest) SetCriteria(v string) *ListSasContainerWebDefenseRuleRequest {
	s.Criteria = &v
	return s
}

func (s *ListSasContainerWebDefenseRuleRequest) SetCurrentPage(v int32) *ListSasContainerWebDefenseRuleRequest {
	s.CurrentPage = &v
	return s
}

func (s *ListSasContainerWebDefenseRuleRequest) SetLogicalExp(v string) *ListSasContainerWebDefenseRuleRequest {
	s.LogicalExp = &v
	return s
}

func (s *ListSasContainerWebDefenseRuleRequest) SetPageSize(v int32) *ListSasContainerWebDefenseRuleRequest {
	s.PageSize = &v
	return s
}

func (s *ListSasContainerWebDefenseRuleRequest) Validate() error {
	return dara.Validate(s)
}
