// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateGroupTextRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *CreateGroupTextRequest
	GetDescription() *string
	SetDirectoryId(v string) *CreateGroupTextRequest
	GetDirectoryId() *string
	SetGroupId(v string) *CreateGroupTextRequest
	GetGroupId() *string
	SetName(v string) *CreateGroupTextRequest
	GetName() *string
	SetSourceTags(v string) *CreateGroupTextRequest
	GetSourceTags() *string
	SetTenantId(v string) *CreateGroupTextRequest
	GetTenantId() *string
	SetTextContent(v string) *CreateGroupTextRequest
	GetTextContent() *string
}

type CreateGroupTextRequest struct {
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
	// 纯文本正文，不能全为空白；Provider沿用去首尾空白规则
	//
	// This parameter is required.
	//
	// example:
	//
	// example
	TextContent *string `json:"textContent,omitempty" xml:"textContent,omitempty"`
}

func (s CreateGroupTextRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateGroupTextRequest) GoString() string {
	return s.String()
}

func (s *CreateGroupTextRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateGroupTextRequest) GetDirectoryId() *string {
	return s.DirectoryId
}

func (s *CreateGroupTextRequest) GetGroupId() *string {
	return s.GroupId
}

func (s *CreateGroupTextRequest) GetName() *string {
	return s.Name
}

func (s *CreateGroupTextRequest) GetSourceTags() *string {
	return s.SourceTags
}

func (s *CreateGroupTextRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *CreateGroupTextRequest) GetTextContent() *string {
	return s.TextContent
}

func (s *CreateGroupTextRequest) SetDescription(v string) *CreateGroupTextRequest {
	s.Description = &v
	return s
}

func (s *CreateGroupTextRequest) SetDirectoryId(v string) *CreateGroupTextRequest {
	s.DirectoryId = &v
	return s
}

func (s *CreateGroupTextRequest) SetGroupId(v string) *CreateGroupTextRequest {
	s.GroupId = &v
	return s
}

func (s *CreateGroupTextRequest) SetName(v string) *CreateGroupTextRequest {
	s.Name = &v
	return s
}

func (s *CreateGroupTextRequest) SetSourceTags(v string) *CreateGroupTextRequest {
	s.SourceTags = &v
	return s
}

func (s *CreateGroupTextRequest) SetTenantId(v string) *CreateGroupTextRequest {
	s.TenantId = &v
	return s
}

func (s *CreateGroupTextRequest) SetTextContent(v string) *CreateGroupTextRequest {
	s.TextContent = &v
	return s
}

func (s *CreateGroupTextRequest) Validate() error {
	return dara.Validate(s)
}
