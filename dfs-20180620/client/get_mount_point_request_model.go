// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMountPointRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFileSystemId(v string) *GetMountPointRequest
	GetFileSystemId() *string
	SetInputRegionId(v string) *GetMountPointRequest
	GetInputRegionId() *string
	SetMountPointId(v string) *GetMountPointRequest
	GetMountPointId() *string
}

type GetMountPointRequest struct {
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
}

func (s GetMountPointRequest) String() string {
	return dara.Prettify(s)
}

func (s GetMountPointRequest) GoString() string {
	return s.String()
}

func (s *GetMountPointRequest) GetFileSystemId() *string {
	return s.FileSystemId
}

func (s *GetMountPointRequest) GetInputRegionId() *string {
	return s.InputRegionId
}

func (s *GetMountPointRequest) GetMountPointId() *string {
	return s.MountPointId
}

func (s *GetMountPointRequest) SetFileSystemId(v string) *GetMountPointRequest {
	s.FileSystemId = &v
	return s
}

func (s *GetMountPointRequest) SetInputRegionId(v string) *GetMountPointRequest {
	s.InputRegionId = &v
	return s
}

func (s *GetMountPointRequest) SetMountPointId(v string) *GetMountPointRequest {
	s.MountPointId = &v
	return s
}

func (s *GetMountPointRequest) Validate() error {
	return dara.Validate(s)
}
