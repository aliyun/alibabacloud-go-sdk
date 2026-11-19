// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListMemoryRulesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListMemoryRulesResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *ListMemoryRulesResponseBody
	GetHttpStatusCode() *int32
	SetItems(v []*ListMemoryRulesResponseBodyItems) *ListMemoryRulesResponseBody
	GetItems() []*ListMemoryRulesResponseBodyItems
	SetMaxResults(v int32) *ListMemoryRulesResponseBody
	GetMaxResults() *int32
	SetMessage(v string) *ListMemoryRulesResponseBody
	GetMessage() *string
	SetNextToken(v string) *ListMemoryRulesResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListMemoryRulesResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListMemoryRulesResponseBody
	GetSuccess() *bool
	SetTotalCount(v int64) *ListMemoryRulesResponseBody
	GetTotalCount() *int64
}

type ListMemoryRulesResponseBody struct {
	// code
	//
	// example:
	//
	// SUCCESS
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// httpStatusCode
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"httpStatusCode,omitempty" xml:"httpStatusCode,omitempty"`
	// items
	Items []*ListMemoryRulesResponseBodyItems `json:"items,omitempty" xml:"items,omitempty" type:"Repeated"`
	// maxResults
	//
	// example:
	//
	// 100
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// message
	//
	// example:
	//
	// success
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// nextToken
	//
	// example:
	//
	// eyJvZmZzZXQiOjIwMH0=
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// requestId
	//
	// example:
	//
	// 01A00000-0000-0000-0000-000000000001
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// success
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
	// totalCount
	//
	// example:
	//
	// 20
	TotalCount *int64 `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
}

func (s ListMemoryRulesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListMemoryRulesResponseBody) GoString() string {
	return s.String()
}

func (s *ListMemoryRulesResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListMemoryRulesResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ListMemoryRulesResponseBody) GetItems() []*ListMemoryRulesResponseBodyItems {
	return s.Items
}

func (s *ListMemoryRulesResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListMemoryRulesResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListMemoryRulesResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListMemoryRulesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListMemoryRulesResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListMemoryRulesResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListMemoryRulesResponseBody) SetCode(v string) *ListMemoryRulesResponseBody {
	s.Code = &v
	return s
}

func (s *ListMemoryRulesResponseBody) SetHttpStatusCode(v int32) *ListMemoryRulesResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListMemoryRulesResponseBody) SetItems(v []*ListMemoryRulesResponseBodyItems) *ListMemoryRulesResponseBody {
	s.Items = v
	return s
}

func (s *ListMemoryRulesResponseBody) SetMaxResults(v int32) *ListMemoryRulesResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListMemoryRulesResponseBody) SetMessage(v string) *ListMemoryRulesResponseBody {
	s.Message = &v
	return s
}

func (s *ListMemoryRulesResponseBody) SetNextToken(v string) *ListMemoryRulesResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListMemoryRulesResponseBody) SetRequestId(v string) *ListMemoryRulesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListMemoryRulesResponseBody) SetSuccess(v bool) *ListMemoryRulesResponseBody {
	s.Success = &v
	return s
}

func (s *ListMemoryRulesResponseBody) SetTotalCount(v int64) *ListMemoryRulesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListMemoryRulesResponseBody) Validate() error {
	if s.Items != nil {
		for _, item := range s.Items {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListMemoryRulesResponseBodyItems struct {
	// createdAt
	//
	// This parameter is required.
	//
	// example:
	//
	// 2026-09-23T10:00:00Z
	CreatedAt *string `json:"createdAt,omitempty" xml:"createdAt,omitempty"`
	// description
	//
	// example:
	//
	// Extract long-term preferences explicitly expressed by the user
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// enabled
	//
	// This parameter is required.
	Enabled *bool `json:"enabled,omitempty" xml:"enabled,omitempty"`
	// memoryStoreId
	//
	// This parameter is required.
	//
	// example:
	//
	// ms-2f5f9c5a8d0e4c3b
	MemoryStoreId *string `json:"memoryStoreId,omitempty" xml:"memoryStoreId,omitempty"`
	// name
	//
	// This parameter is required.
	//
	// example:
	//
	// Long-term preference
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// ruleId
	//
	// This parameter is required.
	//
	// example:
	//
	// rule-2f5f9c5a8d0e4c3b
	RuleId *string `json:"ruleId,omitempty" xml:"ruleId,omitempty"`
	// strategy
	//
	// This parameter is required.
	Strategy *ListMemoryRulesResponseBodyItemsStrategy `json:"strategy,omitempty" xml:"strategy,omitempty" type:"Struct"`
	// updatedAt
	//
	// This parameter is required.
	//
	// example:
	//
	// 2026-09-23T10:00:00Z
	UpdatedAt *string `json:"updatedAt,omitempty" xml:"updatedAt,omitempty"`
}

func (s ListMemoryRulesResponseBodyItems) String() string {
	return dara.Prettify(s)
}

func (s ListMemoryRulesResponseBodyItems) GoString() string {
	return s.String()
}

func (s *ListMemoryRulesResponseBodyItems) GetCreatedAt() *string {
	return s.CreatedAt
}

func (s *ListMemoryRulesResponseBodyItems) GetDescription() *string {
	return s.Description
}

func (s *ListMemoryRulesResponseBodyItems) GetEnabled() *bool {
	return s.Enabled
}

func (s *ListMemoryRulesResponseBodyItems) GetMemoryStoreId() *string {
	return s.MemoryStoreId
}

func (s *ListMemoryRulesResponseBodyItems) GetName() *string {
	return s.Name
}

func (s *ListMemoryRulesResponseBodyItems) GetRuleId() *string {
	return s.RuleId
}

func (s *ListMemoryRulesResponseBodyItems) GetStrategy() *ListMemoryRulesResponseBodyItemsStrategy {
	return s.Strategy
}

func (s *ListMemoryRulesResponseBodyItems) GetUpdatedAt() *string {
	return s.UpdatedAt
}

func (s *ListMemoryRulesResponseBodyItems) SetCreatedAt(v string) *ListMemoryRulesResponseBodyItems {
	s.CreatedAt = &v
	return s
}

func (s *ListMemoryRulesResponseBodyItems) SetDescription(v string) *ListMemoryRulesResponseBodyItems {
	s.Description = &v
	return s
}

func (s *ListMemoryRulesResponseBodyItems) SetEnabled(v bool) *ListMemoryRulesResponseBodyItems {
	s.Enabled = &v
	return s
}

func (s *ListMemoryRulesResponseBodyItems) SetMemoryStoreId(v string) *ListMemoryRulesResponseBodyItems {
	s.MemoryStoreId = &v
	return s
}

func (s *ListMemoryRulesResponseBodyItems) SetName(v string) *ListMemoryRulesResponseBodyItems {
	s.Name = &v
	return s
}

func (s *ListMemoryRulesResponseBodyItems) SetRuleId(v string) *ListMemoryRulesResponseBodyItems {
	s.RuleId = &v
	return s
}

func (s *ListMemoryRulesResponseBodyItems) SetStrategy(v *ListMemoryRulesResponseBodyItemsStrategy) *ListMemoryRulesResponseBodyItems {
	s.Strategy = v
	return s
}

func (s *ListMemoryRulesResponseBodyItems) SetUpdatedAt(v string) *ListMemoryRulesResponseBodyItems {
	s.UpdatedAt = &v
	return s
}

func (s *ListMemoryRulesResponseBodyItems) Validate() error {
	if s.Strategy != nil {
		if err := s.Strategy.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListMemoryRulesResponseBodyItemsStrategy struct {
	// The strategy-level instruction. It takes effect together with the store-level instruction. This field takes precedence in case of conflicts.
	//
	// example:
	//
	// Retain only long-term preferences explicitly expressed by the user
	Instruction *string `json:"instruction,omitempty" xml:"instruction,omitempty"`
	// type
	//
	// This parameter is required.
	//
	// example:
	//
	// auto
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s ListMemoryRulesResponseBodyItemsStrategy) String() string {
	return dara.Prettify(s)
}

func (s ListMemoryRulesResponseBodyItemsStrategy) GoString() string {
	return s.String()
}

func (s *ListMemoryRulesResponseBodyItemsStrategy) GetInstruction() *string {
	return s.Instruction
}

func (s *ListMemoryRulesResponseBodyItemsStrategy) GetType() *string {
	return s.Type
}

func (s *ListMemoryRulesResponseBodyItemsStrategy) SetInstruction(v string) *ListMemoryRulesResponseBodyItemsStrategy {
	s.Instruction = &v
	return s
}

func (s *ListMemoryRulesResponseBodyItemsStrategy) SetType(v string) *ListMemoryRulesResponseBodyItemsStrategy {
	s.Type = &v
	return s
}

func (s *ListMemoryRulesResponseBodyItemsStrategy) Validate() error {
	return dara.Validate(s)
}
