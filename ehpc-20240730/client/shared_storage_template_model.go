// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSharedStorageTemplate interface {
	dara.Model
	String() string
	GoString() string
	SetFileSystemId(v string) *SharedStorageTemplate
	GetFileSystemId() *string
	SetMountDirectory(v string) *SharedStorageTemplate
	GetMountDirectory() *string
	SetMountOptions(v string) *SharedStorageTemplate
	GetMountOptions() *string
	SetMountTargetDomain(v string) *SharedStorageTemplate
	GetMountTargetDomain() *string
	SetNASDirectory(v string) *SharedStorageTemplate
	GetNASDirectory() *string
	SetProtocolType(v string) *SharedStorageTemplate
	GetProtocolType() *string
}

type SharedStorageTemplate struct {
	// The ID of the mounted file system.
	//
	// example:
	//
	// 008b63****
	FileSystemId *string `json:"FileSystemId,omitempty" xml:"FileSystemId,omitempty"`
	// The local mount directory of the attached file system.
	//
	// example:
	//
	// /home
	MountDirectory *string `json:"MountDirectory,omitempty" xml:"MountDirectory,omitempty"`
	// The storage mount options of the mounted file system.
	//
	// example:
	//
	// -t nfs -o vers=3,nolock,noresvport
	MountOptions *string `json:"MountOptions,omitempty" xml:"MountOptions,omitempty"`
	// The address of the mount point of the file system.
	//
	// example:
	//
	// 008b****-sihc.cn-hangzhou.extreme.nas.aliyuncs.com
	MountTargetDomain *string `json:"MountTargetDomain,omitempty" xml:"MountTargetDomain,omitempty"`
	// The remote mount directory of the mounted file system.
	//
	// example:
	//
	// /
	NASDirectory *string `json:"NASDirectory,omitempty" xml:"NASDirectory,omitempty"`
	// The protocol type of the mounted file system. Valid values:
	//
	// 	- NFS
	//
	// 	- SMB
	//
	// example:
	//
	// NFS
	ProtocolType *string `json:"ProtocolType,omitempty" xml:"ProtocolType,omitempty"`
}

func (s SharedStorageTemplate) String() string {
	return dara.Prettify(s)
}

func (s SharedStorageTemplate) GoString() string {
	return s.String()
}

func (s *SharedStorageTemplate) GetFileSystemId() *string {
	return s.FileSystemId
}

func (s *SharedStorageTemplate) GetMountDirectory() *string {
	return s.MountDirectory
}

func (s *SharedStorageTemplate) GetMountOptions() *string {
	return s.MountOptions
}

func (s *SharedStorageTemplate) GetMountTargetDomain() *string {
	return s.MountTargetDomain
}

func (s *SharedStorageTemplate) GetNASDirectory() *string {
	return s.NASDirectory
}

func (s *SharedStorageTemplate) GetProtocolType() *string {
	return s.ProtocolType
}

func (s *SharedStorageTemplate) SetFileSystemId(v string) *SharedStorageTemplate {
	s.FileSystemId = &v
	return s
}

func (s *SharedStorageTemplate) SetMountDirectory(v string) *SharedStorageTemplate {
	s.MountDirectory = &v
	return s
}

func (s *SharedStorageTemplate) SetMountOptions(v string) *SharedStorageTemplate {
	s.MountOptions = &v
	return s
}

func (s *SharedStorageTemplate) SetMountTargetDomain(v string) *SharedStorageTemplate {
	s.MountTargetDomain = &v
	return s
}

func (s *SharedStorageTemplate) SetNASDirectory(v string) *SharedStorageTemplate {
	s.NASDirectory = &v
	return s
}

func (s *SharedStorageTemplate) SetProtocolType(v string) *SharedStorageTemplate {
	s.ProtocolType = &v
	return s
}

func (s *SharedStorageTemplate) Validate() error {
	return dara.Validate(s)
}
