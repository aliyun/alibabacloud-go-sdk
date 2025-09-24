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
	// The dataset tag keys, which are used to filter datasets. Datasets whose tag keys or tag values contain a specified string are filtered.
	//
	// example:
	//
	// key1,key2
	LabelKeys *string `json:"LabelKeys,omitempty" xml:"LabelKeys,omitempty"`
	// The dataset tag values, which are used to filter datasets. Datasets whose tag keys or tag values contain a specified string are filtered.
	//
	// example:
	//
	// value1,value2
	LabelValues *string `json:"LabelValues,omitempty" xml:"LabelValues,omitempty"`
	// The order in which the entries are sorted by the specific field on the returned page. Default value: ASC. Valid values:
	//
	// 	- ASC: ascending order
	//
	// 	- DESC: descending order.
	//
	// example:
	//
	// ASC
	Order *string `json:"Order,omitempty" xml:"Order,omitempty"`
	// The page number. Pages start from page 1. Default value: 1.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Default value: 10.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The dataset properties. Valid values:
	//
	// 	- DIRECTORY
	//
	// 	- FILE
	//
	// example:
	//
	// DIRECTORY
	Properties *string `json:"Properties,omitempty" xml:"Properties,omitempty"`
	// The field used to sort the results in queries by page. Default value: GmtCreateTime.
	//
	// Valid values:
	//
	// 	- SourceType
	//
	// 	- DataSourceType
	//
	// 	- DataSize
	//
	// 	- DataCount
	//
	// 	- Property
	//
	// 	- GmtCreateTime: The results are sorted by creation time. This is the default value.
	//
	// 	- GmtModifiedTime: The results are sorted by modification time.
	//
	// 	- DatasetId
	//
	// example:
	//
	// GmtCreateTime
	SortBy *string `json:"SortBy,omitempty" xml:"SortBy,omitempty"`
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
	// d-a0xbe5n03bhqof46ce
	SourceId *string `json:"SourceId,omitempty" xml:"SourceId,omitempty"`
	// The source type. Valid values:
	//
	// 	- PAI-PUBLIC-DATASET: a public dataset of Platform for AI (PAI).
	//
	// 	- ITAG: a dataset generated from a labeling job of iTAG.
	//
	// 	- USER: a dataset registered by a user.
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
