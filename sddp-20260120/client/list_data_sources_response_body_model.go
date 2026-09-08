// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDataSourcesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCurrentPage(v int32) *ListDataSourcesResponseBody
	GetCurrentPage() *int32
	SetItems(v []*ListDataSourcesResponseBodyItems) *ListDataSourcesResponseBody
	GetItems() []*ListDataSourcesResponseBodyItems
	SetMaxResults(v int32) *ListDataSourcesResponseBody
	GetMaxResults() *int32
	SetNextToken(v string) *ListDataSourcesResponseBody
	GetNextToken() *string
	SetPageSize(v int32) *ListDataSourcesResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *ListDataSourcesResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *ListDataSourcesResponseBody
	GetTotalCount() *int32
}

type ListDataSourcesResponseBody struct {
	// example:
	//
	// 1
	CurrentPage *int32                              `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	Items       []*ListDataSourcesResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Repeated"`
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
	// 7C6D8E9F-1234-5678-ABCD-0123456789AB
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 100
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListDataSourcesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListDataSourcesResponseBody) GoString() string {
	return s.String()
}

func (s *ListDataSourcesResponseBody) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *ListDataSourcesResponseBody) GetItems() []*ListDataSourcesResponseBodyItems {
	return s.Items
}

func (s *ListDataSourcesResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListDataSourcesResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListDataSourcesResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListDataSourcesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListDataSourcesResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListDataSourcesResponseBody) SetCurrentPage(v int32) *ListDataSourcesResponseBody {
	s.CurrentPage = &v
	return s
}

func (s *ListDataSourcesResponseBody) SetItems(v []*ListDataSourcesResponseBodyItems) *ListDataSourcesResponseBody {
	s.Items = v
	return s
}

func (s *ListDataSourcesResponseBody) SetMaxResults(v int32) *ListDataSourcesResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListDataSourcesResponseBody) SetNextToken(v string) *ListDataSourcesResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListDataSourcesResponseBody) SetPageSize(v int32) *ListDataSourcesResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListDataSourcesResponseBody) SetRequestId(v string) *ListDataSourcesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListDataSourcesResponseBody) SetTotalCount(v int32) *ListDataSourcesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListDataSourcesResponseBody) Validate() error {
	if s.Items != nil {
		for _, item := range s.Items {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListDataSourcesResponseBodyItems struct {
	// example:
	//
	// Success
	ConnectStatus *string `json:"ConnectStatus,omitempty" xml:"ConnectStatus,omitempty"`
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
	// 1024
	DataSourceSize *int64 `json:"DataSourceSize,omitempty" xml:"DataSourceSize,omitempty"`
	// example:
	//
	// business_db
	DbName *string `json:"DbName,omitempty" xml:"DbName,omitempty"`
	// example:
	//
	// 业务数据源
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// MySQL
	EngineType *string `json:"EngineType,omitempty" xml:"EngineType,omitempty"`
	// example:
	//
	// data_masking_not_running
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// example:
	//
	// 实例未处于运行状态
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// example:
	//
	// 1001
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
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
	// 123456789012****
	MemberAccount *int64 `json:"MemberAccount,omitempty" xml:"MemberAccount,omitempty"`
	// example:
	//
	// 3306
	Port *int64 `json:"Port,omitempty" xml:"Port,omitempty"`
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
	// cn-zhangjiakou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// 华北 3（张家口）
	RegionName *string `json:"RegionName,omitempty" xml:"RegionName,omitempty"`
	// example:
	//
	// rg-acfmexample****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// example:
	//
	// tenant-example
	TenantId *string `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
	// example:
	//
	// dsc_reader
	UserName *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
}

func (s ListDataSourcesResponseBodyItems) String() string {
	return dara.Prettify(s)
}

func (s ListDataSourcesResponseBodyItems) GoString() string {
	return s.String()
}

func (s *ListDataSourcesResponseBodyItems) GetConnectStatus() *string {
	return s.ConnectStatus
}

func (s *ListDataSourcesResponseBodyItems) GetDataAssetId() *string {
	return s.DataAssetId
}

func (s *ListDataSourcesResponseBodyItems) GetDataSourceId() *string {
	return s.DataSourceId
}

func (s *ListDataSourcesResponseBodyItems) GetDataSourceSize() *int64 {
	return s.DataSourceSize
}

func (s *ListDataSourcesResponseBodyItems) GetDbName() *string {
	return s.DbName
}

func (s *ListDataSourcesResponseBodyItems) GetDescription() *string {
	return s.Description
}

func (s *ListDataSourcesResponseBodyItems) GetEngineType() *string {
	return s.EngineType
}

func (s *ListDataSourcesResponseBodyItems) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *ListDataSourcesResponseBodyItems) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *ListDataSourcesResponseBodyItems) GetId() *int64 {
	return s.Id
}

func (s *ListDataSourcesResponseBodyItems) GetIdentifyStatus() *string {
	return s.IdentifyStatus
}

func (s *ListDataSourcesResponseBodyItems) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListDataSourcesResponseBodyItems) GetMemberAccount() *int64 {
	return s.MemberAccount
}

func (s *ListDataSourcesResponseBodyItems) GetPort() *int64 {
	return s.Port
}

func (s *ListDataSourcesResponseBodyItems) GetProductCode() *string {
	return s.ProductCode
}

func (s *ListDataSourcesResponseBodyItems) GetProductId() *int64 {
	return s.ProductId
}

func (s *ListDataSourcesResponseBodyItems) GetRegionId() *string {
	return s.RegionId
}

func (s *ListDataSourcesResponseBodyItems) GetRegionName() *string {
	return s.RegionName
}

func (s *ListDataSourcesResponseBodyItems) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *ListDataSourcesResponseBodyItems) GetTenantId() *string {
	return s.TenantId
}

func (s *ListDataSourcesResponseBodyItems) GetUserName() *string {
	return s.UserName
}

func (s *ListDataSourcesResponseBodyItems) SetConnectStatus(v string) *ListDataSourcesResponseBodyItems {
	s.ConnectStatus = &v
	return s
}

func (s *ListDataSourcesResponseBodyItems) SetDataAssetId(v string) *ListDataSourcesResponseBodyItems {
	s.DataAssetId = &v
	return s
}

func (s *ListDataSourcesResponseBodyItems) SetDataSourceId(v string) *ListDataSourcesResponseBodyItems {
	s.DataSourceId = &v
	return s
}

func (s *ListDataSourcesResponseBodyItems) SetDataSourceSize(v int64) *ListDataSourcesResponseBodyItems {
	s.DataSourceSize = &v
	return s
}

func (s *ListDataSourcesResponseBodyItems) SetDbName(v string) *ListDataSourcesResponseBodyItems {
	s.DbName = &v
	return s
}

func (s *ListDataSourcesResponseBodyItems) SetDescription(v string) *ListDataSourcesResponseBodyItems {
	s.Description = &v
	return s
}

func (s *ListDataSourcesResponseBodyItems) SetEngineType(v string) *ListDataSourcesResponseBodyItems {
	s.EngineType = &v
	return s
}

func (s *ListDataSourcesResponseBodyItems) SetErrorCode(v string) *ListDataSourcesResponseBodyItems {
	s.ErrorCode = &v
	return s
}

func (s *ListDataSourcesResponseBodyItems) SetErrorMessage(v string) *ListDataSourcesResponseBodyItems {
	s.ErrorMessage = &v
	return s
}

func (s *ListDataSourcesResponseBodyItems) SetId(v int64) *ListDataSourcesResponseBodyItems {
	s.Id = &v
	return s
}

func (s *ListDataSourcesResponseBodyItems) SetIdentifyStatus(v string) *ListDataSourcesResponseBodyItems {
	s.IdentifyStatus = &v
	return s
}

func (s *ListDataSourcesResponseBodyItems) SetInstanceId(v string) *ListDataSourcesResponseBodyItems {
	s.InstanceId = &v
	return s
}

func (s *ListDataSourcesResponseBodyItems) SetMemberAccount(v int64) *ListDataSourcesResponseBodyItems {
	s.MemberAccount = &v
	return s
}

func (s *ListDataSourcesResponseBodyItems) SetPort(v int64) *ListDataSourcesResponseBodyItems {
	s.Port = &v
	return s
}

func (s *ListDataSourcesResponseBodyItems) SetProductCode(v string) *ListDataSourcesResponseBodyItems {
	s.ProductCode = &v
	return s
}

func (s *ListDataSourcesResponseBodyItems) SetProductId(v int64) *ListDataSourcesResponseBodyItems {
	s.ProductId = &v
	return s
}

func (s *ListDataSourcesResponseBodyItems) SetRegionId(v string) *ListDataSourcesResponseBodyItems {
	s.RegionId = &v
	return s
}

func (s *ListDataSourcesResponseBodyItems) SetRegionName(v string) *ListDataSourcesResponseBodyItems {
	s.RegionName = &v
	return s
}

func (s *ListDataSourcesResponseBodyItems) SetResourceGroupId(v string) *ListDataSourcesResponseBodyItems {
	s.ResourceGroupId = &v
	return s
}

func (s *ListDataSourcesResponseBodyItems) SetTenantId(v string) *ListDataSourcesResponseBodyItems {
	s.TenantId = &v
	return s
}

func (s *ListDataSourcesResponseBodyItems) SetUserName(v string) *ListDataSourcesResponseBodyItems {
	s.UserName = &v
	return s
}

func (s *ListDataSourcesResponseBodyItems) Validate() error {
	return dara.Validate(s)
}
