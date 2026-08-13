// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreatePersonalDingtalkMeetingRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCredentialId(v string) *CreatePersonalDingtalkMeetingRequest
	GetCredentialId() *string
	SetDescription(v string) *CreatePersonalDingtalkMeetingRequest
	GetDescription() *string
	SetDirectoryId(v string) *CreatePersonalDingtalkMeetingRequest
	GetDirectoryId() *string
	SetName(v string) *CreatePersonalDingtalkMeetingRequest
	GetName() *string
	SetNotes(v string) *CreatePersonalDingtalkMeetingRequest
	GetNotes() *string
	SetOperatingObjectName(v string) *CreatePersonalDingtalkMeetingRequest
	GetOperatingObjectName() *string
	SetRoomCode(v string) *CreatePersonalDingtalkMeetingRequest
	GetRoomCode() *string
	SetTenantId(v string) *CreatePersonalDingtalkMeetingRequest
	GetTenantId() *string
}

type CreatePersonalDingtalkMeetingRequest struct {
	// 凭证 ID（不传则使用系统默认配置）
	//
	// example:
	//
	// exampleCredentialId
	CredentialId *string `json:"credentialId,omitempty" xml:"credentialId,omitempty"`
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
	// 钉钉会议号（必填）
	//
	// This parameter is required.
	//
	// example:
	//
	// string_value
	RoomCode *string `json:"roomCode,omitempty" xml:"roomCode,omitempty"`
	// 租户ID，公共参数；winnexo-cli 通过 --tenant-id 显式传入
	//
	// example:
	//
	// 10000
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s CreatePersonalDingtalkMeetingRequest) String() string {
	return dara.Prettify(s)
}

func (s CreatePersonalDingtalkMeetingRequest) GoString() string {
	return s.String()
}

func (s *CreatePersonalDingtalkMeetingRequest) GetCredentialId() *string {
	return s.CredentialId
}

func (s *CreatePersonalDingtalkMeetingRequest) GetDescription() *string {
	return s.Description
}

func (s *CreatePersonalDingtalkMeetingRequest) GetDirectoryId() *string {
	return s.DirectoryId
}

func (s *CreatePersonalDingtalkMeetingRequest) GetName() *string {
	return s.Name
}

func (s *CreatePersonalDingtalkMeetingRequest) GetNotes() *string {
	return s.Notes
}

func (s *CreatePersonalDingtalkMeetingRequest) GetOperatingObjectName() *string {
	return s.OperatingObjectName
}

func (s *CreatePersonalDingtalkMeetingRequest) GetRoomCode() *string {
	return s.RoomCode
}

func (s *CreatePersonalDingtalkMeetingRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *CreatePersonalDingtalkMeetingRequest) SetCredentialId(v string) *CreatePersonalDingtalkMeetingRequest {
	s.CredentialId = &v
	return s
}

func (s *CreatePersonalDingtalkMeetingRequest) SetDescription(v string) *CreatePersonalDingtalkMeetingRequest {
	s.Description = &v
	return s
}

func (s *CreatePersonalDingtalkMeetingRequest) SetDirectoryId(v string) *CreatePersonalDingtalkMeetingRequest {
	s.DirectoryId = &v
	return s
}

func (s *CreatePersonalDingtalkMeetingRequest) SetName(v string) *CreatePersonalDingtalkMeetingRequest {
	s.Name = &v
	return s
}

func (s *CreatePersonalDingtalkMeetingRequest) SetNotes(v string) *CreatePersonalDingtalkMeetingRequest {
	s.Notes = &v
	return s
}

func (s *CreatePersonalDingtalkMeetingRequest) SetOperatingObjectName(v string) *CreatePersonalDingtalkMeetingRequest {
	s.OperatingObjectName = &v
	return s
}

func (s *CreatePersonalDingtalkMeetingRequest) SetRoomCode(v string) *CreatePersonalDingtalkMeetingRequest {
	s.RoomCode = &v
	return s
}

func (s *CreatePersonalDingtalkMeetingRequest) SetTenantId(v string) *CreatePersonalDingtalkMeetingRequest {
	s.TenantId = &v
	return s
}

func (s *CreatePersonalDingtalkMeetingRequest) Validate() error {
	return dara.Validate(s)
}
