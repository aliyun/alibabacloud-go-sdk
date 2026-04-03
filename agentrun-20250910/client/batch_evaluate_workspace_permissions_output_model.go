// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchEvaluateWorkspacePermissionsOutput interface {
	dara.Model
	String() string
	GoString() string
	SetResults(v []*WorkspacePermissionEvaluateResult) *BatchEvaluateWorkspacePermissionsOutput
	GetResults() []*WorkspacePermissionEvaluateResult
}

type BatchEvaluateWorkspacePermissionsOutput struct {
	// 各 workspace 的权限校验结果列表；顺序与请求 workspaceIds 一致
	Results []*WorkspacePermissionEvaluateResult `json:"results" xml:"results" type:"Repeated"`
}

func (s BatchEvaluateWorkspacePermissionsOutput) String() string {
	return dara.Prettify(s)
}

func (s BatchEvaluateWorkspacePermissionsOutput) GoString() string {
	return s.String()
}

func (s *BatchEvaluateWorkspacePermissionsOutput) GetResults() []*WorkspacePermissionEvaluateResult {
	return s.Results
}

func (s *BatchEvaluateWorkspacePermissionsOutput) SetResults(v []*WorkspacePermissionEvaluateResult) *BatchEvaluateWorkspacePermissionsOutput {
	s.Results = v
	return s
}

func (s *BatchEvaluateWorkspacePermissionsOutput) Validate() error {
	if s.Results != nil {
		for _, item := range s.Results {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
