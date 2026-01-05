// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDatasetsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCreatorId(v string) *ListDatasetsRequest
	GetCreatorId() *string
	SetDataTypeList(v []*string) *ListDatasetsRequest
	GetDataTypeList() []*string
	SetName(v string) *ListDatasetsRequest
	GetName() *string
	SetOrder(v string) *ListDatasetsRequest
	GetOrder() *string
	SetOrigin(v string) *ListDatasetsRequest
	GetOrigin() *string
	SetPageNumber(v int32) *ListDatasetsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListDatasetsRequest
	GetPageSize() *int32
	SetProjectId(v int64) *ListDatasetsRequest
	GetProjectId() *int64
	SetSortBy(v string) *ListDatasetsRequest
	GetSortBy() *string
	SetStorageTypeList(v []*string) *ListDatasetsRequest
	GetStorageTypeList() []*string
}

type ListDatasetsRequest struct {
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
	DataTypeList []*string `json:"DataTypeList,omitempty" xml:"DataTypeList,omitempty" type:"Repeated"`
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
	StorageTypeList []*string `json:"StorageTypeList,omitempty" xml:"StorageTypeList,omitempty" type:"Repeated"`
}

func (s ListDatasetsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListDatasetsRequest) GoString() string {
	return s.String()
}

func (s *ListDatasetsRequest) GetCreatorId() *string {
	return s.CreatorId
}

func (s *ListDatasetsRequest) GetDataTypeList() []*string {
	return s.DataTypeList
}

func (s *ListDatasetsRequest) GetName() *string {
	return s.Name
}

func (s *ListDatasetsRequest) GetOrder() *string {
	return s.Order
}

func (s *ListDatasetsRequest) GetOrigin() *string {
	return s.Origin
}

func (s *ListDatasetsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListDatasetsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListDatasetsRequest) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *ListDatasetsRequest) GetSortBy() *string {
	return s.SortBy
}

func (s *ListDatasetsRequest) GetStorageTypeList() []*string {
	return s.StorageTypeList
}

func (s *ListDatasetsRequest) SetCreatorId(v string) *ListDatasetsRequest {
	s.CreatorId = &v
	return s
}

func (s *ListDatasetsRequest) SetDataTypeList(v []*string) *ListDatasetsRequest {
	s.DataTypeList = v
	return s
}

func (s *ListDatasetsRequest) SetName(v string) *ListDatasetsRequest {
	s.Name = &v
	return s
}

func (s *ListDatasetsRequest) SetOrder(v string) *ListDatasetsRequest {
	s.Order = &v
	return s
}

func (s *ListDatasetsRequest) SetOrigin(v string) *ListDatasetsRequest {
	s.Origin = &v
	return s
}

func (s *ListDatasetsRequest) SetPageNumber(v int32) *ListDatasetsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListDatasetsRequest) SetPageSize(v int32) *ListDatasetsRequest {
	s.PageSize = &v
	return s
}

func (s *ListDatasetsRequest) SetProjectId(v int64) *ListDatasetsRequest {
	s.ProjectId = &v
	return s
}

func (s *ListDatasetsRequest) SetSortBy(v string) *ListDatasetsRequest {
	s.SortBy = &v
	return s
}

func (s *ListDatasetsRequest) SetStorageTypeList(v []*string) *ListDatasetsRequest {
	s.StorageTypeList = v
	return s
}

func (s *ListDatasetsRequest) Validate() error {
	return dara.Validate(s)
}
