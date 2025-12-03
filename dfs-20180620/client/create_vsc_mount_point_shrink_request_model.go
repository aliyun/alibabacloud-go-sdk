// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateVscMountPointShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *CreateVscMountPointShrinkRequest
	GetDescription() *string
	SetFileSystemId(v string) *CreateVscMountPointShrinkRequest
	GetFileSystemId() *string
	SetInputRegionId(v string) *CreateVscMountPointShrinkRequest
	GetInputRegionId() *string
	SetInstanceIdsShrink(v string) *CreateVscMountPointShrinkRequest
	GetInstanceIdsShrink() *string
}

type CreateVscMountPointShrinkRequest struct {
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
	InstanceIdsShrink *string `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty"`
}

func (s CreateVscMountPointShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateVscMountPointShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateVscMountPointShrinkRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateVscMountPointShrinkRequest) GetFileSystemId() *string {
	return s.FileSystemId
}

func (s *CreateVscMountPointShrinkRequest) GetInputRegionId() *string {
	return s.InputRegionId
}

func (s *CreateVscMountPointShrinkRequest) GetInstanceIdsShrink() *string {
	return s.InstanceIdsShrink
}

func (s *CreateVscMountPointShrinkRequest) SetDescription(v string) *CreateVscMountPointShrinkRequest {
	s.Description = &v
	return s
}

func (s *CreateVscMountPointShrinkRequest) SetFileSystemId(v string) *CreateVscMountPointShrinkRequest {
	s.FileSystemId = &v
	return s
}

func (s *CreateVscMountPointShrinkRequest) SetInputRegionId(v string) *CreateVscMountPointShrinkRequest {
	s.InputRegionId = &v
	return s
}

func (s *CreateVscMountPointShrinkRequest) SetInstanceIdsShrink(v string) *CreateVscMountPointShrinkRequest {
	s.InstanceIdsShrink = &v
	return s
}

func (s *CreateVscMountPointShrinkRequest) Validate() error {
	return dara.Validate(s)
}
