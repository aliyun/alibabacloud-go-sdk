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
	Accessibility *string `json:"Accessibility,omitempty" xml:"Accessibility,omitempty"`
	// The storage types of the data source. Multiple data source types are separated by commas (,). Valid values:
	//
	// 	- NAS: File Storage NAS (NAS).
	//
	// 	- OSS: Object Storage Service (OSS).
	//
	// example:
	//
	// OSS
	DataSourceTypes *string `json:"DataSourceTypes,omitempty" xml:"DataSourceTypes,omitempty"`
	// The dataset types. Multiple dataset types are separated by commas (,). Valid values:
	//
	// 	- Video: video
	//
	// 	- COMMON: common
	//
	// 	- TEXT: text
	//
	// 	- PIC: picture
	//
	// 	- AUDIO: audio
	//
	// example:
	//
	// COMMON,TEXT
	DataTypes *string `json:"DataTypes,omitempty" xml:"DataTypes,omitempty"`
	Edition   *string `json:"Edition,omitempty" xml:"Edition,omitempty"`
	// The dataset tag, which is used to filter datasets. Datasets whose tag key or tag value contains a specified string are filtered.
	//
	// example:
	//
	// test
	Label *string `json:"Label,omitempty" xml:"Label,omitempty"`
	// The dataset name. Fuzzy search based on the dataset name is supported.
	//
	// example:
	//
	// myName
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The order of specific fields of the entries on the returned page. Valid values: ASC and DESC. Default value: ASC.
	//
	// 	- ASC: The entries are sorted in ascending order.
	//
	// 	- DESC: The entries are sorted in descending order.
	//
	// example:
	//
	// ASC
	Order *string `json:"Order,omitempty" xml:"Order,omitempty"`
	// The page number. Pages start from page 1. Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Default value: 10.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The dataset properties. Multiple properties are separated by commas (,). Valid values:
	//
	// 	- DIRECTORY
	//
	// 	- FILE
	//
	// example:
	//
	// FILE
	Properties *string `json:"Properties,omitempty" xml:"Properties,omitempty"`
	// The dataset provider. If the value pai is returned, the dataset is a public dataset provided by PAI.
	//
	// example:
	//
	// pai
	Provider   *string `json:"Provider,omitempty" xml:"Provider,omitempty"`
	ShareScope *string `json:"ShareScope,omitempty" xml:"ShareScope,omitempty"`
	// The field used for sorting.
	//
	// example:
	//
	// GmtCreateTime
	SortBy *string `json:"SortBy,omitempty" xml:"SortBy,omitempty"`
	// The ID of the iTAG labeled dataset that is used as the source dataset.
	//
	// example:
	//
	// d-rcdg3wxxxxxhc5jk87
	SourceDatasetId *string `json:"SourceDatasetId,omitempty" xml:"SourceDatasetId,omitempty"`
	// The data source ID.
	//
	// 	- If SourceType is set to USER, the value of SourceId is a custom string.
	//
	// 	- If SourceType is set to ITAG, the value of SourceId is the ID of the labeling job of iTAG.
	//
	// 	- If SourceType is set to PAI_PUBLIC_DATASET, SourceId is empty by default.
	//
	// example:
	//
	// d-rbvg5wzljzjhc9ks92
	SourceId *string `json:"SourceId,omitempty" xml:"SourceId,omitempty"`
	// The source types. Multiple source types are separated by commas (,). Valid values:
	//
	// 	- PAI-PUBLIC-DATASET: a public dataset of Platform for AI (PAI).
	//
	// 	- ITAG: a dataset generated from a labeling job of iTAG.
	//
	// 	- USER: a dataset registered by a user.
	//
	// example:
	//
	// USER,ITAG
	SourceTypes *string `json:"SourceTypes,omitempty" xml:"SourceTypes,omitempty"`
	// The ID of the workspace to which the dataset belongs. You can call [ListWorkspaces](https://help.aliyun.com/document_detail/449124.html) to obtain the workspace ID. If you do not specify this parameter, the default workspace is used. If the default workspace does not exist, an error is reported.
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
