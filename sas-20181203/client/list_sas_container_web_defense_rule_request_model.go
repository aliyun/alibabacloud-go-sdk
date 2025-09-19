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
	// The search conditions for assets. The value of this parameter is in the JSON format and contains the following fields:
	//
	// 	- **name**: the name of the search condition.
	//
	// 	- **value**: the value of the search condition.
	//
	// 	- **logicalExp**: the logical relation for multiple search conditions. Valid values:
	//
	//     	- **OR**: The search conditions use a logical **OR**.
	//
	//     	- **AND**: The search conditions use a logical **AND**.
	//
	// example:
	//
	// [{\\"name\\":\\"ruleName\\",\\"value\\":\\"test-1818\\",\\"logicalExp\\":\\"AND\\"}]
	Criteria *string `json:"Criteria,omitempty" xml:"Criteria,omitempty"`
	// The page number. Pages start from page 1. Default value: 1.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The logical relation for multiple search conditions. Valid values:
	//
	// 	- **OR**: The search conditions use a logical **OR**.
	//
	// 	- **AND**: The search conditions use a logical **AND**.
	//
	// example:
	//
	// OR
	LogicalExp *string `json:"LogicalExp,omitempty" xml:"LogicalExp,omitempty"`
	// The number of entries per page. Default value: **20**.
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
