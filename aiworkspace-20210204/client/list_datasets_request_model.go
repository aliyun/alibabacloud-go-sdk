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
	// Specifies the dataset\\"s visibility.
	//
	// - `PUBLIC`: The dataset is publicly accessible.
	//
	// - `PRIVATE`: The dataset is privately accessible.
	//
	// example:
	//
	// PRIVATE
	Accessibility *string `json:"Accessibility,omitempty" xml:"Accessibility,omitempty"`
	// The data source type. To specify multiple types, separate them with commas (,). Valid values:
	//
	// - `NAS`: The data source is NAS.
	//
	// - `OSS`: The data source is OSS.
	//
	// example:
	//
	// OSS
	DataSourceTypes *string `json:"DataSourceTypes,omitempty" xml:"DataSourceTypes,omitempty"`
	// The data type of the dataset. To specify multiple data types, separate them with commas (,). Valid values:
	//
	// - `VIDEO`: video.
	//
	// - `COMMON`: general.
	//
	// - `TEXT`: text.
	//
	// - `PIC`: image.
	//
	// - `AUDIO`: audio.
	//
	// example:
	//
	// COMMON,TEXT
	DataTypes *string `json:"DataTypes,omitempty" xml:"DataTypes,omitempty"`
	// A comma-separated list of dataset IDs.
	//
	// example:
	//
	// d-rcdg3wxxxxxhc5jk87
	DatasetIds *string `json:"DatasetIds,omitempty" xml:"DatasetIds,omitempty"`
	// The dataset edition. Valid values:
	//
	// - `BASIC`: Basic edition. Does not support file metadata management.
	//
	// - `ADVANCED`: Advanced edition. This edition is supported only for OSS datasets. Each version can manage metadata for up to 1 million files.
	//
	// - `LOGICAL`: Logical edition. This edition is supported only for OSS datasets and is suitable for most use cases. Each version can manage metadata for up to 1 million files. You must use an SDK with this edition.
	//
	// example:
	//
	// BASIC
	Edition *string `json:"Edition,omitempty" xml:"Edition,omitempty"`
	// A label used to filter datasets. The operation returns datasets whose label key or value contains the specified string.
	//
	// example:
	//
	// test
	Label *string `json:"Label,omitempty" xml:"Label,omitempty"`
	// The dataset name. Fuzzy search is supported.
	//
	// example:
	//
	// myName
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The sort order for the results, based on the `SortBy` parameter. The default is `ASC`.
	//
	// - `ASC`: ascending order.
	//
	// - `DESC`: descending order.
	//
	// example:
	//
	// ASC
	Order *string `json:"Order,omitempty" xml:"Order,omitempty"`
	// The page number for the paged query. Starts at 1. The default is 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of datasets to return per page. The default is 10.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The dataset properties. To specify multiple properties, separate them with commas (,). Valid values:
	//
	// - `DIRECTORY`: A folder.
	//
	// - `FILE`: A file.
	//
	// example:
	//
	// FILE
	Properties *string `json:"Properties,omitempty" xml:"Properties,omitempty"`
	// The dataset provider. Set this parameter to `pai` to query public datasets on the PAI platform.
	//
	// example:
	//
	// pai
	Provider *string `json:"Provider,omitempty" xml:"Provider,omitempty"`
	// A filter for shared datasets.
	//
	// - `TO_ME`: Returns only datasets shared with you.
	//
	// - `BY_ME`: Returns only datasets that you have shared with others and displays details of the sharing configuration.
	//
	// - If this parameter is omitted or empty, the operation returns all datasets in the current workspace, including those shared with you.
	//
	// example:
	//
	// BY_ME
	ShareScope *string `json:"ShareScope,omitempty" xml:"ShareScope,omitempty"`
	// The sort field.
	//
	// example:
	//
	// GmtCreateTime
	SortBy *string `json:"SortBy,omitempty" xml:"SortBy,omitempty"`
	// The source dataset ID for an iTAG annotation set.
	//
	// example:
	//
	// d-rcdg3wxxxxxhc5jk87
	SourceDatasetId *string `json:"SourceDatasetId,omitempty" xml:"SourceDatasetId,omitempty"`
	// The source ID. The value of this parameter varies based on the `SourceTypes` value:
	//
	// - If `SourceTypes` is `USER`, you can specify a custom value for `SourceId`.
	//
	// - If `SourceTypes` is `ITAG`, `SourceId` is the ID of the iTAG task.
	//
	// - If `SourceTypes` is `PAI_PUBLIC_DATASET`, this parameter is empty by default.
	//
	// example:
	//
	// d-rbvg5wzljzjhc9ks92
	SourceId *string `json:"SourceId,omitempty" xml:"SourceId,omitempty"`
	// The source type. To specify multiple types, separate them with commas (,).
	//
	// example:
	//
	// USER,ITAG
	SourceTypes *string `json:"SourceTypes,omitempty" xml:"SourceTypes,omitempty"`
	// The ID of the workspace that contains the dataset. For information about how to obtain the workspace ID, see [ListWorkspaces](https://help.aliyun.com/document_detail/449124.html).
	//
	// If this parameter is not specified, the default workspace is used. An error is returned if the default workspace does not exist.
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
