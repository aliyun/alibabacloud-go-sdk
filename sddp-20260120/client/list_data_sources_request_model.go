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
	SetProductCode(v string) *ListDataSourcesRequest
	GetProductCode() *string
	SetProductId(v int64) *ListDataSourcesRequest
	GetProductId() *int64
	SetSourceIp(v string) *ListDataSourcesRequest
	GetSourceIp() *string
}

type ListDataSourcesRequest struct {
	// example:
	//
	// Success
	ConnectStatus *string `json:"ConnectStatus,omitempty" xml:"ConnectStatus,omitempty"`
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// example:
	//
	// asset-example-001
	DataAssetId *string `json:"DataAssetId,omitempty" xml:"DataAssetId,omitempty"`
	// example:
	//
	// ds-example-001
	DataSourceId *string `json:"DataSourceId,omitempty" xml:"DataSourceId,omitempty"`
	// example:
	//
	// business_db
	DbName *string `json:"DbName,omitempty" xml:"DbName,omitempty"`
	// example:
	//
	// Enabled
	IdentifyStatus *string `json:"IdentifyStatus,omitempty" xml:"IdentifyStatus,omitempty"`
	// example:
	//
	// rm-2ze1abcdefgh****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// example:
	//
	// token-example
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// RDS
	ProductCode *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	// example:
	//
	// 5
	ProductId *int64 `json:"ProductId,omitempty" xml:"ProductId,omitempty"`
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
