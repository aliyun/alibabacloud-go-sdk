// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeFileSystemsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetFileSystems(v []*DescribeFileSystemsResponseBodyFileSystems) *DescribeFileSystemsResponseBody
	GetFileSystems() []*DescribeFileSystemsResponseBodyFileSystems
	SetPageNumber(v int32) *DescribeFileSystemsResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeFileSystemsResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *DescribeFileSystemsResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *DescribeFileSystemsResponseBody
	GetTotalCount() *int32
}

type DescribeFileSystemsResponseBody struct {
	// The information about the file systems.
	FileSystems []*DescribeFileSystemsResponseBodyFileSystems `json:"FileSystems,omitempty" xml:"FileSystems,omitempty" type:"Repeated"`
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
	// 2
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 762DD527-358B-1E48-8005-CCE3ED4EB9E0
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 1
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeFileSystemsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeFileSystemsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeFileSystemsResponseBody) GetFileSystems() []*DescribeFileSystemsResponseBodyFileSystems {
	return s.FileSystems
}

func (s *DescribeFileSystemsResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeFileSystemsResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeFileSystemsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeFileSystemsResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeFileSystemsResponseBody) SetFileSystems(v []*DescribeFileSystemsResponseBodyFileSystems) *DescribeFileSystemsResponseBody {
	s.FileSystems = v
	return s
}

func (s *DescribeFileSystemsResponseBody) SetPageNumber(v int32) *DescribeFileSystemsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeFileSystemsResponseBody) SetPageSize(v int32) *DescribeFileSystemsResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeFileSystemsResponseBody) SetRequestId(v string) *DescribeFileSystemsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeFileSystemsResponseBody) SetTotalCount(v int32) *DescribeFileSystemsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeFileSystemsResponseBody) Validate() error {
	if s.FileSystems != nil {
		for _, item := range s.FileSystems {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeFileSystemsResponseBodyFileSystems struct {
	// The capacity of the file system. Unit: MiB.
	//
	// example:
	//
	// 100000
	Capacity *int64 `json:"Capacity,omitempty" xml:"Capacity,omitempty"`
	// The time when the file system was created.
	//
	// example:
	//
	// 2022-08-31 12:00:00
	CreationTime *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	Description  *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The ID of the region.
	//
	// example:
	//
	// cn-beijing-cmcc
	EnsRegionId *string `json:"EnsRegionId,omitempty" xml:"EnsRegionId,omitempty"`
	// The ID of the file system.
	//
	// example:
	//
	// c50f8*****
	FileSystemId *string `json:"FileSystemId,omitempty" xml:"FileSystemId,omitempty"`
	// The name of the file system.
	//
	// example:
	//
	// FileSystem1
	FileSystemName *string `json:"FileSystemName,omitempty" xml:"FileSystemName,omitempty"`
	// The storage usage of the file system. The value of this parameter is the maximum storage usage of the file system over the last hour. Unit: bytes.
	//
	// example:
	//
	// 102400
	MeteredSize *int64 `json:"MeteredSize,omitempty" xml:"MeteredSize,omitempty"`
	// The information about mount targets.
	MountTargets []*DescribeFileSystemsResponseBodyFileSystemsMountTargets `json:"MountTargets,omitempty" xml:"MountTargets,omitempty" type:"Repeated"`
	// The billing method. PostPaid is returned. PostPaid indicates the pay-as-you-go billing method.
	//
	// example:
	//
	// PostPaid
	PayType *string `json:"PayType,omitempty" xml:"PayType,omitempty"`
	// The protocol type of the file system. Valid values:
	//
	// 	- NFS: Network File System (NFS)
	//
	// 	- SMB: Server Message Block (SMB)
	//
	// example:
	//
	// NFS
	ProtocolType *string `json:"ProtocolType,omitempty" xml:"ProtocolType,omitempty"`
	// The status of the file system. Valid values:
	//
	// 	- pending: The file system is being created or modified.
	//
	// 	- running: The file system is available. Before you create a mount target for the file system, make sure that the file system is in the running state.
	//
	// 	- stopped: The file system is unavailable.
	//
	// 	- extending: The file system is being scaled out.
	//
	// 	- stopping: The file system is being disabled.
	//
	// 	- deleting: The file system is being deleted.
	//
	// example:
	//
	// running
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The storage type. Valid values:
	//
	// 	- capacity: Capacity NAS file systems
	//
	// 	- performance: Performance NAS file systems
	//
	// example:
	//
	// capacity
	StorageType *string `json:"StorageType,omitempty" xml:"StorageType,omitempty"`
}

func (s DescribeFileSystemsResponseBodyFileSystems) String() string {
	return dara.Prettify(s)
}

func (s DescribeFileSystemsResponseBodyFileSystems) GoString() string {
	return s.String()
}

func (s *DescribeFileSystemsResponseBodyFileSystems) GetCapacity() *int64 {
	return s.Capacity
}

func (s *DescribeFileSystemsResponseBodyFileSystems) GetCreationTime() *string {
	return s.CreationTime
}

func (s *DescribeFileSystemsResponseBodyFileSystems) GetDescription() *string {
	return s.Description
}

func (s *DescribeFileSystemsResponseBodyFileSystems) GetEnsRegionId() *string {
	return s.EnsRegionId
}

func (s *DescribeFileSystemsResponseBodyFileSystems) GetFileSystemId() *string {
	return s.FileSystemId
}

func (s *DescribeFileSystemsResponseBodyFileSystems) GetFileSystemName() *string {
	return s.FileSystemName
}

func (s *DescribeFileSystemsResponseBodyFileSystems) GetMeteredSize() *int64 {
	return s.MeteredSize
}

func (s *DescribeFileSystemsResponseBodyFileSystems) GetMountTargets() []*DescribeFileSystemsResponseBodyFileSystemsMountTargets {
	return s.MountTargets
}

func (s *DescribeFileSystemsResponseBodyFileSystems) GetPayType() *string {
	return s.PayType
}

func (s *DescribeFileSystemsResponseBodyFileSystems) GetProtocolType() *string {
	return s.ProtocolType
}

func (s *DescribeFileSystemsResponseBodyFileSystems) GetStatus() *string {
	return s.Status
}

func (s *DescribeFileSystemsResponseBodyFileSystems) GetStorageType() *string {
	return s.StorageType
}

func (s *DescribeFileSystemsResponseBodyFileSystems) SetCapacity(v int64) *DescribeFileSystemsResponseBodyFileSystems {
	s.Capacity = &v
	return s
}

func (s *DescribeFileSystemsResponseBodyFileSystems) SetCreationTime(v string) *DescribeFileSystemsResponseBodyFileSystems {
	s.CreationTime = &v
	return s
}

func (s *DescribeFileSystemsResponseBodyFileSystems) SetDescription(v string) *DescribeFileSystemsResponseBodyFileSystems {
	s.Description = &v
	return s
}

func (s *DescribeFileSystemsResponseBodyFileSystems) SetEnsRegionId(v string) *DescribeFileSystemsResponseBodyFileSystems {
	s.EnsRegionId = &v
	return s
}

func (s *DescribeFileSystemsResponseBodyFileSystems) SetFileSystemId(v string) *DescribeFileSystemsResponseBodyFileSystems {
	s.FileSystemId = &v
	return s
}

func (s *DescribeFileSystemsResponseBodyFileSystems) SetFileSystemName(v string) *DescribeFileSystemsResponseBodyFileSystems {
	s.FileSystemName = &v
	return s
}

func (s *DescribeFileSystemsResponseBodyFileSystems) SetMeteredSize(v int64) *DescribeFileSystemsResponseBodyFileSystems {
	s.MeteredSize = &v
	return s
}

func (s *DescribeFileSystemsResponseBodyFileSystems) SetMountTargets(v []*DescribeFileSystemsResponseBodyFileSystemsMountTargets) *DescribeFileSystemsResponseBodyFileSystems {
	s.MountTargets = v
	return s
}

func (s *DescribeFileSystemsResponseBodyFileSystems) SetPayType(v string) *DescribeFileSystemsResponseBodyFileSystems {
	s.PayType = &v
	return s
}

func (s *DescribeFileSystemsResponseBodyFileSystems) SetProtocolType(v string) *DescribeFileSystemsResponseBodyFileSystems {
	s.ProtocolType = &v
	return s
}

func (s *DescribeFileSystemsResponseBodyFileSystems) SetStatus(v string) *DescribeFileSystemsResponseBodyFileSystems {
	s.Status = &v
	return s
}

func (s *DescribeFileSystemsResponseBodyFileSystems) SetStorageType(v string) *DescribeFileSystemsResponseBodyFileSystems {
	s.StorageType = &v
	return s
}

func (s *DescribeFileSystemsResponseBodyFileSystems) Validate() error {
	if s.MountTargets != nil {
		for _, item := range s.MountTargets {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeFileSystemsResponseBodyFileSystemsMountTargets struct {
	// The path of the mount target.
	//
	// example:
	//
	// *.*.*.*:/${FileSystemName}/{MountTargetName}
	MountTargetDomain *string `json:"MountTargetDomain,omitempty" xml:"MountTargetDomain,omitempty"`
	// The name of the mount target.
	//
	// example:
	//
	// target1
	MountTargetName *string `json:"MountTargetName,omitempty" xml:"MountTargetName,omitempty"`
	// The ID of the network.
	//
	// example:
	//
	// n-****
	NetWorkId *string `json:"NetWorkId,omitempty" xml:"NetWorkId,omitempty"`
	// The status of the mount target. Valid values:
	//
	// 	- active: The mount target is available.
	//
	// 	- inactive: The mount target is unavailable.
	//
	// 	- pending: The task is running.
	//
	// 	- deleting: The mount target is being deleted.
	//
	// example:
	//
	// active
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeFileSystemsResponseBodyFileSystemsMountTargets) String() string {
	return dara.Prettify(s)
}

func (s DescribeFileSystemsResponseBodyFileSystemsMountTargets) GoString() string {
	return s.String()
}

func (s *DescribeFileSystemsResponseBodyFileSystemsMountTargets) GetMountTargetDomain() *string {
	return s.MountTargetDomain
}

func (s *DescribeFileSystemsResponseBodyFileSystemsMountTargets) GetMountTargetName() *string {
	return s.MountTargetName
}

func (s *DescribeFileSystemsResponseBodyFileSystemsMountTargets) GetNetWorkId() *string {
	return s.NetWorkId
}

func (s *DescribeFileSystemsResponseBodyFileSystemsMountTargets) GetStatus() *string {
	return s.Status
}

func (s *DescribeFileSystemsResponseBodyFileSystemsMountTargets) SetMountTargetDomain(v string) *DescribeFileSystemsResponseBodyFileSystemsMountTargets {
	s.MountTargetDomain = &v
	return s
}

func (s *DescribeFileSystemsResponseBodyFileSystemsMountTargets) SetMountTargetName(v string) *DescribeFileSystemsResponseBodyFileSystemsMountTargets {
	s.MountTargetName = &v
	return s
}

func (s *DescribeFileSystemsResponseBodyFileSystemsMountTargets) SetNetWorkId(v string) *DescribeFileSystemsResponseBodyFileSystemsMountTargets {
	s.NetWorkId = &v
	return s
}

func (s *DescribeFileSystemsResponseBodyFileSystemsMountTargets) SetStatus(v string) *DescribeFileSystemsResponseBodyFileSystemsMountTargets {
	s.Status = &v
	return s
}

func (s *DescribeFileSystemsResponseBodyFileSystemsMountTargets) Validate() error {
	return dara.Validate(s)
}
