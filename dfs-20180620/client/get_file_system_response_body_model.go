// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetFileSystemResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetFileSystem(v *GetFileSystemResponseBodyFileSystem) *GetFileSystemResponseBody
	GetFileSystem() *GetFileSystemResponseBodyFileSystem
	SetRequestId(v string) *GetFileSystemResponseBody
	GetRequestId() *string
}

type GetFileSystemResponseBody struct {
	FileSystem *GetFileSystemResponseBodyFileSystem `json:"FileSystem,omitempty" xml:"FileSystem,omitempty" type:"Struct"`
	// example:
	//
	// 55C5FFD6-BF99-41BD-9C66-FFF39189****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetFileSystemResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetFileSystemResponseBody) GoString() string {
	return s.String()
}

func (s *GetFileSystemResponseBody) GetFileSystem() *GetFileSystemResponseBodyFileSystem {
	return s.FileSystem
}

func (s *GetFileSystemResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetFileSystemResponseBody) SetFileSystem(v *GetFileSystemResponseBodyFileSystem) *GetFileSystemResponseBody {
	s.FileSystem = v
	return s
}

func (s *GetFileSystemResponseBody) SetRequestId(v string) *GetFileSystemResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetFileSystemResponseBody) Validate() error {
	if s.FileSystem != nil {
		if err := s.FileSystem.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetFileSystemResponseBodyFileSystem struct {
	CreateTime   *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Description  *string `json:"Description,omitempty" xml:"Description,omitempty"`
	FileSystemId *string `json:"FileSystemId,omitempty" xml:"FileSystemId,omitempty"`
	// example:
	//
	// MyFirstHDFS
	FileSystemName *string `json:"FileSystemName,omitempty" xml:"FileSystemName,omitempty"`
	// example:
	//
	// 107374182400
	MeteringSpaceSize *float32 `json:"MeteringSpaceSize,omitempty" xml:"MeteringSpaceSize,omitempty"`
	// example:
	//
	// 2
	MountPointCount *int64 `json:"MountPointCount,omitempty" xml:"MountPointCount,omitempty"`
	// example:
	//
	// 100
	NumberOfDirectories *int64 `json:"NumberOfDirectories,omitempty" xml:"NumberOfDirectories,omitempty"`
	// example:
	//
	// 1000
	NumberOfFiles *int64 `json:"NumberOfFiles,omitempty" xml:"NumberOfFiles,omitempty"`
	// example:
	//
	// HDFS
	ProtocolType *string `json:"ProtocolType,omitempty" xml:"ProtocolType,omitempty"`
	// example:
	//
	// 1024
	ProvisionedThroughputInMiBps *int64 `json:"ProvisionedThroughputInMiBps,omitempty" xml:"ProvisionedThroughputInMiBps,omitempty"`
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// 1024
	SpaceCapacity *int64 `json:"SpaceCapacity,omitempty" xml:"SpaceCapacity,omitempty"`
	// example:
	//
	// 1
	StoragePackageId *string `json:"StoragePackageId,omitempty" xml:"StoragePackageId,omitempty"`
	// example:
	//
	// STANDARD
	StorageType *string `json:"StorageType,omitempty" xml:"StorageType,omitempty"`
	// example:
	//
	// Standard
	ThroughputMode *string `json:"ThroughputMode,omitempty" xml:"ThroughputMode,omitempty"`
	// example:
	//
	// 100
	UsedSpaceSize *float32 `json:"UsedSpaceSize,omitempty" xml:"UsedSpaceSize,omitempty"`
	// example:
	//
	// 1.0.0
	Version *string `json:"Version,omitempty" xml:"Version,omitempty"`
	// example:
	//
	// cn-hangzhou-b
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s GetFileSystemResponseBodyFileSystem) String() string {
	return dara.Prettify(s)
}

func (s GetFileSystemResponseBodyFileSystem) GoString() string {
	return s.String()
}

func (s *GetFileSystemResponseBodyFileSystem) GetCreateTime() *string {
	return s.CreateTime
}

func (s *GetFileSystemResponseBodyFileSystem) GetDescription() *string {
	return s.Description
}

func (s *GetFileSystemResponseBodyFileSystem) GetFileSystemId() *string {
	return s.FileSystemId
}

func (s *GetFileSystemResponseBodyFileSystem) GetFileSystemName() *string {
	return s.FileSystemName
}

func (s *GetFileSystemResponseBodyFileSystem) GetMeteringSpaceSize() *float32 {
	return s.MeteringSpaceSize
}

func (s *GetFileSystemResponseBodyFileSystem) GetMountPointCount() *int64 {
	return s.MountPointCount
}

func (s *GetFileSystemResponseBodyFileSystem) GetNumberOfDirectories() *int64 {
	return s.NumberOfDirectories
}

func (s *GetFileSystemResponseBodyFileSystem) GetNumberOfFiles() *int64 {
	return s.NumberOfFiles
}

func (s *GetFileSystemResponseBodyFileSystem) GetProtocolType() *string {
	return s.ProtocolType
}

func (s *GetFileSystemResponseBodyFileSystem) GetProvisionedThroughputInMiBps() *int64 {
	return s.ProvisionedThroughputInMiBps
}

func (s *GetFileSystemResponseBodyFileSystem) GetRegionId() *string {
	return s.RegionId
}

func (s *GetFileSystemResponseBodyFileSystem) GetSpaceCapacity() *int64 {
	return s.SpaceCapacity
}

func (s *GetFileSystemResponseBodyFileSystem) GetStoragePackageId() *string {
	return s.StoragePackageId
}

func (s *GetFileSystemResponseBodyFileSystem) GetStorageType() *string {
	return s.StorageType
}

func (s *GetFileSystemResponseBodyFileSystem) GetThroughputMode() *string {
	return s.ThroughputMode
}

func (s *GetFileSystemResponseBodyFileSystem) GetUsedSpaceSize() *float32 {
	return s.UsedSpaceSize
}

func (s *GetFileSystemResponseBodyFileSystem) GetVersion() *string {
	return s.Version
}

func (s *GetFileSystemResponseBodyFileSystem) GetZoneId() *string {
	return s.ZoneId
}

func (s *GetFileSystemResponseBodyFileSystem) SetCreateTime(v string) *GetFileSystemResponseBodyFileSystem {
	s.CreateTime = &v
	return s
}

func (s *GetFileSystemResponseBodyFileSystem) SetDescription(v string) *GetFileSystemResponseBodyFileSystem {
	s.Description = &v
	return s
}

func (s *GetFileSystemResponseBodyFileSystem) SetFileSystemId(v string) *GetFileSystemResponseBodyFileSystem {
	s.FileSystemId = &v
	return s
}

func (s *GetFileSystemResponseBodyFileSystem) SetFileSystemName(v string) *GetFileSystemResponseBodyFileSystem {
	s.FileSystemName = &v
	return s
}

func (s *GetFileSystemResponseBodyFileSystem) SetMeteringSpaceSize(v float32) *GetFileSystemResponseBodyFileSystem {
	s.MeteringSpaceSize = &v
	return s
}

func (s *GetFileSystemResponseBodyFileSystem) SetMountPointCount(v int64) *GetFileSystemResponseBodyFileSystem {
	s.MountPointCount = &v
	return s
}

func (s *GetFileSystemResponseBodyFileSystem) SetNumberOfDirectories(v int64) *GetFileSystemResponseBodyFileSystem {
	s.NumberOfDirectories = &v
	return s
}

func (s *GetFileSystemResponseBodyFileSystem) SetNumberOfFiles(v int64) *GetFileSystemResponseBodyFileSystem {
	s.NumberOfFiles = &v
	return s
}

func (s *GetFileSystemResponseBodyFileSystem) SetProtocolType(v string) *GetFileSystemResponseBodyFileSystem {
	s.ProtocolType = &v
	return s
}

func (s *GetFileSystemResponseBodyFileSystem) SetProvisionedThroughputInMiBps(v int64) *GetFileSystemResponseBodyFileSystem {
	s.ProvisionedThroughputInMiBps = &v
	return s
}

func (s *GetFileSystemResponseBodyFileSystem) SetRegionId(v string) *GetFileSystemResponseBodyFileSystem {
	s.RegionId = &v
	return s
}

func (s *GetFileSystemResponseBodyFileSystem) SetSpaceCapacity(v int64) *GetFileSystemResponseBodyFileSystem {
	s.SpaceCapacity = &v
	return s
}

func (s *GetFileSystemResponseBodyFileSystem) SetStoragePackageId(v string) *GetFileSystemResponseBodyFileSystem {
	s.StoragePackageId = &v
	return s
}

func (s *GetFileSystemResponseBodyFileSystem) SetStorageType(v string) *GetFileSystemResponseBodyFileSystem {
	s.StorageType = &v
	return s
}

func (s *GetFileSystemResponseBodyFileSystem) SetThroughputMode(v string) *GetFileSystemResponseBodyFileSystem {
	s.ThroughputMode = &v
	return s
}

func (s *GetFileSystemResponseBodyFileSystem) SetUsedSpaceSize(v float32) *GetFileSystemResponseBodyFileSystem {
	s.UsedSpaceSize = &v
	return s
}

func (s *GetFileSystemResponseBodyFileSystem) SetVersion(v string) *GetFileSystemResponseBodyFileSystem {
	s.Version = &v
	return s
}

func (s *GetFileSystemResponseBodyFileSystem) SetZoneId(v string) *GetFileSystemResponseBodyFileSystem {
	s.ZoneId = &v
	return s
}

func (s *GetFileSystemResponseBodyFileSystem) Validate() error {
	return dara.Validate(s)
}
