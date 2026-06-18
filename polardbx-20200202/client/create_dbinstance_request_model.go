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
	ExtraParams map[string]*string `json:"ExtraParams,omitempty" xml:"ExtraParams,omitempty"`
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
