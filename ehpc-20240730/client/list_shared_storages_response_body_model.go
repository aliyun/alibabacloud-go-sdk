// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSharedStoragesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ListSharedStoragesResponseBody
	GetRequestId() *string
	SetSharedStorages(v []*ListSharedStoragesResponseBodySharedStorages) *ListSharedStoragesResponseBody
	GetSharedStorages() []*ListSharedStoragesResponseBodySharedStorages
	SetSuccess(v string) *ListSharedStoragesResponseBody
	GetSuccess() *string
}

type ListSharedStoragesResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// F6757FA4-8FED-4602-B7F5-3550C084****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The information about the attached shared storage.
	SharedStorages []*ListSharedStoragesResponseBodySharedStorages `json:"SharedStorages,omitempty" xml:"SharedStorages,omitempty" type:"Repeated"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- true: The request was successful.
	//
	// 	- false: The request failed.
	//
	// example:
	//
	// true
	Success *string `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListSharedStoragesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListSharedStoragesResponseBody) GoString() string {
	return s.String()
}

func (s *ListSharedStoragesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListSharedStoragesResponseBody) GetSharedStorages() []*ListSharedStoragesResponseBodySharedStorages {
	return s.SharedStorages
}

func (s *ListSharedStoragesResponseBody) GetSuccess() *string {
	return s.Success
}

func (s *ListSharedStoragesResponseBody) SetRequestId(v string) *ListSharedStoragesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListSharedStoragesResponseBody) SetSharedStorages(v []*ListSharedStoragesResponseBodySharedStorages) *ListSharedStoragesResponseBody {
	s.SharedStorages = v
	return s
}

func (s *ListSharedStoragesResponseBody) SetSuccess(v string) *ListSharedStoragesResponseBody {
	s.Success = &v
	return s
}

func (s *ListSharedStoragesResponseBody) Validate() error {
	if s.SharedStorages != nil {
		for _, item := range s.SharedStorages {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListSharedStoragesResponseBodySharedStorages struct {
	// The ID of the attached file system.
	//
	// example:
	//
	// 08c7f4b***
	FileSystemId *string `json:"FileSystemId,omitempty" xml:"FileSystemId,omitempty"`
	// The protocol used by the attached file system. Valid values:
	//
	// 	- nfs3
	//
	// 	- nfs4
	//
	// 	- cpfs
	//
	// example:
	//
	// nfs4
	FileSystemProtocol *string `json:"FileSystemProtocol,omitempty" xml:"FileSystemProtocol,omitempty"`
	// The type of the attached file system. Valid values:
	//
	// 	- nas
	//
	// 	- cpfs
	//
	// example:
	//
	// nas
	FileSystemType *string `json:"FileSystemType,omitempty" xml:"FileSystemType,omitempty"`
	// The mount information.
	MountInfo []*ListSharedStoragesResponseBodySharedStoragesMountInfo `json:"MountInfo,omitempty" xml:"MountInfo,omitempty" type:"Repeated"`
}

func (s ListSharedStoragesResponseBodySharedStorages) String() string {
	return dara.Prettify(s)
}

func (s ListSharedStoragesResponseBodySharedStorages) GoString() string {
	return s.String()
}

func (s *ListSharedStoragesResponseBodySharedStorages) GetFileSystemId() *string {
	return s.FileSystemId
}

func (s *ListSharedStoragesResponseBodySharedStorages) GetFileSystemProtocol() *string {
	return s.FileSystemProtocol
}

func (s *ListSharedStoragesResponseBodySharedStorages) GetFileSystemType() *string {
	return s.FileSystemType
}

func (s *ListSharedStoragesResponseBodySharedStorages) GetMountInfo() []*ListSharedStoragesResponseBodySharedStoragesMountInfo {
	return s.MountInfo
}

func (s *ListSharedStoragesResponseBodySharedStorages) SetFileSystemId(v string) *ListSharedStoragesResponseBodySharedStorages {
	s.FileSystemId = &v
	return s
}

func (s *ListSharedStoragesResponseBodySharedStorages) SetFileSystemProtocol(v string) *ListSharedStoragesResponseBodySharedStorages {
	s.FileSystemProtocol = &v
	return s
}

func (s *ListSharedStoragesResponseBodySharedStorages) SetFileSystemType(v string) *ListSharedStoragesResponseBodySharedStorages {
	s.FileSystemType = &v
	return s
}

func (s *ListSharedStoragesResponseBodySharedStorages) SetMountInfo(v []*ListSharedStoragesResponseBodySharedStoragesMountInfo) *ListSharedStoragesResponseBodySharedStorages {
	s.MountInfo = v
	return s
}

func (s *ListSharedStoragesResponseBodySharedStorages) Validate() error {
	if s.MountInfo != nil {
		for _, item := range s.MountInfo {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListSharedStoragesResponseBodySharedStoragesMountInfo struct {
	// The local mount directory of the attached file system.
	//
	// example:
	//
	// /test
	MountDirectory *string `json:"MountDirectory,omitempty" xml:"MountDirectory,omitempty"`
	// The mount options for the attached file system. Valid values:
	//
	// 	- \\-t nfs -o vers=3,nolock,proto=tcp,noresvport
	//
	// 	- \\-t nfs -o vers=4.0,noresvport
	//
	// example:
	//
	// -t nfs -o vers=4.0,noresvport
	MountOptions *string `json:"MountOptions,omitempty" xml:"MountOptions,omitempty"`
	// The mount target of the attached file system.
	//
	// example:
	//
	// 0bd504b***-ngq26.cn-hangzhou.nas.aliyuncs.com
	MountTarget *string `json:"MountTarget,omitempty" xml:"MountTarget,omitempty"`
	// The protocol used by the mount target of the attached file system. Valid values:
	//
	// 	- nfs3
	//
	// 	- nfs4
	//
	// 	- cpfs
	//
	// example:
	//
	// nfs3
	ProtocolType *string `json:"ProtocolType,omitempty" xml:"ProtocolType,omitempty"`
	// The storage directory of the attached file system.
	//
	// example:
	//
	// /testehpc
	StorageDirectory *string `json:"StorageDirectory,omitempty" xml:"StorageDirectory,omitempty"`
}

func (s ListSharedStoragesResponseBodySharedStoragesMountInfo) String() string {
	return dara.Prettify(s)
}

func (s ListSharedStoragesResponseBodySharedStoragesMountInfo) GoString() string {
	return s.String()
}

func (s *ListSharedStoragesResponseBodySharedStoragesMountInfo) GetMountDirectory() *string {
	return s.MountDirectory
}

func (s *ListSharedStoragesResponseBodySharedStoragesMountInfo) GetMountOptions() *string {
	return s.MountOptions
}

func (s *ListSharedStoragesResponseBodySharedStoragesMountInfo) GetMountTarget() *string {
	return s.MountTarget
}

func (s *ListSharedStoragesResponseBodySharedStoragesMountInfo) GetProtocolType() *string {
	return s.ProtocolType
}

func (s *ListSharedStoragesResponseBodySharedStoragesMountInfo) GetStorageDirectory() *string {
	return s.StorageDirectory
}

func (s *ListSharedStoragesResponseBodySharedStoragesMountInfo) SetMountDirectory(v string) *ListSharedStoragesResponseBodySharedStoragesMountInfo {
	s.MountDirectory = &v
	return s
}

func (s *ListSharedStoragesResponseBodySharedStoragesMountInfo) SetMountOptions(v string) *ListSharedStoragesResponseBodySharedStoragesMountInfo {
	s.MountOptions = &v
	return s
}

func (s *ListSharedStoragesResponseBodySharedStoragesMountInfo) SetMountTarget(v string) *ListSharedStoragesResponseBodySharedStoragesMountInfo {
	s.MountTarget = &v
	return s
}

func (s *ListSharedStoragesResponseBodySharedStoragesMountInfo) SetProtocolType(v string) *ListSharedStoragesResponseBodySharedStoragesMountInfo {
	s.ProtocolType = &v
	return s
}

func (s *ListSharedStoragesResponseBodySharedStoragesMountInfo) SetStorageDirectory(v string) *ListSharedStoragesResponseBodySharedStoragesMountInfo {
	s.StorageDirectory = &v
	return s
}

func (s *ListSharedStoragesResponseBodySharedStoragesMountInfo) Validate() error {
	return dara.Validate(s)
}
