// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListContactFlowsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *ListContactFlowsRequest
	GetInstanceId() *string
	SetOrderByField(v string) *ListContactFlowsRequest
	GetOrderByField() *string
	SetPageNumber(v int32) *ListContactFlowsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListContactFlowsRequest
	GetPageSize() *int32
	SetSearchPattern(v string) *ListContactFlowsRequest
	GetSearchPattern() *string
	SetSortOrder(v string) *ListContactFlowsRequest
	GetSortOrder() *string
	SetType(v string) *ListContactFlowsRequest
	GetType() *string
}

type ListContactFlowsRequest struct {
	// The ID of the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// ccc-test
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The field by which to sort the results. Default value: updatedTime.
	//
	// example:
	//
	// updatedTime
	OrderByField *string `json:"OrderByField,omitempty" xml:"OrderByField,omitempty"`
	// The page number. The value must be an integer that is greater than 0. Default value: 1.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Valid values: 1 to 100. Default value: 10.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// A keyword to search for in the names or descriptions of IVR flows.
	SearchPattern *string `json:"SearchPattern,omitempty" xml:"SearchPattern,omitempty"`
	// The sort order. Valid values:
	//
	// - `ASC`: The results are sorted in ascending order.
	//
	// - `DESC`: The results are sorted in descending order. This is the default value.
	//
	// example:
	//
	// DESC
	SortOrder *string `json:"SortOrder,omitempty" xml:"SortOrder,omitempty"`
	// Specifies the type of IVR flow to list. If you omit this parameter, IVR flows of all types are listed.
	//
	// example:
	//
	// MAIN_FLOW
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListContactFlowsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListContactFlowsRequest) GoString() string {
	return s.String()
}

func (s *ListContactFlowsRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListContactFlowsRequest) GetOrderByField() *string {
	return s.OrderByField
}

func (s *ListContactFlowsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListContactFlowsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListContactFlowsRequest) GetSearchPattern() *string {
	return s.SearchPattern
}

func (s *ListContactFlowsRequest) GetSortOrder() *string {
	return s.SortOrder
}

func (s *ListContactFlowsRequest) GetType() *string {
	return s.Type
}

func (s *ListContactFlowsRequest) SetInstanceId(v string) *ListContactFlowsRequest {
	s.InstanceId = &v
	return s
}

func (s *ListContactFlowsRequest) SetOrderByField(v string) *ListContactFlowsRequest {
	s.OrderByField = &v
	return s
}

func (s *ListContactFlowsRequest) SetPageNumber(v int32) *ListContactFlowsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListContactFlowsRequest) SetPageSize(v int32) *ListContactFlowsRequest {
	s.PageSize = &v
	return s
}

func (s *ListContactFlowsRequest) SetSearchPattern(v string) *ListContactFlowsRequest {
	s.SearchPattern = &v
	return s
}

func (s *ListContactFlowsRequest) SetSortOrder(v string) *ListContactFlowsRequest {
	s.SortOrder = &v
	return s
}

func (s *ListContactFlowsRequest) SetType(v string) *ListContactFlowsRequest {
	s.Type = &v
	return s
}

func (s *ListContactFlowsRequest) Validate() error {
	return dara.Validate(s)
}
