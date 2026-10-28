// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDBInstanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAutoRenew(v bool) *CreateDBInstanceRequest
	GetAutoRenew() *bool
	SetCNNodeCount(v int32) *CreateDBInstanceRequest
	GetCNNodeCount() *int32
	SetClientToken(v string) *CreateDBInstanceRequest
	GetClientToken() *string
	SetCnClass(v string) *CreateDBInstanceRequest
	GetCnClass() *string
	SetDBNodeClass(v string) *CreateDBInstanceRequest
	GetDBNodeClass() *string
	SetDBNodeCount(v int32) *CreateDBInstanceRequest
	GetDBNodeCount() *int32
	SetDNNodeCount(v int32) *CreateDBInstanceRequest
	GetDNNodeCount() *int32
	SetDescription(v string) *CreateDBInstanceRequest
	GetDescription() *string
	SetDnClass(v string) *CreateDBInstanceRequest
	GetDnClass() *string
	SetDnStorageSpace(v string) *CreateDBInstanceRequest
	GetDnStorageSpace() *string
	SetEngineVersion(v string) *CreateDBInstanceRequest
	GetEngineVersion() *string
	SetExtraParams(v map[string]*string) *CreateDBInstanceRequest
	GetExtraParams() map[string]*string
	SetIsColumnarReadDBInstance(v bool) *CreateDBInstanceRequest
	GetIsColumnarReadDBInstance() *bool
	SetIsReadDBInstance(v bool) *CreateDBInstanceRequest
	GetIsReadDBInstance() *bool
	SetNetworkType(v string) *CreateDBInstanceRequest
	GetNetworkType() *string
	SetOriginMinorVersion(v string) *CreateDBInstanceRequest
	GetOriginMinorVersion() *string
	SetPayType(v string) *CreateDBInstanceRequest
	GetPayType() *string
	SetPeriod(v string) *CreateDBInstanceRequest
	GetPeriod() *string
	SetPrimaryDBInstanceName(v string) *CreateDBInstanceRequest
	GetPrimaryDBInstanceName() *string
	SetPrimaryZone(v string) *CreateDBInstanceRequest
	GetPrimaryZone() *string
	SetRegionId(v string) *CreateDBInstanceRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *CreateDBInstanceRequest
	GetResourceGroupId() *string
	SetSecondaryZone(v string) *CreateDBInstanceRequest
	GetSecondaryZone() *string
	SetSeries(v string) *CreateDBInstanceRequest
	GetSeries() *string
	SetStorageType(v string) *CreateDBInstanceRequest
	GetStorageType() *string
	SetTertiaryZone(v string) *CreateDBInstanceRequest
	GetTertiaryZone() *string
	SetTopologyType(v string) *CreateDBInstanceRequest
	GetTopologyType() *string
	SetUsedTime(v int32) *CreateDBInstanceRequest
	GetUsedTime() *int32
	SetVPCId(v string) *CreateDBInstanceRequest
	GetVPCId() *string
	SetVSwitchId(v string) *CreateDBInstanceRequest
	GetVSwitchId() *string
	SetZoneId(v string) *CreateDBInstanceRequest
	GetZoneId() *string
}

type CreateDBInstanceRequest struct {
	// Specifies whether to enable auto-renewal. Default value: true. Valid values:
	//
	// - **true**: Enables auto-renewal.
	//
	// - **false**: Disables auto-renewal.
	//
	// example:
	//
	// true
	AutoRenew *bool `json:"AutoRenew,omitempty" xml:"AutoRenew,omitempty"`
	// The number of compute nodes.
	//
	// example:
	//
	// 2
	CNNodeCount *int32 `json:"CNNodeCount,omitempty" xml:"CNNodeCount,omitempty"`
	// The client token that is used to ensure the idempotence of the request. You must specify a unique token for each request.
	//
	// example:
	//
	// xxxxxx-xxx
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The specifications of compute nodes. Required when you create an Enterprise Edition instance. Not required for a Standard Edition instance.
	//
	// Local node specifications for Enterprise Edition:
	//
	// - **polarx.x4.medium.2e**: 2 vCPUs, 8 GB (general-purpose)
	//
	// - **polarx.x4.large.2e**: 4 vCPUs, 16 GB (general-purpose)
	//
	// - **polarx.x4.xlarge.2e**: 8 vCPUs, 32 GB (general-purpose)
	//
	// - **polarx.x4.2xlarge.2e**: 16 vCPUs, 64 GB (general-purpose)
	//
	// - **polarx.x8.large.2e**: 4 vCPUs, 32 GB (dedicated)
	//
	// - **polarx.x2.large.2x**: 8 vCPUs, 16 GB (dedicated)
	//
	// - **polarx.x4.xlarge.2x**: 8 vCPUs, 32 GB (dedicated)
	//
	// - **polarx.x8.xlarge.2e**: 8 vCPUs, 64 GB (dedicated)
	//
	// - **polarx.x8.2xlarge.2e**: 16 vCPUs, 128 GB (dedicated)
	//
	// - **polarx.x4.4xlarge.2e**: 32 vCPUs, 128 GB (dedicated)
	//
	// - **polarx.x8.4xlarge.2e**: 32 vCPUs, 256 GB (dedicated)
	//
	// - **polarx.st.8xlarge.2e**: 60 vCPUs, 470 GB (dedicated)
	//
	// - **polarx.st.12xlarge.2e**: 90 vCPUs, 720 GB (dedicated)
	//
	// Cloud node specifications for Enterprise Edition:
	//
	// - **polarx.x4.medium.c2e**: 2 vCPUs, 8 GB (general-purpose)
	//
	// - **polarx.x4.large.c2e**: 4 vCPUs, 16 GB (general-purpose)
	//
	// - **polarx.x4.xlarge.c2e**: 8 vCPUs, 32 GB (general-purpose)
	//
	// - **polarx.x4.2xlarge.c2e**: 16 vCPUs, 64 GB (general-purpose)
	//
	// - **polarx.x8.large.c2e**: 4 vCPUs, 32 GB (dedicated)
	//
	// - **polarx.x2.large.c2x**: 8 vCPUs, 16 GB (dedicated)
	//
	// - **polarx.x4.xlarge.c2x**: 8 vCPUs, 32 GB (dedicated)
	//
	// - **polarx.x8.xlarge.c2e**: 8 vCPUs, 64 GB (dedicated)
	//
	// - **polarx.x8.2xlarge.c2e**: 16 vCPUs, 128 GB (dedicated)
	//
	// - **polarx.x4.4xlarge.c2e**: 32 vCPUs, 128 GB (dedicated)
	//
	// - **polarx.x8.4xlarge.c2e**: 32 vCPUs, 256 GB (dedicated)
	//
	// - **polarx.st.8xlarge.c2e**: 60 vCPUs, 470 GB (dedicated)
	//
	// - **polarx.st.12xlarge.c2e**: 90 vCPUs, 720 GB (dedicated)
	//
	// example:
	//
	// polarx.x4.medium.2e
	CnClass *string `json:"CnClass,omitempty" xml:"CnClass,omitempty"`
	// Required when you create a Standard Edition instance. Not required for an Enterprise Edition instance.
	//
	// Local disk specifications for Standard Edition:
	//
	// - **mysql.n8.small.25**: 1 vCPU, 8 GB (general-purpose)
	//
	// - **mysql.n2.medium.25**: 2 vCPUs, 4 GB (general-purpose)
	//
	// - **mysql.n4.medium.25**: 2 vCPUs, 8 GB (general-purpose)
	//
	// - **mysql.n8.medium.25**: 2 vCPUs, 16 GB (general-purpose)
	//
	// - **mysql.n2.large.25**: 4 vCPUs, 8 GB (general-purpose)
	//
	// - **mysql.n4.large.25**: 4 vCPUs, 16 GB (general-purpose)
	//
	// - **mysql.n8.large.25**: 4 vCPUs, 32 GB (general-purpose)
	//
	// - **mysql.n2.xlarge.25**: 8 vCPUs, 16 GB (general-purpose)
	//
	// - **mysql.n4.xlarge.25**: 8 vCPUs, 32 GB (general-purpose)
	//
	// - **mysql.n8.xlarge.25**: 8 vCPUs, 64 GB (general-purpose)
	//
	// - **mysql.n2.2xlarge.25**: 16 vCPUs, 32 GB (general-purpose)
	//
	// - **mysql.n4.2xlarge.25**: 16 vCPUs, 64 GB (general-purpose)
	//
	// - **mysql.n8.2xlarge.25**: 16 vCPUs, 128 GB (general-purpose)
	//
	// - **mysql.x2.medium.25**: 2 vCPUs, 4 GB (dedicated)
	//
	// - **mysql.x4.medium.25**: 2 vCPUs, 8 GB (dedicated)
	//
	// - **mysql.x8.medium.25**: 2 vCPUs, 16 GB (dedicated)
	//
	// - **mysql.x2.large.25**: 4 vCPUs, 8 GB (dedicated)
	//
	// - **mysql.x4.large.25**: 4 vCPUs, 16 GB (dedicated)
	//
	// - **mysql.x8.large.25**: 4 vCPUs, 32 GB (dedicated)
	//
	// - **mysql.x2.xlarge.25**: 8 vCPUs, 16 GB (dedicated)
	//
	// - **mysql.x4.xlarge.25**: 8 vCPUs, 32 GB (dedicated)
	//
	// - **mysql.x8.xlarge.25**: 8 vCPUs, 64 GB (dedicated)
	//
	// - **mysql.x2.2xlarge.25**: 16 vCPUs, 32 GB (dedicated)
	//
	// - **mysql.x4.2xlarge.25**: 16 vCPUs, 64 GB (dedicated)
	//
	// - **mysql.x8.2xlarge.25**: 16 vCPUs, 128 GB (dedicated)
	//
	// - **mysql.x4.4xlarge.25**: 32 vCPUs, 128 GB (dedicated)
	//
	// - **mysql.x8.4xlarge.25**: 32 vCPUs, 256 GB (dedicated)
	//
	// - **mysql.x4.8xlarge.25**: 64 vCPUs, 256 GB (dedicated)
	//
	// - **mysql.x8.8xlarge.25**: 64 vCPUs, 512 GB (dedicated)
	//
	// - **mysql.st.12xlarge.25**: 90 vCPUs, 720 GB (dedicated)
	//
	// Cloud disk specifications for Standard Edition:
	//
	// - **polarx.mysql.n2.medium.c25**: 2 vCPUs, 4 GB (general-purpose)
	//
	// - **polarx.mysql.n4.medium.c25**: 2 vCPUs, 8 GB (general-purpose)
	//
	// - **polarx.mysql.n8.medium.c25**: 2 vCPUs, 16 GB (general-purpose)
	//
	// - **polarx.mysql.n2.large.c25**: 4 vCPUs, 8 GB (general-purpose)
	//
	// - **polarx.mysql.n4.large.c25**: 4 vCPUs, 16 GB (general-purpose)
	//
	// - **polarx.mysql.n8.large.c25**: 4 vCPUs, 32 GB (general-purpose)
	//
	// - **polarx.mysql.n2.xlarge.c25**: 8 vCPUs, 16 GB (general-purpose)
	//
	// - **polarx.mysql.n4.xlarge.c25**: 8 vCPUs, 32 GB (general-purpose)
	//
	// - **polarx.mysql.n8.xlarge.c25**: 8 vCPUs, 64 GB (general-purpose)
	//
	// - **polarx.mysql.x2.medium.c25**: 2 vCPUs, 4 GB (dedicated)
	//
	// - **polarx.mysql.x4.medium.c25**: 2 vCPUs, 8 GB (dedicated)
	//
	// - **polarx.mysql.x8.medium.c25**: 2 vCPUs, 16 GB (dedicated)
	//
	// - **polarx.mysql.x2.large.c25**: 4 vCPUs, 8 GB (dedicated)
	//
	// - **polarx.mysql.x4.large.c25**: 4 vCPUs, 16 GB (dedicated)
	//
	// - **polarx.mysql.x8.large.c25**: 4 vCPUs, 32 GB (dedicated)
	//
	// - **polarx.mysql.x2.xlarge.c25**: 8 vCPUs, 16 GB (dedicated)
	//
	// - **polarx.mysql.x4.xlarge.c25**: 8 vCPUs, 32 GB (dedicated)
	//
	// - **polarx.mysql.x8.xlarge.c25**: 8 vCPUs, 64 GB (dedicated)
	//
	// - **polarx.mysql.x2.2xlarge.c25**: 16 vCPUs, 32 GB (dedicated)
	//
	// - **polarx.mysql.x4.2xlarge.c25**: 16 vCPUs, 64 GB (dedicated)
	//
	// - **polarx.mysql.x8.2xlarge.c25**: 16 vCPUs, 128 GB (dedicated)
	//
	// - **polarx.mysql.x2.4xlarge.c25**: 32 vCPUs, 64 GB (dedicated)
	//
	// - **polarx.mysql.x4.4xlarge.c25**: 32 vCPUs, 128 GB (dedicated)
	//
	// - **polarx.mysql.x8.4xlarge.c25**: 32 vCPUs, 256 GB (dedicated)
	//
	// - **polarx.mysql.x2.8xlarge.c25**: 64 vCPUs, 128 GB (dedicated)
	//
	// - **polarx.mysql.x4.8xlarge.c25**: 64 vCPUs, 256 GB (dedicated)
	//
	// - **polarx.mysql.x8.8xlarge.c25**: 64 vCPUs, 512 GB (dedicated)
	//
	// example:
	//
	// polarx.x4.2xlarge.2d
	DBNodeClass *string `json:"DBNodeClass,omitempty" xml:"DBNodeClass,omitempty"`
	// The number of nodes in the instance. Valid values:
	//
	// 	- Standard Edition: 1
	//
	// 	- Enterprise Edition: 2 or greater
	//
	// example:
	//
	// 3
	DBNodeCount *int32 `json:"DBNodeCount,omitempty" xml:"DBNodeCount,omitempty"`
	// The number of storage nodes.
	//
	// example:
	//
	// 2
	DNNodeCount *int32 `json:"DNNodeCount,omitempty" xml:"DNNodeCount,omitempty"`
	// The description of the database.
	//
	// example:
	//
	// ods_api_apidata_mobilegame_pay_api_di
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The specifications of storage nodes. Required when you create an Enterprise Edition instance. Not required for a Standard Edition instance.
	//
	// Local node specifications for Enterprise Edition:
	//
	// - **mysql.n2.medium.25**: 2 vCPUs, 4 GB (general-purpose)
	//
	// - **mysql.n4.medium.25**: 2 vCPUs, 8 GB (general-purpose)
	//
	// - **mysql.n2.large.25**: 4 vCPUs, 8 GB (general-purpose)
	//
	// - **mysql.n4.large.25**: 4 vCPUs, 16 GB (general-purpose)
	//
	// - **mysql.n4.xlarge.25**: 8 vCPUs, 32 GB (general-purpose)
	//
	// - **mysql.n4.2xlarge.25**: 16 vCPUs, 64 GB (general-purpose)
	//
	// - **mysql.x4.large.25**: 4 vCPUs, 16 GB (dedicated)
	//
	// - **mysql.x8.large.25**: 4 vCPUs, 32 GB (dedicated)
	//
	// - **mysql.x2.xlarge.25**: 8 vCPUs, 16 GB (dedicated)
	//
	// - **mysql.x8.xlarge.25**: 8 vCPUs, 64 GB (dedicated)
	//
	// - **mysql.x8.2xlarge.25**: 16 vCPUs, 128 GB (dedicated)
	//
	// - **mysql.x4.4xlarge.25**: 32 vCPUs, 128 GB (dedicated)
	//
	// - **mysql.x8.4xlarge.25**: 32 vCPUs, 256 GB (dedicated)
	//
	// - **mysql.st.8xlarge.25**: 60 vCPUs, 470 GB (dedicated)
	//
	// - **mysql.st.12xlarge.25**: 90 vCPUs, 720 GB (dedicated)
	//
	// - **mysql.x8.45xlarge.25**: 180 vCPUs, 1440 GB (dedicated)
	//
	// - **mysql.x8.60xlarge.25**: 240 vCPUs, 1920 GB (dedicated)
	//
	// Cloud node specifications for Enterprise Edition:
	//
	// - **polarx.mysql.n2.medium.c25**: 2 vCPUs, 4 GB (general-purpose)
	//
	// - **polarx.mysql.n4.medium.c25**: 2 vCPUs, 8 GB (general-purpose)
	//
	// - **polarx.mysql.n2.large.c25**: 4 vCPUs, 8 GB (general-purpose)
	//
	// - **polarx.mysql.n4.large.c25**: 4 vCPUs, 16 GB (general-purpose)
	//
	// - **polarx.mysql.n4.xlarge.c25**: 8 vCPUs, 32 GB (general-purpose)
	//
	// - **polarx.mysql.n4.2xlarge.c25**: 16 vCPUs, 64 GB (general-purpose)
	//
	// - **polarx.mysql.x4.large.c25**: 4 vCPUs, 16 GB (dedicated)
	//
	// - **polarx.mysql.x8.large.c25**: 4 vCPUs, 32 GB (dedicated)
	//
	// - **polarx.mysql.x2.xlarge.c25**: 8 vCPUs, 16 GB (dedicated)
	//
	// - **polarx.mysql.x8.xlarge.c25**: 8 vCPUs, 64 GB (dedicated)
	//
	// - **polarx.mysql.x8.2xlarge.c25**: 16 vCPUs, 128 GB (dedicated)
	//
	// - **polarx.mysql.x4.4xlarge.c25**: 32 vCPUs, 128 GB (dedicated)
	//
	// - **polarx.mysql.x8.4xlarge.c25**: 32 vCPUs, 256 GB (dedicated)
	//
	// - **polarx.mysql.st.8xlarge.c25**: 60 vCPUs, 470 GB (dedicated)
	//
	// - **polarx.mysql.st.12xlarge.c25**: 90 vCPUs, 720 GB (dedicated)
	//
	// - **polarx.mysql.x8.45xlarge.c25**: 180 vCPUs, 1440 GB (dedicated)
	//
	// - **polarx.mysql.x8.60xlarge.c25**: 240 vCPUs, 1920 GB (dedicated)
	//
	// example:
	//
	// mysql.n4.medium.25
	DnClass *string `json:"DnClass,omitempty" xml:"DnClass,omitempty"`
	// The disk space of the storage nodes.
	//
	// example:
	//
	// 1500
	DnStorageSpace *string `json:"DnStorageSpace,omitempty" xml:"DnStorageSpace,omitempty"`
	// The MySQL DPI engine version. Valid values: 5.7 and 8.0.
	//
	// This parameter is required.
	//
	// example:
	//
	// 5.7
	EngineVersion *string `json:"EngineVersion,omitempty" xml:"EngineVersion,omitempty"`
	// The extra parameters.
	ExtraParams map[string]*string `json:"ExtraParams,omitempty" xml:"ExtraParams,omitempty"`
	// Specifies whether the instance to be created is a column store read-only instance.
	//
	// example:
	//
	// true
	IsColumnarReadDBInstance *bool `json:"IsColumnarReadDBInstance,omitempty" xml:"IsColumnarReadDBInstance,omitempty"`
	// Specifies whether the instance is a read-only instance. Valid values:
	//
	// - **true**: Yes.
	//
	// - **false**: No.
	//
	// example:
	//
	// false
	IsReadDBInstance *bool `json:"IsReadDBInstance,omitempty" xml:"IsReadDBInstance,omitempty"`
	// The network type of the instance. Only VPC is supported.
	//
	// example:
	//
	// vpc
	NetworkType *string `json:"NetworkType,omitempty" xml:"NetworkType,omitempty"`
	// The minor engine version. This parameter is generally not required.
	//
	// example:
	//
	// polardb-2.4.0_standard_xcluster8.4.20-20260331
	OriginMinorVersion *string `json:"OriginMinorVersion,omitempty" xml:"OriginMinorVersion,omitempty"`
	// The billing method of the instance. Valid values:
	//
	// - **PREPAY**: subscription
	//
	// - **POSTPAY**: pay-as-you-go
	//
	// This parameter is required.
	//
	// example:
	//
	// PREPAY
	PayType *string `json:"PayType,omitempty" xml:"PayType,omitempty"`
	// The billing cycle. Valid values for subscription instances: Year and Month. Default value for pay-as-you-go instances: Hour.
	//
	// example:
	//
	// Month
	Period *string `json:"Period,omitempty" xml:"Period,omitempty"`
	// The name of the primary instance. This parameter is required if the instance is a read-only instance.
	//
	// example:
	//
	// pxc-*********
	PrimaryDBInstanceName *string `json:"PrimaryDBInstanceName,omitempty" xml:"PrimaryDBInstanceName,omitempty"`
	// The primary zone.
	//
	// example:
	//
	// cn-shenzhen-e
	PrimaryZone *string `json:"PrimaryZone,omitempty" xml:"PrimaryZone,omitempty"`
	// The region of the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The resource group ID. This parameter is not supported and can be left empty.
	//
	// example:
	//
	// null
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The secondary zone.
	//
	// example:
	//
	// cn-shenzhen-a
	SecondaryZone *string `json:"SecondaryZone,omitempty" xml:"SecondaryZone,omitempty"`
	// The edition of the instance. Valid values:
	//
	// - enterprise: Enterprise Edition
	//
	// - standard: Standard Edition
	//
	// example:
	//
	// enterprise
	Series *string `json:"Series,omitempty" xml:"Series,omitempty"`
	// The storage type. Valid values:
	//
	// 	- custom_local_ssd: local disk
	//
	// 	- cloud_auto: cloud disk
	//
	// example:
	//
	// cloud_auto
	StorageType *string `json:"StorageType,omitempty" xml:"StorageType,omitempty"`
	// The third active zone in a Three-zone deployment.
	//
	// example:
	//
	// cn-shenzhen-e
	TertiaryZone *string `json:"TertiaryZone,omitempty" xml:"TertiaryZone,omitempty"`
	// The topology type. Valid values:
	//
	// - **3azones**: Three-zone deployment (three active zones).
	//
	// - **1azone**: Single active zone.
	//
	// This parameter is required.
	//
	// example:
	//
	// 3azones
	TopologyType *string `json:"TopologyType,omitempty" xml:"TopologyType,omitempty"`
	// The subscription duration. This parameter is required when the billing method is set to subscription.
	//
	// > If Period is set to Year, valid values are 1, 2, and 3.
	//
	// example:
	//
	// 1
	UsedTime *int32 `json:"UsedTime,omitempty" xml:"UsedTime,omitempty"`
	// The VPC ID.
	//
	// example:
	//
	// vpc-*****
	VPCId *string `json:"VPCId,omitempty" xml:"VPCId,omitempty"`
	// The vSwitch ID.
	//
	// example:
	//
	// vsw-*********
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// The zone of the instance.
	//
	// example:
	//
	// cn-hangzhou-a
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s CreateDBInstanceRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateDBInstanceRequest) GoString() string {
	return s.String()
}

func (s *CreateDBInstanceRequest) GetAutoRenew() *bool {
	return s.AutoRenew
}

func (s *CreateDBInstanceRequest) GetCNNodeCount() *int32 {
	return s.CNNodeCount
}

func (s *CreateDBInstanceRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateDBInstanceRequest) GetCnClass() *string {
	return s.CnClass
}

func (s *CreateDBInstanceRequest) GetDBNodeClass() *string {
	return s.DBNodeClass
}

func (s *CreateDBInstanceRequest) GetDBNodeCount() *int32 {
	return s.DBNodeCount
}

func (s *CreateDBInstanceRequest) GetDNNodeCount() *int32 {
	return s.DNNodeCount
}

func (s *CreateDBInstanceRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateDBInstanceRequest) GetDnClass() *string {
	return s.DnClass
}

func (s *CreateDBInstanceRequest) GetDnStorageSpace() *string {
	return s.DnStorageSpace
}

func (s *CreateDBInstanceRequest) GetEngineVersion() *string {
	return s.EngineVersion
}

func (s *CreateDBInstanceRequest) GetExtraParams() map[string]*string {
	return s.ExtraParams
}

func (s *CreateDBInstanceRequest) GetIsColumnarReadDBInstance() *bool {
	return s.IsColumnarReadDBInstance
}

func (s *CreateDBInstanceRequest) GetIsReadDBInstance() *bool {
	return s.IsReadDBInstance
}

func (s *CreateDBInstanceRequest) GetNetworkType() *string {
	return s.NetworkType
}

func (s *CreateDBInstanceRequest) GetOriginMinorVersion() *string {
	return s.OriginMinorVersion
}

func (s *CreateDBInstanceRequest) GetPayType() *string {
	return s.PayType
}

func (s *CreateDBInstanceRequest) GetPeriod() *string {
	return s.Period
}

func (s *CreateDBInstanceRequest) GetPrimaryDBInstanceName() *string {
	return s.PrimaryDBInstanceName
}

func (s *CreateDBInstanceRequest) GetPrimaryZone() *string {
	return s.PrimaryZone
}

func (s *CreateDBInstanceRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateDBInstanceRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *CreateDBInstanceRequest) GetSecondaryZone() *string {
	return s.SecondaryZone
}

func (s *CreateDBInstanceRequest) GetSeries() *string {
	return s.Series
}

func (s *CreateDBInstanceRequest) GetStorageType() *string {
	return s.StorageType
}

func (s *CreateDBInstanceRequest) GetTertiaryZone() *string {
	return s.TertiaryZone
}

func (s *CreateDBInstanceRequest) GetTopologyType() *string {
	return s.TopologyType
}

func (s *CreateDBInstanceRequest) GetUsedTime() *int32 {
	return s.UsedTime
}

func (s *CreateDBInstanceRequest) GetVPCId() *string {
	return s.VPCId
}

func (s *CreateDBInstanceRequest) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *CreateDBInstanceRequest) GetZoneId() *string {
	return s.ZoneId
}

func (s *CreateDBInstanceRequest) SetAutoRenew(v bool) *CreateDBInstanceRequest {
	s.AutoRenew = &v
	return s
}

func (s *CreateDBInstanceRequest) SetCNNodeCount(v int32) *CreateDBInstanceRequest {
	s.CNNodeCount = &v
	return s
}

func (s *CreateDBInstanceRequest) SetClientToken(v string) *CreateDBInstanceRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateDBInstanceRequest) SetCnClass(v string) *CreateDBInstanceRequest {
	s.CnClass = &v
	return s
}

func (s *CreateDBInstanceRequest) SetDBNodeClass(v string) *CreateDBInstanceRequest {
	s.DBNodeClass = &v
	return s
}

func (s *CreateDBInstanceRequest) SetDBNodeCount(v int32) *CreateDBInstanceRequest {
	s.DBNodeCount = &v
	return s
}

func (s *CreateDBInstanceRequest) SetDNNodeCount(v int32) *CreateDBInstanceRequest {
	s.DNNodeCount = &v
	return s
}

func (s *CreateDBInstanceRequest) SetDescription(v string) *CreateDBInstanceRequest {
	s.Description = &v
	return s
}

func (s *CreateDBInstanceRequest) SetDnClass(v string) *CreateDBInstanceRequest {
	s.DnClass = &v
	return s
}

func (s *CreateDBInstanceRequest) SetDnStorageSpace(v string) *CreateDBInstanceRequest {
	s.DnStorageSpace = &v
	return s
}

func (s *CreateDBInstanceRequest) SetEngineVersion(v string) *CreateDBInstanceRequest {
	s.EngineVersion = &v
	return s
}

func (s *CreateDBInstanceRequest) SetExtraParams(v map[string]*string) *CreateDBInstanceRequest {
	s.ExtraParams = v
	return s
}

func (s *CreateDBInstanceRequest) SetIsColumnarReadDBInstance(v bool) *CreateDBInstanceRequest {
	s.IsColumnarReadDBInstance = &v
	return s
}

func (s *CreateDBInstanceRequest) SetIsReadDBInstance(v bool) *CreateDBInstanceRequest {
	s.IsReadDBInstance = &v
	return s
}

func (s *CreateDBInstanceRequest) SetNetworkType(v string) *CreateDBInstanceRequest {
	s.NetworkType = &v
	return s
}

func (s *CreateDBInstanceRequest) SetOriginMinorVersion(v string) *CreateDBInstanceRequest {
	s.OriginMinorVersion = &v
	return s
}

func (s *CreateDBInstanceRequest) SetPayType(v string) *CreateDBInstanceRequest {
	s.PayType = &v
	return s
}

func (s *CreateDBInstanceRequest) SetPeriod(v string) *CreateDBInstanceRequest {
	s.Period = &v
	return s
}

func (s *CreateDBInstanceRequest) SetPrimaryDBInstanceName(v string) *CreateDBInstanceRequest {
	s.PrimaryDBInstanceName = &v
	return s
}

func (s *CreateDBInstanceRequest) SetPrimaryZone(v string) *CreateDBInstanceRequest {
	s.PrimaryZone = &v
	return s
}

func (s *CreateDBInstanceRequest) SetRegionId(v string) *CreateDBInstanceRequest {
	s.RegionId = &v
	return s
}

func (s *CreateDBInstanceRequest) SetResourceGroupId(v string) *CreateDBInstanceRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *CreateDBInstanceRequest) SetSecondaryZone(v string) *CreateDBInstanceRequest {
	s.SecondaryZone = &v
	return s
}

func (s *CreateDBInstanceRequest) SetSeries(v string) *CreateDBInstanceRequest {
	s.Series = &v
	return s
}

func (s *CreateDBInstanceRequest) SetStorageType(v string) *CreateDBInstanceRequest {
	s.StorageType = &v
	return s
}

func (s *CreateDBInstanceRequest) SetTertiaryZone(v string) *CreateDBInstanceRequest {
	s.TertiaryZone = &v
	return s
}

func (s *CreateDBInstanceRequest) SetTopologyType(v string) *CreateDBInstanceRequest {
	s.TopologyType = &v
	return s
}

func (s *CreateDBInstanceRequest) SetUsedTime(v int32) *CreateDBInstanceRequest {
	s.UsedTime = &v
	return s
}

func (s *CreateDBInstanceRequest) SetVPCId(v string) *CreateDBInstanceRequest {
	s.VPCId = &v
	return s
}

func (s *CreateDBInstanceRequest) SetVSwitchId(v string) *CreateDBInstanceRequest {
	s.VSwitchId = &v
	return s
}

func (s *CreateDBInstanceRequest) SetZoneId(v string) *CreateDBInstanceRequest {
	s.ZoneId = &v
	return s
}

func (s *CreateDBInstanceRequest) Validate() error {
	return dara.Validate(s)
}
