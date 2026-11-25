// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetLindormInstanceListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceList(v []*GetLindormInstanceListResponseBodyInstanceList) *GetLindormInstanceListResponseBody
	GetInstanceList() []*GetLindormInstanceListResponseBodyInstanceList
	SetPageNumber(v int32) *GetLindormInstanceListResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *GetLindormInstanceListResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *GetLindormInstanceListResponseBody
	GetRequestId() *string
	SetTotal(v int32) *GetLindormInstanceListResponseBody
	GetTotal() *int32
}

type GetLindormInstanceListResponseBody struct {
	// The list of instances.
	InstanceList []*GetLindormInstanceListResponseBodyInstanceList `json:"InstanceList,omitempty" xml:"InstanceList,omitempty" type:"Repeated"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries returned per page.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 1CA1FAFD-E8DC-51C2-AA7E-CA6E2D049BA0
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of instances returned.
	//
	// example:
	//
	// 1
	Total *int32 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s GetLindormInstanceListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetLindormInstanceListResponseBody) GoString() string {
	return s.String()
}

func (s *GetLindormInstanceListResponseBody) GetInstanceList() []*GetLindormInstanceListResponseBodyInstanceList {
	return s.InstanceList
}

func (s *GetLindormInstanceListResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *GetLindormInstanceListResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *GetLindormInstanceListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetLindormInstanceListResponseBody) GetTotal() *int32 {
	return s.Total
}

func (s *GetLindormInstanceListResponseBody) SetInstanceList(v []*GetLindormInstanceListResponseBodyInstanceList) *GetLindormInstanceListResponseBody {
	s.InstanceList = v
	return s
}

func (s *GetLindormInstanceListResponseBody) SetPageNumber(v int32) *GetLindormInstanceListResponseBody {
	s.PageNumber = &v
	return s
}

func (s *GetLindormInstanceListResponseBody) SetPageSize(v int32) *GetLindormInstanceListResponseBody {
	s.PageSize = &v
	return s
}

func (s *GetLindormInstanceListResponseBody) SetRequestId(v string) *GetLindormInstanceListResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetLindormInstanceListResponseBody) SetTotal(v int32) *GetLindormInstanceListResponseBody {
	s.Total = &v
	return s
}

func (s *GetLindormInstanceListResponseBody) Validate() error {
	if s.InstanceList != nil {
		for _, item := range s.InstanceList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetLindormInstanceListResponseBodyInstanceList struct {
	// The 16-character AliUid of the Alibaba Cloud account.
	//
	// example:
	//
	// 164901546557****
	AliUid *int64 `json:"AliUid,omitempty" xml:"AliUid,omitempty"`
	// The cause of the creation failure.
	//
	// example:
	//
	// Resource is not enough
	CreateErrorCode *string `json:"CreateErrorCode,omitempty" xml:"CreateErrorCode,omitempty"`
	// The number of milliseconds between the instance creation time and 00:00:00 on January 1, 1970.
	//
	// example:
	//
	// 1631772842000
	CreateMilliseconds *int64 `json:"CreateMilliseconds,omitempty" xml:"CreateMilliseconds,omitempty"`
	// The creation time of the instance. The time is in the yyyy-MM-dd HH:mm:ss format.
	//
	// example:
	//
	// 2021-09-16 14:13:13
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// Indicates whether the column store engine is enabled. Valid values:
	//
	// - **true**: Enabled.
	//
	// - **false**: Not enabled.
	//
	// example:
	//
	// true
	EnableColumn *bool `json:"EnableColumn,omitempty" xml:"EnableColumn,omitempty"`
	// Indicates whether the compute engine is enabled for the instance. Valid values:
	//
	// - **true**: Enabled.
	//
	// - **false**: Not enabled.
	//
	// example:
	//
	// true
	EnableCompute *bool `json:"EnableCompute,omitempty" xml:"EnableCompute,omitempty"`
	// Indicates whether Lindorm Tunnel Service (LTS) is enabled. Valid values:
	//
	// - **true**: Enabled.
	//
	// - **false**: Not enabled.
	//
	// example:
	//
	// true
	EnableLts *bool `json:"EnableLts,omitempty" xml:"EnableLts,omitempty"`
	// Indicates whether the message engine is enabled. Valid values:
	//
	// - **true**: Enabled.
	//
	// - **false**: Not enabled.
	//
	// example:
	//
	// true
	EnableMessage *bool `json:"EnableMessage,omitempty" xml:"EnableMessage,omitempty"`
	// Indicates whether LindormTable 3.0 is enabled. Valid values:
	//
	// - **true**: Enabled.
	//
	// - **false**: Not enabled.
	//
	// example:
	//
	// true
	EnableRow *bool `json:"EnableRow,omitempty" xml:"EnableRow,omitempty"`
	// Indicates whether the streaming engine is enabled for the instance. Valid values:
	//
	// - **true**: The streaming engine is enabled.
	//
	// - **false**: The streaming engine is not enabled.
	//
	// example:
	//
	// true
	EnableStream *bool `json:"EnableStream,omitempty" xml:"EnableStream,omitempty"`
	// Indicates whether the vector engine is enabled. Valid values:
	//
	// - **true**: Enabled.
	//
	// - **false**: Not enabled.
	//
	// example:
	//
	// true
	EnableVector *bool `json:"EnableVector,omitempty" xml:"EnableVector,omitempty"`
	// The engine types supported by the instance. The return value is the sum of the following engine type values:
	//
	// - **1**: Search engine
	//
	// - **2**: LindormTSDB
	//
	// - **4**: LindormTable
	//
	// - **8**: File engine
	//
	// > For example, if the value is 15 (15 = 8 + 4 + 2 + 1), the instance supports the search engine, LindormTSDB, LindormTable, and file engine. If the value is 6 (6 = 4 + 2), the instance supports LindormTSDB and LindormTable.
	//
	// example:
	//
	// 15
	EngineType *string `json:"EngineType,omitempty" xml:"EngineType,omitempty"`
	// The expiration time of the instance. The time is in the yyyy-MM-dd HH:mm:ss format.
	//
	// > This parameter is returned only when the billing method of the instance is subscription.
	//
	// example:
	//
	// 2022-04-26 00:00:00
	ExpireTime *string `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	// The number of milliseconds between the time when the instance expires and 00:00:00 on January 1, 1970.
	//
	// example:
	//
	// 1650902400000
	ExpiredMilliseconds *int64 `json:"ExpiredMilliseconds,omitempty" xml:"ExpiredMilliseconds,omitempty"`
	// The instance name.
	//
	// example:
	//
	// test
	InstanceAlias *string `json:"InstanceAlias,omitempty" xml:"InstanceAlias,omitempty"`
	// The instance ID.
	//
	// example:
	//
	// ld-bp17pwu1541ia****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The instance status. Valid values:
	//
	// - **CREATING**: Creating.
	//
	// - **ACTIVATION**: Running.
	//
	// - **COLD_EXPANDING**: Performing storage expansion for storage-optimized cloud storage.
	//
	// - **MINOR_VERSION_TRANSING**: Upgrading the minor version.
	//
	// - **RESIZING**: Scaling out nodes.
	//
	// - **SHRINKING**: Scaling in nodes.
	//
	// - **CLASS_CHANGING**: Upgrading or downgrading specifications.
	//
	// - **SSL_SWITCHING**: Modifying SSL configurations.
	//
	// - **CDC_OPENING**: Enabling the data subscription feature.
	//
	// - **TRANSFER**: Migrating data.
	//
	// - **DATABASE_TRANSFER**: Migrating data to the database.
	//
	// - **GUARD_CREATING**: Creating a disaster recovery instance.
	//
	// - **BACKUP_RECOVERING**: Restoring from a backup.
	//
	// - **DATABASE_IMPORTING**: Importing data.
	//
	// - **NET_MODIFYING**: Modifying network configurations.
	//
	// - **NET_SWITCHING**: Switching between the internal network and the Internet.
	//
	// - **NET_CREATING**: Creating a network connection.
	//
	// - **NET_DELETING**: Performing a delete operation on the network connection.
	//
	// - **DELETING**: Performing a delete operation.
	//
	// - **RESTARTING**: Restarting.
	//
	// - **LOCKED**: Locked because the instance has expired.
	//
	// example:
	//
	// ACTIVATION
	InstanceStatus *string `json:"InstanceStatus,omitempty" xml:"InstanceStatus,omitempty"`
	// The storage capacity of the instance.
	//
	// example:
	//
	// 960
	InstanceStorage *string `json:"InstanceStorage,omitempty" xml:"InstanceStorage,omitempty"`
	// The network type of the instance.
	//
	// example:
	//
	// vpc
	NetworkType *string `json:"NetworkType,omitempty" xml:"NetworkType,omitempty"`
	// The billing method of the instance. Valid values:
	//
	// - **PREPAY**: Subscription
	//
	// - **POSTPAY**: Pay-as-you-go
	//
	// example:
	//
	// PREPAY
	PayType *string `json:"PayType,omitempty" xml:"PayType,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The resource group ID.
	//
	// example:
	//
	// rg-aekzledqeat****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The instance type. Valid values:
	//
	// - **lindorm**: Lindorm instance
	//
	// - **serverless_lindorm**: Lindorm Serverless instance
	//
	// - **lindorm_standalone**: Lindorm single-node instance
	//
	// - **lts**: Lindorm Tunnel Service
	//
	// example:
	//
	// lindorm
	ServiceType *string `json:"ServiceType,omitempty" xml:"ServiceType,omitempty"`
	// The list of tags.
	Tags []*GetLindormInstanceListResponseBodyInstanceListTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	// The VPC ID of the instance.
	//
	// example:
	//
	// vpc-bp1n3i15v90el48nx****
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	// The zone ID.
	//
	// example:
	//
	// cn-hangzhou-h
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s GetLindormInstanceListResponseBodyInstanceList) String() string {
	return dara.Prettify(s)
}

func (s GetLindormInstanceListResponseBodyInstanceList) GoString() string {
	return s.String()
}

func (s *GetLindormInstanceListResponseBodyInstanceList) GetAliUid() *int64 {
	return s.AliUid
}

func (s *GetLindormInstanceListResponseBodyInstanceList) GetCreateErrorCode() *string {
	return s.CreateErrorCode
}

func (s *GetLindormInstanceListResponseBodyInstanceList) GetCreateMilliseconds() *int64 {
	return s.CreateMilliseconds
}

func (s *GetLindormInstanceListResponseBodyInstanceList) GetCreateTime() *string {
	return s.CreateTime
}

func (s *GetLindormInstanceListResponseBodyInstanceList) GetEnableColumn() *bool {
	return s.EnableColumn
}

func (s *GetLindormInstanceListResponseBodyInstanceList) GetEnableCompute() *bool {
	return s.EnableCompute
}

func (s *GetLindormInstanceListResponseBodyInstanceList) GetEnableLts() *bool {
	return s.EnableLts
}

func (s *GetLindormInstanceListResponseBodyInstanceList) GetEnableMessage() *bool {
	return s.EnableMessage
}

func (s *GetLindormInstanceListResponseBodyInstanceList) GetEnableRow() *bool {
	return s.EnableRow
}

func (s *GetLindormInstanceListResponseBodyInstanceList) GetEnableStream() *bool {
	return s.EnableStream
}

func (s *GetLindormInstanceListResponseBodyInstanceList) GetEnableVector() *bool {
	return s.EnableVector
}

func (s *GetLindormInstanceListResponseBodyInstanceList) GetEngineType() *string {
	return s.EngineType
}

func (s *GetLindormInstanceListResponseBodyInstanceList) GetExpireTime() *string {
	return s.ExpireTime
}

func (s *GetLindormInstanceListResponseBodyInstanceList) GetExpiredMilliseconds() *int64 {
	return s.ExpiredMilliseconds
}

func (s *GetLindormInstanceListResponseBodyInstanceList) GetInstanceAlias() *string {
	return s.InstanceAlias
}

func (s *GetLindormInstanceListResponseBodyInstanceList) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetLindormInstanceListResponseBodyInstanceList) GetInstanceStatus() *string {
	return s.InstanceStatus
}

func (s *GetLindormInstanceListResponseBodyInstanceList) GetInstanceStorage() *string {
	return s.InstanceStorage
}

func (s *GetLindormInstanceListResponseBodyInstanceList) GetNetworkType() *string {
	return s.NetworkType
}

func (s *GetLindormInstanceListResponseBodyInstanceList) GetPayType() *string {
	return s.PayType
}

func (s *GetLindormInstanceListResponseBodyInstanceList) GetRegionId() *string {
	return s.RegionId
}

func (s *GetLindormInstanceListResponseBodyInstanceList) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *GetLindormInstanceListResponseBodyInstanceList) GetServiceType() *string {
	return s.ServiceType
}

func (s *GetLindormInstanceListResponseBodyInstanceList) GetTags() []*GetLindormInstanceListResponseBodyInstanceListTags {
	return s.Tags
}

func (s *GetLindormInstanceListResponseBodyInstanceList) GetVpcId() *string {
	return s.VpcId
}

func (s *GetLindormInstanceListResponseBodyInstanceList) GetZoneId() *string {
	return s.ZoneId
}

func (s *GetLindormInstanceListResponseBodyInstanceList) SetAliUid(v int64) *GetLindormInstanceListResponseBodyInstanceList {
	s.AliUid = &v
	return s
}

func (s *GetLindormInstanceListResponseBodyInstanceList) SetCreateErrorCode(v string) *GetLindormInstanceListResponseBodyInstanceList {
	s.CreateErrorCode = &v
	return s
}

func (s *GetLindormInstanceListResponseBodyInstanceList) SetCreateMilliseconds(v int64) *GetLindormInstanceListResponseBodyInstanceList {
	s.CreateMilliseconds = &v
	return s
}

func (s *GetLindormInstanceListResponseBodyInstanceList) SetCreateTime(v string) *GetLindormInstanceListResponseBodyInstanceList {
	s.CreateTime = &v
	return s
}

func (s *GetLindormInstanceListResponseBodyInstanceList) SetEnableColumn(v bool) *GetLindormInstanceListResponseBodyInstanceList {
	s.EnableColumn = &v
	return s
}

func (s *GetLindormInstanceListResponseBodyInstanceList) SetEnableCompute(v bool) *GetLindormInstanceListResponseBodyInstanceList {
	s.EnableCompute = &v
	return s
}

func (s *GetLindormInstanceListResponseBodyInstanceList) SetEnableLts(v bool) *GetLindormInstanceListResponseBodyInstanceList {
	s.EnableLts = &v
	return s
}

func (s *GetLindormInstanceListResponseBodyInstanceList) SetEnableMessage(v bool) *GetLindormInstanceListResponseBodyInstanceList {
	s.EnableMessage = &v
	return s
}

func (s *GetLindormInstanceListResponseBodyInstanceList) SetEnableRow(v bool) *GetLindormInstanceListResponseBodyInstanceList {
	s.EnableRow = &v
	return s
}

func (s *GetLindormInstanceListResponseBodyInstanceList) SetEnableStream(v bool) *GetLindormInstanceListResponseBodyInstanceList {
	s.EnableStream = &v
	return s
}

func (s *GetLindormInstanceListResponseBodyInstanceList) SetEnableVector(v bool) *GetLindormInstanceListResponseBodyInstanceList {
	s.EnableVector = &v
	return s
}

func (s *GetLindormInstanceListResponseBodyInstanceList) SetEngineType(v string) *GetLindormInstanceListResponseBodyInstanceList {
	s.EngineType = &v
	return s
}

func (s *GetLindormInstanceListResponseBodyInstanceList) SetExpireTime(v string) *GetLindormInstanceListResponseBodyInstanceList {
	s.ExpireTime = &v
	return s
}

func (s *GetLindormInstanceListResponseBodyInstanceList) SetExpiredMilliseconds(v int64) *GetLindormInstanceListResponseBodyInstanceList {
	s.ExpiredMilliseconds = &v
	return s
}

func (s *GetLindormInstanceListResponseBodyInstanceList) SetInstanceAlias(v string) *GetLindormInstanceListResponseBodyInstanceList {
	s.InstanceAlias = &v
	return s
}

func (s *GetLindormInstanceListResponseBodyInstanceList) SetInstanceId(v string) *GetLindormInstanceListResponseBodyInstanceList {
	s.InstanceId = &v
	return s
}

func (s *GetLindormInstanceListResponseBodyInstanceList) SetInstanceStatus(v string) *GetLindormInstanceListResponseBodyInstanceList {
	s.InstanceStatus = &v
	return s
}

func (s *GetLindormInstanceListResponseBodyInstanceList) SetInstanceStorage(v string) *GetLindormInstanceListResponseBodyInstanceList {
	s.InstanceStorage = &v
	return s
}

func (s *GetLindormInstanceListResponseBodyInstanceList) SetNetworkType(v string) *GetLindormInstanceListResponseBodyInstanceList {
	s.NetworkType = &v
	return s
}

func (s *GetLindormInstanceListResponseBodyInstanceList) SetPayType(v string) *GetLindormInstanceListResponseBodyInstanceList {
	s.PayType = &v
	return s
}

func (s *GetLindormInstanceListResponseBodyInstanceList) SetRegionId(v string) *GetLindormInstanceListResponseBodyInstanceList {
	s.RegionId = &v
	return s
}

func (s *GetLindormInstanceListResponseBodyInstanceList) SetResourceGroupId(v string) *GetLindormInstanceListResponseBodyInstanceList {
	s.ResourceGroupId = &v
	return s
}

func (s *GetLindormInstanceListResponseBodyInstanceList) SetServiceType(v string) *GetLindormInstanceListResponseBodyInstanceList {
	s.ServiceType = &v
	return s
}

func (s *GetLindormInstanceListResponseBodyInstanceList) SetTags(v []*GetLindormInstanceListResponseBodyInstanceListTags) *GetLindormInstanceListResponseBodyInstanceList {
	s.Tags = v
	return s
}

func (s *GetLindormInstanceListResponseBodyInstanceList) SetVpcId(v string) *GetLindormInstanceListResponseBodyInstanceList {
	s.VpcId = &v
	return s
}

func (s *GetLindormInstanceListResponseBodyInstanceList) SetZoneId(v string) *GetLindormInstanceListResponseBodyInstanceList {
	s.ZoneId = &v
	return s
}

func (s *GetLindormInstanceListResponseBodyInstanceList) Validate() error {
	if s.Tags != nil {
		for _, item := range s.Tags {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetLindormInstanceListResponseBodyInstanceListTags struct {
	// The tag key.
	//
	// example:
	//
	// test
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value.
	//
	// example:
	//
	// 2.2.18
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetLindormInstanceListResponseBodyInstanceListTags) String() string {
	return dara.Prettify(s)
}

func (s GetLindormInstanceListResponseBodyInstanceListTags) GoString() string {
	return s.String()
}

func (s *GetLindormInstanceListResponseBodyInstanceListTags) GetKey() *string {
	return s.Key
}

func (s *GetLindormInstanceListResponseBodyInstanceListTags) GetValue() *string {
	return s.Value
}

func (s *GetLindormInstanceListResponseBodyInstanceListTags) SetKey(v string) *GetLindormInstanceListResponseBodyInstanceListTags {
	s.Key = &v
	return s
}

func (s *GetLindormInstanceListResponseBodyInstanceListTags) SetValue(v string) *GetLindormInstanceListResponseBodyInstanceListTags {
	s.Value = &v
	return s
}

func (s *GetLindormInstanceListResponseBodyInstanceListTags) Validate() error {
	return dara.Validate(s)
}
