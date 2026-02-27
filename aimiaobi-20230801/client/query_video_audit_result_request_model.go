// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryVideoAuditResultRequest interface {
	dara.Model
	String() string
	GoString() string
	SetTaskId(v string) *QueryVideoAuditResultRequest
	GetTaskId() *string
	SetWorkspaceId(v string) *QueryVideoAuditResultRequest
	GetWorkspaceId() *string
}

type QueryVideoAuditResultRequest struct {
	// 视频审校任务的唯一标识，长度固定32位
	//
	// This parameter is required.
	//
	// example:
	//
	// xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// llm-xxxxx
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s QueryVideoAuditResultRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryVideoAuditResultRequest) GoString() string {
	return s.String()
}

func (s *QueryVideoAuditResultRequest) GetTaskId() *string {
	return s.TaskId
}

func (s *QueryVideoAuditResultRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *QueryVideoAuditResultRequest) SetTaskId(v string) *QueryVideoAuditResultRequest {
	s.TaskId = &v
	return s
}

func (s *QueryVideoAuditResultRequest) SetWorkspaceId(v string) *QueryVideoAuditResultRequest {
	s.WorkspaceId = &v
	return s
}

func (s *QueryVideoAuditResultRequest) Validate() error {
	return dara.Validate(s)
}
