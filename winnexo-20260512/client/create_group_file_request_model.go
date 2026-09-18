// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateGroupFileRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *CreateGroupFileRequest
	GetDescription() *string
	SetDirectoryId(v string) *CreateGroupFileRequest
	GetDirectoryId() *string
	SetFileRecordId(v string) *CreateGroupFileRequest
	GetFileRecordId() *string
	SetGroupId(v string) *CreateGroupFileRequest
	GetGroupId() *string
	SetName(v string) *CreateGroupFileRequest
	GetName() *string
	SetSourceTags(v string) *CreateGroupFileRequest
	GetSourceTags() *string
	SetTenantId(v string) *CreateGroupFileRequest
	GetTenantId() *string
}

type CreateGroupFileRequest struct {
	// 资料描述
	//
	// example:
	//
	// example
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// 当前空间物理目录ID；省略/root使用空间根，首次可能初始化根目录；引用目录不可写
	//
	// example:
	//
	// dir_example
	DirectoryId *string `json:"directoryId,omitempty" xml:"directoryId,omitempty"`
	// 当前用户在当前租户上传的SOURCE/OSS文件记录ID；须先完成文件PUT
	//
	// This parameter is required.
	//
	// example:
	//
	// example
	FileRecordId *string `json:"fileRecordId,omitempty" xml:"fileRecordId,omitempty"`
	// 协作空间 ID
	//
	// This parameter is required.
	//
	// example:
	//
	// group_example
	GroupId *string `json:"groupId,omitempty" xml:"groupId,omitempty"`
	// 资料显示名；最终名称沿用Provider规则
	//
	// This parameter is required.
	//
	// example:
	//
	// 项目资料
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// 资料标签，JSON字符串列表
	//
	// example:
	//
	// example
	SourceTags *string `json:"sourceTags,omitempty" xml:"sourceTags,omitempty"`
	// 租户ID，公共参数；缺省时使用调用方默认租户
	//
	// example:
	//
	// 10000
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s CreateGroupFileRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateGroupFileRequest) GoString() string {
	return s.String()
}

func (s *CreateGroupFileRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateGroupFileRequest) GetDirectoryId() *string {
	return s.DirectoryId
}

func (s *CreateGroupFileRequest) GetFileRecordId() *string {
	return s.FileRecordId
}

func (s *CreateGroupFileRequest) GetGroupId() *string {
	return s.GroupId
}

func (s *CreateGroupFileRequest) GetName() *string {
	return s.Name
}

func (s *CreateGroupFileRequest) GetSourceTags() *string {
	return s.SourceTags
}

func (s *CreateGroupFileRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *CreateGroupFileRequest) SetDescription(v string) *CreateGroupFileRequest {
	s.Description = &v
	return s
}

func (s *CreateGroupFileRequest) SetDirectoryId(v string) *CreateGroupFileRequest {
	s.DirectoryId = &v
	return s
}

func (s *CreateGroupFileRequest) SetFileRecordId(v string) *CreateGroupFileRequest {
	s.FileRecordId = &v
	return s
}

func (s *CreateGroupFileRequest) SetGroupId(v string) *CreateGroupFileRequest {
	s.GroupId = &v
	return s
}

func (s *CreateGroupFileRequest) SetName(v string) *CreateGroupFileRequest {
	s.Name = &v
	return s
}

func (s *CreateGroupFileRequest) SetSourceTags(v string) *CreateGroupFileRequest {
	s.SourceTags = &v
	return s
}

func (s *CreateGroupFileRequest) SetTenantId(v string) *CreateGroupFileRequest {
	s.TenantId = &v
	return s
}

func (s *CreateGroupFileRequest) Validate() error {
	return dara.Validate(s)
}
