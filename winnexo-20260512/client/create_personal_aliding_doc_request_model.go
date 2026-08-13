// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreatePersonalAlidingDocRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *CreatePersonalAlidingDocRequest
	GetDescription() *string
	SetDirectoryId(v string) *CreatePersonalAlidingDocRequest
	GetDirectoryId() *string
	SetFilePublicUrl(v string) *CreatePersonalAlidingDocRequest
	GetFilePublicUrl() *string
	SetName(v string) *CreatePersonalAlidingDocRequest
	GetName() *string
	SetOperatingObjectName(v string) *CreatePersonalAlidingDocRequest
	GetOperatingObjectName() *string
	SetTenantId(v string) *CreatePersonalAlidingDocRequest
	GetTenantId() *string
}

type CreatePersonalAlidingDocRequest struct {
	// 资源描述（可选）
	//
	// example:
	//
	// 示例描述
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// 目标个人目录 ID；不传时自动绑定到当前数字员工默认根目录，传入时绑定到该目录（必须是当前用户在当前数字员工下的已有个人目录）
	//
	// example:
	//
	// exampleDirectoryId
	DirectoryId *string `json:"directoryId,omitempty" xml:"directoryId,omitempty"`
	// 阿里钉在线文档的可公开访问 URL
	//
	// This parameter is required.
	//
	// example:
	//
	// https://example.com/winnexo/resource
	FilePublicUrl *string `json:"filePublicUrl,omitempty" xml:"filePublicUrl,omitempty"`
	// 文件名
	//
	// This parameter is required.
	//
	// example:
	//
	// 示例名称.pdf
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// Agent 命名空间标识，可选
	//
	// example:
	//
	// string_value
	OperatingObjectName *string `json:"operatingObjectName,omitempty" xml:"operatingObjectName,omitempty"`
	// 租户ID，公共参数；winnexo-cli 通过 --tenant-id 显式传入
	//
	// example:
	//
	// 10000
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s CreatePersonalAlidingDocRequest) String() string {
	return dara.Prettify(s)
}

func (s CreatePersonalAlidingDocRequest) GoString() string {
	return s.String()
}

func (s *CreatePersonalAlidingDocRequest) GetDescription() *string {
	return s.Description
}

func (s *CreatePersonalAlidingDocRequest) GetDirectoryId() *string {
	return s.DirectoryId
}

func (s *CreatePersonalAlidingDocRequest) GetFilePublicUrl() *string {
	return s.FilePublicUrl
}

func (s *CreatePersonalAlidingDocRequest) GetName() *string {
	return s.Name
}

func (s *CreatePersonalAlidingDocRequest) GetOperatingObjectName() *string {
	return s.OperatingObjectName
}

func (s *CreatePersonalAlidingDocRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *CreatePersonalAlidingDocRequest) SetDescription(v string) *CreatePersonalAlidingDocRequest {
	s.Description = &v
	return s
}

func (s *CreatePersonalAlidingDocRequest) SetDirectoryId(v string) *CreatePersonalAlidingDocRequest {
	s.DirectoryId = &v
	return s
}

func (s *CreatePersonalAlidingDocRequest) SetFilePublicUrl(v string) *CreatePersonalAlidingDocRequest {
	s.FilePublicUrl = &v
	return s
}

func (s *CreatePersonalAlidingDocRequest) SetName(v string) *CreatePersonalAlidingDocRequest {
	s.Name = &v
	return s
}

func (s *CreatePersonalAlidingDocRequest) SetOperatingObjectName(v string) *CreatePersonalAlidingDocRequest {
	s.OperatingObjectName = &v
	return s
}

func (s *CreatePersonalAlidingDocRequest) SetTenantId(v string) *CreatePersonalAlidingDocRequest {
	s.TenantId = &v
	return s
}

func (s *CreatePersonalAlidingDocRequest) Validate() error {
	return dara.Validate(s)
}
