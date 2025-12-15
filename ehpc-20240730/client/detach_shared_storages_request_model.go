// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDetachSharedStoragesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClusterId(v string) *DetachSharedStoragesRequest
	GetClusterId() *string
	SetSharedStorages(v []*DetachSharedStoragesRequestSharedStorages) *DetachSharedStoragesRequest
	GetSharedStorages() []*DetachSharedStoragesRequestSharedStorages
}

type DetachSharedStoragesRequest struct {
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
	// The information about mounted shared storage resources.
	//
	// This parameter is required.
	SharedStorages []*DetachSharedStoragesRequestSharedStorages `json:"SharedStorages,omitempty" xml:"SharedStorages,omitempty" type:"Repeated"`
}

func (s DetachSharedStoragesRequest) String() string {
	return dara.Prettify(s)
}

func (s DetachSharedStoragesRequest) GoString() string {
	return s.String()
}

func (s *DetachSharedStoragesRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *DetachSharedStoragesRequest) GetSharedStorages() []*DetachSharedStoragesRequestSharedStorages {
	return s.SharedStorages
}

func (s *DetachSharedStoragesRequest) SetClusterId(v string) *DetachSharedStoragesRequest {
	s.ClusterId = &v
	return s
}

func (s *DetachSharedStoragesRequest) SetSharedStorages(v []*DetachSharedStoragesRequestSharedStorages) *DetachSharedStoragesRequest {
	s.SharedStorages = v
	return s
}

func (s *DetachSharedStoragesRequest) Validate() error {
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

type DetachSharedStoragesRequestSharedStorages struct {
	// The local mount directory of the mounted file system.
	//
	// This parameter is required.
	//
	// example:
	//
	// /test
	MountDirectory *string `json:"MountDirectory,omitempty" xml:"MountDirectory,omitempty"`
}

func (s DetachSharedStoragesRequestSharedStorages) String() string {
	return dara.Prettify(s)
}

func (s DetachSharedStoragesRequestSharedStorages) GoString() string {
	return s.String()
}

func (s *DetachSharedStoragesRequestSharedStorages) GetMountDirectory() *string {
	return s.MountDirectory
}

func (s *DetachSharedStoragesRequestSharedStorages) SetMountDirectory(v string) *DetachSharedStoragesRequestSharedStorages {
	s.MountDirectory = &v
	return s
}

func (s *DetachSharedStoragesRequestSharedStorages) Validate() error {
	return dara.Validate(s)
}
