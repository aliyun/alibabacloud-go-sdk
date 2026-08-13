// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreatePersonalTextRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *CreatePersonalTextRequest
	GetDescription() *string
	SetDirectoryId(v string) *CreatePersonalTextRequest
	GetDirectoryId() *string
	SetName(v string) *CreatePersonalTextRequest
	GetName() *string
	SetOperatingObjectName(v string) *CreatePersonalTextRequest
	GetOperatingObjectName() *string
	SetTenantId(v string) *CreatePersonalTextRequest
	GetTenantId() *string
	SetTextContent(v string) *CreatePersonalTextRequest
	GetTextContent() *string
}

type CreatePersonalTextRequest struct {
	// 资源描述（可选）
	//
	// example:
	//
	// 示例描述
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// 目标个人目录 ID；不传时自动绑定到用户默认根目录，传入时必须是当前用户的已有个人目录
	//
	// example:
	//
	// exampleDirectoryId
	DirectoryId *string `json:"directoryId,omitempty" xml:"directoryId,omitempty"`
	// 资源显示名称
	//
	// This parameter is required.
	//
	// example:
	//
	// 示例名称.pdf
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// 数字员工名称（已废弃：不再作为个人资源隔离条件，仅保留用于来源追溯）
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
	// 纯文本正文（必填）
	//
	// This parameter is required.
	//
	// example:
	//
	// string_value
	TextContent *string `json:"textContent,omitempty" xml:"textContent,omitempty"`
}

func (s CreatePersonalTextRequest) String() string {
	return dara.Prettify(s)
}

func (s CreatePersonalTextRequest) GoString() string {
	return s.String()
}

func (s *CreatePersonalTextRequest) GetDescription() *string {
	return s.Description
}

func (s *CreatePersonalTextRequest) GetDirectoryId() *string {
	return s.DirectoryId
}

func (s *CreatePersonalTextRequest) GetName() *string {
	return s.Name
}

func (s *CreatePersonalTextRequest) GetOperatingObjectName() *string {
	return s.OperatingObjectName
}

func (s *CreatePersonalTextRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *CreatePersonalTextRequest) GetTextContent() *string {
	return s.TextContent
}

func (s *CreatePersonalTextRequest) SetDescription(v string) *CreatePersonalTextRequest {
	s.Description = &v
	return s
}

func (s *CreatePersonalTextRequest) SetDirectoryId(v string) *CreatePersonalTextRequest {
	s.DirectoryId = &v
	return s
}

func (s *CreatePersonalTextRequest) SetName(v string) *CreatePersonalTextRequest {
	s.Name = &v
	return s
}

func (s *CreatePersonalTextRequest) SetOperatingObjectName(v string) *CreatePersonalTextRequest {
	s.OperatingObjectName = &v
	return s
}

func (s *CreatePersonalTextRequest) SetTenantId(v string) *CreatePersonalTextRequest {
	s.TenantId = &v
	return s
}

func (s *CreatePersonalTextRequest) SetTextContent(v string) *CreatePersonalTextRequest {
	s.TextContent = &v
	return s
}

func (s *CreatePersonalTextRequest) Validate() error {
	return dara.Validate(s)
}
