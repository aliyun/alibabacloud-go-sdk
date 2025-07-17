// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateTaskContentRequest interface {
	dara.Model
	String() string
	GoString() string
	SetNodeContent(v string) *UpdateTaskContentRequest
	GetNodeContent() *string
	SetNodeId(v string) *UpdateTaskContentRequest
	GetNodeId() *string
	SetTid(v int64) *UpdateTaskContentRequest
	GetTid() *int64
}

type UpdateTaskContentRequest struct {
	// The node configurations after modification.
	//
	// example:
	//
	// { "dbId":12****, "sql":"select 	- from test_table",   "dbType":"lindorm_sql"  }
	NodeContent *string `json:"NodeContent,omitempty" xml:"NodeContent,omitempty"`
	// The ID of the task node. You can call the [GetTaskInstanceRelation](https://help.aliyun.com/document_detail/424711.html) operation to query the node ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 43****
	NodeId *string `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	// The ID of the tenant.
	//
	// > To view the ID of the tenant, move the pointer over the profile picture in the upper-right corner of the Data Management (DMS) console. For more information, see the ["View information about the current tenant"](https://help.aliyun.com/document_detail/181330.html) section of the Manage DMS tenants topic.
	//
	// example:
	//
	// 3***
	Tid *int64 `json:"Tid,omitempty" xml:"Tid,omitempty"`
}

func (s UpdateTaskContentRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateTaskContentRequest) GoString() string {
	return s.String()
}

func (s *UpdateTaskContentRequest) GetNodeContent() *string {
	return s.NodeContent
}

func (s *UpdateTaskContentRequest) GetNodeId() *string {
	return s.NodeId
}

func (s *UpdateTaskContentRequest) GetTid() *int64 {
	return s.Tid
}

func (s *UpdateTaskContentRequest) SetNodeContent(v string) *UpdateTaskContentRequest {
	s.NodeContent = &v
	return s
}

func (s *UpdateTaskContentRequest) SetNodeId(v string) *UpdateTaskContentRequest {
	s.NodeId = &v
	return s
}

func (s *UpdateTaskContentRequest) SetTid(v int64) *UpdateTaskContentRequest {
	s.Tid = &v
	return s
}

func (s *UpdateTaskContentRequest) Validate() error {
	return dara.Validate(s)
}
