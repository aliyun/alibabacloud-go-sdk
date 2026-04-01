// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeInstancesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *DescribeInstancesResponseBody
	GetCode() *int32
	SetInstances(v *DescribeInstancesResponseBodyInstances) *DescribeInstancesResponseBody
	GetInstances() *DescribeInstancesResponseBodyInstances
	SetPageNumber(v int32) *DescribeInstancesResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeInstancesResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *DescribeInstancesResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *DescribeInstancesResponseBody
	GetTotalCount() *int32
}

type DescribeInstancesResponseBody struct {
	// The returned service code. 0 indicates that the request was successful.
	//
	// example:
	//
	// 0
	Code      *int32                                  `json:"Code,omitempty" xml:"Code,omitempty"`
	Instances *DescribeInstancesResponseBodyInstances `json:"Instances,omitempty" xml:"Instances,omitempty" type:"Struct"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 60
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3C83E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 2
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeInstancesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstancesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeInstancesResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *DescribeInstancesResponseBody) GetInstances() *DescribeInstancesResponseBodyInstances {
	return s.Instances
}

func (s *DescribeInstancesResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeInstancesResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeInstancesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeInstancesResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeInstancesResponseBody) SetCode(v int32) *DescribeInstancesResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeInstancesResponseBody) SetInstances(v *DescribeInstancesResponseBodyInstances) *DescribeInstancesResponseBody {
	s.Instances = v
	return s
}

func (s *DescribeInstancesResponseBody) SetPageNumber(v int32) *DescribeInstancesResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeInstancesResponseBody) SetPageSize(v int32) *DescribeInstancesResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeInstancesResponseBody) SetRequestId(v string) *DescribeInstancesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeInstancesResponseBody) SetTotalCount(v int32) *DescribeInstancesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeInstancesResponseBody) Validate() error {
	if s.Instances != nil {
		if err := s.Instances.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeInstancesResponseBodyInstances struct {
	Instance []*DescribeInstancesResponseBodyInstancesInstance `json:"Instance,omitempty" xml:"Instance,omitempty" type:"Repeated"`
}

func (s DescribeInstancesResponseBodyInstances) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstancesResponseBodyInstances) GoString() string {
	return s.String()
}

func (s *DescribeInstancesResponseBodyInstances) GetInstance() []*DescribeInstancesResponseBodyInstancesInstance {
	return s.Instance
}

func (s *DescribeInstancesResponseBodyInstances) SetInstance(v []*DescribeInstancesResponseBodyInstancesInstance) *DescribeInstancesResponseBodyInstances {
	s.Instance = v
	return s
}

func (s *DescribeInstancesResponseBodyInstances) Validate() error {
	if s.Instance != nil {
		for _, item := range s.Instance {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeInstancesResponseBodyInstancesInstance struct {
	AutoReleaseTime         *string                                                           `json:"AutoReleaseTime,omitempty" xml:"AutoReleaseTime,omitempty"`
	Cpu                     *string                                                           `json:"Cpu,omitempty" xml:"Cpu,omitempty"`
	CreationTime            *string                                                           `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	DataDisk                *DescribeInstancesResponseBodyInstancesInstanceDataDisk           `json:"DataDisk,omitempty" xml:"DataDisk,omitempty" type:"Struct"`
	DeletionProtection      *bool                                                             `json:"DeletionProtection,omitempty" xml:"DeletionProtection,omitempty"`
	Disk                    *int32                                                            `json:"Disk,omitempty" xml:"Disk,omitempty"`
	EnsRegionId             *string                                                           `json:"EnsRegionId,omitempty" xml:"EnsRegionId,omitempty"`
	ExpiredTime             *string                                                           `json:"ExpiredTime,omitempty" xml:"ExpiredTime,omitempty"`
	HostName                *string                                                           `json:"HostName,omitempty" xml:"HostName,omitempty"`
	ImageId                 *string                                                           `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	InnerIpAddress          *DescribeInstancesResponseBodyInstancesInstanceInnerIpAddress     `json:"InnerIpAddress,omitempty" xml:"InnerIpAddress,omitempty" type:"Struct"`
	InstanceId              *string                                                           `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	InstanceName            *string                                                           `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	InstanceResourceType    *string                                                           `json:"InstanceResourceType,omitempty" xml:"InstanceResourceType,omitempty"`
	InstanceTypeFamily      *string                                                           `json:"InstanceTypeFamily,omitempty" xml:"InstanceTypeFamily,omitempty"`
	InternetMaxBandwidthIn  *int32                                                            `json:"InternetMaxBandwidthIn,omitempty" xml:"InternetMaxBandwidthIn,omitempty"`
	InternetMaxBandwidthOut *int32                                                            `json:"InternetMaxBandwidthOut,omitempty" xml:"InternetMaxBandwidthOut,omitempty"`
	KeyPairName             *string                                                           `json:"KeyPairName,omitempty" xml:"KeyPairName,omitempty"`
	Memory                  *int32                                                            `json:"Memory,omitempty" xml:"Memory,omitempty"`
	NetworkAttributes       *DescribeInstancesResponseBodyInstancesInstanceNetworkAttributes  `json:"NetworkAttributes,omitempty" xml:"NetworkAttributes,omitempty" type:"Struct"`
	NetworkInterfaces       *DescribeInstancesResponseBodyInstancesInstanceNetworkInterfaces  `json:"NetworkInterfaces,omitempty" xml:"NetworkInterfaces,omitempty" type:"Struct"`
	OSName                  *string                                                           `json:"OSName,omitempty" xml:"OSName,omitempty"`
	PrivateIpAddresses      *DescribeInstancesResponseBodyInstancesInstancePrivateIpAddresses `json:"PrivateIpAddresses,omitempty" xml:"PrivateIpAddresses,omitempty" type:"Struct"`
	PublicIpAddress         *DescribeInstancesResponseBodyInstancesInstancePublicIpAddress    `json:"PublicIpAddress,omitempty" xml:"PublicIpAddress,omitempty" type:"Struct"`
	PublicIpAddresses       *DescribeInstancesResponseBodyInstancesInstancePublicIpAddresses  `json:"PublicIpAddresses,omitempty" xml:"PublicIpAddresses,omitempty" type:"Struct"`
	SecurityGroupIds        *DescribeInstancesResponseBodyInstancesInstanceSecurityGroupIds   `json:"SecurityGroupIds,omitempty" xml:"SecurityGroupIds,omitempty" type:"Struct"`
	ServiceStatus           *string                                                           `json:"ServiceStatus,omitempty" xml:"ServiceStatus,omitempty"`
	SpecName                *string                                                           `json:"SpecName,omitempty" xml:"SpecName,omitempty"`
	SpotStrategy            *string                                                           `json:"SpotStrategy,omitempty" xml:"SpotStrategy,omitempty"`
	Status                  *string                                                           `json:"Status,omitempty" xml:"Status,omitempty"`
	SystemDisk              *DescribeInstancesResponseBodyInstancesInstanceSystemDisk         `json:"SystemDisk,omitempty" xml:"SystemDisk,omitempty" type:"Struct"`
	Tags                    *DescribeInstancesResponseBodyInstancesInstanceTags               `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Struct"`
}

func (s DescribeInstancesResponseBodyInstancesInstance) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstancesResponseBodyInstancesInstance) GoString() string {
	return s.String()
}

func (s *DescribeInstancesResponseBodyInstancesInstance) GetAutoReleaseTime() *string {
	return s.AutoReleaseTime
}

func (s *DescribeInstancesResponseBodyInstancesInstance) GetCpu() *string {
	return s.Cpu
}

func (s *DescribeInstancesResponseBodyInstancesInstance) GetCreationTime() *string {
	return s.CreationTime
}

func (s *DescribeInstancesResponseBodyInstancesInstance) GetDataDisk() *DescribeInstancesResponseBodyInstancesInstanceDataDisk {
	return s.DataDisk
}

func (s *DescribeInstancesResponseBodyInstancesInstance) GetDeletionProtection() *bool {
	return s.DeletionProtection
}

func (s *DescribeInstancesResponseBodyInstancesInstance) GetDisk() *int32 {
	return s.Disk
}

func (s *DescribeInstancesResponseBodyInstancesInstance) GetEnsRegionId() *string {
	return s.EnsRegionId
}

func (s *DescribeInstancesResponseBodyInstancesInstance) GetExpiredTime() *string {
	return s.ExpiredTime
}

func (s *DescribeInstancesResponseBodyInstancesInstance) GetHostName() *string {
	return s.HostName
}

func (s *DescribeInstancesResponseBodyInstancesInstance) GetImageId() *string {
	return s.ImageId
}

func (s *DescribeInstancesResponseBodyInstancesInstance) GetInnerIpAddress() *DescribeInstancesResponseBodyInstancesInstanceInnerIpAddress {
	return s.InnerIpAddress
}

func (s *DescribeInstancesResponseBodyInstancesInstance) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeInstancesResponseBodyInstancesInstance) GetInstanceName() *string {
	return s.InstanceName
}

func (s *DescribeInstancesResponseBodyInstancesInstance) GetInstanceResourceType() *string {
	return s.InstanceResourceType
}

func (s *DescribeInstancesResponseBodyInstancesInstance) GetInstanceTypeFamily() *string {
	return s.InstanceTypeFamily
}

func (s *DescribeInstancesResponseBodyInstancesInstance) GetInternetMaxBandwidthIn() *int32 {
	return s.InternetMaxBandwidthIn
}

func (s *DescribeInstancesResponseBodyInstancesInstance) GetInternetMaxBandwidthOut() *int32 {
	return s.InternetMaxBandwidthOut
}

func (s *DescribeInstancesResponseBodyInstancesInstance) GetKeyPairName() *string {
	return s.KeyPairName
}

func (s *DescribeInstancesResponseBodyInstancesInstance) GetMemory() *int32 {
	return s.Memory
}

func (s *DescribeInstancesResponseBodyInstancesInstance) GetNetworkAttributes() *DescribeInstancesResponseBodyInstancesInstanceNetworkAttributes {
	return s.NetworkAttributes
}

func (s *DescribeInstancesResponseBodyInstancesInstance) GetNetworkInterfaces() *DescribeInstancesResponseBodyInstancesInstanceNetworkInterfaces {
	return s.NetworkInterfaces
}

func (s *DescribeInstancesResponseBodyInstancesInstance) GetOSName() *string {
	return s.OSName
}

func (s *DescribeInstancesResponseBodyInstancesInstance) GetPrivateIpAddresses() *DescribeInstancesResponseBodyInstancesInstancePrivateIpAddresses {
	return s.PrivateIpAddresses
}

func (s *DescribeInstancesResponseBodyInstancesInstance) GetPublicIpAddress() *DescribeInstancesResponseBodyInstancesInstancePublicIpAddress {
	return s.PublicIpAddress
}

func (s *DescribeInstancesResponseBodyInstancesInstance) GetPublicIpAddresses() *DescribeInstancesResponseBodyInstancesInstancePublicIpAddresses {
	return s.PublicIpAddresses
}

func (s *DescribeInstancesResponseBodyInstancesInstance) GetSecurityGroupIds() *DescribeInstancesResponseBodyInstancesInstanceSecurityGroupIds {
	return s.SecurityGroupIds
}

func (s *DescribeInstancesResponseBodyInstancesInstance) GetServiceStatus() *string {
	return s.ServiceStatus
}

func (s *DescribeInstancesResponseBodyInstancesInstance) GetSpecName() *string {
	return s.SpecName
}

func (s *DescribeInstancesResponseBodyInstancesInstance) GetSpotStrategy() *string {
	return s.SpotStrategy
}

func (s *DescribeInstancesResponseBodyInstancesInstance) GetStatus() *string {
	return s.Status
}

func (s *DescribeInstancesResponseBodyInstancesInstance) GetSystemDisk() *DescribeInstancesResponseBodyInstancesInstanceSystemDisk {
	return s.SystemDisk
}

func (s *DescribeInstancesResponseBodyInstancesInstance) GetTags() *DescribeInstancesResponseBodyInstancesInstanceTags {
	return s.Tags
}

func (s *DescribeInstancesResponseBodyInstancesInstance) SetAutoReleaseTime(v string) *DescribeInstancesResponseBodyInstancesInstance {
	s.AutoReleaseTime = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesInstance) SetCpu(v string) *DescribeInstancesResponseBodyInstancesInstance {
	s.Cpu = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesInstance) SetCreationTime(v string) *DescribeInstancesResponseBodyInstancesInstance {
	s.CreationTime = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesInstance) SetDataDisk(v *DescribeInstancesResponseBodyInstancesInstanceDataDisk) *DescribeInstancesResponseBodyInstancesInstance {
	s.DataDisk = v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesInstance) SetDeletionProtection(v bool) *DescribeInstancesResponseBodyInstancesInstance {
	s.DeletionProtection = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesInstance) SetDisk(v int32) *DescribeInstancesResponseBodyInstancesInstance {
	s.Disk = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesInstance) SetEnsRegionId(v string) *DescribeInstancesResponseBodyInstancesInstance {
	s.EnsRegionId = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesInstance) SetExpiredTime(v string) *DescribeInstancesResponseBodyInstancesInstance {
	s.ExpiredTime = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesInstance) SetHostName(v string) *DescribeInstancesResponseBodyInstancesInstance {
	s.HostName = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesInstance) SetImageId(v string) *DescribeInstancesResponseBodyInstancesInstance {
	s.ImageId = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesInstance) SetInnerIpAddress(v *DescribeInstancesResponseBodyInstancesInstanceInnerIpAddress) *DescribeInstancesResponseBodyInstancesInstance {
	s.InnerIpAddress = v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesInstance) SetInstanceId(v string) *DescribeInstancesResponseBodyInstancesInstance {
	s.InstanceId = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesInstance) SetInstanceName(v string) *DescribeInstancesResponseBodyInstancesInstance {
	s.InstanceName = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesInstance) SetInstanceResourceType(v string) *DescribeInstancesResponseBodyInstancesInstance {
	s.InstanceResourceType = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesInstance) SetInstanceTypeFamily(v string) *DescribeInstancesResponseBodyInstancesInstance {
	s.InstanceTypeFamily = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesInstance) SetInternetMaxBandwidthIn(v int32) *DescribeInstancesResponseBodyInstancesInstance {
	s.InternetMaxBandwidthIn = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesInstance) SetInternetMaxBandwidthOut(v int32) *DescribeInstancesResponseBodyInstancesInstance {
	s.InternetMaxBandwidthOut = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesInstance) SetKeyPairName(v string) *DescribeInstancesResponseBodyInstancesInstance {
	s.KeyPairName = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesInstance) SetMemory(v int32) *DescribeInstancesResponseBodyInstancesInstance {
	s.Memory = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesInstance) SetNetworkAttributes(v *DescribeInstancesResponseBodyInstancesInstanceNetworkAttributes) *DescribeInstancesResponseBodyInstancesInstance {
	s.NetworkAttributes = v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesInstance) SetNetworkInterfaces(v *DescribeInstancesResponseBodyInstancesInstanceNetworkInterfaces) *DescribeInstancesResponseBodyInstancesInstance {
	s.NetworkInterfaces = v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesInstance) SetOSName(v string) *DescribeInstancesResponseBodyInstancesInstance {
	s.OSName = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesInstance) SetPrivateIpAddresses(v *DescribeInstancesResponseBodyInstancesInstancePrivateIpAddresses) *DescribeInstancesResponseBodyInstancesInstance {
	s.PrivateIpAddresses = v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesInstance) SetPublicIpAddress(v *DescribeInstancesResponseBodyInstancesInstancePublicIpAddress) *DescribeInstancesResponseBodyInstancesInstance {
	s.PublicIpAddress = v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesInstance) SetPublicIpAddresses(v *DescribeInstancesResponseBodyInstancesInstancePublicIpAddresses) *DescribeInstancesResponseBodyInstancesInstance {
	s.PublicIpAddresses = v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesInstance) SetSecurityGroupIds(v *DescribeInstancesResponseBodyInstancesInstanceSecurityGroupIds) *DescribeInstancesResponseBodyInstancesInstance {
	s.SecurityGroupIds = v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesInstance) SetServiceStatus(v string) *DescribeInstancesResponseBodyInstancesInstance {
	s.ServiceStatus = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesInstance) SetSpecName(v string) *DescribeInstancesResponseBodyInstancesInstance {
	s.SpecName = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesInstance) SetSpotStrategy(v string) *DescribeInstancesResponseBodyInstancesInstance {
	s.SpotStrategy = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesInstance) SetStatus(v string) *DescribeInstancesResponseBodyInstancesInstance {
	s.Status = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesInstance) SetSystemDisk(v *DescribeInstancesResponseBodyInstancesInstanceSystemDisk) *DescribeInstancesResponseBodyInstancesInstance {
	s.SystemDisk = v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesInstance) SetTags(v *DescribeInstancesResponseBodyInstancesInstanceTags) *DescribeInstancesResponseBodyInstancesInstance {
	s.Tags = v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesInstance) Validate() error {
	if s.DataDisk != nil {
		if err := s.DataDisk.Validate(); err != nil {
			return err
		}
	}
	if s.InnerIpAddress != nil {
		if err := s.InnerIpAddress.Validate(); err != nil {
			return err
		}
	}
	if s.NetworkAttributes != nil {
		if err := s.NetworkAttributes.Validate(); err != nil {
			return err
		}
	}
	if s.NetworkInterfaces != nil {
		if err := s.NetworkInterfaces.Validate(); err != nil {
			return err
		}
	}
	if s.PrivateIpAddresses != nil {
		if err := s.PrivateIpAddresses.Validate(); err != nil {
			return err
		}
	}
	if s.PublicIpAddress != nil {
		if err := s.PublicIpAddress.Validate(); err != nil {
			return err
		}
	}
	if s.PublicIpAddresses != nil {
		if err := s.PublicIpAddresses.Validate(); err != nil {
			return err
		}
	}
	if s.SecurityGroupIds != nil {
		if err := s.SecurityGroupIds.Validate(); err != nil {
			return err
		}
	}
	if s.SystemDisk != nil {
		if err := s.SystemDisk.Validate(); err != nil {
			return err
		}
	}
	if s.Tags != nil {
		if err := s.Tags.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeInstancesResponseBodyInstancesInstanceDataDisk struct {
	DataDisk []*DescribeInstancesResponseBodyInstancesInstanceDataDiskDataDisk `json:"DataDisk,omitempty" xml:"DataDisk,omitempty" type:"Repeated"`
}

func (s DescribeInstancesResponseBodyInstancesInstanceDataDisk) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstancesResponseBodyInstancesInstanceDataDisk) GoString() string {
	return s.String()
}

func (s *DescribeInstancesResponseBodyInstancesInstanceDataDisk) GetDataDisk() []*DescribeInstancesResponseBodyInstancesInstanceDataDiskDataDisk {
	return s.DataDisk
}

func (s *DescribeInstancesResponseBodyInstancesInstanceDataDisk) SetDataDisk(v []*DescribeInstancesResponseBodyInstancesInstanceDataDiskDataDisk) *DescribeInstancesResponseBodyInstancesInstanceDataDisk {
	s.DataDisk = v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesInstanceDataDisk) Validate() error {
	if s.DataDisk != nil {
		for _, item := range s.DataDisk {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeInstancesResponseBodyInstancesInstanceDataDiskDataDisk struct {
	Category     *string `json:"Category,omitempty" xml:"Category,omitempty"`
	DiskId       *string `json:"DiskId,omitempty" xml:"DiskId,omitempty"`
	DiskName     *string `json:"DiskName,omitempty" xml:"DiskName,omitempty"`
	DiskSize     *int32  `json:"DiskSize,omitempty" xml:"DiskSize,omitempty"`
	EncryptKeyId *string `json:"EncryptKeyId,omitempty" xml:"EncryptKeyId,omitempty"`
	Encrypted    *bool   `json:"Encrypted,omitempty" xml:"Encrypted,omitempty"`
	Size         *int32  `json:"Size,omitempty" xml:"Size,omitempty"`
	DeviceType   *string `json:"device_type,omitempty" xml:"device_type,omitempty"`
	DiskType     *string `json:"disk_type,omitempty" xml:"disk_type,omitempty"`
	Name         *string `json:"name,omitempty" xml:"name,omitempty"`
	Storage      *int32  `json:"storage,omitempty" xml:"storage,omitempty"`
	Uuid         *string `json:"uuid,omitempty" xml:"uuid,omitempty"`
}

func (s DescribeInstancesResponseBodyInstancesInstanceDataDiskDataDisk) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstancesResponseBodyInstancesInstanceDataDiskDataDisk) GoString() string {
	return s.String()
}

func (s *DescribeInstancesResponseBodyInstancesInstanceDataDiskDataDisk) GetCategory() *string {
	return s.Category
}

func (s *DescribeInstancesResponseBodyInstancesInstanceDataDiskDataDisk) GetDiskId() *string {
	return s.DiskId
}

func (s *DescribeInstancesResponseBodyInstancesInstanceDataDiskDataDisk) GetDiskName() *string {
	return s.DiskName
}

func (s *DescribeInstancesResponseBodyInstancesInstanceDataDiskDataDisk) GetDiskSize() *int32 {
	return s.DiskSize
}

func (s *DescribeInstancesResponseBodyInstancesInstanceDataDiskDataDisk) GetEncryptKeyId() *string {
	return s.EncryptKeyId
}

func (s *DescribeInstancesResponseBodyInstancesInstanceDataDiskDataDisk) GetEncrypted() *bool {
	return s.Encrypted
}

func (s *DescribeInstancesResponseBodyInstancesInstanceDataDiskDataDisk) GetSize() *int32 {
	return s.Size
}

func (s *DescribeInstancesResponseBodyInstancesInstanceDataDiskDataDisk) GetDeviceType() *string {
	return s.DeviceType
}

func (s *DescribeInstancesResponseBodyInstancesInstanceDataDiskDataDisk) GetDiskType() *string {
	return s.DiskType
}

func (s *DescribeInstancesResponseBodyInstancesInstanceDataDiskDataDisk) GetName() *string {
	return s.Name
}

func (s *DescribeInstancesResponseBodyInstancesInstanceDataDiskDataDisk) GetStorage() *int32 {
	return s.Storage
}

func (s *DescribeInstancesResponseBodyInstancesInstanceDataDiskDataDisk) GetUuid() *string {
	return s.Uuid
}

func (s *DescribeInstancesResponseBodyInstancesInstanceDataDiskDataDisk) SetCategory(v string) *DescribeInstancesResponseBodyInstancesInstanceDataDiskDataDisk {
	s.Category = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesInstanceDataDiskDataDisk) SetDiskId(v string) *DescribeInstancesResponseBodyInstancesInstanceDataDiskDataDisk {
	s.DiskId = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesInstanceDataDiskDataDisk) SetDiskName(v string) *DescribeInstancesResponseBodyInstancesInstanceDataDiskDataDisk {
	s.DiskName = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesInstanceDataDiskDataDisk) SetDiskSize(v int32) *DescribeInstancesResponseBodyInstancesInstanceDataDiskDataDisk {
	s.DiskSize = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesInstanceDataDiskDataDisk) SetEncryptKeyId(v string) *DescribeInstancesResponseBodyInstancesInstanceDataDiskDataDisk {
	s.EncryptKeyId = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesInstanceDataDiskDataDisk) SetEncrypted(v bool) *DescribeInstancesResponseBodyInstancesInstanceDataDiskDataDisk {
	s.Encrypted = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesInstanceDataDiskDataDisk) SetSize(v int32) *DescribeInstancesResponseBodyInstancesInstanceDataDiskDataDisk {
	s.Size = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesInstanceDataDiskDataDisk) SetDeviceType(v string) *DescribeInstancesResponseBodyInstancesInstanceDataDiskDataDisk {
	s.DeviceType = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesInstanceDataDiskDataDisk) SetDiskType(v string) *DescribeInstancesResponseBodyInstancesInstanceDataDiskDataDisk {
	s.DiskType = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesInstanceDataDiskDataDisk) SetName(v string) *DescribeInstancesResponseBodyInstancesInstanceDataDiskDataDisk {
	s.Name = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesInstanceDataDiskDataDisk) SetStorage(v int32) *DescribeInstancesResponseBodyInstancesInstanceDataDiskDataDisk {
	s.Storage = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesInstanceDataDiskDataDisk) SetUuid(v string) *DescribeInstancesResponseBodyInstancesInstanceDataDiskDataDisk {
	s.Uuid = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesInstanceDataDiskDataDisk) Validate() error {
	return dara.Validate(s)
}

type DescribeInstancesResponseBodyInstancesInstanceInnerIpAddress struct {
	IpAddress []*string `json:"IpAddress,omitempty" xml:"IpAddress,omitempty" type:"Repeated"`
}

func (s DescribeInstancesResponseBodyInstancesInstanceInnerIpAddress) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstancesResponseBodyInstancesInstanceInnerIpAddress) GoString() string {
	return s.String()
}

func (s *DescribeInstancesResponseBodyInstancesInstanceInnerIpAddress) GetIpAddress() []*string {
	return s.IpAddress
}

func (s *DescribeInstancesResponseBodyInstancesInstanceInnerIpAddress) SetIpAddress(v []*string) *DescribeInstancesResponseBodyInstancesInstanceInnerIpAddress {
	s.IpAddress = v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesInstanceInnerIpAddress) Validate() error {
	return dara.Validate(s)
}

type DescribeInstancesResponseBodyInstancesInstanceNetworkAttributes struct {
	NetworkId        *string                                                                          `json:"NetworkId,omitempty" xml:"NetworkId,omitempty"`
	PrivateIpAddress *DescribeInstancesResponseBodyInstancesInstanceNetworkAttributesPrivateIpAddress `json:"PrivateIpAddress,omitempty" xml:"PrivateIpAddress,omitempty" type:"Struct"`
	VSwitchId        *string                                                                          `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
}

func (s DescribeInstancesResponseBodyInstancesInstanceNetworkAttributes) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstancesResponseBodyInstancesInstanceNetworkAttributes) GoString() string {
	return s.String()
}

func (s *DescribeInstancesResponseBodyInstancesInstanceNetworkAttributes) GetNetworkId() *string {
	return s.NetworkId
}

func (s *DescribeInstancesResponseBodyInstancesInstanceNetworkAttributes) GetPrivateIpAddress() *DescribeInstancesResponseBodyInstancesInstanceNetworkAttributesPrivateIpAddress {
	return s.PrivateIpAddress
}

func (s *DescribeInstancesResponseBodyInstancesInstanceNetworkAttributes) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *DescribeInstancesResponseBodyInstancesInstanceNetworkAttributes) SetNetworkId(v string) *DescribeInstancesResponseBodyInstancesInstanceNetworkAttributes {
	s.NetworkId = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesInstanceNetworkAttributes) SetPrivateIpAddress(v *DescribeInstancesResponseBodyInstancesInstanceNetworkAttributesPrivateIpAddress) *DescribeInstancesResponseBodyInstancesInstanceNetworkAttributes {
	s.PrivateIpAddress = v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesInstanceNetworkAttributes) SetVSwitchId(v string) *DescribeInstancesResponseBodyInstancesInstanceNetworkAttributes {
	s.VSwitchId = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesInstanceNetworkAttributes) Validate() error {
	if s.PrivateIpAddress != nil {
		if err := s.PrivateIpAddress.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeInstancesResponseBodyInstancesInstanceNetworkAttributesPrivateIpAddress struct {
	IpAddress []*string `json:"IpAddress,omitempty" xml:"IpAddress,omitempty" type:"Repeated"`
}

func (s DescribeInstancesResponseBodyInstancesInstanceNetworkAttributesPrivateIpAddress) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstancesResponseBodyInstancesInstanceNetworkAttributesPrivateIpAddress) GoString() string {
	return s.String()
}

func (s *DescribeInstancesResponseBodyInstancesInstanceNetworkAttributesPrivateIpAddress) GetIpAddress() []*string {
	return s.IpAddress
}

func (s *DescribeInstancesResponseBodyInstancesInstanceNetworkAttributesPrivateIpAddress) SetIpAddress(v []*string) *DescribeInstancesResponseBodyInstancesInstanceNetworkAttributesPrivateIpAddress {
	s.IpAddress = v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesInstanceNetworkAttributesPrivateIpAddress) Validate() error {
	return dara.Validate(s)
}

type DescribeInstancesResponseBodyInstancesInstanceNetworkInterfaces struct {
	NetworkInterfaces []*DescribeInstancesResponseBodyInstancesInstanceNetworkInterfacesNetworkInterfaces `json:"NetworkInterfaces,omitempty" xml:"NetworkInterfaces,omitempty" type:"Repeated"`
}

func (s DescribeInstancesResponseBodyInstancesInstanceNetworkInterfaces) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstancesResponseBodyInstancesInstanceNetworkInterfaces) GoString() string {
	return s.String()
}

func (s *DescribeInstancesResponseBodyInstancesInstanceNetworkInterfaces) GetNetworkInterfaces() []*DescribeInstancesResponseBodyInstancesInstanceNetworkInterfacesNetworkInterfaces {
	return s.NetworkInterfaces
}

func (s *DescribeInstancesResponseBodyInstancesInstanceNetworkInterfaces) SetNetworkInterfaces(v []*DescribeInstancesResponseBodyInstancesInstanceNetworkInterfacesNetworkInterfaces) *DescribeInstancesResponseBodyInstancesInstanceNetworkInterfaces {
	s.NetworkInterfaces = v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesInstanceNetworkInterfaces) Validate() error {
	if s.NetworkInterfaces != nil {
		for _, item := range s.NetworkInterfaces {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeInstancesResponseBodyInstancesInstanceNetworkInterfacesNetworkInterfaces struct {
	Ipv6Sets           *DescribeInstancesResponseBodyInstancesInstanceNetworkInterfacesNetworkInterfacesIpv6Sets      `json:"Ipv6Sets,omitempty" xml:"Ipv6Sets,omitempty" type:"Struct"`
	MacAddress         *string                                                                                        `json:"MacAddress,omitempty" xml:"MacAddress,omitempty"`
	NetworkInterfaceId *string                                                                                        `json:"NetworkInterfaceId,omitempty" xml:"NetworkInterfaceId,omitempty"`
	PrimaryIpAddress   *string                                                                                        `json:"PrimaryIpAddress,omitempty" xml:"PrimaryIpAddress,omitempty"`
	PrivateIpSets      *DescribeInstancesResponseBodyInstancesInstanceNetworkInterfacesNetworkInterfacesPrivateIpSets `json:"PrivateIpSets,omitempty" xml:"PrivateIpSets,omitempty" type:"Struct"`
	Type               *string                                                                                        `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DescribeInstancesResponseBodyInstancesInstanceNetworkInterfacesNetworkInterfaces) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstancesResponseBodyInstancesInstanceNetworkInterfacesNetworkInterfaces) GoString() string {
	return s.String()
}

func (s *DescribeInstancesResponseBodyInstancesInstanceNetworkInterfacesNetworkInterfaces) GetIpv6Sets() *DescribeInstancesResponseBodyInstancesInstanceNetworkInterfacesNetworkInterfacesIpv6Sets {
	return s.Ipv6Sets
}

func (s *DescribeInstancesResponseBodyInstancesInstanceNetworkInterfacesNetworkInterfaces) GetMacAddress() *string {
	return s.MacAddress
}

func (s *DescribeInstancesResponseBodyInstancesInstanceNetworkInterfacesNetworkInterfaces) GetNetworkInterfaceId() *string {
	return s.NetworkInterfaceId
}

func (s *DescribeInstancesResponseBodyInstancesInstanceNetworkInterfacesNetworkInterfaces) GetPrimaryIpAddress() *string {
	return s.PrimaryIpAddress
}

func (s *DescribeInstancesResponseBodyInstancesInstanceNetworkInterfacesNetworkInterfaces) GetPrivateIpSets() *DescribeInstancesResponseBodyInstancesInstanceNetworkInterfacesNetworkInterfacesPrivateIpSets {
	return s.PrivateIpSets
}

func (s *DescribeInstancesResponseBodyInstancesInstanceNetworkInterfacesNetworkInterfaces) GetType() *string {
	return s.Type
}

func (s *DescribeInstancesResponseBodyInstancesInstanceNetworkInterfacesNetworkInterfaces) SetIpv6Sets(v *DescribeInstancesResponseBodyInstancesInstanceNetworkInterfacesNetworkInterfacesIpv6Sets) *DescribeInstancesResponseBodyInstancesInstanceNetworkInterfacesNetworkInterfaces {
	s.Ipv6Sets = v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesInstanceNetworkInterfacesNetworkInterfaces) SetMacAddress(v string) *DescribeInstancesResponseBodyInstancesInstanceNetworkInterfacesNetworkInterfaces {
	s.MacAddress = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesInstanceNetworkInterfacesNetworkInterfaces) SetNetworkInterfaceId(v string) *DescribeInstancesResponseBodyInstancesInstanceNetworkInterfacesNetworkInterfaces {
	s.NetworkInterfaceId = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesInstanceNetworkInterfacesNetworkInterfaces) SetPrimaryIpAddress(v string) *DescribeInstancesResponseBodyInstancesInstanceNetworkInterfacesNetworkInterfaces {
	s.PrimaryIpAddress = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesInstanceNetworkInterfacesNetworkInterfaces) SetPrivateIpSets(v *DescribeInstancesResponseBodyInstancesInstanceNetworkInterfacesNetworkInterfacesPrivateIpSets) *DescribeInstancesResponseBodyInstancesInstanceNetworkInterfacesNetworkInterfaces {
	s.PrivateIpSets = v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesInstanceNetworkInterfacesNetworkInterfaces) SetType(v string) *DescribeInstancesResponseBodyInstancesInstanceNetworkInterfacesNetworkInterfaces {
	s.Type = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesInstanceNetworkInterfacesNetworkInterfaces) Validate() error {
	if s.Ipv6Sets != nil {
		if err := s.Ipv6Sets.Validate(); err != nil {
			return err
		}
	}
	if s.PrivateIpSets != nil {
		if err := s.PrivateIpSets.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeInstancesResponseBodyInstancesInstanceNetworkInterfacesNetworkInterfacesIpv6Sets struct {
	Ipv6Set []*DescribeInstancesResponseBodyInstancesInstanceNetworkInterfacesNetworkInterfacesIpv6SetsIpv6Set `json:"Ipv6Set,omitempty" xml:"Ipv6Set,omitempty" type:"Repeated"`
}

func (s DescribeInstancesResponseBodyInstancesInstanceNetworkInterfacesNetworkInterfacesIpv6Sets) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstancesResponseBodyInstancesInstanceNetworkInterfacesNetworkInterfacesIpv6Sets) GoString() string {
	return s.String()
}

func (s *DescribeInstancesResponseBodyInstancesInstanceNetworkInterfacesNetworkInterfacesIpv6Sets) GetIpv6Set() []*DescribeInstancesResponseBodyInstancesInstanceNetworkInterfacesNetworkInterfacesIpv6SetsIpv6Set {
	return s.Ipv6Set
}

func (s *DescribeInstancesResponseBodyInstancesInstanceNetworkInterfacesNetworkInterfacesIpv6Sets) SetIpv6Set(v []*DescribeInstancesResponseBodyInstancesInstanceNetworkInterfacesNetworkInterfacesIpv6SetsIpv6Set) *DescribeInstancesResponseBodyInstancesInstanceNetworkInterfacesNetworkInterfacesIpv6Sets {
	s.Ipv6Set = v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesInstanceNetworkInterfacesNetworkInterfacesIpv6Sets) Validate() error {
	if s.Ipv6Set != nil {
		for _, item := range s.Ipv6Set {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeInstancesResponseBodyInstancesInstanceNetworkInterfacesNetworkInterfacesIpv6SetsIpv6Set struct {
	Ipv6Address *string `json:"Ipv6Address,omitempty" xml:"Ipv6Address,omitempty"`
}

func (s DescribeInstancesResponseBodyInstancesInstanceNetworkInterfacesNetworkInterfacesIpv6SetsIpv6Set) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstancesResponseBodyInstancesInstanceNetworkInterfacesNetworkInterfacesIpv6SetsIpv6Set) GoString() string {
	return s.String()
}

func (s *DescribeInstancesResponseBodyInstancesInstanceNetworkInterfacesNetworkInterfacesIpv6SetsIpv6Set) GetIpv6Address() *string {
	return s.Ipv6Address
}

func (s *DescribeInstancesResponseBodyInstancesInstanceNetworkInterfacesNetworkInterfacesIpv6SetsIpv6Set) SetIpv6Address(v string) *DescribeInstancesResponseBodyInstancesInstanceNetworkInterfacesNetworkInterfacesIpv6SetsIpv6Set {
	s.Ipv6Address = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesInstanceNetworkInterfacesNetworkInterfacesIpv6SetsIpv6Set) Validate() error {
	return dara.Validate(s)
}

type DescribeInstancesResponseBodyInstancesInstanceNetworkInterfacesNetworkInterfacesPrivateIpSets struct {
	PrivateIpSet []*DescribeInstancesResponseBodyInstancesInstanceNetworkInterfacesNetworkInterfacesPrivateIpSetsPrivateIpSet `json:"PrivateIpSet,omitempty" xml:"PrivateIpSet,omitempty" type:"Repeated"`
}

func (s DescribeInstancesResponseBodyInstancesInstanceNetworkInterfacesNetworkInterfacesPrivateIpSets) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstancesResponseBodyInstancesInstanceNetworkInterfacesNetworkInterfacesPrivateIpSets) GoString() string {
	return s.String()
}

func (s *DescribeInstancesResponseBodyInstancesInstanceNetworkInterfacesNetworkInterfacesPrivateIpSets) GetPrivateIpSet() []*DescribeInstancesResponseBodyInstancesInstanceNetworkInterfacesNetworkInterfacesPrivateIpSetsPrivateIpSet {
	return s.PrivateIpSet
}

func (s *DescribeInstancesResponseBodyInstancesInstanceNetworkInterfacesNetworkInterfacesPrivateIpSets) SetPrivateIpSet(v []*DescribeInstancesResponseBodyInstancesInstanceNetworkInterfacesNetworkInterfacesPrivateIpSetsPrivateIpSet) *DescribeInstancesResponseBodyInstancesInstanceNetworkInterfacesNetworkInterfacesPrivateIpSets {
	s.PrivateIpSet = v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesInstanceNetworkInterfacesNetworkInterfacesPrivateIpSets) Validate() error {
	if s.PrivateIpSet != nil {
		for _, item := range s.PrivateIpSet {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeInstancesResponseBodyInstancesInstanceNetworkInterfacesNetworkInterfacesPrivateIpSetsPrivateIpSet struct {
	Primary          *bool   `json:"Primary,omitempty" xml:"Primary,omitempty"`
	PrivateIpAddress *string `json:"PrivateIpAddress,omitempty" xml:"PrivateIpAddress,omitempty"`
}

func (s DescribeInstancesResponseBodyInstancesInstanceNetworkInterfacesNetworkInterfacesPrivateIpSetsPrivateIpSet) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstancesResponseBodyInstancesInstanceNetworkInterfacesNetworkInterfacesPrivateIpSetsPrivateIpSet) GoString() string {
	return s.String()
}

func (s *DescribeInstancesResponseBodyInstancesInstanceNetworkInterfacesNetworkInterfacesPrivateIpSetsPrivateIpSet) GetPrimary() *bool {
	return s.Primary
}

func (s *DescribeInstancesResponseBodyInstancesInstanceNetworkInterfacesNetworkInterfacesPrivateIpSetsPrivateIpSet) GetPrivateIpAddress() *string {
	return s.PrivateIpAddress
}

func (s *DescribeInstancesResponseBodyInstancesInstanceNetworkInterfacesNetworkInterfacesPrivateIpSetsPrivateIpSet) SetPrimary(v bool) *DescribeInstancesResponseBodyInstancesInstanceNetworkInterfacesNetworkInterfacesPrivateIpSetsPrivateIpSet {
	s.Primary = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesInstanceNetworkInterfacesNetworkInterfacesPrivateIpSetsPrivateIpSet) SetPrivateIpAddress(v string) *DescribeInstancesResponseBodyInstancesInstanceNetworkInterfacesNetworkInterfacesPrivateIpSetsPrivateIpSet {
	s.PrivateIpAddress = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesInstanceNetworkInterfacesNetworkInterfacesPrivateIpSetsPrivateIpSet) Validate() error {
	return dara.Validate(s)
}

type DescribeInstancesResponseBodyInstancesInstancePrivateIpAddresses struct {
	PrivateIpAddress []*DescribeInstancesResponseBodyInstancesInstancePrivateIpAddressesPrivateIpAddress `json:"PrivateIpAddress,omitempty" xml:"PrivateIpAddress,omitempty" type:"Repeated"`
}

func (s DescribeInstancesResponseBodyInstancesInstancePrivateIpAddresses) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstancesResponseBodyInstancesInstancePrivateIpAddresses) GoString() string {
	return s.String()
}

func (s *DescribeInstancesResponseBodyInstancesInstancePrivateIpAddresses) GetPrivateIpAddress() []*DescribeInstancesResponseBodyInstancesInstancePrivateIpAddressesPrivateIpAddress {
	return s.PrivateIpAddress
}

func (s *DescribeInstancesResponseBodyInstancesInstancePrivateIpAddresses) SetPrivateIpAddress(v []*DescribeInstancesResponseBodyInstancesInstancePrivateIpAddressesPrivateIpAddress) *DescribeInstancesResponseBodyInstancesInstancePrivateIpAddresses {
	s.PrivateIpAddress = v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesInstancePrivateIpAddresses) Validate() error {
	if s.PrivateIpAddress != nil {
		for _, item := range s.PrivateIpAddress {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeInstancesResponseBodyInstancesInstancePrivateIpAddressesPrivateIpAddress struct {
	GateWay *string `json:"GateWay,omitempty" xml:"GateWay,omitempty"`
	Ip      *string `json:"Ip,omitempty" xml:"Ip,omitempty"`
	Isp     *string `json:"Isp,omitempty" xml:"Isp,omitempty"`
}

func (s DescribeInstancesResponseBodyInstancesInstancePrivateIpAddressesPrivateIpAddress) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstancesResponseBodyInstancesInstancePrivateIpAddressesPrivateIpAddress) GoString() string {
	return s.String()
}

func (s *DescribeInstancesResponseBodyInstancesInstancePrivateIpAddressesPrivateIpAddress) GetGateWay() *string {
	return s.GateWay
}

func (s *DescribeInstancesResponseBodyInstancesInstancePrivateIpAddressesPrivateIpAddress) GetIp() *string {
	return s.Ip
}

func (s *DescribeInstancesResponseBodyInstancesInstancePrivateIpAddressesPrivateIpAddress) GetIsp() *string {
	return s.Isp
}

func (s *DescribeInstancesResponseBodyInstancesInstancePrivateIpAddressesPrivateIpAddress) SetGateWay(v string) *DescribeInstancesResponseBodyInstancesInstancePrivateIpAddressesPrivateIpAddress {
	s.GateWay = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesInstancePrivateIpAddressesPrivateIpAddress) SetIp(v string) *DescribeInstancesResponseBodyInstancesInstancePrivateIpAddressesPrivateIpAddress {
	s.Ip = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesInstancePrivateIpAddressesPrivateIpAddress) SetIsp(v string) *DescribeInstancesResponseBodyInstancesInstancePrivateIpAddressesPrivateIpAddress {
	s.Isp = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesInstancePrivateIpAddressesPrivateIpAddress) Validate() error {
	return dara.Validate(s)
}

type DescribeInstancesResponseBodyInstancesInstancePublicIpAddress struct {
	IpAddress []*string `json:"IpAddress,omitempty" xml:"IpAddress,omitempty" type:"Repeated"`
}

func (s DescribeInstancesResponseBodyInstancesInstancePublicIpAddress) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstancesResponseBodyInstancesInstancePublicIpAddress) GoString() string {
	return s.String()
}

func (s *DescribeInstancesResponseBodyInstancesInstancePublicIpAddress) GetIpAddress() []*string {
	return s.IpAddress
}

func (s *DescribeInstancesResponseBodyInstancesInstancePublicIpAddress) SetIpAddress(v []*string) *DescribeInstancesResponseBodyInstancesInstancePublicIpAddress {
	s.IpAddress = v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesInstancePublicIpAddress) Validate() error {
	return dara.Validate(s)
}

type DescribeInstancesResponseBodyInstancesInstancePublicIpAddresses struct {
	PublicIpAddress []*DescribeInstancesResponseBodyInstancesInstancePublicIpAddressesPublicIpAddress `json:"PublicIpAddress,omitempty" xml:"PublicIpAddress,omitempty" type:"Repeated"`
}

func (s DescribeInstancesResponseBodyInstancesInstancePublicIpAddresses) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstancesResponseBodyInstancesInstancePublicIpAddresses) GoString() string {
	return s.String()
}

func (s *DescribeInstancesResponseBodyInstancesInstancePublicIpAddresses) GetPublicIpAddress() []*DescribeInstancesResponseBodyInstancesInstancePublicIpAddressesPublicIpAddress {
	return s.PublicIpAddress
}

func (s *DescribeInstancesResponseBodyInstancesInstancePublicIpAddresses) SetPublicIpAddress(v []*DescribeInstancesResponseBodyInstancesInstancePublicIpAddressesPublicIpAddress) *DescribeInstancesResponseBodyInstancesInstancePublicIpAddresses {
	s.PublicIpAddress = v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesInstancePublicIpAddresses) Validate() error {
	if s.PublicIpAddress != nil {
		for _, item := range s.PublicIpAddress {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeInstancesResponseBodyInstancesInstancePublicIpAddressesPublicIpAddress struct {
	GateWay *string `json:"GateWay,omitempty" xml:"GateWay,omitempty"`
	Ip      *string `json:"Ip,omitempty" xml:"Ip,omitempty"`
	Isp     *string `json:"Isp,omitempty" xml:"Isp,omitempty"`
}

func (s DescribeInstancesResponseBodyInstancesInstancePublicIpAddressesPublicIpAddress) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstancesResponseBodyInstancesInstancePublicIpAddressesPublicIpAddress) GoString() string {
	return s.String()
}

func (s *DescribeInstancesResponseBodyInstancesInstancePublicIpAddressesPublicIpAddress) GetGateWay() *string {
	return s.GateWay
}

func (s *DescribeInstancesResponseBodyInstancesInstancePublicIpAddressesPublicIpAddress) GetIp() *string {
	return s.Ip
}

func (s *DescribeInstancesResponseBodyInstancesInstancePublicIpAddressesPublicIpAddress) GetIsp() *string {
	return s.Isp
}

func (s *DescribeInstancesResponseBodyInstancesInstancePublicIpAddressesPublicIpAddress) SetGateWay(v string) *DescribeInstancesResponseBodyInstancesInstancePublicIpAddressesPublicIpAddress {
	s.GateWay = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesInstancePublicIpAddressesPublicIpAddress) SetIp(v string) *DescribeInstancesResponseBodyInstancesInstancePublicIpAddressesPublicIpAddress {
	s.Ip = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesInstancePublicIpAddressesPublicIpAddress) SetIsp(v string) *DescribeInstancesResponseBodyInstancesInstancePublicIpAddressesPublicIpAddress {
	s.Isp = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesInstancePublicIpAddressesPublicIpAddress) Validate() error {
	return dara.Validate(s)
}

type DescribeInstancesResponseBodyInstancesInstanceSecurityGroupIds struct {
	SecurityGroupId []*string `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty" type:"Repeated"`
}

func (s DescribeInstancesResponseBodyInstancesInstanceSecurityGroupIds) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstancesResponseBodyInstancesInstanceSecurityGroupIds) GoString() string {
	return s.String()
}

func (s *DescribeInstancesResponseBodyInstancesInstanceSecurityGroupIds) GetSecurityGroupId() []*string {
	return s.SecurityGroupId
}

func (s *DescribeInstancesResponseBodyInstancesInstanceSecurityGroupIds) SetSecurityGroupId(v []*string) *DescribeInstancesResponseBodyInstancesInstanceSecurityGroupIds {
	s.SecurityGroupId = v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesInstanceSecurityGroupIds) Validate() error {
	return dara.Validate(s)
}

type DescribeInstancesResponseBodyInstancesInstanceSystemDisk struct {
	Category   *string `json:"Category,omitempty" xml:"Category,omitempty"`
	DiskId     *string `json:"DiskId,omitempty" xml:"DiskId,omitempty"`
	DiskName   *string `json:"DiskName,omitempty" xml:"DiskName,omitempty"`
	Size       *int32  `json:"Size,omitempty" xml:"Size,omitempty"`
	DeviceType *string `json:"device_type,omitempty" xml:"device_type,omitempty"`
	DiskType   *string `json:"disk_type,omitempty" xml:"disk_type,omitempty"`
	Name       *string `json:"name,omitempty" xml:"name,omitempty"`
	Storage    *int32  `json:"storage,omitempty" xml:"storage,omitempty"`
	Uuid       *string `json:"uuid,omitempty" xml:"uuid,omitempty"`
}

func (s DescribeInstancesResponseBodyInstancesInstanceSystemDisk) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstancesResponseBodyInstancesInstanceSystemDisk) GoString() string {
	return s.String()
}

func (s *DescribeInstancesResponseBodyInstancesInstanceSystemDisk) GetCategory() *string {
	return s.Category
}

func (s *DescribeInstancesResponseBodyInstancesInstanceSystemDisk) GetDiskId() *string {
	return s.DiskId
}

func (s *DescribeInstancesResponseBodyInstancesInstanceSystemDisk) GetDiskName() *string {
	return s.DiskName
}

func (s *DescribeInstancesResponseBodyInstancesInstanceSystemDisk) GetSize() *int32 {
	return s.Size
}

func (s *DescribeInstancesResponseBodyInstancesInstanceSystemDisk) GetDeviceType() *string {
	return s.DeviceType
}

func (s *DescribeInstancesResponseBodyInstancesInstanceSystemDisk) GetDiskType() *string {
	return s.DiskType
}

func (s *DescribeInstancesResponseBodyInstancesInstanceSystemDisk) GetName() *string {
	return s.Name
}

func (s *DescribeInstancesResponseBodyInstancesInstanceSystemDisk) GetStorage() *int32 {
	return s.Storage
}

func (s *DescribeInstancesResponseBodyInstancesInstanceSystemDisk) GetUuid() *string {
	return s.Uuid
}

func (s *DescribeInstancesResponseBodyInstancesInstanceSystemDisk) SetCategory(v string) *DescribeInstancesResponseBodyInstancesInstanceSystemDisk {
	s.Category = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesInstanceSystemDisk) SetDiskId(v string) *DescribeInstancesResponseBodyInstancesInstanceSystemDisk {
	s.DiskId = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesInstanceSystemDisk) SetDiskName(v string) *DescribeInstancesResponseBodyInstancesInstanceSystemDisk {
	s.DiskName = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesInstanceSystemDisk) SetSize(v int32) *DescribeInstancesResponseBodyInstancesInstanceSystemDisk {
	s.Size = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesInstanceSystemDisk) SetDeviceType(v string) *DescribeInstancesResponseBodyInstancesInstanceSystemDisk {
	s.DeviceType = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesInstanceSystemDisk) SetDiskType(v string) *DescribeInstancesResponseBodyInstancesInstanceSystemDisk {
	s.DiskType = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesInstanceSystemDisk) SetName(v string) *DescribeInstancesResponseBodyInstancesInstanceSystemDisk {
	s.Name = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesInstanceSystemDisk) SetStorage(v int32) *DescribeInstancesResponseBodyInstancesInstanceSystemDisk {
	s.Storage = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesInstanceSystemDisk) SetUuid(v string) *DescribeInstancesResponseBodyInstancesInstanceSystemDisk {
	s.Uuid = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesInstanceSystemDisk) Validate() error {
	return dara.Validate(s)
}

type DescribeInstancesResponseBodyInstancesInstanceTags struct {
	Tags []*DescribeInstancesResponseBodyInstancesInstanceTagsTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
}

func (s DescribeInstancesResponseBodyInstancesInstanceTags) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstancesResponseBodyInstancesInstanceTags) GoString() string {
	return s.String()
}

func (s *DescribeInstancesResponseBodyInstancesInstanceTags) GetTags() []*DescribeInstancesResponseBodyInstancesInstanceTagsTags {
	return s.Tags
}

func (s *DescribeInstancesResponseBodyInstancesInstanceTags) SetTags(v []*DescribeInstancesResponseBodyInstancesInstanceTagsTags) *DescribeInstancesResponseBodyInstancesInstanceTags {
	s.Tags = v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesInstanceTags) Validate() error {
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

type DescribeInstancesResponseBodyInstancesInstanceTagsTags struct {
	TagKey   *string `json:"TagKey,omitempty" xml:"TagKey,omitempty"`
	TagValue *string `json:"TagValue,omitempty" xml:"TagValue,omitempty"`
}

func (s DescribeInstancesResponseBodyInstancesInstanceTagsTags) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstancesResponseBodyInstancesInstanceTagsTags) GoString() string {
	return s.String()
}

func (s *DescribeInstancesResponseBodyInstancesInstanceTagsTags) GetTagKey() *string {
	return s.TagKey
}

func (s *DescribeInstancesResponseBodyInstancesInstanceTagsTags) GetTagValue() *string {
	return s.TagValue
}

func (s *DescribeInstancesResponseBodyInstancesInstanceTagsTags) SetTagKey(v string) *DescribeInstancesResponseBodyInstancesInstanceTagsTags {
	s.TagKey = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesInstanceTagsTags) SetTagValue(v string) *DescribeInstancesResponseBodyInstancesInstanceTagsTags {
	s.TagValue = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesInstanceTagsTags) Validate() error {
	return dara.Validate(s)
}
