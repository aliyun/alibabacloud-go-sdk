// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDataConnectorsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDataConnectorIds(v []*string) *ListDataConnectorsRequest
	GetDataConnectorIds() []*string
	SetDataConnectorName(v string) *ListDataConnectorsRequest
	GetDataConnectorName() *string
	SetDataConnectorStatus(v string) *ListDataConnectorsRequest
	GetDataConnectorStatus() *string
	SetDataConnectorType(v string) *ListDataConnectorsRequest
	GetDataConnectorType() *string
	SetDestDataSourceId(v string) *ListDataConnectorsRequest
	GetDestDataSourceId() *string
	SetLang(v string) *ListDataConnectorsRequest
	GetLang() *string
	SetMaxResults(v int32) *ListDataConnectorsRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListDataConnectorsRequest
	GetNextToken() *string
	SetOrderField(v string) *ListDataConnectorsRequest
	GetOrderField() *string
	SetOrderType(v string) *ListDataConnectorsRequest
	GetOrderType() *string
	SetPageNumber(v int32) *ListDataConnectorsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListDataConnectorsRequest
	GetPageSize() *int32
	SetRegionId(v string) *ListDataConnectorsRequest
	GetRegionId() *string
	SetRoleFor(v int64) *ListDataConnectorsRequest
	GetRoleFor() *int64
	SetSlsIngestionJobName(v string) *ListDataConnectorsRequest
	GetSlsIngestionJobName() *string
	SetSrcDataType(v string) *ListDataConnectorsRequest
	GetSrcDataType() *string
}

type ListDataConnectorsRequest struct {
	// The list of data connector IDs.
	DataConnectorIds []*string `json:"DataConnectorIds,omitempty" xml:"DataConnectorIds,omitempty" type:"Repeated"`
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

func (s ListDataConnectorsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListDataConnectorsRequest) GoString() string {
	return s.String()
}

func (s *ListDataConnectorsRequest) GetDataConnectorIds() []*string {
	return s.DataConnectorIds
}

func (s *ListDataConnectorsRequest) GetDataConnectorName() *string {
	return s.DataConnectorName
}

func (s *ListDataConnectorsRequest) GetDataConnectorStatus() *string {
	return s.DataConnectorStatus
}

func (s *ListDataConnectorsRequest) GetDataConnectorType() *string {
	return s.DataConnectorType
}

func (s *ListDataConnectorsRequest) GetDestDataSourceId() *string {
	return s.DestDataSourceId
}

func (s *ListDataConnectorsRequest) GetLang() *string {
	return s.Lang
}

func (s *ListDataConnectorsRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListDataConnectorsRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListDataConnectorsRequest) GetOrderField() *string {
	return s.OrderField
}

func (s *ListDataConnectorsRequest) GetOrderType() *string {
	return s.OrderType
}

func (s *ListDataConnectorsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListDataConnectorsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListDataConnectorsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListDataConnectorsRequest) GetRoleFor() *int64 {
	return s.RoleFor
}

func (s *ListDataConnectorsRequest) GetSlsIngestionJobName() *string {
	return s.SlsIngestionJobName
}

func (s *ListDataConnectorsRequest) GetSrcDataType() *string {
	return s.SrcDataType
}

func (s *ListDataConnectorsRequest) SetDataConnectorIds(v []*string) *ListDataConnectorsRequest {
	s.DataConnectorIds = v
	return s
}

func (s *ListDataConnectorsRequest) SetDataConnectorName(v string) *ListDataConnectorsRequest {
	s.DataConnectorName = &v
	return s
}

func (s *ListDataConnectorsRequest) SetDataConnectorStatus(v string) *ListDataConnectorsRequest {
	s.DataConnectorStatus = &v
	return s
}

func (s *ListDataConnectorsRequest) SetDataConnectorType(v string) *ListDataConnectorsRequest {
	s.DataConnectorType = &v
	return s
}

func (s *ListDataConnectorsRequest) SetDestDataSourceId(v string) *ListDataConnectorsRequest {
	s.DestDataSourceId = &v
	return s
}

func (s *ListDataConnectorsRequest) SetLang(v string) *ListDataConnectorsRequest {
	s.Lang = &v
	return s
}

func (s *ListDataConnectorsRequest) SetMaxResults(v int32) *ListDataConnectorsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListDataConnectorsRequest) SetNextToken(v string) *ListDataConnectorsRequest {
	s.NextToken = &v
	return s
}

func (s *ListDataConnectorsRequest) SetOrderField(v string) *ListDataConnectorsRequest {
	s.OrderField = &v
	return s
}

func (s *ListDataConnectorsRequest) SetOrderType(v string) *ListDataConnectorsRequest {
	s.OrderType = &v
	return s
}

func (s *ListDataConnectorsRequest) SetPageNumber(v int32) *ListDataConnectorsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListDataConnectorsRequest) SetPageSize(v int32) *ListDataConnectorsRequest {
	s.PageSize = &v
	return s
}

func (s *ListDataConnectorsRequest) SetRegionId(v string) *ListDataConnectorsRequest {
	s.RegionId = &v
	return s
}

func (s *ListDataConnectorsRequest) SetRoleFor(v int64) *ListDataConnectorsRequest {
	s.RoleFor = &v
	return s
}

func (s *ListDataConnectorsRequest) SetSlsIngestionJobName(v string) *ListDataConnectorsRequest {
	s.SlsIngestionJobName = &v
	return s
}

func (s *ListDataConnectorsRequest) SetSrcDataType(v string) *ListDataConnectorsRequest {
	s.SrcDataType = &v
	return s
}

func (s *ListDataConnectorsRequest) Validate() error {
	return dara.Validate(s)
}
