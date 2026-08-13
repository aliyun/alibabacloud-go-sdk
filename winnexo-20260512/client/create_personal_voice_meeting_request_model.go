// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreatePersonalVoiceMeetingRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *CreatePersonalVoiceMeetingRequest
	GetDescription() *string
	SetDirectoryId(v string) *CreatePersonalVoiceMeetingRequest
	GetDirectoryId() *string
	SetFileUrl(v string) *CreatePersonalVoiceMeetingRequest
	GetFileUrl() *string
	SetName(v string) *CreatePersonalVoiceMeetingRequest
	GetName() *string
	SetOperatingObjectName(v string) *CreatePersonalVoiceMeetingRequest
	GetOperatingObjectName() *string
	SetTenantId(v string) *CreatePersonalVoiceMeetingRequest
	GetTenantId() *string
}

type CreatePersonalVoiceMeetingRequest struct {
	// 资源描述（可选）
	//
	// example:
	//
	// 示例描述
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// 目标个人目录 ID；不传时自动绑定到当前数字员工默认根目录，传入时必须是当前用户在当前数字员工下的已有个人目录
	//
	// example:
	//
	// exampleDirectoryId
	DirectoryId *string `json:"directoryId,omitempty" xml:"directoryId,omitempty"`
	// 离线会议音频文件 URL（必填）
	//
	// This parameter is required.
	//
	// example:
	//
	// https://example.com/winnexo/resource
	FileUrl *string `json:"fileUrl,omitempty" xml:"fileUrl,omitempty"`
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
	// 租户ID，公共参数，缺省时使用调用方默认租户
	//
	// example:
	//
	// 10000
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s CreatePersonalVoiceMeetingRequest) String() string {
	return dara.Prettify(s)
}

func (s CreatePersonalVoiceMeetingRequest) GoString() string {
	return s.String()
}

func (s *CreatePersonalVoiceMeetingRequest) GetDescription() *string {
	return s.Description
}

func (s *CreatePersonalVoiceMeetingRequest) GetDirectoryId() *string {
	return s.DirectoryId
}

func (s *CreatePersonalVoiceMeetingRequest) GetFileUrl() *string {
	return s.FileUrl
}

func (s *CreatePersonalVoiceMeetingRequest) GetName() *string {
	return s.Name
}

func (s *CreatePersonalVoiceMeetingRequest) GetOperatingObjectName() *string {
	return s.OperatingObjectName
}

func (s *CreatePersonalVoiceMeetingRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *CreatePersonalVoiceMeetingRequest) SetDescription(v string) *CreatePersonalVoiceMeetingRequest {
	s.Description = &v
	return s
}

func (s *CreatePersonalVoiceMeetingRequest) SetDirectoryId(v string) *CreatePersonalVoiceMeetingRequest {
	s.DirectoryId = &v
	return s
}

func (s *CreatePersonalVoiceMeetingRequest) SetFileUrl(v string) *CreatePersonalVoiceMeetingRequest {
	s.FileUrl = &v
	return s
}

func (s *CreatePersonalVoiceMeetingRequest) SetName(v string) *CreatePersonalVoiceMeetingRequest {
	s.Name = &v
	return s
}

func (s *CreatePersonalVoiceMeetingRequest) SetOperatingObjectName(v string) *CreatePersonalVoiceMeetingRequest {
	s.OperatingObjectName = &v
	return s
}

func (s *CreatePersonalVoiceMeetingRequest) SetTenantId(v string) *CreatePersonalVoiceMeetingRequest {
	s.TenantId = &v
	return s
}

func (s *CreatePersonalVoiceMeetingRequest) Validate() error {
	return dara.Validate(s)
}
