// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iFetchExportWordTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgentKey(v string) *FetchExportWordTaskRequest
	GetAgentKey() *string
	SetTaskId(v string) *FetchExportWordTaskRequest
	GetTaskId() *string
}

type FetchExportWordTaskRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// xxxxx_p_efm
	AgentKey *string `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	// example:
	//
	// 2e27abb32cb64f80a0c6e829b6c87a09
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s FetchExportWordTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s FetchExportWordTaskRequest) GoString() string {
	return s.String()
}

func (s *FetchExportWordTaskRequest) GetAgentKey() *string {
	return s.AgentKey
}

func (s *FetchExportWordTaskRequest) GetTaskId() *string {
	return s.TaskId
}

func (s *FetchExportWordTaskRequest) SetAgentKey(v string) *FetchExportWordTaskRequest {
	s.AgentKey = &v
	return s
}

func (s *FetchExportWordTaskRequest) SetTaskId(v string) *FetchExportWordTaskRequest {
	s.TaskId = &v
	return s
}

func (s *FetchExportWordTaskRequest) Validate() error {
	return dara.Validate(s)
}
