// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreatePersonalAliDingMeetingRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *CreatePersonalAliDingMeetingRequest
	GetDescription() *string
	SetDirectoryId(v string) *CreatePersonalAliDingMeetingRequest
	GetDirectoryId() *string
	SetName(v string) *CreatePersonalAliDingMeetingRequest
	GetName() *string
	SetNotes(v string) *CreatePersonalAliDingMeetingRequest
	GetNotes() *string
	SetOperatingObjectName(v string) *CreatePersonalAliDingMeetingRequest
	GetOperatingObjectName() *string
	SetShanjiUrl(v string) *CreatePersonalAliDingMeetingRequest
	GetShanjiUrl() *string
	SetTenantId(v string) *CreatePersonalAliDingMeetingRequest
	GetTenantId() *string
}

type CreatePersonalAliDingMeetingRequest struct {
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
	// 资源显示名称
	//
	// This parameter is required.
	//
	// example:
	//
	// 示例名称.pdf
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// 会议笔记内容（可选），会参与辅助分析
	//
	// example:
	//
	// string_value
	Notes *string `json:"notes,omitempty" xml:"notes,omitempty"`
	// 数字员工名称（已废弃：不再作为个人资源隔离条件，仅保留用于来源追溯）
	//
	// example:
	//
	// string_value
	OperatingObjectName *string `json:"operatingObjectName,omitempty" xml:"operatingObjectName,omitempty"`
	// 原始的闪记链接（必填）
	//
	// This parameter is required.
	//
	// example:
	//
	// https://example.com/winnexo/resource
	ShanjiUrl *string `json:"shanjiUrl,omitempty" xml:"shanjiUrl,omitempty"`
	// 租户ID，公共参数；winnexo-cli 通过 --tenant-id 显式传入
	//
	// example:
	//
	// 10000
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s CreatePersonalAliDingMeetingRequest) String() string {
	return dara.Prettify(s)
}

func (s CreatePersonalAliDingMeetingRequest) GoString() string {
	return s.String()
}

func (s *CreatePersonalAliDingMeetingRequest) GetDescription() *string {
	return s.Description
}

func (s *CreatePersonalAliDingMeetingRequest) GetDirectoryId() *string {
	return s.DirectoryId
}

func (s *CreatePersonalAliDingMeetingRequest) GetName() *string {
	return s.Name
}

func (s *CreatePersonalAliDingMeetingRequest) GetNotes() *string {
	return s.Notes
}

func (s *CreatePersonalAliDingMeetingRequest) GetOperatingObjectName() *string {
	return s.OperatingObjectName
}

func (s *CreatePersonalAliDingMeetingRequest) GetShanjiUrl() *string {
	return s.ShanjiUrl
}

func (s *CreatePersonalAliDingMeetingRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *CreatePersonalAliDingMeetingRequest) SetDescription(v string) *CreatePersonalAliDingMeetingRequest {
	s.Description = &v
	return s
}

func (s *CreatePersonalAliDingMeetingRequest) SetDirectoryId(v string) *CreatePersonalAliDingMeetingRequest {
	s.DirectoryId = &v
	return s
}

func (s *CreatePersonalAliDingMeetingRequest) SetName(v string) *CreatePersonalAliDingMeetingRequest {
	s.Name = &v
	return s
}

func (s *CreatePersonalAliDingMeetingRequest) SetNotes(v string) *CreatePersonalAliDingMeetingRequest {
	s.Notes = &v
	return s
}

func (s *CreatePersonalAliDingMeetingRequest) SetOperatingObjectName(v string) *CreatePersonalAliDingMeetingRequest {
	s.OperatingObjectName = &v
	return s
}

func (s *CreatePersonalAliDingMeetingRequest) SetShanjiUrl(v string) *CreatePersonalAliDingMeetingRequest {
	s.ShanjiUrl = &v
	return s
}

func (s *CreatePersonalAliDingMeetingRequest) SetTenantId(v string) *CreatePersonalAliDingMeetingRequest {
	s.TenantId = &v
	return s
}

func (s *CreatePersonalAliDingMeetingRequest) Validate() error {
	return dara.Validate(s)
}
