// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeInstanceSourcesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCurrentPage(v int32) *DescribeInstanceSourcesResponseBody
	GetCurrentPage() *int32
	SetItems(v []*DescribeInstanceSourcesResponseBodyItems) *DescribeInstanceSourcesResponseBody
	GetItems() []*DescribeInstanceSourcesResponseBodyItems
	SetPageSize(v int32) *DescribeInstanceSourcesResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *DescribeInstanceSourcesResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *DescribeInstanceSourcesResponseBody
	GetTotalCount() *int32
}

type DescribeInstanceSourcesResponseBody struct {
	// The page number of the returned page.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The assets.
	Items []*DescribeInstanceSourcesResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Repeated"`
	// The number of entries returned per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// 5A7E8FB9-5011-5A90-97D9-AE587047****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 2
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeInstanceSourcesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstanceSourcesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeInstanceSourcesResponseBody) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *DescribeInstanceSourcesResponseBody) GetItems() []*DescribeInstanceSourcesResponseBodyItems {
	return s.Items
}

func (s *DescribeInstanceSourcesResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeInstanceSourcesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeInstanceSourcesResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeInstanceSourcesResponseBody) SetCurrentPage(v int32) *DescribeInstanceSourcesResponseBody {
	s.CurrentPage = &v
	return s
}

func (s *DescribeInstanceSourcesResponseBody) SetItems(v []*DescribeInstanceSourcesResponseBodyItems) *DescribeInstanceSourcesResponseBody {
	s.Items = v
	return s
}

func (s *DescribeInstanceSourcesResponseBody) SetPageSize(v int32) *DescribeInstanceSourcesResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeInstanceSourcesResponseBody) SetRequestId(v string) *DescribeInstanceSourcesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeInstanceSourcesResponseBody) SetTotalCount(v int32) *DescribeInstanceSourcesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeInstanceSourcesResponseBody) Validate() error {
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

type DescribeInstanceSourcesResponseBodyItems struct {
	// Indicates whether the security audit feature is enabled. Valid values:
	//
	// 	- **1**: yes
	//
	// 	- **0**: no
	//
	// example:
	//
	// 1
	AuditStatus *int32 `json:"AuditStatus,omitempty" xml:"AuditStatus,omitempty"`
	// Indicates whether the automatic scan feature is enabled to detect sensitive data. Valid values:
	//
	// 	- **0**: no
	//
	// 	- **1**: yes
	//
	// example:
	//
	// 0
	AutoScan *int32 `json:"AutoScan,omitempty" xml:"AutoScan,omitempty"`
	// Indicates whether the username and password can be changed. Valid values:
	//
	// 	- **true**: yes
	//
	// 	- **false**: no
	//
	// example:
	//
	// true
	CanModifyUserName *bool `json:"CanModifyUserName,omitempty" xml:"CanModifyUserName,omitempty"`
	// The data detection status. Valid values:
	//
	// 	- **0**: The data detection is ready.
	//
	// 	- **1**: The data detection is running.
	//
	// 	- **2**: The connectivity test is in progress.
	//
	// 	- **3**: The connectivity test passed.
	//
	// 	- **4**: The connectivity test failed.
	//
	// example:
	//
	// 3
	CheckStatus *int32 `json:"CheckStatus,omitempty" xml:"CheckStatus,omitempty"`
	// Indicates whether DSC has the data de-identification permissions on the data asset. Valid values:
	//
	// 	- **1**: yes
	//
	// 	- **0**: no
	//
	// example:
	//
	// 1
	DatamaskStatus *int32 `json:"DatamaskStatus,omitempty" xml:"DatamaskStatus,omitempty"`
	// The name of the database to which the data asset belongs.
	//
	// example:
	//
	// demo
	DbName *string `json:"DbName,omitempty" xml:"DbName,omitempty"`
	// Indicates whether sensitive data detection is enabled for the data asset. Valid values:
	//
	// 	- **1**: yes
	//
	// 	- **0**: no
	//
	// example:
	//
	// 1
	Enable *int32 `json:"Enable,omitempty" xml:"Enable,omitempty"`
	// The engine type. Valid values:
	//
	// 	- **MySQL**
	//
	// 	- **MariaDB**
	//
	// 	- **Oracle**
	//
	// 	- **PostgreSQL**
	//
	// 	- **SQLServer**
	//
	// example:
	//
	// MySQL
	EngineType *string `json:"EngineType,omitempty" xml:"EngineType,omitempty"`
	// The reason for the failure.
	//
	// example:
	//
	// The password is invalid.
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The time when the data asset was created. The value is a UNIX timestamp. Unit: milliseconds.
	//
	// example:
	//
	// 1625587423000
	GmtCreate *int64 `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// The unique ID of the data asset.
	//
	// example:
	//
	// 1
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The description of the instance.
	//
	// example:
	//
	// Test
	InstanceDescription *string `json:"InstanceDescription,omitempty" xml:"InstanceDescription,omitempty"`
	// The ID of the instance
	//
	// example:
	//
	// rm-****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The storage space size of the instance. This parameter is valid only if the value of the ProductId parameter is 2. Unit: bytes.
	//
	// example:
	//
	// 409600
	InstanceSize *int64 `json:"InstanceSize,omitempty" xml:"InstanceSize,omitempty"`
	// The time when the data asset was last modified. Unit: milliseconds.
	//
	// example:
	//
	// 1625587423000
	LastModifyTime *int64 `json:"LastModifyTime,omitempty" xml:"LastModifyTime,omitempty"`
	// The ID of the account that is last used to modify the data asset.
	//
	// example:
	//
	// demo
	LastModifyUserId *string `json:"LastModifyUserId,omitempty" xml:"LastModifyUserId,omitempty"`
	// The retention period of raw logs. Unit: days.
	//
	// example:
	//
	// 30
	LogStoreDay *int32 `json:"LogStoreDay,omitempty" xml:"LogStoreDay,omitempty"`
	// Indicates whether the password is used. Valid values:
	//
	// 	- **1**: yes
	//
	// 	- **0**: no
	//
	// example:
	//
	// 1
	PasswordStatus *int32 `json:"PasswordStatus,omitempty" xml:"PasswordStatus,omitempty"`
	// The ID of the service to which the asset belongs. Valid values:
	//
	// 	- **1**: MaxCompute
	//
	// 	- **2**: OSS
	//
	// 	- **3**: AnalyticDB for MySQL
	//
	// 	- **4**: OTS
	//
	// 	- **5**: ApsaraDB RDS
	//
	// 	- **6**: self-managed databases
	//
	// example:
	//
	// 2
	ProductId *int64 `json:"ProductId,omitempty" xml:"ProductId,omitempty"`
	// The ID of the region where the instance resides.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The name of the region.
	//
	// example:
	//
	// China (Hangzhou)
	RegionName *string `json:"RegionName,omitempty" xml:"RegionName,omitempty"`
	// The number of sensitive data samples. Valid values: **0**, **5**, and **10**. Unit: data entries.
	//
	// example:
	//
	// 10
	SamplingSize *int32 `json:"SamplingSize,omitempty" xml:"SamplingSize,omitempty"`
	// The ID of the tenant.
	//
	// example:
	//
	// 11
	TenantId *string `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
	// The name of the tenant.
	//
	// example:
	//
	// user1
	TenantName *string `json:"TenantName,omitempty" xml:"TenantName,omitempty"`
	// The username of the account.
	//
	// example:
	//
	// User01
	UserName *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
}

func (s DescribeInstanceSourcesResponseBodyItems) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstanceSourcesResponseBodyItems) GoString() string {
	return s.String()
}

func (s *DescribeInstanceSourcesResponseBodyItems) GetAuditStatus() *int32 {
	return s.AuditStatus
}

func (s *DescribeInstanceSourcesResponseBodyItems) GetAutoScan() *int32 {
	return s.AutoScan
}

func (s *DescribeInstanceSourcesResponseBodyItems) GetCanModifyUserName() *bool {
	return s.CanModifyUserName
}

func (s *DescribeInstanceSourcesResponseBodyItems) GetCheckStatus() *int32 {
	return s.CheckStatus
}

func (s *DescribeInstanceSourcesResponseBodyItems) GetDatamaskStatus() *int32 {
	return s.DatamaskStatus
}

func (s *DescribeInstanceSourcesResponseBodyItems) GetDbName() *string {
	return s.DbName
}

func (s *DescribeInstanceSourcesResponseBodyItems) GetEnable() *int32 {
	return s.Enable
}

func (s *DescribeInstanceSourcesResponseBodyItems) GetEngineType() *string {
	return s.EngineType
}

func (s *DescribeInstanceSourcesResponseBodyItems) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *DescribeInstanceSourcesResponseBodyItems) GetGmtCreate() *int64 {
	return s.GmtCreate
}

func (s *DescribeInstanceSourcesResponseBodyItems) GetId() *int64 {
	return s.Id
}

func (s *DescribeInstanceSourcesResponseBodyItems) GetInstanceDescription() *string {
	return s.InstanceDescription
}

func (s *DescribeInstanceSourcesResponseBodyItems) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeInstanceSourcesResponseBodyItems) GetInstanceSize() *int64 {
	return s.InstanceSize
}

func (s *DescribeInstanceSourcesResponseBodyItems) GetLastModifyTime() *int64 {
	return s.LastModifyTime
}

func (s *DescribeInstanceSourcesResponseBodyItems) GetLastModifyUserId() *string {
	return s.LastModifyUserId
}

func (s *DescribeInstanceSourcesResponseBodyItems) GetLogStoreDay() *int32 {
	return s.LogStoreDay
}

func (s *DescribeInstanceSourcesResponseBodyItems) GetPasswordStatus() *int32 {
	return s.PasswordStatus
}

func (s *DescribeInstanceSourcesResponseBodyItems) GetProductId() *int64 {
	return s.ProductId
}

func (s *DescribeInstanceSourcesResponseBodyItems) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeInstanceSourcesResponseBodyItems) GetRegionName() *string {
	return s.RegionName
}

func (s *DescribeInstanceSourcesResponseBodyItems) GetSamplingSize() *int32 {
	return s.SamplingSize
}

func (s *DescribeInstanceSourcesResponseBodyItems) GetTenantId() *string {
	return s.TenantId
}

func (s *DescribeInstanceSourcesResponseBodyItems) GetTenantName() *string {
	return s.TenantName
}

func (s *DescribeInstanceSourcesResponseBodyItems) GetUserName() *string {
	return s.UserName
}

func (s *DescribeInstanceSourcesResponseBodyItems) SetAuditStatus(v int32) *DescribeInstanceSourcesResponseBodyItems {
	s.AuditStatus = &v
	return s
}

func (s *DescribeInstanceSourcesResponseBodyItems) SetAutoScan(v int32) *DescribeInstanceSourcesResponseBodyItems {
	s.AutoScan = &v
	return s
}

func (s *DescribeInstanceSourcesResponseBodyItems) SetCanModifyUserName(v bool) *DescribeInstanceSourcesResponseBodyItems {
	s.CanModifyUserName = &v
	return s
}

func (s *DescribeInstanceSourcesResponseBodyItems) SetCheckStatus(v int32) *DescribeInstanceSourcesResponseBodyItems {
	s.CheckStatus = &v
	return s
}

func (s *DescribeInstanceSourcesResponseBodyItems) SetDatamaskStatus(v int32) *DescribeInstanceSourcesResponseBodyItems {
	s.DatamaskStatus = &v
	return s
}

func (s *DescribeInstanceSourcesResponseBodyItems) SetDbName(v string) *DescribeInstanceSourcesResponseBodyItems {
	s.DbName = &v
	return s
}

func (s *DescribeInstanceSourcesResponseBodyItems) SetEnable(v int32) *DescribeInstanceSourcesResponseBodyItems {
	s.Enable = &v
	return s
}

func (s *DescribeInstanceSourcesResponseBodyItems) SetEngineType(v string) *DescribeInstanceSourcesResponseBodyItems {
	s.EngineType = &v
	return s
}

func (s *DescribeInstanceSourcesResponseBodyItems) SetErrorMessage(v string) *DescribeInstanceSourcesResponseBodyItems {
	s.ErrorMessage = &v
	return s
}

func (s *DescribeInstanceSourcesResponseBodyItems) SetGmtCreate(v int64) *DescribeInstanceSourcesResponseBodyItems {
	s.GmtCreate = &v
	return s
}

func (s *DescribeInstanceSourcesResponseBodyItems) SetId(v int64) *DescribeInstanceSourcesResponseBodyItems {
	s.Id = &v
	return s
}

func (s *DescribeInstanceSourcesResponseBodyItems) SetInstanceDescription(v string) *DescribeInstanceSourcesResponseBodyItems {
	s.InstanceDescription = &v
	return s
}

func (s *DescribeInstanceSourcesResponseBodyItems) SetInstanceId(v string) *DescribeInstanceSourcesResponseBodyItems {
	s.InstanceId = &v
	return s
}

func (s *DescribeInstanceSourcesResponseBodyItems) SetInstanceSize(v int64) *DescribeInstanceSourcesResponseBodyItems {
	s.InstanceSize = &v
	return s
}

func (s *DescribeInstanceSourcesResponseBodyItems) SetLastModifyTime(v int64) *DescribeInstanceSourcesResponseBodyItems {
	s.LastModifyTime = &v
	return s
}

func (s *DescribeInstanceSourcesResponseBodyItems) SetLastModifyUserId(v string) *DescribeInstanceSourcesResponseBodyItems {
	s.LastModifyUserId = &v
	return s
}

func (s *DescribeInstanceSourcesResponseBodyItems) SetLogStoreDay(v int32) *DescribeInstanceSourcesResponseBodyItems {
	s.LogStoreDay = &v
	return s
}

func (s *DescribeInstanceSourcesResponseBodyItems) SetPasswordStatus(v int32) *DescribeInstanceSourcesResponseBodyItems {
	s.PasswordStatus = &v
	return s
}

func (s *DescribeInstanceSourcesResponseBodyItems) SetProductId(v int64) *DescribeInstanceSourcesResponseBodyItems {
	s.ProductId = &v
	return s
}

func (s *DescribeInstanceSourcesResponseBodyItems) SetRegionId(v string) *DescribeInstanceSourcesResponseBodyItems {
	s.RegionId = &v
	return s
}

func (s *DescribeInstanceSourcesResponseBodyItems) SetRegionName(v string) *DescribeInstanceSourcesResponseBodyItems {
	s.RegionName = &v
	return s
}

func (s *DescribeInstanceSourcesResponseBodyItems) SetSamplingSize(v int32) *DescribeInstanceSourcesResponseBodyItems {
	s.SamplingSize = &v
	return s
}

func (s *DescribeInstanceSourcesResponseBodyItems) SetTenantId(v string) *DescribeInstanceSourcesResponseBodyItems {
	s.TenantId = &v
	return s
}

func (s *DescribeInstanceSourcesResponseBodyItems) SetTenantName(v string) *DescribeInstanceSourcesResponseBodyItems {
	s.TenantName = &v
	return s
}

func (s *DescribeInstanceSourcesResponseBodyItems) SetUserName(v string) *DescribeInstanceSourcesResponseBodyItems {
	s.UserName = &v
	return s
}

func (s *DescribeInstanceSourcesResponseBodyItems) Validate() error {
	return dara.Validate(s)
}
