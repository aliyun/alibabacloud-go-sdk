// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreatePersonalDirectoryRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *CreatePersonalDirectoryRequest
	GetDescription() *string
	SetName(v string) *CreatePersonalDirectoryRequest
	GetName() *string
	SetOperatingObjectName(v string) *CreatePersonalDirectoryRequest
	GetOperatingObjectName() *string
	SetParentDirectoryId(v string) *CreatePersonalDirectoryRequest
	GetParentDirectoryId() *string
	SetTenantId(v string) *CreatePersonalDirectoryRequest
	GetTenantId() *string
}

type CreatePersonalDirectoryRequest struct {
	// 目录描述（可选）
	//
	// example:
	//
	// 示例描述
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// 目录名称
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
	// 父目录 ID；不传时新目录挂在用户的默认根目录下，传入时必须是当前用户的已有个人目录
	//
	// example:
	//
	// exampleParentDirectoryId
	ParentDirectoryId *string `json:"parentDirectoryId,omitempty" xml:"parentDirectoryId,omitempty"`
	// 租户ID，公共参数；winnexo-cli 通过 --tenant-id 显式传入
	//
	// example:
	//
	// 10000
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s CreatePersonalDirectoryRequest) String() string {
	return dara.Prettify(s)
}

func (s CreatePersonalDirectoryRequest) GoString() string {
	return s.String()
}

func (s *CreatePersonalDirectoryRequest) GetDescription() *string {
	return s.Description
}

func (s *CreatePersonalDirectoryRequest) GetName() *string {
	return s.Name
}

func (s *CreatePersonalDirectoryRequest) GetOperatingObjectName() *string {
	return s.OperatingObjectName
}

func (s *CreatePersonalDirectoryRequest) GetParentDirectoryId() *string {
	return s.ParentDirectoryId
}

func (s *CreatePersonalDirectoryRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *CreatePersonalDirectoryRequest) SetDescription(v string) *CreatePersonalDirectoryRequest {
	s.Description = &v
	return s
}

func (s *CreatePersonalDirectoryRequest) SetName(v string) *CreatePersonalDirectoryRequest {
	s.Name = &v
	return s
}

func (s *CreatePersonalDirectoryRequest) SetOperatingObjectName(v string) *CreatePersonalDirectoryRequest {
	s.OperatingObjectName = &v
	return s
}

func (s *CreatePersonalDirectoryRequest) SetParentDirectoryId(v string) *CreatePersonalDirectoryRequest {
	s.ParentDirectoryId = &v
	return s
}

func (s *CreatePersonalDirectoryRequest) SetTenantId(v string) *CreatePersonalDirectoryRequest {
	s.TenantId = &v
	return s
}

func (s *CreatePersonalDirectoryRequest) Validate() error {
	return dara.Validate(s)
}
