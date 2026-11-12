// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListInstancesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetInstances(v []*ListInstancesResponseBodyInstances) *ListInstancesResponseBody
	GetInstances() []*ListInstancesResponseBodyInstances
	SetPageNumber(v int32) *ListInstancesResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *ListInstancesResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *ListInstancesResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *ListInstancesResponseBody
	GetTotalCount() *int32
}

type ListInstancesResponseBody struct {
	// The array of simple application server information.
	Instances []*ListInstancesResponseBodyInstances `json:"Instances,omitempty" xml:"Instances,omitempty" type:"Repeated"`
	// The page number of the simple application server list.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The settings for the number of entries per page for paging in a paged query.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 20758A-585D-4A41-A9B2-28DA8F4F****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of instances returned.
	//
	// example:
	//
	// 1
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListInstancesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListInstancesResponseBody) GoString() string {
	return s.String()
}

func (s *ListInstancesResponseBody) GetInstances() []*ListInstancesResponseBodyInstances {
	return s.Instances
}

func (s *ListInstancesResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListInstancesResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListInstancesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListInstancesResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListInstancesResponseBody) SetInstances(v []*ListInstancesResponseBodyInstances) *ListInstancesResponseBody {
	s.Instances = v
	return s
}

func (s *ListInstancesResponseBody) SetPageNumber(v int32) *ListInstancesResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListInstancesResponseBody) SetPageSize(v int32) *ListInstancesResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListInstancesResponseBody) SetRequestId(v string) *ListInstancesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListInstancesResponseBody) SetTotalCount(v int32) *ListInstancesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListInstancesResponseBody) Validate() error {
	if s.Instances != nil {
		for _, item := range s.Instances {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListInstancesResponseBodyInstances struct {
	// The server status. Valid values:
	//
	// - Normal: Normal.
	//
	// - Expired: Expired.
	//
	// - Overdue: Overdue payment.
	//
	// example:
	//
	// Normal
	BusinessStatus *string `json:"BusinessStatus,omitempty" xml:"BusinessStatus,omitempty"`
	// The billing method.
	//
	// example:
	//
	// PrePaid
	ChargeType *string `json:"ChargeType,omitempty" xml:"ChargeType,omitempty"`
	// Specifies whether the plan is a combination plan.
	//
	// example:
	//
	// false
	Combination *bool `json:"Combination,omitempty" xml:"Combination,omitempty"`
	// The combination plan instance ID.
	//
	// example:
	//
	// com-f6c9a22****45b5b8de68ad608af1ba
	CombinationInstanceId *string `json:"CombinationInstanceId,omitempty" xml:"CombinationInstanceId,omitempty"`
	// The commodity code of the instance.
	//
	// example:
	//
	// swas_intl
	CommodityCode *string `json:"CommodityCode,omitempty" xml:"CommodityCode,omitempty"`
	// The creation time. The time is displayed in UTC in the [ISO 8601](https://help.aliyun.com/document_detail/25696.html) standard in the yyyy-MM-ddTHH:mm:ssZ format.
	//
	// example:
	//
	// 2021-03-08T05:31:06Z
	CreationTime *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	// The DDoS status. Valid values:
	//
	// - Normal: Normal.
	//
	// - BlackHole: Blackhole filtering is activated.
	//
	// - Defense: Traffic is being scrubbed.
	//
	// example:
	//
	// Normal
	DdosStatus *string `json:"DdosStatus,omitempty" xml:"DdosStatus,omitempty"`
	// The reason why the resource is disabled. Valid values:
	//
	// - FINANCIAL: The resource is locked due to an overdue payment.
	//
	// - SECURITY: The resource is locked due to security reasons.
	//
	// - EXPIRED: The instance has expired.
	//
	// example:
	//
	// EXPIRED
	DisableReason *string `json:"DisableReason,omitempty" xml:"DisableReason,omitempty"`
	// The array of cloud disk information for the simple application server.
	Disks []*ListInstancesResponseBodyInstancesDisks `json:"Disks,omitempty" xml:"Disks,omitempty" type:"Repeated"`
	// The expiration time. The time is displayed in UTC in the [ISO 8601](https://help.aliyun.com/document_detail/25696.html) standard in the yyyy-MM-ddTHH:mm:ssZ format.
	//
	// example:
	//
	// 2021-05-08T16:00:00Z
	ExpiredTime *string `json:"ExpiredTime,omitempty" xml:"ExpiredTime,omitempty"`
	// The image information.
	Image *ListInstancesResponseBodyInstancesImage `json:"Image,omitempty" xml:"Image,omitempty" type:"Struct"`
	// The image ID.
	//
	// example:
	//
	// fe9c66133a9d4688872869726b52****
	ImageId *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	// The internal IP address.
	//
	// example:
	//
	// 172.26.XX.XX
	InnerIpAddress *string `json:"InnerIpAddress,omitempty" xml:"InnerIpAddress,omitempty"`
	// The instance ID.
	//
	// example:
	//
	// 2ad1ae67295445f598017499dc****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The server name.
	//
	// example:
	//
	// test-InstanceName
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// The network attributes.
	NetworkAttributes []*ListInstancesResponseBodyInstancesNetworkAttributes `json:"NetworkAttributes,omitempty" xml:"NetworkAttributes,omitempty" type:"Repeated"`
	// The plan ID.
	//
	// example:
	//
	// swas.s2.c2m2s50b4t08
	PlanId *string `json:"PlanId,omitempty" xml:"PlanId,omitempty"`
	// The plan type. Valid values:
	//
	// - NORMAL: general-purpose.
	//
	// - CAPACITY: storage-optimized.
	//
	// - PREVIOUS: previous generation.
	//
	// - MULTI_IP: multi-public-IP.
	//
	// - INTERNATIONAL: international.
	//
	// > For more information, see [Instance families](https://help.aliyun.com/document_detail/2861773.html).
	//
	// example:
	//
	// NORMAL
	PlanType *string `json:"PlanType,omitempty" xml:"PlanType,omitempty"`
	// The public IP address.
	//
	// example:
	//
	// 42.1.XX.XX
	PublicIpAddress *string `json:"PublicIpAddress,omitempty" xml:"PublicIpAddress,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group to which the instance belongs.
	//
	// example:
	//
	// rg-aekz7jmhg5s****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The resource specifications.
	ResourceSpec *ListInstancesResponseBodyInstancesResourceSpec `json:"ResourceSpec,omitempty" xml:"ResourceSpec,omitempty" type:"Struct"`
	// The instance status. Valid values:
	//
	// - Pending: The instance is being prepared.
	//
	// - Starting: The instance is starting.
	//
	// - Running: The instance is running.
	//
	// - Stopping: The instance is stopping.
	//
	// - Stopped: The instance is stopped.
	//
	// - Resetting: The instance is being reset.
	//
	// - Upgrading: The instance is being upgraded.
	//
	// - Disabled: The instance is unavailable.
	//
	// example:
	//
	// Running
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The tags of the instance.
	Tags []*ListInstancesResponseBodyInstancesTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	// The UUID of the instance.
	//
	// example:
	//
	// 41f30524-5df7-49c9-9c6e-32****489001
	Uuid *string `json:"Uuid,omitempty" xml:"Uuid,omitempty"`
}

func (s ListInstancesResponseBodyInstances) String() string {
	return dara.Prettify(s)
}

func (s ListInstancesResponseBodyInstances) GoString() string {
	return s.String()
}

func (s *ListInstancesResponseBodyInstances) GetBusinessStatus() *string {
	return s.BusinessStatus
}

func (s *ListInstancesResponseBodyInstances) GetChargeType() *string {
	return s.ChargeType
}

func (s *ListInstancesResponseBodyInstances) GetCombination() *bool {
	return s.Combination
}

func (s *ListInstancesResponseBodyInstances) GetCombinationInstanceId() *string {
	return s.CombinationInstanceId
}

func (s *ListInstancesResponseBodyInstances) GetCommodityCode() *string {
	return s.CommodityCode
}

func (s *ListInstancesResponseBodyInstances) GetCreationTime() *string {
	return s.CreationTime
}

func (s *ListInstancesResponseBodyInstances) GetDdosStatus() *string {
	return s.DdosStatus
}

func (s *ListInstancesResponseBodyInstances) GetDisableReason() *string {
	return s.DisableReason
}

func (s *ListInstancesResponseBodyInstances) GetDisks() []*ListInstancesResponseBodyInstancesDisks {
	return s.Disks
}

func (s *ListInstancesResponseBodyInstances) GetExpiredTime() *string {
	return s.ExpiredTime
}

func (s *ListInstancesResponseBodyInstances) GetImage() *ListInstancesResponseBodyInstancesImage {
	return s.Image
}

func (s *ListInstancesResponseBodyInstances) GetImageId() *string {
	return s.ImageId
}

func (s *ListInstancesResponseBodyInstances) GetInnerIpAddress() *string {
	return s.InnerIpAddress
}

func (s *ListInstancesResponseBodyInstances) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListInstancesResponseBodyInstances) GetInstanceName() *string {
	return s.InstanceName
}

func (s *ListInstancesResponseBodyInstances) GetNetworkAttributes() []*ListInstancesResponseBodyInstancesNetworkAttributes {
	return s.NetworkAttributes
}

func (s *ListInstancesResponseBodyInstances) GetPlanId() *string {
	return s.PlanId
}

func (s *ListInstancesResponseBodyInstances) GetPlanType() *string {
	return s.PlanType
}

func (s *ListInstancesResponseBodyInstances) GetPublicIpAddress() *string {
	return s.PublicIpAddress
}

func (s *ListInstancesResponseBodyInstances) GetRegionId() *string {
	return s.RegionId
}

func (s *ListInstancesResponseBodyInstances) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *ListInstancesResponseBodyInstances) GetResourceSpec() *ListInstancesResponseBodyInstancesResourceSpec {
	return s.ResourceSpec
}

func (s *ListInstancesResponseBodyInstances) GetStatus() *string {
	return s.Status
}

func (s *ListInstancesResponseBodyInstances) GetTags() []*ListInstancesResponseBodyInstancesTags {
	return s.Tags
}

func (s *ListInstancesResponseBodyInstances) GetUuid() *string {
	return s.Uuid
}

func (s *ListInstancesResponseBodyInstances) SetBusinessStatus(v string) *ListInstancesResponseBodyInstances {
	s.BusinessStatus = &v
	return s
}

func (s *ListInstancesResponseBodyInstances) SetChargeType(v string) *ListInstancesResponseBodyInstances {
	s.ChargeType = &v
	return s
}

func (s *ListInstancesResponseBodyInstances) SetCombination(v bool) *ListInstancesResponseBodyInstances {
	s.Combination = &v
	return s
}

func (s *ListInstancesResponseBodyInstances) SetCombinationInstanceId(v string) *ListInstancesResponseBodyInstances {
	s.CombinationInstanceId = &v
	return s
}

func (s *ListInstancesResponseBodyInstances) SetCommodityCode(v string) *ListInstancesResponseBodyInstances {
	s.CommodityCode = &v
	return s
}

func (s *ListInstancesResponseBodyInstances) SetCreationTime(v string) *ListInstancesResponseBodyInstances {
	s.CreationTime = &v
	return s
}

func (s *ListInstancesResponseBodyInstances) SetDdosStatus(v string) *ListInstancesResponseBodyInstances {
	s.DdosStatus = &v
	return s
}

func (s *ListInstancesResponseBodyInstances) SetDisableReason(v string) *ListInstancesResponseBodyInstances {
	s.DisableReason = &v
	return s
}

func (s *ListInstancesResponseBodyInstances) SetDisks(v []*ListInstancesResponseBodyInstancesDisks) *ListInstancesResponseBodyInstances {
	s.Disks = v
	return s
}

func (s *ListInstancesResponseBodyInstances) SetExpiredTime(v string) *ListInstancesResponseBodyInstances {
	s.ExpiredTime = &v
	return s
}

func (s *ListInstancesResponseBodyInstances) SetImage(v *ListInstancesResponseBodyInstancesImage) *ListInstancesResponseBodyInstances {
	s.Image = v
	return s
}

func (s *ListInstancesResponseBodyInstances) SetImageId(v string) *ListInstancesResponseBodyInstances {
	s.ImageId = &v
	return s
}

func (s *ListInstancesResponseBodyInstances) SetInnerIpAddress(v string) *ListInstancesResponseBodyInstances {
	s.InnerIpAddress = &v
	return s
}

func (s *ListInstancesResponseBodyInstances) SetInstanceId(v string) *ListInstancesResponseBodyInstances {
	s.InstanceId = &v
	return s
}

func (s *ListInstancesResponseBodyInstances) SetInstanceName(v string) *ListInstancesResponseBodyInstances {
	s.InstanceName = &v
	return s
}

func (s *ListInstancesResponseBodyInstances) SetNetworkAttributes(v []*ListInstancesResponseBodyInstancesNetworkAttributes) *ListInstancesResponseBodyInstances {
	s.NetworkAttributes = v
	return s
}

func (s *ListInstancesResponseBodyInstances) SetPlanId(v string) *ListInstancesResponseBodyInstances {
	s.PlanId = &v
	return s
}

func (s *ListInstancesResponseBodyInstances) SetPlanType(v string) *ListInstancesResponseBodyInstances {
	s.PlanType = &v
	return s
}

func (s *ListInstancesResponseBodyInstances) SetPublicIpAddress(v string) *ListInstancesResponseBodyInstances {
	s.PublicIpAddress = &v
	return s
}

func (s *ListInstancesResponseBodyInstances) SetRegionId(v string) *ListInstancesResponseBodyInstances {
	s.RegionId = &v
	return s
}

func (s *ListInstancesResponseBodyInstances) SetResourceGroupId(v string) *ListInstancesResponseBodyInstances {
	s.ResourceGroupId = &v
	return s
}

func (s *ListInstancesResponseBodyInstances) SetResourceSpec(v *ListInstancesResponseBodyInstancesResourceSpec) *ListInstancesResponseBodyInstances {
	s.ResourceSpec = v
	return s
}

func (s *ListInstancesResponseBodyInstances) SetStatus(v string) *ListInstancesResponseBodyInstances {
	s.Status = &v
	return s
}

func (s *ListInstancesResponseBodyInstances) SetTags(v []*ListInstancesResponseBodyInstancesTags) *ListInstancesResponseBodyInstances {
	s.Tags = v
	return s
}

func (s *ListInstancesResponseBodyInstances) SetUuid(v string) *ListInstancesResponseBodyInstances {
	s.Uuid = &v
	return s
}

func (s *ListInstancesResponseBodyInstances) Validate() error {
	if s.Disks != nil {
		for _, item := range s.Disks {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Image != nil {
		if err := s.Image.Validate(); err != nil {
			return err
		}
	}
	if s.NetworkAttributes != nil {
		for _, item := range s.NetworkAttributes {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.ResourceSpec != nil {
		if err := s.ResourceSpec.Validate(); err != nil {
			return err
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

type ListInstancesResponseBodyInstancesDisks struct {
	// The cloud disk category. Valid values:
	//
	// - ESSD: ESSD at performance level (PL) 0.
	//
	// - SSD: Standard SSD.
	//
	// - CLOUD_EFFICIENCY: Ultra disk.
	//
	// example:
	//
	// ESSD
	Category *string `json:"Category,omitempty" xml:"Category,omitempty"`
	// The creation time of the instance. The time is displayed in UTC in the [ISO 8601](https://help.aliyun.com/document_detail/25696.html) standard in the yyyy-MM-ddTHH:mm:ssZ format.
	//
	// example:
	//
	// 2023-02-24T02:20:10Z
	CreationTime *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	// The device name of the cloud disk on the server.
	//
	// example:
	//
	// /dev/xvda
	Device *string `json:"Device,omitempty" xml:"Device,omitempty"`
	// The billing method of the cloud disk.
	//
	// example:
	//
	// PrePaid
	DiskChargeType *string `json:"DiskChargeType,omitempty" xml:"DiskChargeType,omitempty"`
	// The cloud disk ID.
	//
	// example:
	//
	// d-bp14wq0149cpp2x****
	DiskId *string `json:"DiskId,omitempty" xml:"DiskId,omitempty"`
	// The cloud disk name.
	//
	// example:
	//
	// SystemDisk
	DiskName *string `json:"DiskName,omitempty" xml:"DiskName,omitempty"`
	// The tag information of the cloud disk.
	DiskTags []*ListInstancesResponseBodyInstancesDisksDiskTags `json:"DiskTags,omitempty" xml:"DiskTags,omitempty" type:"Repeated"`
	// The cloud disk type. Valid values:
	//
	// - system: system cloud disk.
	//
	// - data: data cloud disk.
	//
	// example:
	//
	// system
	DiskType *string `json:"DiskType,omitempty" xml:"DiskType,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The remarks of the cloud disk.
	//
	// example:
	//
	// remark
	Remark *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	// The ID of the resource group to which the cloud disk belongs.
	//
	// example:
	//
	// rg-aek2bti7cf7****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The cloud disk size. Unit: GiB.
	//
	// example:
	//
	// 50
	Size *int32 `json:"Size,omitempty" xml:"Size,omitempty"`
	// The cloud disk status. Valid values:
	//
	// - ReIniting: The cloud disk is being initialized.
	//
	// - Creating: The cloud disk is being created.
	//
	// - In_use: The cloud disk is in use.
	//
	// - Available: The cloud disk is unattached.
	//
	// - Attaching: The cloud disk is being mounted.
	//
	// - Detaching: The cloud disk is being unmounted.
	//
	// example:
	//
	// In_use
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ListInstancesResponseBodyInstancesDisks) String() string {
	return dara.Prettify(s)
}

func (s ListInstancesResponseBodyInstancesDisks) GoString() string {
	return s.String()
}

func (s *ListInstancesResponseBodyInstancesDisks) GetCategory() *string {
	return s.Category
}

func (s *ListInstancesResponseBodyInstancesDisks) GetCreationTime() *string {
	return s.CreationTime
}

func (s *ListInstancesResponseBodyInstancesDisks) GetDevice() *string {
	return s.Device
}

func (s *ListInstancesResponseBodyInstancesDisks) GetDiskChargeType() *string {
	return s.DiskChargeType
}

func (s *ListInstancesResponseBodyInstancesDisks) GetDiskId() *string {
	return s.DiskId
}

func (s *ListInstancesResponseBodyInstancesDisks) GetDiskName() *string {
	return s.DiskName
}

func (s *ListInstancesResponseBodyInstancesDisks) GetDiskTags() []*ListInstancesResponseBodyInstancesDisksDiskTags {
	return s.DiskTags
}

func (s *ListInstancesResponseBodyInstancesDisks) GetDiskType() *string {
	return s.DiskType
}

func (s *ListInstancesResponseBodyInstancesDisks) GetRegionId() *string {
	return s.RegionId
}

func (s *ListInstancesResponseBodyInstancesDisks) GetRemark() *string {
	return s.Remark
}

func (s *ListInstancesResponseBodyInstancesDisks) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *ListInstancesResponseBodyInstancesDisks) GetSize() *int32 {
	return s.Size
}

func (s *ListInstancesResponseBodyInstancesDisks) GetStatus() *string {
	return s.Status
}

func (s *ListInstancesResponseBodyInstancesDisks) SetCategory(v string) *ListInstancesResponseBodyInstancesDisks {
	s.Category = &v
	return s
}

func (s *ListInstancesResponseBodyInstancesDisks) SetCreationTime(v string) *ListInstancesResponseBodyInstancesDisks {
	s.CreationTime = &v
	return s
}

func (s *ListInstancesResponseBodyInstancesDisks) SetDevice(v string) *ListInstancesResponseBodyInstancesDisks {
	s.Device = &v
	return s
}

func (s *ListInstancesResponseBodyInstancesDisks) SetDiskChargeType(v string) *ListInstancesResponseBodyInstancesDisks {
	s.DiskChargeType = &v
	return s
}

func (s *ListInstancesResponseBodyInstancesDisks) SetDiskId(v string) *ListInstancesResponseBodyInstancesDisks {
	s.DiskId = &v
	return s
}

func (s *ListInstancesResponseBodyInstancesDisks) SetDiskName(v string) *ListInstancesResponseBodyInstancesDisks {
	s.DiskName = &v
	return s
}

func (s *ListInstancesResponseBodyInstancesDisks) SetDiskTags(v []*ListInstancesResponseBodyInstancesDisksDiskTags) *ListInstancesResponseBodyInstancesDisks {
	s.DiskTags = v
	return s
}

func (s *ListInstancesResponseBodyInstancesDisks) SetDiskType(v string) *ListInstancesResponseBodyInstancesDisks {
	s.DiskType = &v
	return s
}

func (s *ListInstancesResponseBodyInstancesDisks) SetRegionId(v string) *ListInstancesResponseBodyInstancesDisks {
	s.RegionId = &v
	return s
}

func (s *ListInstancesResponseBodyInstancesDisks) SetRemark(v string) *ListInstancesResponseBodyInstancesDisks {
	s.Remark = &v
	return s
}

func (s *ListInstancesResponseBodyInstancesDisks) SetResourceGroupId(v string) *ListInstancesResponseBodyInstancesDisks {
	s.ResourceGroupId = &v
	return s
}

func (s *ListInstancesResponseBodyInstancesDisks) SetSize(v int32) *ListInstancesResponseBodyInstancesDisks {
	s.Size = &v
	return s
}

func (s *ListInstancesResponseBodyInstancesDisks) SetStatus(v string) *ListInstancesResponseBodyInstancesDisks {
	s.Status = &v
	return s
}

func (s *ListInstancesResponseBodyInstancesDisks) Validate() error {
	if s.DiskTags != nil {
		for _, item := range s.DiskTags {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListInstancesResponseBodyInstancesDisksDiskTags struct {
	// The tag key.
	//
	// example:
	//
	// TestKey
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value.
	//
	// example:
	//
	// TestValue
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListInstancesResponseBodyInstancesDisksDiskTags) String() string {
	return dara.Prettify(s)
}

func (s ListInstancesResponseBodyInstancesDisksDiskTags) GoString() string {
	return s.String()
}

func (s *ListInstancesResponseBodyInstancesDisksDiskTags) GetKey() *string {
	return s.Key
}

func (s *ListInstancesResponseBodyInstancesDisksDiskTags) GetValue() *string {
	return s.Value
}

func (s *ListInstancesResponseBodyInstancesDisksDiskTags) SetKey(v string) *ListInstancesResponseBodyInstancesDisksDiskTags {
	s.Key = &v
	return s
}

func (s *ListInstancesResponseBodyInstancesDisksDiskTags) SetValue(v string) *ListInstancesResponseBodyInstancesDisksDiskTags {
	s.Value = &v
	return s
}

func (s *ListInstancesResponseBodyInstancesDisksDiskTags) Validate() error {
	return dara.Validate(s)
}

type ListInstancesResponseBodyInstancesImage struct {
	// The image provider.
	//
	// example:
	//
	// https://selfs****e.console.aliyun.com/ticket/createIndex
	ImageContact *string `json:"ImageContact,omitempty" xml:"ImageContact,omitempty"`
	// The URL of the image icon.
	//
	// example:
	//
	// https://img.alicdn.com/imgextra/i3/O****1vdh9651ReKqWNMI2I_!!6000000002136****-24-24.svg
	ImageIconUrl *string `json:"ImageIconUrl,omitempty" xml:"ImageIconUrl,omitempty"`
	// The image name.
	//
	// example:
	//
	// test-custom-1686536882356
	ImageName *string `json:"ImageName,omitempty" xml:"ImageName,omitempty"`
	// The image type. Valid values:
	//
	// - system: system image.
	//
	// - app: application image.
	//
	// - custom: custom image.
	//
	// example:
	//
	// system
	ImageType *string `json:"ImageType,omitempty" xml:"ImageType,omitempty"`
	// The image version.
	//
	// example:
	//
	// V3.5
	ImageVersion *string `json:"ImageVersion,omitempty" xml:"ImageVersion,omitempty"`
	// The operating system type.
	//
	// example:
	//
	// windows
	OsType *string `json:"OsType,omitempty" xml:"OsType,omitempty"`
}

func (s ListInstancesResponseBodyInstancesImage) String() string {
	return dara.Prettify(s)
}

func (s ListInstancesResponseBodyInstancesImage) GoString() string {
	return s.String()
}

func (s *ListInstancesResponseBodyInstancesImage) GetImageContact() *string {
	return s.ImageContact
}

func (s *ListInstancesResponseBodyInstancesImage) GetImageIconUrl() *string {
	return s.ImageIconUrl
}

func (s *ListInstancesResponseBodyInstancesImage) GetImageName() *string {
	return s.ImageName
}

func (s *ListInstancesResponseBodyInstancesImage) GetImageType() *string {
	return s.ImageType
}

func (s *ListInstancesResponseBodyInstancesImage) GetImageVersion() *string {
	return s.ImageVersion
}

func (s *ListInstancesResponseBodyInstancesImage) GetOsType() *string {
	return s.OsType
}

func (s *ListInstancesResponseBodyInstancesImage) SetImageContact(v string) *ListInstancesResponseBodyInstancesImage {
	s.ImageContact = &v
	return s
}

func (s *ListInstancesResponseBodyInstancesImage) SetImageIconUrl(v string) *ListInstancesResponseBodyInstancesImage {
	s.ImageIconUrl = &v
	return s
}

func (s *ListInstancesResponseBodyInstancesImage) SetImageName(v string) *ListInstancesResponseBodyInstancesImage {
	s.ImageName = &v
	return s
}

func (s *ListInstancesResponseBodyInstancesImage) SetImageType(v string) *ListInstancesResponseBodyInstancesImage {
	s.ImageType = &v
	return s
}

func (s *ListInstancesResponseBodyInstancesImage) SetImageVersion(v string) *ListInstancesResponseBodyInstancesImage {
	s.ImageVersion = &v
	return s
}

func (s *ListInstancesResponseBodyInstancesImage) SetOsType(v string) *ListInstancesResponseBodyInstancesImage {
	s.OsType = &v
	return s
}

func (s *ListInstancesResponseBodyInstancesImage) Validate() error {
	return dara.Validate(s)
}

type ListInstancesResponseBodyInstancesNetworkAttributes struct {
	// The peak public bandwidth. Unit: Mbit/s.
	//
	// example:
	//
	// 200
	PeakBandwidth *int32 `json:"PeakBandwidth,omitempty" xml:"PeakBandwidth,omitempty"`
	// The private IP address.
	//
	// example:
	//
	// 72.25.XX.XX
	PrivateIpAddress *string `json:"PrivateIpAddress,omitempty" xml:"PrivateIpAddress,omitempty"`
	// The public IP address.
	//
	// example:
	//
	// 116.62.XX.XX
	PublicIpAddress *string `json:"PublicIpAddress,omitempty" xml:"PublicIpAddress,omitempty"`
	// The DDoS status of the public IP address.
	//
	// example:
	//
	// normal
	PublicIpDdosStatus *string `json:"PublicIpDdosStatus,omitempty" xml:"PublicIpDdosStatus,omitempty"`
}

func (s ListInstancesResponseBodyInstancesNetworkAttributes) String() string {
	return dara.Prettify(s)
}

func (s ListInstancesResponseBodyInstancesNetworkAttributes) GoString() string {
	return s.String()
}

func (s *ListInstancesResponseBodyInstancesNetworkAttributes) GetPeakBandwidth() *int32 {
	return s.PeakBandwidth
}

func (s *ListInstancesResponseBodyInstancesNetworkAttributes) GetPrivateIpAddress() *string {
	return s.PrivateIpAddress
}

func (s *ListInstancesResponseBodyInstancesNetworkAttributes) GetPublicIpAddress() *string {
	return s.PublicIpAddress
}

func (s *ListInstancesResponseBodyInstancesNetworkAttributes) GetPublicIpDdosStatus() *string {
	return s.PublicIpDdosStatus
}

func (s *ListInstancesResponseBodyInstancesNetworkAttributes) SetPeakBandwidth(v int32) *ListInstancesResponseBodyInstancesNetworkAttributes {
	s.PeakBandwidth = &v
	return s
}

func (s *ListInstancesResponseBodyInstancesNetworkAttributes) SetPrivateIpAddress(v string) *ListInstancesResponseBodyInstancesNetworkAttributes {
	s.PrivateIpAddress = &v
	return s
}

func (s *ListInstancesResponseBodyInstancesNetworkAttributes) SetPublicIpAddress(v string) *ListInstancesResponseBodyInstancesNetworkAttributes {
	s.PublicIpAddress = &v
	return s
}

func (s *ListInstancesResponseBodyInstancesNetworkAttributes) SetPublicIpDdosStatus(v string) *ListInstancesResponseBodyInstancesNetworkAttributes {
	s.PublicIpDdosStatus = &v
	return s
}

func (s *ListInstancesResponseBodyInstancesNetworkAttributes) Validate() error {
	return dara.Validate(s)
}

type ListInstancesResponseBodyInstancesResourceSpec struct {
	// The bandwidth. Unit: Mbit/s.
	//
	// example:
	//
	// 200
	Bandwidth *int32 `json:"Bandwidth,omitempty" xml:"Bandwidth,omitempty"`
	// The number of vCPUs.
	//
	// example:
	//
	// 2
	Cpu *int32 `json:"Cpu,omitempty" xml:"Cpu,omitempty"`
	// The cloud disk category. Valid values:
	//
	// - ESSD: ESSD at performance level (PL) 0.
	//
	// - SSD: Standard SSD.
	//
	// - CLOUD_EFFICIENCY: Ultra disk.
	//
	// example:
	//
	// ESSD
	DiskCategory *string `json:"DiskCategory,omitempty" xml:"DiskCategory,omitempty"`
	// The cloud disk size. Unit: GiB.
	//
	// example:
	//
	// 60
	DiskSize *int32 `json:"DiskSize,omitempty" xml:"DiskSize,omitempty"`
	// The data transfer quota.
	//
	// - A value of 0 indicates a bandwidth-based instance.
	//
	// - A non-zero value indicates a data transfer plan-based instance.
	//
	// example:
	//
	// 818
	Flow *float64 `json:"Flow,omitempty" xml:"Flow,omitempty"`
	// The memory size. Unit: GiB.
	//
	// example:
	//
	// 2
	Memory *float64 `json:"Memory,omitempty" xml:"Memory,omitempty"`
}

func (s ListInstancesResponseBodyInstancesResourceSpec) String() string {
	return dara.Prettify(s)
}

func (s ListInstancesResponseBodyInstancesResourceSpec) GoString() string {
	return s.String()
}

func (s *ListInstancesResponseBodyInstancesResourceSpec) GetBandwidth() *int32 {
	return s.Bandwidth
}

func (s *ListInstancesResponseBodyInstancesResourceSpec) GetCpu() *int32 {
	return s.Cpu
}

func (s *ListInstancesResponseBodyInstancesResourceSpec) GetDiskCategory() *string {
	return s.DiskCategory
}

func (s *ListInstancesResponseBodyInstancesResourceSpec) GetDiskSize() *int32 {
	return s.DiskSize
}

func (s *ListInstancesResponseBodyInstancesResourceSpec) GetFlow() *float64 {
	return s.Flow
}

func (s *ListInstancesResponseBodyInstancesResourceSpec) GetMemory() *float64 {
	return s.Memory
}

func (s *ListInstancesResponseBodyInstancesResourceSpec) SetBandwidth(v int32) *ListInstancesResponseBodyInstancesResourceSpec {
	s.Bandwidth = &v
	return s
}

func (s *ListInstancesResponseBodyInstancesResourceSpec) SetCpu(v int32) *ListInstancesResponseBodyInstancesResourceSpec {
	s.Cpu = &v
	return s
}

func (s *ListInstancesResponseBodyInstancesResourceSpec) SetDiskCategory(v string) *ListInstancesResponseBodyInstancesResourceSpec {
	s.DiskCategory = &v
	return s
}

func (s *ListInstancesResponseBodyInstancesResourceSpec) SetDiskSize(v int32) *ListInstancesResponseBodyInstancesResourceSpec {
	s.DiskSize = &v
	return s
}

func (s *ListInstancesResponseBodyInstancesResourceSpec) SetFlow(v float64) *ListInstancesResponseBodyInstancesResourceSpec {
	s.Flow = &v
	return s
}

func (s *ListInstancesResponseBodyInstancesResourceSpec) SetMemory(v float64) *ListInstancesResponseBodyInstancesResourceSpec {
	s.Memory = &v
	return s
}

func (s *ListInstancesResponseBodyInstancesResourceSpec) Validate() error {
	return dara.Validate(s)
}

type ListInstancesResponseBodyInstancesTags struct {
	// The tag key of the instance.
	//
	// example:
	//
	// TestKey
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value of the instance.
	//
	// example:
	//
	// TestValue
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListInstancesResponseBodyInstancesTags) String() string {
	return dara.Prettify(s)
}

func (s ListInstancesResponseBodyInstancesTags) GoString() string {
	return s.String()
}

func (s *ListInstancesResponseBodyInstancesTags) GetKey() *string {
	return s.Key
}

func (s *ListInstancesResponseBodyInstancesTags) GetValue() *string {
	return s.Value
}

func (s *ListInstancesResponseBodyInstancesTags) SetKey(v string) *ListInstancesResponseBodyInstancesTags {
	s.Key = &v
	return s
}

func (s *ListInstancesResponseBodyInstancesTags) SetValue(v string) *ListInstancesResponseBodyInstancesTags {
	s.Value = &v
	return s
}

func (s *ListInstancesResponseBodyInstancesTags) Validate() error {
	return dara.Validate(s)
}
