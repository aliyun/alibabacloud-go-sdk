// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyMountPointRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccessGroupId(v string) *ModifyMountPointRequest
	GetAccessGroupId() *string
	SetDescription(v string) *ModifyMountPointRequest
	GetDescription() *string
	SetFileSystemId(v string) *ModifyMountPointRequest
	GetFileSystemId() *string
	SetInputRegionId(v string) *ModifyMountPointRequest
	GetInputRegionId() *string
	SetMountPointId(v string) *ModifyMountPointRequest
	GetMountPointId() *string
	SetStatus(v string) *ModifyMountPointRequest
	GetStatus() *string
}

type ModifyMountPointRequest struct {
	// example:
	//
	// acg-e3755fb0-358d-4286-9942-8d461048****
	AccessGroupId *string `json:"AccessGroupId,omitempty" xml:"AccessGroupId,omitempty"`
	Description   *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// This parameter is required.
	FileSystemId *string `json:"FileSystemId,omitempty" xml:"FileSystemId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	InputRegionId *string `json:"InputRegionId,omitempty" xml:"InputRegionId,omitempty"`
	// This parameter is required.
	MountPointId *string `json:"MountPointId,omitempty" xml:"MountPointId,omitempty"`
	// example:
	//
	// Inactive
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ModifyMountPointRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyMountPointRequest) GoString() string {
	return s.String()
}

func (s *ModifyMountPointRequest) GetAccessGroupId() *string {
	return s.AccessGroupId
}

func (s *ModifyMountPointRequest) GetDescription() *string {
	return s.Description
}

func (s *ModifyMountPointRequest) GetFileSystemId() *string {
	return s.FileSystemId
}

func (s *ModifyMountPointRequest) GetInputRegionId() *string {
	return s.InputRegionId
}

func (s *ModifyMountPointRequest) GetMountPointId() *string {
	return s.MountPointId
}

func (s *ModifyMountPointRequest) GetStatus() *string {
	return s.Status
}

func (s *ModifyMountPointRequest) SetAccessGroupId(v string) *ModifyMountPointRequest {
	s.AccessGroupId = &v
	return s
}

func (s *ModifyMountPointRequest) SetDescription(v string) *ModifyMountPointRequest {
	s.Description = &v
	return s
}

func (s *ModifyMountPointRequest) SetFileSystemId(v string) *ModifyMountPointRequest {
	s.FileSystemId = &v
	return s
}

func (s *ModifyMountPointRequest) SetInputRegionId(v string) *ModifyMountPointRequest {
	s.InputRegionId = &v
	return s
}

func (s *ModifyMountPointRequest) SetMountPointId(v string) *ModifyMountPointRequest {
	s.MountPointId = &v
	return s
}

func (s *ModifyMountPointRequest) SetStatus(v string) *ModifyMountPointRequest {
	s.Status = &v
	return s
}

func (s *ModifyMountPointRequest) Validate() error {
	return dara.Validate(s)
}
