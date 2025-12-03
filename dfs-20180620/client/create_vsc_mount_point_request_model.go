// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateVscMountPointRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *CreateVscMountPointRequest
	GetDescription() *string
	SetFileSystemId(v string) *CreateVscMountPointRequest
	GetFileSystemId() *string
	SetInputRegionId(v string) *CreateVscMountPointRequest
	GetInputRegionId() *string
	SetInstanceIds(v []*string) *CreateVscMountPointRequest
	GetInstanceIds() []*string
}

type CreateVscMountPointRequest struct {
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// e389e5c7-bcb4-4558-846a-e5afc444****
	FileSystemId *string `json:"FileSystemId,omitempty" xml:"FileSystemId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	InputRegionId *string `json:"InputRegionId,omitempty" xml:"InputRegionId,omitempty"`
	// example:
	//
	// [
	//
	//       "i-bp1g6zv0ce8oghu7****",
	//
	//       "i-bp1g6zv0ce8oghu1****"
	//
	// ]
	InstanceIds []*string `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty" type:"Repeated"`
}

func (s CreateVscMountPointRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateVscMountPointRequest) GoString() string {
	return s.String()
}

func (s *CreateVscMountPointRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateVscMountPointRequest) GetFileSystemId() *string {
	return s.FileSystemId
}

func (s *CreateVscMountPointRequest) GetInputRegionId() *string {
	return s.InputRegionId
}

func (s *CreateVscMountPointRequest) GetInstanceIds() []*string {
	return s.InstanceIds
}

func (s *CreateVscMountPointRequest) SetDescription(v string) *CreateVscMountPointRequest {
	s.Description = &v
	return s
}

func (s *CreateVscMountPointRequest) SetFileSystemId(v string) *CreateVscMountPointRequest {
	s.FileSystemId = &v
	return s
}

func (s *CreateVscMountPointRequest) SetInputRegionId(v string) *CreateVscMountPointRequest {
	s.InputRegionId = &v
	return s
}

func (s *CreateVscMountPointRequest) SetInstanceIds(v []*string) *CreateVscMountPointRequest {
	s.InstanceIds = v
	return s
}

func (s *CreateVscMountPointRequest) Validate() error {
	return dara.Validate(s)
}
