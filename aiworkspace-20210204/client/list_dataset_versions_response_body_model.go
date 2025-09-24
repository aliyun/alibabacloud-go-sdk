// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDatasetVersionsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDatasetVersions(v []*DatasetVersion) *ListDatasetVersionsResponseBody
	GetDatasetVersions() []*DatasetVersion
	SetPageNumber(v int32) *ListDatasetVersionsResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *ListDatasetVersionsResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *ListDatasetVersionsResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *ListDatasetVersionsResponseBody
	GetTotalCount() *int32
}

type ListDatasetVersionsResponseBody struct {
	// The dataset versions.
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
	// The ID of the request.
	//
	// example:
	//
	// 0648C5BB-68D0-54D2-92A5-607135B8806B
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The number of dataset versions that meet the filter conditions.
	//
	// example:
	//
	// 5
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListDatasetVersionsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListDatasetVersionsResponseBody) GoString() string {
	return s.String()
}

func (s *ListDatasetVersionsResponseBody) GetDatasetVersions() []*DatasetVersion {
	return s.DatasetVersions
}

func (s *ListDatasetVersionsResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListDatasetVersionsResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListDatasetVersionsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListDatasetVersionsResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListDatasetVersionsResponseBody) SetDatasetVersions(v []*DatasetVersion) *ListDatasetVersionsResponseBody {
	s.DatasetVersions = v
	return s
}

func (s *ListDatasetVersionsResponseBody) SetPageNumber(v int32) *ListDatasetVersionsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListDatasetVersionsResponseBody) SetPageSize(v int32) *ListDatasetVersionsResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListDatasetVersionsResponseBody) SetRequestId(v string) *ListDatasetVersionsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListDatasetVersionsResponseBody) SetTotalCount(v int32) *ListDatasetVersionsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListDatasetVersionsResponseBody) Validate() error {
	return dara.Validate(s)
}
