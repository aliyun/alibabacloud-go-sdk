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
	// The page number of the returned data.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The list of records returned by the query.
	Items []*ListDataSourcesResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Repeated"`
	// The maximum number of records that can be returned for the current cursor-based pagination.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// A compatible pagination cursor field. The ListDataSources operation uses page number-based pagination and does not rely on this field to determine whether more pages exist.
	//
	// example:
	//
	// token-example
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The number of records returned on each page.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request identity, which is active when you invoke this API call.
	//
	// example:
	//
	// 7C6D8E9F-1234-5678-ABCD-0123456789AB
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries that meet the current statistics or query scope.
	//
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
	// The identifier of the data asset in Data Security Center.
	//
	// example:
	//
	// asset-example-001
	DataAssetId *string `json:"DataAssetId,omitempty" xml:"DataAssetId,omitempty"`
	// The identifier of the data source in Data Security Center.
	//
	// example:
	//
	// ds-example-001
	DataSourceId *string `json:"DataSourceId,omitempty" xml:"DataSourceId,omitempty"`
	// The statistical value of the data source capacity. This field is not used in the column encryption process for relational databases.
	//
	// example:
	//
	// 1024
	DataSourceSize *int64 `json:"DataSourceSize,omitempty" xml:"DataSourceSize,omitempty"`
	// The name of the database.
	//
	// example:
	//
	// business_db
	DbName *string `json:"DbName,omitempty" xml:"DbName,omitempty"`
	// The description of the data source.
	//
	// example:
	//
	// Business data source
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The type of the database engine.
	//
	// example:
	//
	// MySQL
	EngineType *string `json:"EngineType,omitempty" xml:"EngineType,omitempty"`
	// The business error code of the record or operation. This field can be empty if no error occurs. Evaluate this field together with ErrorMessage.
	//
	// example:
	//
	// data_masking_not_running
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The business error message of the record or operation.
	//
	// example:
	//
	// The instance is not in the running state
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The identifier of the record.
	//
	// example:
	//
	// 1001
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
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
	// The database instance identity.
	//
	// example:
	//
	// rm-2ze1abcdefgh****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The identity of the member accounts to which the asset belongs.
	//
	// example:
	//
	// 123456789012****
	MemberAccount *int64 `json:"MemberAccount,omitempty" xml:"MemberAccount,omitempty"`
	// The port for the database connection.
	//
	// example:
	//
	// 3306
	Port *int64 `json:"Port,omitempty" xml:"Port,omitempty"`
	// The code of the cloud service.
	//
	// example:
	//
	// RDS
	ProductCode *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	// The identifier of the cloud service type. 5: ApsaraDB RDS; 8: PolarDB; 27: PolarDB-X 2.0.
	//
	// example:
	//
	// 5
	ProductId *int64 `json:"ProductId,omitempty" xml:"ProductId,omitempty"`
	// The ID of the region where the asset resides.
	//
	// example:
	//
	// cn-zhangjiakou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The name of the region where the asset resides.
	//
	// example:
	//
	// China (Zhangjiakou)
	RegionName *string `json:"RegionName,omitempty" xml:"RegionName,omitempty"`
	// The ID of the resource group to which the asset belongs.
	//
	// example:
	//
	// rg-acfmexample****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The identifier of the database tenant.
	//
	// example:
	//
	// tenant-example
	TenantId *string `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
	// The name of the database account used to connect to the data source.
	//
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
