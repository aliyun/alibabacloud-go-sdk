// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeMvRecommendTasksResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *DescribeMvRecommendTasksResponseBodyData) *DescribeMvRecommendTasksResponseBody
	GetData() *DescribeMvRecommendTasksResponseBodyData
	SetPageNumber(v int64) *DescribeMvRecommendTasksResponseBody
	GetPageNumber() *int64
	SetPageSize(v int64) *DescribeMvRecommendTasksResponseBody
	GetPageSize() *int64
	SetRequestId(v string) *DescribeMvRecommendTasksResponseBody
	GetRequestId() *string
	SetTotalCount(v int64) *DescribeMvRecommendTasksResponseBody
	GetTotalCount() *int64
}

type DescribeMvRecommendTasksResponseBody struct {
	// The data returned.
	Data *DescribeMvRecommendTasksResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
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
	// 1AD222E9-E606-4A42-BF6D-8A4442913CEF
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries that are returned.
	//
	// example:
	//
	// 4
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeMvRecommendTasksResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeMvRecommendTasksResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeMvRecommendTasksResponseBody) GetData() *DescribeMvRecommendTasksResponseBodyData {
	return s.Data
}

func (s *DescribeMvRecommendTasksResponseBody) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *DescribeMvRecommendTasksResponseBody) GetPageSize() *int64 {
	return s.PageSize
}

func (s *DescribeMvRecommendTasksResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeMvRecommendTasksResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *DescribeMvRecommendTasksResponseBody) SetData(v *DescribeMvRecommendTasksResponseBodyData) *DescribeMvRecommendTasksResponseBody {
	s.Data = v
	return s
}

func (s *DescribeMvRecommendTasksResponseBody) SetPageNumber(v int64) *DescribeMvRecommendTasksResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeMvRecommendTasksResponseBody) SetPageSize(v int64) *DescribeMvRecommendTasksResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeMvRecommendTasksResponseBody) SetRequestId(v string) *DescribeMvRecommendTasksResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeMvRecommendTasksResponseBody) SetTotalCount(v int64) *DescribeMvRecommendTasksResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeMvRecommendTasksResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeMvRecommendTasksResponseBodyData struct {
	// The list of recommended tasks for materialized views.
	MvRecommendTaskModels []*OpenStructMvRecommendTaskModel `json:"MvRecommendTaskModels,omitempty" xml:"MvRecommendTaskModels,omitempty" type:"Repeated"`
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
	// The total number of entries that are returned.
	//
	// example:
	//
	// 4
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeMvRecommendTasksResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DescribeMvRecommendTasksResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeMvRecommendTasksResponseBodyData) GetMvRecommendTaskModels() []*OpenStructMvRecommendTaskModel {
	return s.MvRecommendTaskModels
}

func (s *DescribeMvRecommendTasksResponseBodyData) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *DescribeMvRecommendTasksResponseBodyData) GetPageSize() *int64 {
	return s.PageSize
}

func (s *DescribeMvRecommendTasksResponseBodyData) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *DescribeMvRecommendTasksResponseBodyData) SetMvRecommendTaskModels(v []*OpenStructMvRecommendTaskModel) *DescribeMvRecommendTasksResponseBodyData {
	s.MvRecommendTaskModels = v
	return s
}

func (s *DescribeMvRecommendTasksResponseBodyData) SetPageNumber(v int64) *DescribeMvRecommendTasksResponseBodyData {
	s.PageNumber = &v
	return s
}

func (s *DescribeMvRecommendTasksResponseBodyData) SetPageSize(v int64) *DescribeMvRecommendTasksResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *DescribeMvRecommendTasksResponseBodyData) SetTotalCount(v int64) *DescribeMvRecommendTasksResponseBodyData {
	s.TotalCount = &v
	return s
}

func (s *DescribeMvRecommendTasksResponseBodyData) Validate() error {
	if s.MvRecommendTaskModels != nil {
		for _, item := range s.MvRecommendTaskModels {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
