// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeMvRecommendSubTasksResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *DescribeMvRecommendSubTasksResponseBodyData) *DescribeMvRecommendSubTasksResponseBody
	GetData() *DescribeMvRecommendSubTasksResponseBodyData
	SetPageNumber(v int64) *DescribeMvRecommendSubTasksResponseBody
	GetPageNumber() *int64
	SetPageSize(v int64) *DescribeMvRecommendSubTasksResponseBody
	GetPageSize() *int64
	SetRequestId(v string) *DescribeMvRecommendSubTasksResponseBody
	GetRequestId() *string
	SetTotalCount(v int64) *DescribeMvRecommendSubTasksResponseBody
	GetTotalCount() *int64
}

type DescribeMvRecommendSubTasksResponseBody struct {
	// The returned data.
	Data *DescribeMvRecommendSubTasksResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
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
	// The total number of entries returned.
	//
	// example:
	//
	// 4
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeMvRecommendSubTasksResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeMvRecommendSubTasksResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeMvRecommendSubTasksResponseBody) GetData() *DescribeMvRecommendSubTasksResponseBodyData {
	return s.Data
}

func (s *DescribeMvRecommendSubTasksResponseBody) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *DescribeMvRecommendSubTasksResponseBody) GetPageSize() *int64 {
	return s.PageSize
}

func (s *DescribeMvRecommendSubTasksResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeMvRecommendSubTasksResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *DescribeMvRecommendSubTasksResponseBody) SetData(v *DescribeMvRecommendSubTasksResponseBodyData) *DescribeMvRecommendSubTasksResponseBody {
	s.Data = v
	return s
}

func (s *DescribeMvRecommendSubTasksResponseBody) SetPageNumber(v int64) *DescribeMvRecommendSubTasksResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeMvRecommendSubTasksResponseBody) SetPageSize(v int64) *DescribeMvRecommendSubTasksResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeMvRecommendSubTasksResponseBody) SetRequestId(v string) *DescribeMvRecommendSubTasksResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeMvRecommendSubTasksResponseBody) SetTotalCount(v int64) *DescribeMvRecommendSubTasksResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeMvRecommendSubTasksResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeMvRecommendSubTasksResponseBodyData struct {
	// The list of recommended tasks to be executed.
	MvRecommendSubTaskModels []*OpenStructMvRecommendSubTaskModel `json:"MvRecommendSubTaskModels,omitempty" xml:"MvRecommendSubTaskModels,omitempty" type:"Repeated"`
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
	// The total number of entries returned.
	//
	// example:
	//
	// 4
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeMvRecommendSubTasksResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DescribeMvRecommendSubTasksResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeMvRecommendSubTasksResponseBodyData) GetMvRecommendSubTaskModels() []*OpenStructMvRecommendSubTaskModel {
	return s.MvRecommendSubTaskModels
}

func (s *DescribeMvRecommendSubTasksResponseBodyData) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *DescribeMvRecommendSubTasksResponseBodyData) GetPageSize() *int64 {
	return s.PageSize
}

func (s *DescribeMvRecommendSubTasksResponseBodyData) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *DescribeMvRecommendSubTasksResponseBodyData) SetMvRecommendSubTaskModels(v []*OpenStructMvRecommendSubTaskModel) *DescribeMvRecommendSubTasksResponseBodyData {
	s.MvRecommendSubTaskModels = v
	return s
}

func (s *DescribeMvRecommendSubTasksResponseBodyData) SetPageNumber(v int64) *DescribeMvRecommendSubTasksResponseBodyData {
	s.PageNumber = &v
	return s
}

func (s *DescribeMvRecommendSubTasksResponseBodyData) SetPageSize(v int64) *DescribeMvRecommendSubTasksResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *DescribeMvRecommendSubTasksResponseBodyData) SetTotalCount(v int64) *DescribeMvRecommendSubTasksResponseBodyData {
	s.TotalCount = &v
	return s
}

func (s *DescribeMvRecommendSubTasksResponseBodyData) Validate() error {
	if s.MvRecommendSubTaskModels != nil {
		for _, item := range s.MvRecommendSubTaskModels {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
