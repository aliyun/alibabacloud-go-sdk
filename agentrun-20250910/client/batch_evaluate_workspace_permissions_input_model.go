// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchEvaluateWorkspacePermissionsInput interface {
	dara.Model
	String() string
	GoString() string
	SetActions(v []*string) *BatchEvaluateWorkspacePermissionsInput
	GetActions() []*string
	SetWorkspaceIds(v []*string) *BatchEvaluateWorkspacePermissionsInput
	GetWorkspaceIds() []*string
}

type BatchEvaluateWorkspacePermissionsInput struct {
	// RAM List 类 action 列表；支持带 agentrun: 前缀或不带（服务端归一化）；顺序与每条 permissions 一致；上限 20 条
	//
	// This parameter is required.
	Actions []*string `json:"actions" xml:"actions" type:"Repeated"`
	// Workspace 资源 ID 列表（UUID 字符串）；顺序与响应 results 一致；上限 50 条
	//
	// This parameter is required.
	WorkspaceIds []*string `json:"workspaceIds" xml:"workspaceIds" type:"Repeated"`
}

func (s BatchEvaluateWorkspacePermissionsInput) String() string {
	return dara.Prettify(s)
}

func (s BatchEvaluateWorkspacePermissionsInput) GoString() string {
	return s.String()
}

func (s *BatchEvaluateWorkspacePermissionsInput) GetActions() []*string {
	return s.Actions
}

func (s *BatchEvaluateWorkspacePermissionsInput) GetWorkspaceIds() []*string {
	return s.WorkspaceIds
}

func (s *BatchEvaluateWorkspacePermissionsInput) SetActions(v []*string) *BatchEvaluateWorkspacePermissionsInput {
	s.Actions = v
	return s
}

func (s *BatchEvaluateWorkspacePermissionsInput) SetWorkspaceIds(v []*string) *BatchEvaluateWorkspacePermissionsInput {
	s.WorkspaceIds = v
	return s
}

func (s *BatchEvaluateWorkspacePermissionsInput) Validate() error {
	return dara.Validate(s)
}
