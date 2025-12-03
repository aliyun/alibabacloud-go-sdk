// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyFileSystemRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *ModifyFileSystemRequest
	GetDescription() *string
	SetFileSystemId(v string) *ModifyFileSystemRequest
	GetFileSystemId() *string
	SetFileSystemName(v string) *ModifyFileSystemRequest
	GetFileSystemName() *string
	SetInputRegionId(v string) *ModifyFileSystemRequest
	GetInputRegionId() *string
	SetProvisionedThroughputInMiBps(v int64) *ModifyFileSystemRequest
	GetProvisionedThroughputInMiBps() *int64
	SetSpaceCapacity(v int64) *ModifyFileSystemRequest
	GetSpaceCapacity() *int64
	SetThroughputMode(v string) *ModifyFileSystemRequest
	GetThroughputMode() *string
}

type ModifyFileSystemRequest struct {
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// This parameter is required.
	FileSystemId *string `json:"FileSystemId,omitempty" xml:"FileSystemId,omitempty"`
	// example:
	//
	// MyFirstModDFS
	FileSystemName *string `json:"FileSystemName,omitempty" xml:"FileSystemName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	InputRegionId *string `json:"InputRegionId,omitempty" xml:"InputRegionId,omitempty"`
	// example:
	//
	// 1024
	ProvisionedThroughputInMiBps *int64 `json:"ProvisionedThroughputInMiBps,omitempty" xml:"ProvisionedThroughputInMiBps,omitempty"`
	// example:
	//
	// 1024
	SpaceCapacity *int64 `json:"SpaceCapacity,omitempty" xml:"SpaceCapacity,omitempty"`
	// example:
	//
	// Standard
	ThroughputMode *string `json:"ThroughputMode,omitempty" xml:"ThroughputMode,omitempty"`
}

func (s ModifyFileSystemRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyFileSystemRequest) GoString() string {
	return s.String()
}

func (s *ModifyFileSystemRequest) GetDescription() *string {
	return s.Description
}

func (s *ModifyFileSystemRequest) GetFileSystemId() *string {
	return s.FileSystemId
}

func (s *ModifyFileSystemRequest) GetFileSystemName() *string {
	return s.FileSystemName
}

func (s *ModifyFileSystemRequest) GetInputRegionId() *string {
	return s.InputRegionId
}

func (s *ModifyFileSystemRequest) GetProvisionedThroughputInMiBps() *int64 {
	return s.ProvisionedThroughputInMiBps
}

func (s *ModifyFileSystemRequest) GetSpaceCapacity() *int64 {
	return s.SpaceCapacity
}

func (s *ModifyFileSystemRequest) GetThroughputMode() *string {
	return s.ThroughputMode
}

func (s *ModifyFileSystemRequest) SetDescription(v string) *ModifyFileSystemRequest {
	s.Description = &v
	return s
}

func (s *ModifyFileSystemRequest) SetFileSystemId(v string) *ModifyFileSystemRequest {
	s.FileSystemId = &v
	return s
}

func (s *ModifyFileSystemRequest) SetFileSystemName(v string) *ModifyFileSystemRequest {
	s.FileSystemName = &v
	return s
}

func (s *ModifyFileSystemRequest) SetInputRegionId(v string) *ModifyFileSystemRequest {
	s.InputRegionId = &v
	return s
}

func (s *ModifyFileSystemRequest) SetProvisionedThroughputInMiBps(v int64) *ModifyFileSystemRequest {
	s.ProvisionedThroughputInMiBps = &v
	return s
}

func (s *ModifyFileSystemRequest) SetSpaceCapacity(v int64) *ModifyFileSystemRequest {
	s.SpaceCapacity = &v
	return s
}

func (s *ModifyFileSystemRequest) SetThroughputMode(v string) *ModifyFileSystemRequest {
	s.ThroughputMode = &v
	return s
}

func (s *ModifyFileSystemRequest) Validate() error {
	return dara.Validate(s)
}
