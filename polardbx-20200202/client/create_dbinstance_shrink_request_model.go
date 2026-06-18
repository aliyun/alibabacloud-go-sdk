// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDBInstanceShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAutoRenew(v bool) *CreateDBInstanceShrinkRequest
	GetAutoRenew() *bool
	SetCNNodeCount(v int32) *CreateDBInstanceShrinkRequest
	GetCNNodeCount() *int32
	SetClientToken(v string) *CreateDBInstanceShrinkRequest
	GetClientToken() *string
	SetCnClass(v string) *CreateDBInstanceShrinkRequest
	GetCnClass() *string
	SetDBNodeClass(v string) *CreateDBInstanceShrinkRequest
	GetDBNodeClass() *string
	SetDBNodeCount(v int32) *CreateDBInstanceShrinkRequest
	GetDBNodeCount() *int32
	SetDNNodeCount(v int32) *CreateDBInstanceShrinkRequest
	GetDNNodeCount() *int32
	SetDescription(v string) *CreateDBInstanceShrinkRequest
	GetDescription() *string
	SetDnClass(v string) *CreateDBInstanceShrinkRequest
	GetDnClass() *string
	SetDnStorageSpace(v string) *CreateDBInstanceShrinkRequest
	GetDnStorageSpace() *string
	SetEngineVersion(v string) *CreateDBInstanceShrinkRequest
	GetEngineVersion() *string
	SetExtraParamsShrink(v string) *CreateDBInstanceShrinkRequest
	GetExtraParamsShrink() *string
	SetIsColumnarReadDBInstance(v bool) *CreateDBInstanceShrinkRequest
	GetIsColumnarReadDBInstance() *bool
	SetIsReadDBInstance(v bool) *CreateDBInstanceShrinkRequest
	GetIsReadDBInstance() *bool
	SetNetworkType(v string) *CreateDBInstanceShrinkRequest
	GetNetworkType() *string
	SetOriginMinorVersion(v string) *CreateDBInstanceShrinkRequest
	GetOriginMinorVersion() *string
	SetPayType(v string) *CreateDBInstanceShrinkRequest
	GetPayType() *string
	SetPeriod(v string) *CreateDBInstanceShrinkRequest
	GetPeriod() *string
	SetPrimaryDBInstanceName(v string) *CreateDBInstanceShrinkRequest
	GetPrimaryDBInstanceName() *string
	SetPrimaryZone(v string) *CreateDBInstanceShrinkRequest
	GetPrimaryZone() *string
	SetRegionId(v string) *CreateDBInstanceShrinkRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *CreateDBInstanceShrinkRequest
	GetResourceGroupId() *string
	SetSecondaryZone(v string) *CreateDBInstanceShrinkRequest
	GetSecondaryZone() *string
	SetSeries(v string) *CreateDBInstanceShrinkRequest
	GetSeries() *string
	SetStorageType(v string) *CreateDBInstanceShrinkRequest
	GetStorageType() *string
	SetTertiaryZone(v string) *CreateDBInstanceShrinkRequest
	GetTertiaryZone() *string
	SetTopologyType(v string) *CreateDBInstanceShrinkRequest
	GetTopologyType() *string
	SetUsedTime(v int32) *CreateDBInstanceShrinkRequest
	GetUsedTime() *int32
	SetVPCId(v string) *CreateDBInstanceShrinkRequest
	GetVPCId() *string
	SetVSwitchId(v string) *CreateDBInstanceShrinkRequest
	GetVSwitchId() *string
	SetZoneId(v string) *CreateDBInstanceShrinkRequest
	GetZoneId() *string
}

type CreateDBInstanceShrinkRequest struct {
	// Specifies whether to enable auto-renewal. Default value: true.
	//
	// - **true**: Enable
	//
	// - **false**: Disable
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
	// The idempotency token. Ensure that a unique value is used for each creation request.
	//
	// example:
	//
	// xxxxxx-xxx
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The compute node specification. Required for creating Enterprise Edition instances. Not required for Standard Edition.
	//
	// Enterprise Edition local disk:
	//
	// - **polarx.x4.medium.2e**:	2 Cores 8 GB (General Purpose)
	//
	// - **polarx.x4.large.2e**:	4 Cores 16 GB (General Purpose)
	//
	// - **polarx.x4.xlarge.2e**:	8 Cores 32 GB (General Purpose)
	//
	// - **polarx.x4.2xlarge.2e**:	16 Cores 64 GB (General Purpose)
	//
	// - **polarx.x8.large.2e**:	4 Cores 32 GB (Dedicated)
	//
	// - **polarx.x2.large.2x**:	8 Cores 16 GB (Dedicated)
	//
	// - **polarx.x4.xlarge.2x**:	8 Cores 32 GB (Dedicated)
	//
	// - **polarx.x8.xlarge.2e**:	8 Cores 64 GB (Dedicated)
	//
	// - **polarx.x8.2xlarge.2e**:	16 Cores 128 GB (Dedicated)
	//
	// - **polarx.x4.4xlarge.2e**:	32 Cores 128 GB (Dedicated)
	//
	// - **polarx.x8.4xlarge.2e**:	32 Cores 256 GB (Dedicated)
	//
	// - **polarx.st.8xlarge.2e**:	60 Cores 470 GB (Dedicated)
	//
	// - **polarx.st.12xlarge.2e**:	90 Cores 720 GB (Dedicated)
	//
	//
	// Enterprise Edition cloud disk:
	//
	// - **polarx.x4.medium.c2e**:	2 Cores 8 GB (General Purpose)
	//
	// - **polarx.x4.large.c2e**:	4 Cores 16 GB (General Purpose)
	//
	// - **polarx.x4.xlarge.c2e**:	8 Cores 32 GB (General Purpose)
	//
	// - **polarx.x4.2xlarge.c2e**:	16 Cores 64 GB (General Purpose)
	//
	// - **polarx.x8.large.c2e**:	4 Cores 32 GB (Dedicated)
	//
	// - **polarx.x2.large.c2x**:	8 Cores 16 GB (Dedicated)
	//
	// - **polarx.x4.xlarge.c2x**:	8 Cores 32 GB (Dedicated)
	//
	// - **polarx.x8.xlarge.c2e**:	8 Cores 64 GB (Dedicated)
	//
	// - **polarx.x8.2xlarge.c2e**:	16 Cores 128 GB (Dedicated)
	//
	// - **polarx.x4.4xlarge.c2e**:	32 Cores 128 GB (Dedicated)
	//
	// - **polarx.x8.4xlarge.c2e**:	32 Cores 256 GB (Dedicated)
	//
	// - **polarx.st.8xlarge.c2e**:	60 Cores 470 GB (Dedicated)
	//
	// - **polarx.st.12xlarge.c2e**:	90 Cores 720 GB (Dedicated)
	//
	// example:
	//
	// polarx.x4.medium.2e
	CnClass *string `json:"CnClass,omitempty" xml:"CnClass,omitempty"`
	// Required for creating Standard Edition instances. Not required for Enterprise Edition.
	//
	// Standard Edition local disk specifications:
	//
	// - **mysql.n8.small.25**:	1 Core 8 GB (General Purpose)
	//
	// - **mysql.n2.medium.25**:	2 Cores 4 GB (General Purpose)
	//
	// - **mysql.n4.medium.25**:	2 Cores 8 GB (General Purpose)
	//
	// - **mysql.n8.medium.25**:	2 Cores 16 GB (General Purpose)
	//
	// - **mysql.n2.large.25**:	4 Cores 8 GB (General Purpose)
	//
	// - **mysql.n4.large.25**:	4 Cores 16 GB (General Purpose)
	//
	// - **mysql.n8.large.25**:	4 Cores 32 GB (General Purpose)
	//
	// - **mysql.n2.xlarge.25**:	8 Cores 16 GB (General Purpose)
	//
	// - **mysql.n4.xlarge.25**:	8 Cores 32 GB (General Purpose)
	//
	// - **mysql.n8.xlarge.25**:	8 Cores 64 GB (General Purpose)
	//
	// - **mysql.n2.2xlarge.25**:	16 Cores 32 GB (General Purpose)
	//
	// - **mysql.n4.2xlarge.25**:	16 Cores 64 GB (General Purpose)
	//
	// - **mysql.n8.2xlarge.25**:	16 Cores 128 GB (General Purpose)
	//
	// - **mysql.x2.medium.25**:	2 Cores 4 GB (Dedicated)
	//
	// - **mysql.x4.medium.25**:	2 Cores 8 GB (Dedicated)
	//
	// - **mysql.x8.medium.25**:	2 Cores 16 GB (Dedicated)
	//
	// - **mysql.x2.large.25**:	4 Cores 8 GB (Dedicated)
	//
	// - **mysql.x4.large.25**:	4 Cores 16 GB (Dedicated)
	//
	// - **mysql.x8.large.25**:	4 Cores 32 GB (Dedicated)
	//
	// - **mysql.x2.xlarge.25**:	8 Cores 16 GB (Dedicated)
	//
	// - **mysql.x4.xlarge.25**:	8 Cores 32 GB (Dedicated)
	//
	// - **mysql.x8.xlarge.25**:	8 Cores 64 GB (Dedicated)
	//
	// - **mysql.x2.2xlarge.25**:	16 Cores 32 GB (Dedicated)
	//
	// - **mysql.x4.2xlarge.25**:	16 Cores 64 GB (Dedicated)
	//
	// - **mysql.x8.2xlarge.25**:	16 Cores 128 GB (Dedicated)
	//
	// - **mysql.x4.4xlarge.25**:	32 Cores 128 GB (Dedicated)
	//
	// - **mysql.x8.4xlarge.25**:	32 Cores 256 GB (Dedicated)
	//
	// - **mysql.x4.8xlarge.25**:	64 Cores 256 GB (Dedicated)
	//
	// - **mysql.x8.8xlarge.25**:	64 Cores 512 GB (Dedicated)
	//
	// - **mysql.st.12xlarge.25**:	90 Cores 720 GB (Dedicated)
	//
	// Standard Edition cloud disk specifications:
	//
	// - **polarx.mysql.n2.medium.c25**:	2 Cores 4 GB (General Purpose)
	//
	// - **polarx.mysql.n4.medium.c25**:	2 Cores 8 GB (General Purpose)
	//
	// - **polarx.mysql.n8.medium.c25**:	2 Cores 16 GB (General Purpose)
	//
	// - **polarx.mysql.n2.large.c25**:	4 Cores 8 GB (General Purpose)
	//
	// - **polarx.mysql.n4.large.c25**:	4 Cores 16 GB (General Purpose)
	//
	// - **polarx.mysql.n8.large.c25**:	4 Cores 32 GB (General Purpose)
	//
	// - **polarx.mysql.n2.xlarge.c25**:	8 Cores 16 GB (General Purpose)
	//
	// - **polarx.mysql.n4.xlarge.c25**:	8 Cores 32 GB (General Purpose)
	//
	// - **polarx.mysql.n8.xlarge.c25**:	8 Cores 64 GB (General Purpose)
	//
	// - **polarx.mysql.x2.medium.c25**:	2 Cores 4 GB (Dedicated)
	//
	// - **polarx.mysql.x4.medium.c25**:	2 Cores 8 GB (Dedicated)
	//
	// - **polarx.mysql.x8.medium.c25**:	2 Cores 16 GB (Dedicated)
	//
	// - **polarx.mysql.x2.large.c25**:	4 Cores 8 GB (Dedicated)
	//
	// - **polarx.mysql.x4.large.c25**:	4 Cores 16 GB (Dedicated)
	//
	// - **polarx.mysql.x8.large.c25**:	4 Cores 32 GB (Dedicated)
	//
	// - **polarx.mysql.x2.xlarge.c25**:	8 Cores 16 GB (Dedicated)
	//
	// - **polarx.mysql.x4.xlarge.c25**:	8 Cores 32 GB (Dedicated)
	//
	// - **polarx.mysql.x8.xlarge.c25**:	8 Cores 64 GB (Dedicated)
	//
	// - **polarx.mysql.x2.2xlarge.c25**:	16 Cores 32 GB (Dedicated)
	//
	// - **polarx.mysql.x4.2xlarge.c25**:	16 Cores 64 GB (Dedicated)
	//
	// - **polarx.mysql.x8.2xlarge.c25**:	16 Cores 128 GB (Dedicated)
	//
	// - **polarx.mysql.x2.4xlarge.c25**:	32 Cores 64 GB (Dedicated)
	//
	// - **polarx.mysql.x4.4xlarge.c25**:	32 Cores 128 GB (Dedicated)
	//
	// - **polarx.mysql.x8.4xlarge.c25**:	32 Cores 256 GB (Dedicated)
	//
	// - **polarx.mysql.x2.8xlarge.c25**:	64 Cores 128 GB (Dedicated)
	//
	// - **polarx.mysql.x4.8xlarge.c25**:	64 Cores 256 GB (Dedicated)
	//
	// - **polarx.mysql.x8.8xlarge.c25**:	64 Cores 512 GB (Dedicated)
	//
	// example:
	//
	// polarx.x4.2xlarge.2d
	DBNodeClass *string `json:"DBNodeClass,omitempty" xml:"DBNodeClass,omitempty"`
	// The number of instance nodes:
	//
	// 	- Standard Edition: 1
	//
	// 	- Enterprise Edition: >=2
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
	// The storage node specification. Required for creating Enterprise Edition instances. Not required for Standard Edition.
	//
	// Enterprise Edition local disk:
	//
	// - **mysql.n2.medium.25**:	2 Cores 4 GB (General Purpose)
	//
	// - **mysql.n4.medium.25**:	2 Cores 8 GB (General Purpose)
	//
	// - **mysql.n2.large.25**:	4 Cores 8 GB (General Purpose)
	//
	// - **mysql.n4.large.25**:	4 Cores 16 GB (General Purpose)
	//
	// - **mysql.n4.xlarge.25**:	8 Cores 32 GB (General Purpose)
	//
	// - **mysql.n4.2xlarge.25**:	16 Cores 64 GB (General Purpose)
	//
	// - **mysql.x4.large.25**:	4 Cores 16 GB (Dedicated)
	//
	// - **mysql.x8.large.25**:	4 Cores 32 GB (Dedicated)
	//
	// - **mysql.x2.xlarge.25**:	8 Cores 16 GB (Dedicated)
	//
	// - **mysql.x8.xlarge.25**:	8 Cores 64 GB (Dedicated)
	//
	// - **mysql.x8.2xlarge.25**:	16 Cores 128 GB (Dedicated)
	//
	// - **mysql.x4.4xlarge.25**:	32 Cores 128 GB (Dedicated)
	//
	// - **mysql.x8.4xlarge.25**:	32 Cores 256 GB (Dedicated)
	//
	// - **mysql.st.8xlarge.25**:	60 Cores 470 GB (Dedicated)
	//
	// - **mysql.st.12xlarge.25**:	90 Cores 720 GB (Dedicated)
	//
	// - **mysql.x8.45xlarge.25**:	180 Cores 1440 GB (Dedicated)
	//
	// - **mysql.x8.60xlarge.25**:	240 Cores 1920 GB (Dedicated)
	//
	//
	// Enterprise Edition cloud disk:
	//
	// - **polarx.mysql.n2.medium.c25**:	2 Cores 4 GB (General Purpose)
	//
	// - **polarx.mysql.n4.medium.c25**:	2 Cores 8 GB (General Purpose)
	//
	// - **polarx.mysql.n2.large.c25**:	4 Cores 8 GB (General Purpose)
	//
	// - **polarx.mysql.n4.large.c25**:	4 Cores 16 GB (General Purpose)
	//
	// - **polarx.mysql.n4.xlarge.c25**:	8 Cores 32 GB (General Purpose)
	//
	// - **polarx.mysql.n4.2xlarge.c25**:	16 Cores 64 GB (General Purpose)
	//
	// - **polarx.mysql.x4.large.c25**:	4 Cores 16 GB (Dedicated)
	//
	// - **polarx.mysql.x8.large.c25**:	4 Cores 32 GB (Dedicated)
	//
	// - **polarx.mysql.x2.xlarge.c25**:	8 Cores 16 GB (Dedicated)
	//
	// - **polarx.mysql.x8.xlarge.c25**:	8 Cores 64 GB (Dedicated)
	//
	// - **polarx.mysql.x8.2xlarge.c25**:	16 Cores 128 GB (Dedicated)
	//
	// - **polarx.mysql.x4.4xlarge.c25**:	32 Cores 128 GB (Dedicated)
	//
	// - **polarx.mysql.x8.4xlarge.c25**:	32 Cores 256 GB (Dedicated)
	//
	// - **polarx.mysql.st.8xlarge.c25**:	60 Cores 470 GB (Dedicated)
	//
	// - **polarx.mysql.st.12xlarge.c25**: 90 Cores 720 GB (Dedicated)
	//
	// - **polarx.mysql.x8.45xlarge.c25**: 180 Cores 1440 GB (Dedicated)
	//
	// - **polarx.mysql.x8.60xlarge.c25**: 240 Cores 1920 GB (Dedicated)
	//
	// example:
	//
	// mysql.n4.medium.25
	DnClass *string `json:"DnClass,omitempty" xml:"DnClass,omitempty"`
	// The disk space size of the storage node.
	//
	// example:
	//
	// 1500
	DnStorageSpace *string `json:"DnStorageSpace,omitempty" xml:"DnStorageSpace,omitempty"`
	// The MySQL engine version. Valid values: 5.7 and 8.0.
	//
	// This parameter is required.
	//
	// example:
	//
	// 5.7
	EngineVersion *string `json:"EngineVersion,omitempty" xml:"EngineVersion,omitempty"`
	// Additional parameters.
	ExtraParamsShrink *string `json:"ExtraParams,omitempty" xml:"ExtraParams,omitempty"`
	// Specifies whether the instance to be created is a columnar read-only instance.
	//
	// example:
	//
	// true
	IsColumnarReadDBInstance *bool `json:"IsColumnarReadDBInstance,omitempty" xml:"IsColumnarReadDBInstance,omitempty"`
	// Specifies whether the instance is a read-only instance.
	//
	// - **true**: Yes
	//
	// - **false**: No
	//
	// example:
	//
	// false
	IsReadDBInstance *bool `json:"IsReadDBInstance,omitempty" xml:"IsReadDBInstance,omitempty"`
	// The network type. Only VPC is supported.
	//
	// example:
	//
	// vpc
	NetworkType *string `json:"NetworkType,omitempty" xml:"NetworkType,omitempty"`
	// The minor version number. This parameter is generally not required.
	//
	// example:
	//
	// polardb-2.4.0_standard_xcluster8.4.20-20260331
	OriginMinorVersion *string `json:"OriginMinorVersion,omitempty" xml:"OriginMinorVersion,omitempty"`
	// The billing method of the instance.
	//
	// - **PREPAY**: Subscription
	//
	// - **POSTPAY**: Pay-as-you-go
	//
	// This parameter is required.
	//
	// example:
	//
	// PREPAY
	PayType *string `json:"PayType,omitempty" xml:"PayType,omitempty"`
	// The billing cycle. For subscription instances, valid values are Year and Month. For pay-as-you-go instances, the default value is Hour.
	//
	// example:
	//
	// Month
	Period *string `json:"Period,omitempty" xml:"Period,omitempty"`
	// The primary instance must be specified if the instance is a read-only instance.
	//
	// example:
	//
	// pxc-*********
	PrimaryDBInstanceName *string `json:"PrimaryDBInstanceName,omitempty" xml:"PrimaryDBInstanceName,omitempty"`
	// The primary availability zone.
	//
	// example:
	//
	// cn-shenzhen-e
	PrimaryZone *string `json:"PrimaryZone,omitempty" xml:"PrimaryZone,omitempty"`
	// The region where the instance resides.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The resource group ID. This parameter can be left empty. This feature is currently not supported.
	//
	// example:
	//
	// null
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The secondary availability zone.
	//
	// example:
	//
	// cn-shenzhen-a
	SecondaryZone *string `json:"SecondaryZone,omitempty" xml:"SecondaryZone,omitempty"`
	// Enterprise Edition: enterprise
	//
	// Standard Edition: standard
	//
	// example:
	//
	// enterprise
	Series *string `json:"Series,omitempty" xml:"Series,omitempty"`
	// The storage type.
	//
	// 	- Local disk: custom_local_ssd
	//
	// 	- Cloud disk: cloud_auto
	//
	// example:
	//
	// cloud_auto
	StorageType *string `json:"StorageType,omitempty" xml:"StorageType,omitempty"`
	// The tertiary availability zone.
	//
	// example:
	//
	// cn-shenzhen-e
	TertiaryZone *string `json:"TertiaryZone,omitempty" xml:"TertiaryZone,omitempty"`
	// The topology type:
	//
	// - **3azones**: Three availability zones
	//
	// - **1azone**: Single availability zone
	//
	// This parameter is required.
	//
	// example:
	//
	// 3azones
	TopologyType *string `json:"TopologyType,omitempty" xml:"TopologyType,omitempty"`
	// The subscription duration. You can specify the number of months or years for prepaid instances.
	//
	// > When Period is set to Year, valid values for this parameter are 1, 2, and 3.
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
	// The availability zone of the instance.
	//
	// example:
	//
	// cn-hangzhou-a
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s CreateDBInstanceShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateDBInstanceShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateDBInstanceShrinkRequest) GetAutoRenew() *bool {
	return s.AutoRenew
}

func (s *CreateDBInstanceShrinkRequest) GetCNNodeCount() *int32 {
	return s.CNNodeCount
}

func (s *CreateDBInstanceShrinkRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateDBInstanceShrinkRequest) GetCnClass() *string {
	return s.CnClass
}

func (s *CreateDBInstanceShrinkRequest) GetDBNodeClass() *string {
	return s.DBNodeClass
}

func (s *CreateDBInstanceShrinkRequest) GetDBNodeCount() *int32 {
	return s.DBNodeCount
}

func (s *CreateDBInstanceShrinkRequest) GetDNNodeCount() *int32 {
	return s.DNNodeCount
}

func (s *CreateDBInstanceShrinkRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateDBInstanceShrinkRequest) GetDnClass() *string {
	return s.DnClass
}

func (s *CreateDBInstanceShrinkRequest) GetDnStorageSpace() *string {
	return s.DnStorageSpace
}

func (s *CreateDBInstanceShrinkRequest) GetEngineVersion() *string {
	return s.EngineVersion
}

func (s *CreateDBInstanceShrinkRequest) GetExtraParamsShrink() *string {
	return s.ExtraParamsShrink
}

func (s *CreateDBInstanceShrinkRequest) GetIsColumnarReadDBInstance() *bool {
	return s.IsColumnarReadDBInstance
}

func (s *CreateDBInstanceShrinkRequest) GetIsReadDBInstance() *bool {
	return s.IsReadDBInstance
}

func (s *CreateDBInstanceShrinkRequest) GetNetworkType() *string {
	return s.NetworkType
}

func (s *CreateDBInstanceShrinkRequest) GetOriginMinorVersion() *string {
	return s.OriginMinorVersion
}

func (s *CreateDBInstanceShrinkRequest) GetPayType() *string {
	return s.PayType
}

func (s *CreateDBInstanceShrinkRequest) GetPeriod() *string {
	return s.Period
}

func (s *CreateDBInstanceShrinkRequest) GetPrimaryDBInstanceName() *string {
	return s.PrimaryDBInstanceName
}

func (s *CreateDBInstanceShrinkRequest) GetPrimaryZone() *string {
	return s.PrimaryZone
}

func (s *CreateDBInstanceShrinkRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateDBInstanceShrinkRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *CreateDBInstanceShrinkRequest) GetSecondaryZone() *string {
	return s.SecondaryZone
}

func (s *CreateDBInstanceShrinkRequest) GetSeries() *string {
	return s.Series
}

func (s *CreateDBInstanceShrinkRequest) GetStorageType() *string {
	return s.StorageType
}

func (s *CreateDBInstanceShrinkRequest) GetTertiaryZone() *string {
	return s.TertiaryZone
}

func (s *CreateDBInstanceShrinkRequest) GetTopologyType() *string {
	return s.TopologyType
}

func (s *CreateDBInstanceShrinkRequest) GetUsedTime() *int32 {
	return s.UsedTime
}

func (s *CreateDBInstanceShrinkRequest) GetVPCId() *string {
	return s.VPCId
}

func (s *CreateDBInstanceShrinkRequest) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *CreateDBInstanceShrinkRequest) GetZoneId() *string {
	return s.ZoneId
}

func (s *CreateDBInstanceShrinkRequest) SetAutoRenew(v bool) *CreateDBInstanceShrinkRequest {
	s.AutoRenew = &v
	return s
}

func (s *CreateDBInstanceShrinkRequest) SetCNNodeCount(v int32) *CreateDBInstanceShrinkRequest {
	s.CNNodeCount = &v
	return s
}

func (s *CreateDBInstanceShrinkRequest) SetClientToken(v string) *CreateDBInstanceShrinkRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateDBInstanceShrinkRequest) SetCnClass(v string) *CreateDBInstanceShrinkRequest {
	s.CnClass = &v
	return s
}

func (s *CreateDBInstanceShrinkRequest) SetDBNodeClass(v string) *CreateDBInstanceShrinkRequest {
	s.DBNodeClass = &v
	return s
}

func (s *CreateDBInstanceShrinkRequest) SetDBNodeCount(v int32) *CreateDBInstanceShrinkRequest {
	s.DBNodeCount = &v
	return s
}

func (s *CreateDBInstanceShrinkRequest) SetDNNodeCount(v int32) *CreateDBInstanceShrinkRequest {
	s.DNNodeCount = &v
	return s
}

func (s *CreateDBInstanceShrinkRequest) SetDescription(v string) *CreateDBInstanceShrinkRequest {
	s.Description = &v
	return s
}

func (s *CreateDBInstanceShrinkRequest) SetDnClass(v string) *CreateDBInstanceShrinkRequest {
	s.DnClass = &v
	return s
}

func (s *CreateDBInstanceShrinkRequest) SetDnStorageSpace(v string) *CreateDBInstanceShrinkRequest {
	s.DnStorageSpace = &v
	return s
}

func (s *CreateDBInstanceShrinkRequest) SetEngineVersion(v string) *CreateDBInstanceShrinkRequest {
	s.EngineVersion = &v
	return s
}

func (s *CreateDBInstanceShrinkRequest) SetExtraParamsShrink(v string) *CreateDBInstanceShrinkRequest {
	s.ExtraParamsShrink = &v
	return s
}

func (s *CreateDBInstanceShrinkRequest) SetIsColumnarReadDBInstance(v bool) *CreateDBInstanceShrinkRequest {
	s.IsColumnarReadDBInstance = &v
	return s
}

func (s *CreateDBInstanceShrinkRequest) SetIsReadDBInstance(v bool) *CreateDBInstanceShrinkRequest {
	s.IsReadDBInstance = &v
	return s
}

func (s *CreateDBInstanceShrinkRequest) SetNetworkType(v string) *CreateDBInstanceShrinkRequest {
	s.NetworkType = &v
	return s
}

func (s *CreateDBInstanceShrinkRequest) SetOriginMinorVersion(v string) *CreateDBInstanceShrinkRequest {
	s.OriginMinorVersion = &v
	return s
}

func (s *CreateDBInstanceShrinkRequest) SetPayType(v string) *CreateDBInstanceShrinkRequest {
	s.PayType = &v
	return s
}

func (s *CreateDBInstanceShrinkRequest) SetPeriod(v string) *CreateDBInstanceShrinkRequest {
	s.Period = &v
	return s
}

func (s *CreateDBInstanceShrinkRequest) SetPrimaryDBInstanceName(v string) *CreateDBInstanceShrinkRequest {
	s.PrimaryDBInstanceName = &v
	return s
}

func (s *CreateDBInstanceShrinkRequest) SetPrimaryZone(v string) *CreateDBInstanceShrinkRequest {
	s.PrimaryZone = &v
	return s
}

func (s *CreateDBInstanceShrinkRequest) SetRegionId(v string) *CreateDBInstanceShrinkRequest {
	s.RegionId = &v
	return s
}

func (s *CreateDBInstanceShrinkRequest) SetResourceGroupId(v string) *CreateDBInstanceShrinkRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *CreateDBInstanceShrinkRequest) SetSecondaryZone(v string) *CreateDBInstanceShrinkRequest {
	s.SecondaryZone = &v
	return s
}

func (s *CreateDBInstanceShrinkRequest) SetSeries(v string) *CreateDBInstanceShrinkRequest {
	s.Series = &v
	return s
}

func (s *CreateDBInstanceShrinkRequest) SetStorageType(v string) *CreateDBInstanceShrinkRequest {
	s.StorageType = &v
	return s
}

func (s *CreateDBInstanceShrinkRequest) SetTertiaryZone(v string) *CreateDBInstanceShrinkRequest {
	s.TertiaryZone = &v
	return s
}

func (s *CreateDBInstanceShrinkRequest) SetTopologyType(v string) *CreateDBInstanceShrinkRequest {
	s.TopologyType = &v
	return s
}

func (s *CreateDBInstanceShrinkRequest) SetUsedTime(v int32) *CreateDBInstanceShrinkRequest {
	s.UsedTime = &v
	return s
}

func (s *CreateDBInstanceShrinkRequest) SetVPCId(v string) *CreateDBInstanceShrinkRequest {
	s.VPCId = &v
	return s
}

func (s *CreateDBInstanceShrinkRequest) SetVSwitchId(v string) *CreateDBInstanceShrinkRequest {
	s.VSwitchId = &v
	return s
}

func (s *CreateDBInstanceShrinkRequest) SetZoneId(v string) *CreateDBInstanceShrinkRequest {
	s.ZoneId = &v
	return s
}

func (s *CreateDBInstanceShrinkRequest) Validate() error {
	return dara.Validate(s)
}
