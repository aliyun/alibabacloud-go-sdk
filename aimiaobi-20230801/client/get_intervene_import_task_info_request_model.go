// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetInterveneImportTaskInfoRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgentKey(v string) *GetInterveneImportTaskInfoRequest
	GetAgentKey() *string
	SetTaskId(v string) *GetInterveneImportTaskInfoRequest
	GetTaskId() *string
}

type GetInterveneImportTaskInfoRequest struct {
	// The unique identifier of the workspace. For more information, see [AgentKey](https://help.aliyun.com/document_detail/2587494.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// 67c520d1fa43455ea44fb69fa402d54d_p_beebot_public
	AgentKey *string `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	// The unique identifier of the task.
	//
	// > This parameter is optional. The system automatically generates a task ID if you do not specify this parameter. Tasks that have the same \\`TaskId\\` belong to the same conversation group.
	//
	// example:
	//
	// 19162157
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s GetInterveneImportTaskInfoRequest) String() string {
	return dara.Prettify(s)
}

func (s GetInterveneImportTaskInfoRequest) GoString() string {
	return s.String()
}

func (s *GetInterveneImportTaskInfoRequest) GetAgentKey() *string {
	return s.AgentKey
}

func (s *GetInterveneImportTaskInfoRequest) GetTaskId() *string {
	return s.TaskId
}

func (s *GetInterveneImportTaskInfoRequest) SetAgentKey(v string) *GetInterveneImportTaskInfoRequest {
	s.AgentKey = &v
	return s
}

func (s *GetInterveneImportTaskInfoRequest) SetTaskId(v string) *GetInterveneImportTaskInfoRequest {
	s.TaskId = &v
	return s
}

func (s *GetInterveneImportTaskInfoRequest) Validate() error {
	return dara.Validate(s)
}
