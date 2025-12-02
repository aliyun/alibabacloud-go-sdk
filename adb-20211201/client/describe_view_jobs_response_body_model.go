// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeViewJobsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *DescribeViewJobsResponseBodyData) *DescribeViewJobsResponseBody
	GetData() *DescribeViewJobsResponseBodyData
	SetPageNumber(v int64) *DescribeViewJobsResponseBody
	GetPageNumber() *int64
	SetPageSize(v int64) *DescribeViewJobsResponseBody
	GetPageSize() *int64
	SetRequestId(v string) *DescribeViewJobsResponseBody
	GetRequestId() *string
	SetTotalCount(v int64) *DescribeViewJobsResponseBody
	GetTotalCount() *int64
}

type DescribeViewJobsResponseBody struct {
	// The returned data.
	Data *DescribeViewJobsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries returned per page.
	//
	// example:
	//
	// 30
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 2A318054-6815-528A-AA94-8AC921******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries that are returned.
	//
	// example:
	//
	// 44
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeViewJobsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeViewJobsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeViewJobsResponseBody) GetData() *DescribeViewJobsResponseBodyData {
	return s.Data
}

func (s *DescribeViewJobsResponseBody) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *DescribeViewJobsResponseBody) GetPageSize() *int64 {
	return s.PageSize
}

func (s *DescribeViewJobsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeViewJobsResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *DescribeViewJobsResponseBody) SetData(v *DescribeViewJobsResponseBodyData) *DescribeViewJobsResponseBody {
	s.Data = v
	return s
}

func (s *DescribeViewJobsResponseBody) SetPageNumber(v int64) *DescribeViewJobsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeViewJobsResponseBody) SetPageSize(v int64) *DescribeViewJobsResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeViewJobsResponseBody) SetRequestId(v string) *DescribeViewJobsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeViewJobsResponseBody) SetTotalCount(v int64) *DescribeViewJobsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeViewJobsResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeViewJobsResponseBodyData struct {
	// The refresh tasks.
	MvRefreshJobModels []*OpenStructRefreshJobModel `json:"MvRefreshJobModels,omitempty" xml:"MvRefreshJobModels,omitempty" type:"Repeated"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries returned on each page.
	//
	// example:
	//
	// 30
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The total number of entries that are returned.
	//
	// example:
	//
	// 44
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeViewJobsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DescribeViewJobsResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeViewJobsResponseBodyData) GetMvRefreshJobModels() []*OpenStructRefreshJobModel {
	return s.MvRefreshJobModels
}

func (s *DescribeViewJobsResponseBodyData) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *DescribeViewJobsResponseBodyData) GetPageSize() *int64 {
	return s.PageSize
}

func (s *DescribeViewJobsResponseBodyData) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *DescribeViewJobsResponseBodyData) SetMvRefreshJobModels(v []*OpenStructRefreshJobModel) *DescribeViewJobsResponseBodyData {
	s.MvRefreshJobModels = v
	return s
}

func (s *DescribeViewJobsResponseBodyData) SetPageNumber(v int64) *DescribeViewJobsResponseBodyData {
	s.PageNumber = &v
	return s
}

func (s *DescribeViewJobsResponseBodyData) SetPageSize(v int64) *DescribeViewJobsResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *DescribeViewJobsResponseBodyData) SetTotalCount(v int64) *DescribeViewJobsResponseBodyData {
	s.TotalCount = &v
	return s
}

func (s *DescribeViewJobsResponseBodyData) Validate() error {
	if s.MvRefreshJobModels != nil {
		for _, item := range s.MvRefreshJobModels {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
