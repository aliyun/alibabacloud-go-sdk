// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListFileSystemsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetFileSystems(v []*ListFileSystemsResponseBodyFileSystems) *ListFileSystemsResponseBody
	GetFileSystems() []*ListFileSystemsResponseBodyFileSystems
	SetNextToken(v string) *ListFileSystemsResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListFileSystemsResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *ListFileSystemsResponseBody
	GetTotalCount() *int32
}

type ListFileSystemsResponseBody struct {
	FileSystems []*ListFileSystemsResponseBodyFileSystems `json:"FileSystems,omitempty" xml:"FileSystems,omitempty" type:"Repeated"`
	NextToken   *string                                   `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// example:
	//
	// 55C5FFD6-BF99-41BD-9C66-FFF39189****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 1
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListFileSystemsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListFileSystemsResponseBody) GoString() string {
	return s.String()
}

func (s *ListFileSystemsResponseBody) GetFileSystems() []*ListFileSystemsResponseBodyFileSystems {
	return s.FileSystems
}

func (s *ListFileSystemsResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListFileSystemsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListFileSystemsResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListFileSystemsResponseBody) SetFileSystems(v []*ListFileSystemsResponseBodyFileSystems) *ListFileSystemsResponseBody {
	s.FileSystems = v
	return s
}

func (s *ListFileSystemsResponseBody) SetNextToken(v string) *ListFileSystemsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListFileSystemsResponseBody) SetRequestId(v string) *ListFileSystemsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListFileSystemsResponseBody) SetTotalCount(v int32) *ListFileSystemsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListFileSystemsResponseBody) Validate() error {
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

type ListFileSystemsResponseBodyFileSystems struct {
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

func (s ListFileSystemsResponseBodyFileSystems) String() string {
	return dara.Prettify(s)
}

func (s ListFileSystemsResponseBodyFileSystems) GoString() string {
	return s.String()
}

func (s *ListFileSystemsResponseBodyFileSystems) GetCreateTime() *string {
	return s.CreateTime
}

func (s *ListFileSystemsResponseBodyFileSystems) GetDescription() *string {
	return s.Description
}

func (s *ListFileSystemsResponseBodyFileSystems) GetFileSystemId() *string {
	return s.FileSystemId
}

func (s *ListFileSystemsResponseBodyFileSystems) GetFileSystemName() *string {
	return s.FileSystemName
}

func (s *ListFileSystemsResponseBodyFileSystems) GetMeteringSpaceSize() *float32 {
	return s.MeteringSpaceSize
}

func (s *ListFileSystemsResponseBodyFileSystems) GetMountPointCount() *int64 {
	return s.MountPointCount
}

func (s *ListFileSystemsResponseBodyFileSystems) GetNumberOfDirectories() *int64 {
	return s.NumberOfDirectories
}

func (s *ListFileSystemsResponseBodyFileSystems) GetNumberOfFiles() *int64 {
	return s.NumberOfFiles
}

func (s *ListFileSystemsResponseBodyFileSystems) GetProtocolType() *string {
	return s.ProtocolType
}

func (s *ListFileSystemsResponseBodyFileSystems) GetProvisionedThroughputInMiBps() *int64 {
	return s.ProvisionedThroughputInMiBps
}

func (s *ListFileSystemsResponseBodyFileSystems) GetRegionId() *string {
	return s.RegionId
}

func (s *ListFileSystemsResponseBodyFileSystems) GetSpaceCapacity() *int64 {
	return s.SpaceCapacity
}

func (s *ListFileSystemsResponseBodyFileSystems) GetStoragePackageId() *string {
	return s.StoragePackageId
}

func (s *ListFileSystemsResponseBodyFileSystems) GetStorageType() *string {
	return s.StorageType
}

func (s *ListFileSystemsResponseBodyFileSystems) GetThroughputMode() *string {
	return s.ThroughputMode
}

func (s *ListFileSystemsResponseBodyFileSystems) GetUsedSpaceSize() *float32 {
	return s.UsedSpaceSize
}

func (s *ListFileSystemsResponseBodyFileSystems) GetVersion() *string {
	return s.Version
}

func (s *ListFileSystemsResponseBodyFileSystems) GetZoneId() *string {
	return s.ZoneId
}

func (s *ListFileSystemsResponseBodyFileSystems) SetCreateTime(v string) *ListFileSystemsResponseBodyFileSystems {
	s.CreateTime = &v
	return s
}

func (s *ListFileSystemsResponseBodyFileSystems) SetDescription(v string) *ListFileSystemsResponseBodyFileSystems {
	s.Description = &v
	return s
}

func (s *ListFileSystemsResponseBodyFileSystems) SetFileSystemId(v string) *ListFileSystemsResponseBodyFileSystems {
	s.FileSystemId = &v
	return s
}

func (s *ListFileSystemsResponseBodyFileSystems) SetFileSystemName(v string) *ListFileSystemsResponseBodyFileSystems {
	s.FileSystemName = &v
	return s
}

func (s *ListFileSystemsResponseBodyFileSystems) SetMeteringSpaceSize(v float32) *ListFileSystemsResponseBodyFileSystems {
	s.MeteringSpaceSize = &v
	return s
}

func (s *ListFileSystemsResponseBodyFileSystems) SetMountPointCount(v int64) *ListFileSystemsResponseBodyFileSystems {
	s.MountPointCount = &v
	return s
}

func (s *ListFileSystemsResponseBodyFileSystems) SetNumberOfDirectories(v int64) *ListFileSystemsResponseBodyFileSystems {
	s.NumberOfDirectories = &v
	return s
}

func (s *ListFileSystemsResponseBodyFileSystems) SetNumberOfFiles(v int64) *ListFileSystemsResponseBodyFileSystems {
	s.NumberOfFiles = &v
	return s
}

func (s *ListFileSystemsResponseBodyFileSystems) SetProtocolType(v string) *ListFileSystemsResponseBodyFileSystems {
	s.ProtocolType = &v
	return s
}

func (s *ListFileSystemsResponseBodyFileSystems) SetProvisionedThroughputInMiBps(v int64) *ListFileSystemsResponseBodyFileSystems {
	s.ProvisionedThroughputInMiBps = &v
	return s
}

func (s *ListFileSystemsResponseBodyFileSystems) SetRegionId(v string) *ListFileSystemsResponseBodyFileSystems {
	s.RegionId = &v
	return s
}

func (s *ListFileSystemsResponseBodyFileSystems) SetSpaceCapacity(v int64) *ListFileSystemsResponseBodyFileSystems {
	s.SpaceCapacity = &v
	return s
}

func (s *ListFileSystemsResponseBodyFileSystems) SetStoragePackageId(v string) *ListFileSystemsResponseBodyFileSystems {
	s.StoragePackageId = &v
	return s
}

func (s *ListFileSystemsResponseBodyFileSystems) SetStorageType(v string) *ListFileSystemsResponseBodyFileSystems {
	s.StorageType = &v
	return s
}

func (s *ListFileSystemsResponseBodyFileSystems) SetThroughputMode(v string) *ListFileSystemsResponseBodyFileSystems {
	s.ThroughputMode = &v
	return s
}

func (s *ListFileSystemsResponseBodyFileSystems) SetUsedSpaceSize(v float32) *ListFileSystemsResponseBodyFileSystems {
	s.UsedSpaceSize = &v
	return s
}

func (s *ListFileSystemsResponseBodyFileSystems) SetVersion(v string) *ListFileSystemsResponseBodyFileSystems {
	s.Version = &v
	return s
}

func (s *ListFileSystemsResponseBodyFileSystems) SetZoneId(v string) *ListFileSystemsResponseBodyFileSystems {
	s.ZoneId = &v
	return s
}

func (s *ListFileSystemsResponseBodyFileSystems) Validate() error {
	return dara.Validate(s)
}
