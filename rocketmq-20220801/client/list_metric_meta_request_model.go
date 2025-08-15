// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListMetricMetaRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPageNumber(v int32) *ListMetricMetaRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListMetricMetaRequest
	GetPageSize() *int32
}

type ListMetricMetaRequest struct {
	// Page number, indicating which page of the results to return.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// Page size, indicating the maximum number of results per page.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
}

func (s ListMetricMetaRequest) String() string {
	return dara.Prettify(s)
}

func (s ListMetricMetaRequest) GoString() string {
	return s.String()
}

func (s *ListMetricMetaRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListMetricMetaRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListMetricMetaRequest) SetPageNumber(v int32) *ListMetricMetaRequest {
	s.PageNumber = &v
	return s
}

func (s *ListMetricMetaRequest) SetPageSize(v int32) *ListMetricMetaRequest {
	s.PageSize = &v
	return s
}

func (s *ListMetricMetaRequest) Validate() error {
	return dara.Validate(s)
}
