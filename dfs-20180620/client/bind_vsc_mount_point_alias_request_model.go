// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBindVscMountPointAliasRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAliasPrefix(v string) *BindVscMountPointAliasRequest
	GetAliasPrefix() *string
	SetFileSystemId(v string) *BindVscMountPointAliasRequest
	GetFileSystemId() *string
	SetInputRegionId(v string) *BindVscMountPointAliasRequest
	GetInputRegionId() *string
	SetMountPointId(v string) *BindVscMountPointAliasRequest
	GetMountPointId() *string
}

type BindVscMountPointAliasRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// sdfe
	AliasPrefix *string `json:"AliasPrefix,omitempty" xml:"AliasPrefix,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 037****e1d
	FileSystemId *string `json:"FileSystemId,omitempty" xml:"FileSystemId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	InputRegionId *string `json:"InputRegionId,omitempty" xml:"InputRegionId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 037cb49e1d-c***5
	MountPointId *string `json:"MountPointId,omitempty" xml:"MountPointId,omitempty"`
}

func (s BindVscMountPointAliasRequest) String() string {
	return dara.Prettify(s)
}

func (s BindVscMountPointAliasRequest) GoString() string {
	return s.String()
}

func (s *BindVscMountPointAliasRequest) GetAliasPrefix() *string {
	return s.AliasPrefix
}

func (s *BindVscMountPointAliasRequest) GetFileSystemId() *string {
	return s.FileSystemId
}

func (s *BindVscMountPointAliasRequest) GetInputRegionId() *string {
	return s.InputRegionId
}

func (s *BindVscMountPointAliasRequest) GetMountPointId() *string {
	return s.MountPointId
}

func (s *BindVscMountPointAliasRequest) SetAliasPrefix(v string) *BindVscMountPointAliasRequest {
	s.AliasPrefix = &v
	return s
}

func (s *BindVscMountPointAliasRequest) SetFileSystemId(v string) *BindVscMountPointAliasRequest {
	s.FileSystemId = &v
	return s
}

func (s *BindVscMountPointAliasRequest) SetInputRegionId(v string) *BindVscMountPointAliasRequest {
	s.InputRegionId = &v
	return s
}

func (s *BindVscMountPointAliasRequest) SetMountPointId(v string) *BindVscMountPointAliasRequest {
	s.MountPointId = &v
	return s
}

func (s *BindVscMountPointAliasRequest) Validate() error {
	return dara.Validate(s)
}
