// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDataSourcesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConnectStatus(v string) *ListDataSourcesRequest
	GetConnectStatus() *string
	SetCurrentPage(v int32) *ListDataSourcesRequest
	GetCurrentPage() *int32
	SetDataAssetId(v string) *ListDataSourcesRequest
	GetDataAssetId() *string
	SetDataSourceId(v string) *ListDataSourcesRequest
	GetDataSourceId() *string
	SetDbName(v string) *ListDataSourcesRequest
	GetDbName() *string
	SetEngineType(v string) *ListDataSourcesRequest
	GetEngineType() *string
	SetIdentifyStatus(v string) *ListDataSourcesRequest
	GetIdentifyStatus() *string
	SetInstanceId(v string) *ListDataSourcesRequest
	GetInstanceId() *string
	SetLang(v string) *ListDataSourcesRequest
	GetLang() *string
	SetMaxResults(v int32) *ListDataSourcesRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListDataSourcesRequest
	GetNextToken() *string
	SetPageSize(v int32) *ListDataSourcesRequest
	GetPageSize() *int32
	SetPort(v int32) *ListDataSourcesRequest
	GetPort() *int32
	SetProductCode(v string) *ListDataSourcesRequest
	GetProductCode() *string
	SetProductId(v int64) *ListDataSourcesRequest
	GetProductId() *int64
	SetSourceIp(v string) *ListDataSourcesRequest
	GetSourceIp() *string
}

type ListDataSourcesRequest struct {
	// The connection status of the data source. Valid values:
	//
	// - UnConnected: Not connected.
	//
	// - Processing: Connecting.
	//
	// - Success: Connected.
	//
	// - Failed: Connection failed.
	//
	// example:
	//
	// Success
	ConnectStatus *string `json:"ConnectStatus,omitempty" xml:"ConnectStatus,omitempty"`
	// The current page number. Pages start from 1.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The identifier of the data asset in Data Security Center.
	//
	// example:
	//
	// asset-example-001
	DataAssetId *string `json:"DataAssetId,omitempty" xml:"DataAssetId,omitempty"`
	// The identifier of the data source in Data Security Center. You can obtain the value from the Items[].DataSourceId parameter in the response of the ListDataSources operation.
	//
	// example:
	//
	// ds-example-001
	DataSourceId *string `json:"DataSourceId,omitempty" xml:"DataSourceId,omitempty"`
	// The name of the database.
	//
	// example:
	//
	// business_db
	DbName *string `json:"DbName,omitempty" xml:"DbName,omitempty"`
	// The engine type. Valid values:
	//
	// - **MySQL**
	//
	// - **MariaDB**
	//
	// - **Oracle**
	//
	// - **PostgreSQL**
	//
	// - **SQLServer**
	//
	// example:
	//
	// MySQL
	EngineType *string `json:"EngineType,omitempty" xml:"EngineType,omitempty"`
	// The status of the data source detection switch. Valid values:
	//
	// - Enable: Submitted for enabling.
	//
	// - Enabling: Enabling.
	//
	// - Enabled: Enabled.
	//
	// - EnableFailed: Enabling failed.
	//
	// - Disable: Submitted for disabling.
	//
	// - Disabling: Disabling.
	//
	// - Disabled: Disabled.
	//
	// - DisableFailed: Disabling failed.
	//
	// example:
	//
	// Enabled
	IdentifyStatus *string `json:"IdentifyStatus,omitempty" xml:"IdentifyStatus,omitempty"`
	// The database instance identity. You can obtain the value from the Items[].InstanceId parameter in the response of the ListDataMaskingInstances operation.
	//
	// example:
	//
	// rm-2ze1abcdefgh****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The language of the response message. Valid values:
	//
	// - zh: Chinese.
	//
	// - en: English.
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// A reserved pagination parameter. The ListDataSources operation uses CurrentPage and PageSize for pagination. Omit this parameter when you call the operation.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// A reserved pagination cursor parameter. The ListDataSources operation uses CurrentPage and PageSize for pagination. Omit this parameter when you call the operation.
	//
	// example:
	//
	// token-example
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The number of entries to return on each page. Keep other filter conditions consistent when you query the next page.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The port for the database connection.
	//
	// example:
	//
	// 3306
	Port *int32 `json:"Port,omitempty" xml:"Port,omitempty"`
	// The code of the cloud service. Column encryption supports ApsaraDB RDS, PolarDB, and PolarDB-X 2.0. If this parameter is specified together with ProductId, the values must match.
	//
	// example:
	//
	// RDS
	ProductCode *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	// The identifier of the cloud service type. 5: ApsaraDB RDS; 8: PolarDB; 27: PolarDB-X 2.0. If this parameter is specified together with ProductCode, the values must match.
	//
	// example:
	//
	// 5
	ProductId *int64 `json:"ProductId,omitempty" xml:"ProductId,omitempty"`
	// The source IP address of the call.
	//
	// example:
	//
	// 192.0.2.1
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
}

func (s ListDataSourcesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListDataSourcesRequest) GoString() string {
	return s.String()
}

func (s *ListDataSourcesRequest) GetConnectStatus() *string {
	return s.ConnectStatus
}

func (s *ListDataSourcesRequest) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *ListDataSourcesRequest) GetDataAssetId() *string {
	return s.DataAssetId
}

func (s *ListDataSourcesRequest) GetDataSourceId() *string {
	return s.DataSourceId
}

func (s *ListDataSourcesRequest) GetDbName() *string {
	return s.DbName
}

func (s *ListDataSourcesRequest) GetEngineType() *string {
	return s.EngineType
}

func (s *ListDataSourcesRequest) GetIdentifyStatus() *string {
	return s.IdentifyStatus
}

func (s *ListDataSourcesRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListDataSourcesRequest) GetLang() *string {
	return s.Lang
}

func (s *ListDataSourcesRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListDataSourcesRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListDataSourcesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListDataSourcesRequest) GetPort() *int32 {
	return s.Port
}

func (s *ListDataSourcesRequest) GetProductCode() *string {
	return s.ProductCode
}

func (s *ListDataSourcesRequest) GetProductId() *int64 {
	return s.ProductId
}

func (s *ListDataSourcesRequest) GetSourceIp() *string {
	return s.SourceIp
}

func (s *ListDataSourcesRequest) SetConnectStatus(v string) *ListDataSourcesRequest {
	s.ConnectStatus = &v
	return s
}

func (s *ListDataSourcesRequest) SetCurrentPage(v int32) *ListDataSourcesRequest {
	s.CurrentPage = &v
	return s
}

func (s *ListDataSourcesRequest) SetDataAssetId(v string) *ListDataSourcesRequest {
	s.DataAssetId = &v
	return s
}

func (s *ListDataSourcesRequest) SetDataSourceId(v string) *ListDataSourcesRequest {
	s.DataSourceId = &v
	return s
}

func (s *ListDataSourcesRequest) SetDbName(v string) *ListDataSourcesRequest {
	s.DbName = &v
	return s
}

func (s *ListDataSourcesRequest) SetEngineType(v string) *ListDataSourcesRequest {
	s.EngineType = &v
	return s
}

func (s *ListDataSourcesRequest) SetIdentifyStatus(v string) *ListDataSourcesRequest {
	s.IdentifyStatus = &v
	return s
}

func (s *ListDataSourcesRequest) SetInstanceId(v string) *ListDataSourcesRequest {
	s.InstanceId = &v
	return s
}

func (s *ListDataSourcesRequest) SetLang(v string) *ListDataSourcesRequest {
	s.Lang = &v
	return s
}

func (s *ListDataSourcesRequest) SetMaxResults(v int32) *ListDataSourcesRequest {
	s.MaxResults = &v
	return s
}

func (s *ListDataSourcesRequest) SetNextToken(v string) *ListDataSourcesRequest {
	s.NextToken = &v
	return s
}

func (s *ListDataSourcesRequest) SetPageSize(v int32) *ListDataSourcesRequest {
	s.PageSize = &v
	return s
}

func (s *ListDataSourcesRequest) SetPort(v int32) *ListDataSourcesRequest {
	s.Port = &v
	return s
}

func (s *ListDataSourcesRequest) SetProductCode(v string) *ListDataSourcesRequest {
	s.ProductCode = &v
	return s
}

func (s *ListDataSourcesRequest) SetProductId(v int64) *ListDataSourcesRequest {
	s.ProductId = &v
	return s
}

func (s *ListDataSourcesRequest) SetSourceIp(v string) *ListDataSourcesRequest {
	s.SourceIp = &v
	return s
}

func (s *ListDataSourcesRequest) Validate() error {
	return dara.Validate(s)
}
