// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetStorageAnalysisResultRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *GetStorageAnalysisResultRequest
	GetInstanceId() *string
	SetNodeId(v string) *GetStorageAnalysisResultRequest
	GetNodeId() *string
	SetTaskId(v string) *GetStorageAnalysisResultRequest
	GetTaskId() *string
}

type GetStorageAnalysisResultRequest struct {
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// rm-bp10xxxxxxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The node ID.
	//
	// >  This parameter is reserved.
	//
	// example:
	//
	// 202****
	NodeId *string `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	// The task ID. You can obtain the task ID from the response of the [CreateStorageAnalysisTask](https://help.aliyun.com/document_detail/2639140.html) operation.
	//
	// This parameter is required.
	//
	// example:
	//
	// 910f83f4b96df0524ddc5749f615****
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s GetStorageAnalysisResultRequest) String() string {
	return dara.Prettify(s)
}

func (s GetStorageAnalysisResultRequest) GoString() string {
	return s.String()
}

func (s *GetStorageAnalysisResultRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetStorageAnalysisResultRequest) GetNodeId() *string {
	return s.NodeId
}

func (s *GetStorageAnalysisResultRequest) GetTaskId() *string {
	return s.TaskId
}

func (s *GetStorageAnalysisResultRequest) SetInstanceId(v string) *GetStorageAnalysisResultRequest {
	s.InstanceId = &v
	return s
}

func (s *GetStorageAnalysisResultRequest) SetNodeId(v string) *GetStorageAnalysisResultRequest {
	s.NodeId = &v
	return s
}

func (s *GetStorageAnalysisResultRequest) SetTaskId(v string) *GetStorageAnalysisResultRequest {
	s.TaskId = &v
	return s
}

func (s *GetStorageAnalysisResultRequest) Validate() error {
	return dara.Validate(s)
}
