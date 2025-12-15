// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAvailableFileSystemsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetFileSystemList(v []*ListAvailableFileSystemsResponseBodyFileSystemList) *ListAvailableFileSystemsResponseBody
	GetFileSystemList() []*ListAvailableFileSystemsResponseBodyFileSystemList
	SetPageNumber(v int32) *ListAvailableFileSystemsResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *ListAvailableFileSystemsResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *ListAvailableFileSystemsResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *ListAvailableFileSystemsResponseBody
	GetTotalCount() *int32
}

type ListAvailableFileSystemsResponseBody struct {
	// The file systems.
	FileSystemList []*ListAvailableFileSystemsResponseBodyFileSystemList `json:"FileSystemList,omitempty" xml:"FileSystemList,omitempty" type:"Repeated"`
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
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID.
	//
	// example:
	//
	// BF4E8AB1-02A3-5ECF-87CC-3AB7BE98**
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 65
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListAvailableFileSystemsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListAvailableFileSystemsResponseBody) GoString() string {
	return s.String()
}

func (s *ListAvailableFileSystemsResponseBody) GetFileSystemList() []*ListAvailableFileSystemsResponseBodyFileSystemList {
	return s.FileSystemList
}

func (s *ListAvailableFileSystemsResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListAvailableFileSystemsResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListAvailableFileSystemsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListAvailableFileSystemsResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListAvailableFileSystemsResponseBody) SetFileSystemList(v []*ListAvailableFileSystemsResponseBodyFileSystemList) *ListAvailableFileSystemsResponseBody {
	s.FileSystemList = v
	return s
}

func (s *ListAvailableFileSystemsResponseBody) SetPageNumber(v int32) *ListAvailableFileSystemsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListAvailableFileSystemsResponseBody) SetPageSize(v int32) *ListAvailableFileSystemsResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListAvailableFileSystemsResponseBody) SetRequestId(v string) *ListAvailableFileSystemsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListAvailableFileSystemsResponseBody) SetTotalCount(v int32) *ListAvailableFileSystemsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListAvailableFileSystemsResponseBody) Validate() error {
	if s.FileSystemList != nil {
		for _, item := range s.FileSystemList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListAvailableFileSystemsResponseBodyFileSystemList struct {
	// The time when the file system was created.
	//
	// example:
	//
	// 2024-7-29 15:43:53
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The ID of the file system.
	//
	// example:
	//
	// 2fa0248***
	FileSystemId *string `json:"FileSystemId,omitempty" xml:"FileSystemId,omitempty"`
	// The type of the file system. Valid values:
	//
	// 	- standard: general-purpose network-attached storage (NAS) file system
	//
	// 	- extreme: extreme NAS file system
	//
	// example:
	//
	// cpfs
	FileSystemType *string `json:"FileSystemType,omitempty" xml:"FileSystemType,omitempty"`
	// The mount targets of the file system.
	MountTargetList []*ListAvailableFileSystemsResponseBodyFileSystemListMountTargetList `json:"MountTargetList,omitempty" xml:"MountTargetList,omitempty" type:"Repeated"`
	// The protocol type of the file system. Valid values:
	//
	// 	- nfs
	//
	// 	- smb
	//
	// 	- cpfs
	//
	// example:
	//
	// cpfs
	ProtocolType *string `json:"ProtocolType,omitempty" xml:"ProtocolType,omitempty"`
	// The state of the file system. Valid values:
	//
	// 	- Pending: The file system is processing a task.
	//
	// 	- Running: The file system is available. You can perform subsequent operations, such as creating a mount target, only when the file system is in the Running state.
	//
	// 	- Stopped: The file system is unavailable.
	//
	// 	- Extending: The file system is being scaled out.
	//
	// 	- Stopping: The file system is being stopped.
	//
	// 	- Deleting: The file system is being deleted.
	//
	// example:
	//
	// Running
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The storage type of the file system.
	//
	// 	- Valid values if FileSystemType is set to standard: Capacity and Performance.
	//
	// 	- Valid values if FileSystemType is set to extreme: standard and advance.
	//
	// example:
	//
	// Performance
	StorageType *string `json:"StorageType,omitempty" xml:"StorageType,omitempty"`
	// The VPC ID.
	//
	// example:
	//
	// vpc-bp132kgui8n0targb****
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s ListAvailableFileSystemsResponseBodyFileSystemList) String() string {
	return dara.Prettify(s)
}

func (s ListAvailableFileSystemsResponseBodyFileSystemList) GoString() string {
	return s.String()
}

func (s *ListAvailableFileSystemsResponseBodyFileSystemList) GetCreateTime() *string {
	return s.CreateTime
}

func (s *ListAvailableFileSystemsResponseBodyFileSystemList) GetFileSystemId() *string {
	return s.FileSystemId
}

func (s *ListAvailableFileSystemsResponseBodyFileSystemList) GetFileSystemType() *string {
	return s.FileSystemType
}

func (s *ListAvailableFileSystemsResponseBodyFileSystemList) GetMountTargetList() []*ListAvailableFileSystemsResponseBodyFileSystemListMountTargetList {
	return s.MountTargetList
}

func (s *ListAvailableFileSystemsResponseBodyFileSystemList) GetProtocolType() *string {
	return s.ProtocolType
}

func (s *ListAvailableFileSystemsResponseBodyFileSystemList) GetStatus() *string {
	return s.Status
}

func (s *ListAvailableFileSystemsResponseBodyFileSystemList) GetStorageType() *string {
	return s.StorageType
}

func (s *ListAvailableFileSystemsResponseBodyFileSystemList) GetVpcId() *string {
	return s.VpcId
}

func (s *ListAvailableFileSystemsResponseBodyFileSystemList) SetCreateTime(v string) *ListAvailableFileSystemsResponseBodyFileSystemList {
	s.CreateTime = &v
	return s
}

func (s *ListAvailableFileSystemsResponseBodyFileSystemList) SetFileSystemId(v string) *ListAvailableFileSystemsResponseBodyFileSystemList {
	s.FileSystemId = &v
	return s
}

func (s *ListAvailableFileSystemsResponseBodyFileSystemList) SetFileSystemType(v string) *ListAvailableFileSystemsResponseBodyFileSystemList {
	s.FileSystemType = &v
	return s
}

func (s *ListAvailableFileSystemsResponseBodyFileSystemList) SetMountTargetList(v []*ListAvailableFileSystemsResponseBodyFileSystemListMountTargetList) *ListAvailableFileSystemsResponseBodyFileSystemList {
	s.MountTargetList = v
	return s
}

func (s *ListAvailableFileSystemsResponseBodyFileSystemList) SetProtocolType(v string) *ListAvailableFileSystemsResponseBodyFileSystemList {
	s.ProtocolType = &v
	return s
}

func (s *ListAvailableFileSystemsResponseBodyFileSystemList) SetStatus(v string) *ListAvailableFileSystemsResponseBodyFileSystemList {
	s.Status = &v
	return s
}

func (s *ListAvailableFileSystemsResponseBodyFileSystemList) SetStorageType(v string) *ListAvailableFileSystemsResponseBodyFileSystemList {
	s.StorageType = &v
	return s
}

func (s *ListAvailableFileSystemsResponseBodyFileSystemList) SetVpcId(v string) *ListAvailableFileSystemsResponseBodyFileSystemList {
	s.VpcId = &v
	return s
}

func (s *ListAvailableFileSystemsResponseBodyFileSystemList) Validate() error {
	if s.MountTargetList != nil {
		for _, item := range s.MountTargetList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListAvailableFileSystemsResponseBodyFileSystemListMountTargetList struct {
	// The address of the mount target.
	//
	// example:
	//
	// c0967****.cn-hangzhou.cpfs.nas.aliyuncs.com
	MountTargetDomain *string `json:"MountTargetDomain,omitempty" xml:"MountTargetDomain,omitempty"`
	// The network type. Valid values: Valid values:
	//
	// 	- vpc
	//
	// example:
	//
	// vpc
	NetworkType *string `json:"NetworkType,omitempty" xml:"NetworkType,omitempty"`
	// The state of the mount target. Valid values:
	//
	// 	- Active: The mount target is available.
	//
	// 	- Inactive: The mount target is unavailable.
	//
	// 	- Pending: The mount target is being mounted.
	//
	// 	- Deleting: The mount target is being deleted.
	//
	// example:
	//
	// Active
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The vSwitch ID.
	//
	// example:
	//
	// vsw-2ze0c41hwu7lc29ris***
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// The virtual private cloud (VPC) ID.
	//
	// example:
	//
	// vpc-8vbvb34rtyh6xb3zrehs1****
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s ListAvailableFileSystemsResponseBodyFileSystemListMountTargetList) String() string {
	return dara.Prettify(s)
}

func (s ListAvailableFileSystemsResponseBodyFileSystemListMountTargetList) GoString() string {
	return s.String()
}

func (s *ListAvailableFileSystemsResponseBodyFileSystemListMountTargetList) GetMountTargetDomain() *string {
	return s.MountTargetDomain
}

func (s *ListAvailableFileSystemsResponseBodyFileSystemListMountTargetList) GetNetworkType() *string {
	return s.NetworkType
}

func (s *ListAvailableFileSystemsResponseBodyFileSystemListMountTargetList) GetStatus() *string {
	return s.Status
}

func (s *ListAvailableFileSystemsResponseBodyFileSystemListMountTargetList) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *ListAvailableFileSystemsResponseBodyFileSystemListMountTargetList) GetVpcId() *string {
	return s.VpcId
}

func (s *ListAvailableFileSystemsResponseBodyFileSystemListMountTargetList) SetMountTargetDomain(v string) *ListAvailableFileSystemsResponseBodyFileSystemListMountTargetList {
	s.MountTargetDomain = &v
	return s
}

func (s *ListAvailableFileSystemsResponseBodyFileSystemListMountTargetList) SetNetworkType(v string) *ListAvailableFileSystemsResponseBodyFileSystemListMountTargetList {
	s.NetworkType = &v
	return s
}

func (s *ListAvailableFileSystemsResponseBodyFileSystemListMountTargetList) SetStatus(v string) *ListAvailableFileSystemsResponseBodyFileSystemListMountTargetList {
	s.Status = &v
	return s
}

func (s *ListAvailableFileSystemsResponseBodyFileSystemListMountTargetList) SetVSwitchId(v string) *ListAvailableFileSystemsResponseBodyFileSystemListMountTargetList {
	s.VSwitchId = &v
	return s
}

func (s *ListAvailableFileSystemsResponseBodyFileSystemListMountTargetList) SetVpcId(v string) *ListAvailableFileSystemsResponseBodyFileSystemListMountTargetList {
	s.VpcId = &v
	return s
}

func (s *ListAvailableFileSystemsResponseBodyFileSystemListMountTargetList) Validate() error {
	return dara.Validate(s)
}
