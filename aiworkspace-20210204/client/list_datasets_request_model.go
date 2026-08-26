// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDatasetsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccessibility(v string) *ListDatasetsRequest
	GetAccessibility() *string
	SetDataSourceTypes(v string) *ListDatasetsRequest
	GetDataSourceTypes() *string
	SetDataTypes(v string) *ListDatasetsRequest
	GetDataTypes() *string
	SetDatasetIds(v string) *ListDatasetsRequest
	GetDatasetIds() *string
	SetEdition(v string) *ListDatasetsRequest
	GetEdition() *string
	SetLabel(v string) *ListDatasetsRequest
	GetLabel() *string
	SetName(v string) *ListDatasetsRequest
	GetName() *string
	SetOrder(v string) *ListDatasetsRequest
	GetOrder() *string
	SetPageNumber(v int32) *ListDatasetsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListDatasetsRequest
	GetPageSize() *int32
	SetProperties(v string) *ListDatasetsRequest
	GetProperties() *string
	SetProvider(v string) *ListDatasetsRequest
	GetProvider() *string
	SetShareScope(v string) *ListDatasetsRequest
	GetShareScope() *string
	SetSortBy(v string) *ListDatasetsRequest
	GetSortBy() *string
	SetSourceDatasetId(v string) *ListDatasetsRequest
	GetSourceDatasetId() *string
	SetSourceId(v string) *ListDatasetsRequest
	GetSourceId() *string
	SetSourceTypes(v string) *ListDatasetsRequest
	GetSourceTypes() *string
	SetWorkspaceId(v string) *ListDatasetsRequest
	GetWorkspaceId() *string
}

type ListDatasetsRequest struct {
	// The visibility of the dataset.
	//
	// - PUBLIC: public.
	//
	// - PRIVATE: private.
	//
	// example:
	//
	// PRIVATE
	Accessibility *string `json:"Accessibility,omitempty" xml:"Accessibility,omitempty"`
	// The data source types. Separate multiple values with commas (,). Valid values:
	//
	// - NAS: Alibaba Cloud Network Attached Storage (NAS).
	//
	// - OSS: Alibaba Cloud Object Storage Service (OSS).
	//
	// example:
	//
	// OSS
	DataSourceTypes *string `json:"DataSourceTypes,omitempty" xml:"DataSourceTypes,omitempty"`
	// The data types of the dataset. Separate multiple values with commas (,). Valid values:
	//
	// - VIDEO: video.
	//
	// - COMMON: common.
	//
	// - TEXT: text.
	//
	// - PIC: image.
	//
	// - AUDIO: audio.
	//
	// example:
	//
	// COMMON,TEXT
	DataTypes *string `json:"DataTypes,omitempty" xml:"DataTypes,omitempty"`
	// The dataset IDs. You can specify multiple dataset IDs separated by commas (,).
	//
	// example:
	//
	// d-rcdg3wxxxxxhc5jk87
	DatasetIds *string `json:"DatasetIds,omitempty" xml:"DatasetIds,omitempty"`
	// The dataset edition. Valid values:
	//
	// - BASIC: Basic Edition. Does not support dataset file metadata management.
	//
	// - ADVANCED: Advanced Edition. Supported only for OSS type. Each version supports up to 1 million file metadata entries.
	//
	// - LOGICAL: Logical Edition. Supported only for OSS type. Each version supports up to 1 million file metadata entries. Applicable to most scenarios and requires the use of the SDK.
	//
	// example:
	//
	// BASIC
	Edition *string `json:"Edition,omitempty" xml:"Edition,omitempty"`
	// The dataset label used to filter the dataset list. Datasets whose label key or value contains the specified string are returned.
	//
	// example:
	//
	// test
	Label *string `json:"Label,omitempty" xml:"Label,omitempty"`
	// The dataset name. Fuzzy match is supported based on the dataset name.
	//
	// example:
	//
	// myName
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The sorting order for the specified sort field in paging queries. Default value: ASC.
	//
	// - ASC: ascending order.
	//
	// - DESC: descending order.
	//
	// example:
	//
	// ASC
	Order *string `json:"Order,omitempty" xml:"Order,omitempty"`
	// The page number of the dataset list. Minimum value: 1. Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page settings for paging queries. Default value: 10.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The dataset properties. Separate multiple values with commas (,). Valid values:
	//
	// - DIRECTORY: folder.
	//
	// - FILE: file.
	//
	// example:
	//
	// FILE
	Properties *string `json:"Properties,omitempty" xml:"Properties,omitempty"`
	// The dataset provider. A value of "pai" indicates that the dataset is a PAI platform public dataset.
	//
	// example:
	//
	// pai
	Provider *string `json:"Provider,omitempty" xml:"Provider,omitempty"`
	// The sharing filter for datasets:
	//
	// 	- TO_ME: returns only datasets shared with you.
	//
	// 	- BY_ME: returns only datasets you shared with others, with sharing configuration details displayed.
	//
	// 	- If this parameter is not set or is set to empty: returns all datasets in the current workspace, including TO_ME.
	//
	// example:
	//
	// BY_ME
	ShareScope *string `json:"ShareScope,omitempty" xml:"ShareScope,omitempty"`
	// The field by which to sort the results.
	//
	// example:
	//
	// GmtCreateTime
	SortBy *string `json:"SortBy,omitempty" xml:"SortBy,omitempty"`
	// The source dataset ID of the iTAG annotation set.
	//
	// example:
	//
	// d-rcdg3wxxxxxhc5jk87
	SourceDatasetId *string `json:"SourceDatasetId,omitempty" xml:"SourceDatasetId,omitempty"`
	// The data source ID.
	//
	// - If SourceTypes is set to USER, you can customize the SourceId value.
	//
	// - If SourceTypes is set to ITAG, which indicates a dataset generated from iTAG annotation results, SourceId is the iTAG task ID.
	//
	// - If SourceTypes is set to PAI_PUBLIC_DATASET, which indicates a dataset created from a PAI public dataset, SourceId is empty by default.
	//
	// example:
	//
	// d-rbvg5wzljzjhc9ks92
	SourceId *string `json:"SourceId,omitempty" xml:"SourceId,omitempty"`
	// The source types. Separate multiple values with commas (,).
	//
	// example:
	//
	// USER,ITAG
	SourceTypes *string `json:"SourceTypes,omitempty" xml:"SourceTypes,omitempty"`
	// The ID of the workspace where the dataset resides. For information about how to obtain the workspace ID, see [ListWorkspaces](https://help.aliyun.com/document_detail/449124.html).
	//
	// If you do not specify this parameter, the default workspace is used. If the default workspace does not exist, an error is returned.
	//
	// example:
	//
	// 324**
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s ListDatasetsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListDatasetsRequest) GoString() string {
	return s.String()
}

func (s *ListDatasetsRequest) GetAccessibility() *string {
	return s.Accessibility
}

func (s *ListDatasetsRequest) GetDataSourceTypes() *string {
	return s.DataSourceTypes
}

func (s *ListDatasetsRequest) GetDataTypes() *string {
	return s.DataTypes
}

func (s *ListDatasetsRequest) GetDatasetIds() *string {
	return s.DatasetIds
}

func (s *ListDatasetsRequest) GetEdition() *string {
	return s.Edition
}

func (s *ListDatasetsRequest) GetLabel() *string {
	return s.Label
}

func (s *ListDatasetsRequest) GetName() *string {
	return s.Name
}

func (s *ListDatasetsRequest) GetOrder() *string {
	return s.Order
}

func (s *ListDatasetsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListDatasetsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListDatasetsRequest) GetProperties() *string {
	return s.Properties
}

func (s *ListDatasetsRequest) GetProvider() *string {
	return s.Provider
}

func (s *ListDatasetsRequest) GetShareScope() *string {
	return s.ShareScope
}

func (s *ListDatasetsRequest) GetSortBy() *string {
	return s.SortBy
}

func (s *ListDatasetsRequest) GetSourceDatasetId() *string {
	return s.SourceDatasetId
}

func (s *ListDatasetsRequest) GetSourceId() *string {
	return s.SourceId
}

func (s *ListDatasetsRequest) GetSourceTypes() *string {
	return s.SourceTypes
}

func (s *ListDatasetsRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *ListDatasetsRequest) SetAccessibility(v string) *ListDatasetsRequest {
	s.Accessibility = &v
	return s
}

func (s *ListDatasetsRequest) SetDataSourceTypes(v string) *ListDatasetsRequest {
	s.DataSourceTypes = &v
	return s
}

func (s *ListDatasetsRequest) SetDataTypes(v string) *ListDatasetsRequest {
	s.DataTypes = &v
	return s
}

func (s *ListDatasetsRequest) SetDatasetIds(v string) *ListDatasetsRequest {
	s.DatasetIds = &v
	return s
}

func (s *ListDatasetsRequest) SetEdition(v string) *ListDatasetsRequest {
	s.Edition = &v
	return s
}

func (s *ListDatasetsRequest) SetLabel(v string) *ListDatasetsRequest {
	s.Label = &v
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

func (s *ListDatasetsRequest) SetPageNumber(v int32) *ListDatasetsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListDatasetsRequest) SetPageSize(v int32) *ListDatasetsRequest {
	s.PageSize = &v
	return s
}

func (s *ListDatasetsRequest) SetProperties(v string) *ListDatasetsRequest {
	s.Properties = &v
	return s
}

func (s *ListDatasetsRequest) SetProvider(v string) *ListDatasetsRequest {
	s.Provider = &v
	return s
}

func (s *ListDatasetsRequest) SetShareScope(v string) *ListDatasetsRequest {
	s.ShareScope = &v
	return s
}

func (s *ListDatasetsRequest) SetSortBy(v string) *ListDatasetsRequest {
	s.SortBy = &v
	return s
}

func (s *ListDatasetsRequest) SetSourceDatasetId(v string) *ListDatasetsRequest {
	s.SourceDatasetId = &v
	return s
}

func (s *ListDatasetsRequest) SetSourceId(v string) *ListDatasetsRequest {
	s.SourceId = &v
	return s
}

func (s *ListDatasetsRequest) SetSourceTypes(v string) *ListDatasetsRequest {
	s.SourceTypes = &v
	return s
}

func (s *ListDatasetsRequest) SetWorkspaceId(v string) *ListDatasetsRequest {
	s.WorkspaceId = &v
	return s
}

func (s *ListDatasetsRequest) Validate() error {
	return dara.Validate(s)
}
