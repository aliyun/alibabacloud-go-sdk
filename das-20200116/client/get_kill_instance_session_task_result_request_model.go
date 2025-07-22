// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetKillInstanceSessionTaskResultRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *GetKillInstanceSessionTaskResultRequest
	GetInstanceId() *string
	SetNodeId(v string) *GetKillInstanceSessionTaskResultRequest
	GetNodeId() *string
	SetTaskId(v string) *GetKillInstanceSessionTaskResultRequest
	GetTaskId() *string
}

type GetKillInstanceSessionTaskResultRequest struct {
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// rm-2ze1jdv45i7l6****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The node ID.
	//
	// >  You must specify this parameter if your database instance is a PolarDB for MySQL cluster.
	//
	// example:
	//
	// pi-8vbkfj5a756um****
	NodeId *string `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	// The task ID. You can obtain the task ID from the response parameters of the [CreateKillInstanceSessionTask](https://help.aliyun.com/document_detail/609246.html) operation.
	//
	// This parameter is required.
	//
	// example:
	//
	// f77d535b45405bd462b21caa3ee8****
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s GetKillInstanceSessionTaskResultRequest) String() string {
	return dara.Prettify(s)
}

func (s GetKillInstanceSessionTaskResultRequest) GoString() string {
	return s.String()
}

func (s *GetKillInstanceSessionTaskResultRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetKillInstanceSessionTaskResultRequest) GetNodeId() *string {
	return s.NodeId
}

func (s *GetKillInstanceSessionTaskResultRequest) GetTaskId() *string {
	return s.TaskId
}

func (s *GetKillInstanceSessionTaskResultRequest) SetInstanceId(v string) *GetKillInstanceSessionTaskResultRequest {
	s.InstanceId = &v
	return s
}

func (s *GetKillInstanceSessionTaskResultRequest) SetNodeId(v string) *GetKillInstanceSessionTaskResultRequest {
	s.NodeId = &v
	return s
}

func (s *GetKillInstanceSessionTaskResultRequest) SetTaskId(v string) *GetKillInstanceSessionTaskResultRequest {
	s.TaskId = &v
	return s
}

func (s *GetKillInstanceSessionTaskResultRequest) Validate() error {
	return dara.Validate(s)
}
