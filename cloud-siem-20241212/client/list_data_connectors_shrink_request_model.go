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
	// The list of data connector IDs.
	DataConnectorIdsShrink *string `json:"DataConnectorIds,omitempty" xml:"DataConnectorIds,omitempty"`
	// The name of the data connector.
	//
	// example:
	//
	// SAS-CTDR-2026070210****
	DataConnectorName *string `json:"DataConnectorName,omitempty" xml:"DataConnectorName,omitempty"`
	// The enabled status of the data connector. Valid values:
	//
	// - "enabled": Enabled.
	//
	// - "disabled" (default): Disabled.
	//
	// example:
	//
	// enabled
	DataConnectorStatus *string `json:"DataConnectorStatus,omitempty" xml:"DataConnectorStatus,omitempty"`
	// The type of the data connector. Valid values:
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
	// The ID of the destination data source. This parameter is required only for synchronization.
	//
	// example:
	//
	// ds-5sfe68t122pxnti1cjpl
	DestDataSourceId *string `json:"DestDataSourceId,omitempty" xml:"DestDataSourceId,omitempty"`
	// The language of the response message. Valid values:
	//
	// - **zh*	- (default): Chinese.
	//
	// - **en**: English.
	//
	// Note: This parameter is required in actual API calls. If this parameter is not provided, the service returns a 400 error with the message "The specified Lang can not be empty." Explicitly pass the language code.
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The maximum number of entries to return for a single query when you use the NextToken method. Valid values: 1 to 100. Default value: 50.
	//
	// example:
	//
	// 50
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The token used to query the next page. You do not need to specify this parameter for the first query or when no next page exists. If a next page exists, set this parameter to the NextToken value returned by the previous API call.
	//
	// example:
	//
	// AAAAAUqcj6VO4E3ECWIrFczs****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The field used for sorting.
	//
	// Currently, only sorting by updateTime is supported.
	//
	// If OrderField is empty, the default order returned by the database is used.
	//
	// example:
	//
	// Id
	OrderField *string `json:"OrderField,omitempty" xml:"OrderField,omitempty"`
	// The sort order. Valid values:
	//
	// - "asc"
	//
	// - "desc" (default)
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
	// The number of entries per page. Default value: 1000.
	//
	// example:
	//
	// 1
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The region where the data management center for threat analysis is located. Select the region of the management center based on the region where your assets are located. Valid values:
	//
	// - cn-hangzhou: The assets are located in the Chinese mainland.
	//
	// - ap-southeast-1: The assets are located outside the Chinese mainland.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the user whose perspective the administrator assumes.
	//
	// example:
	//
	// 113091674488****
	RoleFor *int64 `json:"RoleFor,omitempty" xml:"RoleFor,omitempty"`
	// The name of the Simple Log Service (SLS) data import job for the data connector.
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
