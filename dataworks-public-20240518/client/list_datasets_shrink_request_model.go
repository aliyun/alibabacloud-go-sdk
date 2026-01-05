// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDatasetsShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCreatorId(v string) *ListDatasetsShrinkRequest
	GetCreatorId() *string
	SetDataTypeListShrink(v string) *ListDatasetsShrinkRequest
	GetDataTypeListShrink() *string
	SetName(v string) *ListDatasetsShrinkRequest
	GetName() *string
	SetOrder(v string) *ListDatasetsShrinkRequest
	GetOrder() *string
	SetOrigin(v string) *ListDatasetsShrinkRequest
	GetOrigin() *string
	SetPageNumber(v int32) *ListDatasetsShrinkRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListDatasetsShrinkRequest
	GetPageSize() *int32
	SetProjectId(v int64) *ListDatasetsShrinkRequest
	GetProjectId() *int64
	SetSortBy(v string) *ListDatasetsShrinkRequest
	GetSortBy() *string
	SetStorageTypeListShrink(v string) *ListDatasetsShrinkRequest
	GetStorageTypeListShrink() *string
}

type ListDatasetsShrinkRequest struct {
	// The creator ID.
	//
	// example:
	//
	// 12103XXX46492139
	CreatorId *string `json:"CreatorId,omitempty" xml:"CreatorId,omitempty"`
	// The data type. Multiple selections are allowed. Valid values:
	//
	// 	- COMMON
	//
	// 	- PIC
	//
	// 	- TEXT
	//
	// 	- TABLE
	//
	// 	- VIDEO
	//
	// 	- AUDIO
	//
	// 	- INDEX
	DataTypeListShrink *string `json:"DataTypeList,omitempty" xml:"DataTypeList,omitempty"`
	// The dataset name. Supports fuzzy search.
	//
	// example:
	//
	// test_dataset
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
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
	// Asc
	Order *string `json:"Order,omitempty" xml:"Order,omitempty"`
	// The dataset source. Valid values:
	//
	// 	- DataWorks
	//
	// 	- PAI
	//
	// example:
	//
	// DataWorks
	Origin *string `json:"Origin,omitempty" xml:"Origin,omitempty"`
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
	// The DataWorks workspace ID.
	//
	// example:
	//
	// 251363
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The sort field. Default: CreateTime.
	//
	// Valid values:
	//
	// 	- ModifyTime: Modification time.
	//
	// 	- CreateTime: Creation time.
	//
	// 	- Name
	//
	// example:
	//
	// CreateTime
	SortBy *string `json:"SortBy,omitempty" xml:"SortBy,omitempty"`
	// The storage type. Multiple selections are allowed. Supported values:
	//
	// 	- OSS
	//
	// 	- NAS: General-purpose NAS file systems
	//
	// 	- EXTREMENAS: Extreme NAS file systems
	//
	// 	- DLF_LANCE: Data Lake Formation
	//
	// 	- CPFS: Cloud Parallel File Storage
	//
	// 	- BMCPFS: CPFS for Lingjun
	//
	// 	- MAXCOMPUTE: MaxCompute table
	StorageTypeListShrink *string `json:"StorageTypeList,omitempty" xml:"StorageTypeList,omitempty"`
}

func (s ListDatasetsShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ListDatasetsShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListDatasetsShrinkRequest) GetCreatorId() *string {
	return s.CreatorId
}

func (s *ListDatasetsShrinkRequest) GetDataTypeListShrink() *string {
	return s.DataTypeListShrink
}

func (s *ListDatasetsShrinkRequest) GetName() *string {
	return s.Name
}

func (s *ListDatasetsShrinkRequest) GetOrder() *string {
	return s.Order
}

func (s *ListDatasetsShrinkRequest) GetOrigin() *string {
	return s.Origin
}

func (s *ListDatasetsShrinkRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListDatasetsShrinkRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListDatasetsShrinkRequest) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *ListDatasetsShrinkRequest) GetSortBy() *string {
	return s.SortBy
}

func (s *ListDatasetsShrinkRequest) GetStorageTypeListShrink() *string {
	return s.StorageTypeListShrink
}

func (s *ListDatasetsShrinkRequest) SetCreatorId(v string) *ListDatasetsShrinkRequest {
	s.CreatorId = &v
	return s
}

func (s *ListDatasetsShrinkRequest) SetDataTypeListShrink(v string) *ListDatasetsShrinkRequest {
	s.DataTypeListShrink = &v
	return s
}

func (s *ListDatasetsShrinkRequest) SetName(v string) *ListDatasetsShrinkRequest {
	s.Name = &v
	return s
}

func (s *ListDatasetsShrinkRequest) SetOrder(v string) *ListDatasetsShrinkRequest {
	s.Order = &v
	return s
}

func (s *ListDatasetsShrinkRequest) SetOrigin(v string) *ListDatasetsShrinkRequest {
	s.Origin = &v
	return s
}

func (s *ListDatasetsShrinkRequest) SetPageNumber(v int32) *ListDatasetsShrinkRequest {
	s.PageNumber = &v
	return s
}

func (s *ListDatasetsShrinkRequest) SetPageSize(v int32) *ListDatasetsShrinkRequest {
	s.PageSize = &v
	return s
}

func (s *ListDatasetsShrinkRequest) SetProjectId(v int64) *ListDatasetsShrinkRequest {
	s.ProjectId = &v
	return s
}

func (s *ListDatasetsShrinkRequest) SetSortBy(v string) *ListDatasetsShrinkRequest {
	s.SortBy = &v
	return s
}

func (s *ListDatasetsShrinkRequest) SetStorageTypeListShrink(v string) *ListDatasetsShrinkRequest {
	s.StorageTypeListShrink = &v
	return s
}

func (s *ListDatasetsShrinkRequest) Validate() error {
	return dara.Validate(s)
}
