// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iWorkspacePermissionItem interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v *AccessDeniedDetail) *WorkspacePermissionItem
	GetAccessDeniedDetail() *AccessDeniedDetail
	SetAction(v string) *WorkspacePermissionItem
	GetAction() *string
	SetAllowed(v bool) *WorkspacePermissionItem
	GetAllowed() *bool
	SetMessage(v string) *WorkspacePermissionItem
	GetMessage() *string
}

type WorkspacePermissionItem struct {
	// RAM 明确拒绝且可构造 detail 时返回；通过或非 RAM 拒绝场景省略
	//
	// if can be null:
	// true
	AccessDeniedDetail *AccessDeniedDetail `json:"accessDeniedDetail,omitempty" xml:"accessDeniedDetail,omitempty"`
	// 归一化后的 RAM action，始终含 agentrun: 前缀，如 agentrun:ListTemplates
	Action *string `json:"action,omitempty" xml:"action,omitempty"`
	// RAM 判定通过为 true；workspace 不存在、非法/不支持 action、RAM SDK 报错等均为 false
	Allowed *bool `json:"allowed,omitempty" xml:"allowed,omitempty"`
	// 人类可读原因；通过时通常为空字符串
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
}

func (s WorkspacePermissionItem) String() string {
	return dara.Prettify(s)
}

func (s WorkspacePermissionItem) GoString() string {
	return s.String()
}

func (s *WorkspacePermissionItem) GetAccessDeniedDetail() *AccessDeniedDetail {
	return s.AccessDeniedDetail
}

func (s *WorkspacePermissionItem) GetAction() *string {
	return s.Action
}

func (s *WorkspacePermissionItem) GetAllowed() *bool {
	return s.Allowed
}

func (s *WorkspacePermissionItem) GetMessage() *string {
	return s.Message
}

func (s *WorkspacePermissionItem) SetAccessDeniedDetail(v *AccessDeniedDetail) *WorkspacePermissionItem {
	s.AccessDeniedDetail = v
	return s
}

func (s *WorkspacePermissionItem) SetAction(v string) *WorkspacePermissionItem {
	s.Action = &v
	return s
}

func (s *WorkspacePermissionItem) SetAllowed(v bool) *WorkspacePermissionItem {
	s.Allowed = &v
	return s
}

func (s *WorkspacePermissionItem) SetMessage(v string) *WorkspacePermissionItem {
	s.Message = &v
	return s
}

func (s *WorkspacePermissionItem) Validate() error {
	if s.AccessDeniedDetail != nil {
		if err := s.AccessDeniedDetail.Validate(); err != nil {
			return err
		}
	}
	return nil
}
