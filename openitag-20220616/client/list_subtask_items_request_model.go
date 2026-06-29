// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSubtaskItemsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPageNumber(v int32) *ListSubtaskItemsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListSubtaskItemsRequest
	GetPageSize() *int32
}

type ListSubtaskItemsRequest struct {
	// The page number of the annotated data for task packages. The starting value is 1, and the default value is 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of annotated data entries per page to display in a paged query. Default value is 20.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListSubtaskItemsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListSubtaskItemsRequest) GoString() string {
	return s.String()
}

func (s *ListSubtaskItemsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListSubtaskItemsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListSubtaskItemsRequest) SetPageNumber(v int32) *ListSubtaskItemsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListSubtaskItemsRequest) SetPageSize(v int32) *ListSubtaskItemsRequest {
	s.PageSize = &v
	return s
}

func (s *ListSubtaskItemsRequest) Validate() error {
	return dara.Validate(s)
}
