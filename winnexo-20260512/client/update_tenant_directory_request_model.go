// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateTenantDirectoryRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *UpdateTenantDirectoryRequest
	GetDescription() *string
	SetDirectoryId(v string) *UpdateTenantDirectoryRequest
	GetDirectoryId() *string
	SetName(v string) *UpdateTenantDirectoryRequest
	GetName() *string
	SetParentId(v int64) *UpdateTenantDirectoryRequest
	GetParentId() *int64
	SetPath(v string) *UpdateTenantDirectoryRequest
	GetPath() *string
	SetTenantId(v string) *UpdateTenantDirectoryRequest
	GetTenantId() *string
}

type UpdateTenantDirectoryRequest struct {
	// 新目录描述
	//
	// example:
	//
	// 示例描述
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// 目录唯一标识
	//
	// example:
	//
	// exampleDirectoryId
	DirectoryId *string `json:"directoryId,omitempty" xml:"directoryId,omitempty"`
	// 文件名
	//
	// example:
	//
	// 示例名称.pdf
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// 新父目录内部主键
	//
	// example:
	//
	// 1
	ParentId *int64 `json:"parentId,omitempty" xml:"parentId,omitempty"`
	// 文件 OSS URL
	//
	// example:
	//
	// https://example.com/oss/file.pdf
	Path *string `json:"path,omitempty" xml:"path,omitempty"`
	// 租户ID，公共参数；winnexo-cli 通过 --tenant-id 显式传入
	//
	// example:
	//
	// 10000
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s UpdateTenantDirectoryRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateTenantDirectoryRequest) GoString() string {
	return s.String()
}

func (s *UpdateTenantDirectoryRequest) GetDescription() *string {
	return s.Description
}

func (s *UpdateTenantDirectoryRequest) GetDirectoryId() *string {
	return s.DirectoryId
}

func (s *UpdateTenantDirectoryRequest) GetName() *string {
	return s.Name
}

func (s *UpdateTenantDirectoryRequest) GetParentId() *int64 {
	return s.ParentId
}

func (s *UpdateTenantDirectoryRequest) GetPath() *string {
	return s.Path
}

func (s *UpdateTenantDirectoryRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *UpdateTenantDirectoryRequest) SetDescription(v string) *UpdateTenantDirectoryRequest {
	s.Description = &v
	return s
}

func (s *UpdateTenantDirectoryRequest) SetDirectoryId(v string) *UpdateTenantDirectoryRequest {
	s.DirectoryId = &v
	return s
}

func (s *UpdateTenantDirectoryRequest) SetName(v string) *UpdateTenantDirectoryRequest {
	s.Name = &v
	return s
}

func (s *UpdateTenantDirectoryRequest) SetParentId(v int64) *UpdateTenantDirectoryRequest {
	s.ParentId = &v
	return s
}

func (s *UpdateTenantDirectoryRequest) SetPath(v string) *UpdateTenantDirectoryRequest {
	s.Path = &v
	return s
}

func (s *UpdateTenantDirectoryRequest) SetTenantId(v string) *UpdateTenantDirectoryRequest {
	s.TenantId = &v
	return s
}

func (s *UpdateTenantDirectoryRequest) Validate() error {
	return dara.Validate(s)
}
