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
	// The instances.
	InstanceList []*GetLindormInstanceListResponseBodyInstanceList `json:"InstanceList,omitempty" xml:"InstanceList,omitempty" type:"Repeated"`
	// The number of returned pages.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of instances that are returned on each page.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 1CA1FAFD-E8DC-51C2-AA7E-CA6E2D049BA0
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of returned instances.
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
	// The 16-digit AliUid of the Alibaba Cloud account that owns the instance.
	//
	// example:
	//
	// 164901546557****
	AliUid          *int64  `json:"AliUid,omitempty" xml:"AliUid,omitempty"`
	CreateErrorCode *string `json:"CreateErrorCode,omitempty" xml:"CreateErrorCode,omitempty"`
	// The time when the instance is created. This value is a UNIX timestamp that indicates the number of milliseconds that have elapsed since the epoch time January 1, 1970, 00:00:00 UTC.
	//
	// example:
	//
	// 1631772842000
	CreateMilliseconds *int64 `json:"CreateMilliseconds,omitempty" xml:"CreateMilliseconds,omitempty"`
	// The time when the instance is created.
	//
	// example:
	//
	// 2021-09-16 14:13:13
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// Indicates whether the column storage engine is enabled, returning:
	//
	// - **true**: Enabled. - **false**: Not enabled.
	//
	// example:
	//
	// true
	EnableColumn *bool `json:"EnableColumn,omitempty" xml:"EnableColumn,omitempty"`
	// Indicates whether LDPS is activated for the instance. Valid values:
	//
	// 	- **true**: LDPS is activated for the instance.
	//
	// 	- **false**: LDPS is not activated for the instance.
	//
	// example:
	//
	// true
	EnableCompute *bool `json:"EnableCompute,omitempty" xml:"EnableCompute,omitempty"`
	// Indicates whether the LTS engine is enabled, returning:
	//
	// - **true**: Enabled. - **false**: Not enabled.
	//
	// example:
	//
	// true
	EnableLts *bool `json:"EnableLts,omitempty" xml:"EnableLts,omitempty"`
	// Indicates whether the message engine is enabled, returning:
	//
	// - **true**: Enabled. - **false**: Not enabled.
	//
	// example:
	//
	// true
	EnableMessage *bool `json:"EnableMessage,omitempty" xml:"EnableMessage,omitempty"`
	// Indicates whether the table 3.0 storage engine is enabled, returning:
	//
	// true: Enabled. - false: Not enabled.
	//
	// example:
	//
	// true
	EnableRow *bool `json:"EnableRow,omitempty" xml:"EnableRow,omitempty"`
	// Indicates whether the Lindorm streaming engine is activated for the instance. Valid values:
	//
	// 	- **true**: The Lindorm streaming engine is activated for the instance.
	//
	// 	- **false**: The Lindorm streaming engine is not activated for the instance.
	//
	// example:
	//
	// true
	EnableStream *bool `json:"EnableStream,omitempty" xml:"EnableStream,omitempty"`
	// Whether the vector engine is enabled, returns:
	//
	// - **true**: Enabled. - **false**: Not enabled.
	//
	// example:
	//
	// true
	EnableVector *bool `json:"EnableVector,omitempty" xml:"EnableVector,omitempty"`
	// The engine supported by the instance. The engines are indicated by different numbers:
	//
	// 	- **1**: LindormSearch.
	//
	// 	- **2**: LindormTSDB.
	//
	// 	- **4**: LindormTable.
	//
	// 	- **8**: LindormDFS.
	//
	// > The value of this parameter is the sum of all numbers that indicate the engines supported by the instance. For example, if the value of this parameter is 15, which is the sum of 1, 2, 4, and 8, the instance supports all four engines. If the value of this parameter is 6, which is the sum of 2 and 4, the instance supports LindormTSDB and LindormTable.
	//
	// example:
	//
	// 15
	EngineType *string `json:"EngineType,omitempty" xml:"EngineType,omitempty"`
	// The time when the instance expires.
	//
	// > This parameter is returned only if the billing method of the instance is subscription.
	//
	// example:
	//
	// 2022-04-26 00:00:00
	ExpireTime *string `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	// The time when the instance expires. This value is a UNIX timestamp that indicates the number of milliseconds that have elapsed since the epoch time January 1, 1970, 00:00:00 UTC.
	//
	// example:
	//
	// 1650902400000
	ExpiredMilliseconds *int64 `json:"ExpiredMilliseconds,omitempty" xml:"ExpiredMilliseconds,omitempty"`
	// The name of the VPC.
	//
	// example:
	//
	// test
	InstanceAlias *string `json:"InstanceAlias,omitempty" xml:"InstanceAlias,omitempty"`
	// The ID of the instance
	//
	// example:
	//
	// ld-bp17pwu1541ia****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The status of the instance. Valid values:
	//
	// 	- **CREATING**: The instance is being created.
	//
	// 	- **ACTIVATION**: The instance is running.
	//
	// 	- **COLD_EXPANDING**: The Capacity storage of the instance is being scaled up.
	//
	// 	- **MINOR_VERSION_TRANSING**: The minor version of the instance is being updated.
	//
	// 	- **RESIZING**: The nodes in the instance are being scaled up.
	//
	// 	- **SHRINKING**: The nodes in the instance are being scaled down.
	//
	// 	- **CLASS_CHANGING**: The specification of the instance is being changed.
	//
	// 	- **SSL_SWITCHING: SSL**: The SSL configurations of the instance are being changed.
	//
	// 	- **CDC_OPENING**: Data subscription is being enabled for the instance.
	//
	// 	- **TRANSFER**: The data of the instance is being transferred.
	//
	// 	- **DATABASE_TRANSFER**: The data of the instance is being transferred to databases.
	//
	// 	- **GUARD_CREATING**: A disaster recovery instance is being created.
	//
	// 	- **BACKUP_RECOVERING**: The data of the instance is being restored from a backup.
	//
	// 	- **DATABASE_IMPORTING**: Data is being imported to the instance.
	//
	// 	- **NET_MODIFYING**: The network configurations of the instance are being changed.
	//
	// 	- **NET_SWITCHING**: The network of the instance is being switched between a virtual private cloud (VPC) and the Internet.
	//
	// 	- **NET_CREATING**: The connection to the instance is being created.
	//
	// 	- **NET_DELETING**: The connection to the instance is being deleted.
	//
	// 	- **DELETING**: The instance is being deleted.
	//
	// 	- **RESTARTING**: The instance is restarting.
	//
	// 	- **LOCKED**: The instance is locked because it expires.
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
	// 	- **PREPAY**: subscription.
	//
	// 	- **POSTPAY**: pay-as-you-go.
	//
	// example:
	//
	// PREPAY
	PayType *string `json:"PayType,omitempty" xml:"PayType,omitempty"`
	// The region ID of the instance.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group to which the instance belongs.
	//
	// example:
	//
	// rg-aekzledqeat****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The series of the instance. Valid values:
	//
	// 	- **lindorm**: The instance is a Lindorm instance.
	//
	// 	- **serverless_lindorm**: The instance is a Lindorm Serverless instance.
	//
	// 	- **lindorm_standalone**: The instance is a single-node Lindorm instance.
	//
	// 	- **lts**: The instance is an LTS instance.
	//
	// example:
	//
	// lindorm
	ServiceType *string `json:"ServiceType,omitempty" xml:"ServiceType,omitempty"`
	// The list of tags associated with the specified instances.
	Tags []*GetLindormInstanceListResponseBodyInstanceListTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	// The ID of the VPC in which the instance is deployed.
	//
	// example:
	//
	// vpc-bp1n3i15v90el48nx****
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	// The ID of the zone in which the instance is created.
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
	// The key of the tag.
	//
	// example:
	//
	// test
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of the tag.
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
