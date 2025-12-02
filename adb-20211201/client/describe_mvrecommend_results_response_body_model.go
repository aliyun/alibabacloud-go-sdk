// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeMVRecommendResultsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *DescribeMVRecommendResultsResponseBodyData) *DescribeMVRecommendResultsResponseBody
	GetData() *DescribeMVRecommendResultsResponseBodyData
	SetPageNumber(v int64) *DescribeMVRecommendResultsResponseBody
	GetPageNumber() *int64
	SetPageSize(v int64) *DescribeMVRecommendResultsResponseBody
	GetPageSize() *int64
	SetRequestId(v string) *DescribeMVRecommendResultsResponseBody
	GetRequestId() *string
	SetTotalCount(v int64) *DescribeMVRecommendResultsResponseBody
	GetTotalCount() *int64
}

type DescribeMVRecommendResultsResponseBody struct {
	// The returned data.
	Data *DescribeMVRecommendResultsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
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
	// 20
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeMVRecommendResultsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeMVRecommendResultsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeMVRecommendResultsResponseBody) GetData() *DescribeMVRecommendResultsResponseBodyData {
	return s.Data
}

func (s *DescribeMVRecommendResultsResponseBody) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *DescribeMVRecommendResultsResponseBody) GetPageSize() *int64 {
	return s.PageSize
}

func (s *DescribeMVRecommendResultsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeMVRecommendResultsResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *DescribeMVRecommendResultsResponseBody) SetData(v *DescribeMVRecommendResultsResponseBodyData) *DescribeMVRecommendResultsResponseBody {
	s.Data = v
	return s
}

func (s *DescribeMVRecommendResultsResponseBody) SetPageNumber(v int64) *DescribeMVRecommendResultsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeMVRecommendResultsResponseBody) SetPageSize(v int64) *DescribeMVRecommendResultsResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeMVRecommendResultsResponseBody) SetRequestId(v string) *DescribeMVRecommendResultsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeMVRecommendResultsResponseBody) SetTotalCount(v int64) *DescribeMVRecommendResultsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeMVRecommendResultsResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeMVRecommendResultsResponseBodyData struct {
	MvRecommendResultModels []*OpenStructMVRecommendResultModel `json:"MvRecommendResultModels,omitempty" xml:"MvRecommendResultModels,omitempty" type:"Repeated"`
	// The page number. Default value: **1**.
	//
	// example:
	//
	// 1
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries returned per page. Valid values:
	//
	// 	- **30*	- (default).
	//
	// 	- **50**.
	//
	// 	- **100**.
	//
	// example:
	//
	// 20
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The total number of entries.
	//
	// example:
	//
	// 1
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeMVRecommendResultsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DescribeMVRecommendResultsResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeMVRecommendResultsResponseBodyData) GetMvRecommendResultModels() []*OpenStructMVRecommendResultModel {
	return s.MvRecommendResultModels
}

func (s *DescribeMVRecommendResultsResponseBodyData) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *DescribeMVRecommendResultsResponseBodyData) GetPageSize() *int64 {
	return s.PageSize
}

func (s *DescribeMVRecommendResultsResponseBodyData) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *DescribeMVRecommendResultsResponseBodyData) SetMvRecommendResultModels(v []*OpenStructMVRecommendResultModel) *DescribeMVRecommendResultsResponseBodyData {
	s.MvRecommendResultModels = v
	return s
}

func (s *DescribeMVRecommendResultsResponseBodyData) SetPageNumber(v int64) *DescribeMVRecommendResultsResponseBodyData {
	s.PageNumber = &v
	return s
}

func (s *DescribeMVRecommendResultsResponseBodyData) SetPageSize(v int64) *DescribeMVRecommendResultsResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *DescribeMVRecommendResultsResponseBodyData) SetTotalCount(v int64) *DescribeMVRecommendResultsResponseBodyData {
	s.TotalCount = &v
	return s
}

func (s *DescribeMVRecommendResultsResponseBodyData) Validate() error {
	if s.MvRecommendResultModels != nil {
		for _, item := range s.MvRecommendResultModels {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
