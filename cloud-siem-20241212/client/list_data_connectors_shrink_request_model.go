// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDataConnectorsShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDataConnectorIdsShrink(v string) *ListDataConnectorsShrinkRequest
	GetDataConnectorIdsShrink() *string
	SetDataConnectorName(v string) *ListDataConnectorsShrinkRequest
	GetDataConnectorName() *string
	SetDataConnectorStatus(v string) *ListDataConnectorsShrinkRequest
	GetDataConnectorStatus() *string
	SetDataConnectorType(v string) *ListDataConnectorsShrinkRequest
	GetDataConnectorType() *string
	SetDestDataSourceId(v string) *ListDataConnectorsShrinkRequest
	GetDestDataSourceId() *string
	SetLang(v string) *ListDataConnectorsShrinkRequest
	GetLang() *string
	SetMaxResults(v int32) *ListDataConnectorsShrinkRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListDataConnectorsShrinkRequest
	GetNextToken() *string
	SetOrderField(v string) *ListDataConnectorsShrinkRequest
	GetOrderField() *string
	SetOrderType(v string) *ListDataConnectorsShrinkRequest
	GetOrderType() *string
	SetPageNumber(v int32) *ListDataConnectorsShrinkRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListDataConnectorsShrinkRequest
	GetPageSize() *int32
	SetRegionId(v string) *ListDataConnectorsShrinkRequest
	GetRegionId() *string
	SetRoleFor(v int64) *ListDataConnectorsShrinkRequest
	GetRoleFor() *int64
	SetSlsIngestionJobName(v string) *ListDataConnectorsShrinkRequest
	GetSlsIngestionJobName() *string
	SetSrcDataType(v string) *ListDataConnectorsShrinkRequest
	GetSrcDataType() *string
}

type ListDataConnectorsShrinkRequest struct {
	// The list of collector IDs.
	DataConnectorIdsShrink *string `json:"DataConnectorIds,omitempty" xml:"DataConnectorIds,omitempty"`
	// The collector name.
	//
	// example:
	//
	// SAS-CTDR-2026070210****
	DataConnectorName *string `json:"DataConnectorName,omitempty" xml:"DataConnectorName,omitempty"`
	// The collector status. Valid values:
	//
	// - "enabled": enabled.
	//
	// - "disabled" (default): disabled.
	//
	// example:
	//
	// enabled
	DataConnectorStatus *string `json:"DataConnectorStatus,omitempty" xml:"DataConnectorStatus,omitempty"`
	// The collector type. Valid values:
	//
	// - oss
	//
	// - s3
	//
	// - kafka
	//
	// example:
	//
	// s3
	DataConnectorType *string `json:"DataConnectorType,omitempty" xml:"DataConnectorType,omitempty"`
	// The destination data source ID. This parameter is required only for synchronization.
	//
	// example:
	//
	// ds-5sfe68t122pxnti1cjpl
	DestDataSourceId *string `json:"DestDataSourceId,omitempty" xml:"DestDataSourceId,omitempty"`
	// The language of the response. Valid values:
	//
	// - **zh*	- (default): Chinese.
	//
	// - **en**: English.
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The maximum number of results to return when you use the NextToken-based pagination method. Valid values: 1 to 100. Default value: 50.
	//
	// example:
	//
	// 50
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The pagination token. You do not need to set this parameter for the first request or if no more results exist. If more results exist, set this parameter to the NextToken value returned by the previous API call.
	//
	// example:
	//
	// AAAAAUqcj6VO4E3ECWIrFczs****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The sort field. Currently, only sorting by updateTime is supported. If OrderField is left empty, the default order returned by the database is used.
	//
	// example:
	//
	// Id
	OrderField *string `json:"OrderField,omitempty" xml:"OrderField,omitempty"`
	// The sort order. Valid values:
	//
	// - "asc": ascending order.
	//
	// - "desc" (default): descending order.
	//
	// example:
	//
	// desc
	OrderType *string `json:"OrderType,omitempty" xml:"OrderType,omitempty"`
	// The current page number. Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of records per page. Default value: 1000.
	//
	// example:
	//
	// 1
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The region where the threat analysis data management center resides. Specify the management center region based on the region of your assets. Valid values:
	//
	// - cn-hangzhou: The assets reside in the Chinese mainland.
	//
	// - ap-southeast-1: The assets reside outside China.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The user ID that the administrator switches to when viewing as another member.
	//
	// example:
	//
	// 113091674488****
	RoleFor *int64 `json:"RoleFor,omitempty" xml:"RoleFor,omitempty"`
	// The name of the Simple Log Service (SLS) data import job for the collector.
	//
	// example:
	//
	// ingest-oss-dc-1a2b3c4d5e6f7a8****
	SlsIngestionJobName *string `json:"SlsIngestionJobName,omitempty" xml:"SlsIngestionJobName,omitempty"`
	// The source data type.
	//
	// example:
	//
	// s3
	SrcDataType *string `json:"SrcDataType,omitempty" xml:"SrcDataType,omitempty"`
}

func (s ListDataConnectorsShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ListDataConnectorsShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListDataConnectorsShrinkRequest) GetDataConnectorIdsShrink() *string {
	return s.DataConnectorIdsShrink
}

func (s *ListDataConnectorsShrinkRequest) GetDataConnectorName() *string {
	return s.DataConnectorName
}

func (s *ListDataConnectorsShrinkRequest) GetDataConnectorStatus() *string {
	return s.DataConnectorStatus
}

func (s *ListDataConnectorsShrinkRequest) GetDataConnectorType() *string {
	return s.DataConnectorType
}

func (s *ListDataConnectorsShrinkRequest) GetDestDataSourceId() *string {
	return s.DestDataSourceId
}

func (s *ListDataConnectorsShrinkRequest) GetLang() *string {
	return s.Lang
}

func (s *ListDataConnectorsShrinkRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListDataConnectorsShrinkRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListDataConnectorsShrinkRequest) GetOrderField() *string {
	return s.OrderField
}

func (s *ListDataConnectorsShrinkRequest) GetOrderType() *string {
	return s.OrderType
}

func (s *ListDataConnectorsShrinkRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListDataConnectorsShrinkRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListDataConnectorsShrinkRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListDataConnectorsShrinkRequest) GetRoleFor() *int64 {
	return s.RoleFor
}

func (s *ListDataConnectorsShrinkRequest) GetSlsIngestionJobName() *string {
	return s.SlsIngestionJobName
}

func (s *ListDataConnectorsShrinkRequest) GetSrcDataType() *string {
	return s.SrcDataType
}

func (s *ListDataConnectorsShrinkRequest) SetDataConnectorIdsShrink(v string) *ListDataConnectorsShrinkRequest {
	s.DataConnectorIdsShrink = &v
	return s
}

func (s *ListDataConnectorsShrinkRequest) SetDataConnectorName(v string) *ListDataConnectorsShrinkRequest {
	s.DataConnectorName = &v
	return s
}

func (s *ListDataConnectorsShrinkRequest) SetDataConnectorStatus(v string) *ListDataConnectorsShrinkRequest {
	s.DataConnectorStatus = &v
	return s
}

func (s *ListDataConnectorsShrinkRequest) SetDataConnectorType(v string) *ListDataConnectorsShrinkRequest {
	s.DataConnectorType = &v
	return s
}

func (s *ListDataConnectorsShrinkRequest) SetDestDataSourceId(v string) *ListDataConnectorsShrinkRequest {
	s.DestDataSourceId = &v
	return s
}

func (s *ListDataConnectorsShrinkRequest) SetLang(v string) *ListDataConnectorsShrinkRequest {
	s.Lang = &v
	return s
}

func (s *ListDataConnectorsShrinkRequest) SetMaxResults(v int32) *ListDataConnectorsShrinkRequest {
	s.MaxResults = &v
	return s
}

func (s *ListDataConnectorsShrinkRequest) SetNextToken(v string) *ListDataConnectorsShrinkRequest {
	s.NextToken = &v
	return s
}

func (s *ListDataConnectorsShrinkRequest) SetOrderField(v string) *ListDataConnectorsShrinkRequest {
	s.OrderField = &v
	return s
}

func (s *ListDataConnectorsShrinkRequest) SetOrderType(v string) *ListDataConnectorsShrinkRequest {
	s.OrderType = &v
	return s
}

func (s *ListDataConnectorsShrinkRequest) SetPageNumber(v int32) *ListDataConnectorsShrinkRequest {
	s.PageNumber = &v
	return s
}

func (s *ListDataConnectorsShrinkRequest) SetPageSize(v int32) *ListDataConnectorsShrinkRequest {
	s.PageSize = &v
	return s
}

func (s *ListDataConnectorsShrinkRequest) SetRegionId(v string) *ListDataConnectorsShrinkRequest {
	s.RegionId = &v
	return s
}

func (s *ListDataConnectorsShrinkRequest) SetRoleFor(v int64) *ListDataConnectorsShrinkRequest {
	s.RoleFor = &v
	return s
}

func (s *ListDataConnectorsShrinkRequest) SetSlsIngestionJobName(v string) *ListDataConnectorsShrinkRequest {
	s.SlsIngestionJobName = &v
	return s
}

func (s *ListDataConnectorsShrinkRequest) SetSrcDataType(v string) *ListDataConnectorsShrinkRequest {
	s.SrcDataType = &v
	return s
}

func (s *ListDataConnectorsShrinkRequest) Validate() error {
	return dara.Validate(s)
}
