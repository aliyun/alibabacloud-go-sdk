// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDatasetVersionsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCreatorId(v string) *ListDatasetVersionsRequest
	GetCreatorId() *string
	SetDatasetId(v string) *ListDatasetVersionsRequest
	GetDatasetId() *string
	SetOrder(v string) *ListDatasetVersionsRequest
	GetOrder() *string
	SetPageNumber(v int32) *ListDatasetVersionsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListDatasetVersionsRequest
	GetPageSize() *int32
	SetSortBy(v string) *ListDatasetVersionsRequest
	GetSortBy() *string
}

type ListDatasetVersionsRequest struct {
	// The creator ID.
	//
	// example:
	//
	// 12103XXX46492139
	CreatorId *string `json:"CreatorId,omitempty" xml:"CreatorId,omitempty"`
	// The dataset ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// dataworks-dataset:3pXXXb8o0ngr07njhps1
	DatasetId *string `json:"DatasetId,omitempty" xml:"DatasetId,omitempty"`
	// The sort order. Default: Desc.
	//
	// Valid values:
	//
	// 	- Asc: Ascending.
	//
	// 	- Desc: Descending.
	//
	// example:
	//
	// Desc
	Order *string `json:"Order,omitempty" xml:"Order,omitempty"`
	// The page number. Default: 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Default: 10. Maximum: 100.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The sort field. Default: VersionNumber.
	//
	// Valid values:
	//
	// 	- ModifyTime: Modification time.
	//
	// 	- CreateTime: Creation time.
	//
	// 	- VersionNumber: Version number.
	//
	// example:
	//
	// CreateTime
	SortBy *string `json:"SortBy,omitempty" xml:"SortBy,omitempty"`
}

func (s ListDatasetVersionsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListDatasetVersionsRequest) GoString() string {
	return s.String()
}

func (s *ListDatasetVersionsRequest) GetCreatorId() *string {
	return s.CreatorId
}

func (s *ListDatasetVersionsRequest) GetDatasetId() *string {
	return s.DatasetId
}

func (s *ListDatasetVersionsRequest) GetOrder() *string {
	return s.Order
}

func (s *ListDatasetVersionsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListDatasetVersionsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListDatasetVersionsRequest) GetSortBy() *string {
	return s.SortBy
}

func (s *ListDatasetVersionsRequest) SetCreatorId(v string) *ListDatasetVersionsRequest {
	s.CreatorId = &v
	return s
}

func (s *ListDatasetVersionsRequest) SetDatasetId(v string) *ListDatasetVersionsRequest {
	s.DatasetId = &v
	return s
}

func (s *ListDatasetVersionsRequest) SetOrder(v string) *ListDatasetVersionsRequest {
	s.Order = &v
	return s
}

func (s *ListDatasetVersionsRequest) SetPageNumber(v int32) *ListDatasetVersionsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListDatasetVersionsRequest) SetPageSize(v int32) *ListDatasetVersionsRequest {
	s.PageSize = &v
	return s
}

func (s *ListDatasetVersionsRequest) SetSortBy(v string) *ListDatasetVersionsRequest {
	s.SortBy = &v
	return s
}

func (s *ListDatasetVersionsRequest) Validate() error {
	return dara.Validate(s)
}
