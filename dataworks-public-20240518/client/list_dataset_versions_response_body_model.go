// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDatasetVersionsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPagingInfo(v *ListDatasetVersionsResponseBodyPagingInfo) *ListDatasetVersionsResponseBody
	GetPagingInfo() *ListDatasetVersionsResponseBodyPagingInfo
	SetRequestId(v string) *ListDatasetVersionsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListDatasetVersionsResponseBody
	GetSuccess() *bool
}

type ListDatasetVersionsResponseBody struct {
	// The pagination information.
	PagingInfo *ListDatasetVersionsResponseBodyPagingInfo `json:"PagingInfo,omitempty" xml:"PagingInfo,omitempty" type:"Struct"`
	// RequestId
	//
	// example:
	//
	// E25887B7-579C-54A5-9C4F-83A0DE367XXX
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListDatasetVersionsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListDatasetVersionsResponseBody) GoString() string {
	return s.String()
}

func (s *ListDatasetVersionsResponseBody) GetPagingInfo() *ListDatasetVersionsResponseBodyPagingInfo {
	return s.PagingInfo
}

func (s *ListDatasetVersionsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListDatasetVersionsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListDatasetVersionsResponseBody) SetPagingInfo(v *ListDatasetVersionsResponseBodyPagingInfo) *ListDatasetVersionsResponseBody {
	s.PagingInfo = v
	return s
}

func (s *ListDatasetVersionsResponseBody) SetRequestId(v string) *ListDatasetVersionsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListDatasetVersionsResponseBody) SetSuccess(v bool) *ListDatasetVersionsResponseBody {
	s.Success = &v
	return s
}

func (s *ListDatasetVersionsResponseBody) Validate() error {
	if s.PagingInfo != nil {
		if err := s.PagingInfo.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListDatasetVersionsResponseBodyPagingInfo struct {
	// The dataset version list.
	DatasetVersions []*DatasetVersion `json:"DatasetVersions,omitempty" xml:"DatasetVersions,omitempty" type:"Repeated"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The total number of entries.
	//
	// example:
	//
	// 100
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListDatasetVersionsResponseBodyPagingInfo) String() string {
	return dara.Prettify(s)
}

func (s ListDatasetVersionsResponseBodyPagingInfo) GoString() string {
	return s.String()
}

func (s *ListDatasetVersionsResponseBodyPagingInfo) GetDatasetVersions() []*DatasetVersion {
	return s.DatasetVersions
}

func (s *ListDatasetVersionsResponseBodyPagingInfo) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListDatasetVersionsResponseBodyPagingInfo) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListDatasetVersionsResponseBodyPagingInfo) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListDatasetVersionsResponseBodyPagingInfo) SetDatasetVersions(v []*DatasetVersion) *ListDatasetVersionsResponseBodyPagingInfo {
	s.DatasetVersions = v
	return s
}

func (s *ListDatasetVersionsResponseBodyPagingInfo) SetPageNumber(v int32) *ListDatasetVersionsResponseBodyPagingInfo {
	s.PageNumber = &v
	return s
}

func (s *ListDatasetVersionsResponseBodyPagingInfo) SetPageSize(v int32) *ListDatasetVersionsResponseBodyPagingInfo {
	s.PageSize = &v
	return s
}

func (s *ListDatasetVersionsResponseBodyPagingInfo) SetTotalCount(v int64) *ListDatasetVersionsResponseBodyPagingInfo {
	s.TotalCount = &v
	return s
}

func (s *ListDatasetVersionsResponseBodyPagingInfo) Validate() error {
	if s.DatasetVersions != nil {
		for _, item := range s.DatasetVersions {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
