// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListMultiDimTableRecordsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBaseId(v string) *ListMultiDimTableRecordsRequest
	GetBaseId() *string
	SetFilter(v *ListMultiDimTableRecordsRequestFilter) *ListMultiDimTableRecordsRequest
	GetFilter() *ListMultiDimTableRecordsRequestFilter
	SetMaxResults(v int32) *ListMultiDimTableRecordsRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListMultiDimTableRecordsRequest
	GetNextToken() *string
	SetSheetIdOrName(v string) *ListMultiDimTableRecordsRequest
	GetSheetIdOrName() *string
	SetTenantContext(v *ListMultiDimTableRecordsRequestTenantContext) *ListMultiDimTableRecordsRequest
	GetTenantContext() *ListMultiDimTableRecordsRequestTenantContext
}

type ListMultiDimTableRecordsRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 101114
	BaseId *string                                `json:"BaseId,omitempty" xml:"BaseId,omitempty"`
	Filter *ListMultiDimTableRecordsRequestFilter `json:"Filter,omitempty" xml:"Filter,omitempty" type:"Struct"`
	// example:
	//
	// 100
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// example:
	//
	// AAAAAUUg5QSTWwHyeElt8z5z4Qo=
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// xxx
	SheetIdOrName *string                                       `json:"SheetIdOrName,omitempty" xml:"SheetIdOrName,omitempty"`
	TenantContext *ListMultiDimTableRecordsRequestTenantContext `json:"TenantContext,omitempty" xml:"TenantContext,omitempty" type:"Struct"`
}

func (s ListMultiDimTableRecordsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListMultiDimTableRecordsRequest) GoString() string {
	return s.String()
}

func (s *ListMultiDimTableRecordsRequest) GetBaseId() *string {
	return s.BaseId
}

func (s *ListMultiDimTableRecordsRequest) GetFilter() *ListMultiDimTableRecordsRequestFilter {
	return s.Filter
}

func (s *ListMultiDimTableRecordsRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListMultiDimTableRecordsRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListMultiDimTableRecordsRequest) GetSheetIdOrName() *string {
	return s.SheetIdOrName
}

func (s *ListMultiDimTableRecordsRequest) GetTenantContext() *ListMultiDimTableRecordsRequestTenantContext {
	return s.TenantContext
}

func (s *ListMultiDimTableRecordsRequest) SetBaseId(v string) *ListMultiDimTableRecordsRequest {
	s.BaseId = &v
	return s
}

func (s *ListMultiDimTableRecordsRequest) SetFilter(v *ListMultiDimTableRecordsRequestFilter) *ListMultiDimTableRecordsRequest {
	s.Filter = v
	return s
}

func (s *ListMultiDimTableRecordsRequest) SetMaxResults(v int32) *ListMultiDimTableRecordsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListMultiDimTableRecordsRequest) SetNextToken(v string) *ListMultiDimTableRecordsRequest {
	s.NextToken = &v
	return s
}

func (s *ListMultiDimTableRecordsRequest) SetSheetIdOrName(v string) *ListMultiDimTableRecordsRequest {
	s.SheetIdOrName = &v
	return s
}

func (s *ListMultiDimTableRecordsRequest) SetTenantContext(v *ListMultiDimTableRecordsRequestTenantContext) *ListMultiDimTableRecordsRequest {
	s.TenantContext = v
	return s
}

func (s *ListMultiDimTableRecordsRequest) Validate() error {
	if s.Filter != nil {
		if err := s.Filter.Validate(); err != nil {
			return err
		}
	}
	if s.TenantContext != nil {
		if err := s.TenantContext.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListMultiDimTableRecordsRequestFilter struct {
	// example:
	//
	// and
	Combination *string                                            `json:"Combination,omitempty" xml:"Combination,omitempty"`
	Conditions  []*ListMultiDimTableRecordsRequestFilterConditions `json:"Conditions,omitempty" xml:"Conditions,omitempty" type:"Repeated"`
}

func (s ListMultiDimTableRecordsRequestFilter) String() string {
	return dara.Prettify(s)
}

func (s ListMultiDimTableRecordsRequestFilter) GoString() string {
	return s.String()
}

func (s *ListMultiDimTableRecordsRequestFilter) GetCombination() *string {
	return s.Combination
}

func (s *ListMultiDimTableRecordsRequestFilter) GetConditions() []*ListMultiDimTableRecordsRequestFilterConditions {
	return s.Conditions
}

func (s *ListMultiDimTableRecordsRequestFilter) SetCombination(v string) *ListMultiDimTableRecordsRequestFilter {
	s.Combination = &v
	return s
}

func (s *ListMultiDimTableRecordsRequestFilter) SetConditions(v []*ListMultiDimTableRecordsRequestFilterConditions) *ListMultiDimTableRecordsRequestFilter {
	s.Conditions = v
	return s
}

func (s *ListMultiDimTableRecordsRequestFilter) Validate() error {
	if s.Conditions != nil {
		for _, item := range s.Conditions {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListMultiDimTableRecordsRequestFilterConditions struct {
	// example:
	//
	// Sandbox
	Field *string `json:"Field,omitempty" xml:"Field,omitempty"`
	// example:
	//
	// equal
	Operator *string       `json:"Operator,omitempty" xml:"Operator,omitempty"`
	Value    []interface{} `json:"Value,omitempty" xml:"Value,omitempty" type:"Repeated"`
}

func (s ListMultiDimTableRecordsRequestFilterConditions) String() string {
	return dara.Prettify(s)
}

func (s ListMultiDimTableRecordsRequestFilterConditions) GoString() string {
	return s.String()
}

func (s *ListMultiDimTableRecordsRequestFilterConditions) GetField() *string {
	return s.Field
}

func (s *ListMultiDimTableRecordsRequestFilterConditions) GetOperator() *string {
	return s.Operator
}

func (s *ListMultiDimTableRecordsRequestFilterConditions) GetValue() []interface{} {
	return s.Value
}

func (s *ListMultiDimTableRecordsRequestFilterConditions) SetField(v string) *ListMultiDimTableRecordsRequestFilterConditions {
	s.Field = &v
	return s
}

func (s *ListMultiDimTableRecordsRequestFilterConditions) SetOperator(v string) *ListMultiDimTableRecordsRequestFilterConditions {
	s.Operator = &v
	return s
}

func (s *ListMultiDimTableRecordsRequestFilterConditions) SetValue(v []interface{}) *ListMultiDimTableRecordsRequestFilterConditions {
	s.Value = v
	return s
}

func (s *ListMultiDimTableRecordsRequestFilterConditions) Validate() error {
	return dara.Validate(s)
}

type ListMultiDimTableRecordsRequestTenantContext struct {
	// example:
	//
	// xxxxxx
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s ListMultiDimTableRecordsRequestTenantContext) String() string {
	return dara.Prettify(s)
}

func (s ListMultiDimTableRecordsRequestTenantContext) GoString() string {
	return s.String()
}

func (s *ListMultiDimTableRecordsRequestTenantContext) GetTenantId() *string {
	return s.TenantId
}

func (s *ListMultiDimTableRecordsRequestTenantContext) SetTenantId(v string) *ListMultiDimTableRecordsRequestTenantContext {
	s.TenantId = &v
	return s
}

func (s *ListMultiDimTableRecordsRequestTenantContext) Validate() error {
	return dara.Validate(s)
}
