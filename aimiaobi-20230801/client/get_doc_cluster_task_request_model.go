// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDocClusterTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgentKey(v string) *GetDocClusterTaskRequest
	GetAgentKey() *string
	SetTaskId(v string) *GetDocClusterTaskRequest
	GetTaskId() *string
}

type GetDocClusterTaskRequest struct {
	// The unique identifier of the workspace. For more information, see [AgentKey](https://help.aliyun.com/document_detail/2587494.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// xxxxx_p_efm
	AgentKey *string `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	// The unique ID of the task.
	//
	// > This parameter is optional. The system automatically generates a task ID. If subsequent tasks use the same task ID, they are considered part of the same conversation group.
	//
	// This parameter is required.
	//
	// example:
	//
	// 93771c8e1142467fb1aedf1763feba1e
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s GetDocClusterTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s GetDocClusterTaskRequest) GoString() string {
	return s.String()
}

func (s *GetDocClusterTaskRequest) GetAgentKey() *string {
	return s.AgentKey
}

func (s *GetDocClusterTaskRequest) GetTaskId() *string {
	return s.TaskId
}

func (s *GetDocClusterTaskRequest) SetAgentKey(v string) *GetDocClusterTaskRequest {
	s.AgentKey = &v
	return s
}

func (s *GetDocClusterTaskRequest) SetTaskId(v string) *GetDocClusterTaskRequest {
	s.TaskId = &v
	return s
}

func (s *GetDocClusterTaskRequest) Validate() error {
	return dara.Validate(s)
}
