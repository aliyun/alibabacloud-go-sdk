// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSharedStoragesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClusterId(v string) *ListSharedStoragesRequest
	GetClusterId() *string
	SetFileSystemId(v string) *ListSharedStoragesRequest
	GetFileSystemId() *string
	SetFileSystemType(v string) *ListSharedStoragesRequest
	GetFileSystemType() *string
}

type ListSharedStoragesRequest struct {
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
	// The ID of the attached file system.
	//
	// example:
	//
	// 0bd504b0**
	FileSystemId *string `json:"FileSystemId,omitempty" xml:"FileSystemId,omitempty"`
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
}

func (s ListSharedStoragesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListSharedStoragesRequest) GoString() string {
	return s.String()
}

func (s *ListSharedStoragesRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *ListSharedStoragesRequest) GetFileSystemId() *string {
	return s.FileSystemId
}

func (s *ListSharedStoragesRequest) GetFileSystemType() *string {
	return s.FileSystemType
}

func (s *ListSharedStoragesRequest) SetClusterId(v string) *ListSharedStoragesRequest {
	s.ClusterId = &v
	return s
}

func (s *ListSharedStoragesRequest) SetFileSystemId(v string) *ListSharedStoragesRequest {
	s.FileSystemId = &v
	return s
}

func (s *ListSharedStoragesRequest) SetFileSystemType(v string) *ListSharedStoragesRequest {
	s.FileSystemType = &v
	return s
}

func (s *ListSharedStoragesRequest) Validate() error {
	return dara.Validate(s)
}
