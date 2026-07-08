// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetEnterpriseVocAnalysisTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetTaskId(v string) *GetEnterpriseVocAnalysisTaskRequest
	GetTaskId() *string
	SetWorkspaceId(v string) *GetEnterpriseVocAnalysisTaskRequest
	GetWorkspaceId() *string
}

type GetEnterpriseVocAnalysisTaskRequest struct {
	// The unique ID of the task.
	//
	// > You do not need to specify this parameter. The system automatically generates a TaskId. If you specify the same TaskId for subsequent tasks, the tasks are considered part of the same conversation group.
	//
	// This parameter is required.
	//
	// example:
	//
	// xxx
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// The unique ID of the Alibaba Cloud Model Studio workspace. For more information, see [Get a workspace ID](https://help.aliyun.com/document_detail/2782167.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// xxxx
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s GetEnterpriseVocAnalysisTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s GetEnterpriseVocAnalysisTaskRequest) GoString() string {
	return s.String()
}

func (s *GetEnterpriseVocAnalysisTaskRequest) GetTaskId() *string {
	return s.TaskId
}

func (s *GetEnterpriseVocAnalysisTaskRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *GetEnterpriseVocAnalysisTaskRequest) SetTaskId(v string) *GetEnterpriseVocAnalysisTaskRequest {
	s.TaskId = &v
	return s
}

func (s *GetEnterpriseVocAnalysisTaskRequest) SetWorkspaceId(v string) *GetEnterpriseVocAnalysisTaskRequest {
	s.WorkspaceId = &v
	return s
}

func (s *GetEnterpriseVocAnalysisTaskRequest) Validate() error {
	return dara.Validate(s)
}
