// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetNodeId(v int64) *GetTaskRequest
	GetNodeId() *int64
	SetTid(v int64) *GetTaskRequest
	GetTid() *int64
}

type GetTaskRequest struct {
	// The ID of the task node. You can call the [GetTaskInstanceRelation](https://help.aliyun.com/document_detail/424711.html) operation to query the node ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 51***
	NodeId *int64 `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	// The ID of the tenant.
	//
	// > : To view the ID of the tenant, go to the Data Management (DMS) console and move the pointer over the profile picture in the upper-right corner. For more information, see [View information about the current tenant](https://help.aliyun.com/document_detail/181330.html).
	//
	// example:
	//
	// 3***
	Tid *int64 `json:"Tid,omitempty" xml:"Tid,omitempty"`
}

func (s GetTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s GetTaskRequest) GoString() string {
	return s.String()
}

func (s *GetTaskRequest) GetNodeId() *int64 {
	return s.NodeId
}

func (s *GetTaskRequest) GetTid() *int64 {
	return s.Tid
}

func (s *GetTaskRequest) SetNodeId(v int64) *GetTaskRequest {
	s.NodeId = &v
	return s
}

func (s *GetTaskRequest) SetTid(v int64) *GetTaskRequest {
	s.Tid = &v
	return s
}

func (s *GetTaskRequest) Validate() error {
	return dara.Validate(s)
}
