// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteTenantDirectoryRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDeleteMode(v string) *DeleteTenantDirectoryRequest
	GetDeleteMode() *string
	SetDirectoryId(v string) *DeleteTenantDirectoryRequest
	GetDirectoryId() *string
	SetTenantId(v string) *DeleteTenantDirectoryRequest
	GetTenantId() *string
}

type DeleteTenantDirectoryRequest struct {
	// 删除模式：reject / recursive / move_to_root
	//
	// example:
	//
	// reject
	DeleteMode *string `json:"deleteMode,omitempty" xml:"deleteMode,omitempty"`
	// 目录唯一标识
	//
	// example:
	//
	// exampleDirectoryId
	DirectoryId *string `json:"directoryId,omitempty" xml:"directoryId,omitempty"`
	// 租户ID，公共参数；winnexo-cli 通过 --tenant-id 显式传入
	//
	// example:
	//
	// 10000
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s DeleteTenantDirectoryRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteTenantDirectoryRequest) GoString() string {
	return s.String()
}

func (s *DeleteTenantDirectoryRequest) GetDeleteMode() *string {
	return s.DeleteMode
}

func (s *DeleteTenantDirectoryRequest) GetDirectoryId() *string {
	return s.DirectoryId
}

func (s *DeleteTenantDirectoryRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *DeleteTenantDirectoryRequest) SetDeleteMode(v string) *DeleteTenantDirectoryRequest {
	s.DeleteMode = &v
	return s
}

func (s *DeleteTenantDirectoryRequest) SetDirectoryId(v string) *DeleteTenantDirectoryRequest {
	s.DirectoryId = &v
	return s
}

func (s *DeleteTenantDirectoryRequest) SetTenantId(v string) *DeleteTenantDirectoryRequest {
	s.TenantId = &v
	return s
}

func (s *DeleteTenantDirectoryRequest) Validate() error {
	return dara.Validate(s)
}
