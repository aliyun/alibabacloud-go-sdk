// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateTaskNameRequest interface {
	dara.Model
	String() string
	GoString() string
	SetNodeId(v string) *UpdateTaskNameRequest
	GetNodeId() *string
	SetNodeName(v string) *UpdateTaskNameRequest
	GetNodeName() *string
	SetTid(v int64) *UpdateTaskNameRequest
	GetTid() *int64
}

type UpdateTaskNameRequest struct {
	// The ID of the node. You can call the [GetTaskInstanceRelation](https://help.aliyun.com/document_detail/424711.html) operation to query the node ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 449***
	NodeId *string `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	// The name of the node. You can call the [GetTaskInstanceRelation](https://help.aliyun.com/document_detail/424711.html) operation to query the node name.
	//
	// This parameter is required.
	//
	// example:
	//
	// Spark-test
	NodeName *string `json:"NodeName,omitempty" xml:"NodeName,omitempty"`
	// The ID of the tenant.
	//
	// >  To view the ID of the tenant, go to the Data Management (DMS) console and move the pointer over the profile picture in the upper-right corner. For more information, see [View information about the current tenant](https://help.aliyun.com/document_detail/181330.html).
	//
	// example:
	//
	// 3***
	Tid *int64 `json:"Tid,omitempty" xml:"Tid,omitempty"`
}

func (s UpdateTaskNameRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateTaskNameRequest) GoString() string {
	return s.String()
}

func (s *UpdateTaskNameRequest) GetNodeId() *string {
	return s.NodeId
}

func (s *UpdateTaskNameRequest) GetNodeName() *string {
	return s.NodeName
}

func (s *UpdateTaskNameRequest) GetTid() *int64 {
	return s.Tid
}

func (s *UpdateTaskNameRequest) SetNodeId(v string) *UpdateTaskNameRequest {
	s.NodeId = &v
	return s
}

func (s *UpdateTaskNameRequest) SetNodeName(v string) *UpdateTaskNameRequest {
	s.NodeName = &v
	return s
}

func (s *UpdateTaskNameRequest) SetTid(v int64) *UpdateTaskNameRequest {
	s.Tid = &v
	return s
}

func (s *UpdateTaskNameRequest) Validate() error {
	return dara.Validate(s)
}
