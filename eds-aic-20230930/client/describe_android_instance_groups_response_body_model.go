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
	// The instance group.
	InstanceGroupModel []*DescribeAndroidInstanceGroupsResponseBodyInstanceGroupModel `json:"InstanceGroupModel,omitempty" xml:"InstanceGroupModel,omitempty" type:"Repeated"`
	// A pagination token. It can be used in the next request to retrieve a new page of results. If NextToken is empty, no next page exists.
	//
	// example:
	//
	// AAAAAV3MpHK1AP0pfERHZN5pu6l5V9uONHqPtDLM2U8s****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// F07A1DA1-E1EB-5CCA-8EED-12F85D32****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
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
	return dara.Validate(s)
}

type DescribeAndroidInstanceGroupsResponseBodyInstanceGroupModel struct {
	// The ID of the delivery group.
	//
	// example:
	//
	// aig-48xr63m4dybjk****
	AppInstanceGroupId *string `json:"AppInstanceGroupId,omitempty" xml:"AppInstanceGroupId,omitempty"`
	// The type of the architecture.
	//
	// example:
	//
	// ARM
	ArchitectureType *string `json:"ArchitectureType,omitempty" xml:"ArchitectureType,omitempty"`
	// The number of available instances.
	//
	// >  Available instances are those not in the Deleting or Deleted state.
	//
	// example:
	//
	// 5
	AvailableInstanceAmount *int32 `json:"AvailableInstanceAmount,omitempty" xml:"AvailableInstanceAmount,omitempty"`
	// The billing method.
	//
	// example:
	//
	// PrePaid
	ChargeType *string `json:"ChargeType,omitempty" xml:"ChargeType,omitempty"`
	// The number of vCPUs.
	//
	// example:
	//
	// 8
	Cpu *string `json:"Cpu,omitempty" xml:"Cpu,omitempty"`
	// The disks.
	Disks []*DescribeAndroidInstanceGroupsResponseBodyInstanceGroupModelDisks `json:"Disks,omitempty" xml:"Disks,omitempty" type:"Repeated"`
	// example:
	//
	// true
	EnableIpv6 *bool `json:"EnableIpv6,omitempty" xml:"EnableIpv6,omitempty"`
	// The cause of the creation failure.
	//
	// example:
	//
	// 0
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The time when the instance group was created.
	//
	// example:
	//
	// 2024-02-01 10:56:36
	GmtCreate *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// The time when the subscription instance group expires.
	//
	// example:
	//
	// 2027-06-29 07:25:31
	GmtExpired *string `json:"GmtExpired,omitempty" xml:"GmtExpired,omitempty"`
	// The time when the instance group was updated.
	//
	// example:
	//
	// 2024-02-01 10:56:36
	GmtModified *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// The ID of the image.
	//
	// example:
	//
	// imgc-06zyt9m93zwax****
	ImageId *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	// The list of installed applications.
	//
	// example:
	//
	// "TikTok","WeChat"
	InstalledAppList *string `json:"InstalledAppList,omitempty" xml:"InstalledAppList,omitempty"`
	// The ID of the instance group.
	//
	// example:
	//
	// ag-h67a2cs0zprfdh****
	InstanceGroupId *string `json:"InstanceGroupId,omitempty" xml:"InstanceGroupId,omitempty"`
	// The name of the instance group.
	//
	// example:
	//
	// defaultInstanceGroup
	InstanceGroupName *string `json:"InstanceGroupName,omitempty" xml:"InstanceGroupName,omitempty"`
	// The specifications of the instance group.
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
	// The status of the instance group.
	//
	// example:
	//
	// RUNNING
	InstanceGroupStatus *string `json:"InstanceGroupStatus,omitempty" xml:"InstanceGroupStatus,omitempty"`
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
	// The number of instances in the instance group.
	//
	// example:
	//
	// 10
	NumberOfInstances *string `json:"NumberOfInstances,omitempty" xml:"NumberOfInstances,omitempty"`
	// The ID of the network.
	//
	// example:
	//
	// cn-shanghai+dir-030598****
	OfficeSiteId *string `json:"OfficeSiteId,omitempty" xml:"OfficeSiteId,omitempty"`
	// The ID of the policy.
	//
	// example:
	//
	// pg-c6n38xucps8kl****
	PolicyGroupId *string `json:"PolicyGroupId,omitempty" xml:"PolicyGroupId,omitempty"`
	// The ID of the region.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The rendering mode of the instance group.
	//
	// Valid values:
	//
	// 	- GPURemote: GPU remote rendering.
	//
	// 	- CPU: CPU rendering.
	//
	// 	- GPUocal: GPU local rendering.
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
	// The version of the operating system.
	//
	// example:
	//
	// Android 12
	SystemVersion *string `json:"SystemVersion,omitempty" xml:"SystemVersion,omitempty"`
	// The ID of the vSwitch.
	//
	// example:
	//
	// vsw-t4n0yqs009ho024wt****
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
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

func (s *DescribeAndroidInstanceGroupsResponseBodyInstanceGroupModel) GetNumberOfInstances() *string {
	return s.NumberOfInstances
}

func (s *DescribeAndroidInstanceGroupsResponseBodyInstanceGroupModel) GetOfficeSiteId() *string {
	return s.OfficeSiteId
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

func (s *DescribeAndroidInstanceGroupsResponseBodyInstanceGroupModel) GetVSwitchId() *string {
	return s.VSwitchId
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

func (s *DescribeAndroidInstanceGroupsResponseBodyInstanceGroupModel) SetNumberOfInstances(v string) *DescribeAndroidInstanceGroupsResponseBodyInstanceGroupModel {
	s.NumberOfInstances = &v
	return s
}

func (s *DescribeAndroidInstanceGroupsResponseBodyInstanceGroupModel) SetOfficeSiteId(v string) *DescribeAndroidInstanceGroupsResponseBodyInstanceGroupModel {
	s.OfficeSiteId = &v
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

func (s *DescribeAndroidInstanceGroupsResponseBodyInstanceGroupModel) SetVSwitchId(v string) *DescribeAndroidInstanceGroupsResponseBodyInstanceGroupModel {
	s.VSwitchId = &v
	return s
}

func (s *DescribeAndroidInstanceGroupsResponseBodyInstanceGroupModel) Validate() error {
	return dara.Validate(s)
}

type DescribeAndroidInstanceGroupsResponseBodyInstanceGroupModelDisks struct {
	// The size of the disk. Unit: GB.
	//
	// example:
	//
	// 32
	DiskSize *int32 `json:"DiskSize,omitempty" xml:"DiskSize,omitempty"`
	// The type of the disk.
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
