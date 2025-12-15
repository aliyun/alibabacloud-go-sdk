// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAttachSharedStoragesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClusterId(v string) *AttachSharedStoragesRequest
	GetClusterId() *string
	SetSharedStorages(v []*AttachSharedStoragesRequestSharedStorages) *AttachSharedStoragesRequest
	GetSharedStorages() []*AttachSharedStoragesRequestSharedStorages
}

type AttachSharedStoragesRequest struct {
	// The cluster ID.
	//
	// You can call the [ListClusters](https://help.aliyun.com/document_detail/87116.html) operation to query the cluster ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// ehpc-hz-FYUr32****
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The information about the shared storage resources that you want to attach to the cluster.
	//
	// This parameter is required.
	SharedStorages []*AttachSharedStoragesRequestSharedStorages `json:"SharedStorages,omitempty" xml:"SharedStorages,omitempty" type:"Repeated"`
}

func (s AttachSharedStoragesRequest) String() string {
	return dara.Prettify(s)
}

func (s AttachSharedStoragesRequest) GoString() string {
	return s.String()
}

func (s *AttachSharedStoragesRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *AttachSharedStoragesRequest) GetSharedStorages() []*AttachSharedStoragesRequestSharedStorages {
	return s.SharedStorages
}

func (s *AttachSharedStoragesRequest) SetClusterId(v string) *AttachSharedStoragesRequest {
	s.ClusterId = &v
	return s
}

func (s *AttachSharedStoragesRequest) SetSharedStorages(v []*AttachSharedStoragesRequestSharedStorages) *AttachSharedStoragesRequest {
	s.SharedStorages = v
	return s
}

func (s *AttachSharedStoragesRequest) Validate() error {
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

type AttachSharedStoragesRequestSharedStorages struct {
	// The ID of the file system to be attached.
	//
	// This parameter is required.
	//
	// example:
	//
	// 0bd504b0**
	FileSystemId *string `json:"FileSystemId,omitempty" xml:"FileSystemId,omitempty"`
	// The storage location of the file system to be attached. Valid values:
	//
	// 	- OnPremise: The file system is deployed on a hybrid cloud.
	//
	// 	- PublicCloud: The file system is deployed on a public cloud.
	//
	// example:
	//
	// PublicCloud
	Location *string `json:"Location,omitempty" xml:"Location,omitempty"`
	// The local mount directory of the file system that you want to attach.
	//
	// This parameter is required.
	//
	// example:
	//
	// /test
	MountDirectory *string `json:"MountDirectory,omitempty" xml:"MountDirectory,omitempty"`
	// The attaching options of the file system to be attached. Valid values:
	//
	// 	- \\-t nfs -o vers=3,nolock,proto=tcp,noresvport
	//
	// 	- \\-t nfs -o vers=4.0,noresvport
	//
	// Default value:-t nfs -o vers=3,nolock,proto=tcp,noresvport
	//
	// >  The v3 version is recommended for higher performance if multiple Elastic Compute Service (ECS) instances do not edit the same file at the same time.
	//
	// example:
	//
	// -t nfs -o vers=3,nolock,proto=tcp,noresvport
	MountOptions *string `json:"MountOptions,omitempty" xml:"MountOptions,omitempty"`
	// The address of the mount point of the file system to be attached.
	//
	// This parameter is required.
	//
	// example:
	//
	// 0bd504b***-ngq26.cn-hangzhou.nas.aliyuncs.com
	MountTarget *string `json:"MountTarget,omitempty" xml:"MountTarget,omitempty"`
	// The protocol type of the file system to be attached. Valid values:
	//
	// 	- NFS
	//
	// 	- SMB
	//
	// This parameter is required.
	//
	// example:
	//
	// NFS
	ProtocolType *string `json:"ProtocolType,omitempty" xml:"ProtocolType,omitempty"`
	// The storage directory of the file system. You can mount any directory in the file system to the specified cluster directory.
	//
	// This parameter is required.
	//
	// example:
	//
	// /testehpc
	StorageDirectory *string `json:"StorageDirectory,omitempty" xml:"StorageDirectory,omitempty"`
	// The type of the file system to be attached. Valid values:
	//
	// 	- nas
	//
	// 	- cpfs
	//
	// This parameter is required.
	//
	// example:
	//
	// nas
	VolumeType *string `json:"VolumeType,omitempty" xml:"VolumeType,omitempty"`
}

func (s AttachSharedStoragesRequestSharedStorages) String() string {
	return dara.Prettify(s)
}

func (s AttachSharedStoragesRequestSharedStorages) GoString() string {
	return s.String()
}

func (s *AttachSharedStoragesRequestSharedStorages) GetFileSystemId() *string {
	return s.FileSystemId
}

func (s *AttachSharedStoragesRequestSharedStorages) GetLocation() *string {
	return s.Location
}

func (s *AttachSharedStoragesRequestSharedStorages) GetMountDirectory() *string {
	return s.MountDirectory
}

func (s *AttachSharedStoragesRequestSharedStorages) GetMountOptions() *string {
	return s.MountOptions
}

func (s *AttachSharedStoragesRequestSharedStorages) GetMountTarget() *string {
	return s.MountTarget
}

func (s *AttachSharedStoragesRequestSharedStorages) GetProtocolType() *string {
	return s.ProtocolType
}

func (s *AttachSharedStoragesRequestSharedStorages) GetStorageDirectory() *string {
	return s.StorageDirectory
}

func (s *AttachSharedStoragesRequestSharedStorages) GetVolumeType() *string {
	return s.VolumeType
}

func (s *AttachSharedStoragesRequestSharedStorages) SetFileSystemId(v string) *AttachSharedStoragesRequestSharedStorages {
	s.FileSystemId = &v
	return s
}

func (s *AttachSharedStoragesRequestSharedStorages) SetLocation(v string) *AttachSharedStoragesRequestSharedStorages {
	s.Location = &v
	return s
}

func (s *AttachSharedStoragesRequestSharedStorages) SetMountDirectory(v string) *AttachSharedStoragesRequestSharedStorages {
	s.MountDirectory = &v
	return s
}

func (s *AttachSharedStoragesRequestSharedStorages) SetMountOptions(v string) *AttachSharedStoragesRequestSharedStorages {
	s.MountOptions = &v
	return s
}

func (s *AttachSharedStoragesRequestSharedStorages) SetMountTarget(v string) *AttachSharedStoragesRequestSharedStorages {
	s.MountTarget = &v
	return s
}

func (s *AttachSharedStoragesRequestSharedStorages) SetProtocolType(v string) *AttachSharedStoragesRequestSharedStorages {
	s.ProtocolType = &v
	return s
}

func (s *AttachSharedStoragesRequestSharedStorages) SetStorageDirectory(v string) *AttachSharedStoragesRequestSharedStorages {
	s.StorageDirectory = &v
	return s
}

func (s *AttachSharedStoragesRequestSharedStorages) SetVolumeType(v string) *AttachSharedStoragesRequestSharedStorages {
	s.VolumeType = &v
	return s
}

func (s *AttachSharedStoragesRequestSharedStorages) Validate() error {
	return dara.Validate(s)
}
