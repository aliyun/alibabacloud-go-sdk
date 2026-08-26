// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDatasetVersionsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLabelKeys(v string) *ListDatasetVersionsRequest
	GetLabelKeys() *string
	SetLabelValues(v string) *ListDatasetVersionsRequest
	GetLabelValues() *string
	SetOrder(v string) *ListDatasetVersionsRequest
	GetOrder() *string
	SetPageNumber(v int32) *ListDatasetVersionsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListDatasetVersionsRequest
	GetPageSize() *int32
	SetProperties(v string) *ListDatasetVersionsRequest
	GetProperties() *string
	SetSortBy(v string) *ListDatasetVersionsRequest
	GetSortBy() *string
	SetSourceId(v string) *ListDatasetVersionsRequest
	GetSourceId() *string
	SetSourceTypes(v string) *ListDatasetVersionsRequest
	GetSourceTypes() *string
}

type ListDatasetVersionsRequest struct {
	// The dataset label used to filter the dataset list. Datasets whose label key or value contains the specified string are returned.
	//
	// example:
	//
	// key1,key2
	LabelKeys *string `json:"LabelKeys,omitempty" xml:"LabelKeys,omitempty"`
	// The dataset label used to filter the dataset list. Datasets whose label key or value contains the specified string are returned.
	//
	// example:
	//
	// value1,value2
	LabelValues *string `json:"LabelValues,omitempty" xml:"LabelValues,omitempty"`
	// The order in which entries are sorted by the specified field in a paged query. Default value: ASC.
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
	// This parameter is required.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page for a paged query. Default value: 10.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The dataset property. Valid values:
	//
	// - DIRECTORY: folder.
	//
	// - FILE: file.
	//
	// example:
	//
	// DIRECTORY
	Properties *string `json:"Properties,omitempty" xml:"Properties,omitempty"`
	// The field by which entries are sorted in a paged query. Default value: GmtCreateTime. Valid values:
	//
	// - GmtCreateTime (default): sort by creation time.
	//
	// - GmtModifiedTime: sort by modification time.
	//
	// - SourceType
	//
	// - DataSourceType
	//
	// - Property
	//
	// - DataSize
	//
	// - DataCount
	//
	// example:
	//
	// GmtCreateTime
	SortBy *string `json:"SortBy,omitempty" xml:"SortBy,omitempty"`
	// The data source ID.
	//
	// - If SourceTypes is set to USER, SourceId can be customized.
	//
	// - If SourceTypes is set to ITAG, which indicates a dataset generated from iTAG labeling results, SourceId is the iTAG task ID.
	//
	// - If SourceTypes is set to PAI_PUBLIC_DATASET, which indicates a dataset created from a PAI public dataset, SourceId is empty by default.
	//
	// example:
	//
	// d-a0xbe5n03bhqof46ce
	SourceId *string `json:"SourceId,omitempty" xml:"SourceId,omitempty"`
	// The source type. Valid values:
	//
	// - PAI-PUBLIC-DATASET: PAI public dataset.
	//
	// - ITAG: dataset generated from iTAG labeling results.
	//
	// - USER: dataset registered by a user.
	//
	// example:
	//
	// USER
	SourceTypes *string `json:"SourceTypes,omitempty" xml:"SourceTypes,omitempty"`
}

func (s ListDatasetVersionsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListDatasetVersionsRequest) GoString() string {
	return s.String()
}

func (s *ListDatasetVersionsRequest) GetLabelKeys() *string {
	return s.LabelKeys
}

func (s *ListDatasetVersionsRequest) GetLabelValues() *string {
	return s.LabelValues
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

func (s *ListDatasetVersionsRequest) GetProperties() *string {
	return s.Properties
}

func (s *ListDatasetVersionsRequest) GetSortBy() *string {
	return s.SortBy
}

func (s *ListDatasetVersionsRequest) GetSourceId() *string {
	return s.SourceId
}

func (s *ListDatasetVersionsRequest) GetSourceTypes() *string {
	return s.SourceTypes
}

func (s *ListDatasetVersionsRequest) SetLabelKeys(v string) *ListDatasetVersionsRequest {
	s.LabelKeys = &v
	return s
}

func (s *ListDatasetVersionsRequest) SetLabelValues(v string) *ListDatasetVersionsRequest {
	s.LabelValues = &v
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

func (s *ListDatasetVersionsRequest) SetProperties(v string) *ListDatasetVersionsRequest {
	s.Properties = &v
	return s
}

func (s *ListDatasetVersionsRequest) SetSortBy(v string) *ListDatasetVersionsRequest {
	s.SortBy = &v
	return s
}

func (s *ListDatasetVersionsRequest) SetSourceId(v string) *ListDatasetVersionsRequest {
	s.SourceId = &v
	return s
}

func (s *ListDatasetVersionsRequest) SetSourceTypes(v string) *ListDatasetVersionsRequest {
	s.SourceTypes = &v
	return s
}

func (s *ListDatasetVersionsRequest) Validate() error {
	return dara.Validate(s)
}
