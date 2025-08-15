// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDeliveryHistoryJobsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPageNumber(v int32) *ListDeliveryHistoryJobsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListDeliveryHistoryJobsRequest
	GetPageSize() *int32
}

type ListDeliveryHistoryJobsRequest struct {
	// The page number.
	//
	// 	- Pages start from page 1.
	//
	// 	- Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// 	- Valid values: 1 to 100.
	//
	// 	- Default value: 20.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListDeliveryHistoryJobsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListDeliveryHistoryJobsRequest) GoString() string {
	return s.String()
}

func (s *ListDeliveryHistoryJobsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListDeliveryHistoryJobsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListDeliveryHistoryJobsRequest) SetPageNumber(v int32) *ListDeliveryHistoryJobsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListDeliveryHistoryJobsRequest) SetPageSize(v int32) *ListDeliveryHistoryJobsRequest {
	s.PageSize = &v
	return s
}

func (s *ListDeliveryHistoryJobsRequest) Validate() error {
	return dara.Validate(s)
}
