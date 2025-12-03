// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateFileSystemRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDataRedundancyType(v string) *CreateFileSystemRequest
	GetDataRedundancyType() *string
	SetDedicatedClusterId(v string) *CreateFileSystemRequest
	GetDedicatedClusterId() *string
	SetDescription(v string) *CreateFileSystemRequest
	GetDescription() *string
	SetFileSystemName(v string) *CreateFileSystemRequest
	GetFileSystemName() *string
	SetInputRegionId(v string) *CreateFileSystemRequest
	GetInputRegionId() *string
	SetPartitionNumber(v int32) *CreateFileSystemRequest
	GetPartitionNumber() *int32
	SetProtocolType(v string) *CreateFileSystemRequest
	GetProtocolType() *string
	SetProvisionedThroughputInMiBps(v int64) *CreateFileSystemRequest
	GetProvisionedThroughputInMiBps() *int64
	SetSpaceCapacity(v int64) *CreateFileSystemRequest
	GetSpaceCapacity() *int64
	SetStorageSetName(v string) *CreateFileSystemRequest
	GetStorageSetName() *string
	SetStorageType(v string) *CreateFileSystemRequest
	GetStorageType() *string
	SetThroughputMode(v string) *CreateFileSystemRequest
	GetThroughputMode() *string
	SetZoneId(v string) *CreateFileSystemRequest
	GetZoneId() *string
}

type CreateFileSystemRequest struct {
	// example:
	//
	// LRS
	DataRedundancyType *string `json:"DataRedundancyType,omitempty" xml:"DataRedundancyType,omitempty"`
	DedicatedClusterId *string `json:"DedicatedClusterId,omitempty" xml:"DedicatedClusterId,omitempty"`
	Description        *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// MyFirstHDFS
	FileSystemName *string `json:"FileSystemName,omitempty" xml:"FileSystemName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	InputRegionId *string `json:"InputRegionId,omitempty" xml:"InputRegionId,omitempty"`
	// example:
	//
	// 0
	PartitionNumber *int32 `json:"PartitionNumber,omitempty" xml:"PartitionNumber,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// HDFS
	ProtocolType *string `json:"ProtocolType,omitempty" xml:"ProtocolType,omitempty"`
	// example:
	//
	// 1024
	ProvisionedThroughputInMiBps *int64 `json:"ProvisionedThroughputInMiBps,omitempty" xml:"ProvisionedThroughputInMiBps,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1024
	SpaceCapacity *int64 `json:"SpaceCapacity,omitempty" xml:"SpaceCapacity,omitempty"`
	// example:
	//
	// AligroupStorageSet
	StorageSetName *string `json:"StorageSetName,omitempty" xml:"StorageSetName,omitempty"`
	// This parameter is required.
	//
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
	// cn-hangzhou-b
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s CreateFileSystemRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateFileSystemRequest) GoString() string {
	return s.String()
}

func (s *CreateFileSystemRequest) GetDataRedundancyType() *string {
	return s.DataRedundancyType
}

func (s *CreateFileSystemRequest) GetDedicatedClusterId() *string {
	return s.DedicatedClusterId
}

func (s *CreateFileSystemRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateFileSystemRequest) GetFileSystemName() *string {
	return s.FileSystemName
}

func (s *CreateFileSystemRequest) GetInputRegionId() *string {
	return s.InputRegionId
}

func (s *CreateFileSystemRequest) GetPartitionNumber() *int32 {
	return s.PartitionNumber
}

func (s *CreateFileSystemRequest) GetProtocolType() *string {
	return s.ProtocolType
}

func (s *CreateFileSystemRequest) GetProvisionedThroughputInMiBps() *int64 {
	return s.ProvisionedThroughputInMiBps
}

func (s *CreateFileSystemRequest) GetSpaceCapacity() *int64 {
	return s.SpaceCapacity
}

func (s *CreateFileSystemRequest) GetStorageSetName() *string {
	return s.StorageSetName
}

func (s *CreateFileSystemRequest) GetStorageType() *string {
	return s.StorageType
}

func (s *CreateFileSystemRequest) GetThroughputMode() *string {
	return s.ThroughputMode
}

func (s *CreateFileSystemRequest) GetZoneId() *string {
	return s.ZoneId
}

func (s *CreateFileSystemRequest) SetDataRedundancyType(v string) *CreateFileSystemRequest {
	s.DataRedundancyType = &v
	return s
}

func (s *CreateFileSystemRequest) SetDedicatedClusterId(v string) *CreateFileSystemRequest {
	s.DedicatedClusterId = &v
	return s
}

func (s *CreateFileSystemRequest) SetDescription(v string) *CreateFileSystemRequest {
	s.Description = &v
	return s
}

func (s *CreateFileSystemRequest) SetFileSystemName(v string) *CreateFileSystemRequest {
	s.FileSystemName = &v
	return s
}

func (s *CreateFileSystemRequest) SetInputRegionId(v string) *CreateFileSystemRequest {
	s.InputRegionId = &v
	return s
}

func (s *CreateFileSystemRequest) SetPartitionNumber(v int32) *CreateFileSystemRequest {
	s.PartitionNumber = &v
	return s
}

func (s *CreateFileSystemRequest) SetProtocolType(v string) *CreateFileSystemRequest {
	s.ProtocolType = &v
	return s
}

func (s *CreateFileSystemRequest) SetProvisionedThroughputInMiBps(v int64) *CreateFileSystemRequest {
	s.ProvisionedThroughputInMiBps = &v
	return s
}

func (s *CreateFileSystemRequest) SetSpaceCapacity(v int64) *CreateFileSystemRequest {
	s.SpaceCapacity = &v
	return s
}

func (s *CreateFileSystemRequest) SetStorageSetName(v string) *CreateFileSystemRequest {
	s.StorageSetName = &v
	return s
}

func (s *CreateFileSystemRequest) SetStorageType(v string) *CreateFileSystemRequest {
	s.StorageType = &v
	return s
}

func (s *CreateFileSystemRequest) SetThroughputMode(v string) *CreateFileSystemRequest {
	s.ThroughputMode = &v
	return s
}

func (s *CreateFileSystemRequest) SetZoneId(v string) *CreateFileSystemRequest {
	s.ZoneId = &v
	return s
}

func (s *CreateFileSystemRequest) Validate() error {
	return dara.Validate(s)
}
