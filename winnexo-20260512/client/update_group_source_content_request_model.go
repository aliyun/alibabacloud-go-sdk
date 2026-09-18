// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateGroupSourceContentRequest interface {
	dara.Model
	String() string
	GoString() string
	SetContent(v string) *UpdateGroupSourceContentRequest
	GetContent() *string
	SetForceSync(v bool) *UpdateGroupSourceContentRequest
	GetForceSync() *bool
	SetGroupId(v string) *UpdateGroupSourceContentRequest
	GetGroupId() *string
	SetSourceId(v string) *UpdateGroupSourceContentRequest
	GetSourceId() *string
	SetTenantId(v string) *UpdateGroupSourceContentRequest
	GetTenantId() *string
}

type UpdateGroupSourceContentRequest struct {
	// 更新后的完整正文，可为空字符串；TEXT 存储时去首尾空白；支持 TEXT/本地 txt、md FILE，已有 skip_parse 资料沿用免解析与本地文件扩展名规则
	//
	// This parameter is required.
	//
	// example:
	//
	// 更新后的正文
	Content *string `json:"content,omitempty" xml:"content,omitempty"`
	// 是否等待解析完成；默认 false 异步受理，true 同步等待，网关超时 300000ms
	//
	// example:
	//
	// false
	ForceSync *bool `json:"forceSync,omitempty" xml:"forceSync,omitempty"`
	// 资料所属协作空间 ID
	//
	// This parameter is required.
	//
	// example:
	//
	// group_example
	GroupId *string `json:"groupId,omitempty" xml:"groupId,omitempty"`
	// 当前空间物理 GROUP 资料 ID；引用资料只读
	//
	// This parameter is required.
	//
	// example:
	//
	// source_example
	SourceId *string `json:"sourceId,omitempty" xml:"sourceId,omitempty"`
	// 租户ID，公共参数；缺省时使用调用方默认租户
	//
	// example:
	//
	// 10000
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s UpdateGroupSourceContentRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateGroupSourceContentRequest) GoString() string {
	return s.String()
}

func (s *UpdateGroupSourceContentRequest) GetContent() *string {
	return s.Content
}

func (s *UpdateGroupSourceContentRequest) GetForceSync() *bool {
	return s.ForceSync
}

func (s *UpdateGroupSourceContentRequest) GetGroupId() *string {
	return s.GroupId
}

func (s *UpdateGroupSourceContentRequest) GetSourceId() *string {
	return s.SourceId
}

func (s *UpdateGroupSourceContentRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *UpdateGroupSourceContentRequest) SetContent(v string) *UpdateGroupSourceContentRequest {
	s.Content = &v
	return s
}

func (s *UpdateGroupSourceContentRequest) SetForceSync(v bool) *UpdateGroupSourceContentRequest {
	s.ForceSync = &v
	return s
}

func (s *UpdateGroupSourceContentRequest) SetGroupId(v string) *UpdateGroupSourceContentRequest {
	s.GroupId = &v
	return s
}

func (s *UpdateGroupSourceContentRequest) SetSourceId(v string) *UpdateGroupSourceContentRequest {
	s.SourceId = &v
	return s
}

func (s *UpdateGroupSourceContentRequest) SetTenantId(v string) *UpdateGroupSourceContentRequest {
	s.TenantId = &v
	return s
}

func (s *UpdateGroupSourceContentRequest) Validate() error {
	return dara.Validate(s)
}
