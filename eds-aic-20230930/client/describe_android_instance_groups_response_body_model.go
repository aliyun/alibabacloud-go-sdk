// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAndroidInstanceGroupsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceGroupModel(v []*DescribeAndroidInstanceGroupsResponseBodyInstanceGroupModel) *DescribeAndroidInstanceGroupsResponseBody
	GetInstanceGroupModel() []*DescribeAndroidInstanceGroupsResponseBodyInstanceGroupModel
	SetNextToken(v string) *DescribeAndroidInstanceGroupsResponseBody
	GetNextToken() *string
	SetRequestId(v string) *DescribeAndroidInstanceGroupsResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *DescribeAndroidInstanceGroupsResponseBody
	GetTotalCount() *int32
}

type DescribeAndroidInstanceGroupsResponseBody struct {
	// The details of the instance group.
	InstanceGroupModel []*DescribeAndroidInstanceGroupsResponseBodyInstanceGroupModel `json:"InstanceGroupModel,omitempty" xml:"InstanceGroupModel,omitempty" type:"Repeated"`
	// The pagination token that indicates the position where the current call returns. An empty value indicates that all data has been read.
	//
	// example:
	//
	// AAAAAV3MpHK1AP0pfERHZN5pu6l5V9uONHqPtDLM2U8s****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The request ID.
	//
	// example:
	//
	// F07A1DA1-E1EB-5CCA-8EED-12F85D32****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries.
	//
	// example:
	//
	// 10
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeAndroidInstanceGroupsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeAndroidInstanceGroupsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAndroidInstanceGroupsResponseBody) GetInstanceGroupModel() []*DescribeAndroidInstanceGroupsResponseBodyInstanceGroupModel {
	return s.InstanceGroupModel
}

func (s *DescribeAndroidInstanceGroupsResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *DescribeAndroidInstanceGroupsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeAndroidInstanceGroupsResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeAndroidInstanceGroupsResponseBody) SetInstanceGroupModel(v []*DescribeAndroidInstanceGroupsResponseBodyInstanceGroupModel) *DescribeAndroidInstanceGroupsResponseBody {
	s.InstanceGroupModel = v
	return s
}

func (s *DescribeAndroidInstanceGroupsResponseBody) SetNextToken(v string) *DescribeAndroidInstanceGroupsResponseBody {
	s.NextToken = &v
	return s
}

func (s *DescribeAndroidInstanceGroupsResponseBody) SetRequestId(v string) *DescribeAndroidInstanceGroupsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeAndroidInstanceGroupsResponseBody) SetTotalCount(v int32) *DescribeAndroidInstanceGroupsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeAndroidInstanceGroupsResponseBody) Validate() error {
	if s.InstanceGroupModel != nil {
		for _, item := range s.InstanceGroupModel {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeAndroidInstanceGroupsResponseBodyInstanceGroupModel struct {
	// The delivery group ID.
	//
	// example:
	//
	// aig-48xr63m4dybjk****
	AppInstanceGroupId *string `json:"AppInstanceGroupId,omitempty" xml:"AppInstanceGroupId,omitempty"`
	// The architecture type.
	//
	// example:
	//
	// ARM
	ArchitectureType *string `json:"ArchitectureType,omitempty" xml:"ArchitectureType,omitempty"`
	// The number of active instances.
	//
	// > An instance is considered active if its instance status is not "Deleting" or "Deleted".
	//
	// example:
	//
	// 2
	AvailableInstanceAmount *int32 `json:"AvailableInstanceAmount,omitempty" xml:"AvailableInstanceAmount,omitempty"`
	// The ID of the bandwidth package.
	//
	// example:
	//
	// np-0q6ixs7vpxciz****
	BandwidthPackageId *string `json:"BandwidthPackageId,omitempty" xml:"BandwidthPackageId,omitempty"`
	// The status of the bandwidth package.
	//
	// Valid values:
	//
	// - Creating: being created.
	//
	// - Releasing: being released.
	//
	// - InUse: in use.
	//
	// - Failed: failed.
	//
	// - Expired: expired.
	//
	// - Available: unbound and being billed.
	//
	// example:
	//
	// Creating
	BandwidthPackageStatus *string `json:"BandwidthPackageStatus,omitempty" xml:"BandwidthPackageStatus,omitempty"`
	// The type of the bandwidth package.
	//
	// example:
	//
	// cbwp_ecd
	BandwidthPackageType *string `json:"BandwidthPackageType,omitempty" xml:"BandwidthPackageType,omitempty"`
	// The public network bandwidth throttling rules for the instance group.
	BindQosRules *DescribeAndroidInstanceGroupsResponseBodyInstanceGroupModelBindQosRules `json:"BindQosRules,omitempty" xml:"BindQosRules,omitempty" type:"Struct"`
	Channel      *string                                                                  `json:"Channel,omitempty" xml:"Channel,omitempty"`
	// The billing method.
	//
	// example:
	//
	// PrePaid
	ChargeType *string `json:"ChargeType,omitempty" xml:"ChargeType,omitempty"`
	// The number of CPU cores.
	//
	// example:
	//
	// 8
	Cpu *string `json:"Cpu,omitempty" xml:"Cpu,omitempty"`
	// The disk information.
	Disks []*DescribeAndroidInstanceGroupsResponseBodyInstanceGroupModelDisks `json:"Disks,omitempty" xml:"Disks,omitempty" type:"Repeated"`
	// > This parameter is not publicly available.
	//
	// example:
	//
	// true
	EnableIpv6 *bool `json:"EnableIpv6,omitempty" xml:"EnableIpv6,omitempty"`
	// The reason for the creation failure.
	//
	// example:
	//
	// InternalError
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The creation time.
	//
	// example:
	//
	// 2024-02-01 10:56:36
	GmtCreate *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// The expiration time of the subscription instance group.
	//
	// example:
	//
	// 2027-06-29 07:25:31
	GmtExpired *string `json:"GmtExpired,omitempty" xml:"GmtExpired,omitempty"`
	// The update time.
	//
	// example:
	//
	// 2024-02-01 10:56:36
	GmtModified *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// The image ID.
	//
	// example:
	//
	// imgc-06zyt9m93zwax****
	ImageId *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	// The image version.
	//
	// example:
	//
	// 25.09.2
	ImageVersion *string `json:"ImageVersion,omitempty" xml:"ImageVersion,omitempty"`
	// The list of installed applications.
	//
	// example:
	//
	// "抖音","淘宝"
	InstalledAppList *string `json:"InstalledAppList,omitempty" xml:"InstalledAppList,omitempty"`
	// The instance group ID.
	//
	// example:
	//
	// ag-h67a2cs0zprfdh****
	InstanceGroupId *string `json:"InstanceGroupId,omitempty" xml:"InstanceGroupId,omitempty"`
	// The instance group name.
	//
	// example:
	//
	// Cloud phoneA
	InstanceGroupName *string `json:"InstanceGroupName,omitempty" xml:"InstanceGroupName,omitempty"`
	// The instance group specifications.
	//
	// example:
	//
	// acp.basic.small
	InstanceGroupSpec *string `json:"InstanceGroupSpec,omitempty" xml:"InstanceGroupSpec,omitempty"`
	// The description of the instance group specifications.
	//
	// example:
	//
	// ARM-2vCPU4GiB 32GiB
	InstanceGroupSpecDescribe *string `json:"InstanceGroupSpecDescribe,omitempty" xml:"InstanceGroupSpecDescribe,omitempty"`
	// The instance group status.
	//
	// example:
	//
	// RUNNING
	InstanceGroupStatus *string `json:"InstanceGroupStatus,omitempty" xml:"InstanceGroupStatus,omitempty"`
	// > This parameter is not publicly available.
	//
	// example:
	//
	// 50
	Ipv6Bandwidth *int32 `json:"Ipv6Bandwidth,omitempty" xml:"Ipv6Bandwidth,omitempty"`
	// The memory size.
	//
	// example:
	//
	// 8
	Memory *int32 `json:"Memory,omitempty" xml:"Memory,omitempty"`
	// The network type of the instance.
	//
	// > This field is returned only for instance groups with a standard network.
	//
	// example:
	//
	// network_pro_ecd
	NetworkType *string `json:"NetworkType,omitempty" xml:"NetworkType,omitempty"`
	// The number of instances in the instance group.
	//
	// example:
	//
	// 10
	NumberOfInstances *string `json:"NumberOfInstances,omitempty" xml:"NumberOfInstances,omitempty"`
	// The network ID.
	//
	// example:
	//
	// cn-shanghai+dir-030598****
	OfficeSiteId *string `json:"OfficeSiteId,omitempty" xml:"OfficeSiteId,omitempty"`
	PackageId    *string `json:"PackageId,omitempty" xml:"PackageId,omitempty"`
	// The policy ID.
	//
	// example:
	//
	// pg-c6n38xucps8kl****
	PolicyGroupId *string `json:"PolicyGroupId,omitempty" xml:"PolicyGroupId,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The rendering type of the instance group.
	//
	// example:
	//
	// CPU
	RenderingType *string `json:"RenderingType,omitempty" xml:"RenderingType,omitempty"`
	// The height of the resolution.
	//
	// example:
	//
	// 1280
	ResolutionHeight *int32 `json:"ResolutionHeight,omitempty" xml:"ResolutionHeight,omitempty"`
	// The width of the resolution.
	//
	// example:
	//
	// 720
	ResolutionWidth *int32 `json:"ResolutionWidth,omitempty" xml:"ResolutionWidth,omitempty"`
	// The sales mode.
	//
	// example:
	//
	// standard
	SaleMode *string `json:"SaleMode,omitempty" xml:"SaleMode,omitempty"`
	// The system version.
	//
	// example:
	//
	// Android 12
	SystemVersion *string `json:"SystemVersion,omitempty" xml:"SystemVersion,omitempty"`
	// The tag information.
	Tags []*DescribeAndroidInstanceGroupsResponseBodyInstanceGroupModelTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	// The vSwitch ID in the VPC.
	//
	// example:
	//
	// vsw-t4n0yqs009ho024wt****
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// The zone ID.
	//
	// example:
	//
	// cn-hangzhou-j
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s DescribeAndroidInstanceGroupsResponseBodyInstanceGroupModel) String() string {
	return dara.Prettify(s)
}

func (s DescribeAndroidInstanceGroupsResponseBodyInstanceGroupModel) GoString() string {
	return s.String()
}

func (s *DescribeAndroidInstanceGroupsResponseBodyInstanceGroupModel) GetAppInstanceGroupId() *string {
	return s.AppInstanceGroupId
}

func (s *DescribeAndroidInstanceGroupsResponseBodyInstanceGroupModel) GetArchitectureType() *string {
	return s.ArchitectureType
}

func (s *DescribeAndroidInstanceGroupsResponseBodyInstanceGroupModel) GetAvailableInstanceAmount() *int32 {
	return s.AvailableInstanceAmount
}

func (s *DescribeAndroidInstanceGroupsResponseBodyInstanceGroupModel) GetBandwidthPackageId() *string {
	return s.BandwidthPackageId
}

func (s *DescribeAndroidInstanceGroupsResponseBodyInstanceGroupModel) GetBandwidthPackageStatus() *string {
	return s.BandwidthPackageStatus
}

func (s *DescribeAndroidInstanceGroupsResponseBodyInstanceGroupModel) GetBandwidthPackageType() *string {
	return s.BandwidthPackageType
}

func (s *DescribeAndroidInstanceGroupsResponseBodyInstanceGroupModel) GetBindQosRules() *DescribeAndroidInstanceGroupsResponseBodyInstanceGroupModelBindQosRules {
	return s.BindQosRules
}

func (s *DescribeAndroidInstanceGroupsResponseBodyInstanceGroupModel) GetChannel() *string {
	return s.Channel
}

func (s *DescribeAndroidInstanceGroupsResponseBodyInstanceGroupModel) GetChargeType() *string {
	return s.ChargeType
}

func (s *DescribeAndroidInstanceGroupsResponseBodyInstanceGroupModel) GetCpu() *string {
	return s.Cpu
}

func (s *DescribeAndroidInstanceGroupsResponseBodyInstanceGroupModel) GetDisks() []*DescribeAndroidInstanceGroupsResponseBodyInstanceGroupModelDisks {
	return s.Disks
}

func (s *DescribeAndroidInstanceGroupsResponseBodyInstanceGroupModel) GetEnableIpv6() *bool {
	return s.EnableIpv6
}

func (s *DescribeAndroidInstanceGroupsResponseBodyInstanceGroupModel) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *DescribeAndroidInstanceGroupsResponseBodyInstanceGroupModel) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *DescribeAndroidInstanceGroupsResponseBodyInstanceGroupModel) GetGmtExpired() *string {
	return s.GmtExpired
}

func (s *DescribeAndroidInstanceGroupsResponseBodyInstanceGroupModel) GetGmtModified() *string {
	return s.GmtModified
}

func (s *DescribeAndroidInstanceGroupsResponseBodyInstanceGroupModel) GetImageId() *string {
	return s.ImageId
}

func (s *DescribeAndroidInstanceGroupsResponseBodyInstanceGroupModel) GetImageVersion() *string {
	return s.ImageVersion
}

func (s *DescribeAndroidInstanceGroupsResponseBodyInstanceGroupModel) GetInstalledAppList() *string {
	return s.InstalledAppList
}

func (s *DescribeAndroidInstanceGroupsResponseBodyInstanceGroupModel) GetInstanceGroupId() *string {
	return s.InstanceGroupId
}

func (s *DescribeAndroidInstanceGroupsResponseBodyInstanceGroupModel) GetInstanceGroupName() *string {
	return s.InstanceGroupName
}

func (s *DescribeAndroidInstanceGroupsResponseBodyInstanceGroupModel) GetInstanceGroupSpec() *string {
	return s.InstanceGroupSpec
}

func (s *DescribeAndroidInstanceGroupsResponseBodyInstanceGroupModel) GetInstanceGroupSpecDescribe() *string {
	return s.InstanceGroupSpecDescribe
}

func (s *DescribeAndroidInstanceGroupsResponseBodyInstanceGroupModel) GetInstanceGroupStatus() *string {
	return s.InstanceGroupStatus
}

func (s *DescribeAndroidInstanceGroupsResponseBodyInstanceGroupModel) GetIpv6Bandwidth() *int32 {
	return s.Ipv6Bandwidth
}

func (s *DescribeAndroidInstanceGroupsResponseBodyInstanceGroupModel) GetMemory() *int32 {
	return s.Memory
}

func (s *DescribeAndroidInstanceGroupsResponseBodyInstanceGroupModel) GetNetworkType() *string {
	return s.NetworkType
}

func (s *DescribeAndroidInstanceGroupsResponseBodyInstanceGroupModel) GetNumberOfInstances() *string {
	return s.NumberOfInstances
}

func (s *DescribeAndroidInstanceGroupsResponseBodyInstanceGroupModel) GetOfficeSiteId() *string {
	return s.OfficeSiteId
}

func (s *DescribeAndroidInstanceGroupsResponseBodyInstanceGroupModel) GetPackageId() *string {
	return s.PackageId
}

func (s *DescribeAndroidInstanceGroupsResponseBodyInstanceGroupModel) GetPolicyGroupId() *string {
	return s.PolicyGroupId
}

func (s *DescribeAndroidInstanceGroupsResponseBodyInstanceGroupModel) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeAndroidInstanceGroupsResponseBodyInstanceGroupModel) GetRenderingType() *string {
	return s.RenderingType
}

func (s *DescribeAndroidInstanceGroupsResponseBodyInstanceGroupModel) GetResolutionHeight() *int32 {
	return s.ResolutionHeight
}

func (s *DescribeAndroidInstanceGroupsResponseBodyInstanceGroupModel) GetResolutionWidth() *int32 {
	return s.ResolutionWidth
}

func (s *DescribeAndroidInstanceGroupsResponseBodyInstanceGroupModel) GetSaleMode() *string {
	return s.SaleMode
}

func (s *DescribeAndroidInstanceGroupsResponseBodyInstanceGroupModel) GetSystemVersion() *string {
	return s.SystemVersion
}

func (s *DescribeAndroidInstanceGroupsResponseBodyInstanceGroupModel) GetTags() []*DescribeAndroidInstanceGroupsResponseBodyInstanceGroupModelTags {
	return s.Tags
}

func (s *DescribeAndroidInstanceGroupsResponseBodyInstanceGroupModel) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *DescribeAndroidInstanceGroupsResponseBodyInstanceGroupModel) GetZoneId() *string {
	return s.ZoneId
}

func (s *DescribeAndroidInstanceGroupsResponseBodyInstanceGroupModel) SetAppInstanceGroupId(v string) *DescribeAndroidInstanceGroupsResponseBodyInstanceGroupModel {
	s.AppInstanceGroupId = &v
	return s
}

func (s *DescribeAndroidInstanceGroupsResponseBodyInstanceGroupModel) SetArchitectureType(v string) *DescribeAndroidInstanceGroupsResponseBodyInstanceGroupModel {
	s.ArchitectureType = &v
	return s
}

func (s *DescribeAndroidInstanceGroupsResponseBodyInstanceGroupModel) SetAvailableInstanceAmount(v int32) *DescribeAndroidInstanceGroupsResponseBodyInstanceGroupModel {
	s.AvailableInstanceAmount = &v
	return s
}

func (s *DescribeAndroidInstanceGroupsResponseBodyInstanceGroupModel) SetBandwidthPackageId(v string) *DescribeAndroidInstanceGroupsResponseBodyInstanceGroupModel {
	s.BandwidthPackageId = &v
	return s
}

func (s *DescribeAndroidInstanceGroupsResponseBodyInstanceGroupModel) SetBandwidthPackageStatus(v string) *DescribeAndroidInstanceGroupsResponseBodyInstanceGroupModel {
	s.BandwidthPackageStatus = &v
	return s
}

func (s *DescribeAndroidInstanceGroupsResponseBodyInstanceGroupModel) SetBandwidthPackageType(v string) *DescribeAndroidInstanceGroupsResponseBodyInstanceGroupModel {
	s.BandwidthPackageType = &v
	return s
}

func (s *DescribeAndroidInstanceGroupsResponseBodyInstanceGroupModel) SetBindQosRules(v *DescribeAndroidInstanceGroupsResponseBodyInstanceGroupModelBindQosRules) *DescribeAndroidInstanceGroupsResponseBodyInstanceGroupModel {
	s.BindQosRules = v
	return s
}

func (s *DescribeAndroidInstanceGroupsResponseBodyInstanceGroupModel) SetChannel(v string) *DescribeAndroidInstanceGroupsResponseBodyInstanceGroupModel {
	s.Channel = &v
	return s
}

func (s *DescribeAndroidInstanceGroupsResponseBodyInstanceGroupModel) SetChargeType(v string) *DescribeAndroidInstanceGroupsResponseBodyInstanceGroupModel {
	s.ChargeType = &v
	return s
}

func (s *DescribeAndroidInstanceGroupsResponseBodyInstanceGroupModel) SetCpu(v string) *DescribeAndroidInstanceGroupsResponseBodyInstanceGroupModel {
	s.Cpu = &v
	return s
}

func (s *DescribeAndroidInstanceGroupsResponseBodyInstanceGroupModel) SetDisks(v []*DescribeAndroidInstanceGroupsResponseBodyInstanceGroupModelDisks) *DescribeAndroidInstanceGroupsResponseBodyInstanceGroupModel {
	s.Disks = v
	return s
}

func (s *DescribeAndroidInstanceGroupsResponseBodyInstanceGroupModel) SetEnableIpv6(v bool) *DescribeAndroidInstanceGroupsResponseBodyInstanceGroupModel {
	s.EnableIpv6 = &v
	return s
}

func (s *DescribeAndroidInstanceGroupsResponseBodyInstanceGroupModel) SetErrorCode(v string) *DescribeAndroidInstanceGroupsResponseBodyInstanceGroupModel {
	s.ErrorCode = &v
	return s
}

func (s *DescribeAndroidInstanceGroupsResponseBodyInstanceGroupModel) SetGmtCreate(v string) *DescribeAndroidInstanceGroupsResponseBodyInstanceGroupModel {
	s.GmtCreate = &v
	return s
}

func (s *DescribeAndroidInstanceGroupsResponseBodyInstanceGroupModel) SetGmtExpired(v string) *DescribeAndroidInstanceGroupsResponseBodyInstanceGroupModel {
	s.GmtExpired = &v
	return s
}

func (s *DescribeAndroidInstanceGroupsResponseBodyInstanceGroupModel) SetGmtModified(v string) *DescribeAndroidInstanceGroupsResponseBodyInstanceGroupModel {
	s.GmtModified = &v
	return s
}

func (s *DescribeAndroidInstanceGroupsResponseBodyInstanceGroupModel) SetImageId(v string) *DescribeAndroidInstanceGroupsResponseBodyInstanceGroupModel {
	s.ImageId = &v
	return s
}

func (s *DescribeAndroidInstanceGroupsResponseBodyInstanceGroupModel) SetImageVersion(v string) *DescribeAndroidInstanceGroupsResponseBodyInstanceGroupModel {
	s.ImageVersion = &v
	return s
}

func (s *DescribeAndroidInstanceGroupsResponseBodyInstanceGroupModel) SetInstalledAppList(v string) *DescribeAndroidInstanceGroupsResponseBodyInstanceGroupModel {
	s.InstalledAppList = &v
	return s
}

func (s *DescribeAndroidInstanceGroupsResponseBodyInstanceGroupModel) SetInstanceGroupId(v string) *DescribeAndroidInstanceGroupsResponseBodyInstanceGroupModel {
	s.InstanceGroupId = &v
	return s
}

func (s *DescribeAndroidInstanceGroupsResponseBodyInstanceGroupModel) SetInstanceGroupName(v string) *DescribeAndroidInstanceGroupsResponseBodyInstanceGroupModel {
	s.InstanceGroupName = &v
	return s
}

func (s *DescribeAndroidInstanceGroupsResponseBodyInstanceGroupModel) SetInstanceGroupSpec(v string) *DescribeAndroidInstanceGroupsResponseBodyInstanceGroupModel {
	s.InstanceGroupSpec = &v
	return s
}

func (s *DescribeAndroidInstanceGroupsResponseBodyInstanceGroupModel) SetInstanceGroupSpecDescribe(v string) *DescribeAndroidInstanceGroupsResponseBodyInstanceGroupModel {
	s.InstanceGroupSpecDescribe = &v
	return s
}

func (s *DescribeAndroidInstanceGroupsResponseBodyInstanceGroupModel) SetInstanceGroupStatus(v string) *DescribeAndroidInstanceGroupsResponseBodyInstanceGroupModel {
	s.InstanceGroupStatus = &v
	return s
}

func (s *DescribeAndroidInstanceGroupsResponseBodyInstanceGroupModel) SetIpv6Bandwidth(v int32) *DescribeAndroidInstanceGroupsResponseBodyInstanceGroupModel {
	s.Ipv6Bandwidth = &v
	return s
}

func (s *DescribeAndroidInstanceGroupsResponseBodyInstanceGroupModel) SetMemory(v int32) *DescribeAndroidInstanceGroupsResponseBodyInstanceGroupModel {
	s.Memory = &v
	return s
}

func (s *DescribeAndroidInstanceGroupsResponseBodyInstanceGroupModel) SetNetworkType(v string) *DescribeAndroidInstanceGroupsResponseBodyInstanceGroupModel {
	s.NetworkType = &v
	return s
}

func (s *DescribeAndroidInstanceGroupsResponseBodyInstanceGroupModel) SetNumberOfInstances(v string) *DescribeAndroidInstanceGroupsResponseBodyInstanceGroupModel {
	s.NumberOfInstances = &v
	return s
}

func (s *DescribeAndroidInstanceGroupsResponseBodyInstanceGroupModel) SetOfficeSiteId(v string) *DescribeAndroidInstanceGroupsResponseBodyInstanceGroupModel {
	s.OfficeSiteId = &v
	return s
}

func (s *DescribeAndroidInstanceGroupsResponseBodyInstanceGroupModel) SetPackageId(v string) *DescribeAndroidInstanceGroupsResponseBodyInstanceGroupModel {
	s.PackageId = &v
	return s
}

func (s *DescribeAndroidInstanceGroupsResponseBodyInstanceGroupModel) SetPolicyGroupId(v string) *DescribeAndroidInstanceGroupsResponseBodyInstanceGroupModel {
	s.PolicyGroupId = &v
	return s
}

func (s *DescribeAndroidInstanceGroupsResponseBodyInstanceGroupModel) SetRegionId(v string) *DescribeAndroidInstanceGroupsResponseBodyInstanceGroupModel {
	s.RegionId = &v
	return s
}

func (s *DescribeAndroidInstanceGroupsResponseBodyInstanceGroupModel) SetRenderingType(v string) *DescribeAndroidInstanceGroupsResponseBodyInstanceGroupModel {
	s.RenderingType = &v
	return s
}

func (s *DescribeAndroidInstanceGroupsResponseBodyInstanceGroupModel) SetResolutionHeight(v int32) *DescribeAndroidInstanceGroupsResponseBodyInstanceGroupModel {
	s.ResolutionHeight = &v
	return s
}

func (s *DescribeAndroidInstanceGroupsResponseBodyInstanceGroupModel) SetResolutionWidth(v int32) *DescribeAndroidInstanceGroupsResponseBodyInstanceGroupModel {
	s.ResolutionWidth = &v
	return s
}

func (s *DescribeAndroidInstanceGroupsResponseBodyInstanceGroupModel) SetSaleMode(v string) *DescribeAndroidInstanceGroupsResponseBodyInstanceGroupModel {
	s.SaleMode = &v
	return s
}

func (s *DescribeAndroidInstanceGroupsResponseBodyInstanceGroupModel) SetSystemVersion(v string) *DescribeAndroidInstanceGroupsResponseBodyInstanceGroupModel {
	s.SystemVersion = &v
	return s
}

func (s *DescribeAndroidInstanceGroupsResponseBodyInstanceGroupModel) SetTags(v []*DescribeAndroidInstanceGroupsResponseBodyInstanceGroupModelTags) *DescribeAndroidInstanceGroupsResponseBodyInstanceGroupModel {
	s.Tags = v
	return s
}

func (s *DescribeAndroidInstanceGroupsResponseBodyInstanceGroupModel) SetVSwitchId(v string) *DescribeAndroidInstanceGroupsResponseBodyInstanceGroupModel {
	s.VSwitchId = &v
	return s
}

func (s *DescribeAndroidInstanceGroupsResponseBodyInstanceGroupModel) SetZoneId(v string) *DescribeAndroidInstanceGroupsResponseBodyInstanceGroupModel {
	s.ZoneId = &v
	return s
}

func (s *DescribeAndroidInstanceGroupsResponseBodyInstanceGroupModel) Validate() error {
	if s.BindQosRules != nil {
		if err := s.BindQosRules.Validate(); err != nil {
			return err
		}
	}
	if s.Disks != nil {
		for _, item := range s.Disks {
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
	return nil
}

type DescribeAndroidInstanceGroupsResponseBodyInstanceGroupModelBindQosRules struct {
	// The public network bandwidth throttling rules bound to the instance.
	InstanceQosRule []*DescribeAndroidInstanceGroupsResponseBodyInstanceGroupModelBindQosRulesInstanceQosRule `json:"InstanceQosRule,omitempty" xml:"InstanceQosRule,omitempty" type:"Repeated"`
	// The total number of public network bandwidth throttling rules for the instance group.
	//
	// example:
	//
	// 2
	TotalCount *int32 `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
}

func (s DescribeAndroidInstanceGroupsResponseBodyInstanceGroupModelBindQosRules) String() string {
	return dara.Prettify(s)
}

func (s DescribeAndroidInstanceGroupsResponseBodyInstanceGroupModelBindQosRules) GoString() string {
	return s.String()
}

func (s *DescribeAndroidInstanceGroupsResponseBodyInstanceGroupModelBindQosRules) GetInstanceQosRule() []*DescribeAndroidInstanceGroupsResponseBodyInstanceGroupModelBindQosRulesInstanceQosRule {
	return s.InstanceQosRule
}

func (s *DescribeAndroidInstanceGroupsResponseBodyInstanceGroupModelBindQosRules) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeAndroidInstanceGroupsResponseBodyInstanceGroupModelBindQosRules) SetInstanceQosRule(v []*DescribeAndroidInstanceGroupsResponseBodyInstanceGroupModelBindQosRulesInstanceQosRule) *DescribeAndroidInstanceGroupsResponseBodyInstanceGroupModelBindQosRules {
	s.InstanceQosRule = v
	return s
}

func (s *DescribeAndroidInstanceGroupsResponseBodyInstanceGroupModelBindQosRules) SetTotalCount(v int32) *DescribeAndroidInstanceGroupsResponseBodyInstanceGroupModelBindQosRules {
	s.TotalCount = &v
	return s
}

func (s *DescribeAndroidInstanceGroupsResponseBodyInstanceGroupModelBindQosRules) Validate() error {
	if s.InstanceQosRule != nil {
		for _, item := range s.InstanceQosRule {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeAndroidInstanceGroupsResponseBodyInstanceGroupModelBindQosRulesInstanceQosRule struct {
	// The instance ID.
	//
	// example:
	//
	// acp-h3m8b5dusopp5****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The ID of the public network bandwidth throttling rule. This rule applies only to premium bandwidth.
	//
	// example:
	//
	// qos-3kh93uu0vdbka****
	QosRuleId *string `json:"QosRuleId,omitempty" xml:"QosRuleId,omitempty"`
}

func (s DescribeAndroidInstanceGroupsResponseBodyInstanceGroupModelBindQosRulesInstanceQosRule) String() string {
	return dara.Prettify(s)
}

func (s DescribeAndroidInstanceGroupsResponseBodyInstanceGroupModelBindQosRulesInstanceQosRule) GoString() string {
	return s.String()
}

func (s *DescribeAndroidInstanceGroupsResponseBodyInstanceGroupModelBindQosRulesInstanceQosRule) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeAndroidInstanceGroupsResponseBodyInstanceGroupModelBindQosRulesInstanceQosRule) GetQosRuleId() *string {
	return s.QosRuleId
}

func (s *DescribeAndroidInstanceGroupsResponseBodyInstanceGroupModelBindQosRulesInstanceQosRule) SetInstanceId(v string) *DescribeAndroidInstanceGroupsResponseBodyInstanceGroupModelBindQosRulesInstanceQosRule {
	s.InstanceId = &v
	return s
}

func (s *DescribeAndroidInstanceGroupsResponseBodyInstanceGroupModelBindQosRulesInstanceQosRule) SetQosRuleId(v string) *DescribeAndroidInstanceGroupsResponseBodyInstanceGroupModelBindQosRulesInstanceQosRule {
	s.QosRuleId = &v
	return s
}

func (s *DescribeAndroidInstanceGroupsResponseBodyInstanceGroupModelBindQosRulesInstanceQosRule) Validate() error {
	return dara.Validate(s)
}

type DescribeAndroidInstanceGroupsResponseBodyInstanceGroupModelDisks struct {
	// The disk size, in GB.
	//
	// example:
	//
	// 32
	DiskSize *int32 `json:"DiskSize,omitempty" xml:"DiskSize,omitempty"`
	// The disk type.
	//
	// example:
	//
	// SYSTEM
	DiskType *string `json:"DiskType,omitempty" xml:"DiskType,omitempty"`
}

func (s DescribeAndroidInstanceGroupsResponseBodyInstanceGroupModelDisks) String() string {
	return dara.Prettify(s)
}

func (s DescribeAndroidInstanceGroupsResponseBodyInstanceGroupModelDisks) GoString() string {
	return s.String()
}

func (s *DescribeAndroidInstanceGroupsResponseBodyInstanceGroupModelDisks) GetDiskSize() *int32 {
	return s.DiskSize
}

func (s *DescribeAndroidInstanceGroupsResponseBodyInstanceGroupModelDisks) GetDiskType() *string {
	return s.DiskType
}

func (s *DescribeAndroidInstanceGroupsResponseBodyInstanceGroupModelDisks) SetDiskSize(v int32) *DescribeAndroidInstanceGroupsResponseBodyInstanceGroupModelDisks {
	s.DiskSize = &v
	return s
}

func (s *DescribeAndroidInstanceGroupsResponseBodyInstanceGroupModelDisks) SetDiskType(v string) *DescribeAndroidInstanceGroupsResponseBodyInstanceGroupModelDisks {
	s.DiskType = &v
	return s
}

func (s *DescribeAndroidInstanceGroupsResponseBodyInstanceGroupModelDisks) Validate() error {
	return dara.Validate(s)
}

type DescribeAndroidInstanceGroupsResponseBodyInstanceGroupModelTags struct {
	// The tag key.
	//
	// example:
	//
	// phone
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value.
	//
	// example:
	//
	// 2025
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeAndroidInstanceGroupsResponseBodyInstanceGroupModelTags) String() string {
	return dara.Prettify(s)
}

func (s DescribeAndroidInstanceGroupsResponseBodyInstanceGroupModelTags) GoString() string {
	return s.String()
}

func (s *DescribeAndroidInstanceGroupsResponseBodyInstanceGroupModelTags) GetKey() *string {
	return s.Key
}

func (s *DescribeAndroidInstanceGroupsResponseBodyInstanceGroupModelTags) GetValue() *string {
	return s.Value
}

func (s *DescribeAndroidInstanceGroupsResponseBodyInstanceGroupModelTags) SetKey(v string) *DescribeAndroidInstanceGroupsResponseBodyInstanceGroupModelTags {
	s.Key = &v
	return s
}

func (s *DescribeAndroidInstanceGroupsResponseBodyInstanceGroupModelTags) SetValue(v string) *DescribeAndroidInstanceGroupsResponseBodyInstanceGroupModelTags {
	s.Value = &v
	return s
}

func (s *DescribeAndroidInstanceGroupsResponseBodyInstanceGroupModelTags) Validate() error {
	return dara.Validate(s)
}
