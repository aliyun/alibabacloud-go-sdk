// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateTenantDirectoryRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *CreateTenantDirectoryRequest
	GetDescription() *string
	SetName(v string) *CreateTenantDirectoryRequest
	GetName() *string
	SetParentId(v int64) *CreateTenantDirectoryRequest
	GetParentId() *int64
	SetPath(v string) *CreateTenantDirectoryRequest
	GetPath() *string
	SetTenantId(v string) *CreateTenantDirectoryRequest
	GetTenantId() *string
}

type CreateTenantDirectoryRequest struct {
	// 目录描述
	//
	// example:
	//
	// 示例描述
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// 文件名
	//
	// This parameter is required.
	//
	// example:
	//
	// 示例名称.pdf
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// 父目录内部主键；不传表示创建企业知识库根目录
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

func (s CreateTenantDirectoryRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateTenantDirectoryRequest) GoString() string {
	return s.String()
}

func (s *CreateTenantDirectoryRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateTenantDirectoryRequest) GetName() *string {
	return s.Name
}

func (s *CreateTenantDirectoryRequest) GetParentId() *int64 {
	return s.ParentId
}

func (s *CreateTenantDirectoryRequest) GetPath() *string {
	return s.Path
}

func (s *CreateTenantDirectoryRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *CreateTenantDirectoryRequest) SetDescription(v string) *CreateTenantDirectoryRequest {
	s.Description = &v
	return s
}

func (s *CreateTenantDirectoryRequest) SetName(v string) *CreateTenantDirectoryRequest {
	s.Name = &v
	return s
}

func (s *CreateTenantDirectoryRequest) SetParentId(v int64) *CreateTenantDirectoryRequest {
	s.ParentId = &v
	return s
}

func (s *CreateTenantDirectoryRequest) SetPath(v string) *CreateTenantDirectoryRequest {
	s.Path = &v
	return s
}

func (s *CreateTenantDirectoryRequest) SetTenantId(v string) *CreateTenantDirectoryRequest {
	s.TenantId = &v
	return s
}

func (s *CreateTenantDirectoryRequest) Validate() error {
	return dara.Validate(s)
}
