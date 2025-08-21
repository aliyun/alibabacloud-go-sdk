// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeMountTargetsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMountTargets(v []*DescribeMountTargetsResponseBodyMountTargets) *DescribeMountTargetsResponseBody
	GetMountTargets() []*DescribeMountTargetsResponseBodyMountTargets
	SetPageNumber(v int32) *DescribeMountTargetsResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeMountTargetsResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *DescribeMountTargetsResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *DescribeMountTargetsResponseBody
	GetTotalCount() *int32
}

type DescribeMountTargetsResponseBody struct {
	// The information about mount targets.
	MountTargets []*DescribeMountTargetsResponseBodyMountTargets `json:"MountTargets,omitempty" xml:"MountTargets,omitempty" type:"Repeated"`
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
	// The total number of mount targets.
	//
	// example:
	//
	// 1
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeMountTargetsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeMountTargetsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeMountTargetsResponseBody) GetMountTargets() []*DescribeMountTargetsResponseBodyMountTargets {
	return s.MountTargets
}

func (s *DescribeMountTargetsResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeMountTargetsResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeMountTargetsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeMountTargetsResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeMountTargetsResponseBody) SetMountTargets(v []*DescribeMountTargetsResponseBodyMountTargets) *DescribeMountTargetsResponseBody {
	s.MountTargets = v
	return s
}

func (s *DescribeMountTargetsResponseBody) SetPageNumber(v int32) *DescribeMountTargetsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeMountTargetsResponseBody) SetPageSize(v int32) *DescribeMountTargetsResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeMountTargetsResponseBody) SetRequestId(v string) *DescribeMountTargetsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeMountTargetsResponseBody) SetTotalCount(v int32) *DescribeMountTargetsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeMountTargetsResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeMountTargetsResponseBodyMountTargets struct {
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
	// The path of the mount target.
	//
	// example:
	//
	// LB:/fileSystemName/mountTargetName
	MountTargetDomain *string `json:"MountTargetDomain,omitempty" xml:"MountTargetDomain,omitempty"`
	// The name of the mount target.
	//
	// example:
	//
	// TestMountPath
	MountTargetName *string `json:"MountTargetName,omitempty" xml:"MountTargetName,omitempty"`
	// The ID of the network.
	//
	// example:
	//
	// n-***
	NetWorkId *string `json:"NetWorkId,omitempty" xml:"NetWorkId,omitempty"`
	// The state of the mount target. Valid values:
	//
	// 	- active: The mount target is available.
	//
	// 	- inactive: The mount target is unavailable.
	//
	// 	- pending: A task is being queued for the mount target.
	//
	// 	- deleting: The mount target is being deleted.
	//
	// example:
	//
	// active
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeMountTargetsResponseBodyMountTargets) String() string {
	return dara.Prettify(s)
}

func (s DescribeMountTargetsResponseBodyMountTargets) GoString() string {
	return s.String()
}

func (s *DescribeMountTargetsResponseBodyMountTargets) GetEnsRegionId() *string {
	return s.EnsRegionId
}

func (s *DescribeMountTargetsResponseBodyMountTargets) GetFileSystemId() *string {
	return s.FileSystemId
}

func (s *DescribeMountTargetsResponseBodyMountTargets) GetMountTargetDomain() *string {
	return s.MountTargetDomain
}

func (s *DescribeMountTargetsResponseBodyMountTargets) GetMountTargetName() *string {
	return s.MountTargetName
}

func (s *DescribeMountTargetsResponseBodyMountTargets) GetNetWorkId() *string {
	return s.NetWorkId
}

func (s *DescribeMountTargetsResponseBodyMountTargets) GetStatus() *string {
	return s.Status
}

func (s *DescribeMountTargetsResponseBodyMountTargets) SetEnsRegionId(v string) *DescribeMountTargetsResponseBodyMountTargets {
	s.EnsRegionId = &v
	return s
}

func (s *DescribeMountTargetsResponseBodyMountTargets) SetFileSystemId(v string) *DescribeMountTargetsResponseBodyMountTargets {
	s.FileSystemId = &v
	return s
}

func (s *DescribeMountTargetsResponseBodyMountTargets) SetMountTargetDomain(v string) *DescribeMountTargetsResponseBodyMountTargets {
	s.MountTargetDomain = &v
	return s
}

func (s *DescribeMountTargetsResponseBodyMountTargets) SetMountTargetName(v string) *DescribeMountTargetsResponseBodyMountTargets {
	s.MountTargetName = &v
	return s
}

func (s *DescribeMountTargetsResponseBodyMountTargets) SetNetWorkId(v string) *DescribeMountTargetsResponseBodyMountTargets {
	s.NetWorkId = &v
	return s
}

func (s *DescribeMountTargetsResponseBodyMountTargets) SetStatus(v string) *DescribeMountTargetsResponseBodyMountTargets {
	s.Status = &v
	return s
}

func (s *DescribeMountTargetsResponseBodyMountTargets) Validate() error {
	return dara.Validate(s)
}
