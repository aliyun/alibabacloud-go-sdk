// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateTaskConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetNodeConfig(v string) *UpdateTaskConfigRequest
	GetNodeConfig() *string
	SetNodeId(v string) *UpdateTaskConfigRequest
	GetNodeId() *string
	SetTid(v int64) *UpdateTaskConfigRequest
	GetTid() *int64
}

type UpdateTaskConfigRequest struct {
	// The advanced configuration for the node. The value of this parameter must be a JSON string.
	//
	// This parameter is required.
	//
	// example:
	//
	// { "rerun":{ "rerunEnable":true,  "rerunCount":1,   "rerunInterval":10 } }
	NodeConfig *string `json:"NodeConfig,omitempty" xml:"NodeConfig,omitempty"`
	// The ID of the task node. You can call the [GetTaskInstanceRelation](https://help.aliyun.com/document_detail/424711.html) operation to query the node ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 43****
	NodeId *string `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	// The ID of the tenant. You can call the [GetUserActiveTenant](https://help.aliyun.com/document_detail/198073.html) or [ListUserTenants](https://help.aliyun.com/document_detail/198074.html) operation to query the tenant ID.
	//
	// example:
	//
	// 3***
	Tid *int64 `json:"Tid,omitempty" xml:"Tid,omitempty"`
}

func (s UpdateTaskConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateTaskConfigRequest) GoString() string {
	return s.String()
}

func (s *UpdateTaskConfigRequest) GetNodeConfig() *string {
	return s.NodeConfig
}

func (s *UpdateTaskConfigRequest) GetNodeId() *string {
	return s.NodeId
}

func (s *UpdateTaskConfigRequest) GetTid() *int64 {
	return s.Tid
}

func (s *UpdateTaskConfigRequest) SetNodeConfig(v string) *UpdateTaskConfigRequest {
	s.NodeConfig = &v
	return s
}

func (s *UpdateTaskConfigRequest) SetNodeId(v string) *UpdateTaskConfigRequest {
	s.NodeId = &v
	return s
}

func (s *UpdateTaskConfigRequest) SetTid(v int64) *UpdateTaskConfigRequest {
	s.Tid = &v
	return s
}

func (s *UpdateTaskConfigRequest) Validate() error {
	return dara.Validate(s)
}
