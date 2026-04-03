// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iWorkspacePermissionEvaluateResult interface {
	dara.Model
	String() string
	GoString() string
	SetPermissions(v []*WorkspacePermissionItem) *WorkspacePermissionEvaluateResult
	GetPermissions() []*WorkspacePermissionItem
	SetWorkspaceFound(v bool) *WorkspacePermissionEvaluateResult
	GetWorkspaceFound() *bool
	SetWorkspaceId(v string) *WorkspacePermissionEvaluateResult
	GetWorkspaceId() *string
}

type WorkspacePermissionEvaluateResult struct {
	// 各 action 的校验结果；顺序与请求 actions 一致
	Permissions []*WorkspacePermissionItem `json:"permissions" xml:"permissions" type:"Repeated"`
	// 当前租户下是否解析到该 workspace；为 false 时各 permissions 一般为 allowed: false，不会调用 RAM
	WorkspaceFound *bool `json:"workspaceFound,omitempty" xml:"workspaceFound,omitempty"`
	// 回显请求中的 workspace ID（trim 后）
	WorkspaceId *string `json:"workspaceId,omitempty" xml:"workspaceId,omitempty"`
}

func (s WorkspacePermissionEvaluateResult) String() string {
	return dara.Prettify(s)
}

func (s WorkspacePermissionEvaluateResult) GoString() string {
	return s.String()
}

func (s *WorkspacePermissionEvaluateResult) GetPermissions() []*WorkspacePermissionItem {
	return s.Permissions
}

func (s *WorkspacePermissionEvaluateResult) GetWorkspaceFound() *bool {
	return s.WorkspaceFound
}

func (s *WorkspacePermissionEvaluateResult) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *WorkspacePermissionEvaluateResult) SetPermissions(v []*WorkspacePermissionItem) *WorkspacePermissionEvaluateResult {
	s.Permissions = v
	return s
}

func (s *WorkspacePermissionEvaluateResult) SetWorkspaceFound(v bool) *WorkspacePermissionEvaluateResult {
	s.WorkspaceFound = &v
	return s
}

func (s *WorkspacePermissionEvaluateResult) SetWorkspaceId(v string) *WorkspacePermissionEvaluateResult {
	s.WorkspaceId = &v
	return s
}

func (s *WorkspacePermissionEvaluateResult) Validate() error {
	if s.Permissions != nil {
		for _, item := range s.Permissions {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
