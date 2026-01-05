// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDatasetsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPagingInfo(v *ListDatasetsResponseBodyPagingInfo) *ListDatasetsResponseBody
	GetPagingInfo() *ListDatasetsResponseBodyPagingInfo
	SetRequestId(v string) *ListDatasetsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListDatasetsResponseBody
	GetSuccess() *bool
}

type ListDatasetsResponseBody struct {
	// The pagination information.
	PagingInfo *ListDatasetsResponseBodyPagingInfo `json:"PagingInfo,omitempty" xml:"PagingInfo,omitempty" type:"Struct"`
	// RequestId
	//
	// example:
	//
	// 6A6CBE87-9F91-1323-B680-E7A7065XXXXX
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListDatasetsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListDatasetsResponseBody) GoString() string {
	return s.String()
}

func (s *ListDatasetsResponseBody) GetPagingInfo() *ListDatasetsResponseBodyPagingInfo {
	return s.PagingInfo
}

func (s *ListDatasetsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListDatasetsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListDatasetsResponseBody) SetPagingInfo(v *ListDatasetsResponseBodyPagingInfo) *ListDatasetsResponseBody {
	s.PagingInfo = v
	return s
}

func (s *ListDatasetsResponseBody) SetRequestId(v string) *ListDatasetsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListDatasetsResponseBody) SetSuccess(v bool) *ListDatasetsResponseBody {
	s.Success = &v
	return s
}

func (s *ListDatasetsResponseBody) Validate() error {
	if s.PagingInfo != nil {
		if err := s.PagingInfo.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListDatasetsResponseBodyPagingInfo struct {
	// The dataset list.
	Datasets []*Dataset `json:"Datasets,omitempty" xml:"Datasets,omitempty" type:"Repeated"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries on this page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 100
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListDatasetsResponseBodyPagingInfo) String() string {
	return dara.Prettify(s)
}

func (s ListDatasetsResponseBodyPagingInfo) GoString() string {
	return s.String()
}

func (s *ListDatasetsResponseBodyPagingInfo) GetDatasets() []*Dataset {
	return s.Datasets
}

func (s *ListDatasetsResponseBodyPagingInfo) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListDatasetsResponseBodyPagingInfo) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListDatasetsResponseBodyPagingInfo) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListDatasetsResponseBodyPagingInfo) SetDatasets(v []*Dataset) *ListDatasetsResponseBodyPagingInfo {
	s.Datasets = v
	return s
}

func (s *ListDatasetsResponseBodyPagingInfo) SetPageNumber(v int32) *ListDatasetsResponseBodyPagingInfo {
	s.PageNumber = &v
	return s
}

func (s *ListDatasetsResponseBodyPagingInfo) SetPageSize(v int32) *ListDatasetsResponseBodyPagingInfo {
	s.PageSize = &v
	return s
}

func (s *ListDatasetsResponseBodyPagingInfo) SetTotalCount(v int64) *ListDatasetsResponseBodyPagingInfo {
	s.TotalCount = &v
	return s
}

func (s *ListDatasetsResponseBodyPagingInfo) Validate() error {
	if s.Datasets != nil {
		for _, item := range s.Datasets {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
