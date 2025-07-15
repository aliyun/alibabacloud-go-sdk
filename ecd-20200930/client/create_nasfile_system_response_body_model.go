// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateNASFileSystemResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetFileSystemId(v string) *CreateNASFileSystemResponseBody
	GetFileSystemId() *string
	SetFileSystemName(v string) *CreateNASFileSystemResponseBody
	GetFileSystemName() *string
	SetMountTargetDomain(v string) *CreateNASFileSystemResponseBody
	GetMountTargetDomain() *string
	SetOfficeSiteId(v string) *CreateNASFileSystemResponseBody
	GetOfficeSiteId() *string
	SetRequestId(v string) *CreateNASFileSystemResponseBody
	GetRequestId() *string
}

type CreateNASFileSystemResponseBody struct {
	// ID of the NAS file system.
	//
	// example:
	//
	// 04f314****
	FileSystemId *string `json:"FileSystemId,omitempty" xml:"FileSystemId,omitempty"`
	// Name of the NAS file system.
	//
	// example:
	//
	// testNAS
	FileSystemName *string `json:"FileSystemName,omitempty" xml:"FileSystemName,omitempty"`
	// Mount point domain.
	//
	// example:
	//
	// 04f314****-at***.cn-hangzhou.nas.aliyuncs.com
	MountTargetDomain *string `json:"MountTargetDomain,omitempty" xml:"MountTargetDomain,omitempty"`
	// Workspace ID.
	//
	// example:
	//
	// cn-hangzhou+dir-363353****
	OfficeSiteId *string `json:"OfficeSiteId,omitempty" xml:"OfficeSiteId,omitempty"`
	// Request ID.
	//
	// example:
	//
	// 269BDB16-2CD8-4865-84BD-11C40BC21DB0
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateNASFileSystemResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateNASFileSystemResponseBody) GoString() string {
	return s.String()
}

func (s *CreateNASFileSystemResponseBody) GetFileSystemId() *string {
	return s.FileSystemId
}

func (s *CreateNASFileSystemResponseBody) GetFileSystemName() *string {
	return s.FileSystemName
}

func (s *CreateNASFileSystemResponseBody) GetMountTargetDomain() *string {
	return s.MountTargetDomain
}

func (s *CreateNASFileSystemResponseBody) GetOfficeSiteId() *string {
	return s.OfficeSiteId
}

func (s *CreateNASFileSystemResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateNASFileSystemResponseBody) SetFileSystemId(v string) *CreateNASFileSystemResponseBody {
	s.FileSystemId = &v
	return s
}

func (s *CreateNASFileSystemResponseBody) SetFileSystemName(v string) *CreateNASFileSystemResponseBody {
	s.FileSystemName = &v
	return s
}

func (s *CreateNASFileSystemResponseBody) SetMountTargetDomain(v string) *CreateNASFileSystemResponseBody {
	s.MountTargetDomain = &v
	return s
}

func (s *CreateNASFileSystemResponseBody) SetOfficeSiteId(v string) *CreateNASFileSystemResponseBody {
	s.OfficeSiteId = &v
	return s
}

func (s *CreateNASFileSystemResponseBody) SetRequestId(v string) *CreateNASFileSystemResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateNASFileSystemResponseBody) Validate() error {
	return dara.Validate(s)
}
