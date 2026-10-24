// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateInstanceV1Request interface {
	dara.Model
	String() string
	GoString() string
	SetAdminPassword(v string) *CreateInstanceV1Request
	GetAdminPassword() *string
	SetAgentNodeGroup(v *CreateInstanceV1RequestAgentNodeGroup) *CreateInstanceV1Request
	GetAgentNodeGroup() *CreateInstanceV1RequestAgentNodeGroup
	SetAutoPay(v bool) *CreateInstanceV1Request
	GetAutoPay() *bool
	SetAutoRenew(v bool) *CreateInstanceV1Request
	GetAutoRenew() *bool
	SetBackendNodeGroups(v []*CreateInstanceV1RequestBackendNodeGroups) *CreateInstanceV1Request
	GetBackendNodeGroups() []*CreateInstanceV1RequestBackendNodeGroups
	SetClientToken(v string) *CreateInstanceV1Request
	GetClientToken() *string
	SetDlfCatalogName(v string) *CreateInstanceV1Request
	GetDlfCatalogName() *string
	SetDlfCatalogType(v string) *CreateInstanceV1Request
	GetDlfCatalogType() *string
	SetDuration(v int32) *CreateInstanceV1Request
	GetDuration() *int32
	SetEnableAiFunction(v bool) *CreateInstanceV1Request
	GetEnableAiFunction() *bool
	SetEnableMultiAz(v bool) *CreateInstanceV1Request
	GetEnableMultiAz() *bool
	SetEncrypted(v bool) *CreateInstanceV1Request
	GetEncrypted() *bool
	SetFrontendNodeGroups(v []*CreateInstanceV1RequestFrontendNodeGroups) *CreateInstanceV1Request
	GetFrontendNodeGroups() []*CreateInstanceV1RequestFrontendNodeGroups
	SetGatewayType(v string) *CreateInstanceV1Request
	GetGatewayType() *string
	SetInstanceName(v string) *CreateInstanceV1Request
	GetInstanceName() *string
	SetKmsKeyId(v string) *CreateInstanceV1Request
	GetKmsKeyId() *string
	SetLinkedRamUserName(v string) *CreateInstanceV1Request
	GetLinkedRamUserName() *string
	SetObserverNodeGroups(v []*CreateInstanceV1RequestObserverNodeGroups) *CreateInstanceV1Request
	GetObserverNodeGroups() []*CreateInstanceV1RequestObserverNodeGroups
	SetOssAccessingRoleName(v string) *CreateInstanceV1Request
	GetOssAccessingRoleName() *string
	SetPackageType(v string) *CreateInstanceV1Request
	GetPackageType() *string
	SetPayType(v string) *CreateInstanceV1Request
	GetPayType() *string
	SetPricingCycle(v string) *CreateInstanceV1Request
	GetPricingCycle() *string
	SetPrincipalType(v string) *CreateInstanceV1Request
	GetPrincipalType() *string
	SetPromotionOptionNo(v string) *CreateInstanceV1Request
	GetPromotionOptionNo() *string
	SetRamUserId(v string) *CreateInstanceV1Request
	GetRamUserId() *string
	SetRegionId(v string) *CreateInstanceV1Request
	GetRegionId() *string
	SetResourceGroupId(v string) *CreateInstanceV1Request
	GetResourceGroupId() *string
	SetRunMode(v string) *CreateInstanceV1Request
	GetRunMode() *string
	SetTags(v []*CreateInstanceV1RequestTags) *CreateInstanceV1Request
	GetTags() []*CreateInstanceV1RequestTags
	SetVSwitches(v []*CreateInstanceV1RequestVSwitches) *CreateInstanceV1Request
	GetVSwitches() []*CreateInstanceV1RequestVSwitches
	SetVersion(v string) *CreateInstanceV1Request
	GetVersion() *string
	SetVpcId(v string) *CreateInstanceV1Request
	GetVpcId() *string
	SetZoneId(v string) *CreateInstanceV1Request
	GetZoneId() *string
}

type CreateInstanceV1Request struct {
	// The password of the instance administrator.
	//
	// This parameter is required.
	//
	// example:
	//
	// password_example
	AdminPassword *string `json:"AdminPassword,omitempty" xml:"AdminPassword,omitempty"`
	// The agent compute group.
	AgentNodeGroup *CreateInstanceV1RequestAgentNodeGroup `json:"AgentNodeGroup,omitempty" xml:"AgentNodeGroup,omitempty" type:"Struct"`
	// Specifies whether automatic payment is enabled for all products specified in the Products parameter. Valid values:
	//
	// - true: Enables automatic payment.
	//
	// - false: Disables automatic payment.
	//
	// example:
	//
	// false
	AutoPay *bool `json:"AutoPay,omitempty" xml:"AutoPay,omitempty"`
	// Specifies whether to enable auto-renewal. This parameter takes effect only when PayType is set to PrePaid. Default value: false.
	//
	// example:
	//
	// true
	AutoRenew *bool `json:"AutoRenew,omitempty" xml:"AutoRenew,omitempty"`
	// The information about the BE or CN compute groups.
	BackendNodeGroups []*CreateInstanceV1RequestBackendNodeGroups `json:"BackendNodeGroups,omitempty" xml:"BackendNodeGroups,omitempty" type:"Repeated"`
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the value, but you must make sure that the value is unique among different requests. The token can contain only ASCII characters and cannot exceed 64 characters in length.
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-426655440000
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The name of the DLF catalog.
	//
	// example:
	//
	// paimon_test
	DlfCatalogName *string `json:"DlfCatalogName,omitempty" xml:"DlfCatalogName,omitempty"`
	// The type of the DLF catalog, such as paimon or iceberg.
	//
	// example:
	//
	// paimon
	DlfCatalogType *string `json:"DlfCatalogType,omitempty" xml:"DlfCatalogType,omitempty"`
	// The subscription duration. This parameter takes effect only when PayType is set to PrePaid.
	//
	// example:
	//
	// 1
	Duration *int32 `json:"Duration,omitempty" xml:"Duration,omitempty"`
	// Specifies whether to enable the AI center. Default value: false.
	EnableAiFunction *bool `json:"EnableAiFunction,omitempty" xml:"EnableAiFunction,omitempty"`
	// Specifies whether to enable disaster recovery.
	EnableMultiAz *bool `json:"EnableMultiAz,omitempty" xml:"EnableMultiAz,omitempty"`
	// Specifies whether to enable encryption.
	//
	// example:
	//
	// true
	Encrypted *bool `json:"Encrypted,omitempty" xml:"Encrypted,omitempty"`
	// The information about the FE node groups.
	FrontendNodeGroups []*CreateInstanceV1RequestFrontendNodeGroups `json:"FrontendNodeGroups,omitempty" xml:"FrontendNodeGroups,omitempty" type:"Repeated"`
	// The gateway type. Generally, the value is Slb or Privatezone.
	//
	// example:
	//
	// Traditional\\"\\"
	GatewayType *string `json:"GatewayType,omitempty" xml:"GatewayType,omitempty"`
	// The instance name.
	//
	// This parameter is required.
	//
	// example:
	//
	// c1
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// The KMS key ID.
	//
	// example:
	//
	// fdsdf****
	KmsKeyId *string `json:"KmsKeyId,omitempty" xml:"KmsKeyId,omitempty"`
	// The username of the associated Resource Access Management (RAM) user or the name of the RAM role.
	//
	// example:
	//
	// aliyun.test1
	LinkedRamUserName *string `json:"LinkedRamUserName,omitempty" xml:"LinkedRamUserName,omitempty"`
	// The information about the Observer compute group. Specify this parameter only when you need to enable multi-zone disaster recovery. The configuration of the Observer compute group must be consistent with the specifications of the FE node group.
	ObserverNodeGroups []*CreateInstanceV1RequestObserverNodeGroups `json:"ObserverNodeGroups,omitempty" xml:"ObserverNodeGroups,omitempty" type:"Repeated"`
	// The name of the role used for password-free access to OSS.
	//
	// example:
	//
	// AliyunEMRStarRocksAccessingOSSRole
	OssAccessingRoleName *string `json:"OssAccessingRoleName,omitempty" xml:"OssAccessingRoleName,omitempty"`
	// The instance edition. Valid values:
	//
	// <ul>
	//
	// <li>trial: Trial Edition.</li>
	//
	// <li>official: Standard Edition.</li>
	//
	// </ul>
	//
	// This parameter is required.
	//
	// example:
	//
	// official
	PackageType *string `json:"PackageType,omitempty" xml:"PackageType,omitempty"`
	// The billing method of the instance. Valid values:
	//
	// <ol>
	//
	// <li>prePaid: subscription.</li>
	//
	// <li>postPaid: pay-as-you-go.</li>
	//
	// </ol>
	//
	// This parameter is required.
	//
	// example:
	//
	// prePaid
	PayType *string `json:"PayType,omitempty" xml:"PayType,omitempty"`
	// The unit of the billing cycle. Valid values:
	//
	// <ul>
	//
	// <li>Month</li>
	//
	// <li>Year</li>
	//
	// </ul>
	//
	// This parameter takes effect only when PayType is set to PrePaid.
	//
	// example:
	//
	// Month
	PricingCycle *string `json:"PricingCycle,omitempty" xml:"PricingCycle,omitempty"`
	// The RAM authentication type. A value of RS indicates a Resource Access Management (RAM) user, and a value of RR indicates a RAM role.
	//
	// example:
	//
	// RR
	PrincipalType *string `json:"PrincipalType,omitempty" xml:"PrincipalType,omitempty"`
	// The coupon ID.
	//
	// example:
	//
	// 165445235634
	PromotionOptionNo *string `json:"PromotionOptionNo,omitempty" xml:"PromotionOptionNo,omitempty"`
	// The ID of the Resource Access Management (RAM) user or RAM role.
	//
	// example:
	//
	// 123456
	RamUserId *string `json:"RamUserId,omitempty" xml:"RamUserId,omitempty"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The resource group ID.
	//
	// example:
	//
	// rg-aekzllkih7jqxxx
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The run mode of the cluster. Valid values:
	//
	// <ul>
	//
	// <li>shared_nothing: shared-nothing architecture.</li>
	//
	// <li>shared_data: shared-data architecture.</li>
	//
	// </ul>
	//
	// This parameter is required.
	//
	// example:
	//
	// shared_data
	RunMode *string `json:"RunMode,omitempty" xml:"RunMode,omitempty"`
	// The tags of the instance.
	Tags []*CreateInstanceV1RequestTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	// The vSwitch and zone information.
	VSwitches []*CreateInstanceV1RequestVSwitches `json:"VSwitches,omitempty" xml:"VSwitches,omitempty" type:"Repeated"`
	// The major version of the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// 3
	Version *string `json:"Version,omitempty" xml:"Version,omitempty"`
	// The VPC ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// vpc-bp1fll2mci6d7pw8m****
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	// The primary zone ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou-i
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s CreateInstanceV1Request) String() string {
	return dara.Prettify(s)
}

func (s CreateInstanceV1Request) GoString() string {
	return s.String()
}

func (s *CreateInstanceV1Request) GetAdminPassword() *string {
	return s.AdminPassword
}

func (s *CreateInstanceV1Request) GetAgentNodeGroup() *CreateInstanceV1RequestAgentNodeGroup {
	return s.AgentNodeGroup
}

func (s *CreateInstanceV1Request) GetAutoPay() *bool {
	return s.AutoPay
}

func (s *CreateInstanceV1Request) GetAutoRenew() *bool {
	return s.AutoRenew
}

func (s *CreateInstanceV1Request) GetBackendNodeGroups() []*CreateInstanceV1RequestBackendNodeGroups {
	return s.BackendNodeGroups
}

func (s *CreateInstanceV1Request) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateInstanceV1Request) GetDlfCatalogName() *string {
	return s.DlfCatalogName
}

func (s *CreateInstanceV1Request) GetDlfCatalogType() *string {
	return s.DlfCatalogType
}

func (s *CreateInstanceV1Request) GetDuration() *int32 {
	return s.Duration
}

func (s *CreateInstanceV1Request) GetEnableAiFunction() *bool {
	return s.EnableAiFunction
}

func (s *CreateInstanceV1Request) GetEnableMultiAz() *bool {
	return s.EnableMultiAz
}

func (s *CreateInstanceV1Request) GetEncrypted() *bool {
	return s.Encrypted
}

func (s *CreateInstanceV1Request) GetFrontendNodeGroups() []*CreateInstanceV1RequestFrontendNodeGroups {
	return s.FrontendNodeGroups
}

func (s *CreateInstanceV1Request) GetGatewayType() *string {
	return s.GatewayType
}

func (s *CreateInstanceV1Request) GetInstanceName() *string {
	return s.InstanceName
}

func (s *CreateInstanceV1Request) GetKmsKeyId() *string {
	return s.KmsKeyId
}

func (s *CreateInstanceV1Request) GetLinkedRamUserName() *string {
	return s.LinkedRamUserName
}

func (s *CreateInstanceV1Request) GetObserverNodeGroups() []*CreateInstanceV1RequestObserverNodeGroups {
	return s.ObserverNodeGroups
}

func (s *CreateInstanceV1Request) GetOssAccessingRoleName() *string {
	return s.OssAccessingRoleName
}

func (s *CreateInstanceV1Request) GetPackageType() *string {
	return s.PackageType
}

func (s *CreateInstanceV1Request) GetPayType() *string {
	return s.PayType
}

func (s *CreateInstanceV1Request) GetPricingCycle() *string {
	return s.PricingCycle
}

func (s *CreateInstanceV1Request) GetPrincipalType() *string {
	return s.PrincipalType
}

func (s *CreateInstanceV1Request) GetPromotionOptionNo() *string {
	return s.PromotionOptionNo
}

func (s *CreateInstanceV1Request) GetRamUserId() *string {
	return s.RamUserId
}

func (s *CreateInstanceV1Request) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateInstanceV1Request) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *CreateInstanceV1Request) GetRunMode() *string {
	return s.RunMode
}

func (s *CreateInstanceV1Request) GetTags() []*CreateInstanceV1RequestTags {
	return s.Tags
}

func (s *CreateInstanceV1Request) GetVSwitches() []*CreateInstanceV1RequestVSwitches {
	return s.VSwitches
}

func (s *CreateInstanceV1Request) GetVersion() *string {
	return s.Version
}

func (s *CreateInstanceV1Request) GetVpcId() *string {
	return s.VpcId
}

func (s *CreateInstanceV1Request) GetZoneId() *string {
	return s.ZoneId
}

func (s *CreateInstanceV1Request) SetAdminPassword(v string) *CreateInstanceV1Request {
	s.AdminPassword = &v
	return s
}

func (s *CreateInstanceV1Request) SetAgentNodeGroup(v *CreateInstanceV1RequestAgentNodeGroup) *CreateInstanceV1Request {
	s.AgentNodeGroup = v
	return s
}

func (s *CreateInstanceV1Request) SetAutoPay(v bool) *CreateInstanceV1Request {
	s.AutoPay = &v
	return s
}

func (s *CreateInstanceV1Request) SetAutoRenew(v bool) *CreateInstanceV1Request {
	s.AutoRenew = &v
	return s
}

func (s *CreateInstanceV1Request) SetBackendNodeGroups(v []*CreateInstanceV1RequestBackendNodeGroups) *CreateInstanceV1Request {
	s.BackendNodeGroups = v
	return s
}

func (s *CreateInstanceV1Request) SetClientToken(v string) *CreateInstanceV1Request {
	s.ClientToken = &v
	return s
}

func (s *CreateInstanceV1Request) SetDlfCatalogName(v string) *CreateInstanceV1Request {
	s.DlfCatalogName = &v
	return s
}

func (s *CreateInstanceV1Request) SetDlfCatalogType(v string) *CreateInstanceV1Request {
	s.DlfCatalogType = &v
	return s
}

func (s *CreateInstanceV1Request) SetDuration(v int32) *CreateInstanceV1Request {
	s.Duration = &v
	return s
}

func (s *CreateInstanceV1Request) SetEnableAiFunction(v bool) *CreateInstanceV1Request {
	s.EnableAiFunction = &v
	return s
}

func (s *CreateInstanceV1Request) SetEnableMultiAz(v bool) *CreateInstanceV1Request {
	s.EnableMultiAz = &v
	return s
}

func (s *CreateInstanceV1Request) SetEncrypted(v bool) *CreateInstanceV1Request {
	s.Encrypted = &v
	return s
}

func (s *CreateInstanceV1Request) SetFrontendNodeGroups(v []*CreateInstanceV1RequestFrontendNodeGroups) *CreateInstanceV1Request {
	s.FrontendNodeGroups = v
	return s
}

func (s *CreateInstanceV1Request) SetGatewayType(v string) *CreateInstanceV1Request {
	s.GatewayType = &v
	return s
}

func (s *CreateInstanceV1Request) SetInstanceName(v string) *CreateInstanceV1Request {
	s.InstanceName = &v
	return s
}

func (s *CreateInstanceV1Request) SetKmsKeyId(v string) *CreateInstanceV1Request {
	s.KmsKeyId = &v
	return s
}

func (s *CreateInstanceV1Request) SetLinkedRamUserName(v string) *CreateInstanceV1Request {
	s.LinkedRamUserName = &v
	return s
}

func (s *CreateInstanceV1Request) SetObserverNodeGroups(v []*CreateInstanceV1RequestObserverNodeGroups) *CreateInstanceV1Request {
	s.ObserverNodeGroups = v
	return s
}

func (s *CreateInstanceV1Request) SetOssAccessingRoleName(v string) *CreateInstanceV1Request {
	s.OssAccessingRoleName = &v
	return s
}

func (s *CreateInstanceV1Request) SetPackageType(v string) *CreateInstanceV1Request {
	s.PackageType = &v
	return s
}

func (s *CreateInstanceV1Request) SetPayType(v string) *CreateInstanceV1Request {
	s.PayType = &v
	return s
}

func (s *CreateInstanceV1Request) SetPricingCycle(v string) *CreateInstanceV1Request {
	s.PricingCycle = &v
	return s
}

func (s *CreateInstanceV1Request) SetPrincipalType(v string) *CreateInstanceV1Request {
	s.PrincipalType = &v
	return s
}

func (s *CreateInstanceV1Request) SetPromotionOptionNo(v string) *CreateInstanceV1Request {
	s.PromotionOptionNo = &v
	return s
}

func (s *CreateInstanceV1Request) SetRamUserId(v string) *CreateInstanceV1Request {
	s.RamUserId = &v
	return s
}

func (s *CreateInstanceV1Request) SetRegionId(v string) *CreateInstanceV1Request {
	s.RegionId = &v
	return s
}

func (s *CreateInstanceV1Request) SetResourceGroupId(v string) *CreateInstanceV1Request {
	s.ResourceGroupId = &v
	return s
}

func (s *CreateInstanceV1Request) SetRunMode(v string) *CreateInstanceV1Request {
	s.RunMode = &v
	return s
}

func (s *CreateInstanceV1Request) SetTags(v []*CreateInstanceV1RequestTags) *CreateInstanceV1Request {
	s.Tags = v
	return s
}

func (s *CreateInstanceV1Request) SetVSwitches(v []*CreateInstanceV1RequestVSwitches) *CreateInstanceV1Request {
	s.VSwitches = v
	return s
}

func (s *CreateInstanceV1Request) SetVersion(v string) *CreateInstanceV1Request {
	s.Version = &v
	return s
}

func (s *CreateInstanceV1Request) SetVpcId(v string) *CreateInstanceV1Request {
	s.VpcId = &v
	return s
}

func (s *CreateInstanceV1Request) SetZoneId(v string) *CreateInstanceV1Request {
	s.ZoneId = &v
	return s
}

func (s *CreateInstanceV1Request) Validate() error {
	if s.AgentNodeGroup != nil {
		if err := s.AgentNodeGroup.Validate(); err != nil {
			return err
		}
	}
	if s.BackendNodeGroups != nil {
		for _, item := range s.BackendNodeGroups {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.FrontendNodeGroups != nil {
		for _, item := range s.FrontendNodeGroups {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.ObserverNodeGroups != nil {
		for _, item := range s.ObserverNodeGroups {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Tags != nil {
		for _, item := range s.Tags {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.VSwitches != nil {
		for _, item := range s.VSwitches {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type CreateInstanceV1RequestAgentNodeGroup struct {
	// The number of compute units (CUs). A CU is the basic metering unit of the service. 1 CU equals 1 vCPU and 4 GiB of memory.
	//
	// example:
	//
	// 2
	Cu *int32 `json:"cu,omitempty" xml:"cu,omitempty"`
}

func (s CreateInstanceV1RequestAgentNodeGroup) String() string {
	return dara.Prettify(s)
}

func (s CreateInstanceV1RequestAgentNodeGroup) GoString() string {
	return s.String()
}

func (s *CreateInstanceV1RequestAgentNodeGroup) GetCu() *int32 {
	return s.Cu
}

func (s *CreateInstanceV1RequestAgentNodeGroup) SetCu(v int32) *CreateInstanceV1RequestAgentNodeGroup {
	s.Cu = &v
	return s
}

func (s *CreateInstanceV1RequestAgentNodeGroup) Validate() error {
	return dara.Validate(s)
}

type CreateInstanceV1RequestBackendNodeGroups struct {
	// The number of CUs. A CU (Compute Unit) is the basic metering unit of the service. 1 CU equals 1 vCPU and 4 GiB of memory. If SpecType is set to the memory-enhanced instance family, 1 CU equals 1 vCPU and 8 GiB of memory.
	//
	// example:
	//
	// 8
	Cu *int32 `json:"cu,omitempty" xml:"cu,omitempty"`
	// The number of disks.
	//
	// if can be null:
	// false
	//
	// example:
	//
	// 1
	DiskNumber *int32 `json:"diskNumber,omitempty" xml:"diskNumber,omitempty"`
	// The local SSD instance type of the node group. This parameter is valid only when the instance is based on ECS and SpecType is set to localSSD or bigData.
	//
	// example:
	//
	// local_ssd_4_4xlarge
	LocalStorageInstanceType *string `json:"localStorageInstanceType,omitempty" xml:"localStorageInstanceType,omitempty"`
	// The number of nodes.
	//
	// example:
	//
	// 3
	ResidentNodeNumber *int32 `json:"residentNodeNumber,omitempty" xml:"residentNodeNumber,omitempty"`
	// The instance family category of the compute group. Valid values:
	//
	// - standard: Standard.
	//
	// - localSSD: Local SSD.
	//
	// - bigData: Large storage.
	//
	// - ramEnhanced: Memory-enhanced instance family.
	//
	// - networkEnhanced: Network-enhanced.
	//
	// example:
	//
	// standard
	SpecType *string `json:"specType,omitempty" xml:"specType,omitempty"`
	// The performance level (PL) of the cloud disk. Valid values:
	//
	// - pl0: A single disk delivers up to 10,000 random read/write IOPS.
	//
	// - pl1: A single disk delivers up to 50,000 random read/write IOPS.
	//
	// - pl2: A single disk delivers up to 100,000 random read/write IOPS.
	//
	// - pl3: A single disk delivers up to 1,000,000 random read/write IOPS.
	//
	// example:
	//
	// pl1
	StoragePerformanceLevel *string `json:"storagePerformanceLevel,omitempty" xml:"storagePerformanceLevel,omitempty"`
	// The storage size. Unit: GiB.
	//
	// example:
	//
	// 100
	StorageSize *int32 `json:"storageSize,omitempty" xml:"storageSize,omitempty"`
	// The zone ID.
	//
	// example:
	//
	// cn-hangzhou-i
	ZoneId *string `json:"zoneId,omitempty" xml:"zoneId,omitempty"`
}

func (s CreateInstanceV1RequestBackendNodeGroups) String() string {
	return dara.Prettify(s)
}

func (s CreateInstanceV1RequestBackendNodeGroups) GoString() string {
	return s.String()
}

func (s *CreateInstanceV1RequestBackendNodeGroups) GetCu() *int32 {
	return s.Cu
}

func (s *CreateInstanceV1RequestBackendNodeGroups) GetDiskNumber() *int32 {
	return s.DiskNumber
}

func (s *CreateInstanceV1RequestBackendNodeGroups) GetLocalStorageInstanceType() *string {
	return s.LocalStorageInstanceType
}

func (s *CreateInstanceV1RequestBackendNodeGroups) GetResidentNodeNumber() *int32 {
	return s.ResidentNodeNumber
}

func (s *CreateInstanceV1RequestBackendNodeGroups) GetSpecType() *string {
	return s.SpecType
}

func (s *CreateInstanceV1RequestBackendNodeGroups) GetStoragePerformanceLevel() *string {
	return s.StoragePerformanceLevel
}

func (s *CreateInstanceV1RequestBackendNodeGroups) GetStorageSize() *int32 {
	return s.StorageSize
}

func (s *CreateInstanceV1RequestBackendNodeGroups) GetZoneId() *string {
	return s.ZoneId
}

func (s *CreateInstanceV1RequestBackendNodeGroups) SetCu(v int32) *CreateInstanceV1RequestBackendNodeGroups {
	s.Cu = &v
	return s
}

func (s *CreateInstanceV1RequestBackendNodeGroups) SetDiskNumber(v int32) *CreateInstanceV1RequestBackendNodeGroups {
	s.DiskNumber = &v
	return s
}

func (s *CreateInstanceV1RequestBackendNodeGroups) SetLocalStorageInstanceType(v string) *CreateInstanceV1RequestBackendNodeGroups {
	s.LocalStorageInstanceType = &v
	return s
}

func (s *CreateInstanceV1RequestBackendNodeGroups) SetResidentNodeNumber(v int32) *CreateInstanceV1RequestBackendNodeGroups {
	s.ResidentNodeNumber = &v
	return s
}

func (s *CreateInstanceV1RequestBackendNodeGroups) SetSpecType(v string) *CreateInstanceV1RequestBackendNodeGroups {
	s.SpecType = &v
	return s
}

func (s *CreateInstanceV1RequestBackendNodeGroups) SetStoragePerformanceLevel(v string) *CreateInstanceV1RequestBackendNodeGroups {
	s.StoragePerformanceLevel = &v
	return s
}

func (s *CreateInstanceV1RequestBackendNodeGroups) SetStorageSize(v int32) *CreateInstanceV1RequestBackendNodeGroups {
	s.StorageSize = &v
	return s
}

func (s *CreateInstanceV1RequestBackendNodeGroups) SetZoneId(v string) *CreateInstanceV1RequestBackendNodeGroups {
	s.ZoneId = &v
	return s
}

func (s *CreateInstanceV1RequestBackendNodeGroups) Validate() error {
	return dara.Validate(s)
}

type CreateInstanceV1RequestFrontendNodeGroups struct {
	// The number of compute units (CUs). A CU is the basic metering unit of the service. 1 CU equals 1 vCPU and 4 GiB of memory.
	//
	// example:
	//
	// 8
	Cu *int32 `json:"cu,omitempty" xml:"cu,omitempty"`
	// The number of disks.
	//
	// example:
	//
	// 1
	DiskNumber *int32 `json:"diskNumber,omitempty" xml:"diskNumber,omitempty"`
	// The local SSD instance type. You do not need to set this parameter for FE compute groups.
	//
	// example:
	//
	// null
	LocalStorageInstanceType *string `json:"localStorageInstanceType,omitempty" xml:"localStorageInstanceType,omitempty"`
	// The number of nodes.
	//
	// example:
	//
	// 3
	ResidentNodeNumber *int32 `json:"residentNodeNumber,omitempty" xml:"residentNodeNumber,omitempty"`
	// The specification type of the compute group. Only standard is supported.
	//
	// example:
	//
	// standard
	SpecType *string `json:"specType,omitempty" xml:"specType,omitempty"`
	// The performance level (PL) of the cloud disk. Only pl1 is supported, where a single disk delivers up to 50,000 random read/write IOPS.
	//
	// example:
	//
	// pl1
	StoragePerformanceLevel *string `json:"storagePerformanceLevel,omitempty" xml:"storagePerformanceLevel,omitempty"`
	// The storage size. Unit: GiB.
	//
	// example:
	//
	// 100
	StorageSize *int32 `json:"storageSize,omitempty" xml:"storageSize,omitempty"`
	// The zone ID.
	//
	// example:
	//
	// cn-hangzhou-i
	ZoneId *string `json:"zoneId,omitempty" xml:"zoneId,omitempty"`
}

func (s CreateInstanceV1RequestFrontendNodeGroups) String() string {
	return dara.Prettify(s)
}

func (s CreateInstanceV1RequestFrontendNodeGroups) GoString() string {
	return s.String()
}

func (s *CreateInstanceV1RequestFrontendNodeGroups) GetCu() *int32 {
	return s.Cu
}

func (s *CreateInstanceV1RequestFrontendNodeGroups) GetDiskNumber() *int32 {
	return s.DiskNumber
}

func (s *CreateInstanceV1RequestFrontendNodeGroups) GetLocalStorageInstanceType() *string {
	return s.LocalStorageInstanceType
}

func (s *CreateInstanceV1RequestFrontendNodeGroups) GetResidentNodeNumber() *int32 {
	return s.ResidentNodeNumber
}

func (s *CreateInstanceV1RequestFrontendNodeGroups) GetSpecType() *string {
	return s.SpecType
}

func (s *CreateInstanceV1RequestFrontendNodeGroups) GetStoragePerformanceLevel() *string {
	return s.StoragePerformanceLevel
}

func (s *CreateInstanceV1RequestFrontendNodeGroups) GetStorageSize() *int32 {
	return s.StorageSize
}

func (s *CreateInstanceV1RequestFrontendNodeGroups) GetZoneId() *string {
	return s.ZoneId
}

func (s *CreateInstanceV1RequestFrontendNodeGroups) SetCu(v int32) *CreateInstanceV1RequestFrontendNodeGroups {
	s.Cu = &v
	return s
}

func (s *CreateInstanceV1RequestFrontendNodeGroups) SetDiskNumber(v int32) *CreateInstanceV1RequestFrontendNodeGroups {
	s.DiskNumber = &v
	return s
}

func (s *CreateInstanceV1RequestFrontendNodeGroups) SetLocalStorageInstanceType(v string) *CreateInstanceV1RequestFrontendNodeGroups {
	s.LocalStorageInstanceType = &v
	return s
}

func (s *CreateInstanceV1RequestFrontendNodeGroups) SetResidentNodeNumber(v int32) *CreateInstanceV1RequestFrontendNodeGroups {
	s.ResidentNodeNumber = &v
	return s
}

func (s *CreateInstanceV1RequestFrontendNodeGroups) SetSpecType(v string) *CreateInstanceV1RequestFrontendNodeGroups {
	s.SpecType = &v
	return s
}

func (s *CreateInstanceV1RequestFrontendNodeGroups) SetStoragePerformanceLevel(v string) *CreateInstanceV1RequestFrontendNodeGroups {
	s.StoragePerformanceLevel = &v
	return s
}

func (s *CreateInstanceV1RequestFrontendNodeGroups) SetStorageSize(v int32) *CreateInstanceV1RequestFrontendNodeGroups {
	s.StorageSize = &v
	return s
}

func (s *CreateInstanceV1RequestFrontendNodeGroups) SetZoneId(v string) *CreateInstanceV1RequestFrontendNodeGroups {
	s.ZoneId = &v
	return s
}

func (s *CreateInstanceV1RequestFrontendNodeGroups) Validate() error {
	return dara.Validate(s)
}

type CreateInstanceV1RequestObserverNodeGroups struct {
	// The number of compute units (CUs). A CU is the basic metering unit of the service. 1 CU equals 1 vCPU and 4 GiB of memory.
	//
	// example:
	//
	// 8
	Cu *int32 `json:"cu,omitempty" xml:"cu,omitempty"`
	// The number of disks.
	//
	// example:
	//
	// 1
	DiskNumber *int32 `json:"diskNumber,omitempty" xml:"diskNumber,omitempty"`
	// The instance specification for local SSDs. You do not need to specify this parameter for the Observer compute group.
	//
	// example:
	//
	// null
	LocalStorageInstanceType *string `json:"localStorageInstanceType,omitempty" xml:"localStorageInstanceType,omitempty"`
	// The number of nodes.
	//
	// example:
	//
	// 3
	ResidentNodeNumber *int32 `json:"residentNodeNumber,omitempty" xml:"residentNodeNumber,omitempty"`
	// The specification type of the compute group. Only standard is supported.
	//
	// example:
	//
	// standard
	SpecType *string `json:"specType,omitempty" xml:"specType,omitempty"`
	// The performance level (PL) of the cloud disk. Only PL1 is supported, which delivers up to 50,000 random read/write IOPS per disk.
	//
	// example:
	//
	// pl1
	StoragePerformanceLevel *string `json:"storagePerformanceLevel,omitempty" xml:"storagePerformanceLevel,omitempty"`
	// The storage size. Unit: GiB.
	//
	// example:
	//
	// 100
	StorageSize *int32 `json:"storageSize,omitempty" xml:"storageSize,omitempty"`
	// The zone ID.
	//
	// example:
	//
	// cn-hangzhou-i
	ZoneId *string `json:"zoneId,omitempty" xml:"zoneId,omitempty"`
}

func (s CreateInstanceV1RequestObserverNodeGroups) String() string {
	return dara.Prettify(s)
}

func (s CreateInstanceV1RequestObserverNodeGroups) GoString() string {
	return s.String()
}

func (s *CreateInstanceV1RequestObserverNodeGroups) GetCu() *int32 {
	return s.Cu
}

func (s *CreateInstanceV1RequestObserverNodeGroups) GetDiskNumber() *int32 {
	return s.DiskNumber
}

func (s *CreateInstanceV1RequestObserverNodeGroups) GetLocalStorageInstanceType() *string {
	return s.LocalStorageInstanceType
}

func (s *CreateInstanceV1RequestObserverNodeGroups) GetResidentNodeNumber() *int32 {
	return s.ResidentNodeNumber
}

func (s *CreateInstanceV1RequestObserverNodeGroups) GetSpecType() *string {
	return s.SpecType
}

func (s *CreateInstanceV1RequestObserverNodeGroups) GetStoragePerformanceLevel() *string {
	return s.StoragePerformanceLevel
}

func (s *CreateInstanceV1RequestObserverNodeGroups) GetStorageSize() *int32 {
	return s.StorageSize
}

func (s *CreateInstanceV1RequestObserverNodeGroups) GetZoneId() *string {
	return s.ZoneId
}

func (s *CreateInstanceV1RequestObserverNodeGroups) SetCu(v int32) *CreateInstanceV1RequestObserverNodeGroups {
	s.Cu = &v
	return s
}

func (s *CreateInstanceV1RequestObserverNodeGroups) SetDiskNumber(v int32) *CreateInstanceV1RequestObserverNodeGroups {
	s.DiskNumber = &v
	return s
}

func (s *CreateInstanceV1RequestObserverNodeGroups) SetLocalStorageInstanceType(v string) *CreateInstanceV1RequestObserverNodeGroups {
	s.LocalStorageInstanceType = &v
	return s
}

func (s *CreateInstanceV1RequestObserverNodeGroups) SetResidentNodeNumber(v int32) *CreateInstanceV1RequestObserverNodeGroups {
	s.ResidentNodeNumber = &v
	return s
}

func (s *CreateInstanceV1RequestObserverNodeGroups) SetSpecType(v string) *CreateInstanceV1RequestObserverNodeGroups {
	s.SpecType = &v
	return s
}

func (s *CreateInstanceV1RequestObserverNodeGroups) SetStoragePerformanceLevel(v string) *CreateInstanceV1RequestObserverNodeGroups {
	s.StoragePerformanceLevel = &v
	return s
}

func (s *CreateInstanceV1RequestObserverNodeGroups) SetStorageSize(v int32) *CreateInstanceV1RequestObserverNodeGroups {
	s.StorageSize = &v
	return s
}

func (s *CreateInstanceV1RequestObserverNodeGroups) SetZoneId(v string) *CreateInstanceV1RequestObserverNodeGroups {
	s.ZoneId = &v
	return s
}

func (s *CreateInstanceV1RequestObserverNodeGroups) Validate() error {
	return dara.Validate(s)
}

type CreateInstanceV1RequestTags struct {
	// The tag key.
	//
	// example:
	//
	// k1
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value.
	//
	// example:
	//
	// v1
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateInstanceV1RequestTags) String() string {
	return dara.Prettify(s)
}

func (s CreateInstanceV1RequestTags) GoString() string {
	return s.String()
}

func (s *CreateInstanceV1RequestTags) GetKey() *string {
	return s.Key
}

func (s *CreateInstanceV1RequestTags) GetValue() *string {
	return s.Value
}

func (s *CreateInstanceV1RequestTags) SetKey(v string) *CreateInstanceV1RequestTags {
	s.Key = &v
	return s
}

func (s *CreateInstanceV1RequestTags) SetValue(v string) *CreateInstanceV1RequestTags {
	s.Value = &v
	return s
}

func (s *CreateInstanceV1RequestTags) Validate() error {
	return dara.Validate(s)
}

type CreateInstanceV1RequestVSwitches struct {
	// The vSwitch ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// vsw-bp19mlh98tm9teyyd****
	VswId *string `json:"VswId,omitempty" xml:"VswId,omitempty"`
	// The zone ID of the vSwitch.
	//
	// example:
	//
	// cn-hangzhou-i
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s CreateInstanceV1RequestVSwitches) String() string {
	return dara.Prettify(s)
}

func (s CreateInstanceV1RequestVSwitches) GoString() string {
	return s.String()
}

func (s *CreateInstanceV1RequestVSwitches) GetVswId() *string {
	return s.VswId
}

func (s *CreateInstanceV1RequestVSwitches) GetZoneId() *string {
	return s.ZoneId
}

func (s *CreateInstanceV1RequestVSwitches) SetVswId(v string) *CreateInstanceV1RequestVSwitches {
	s.VswId = &v
	return s
}

func (s *CreateInstanceV1RequestVSwitches) SetZoneId(v string) *CreateInstanceV1RequestVSwitches {
	s.ZoneId = &v
	return s
}

func (s *CreateInstanceV1RequestVSwitches) Validate() error {
	return dara.Validate(s)
}
