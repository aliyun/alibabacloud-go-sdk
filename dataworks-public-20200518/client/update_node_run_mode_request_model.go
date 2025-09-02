// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateNodeRunModeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetNodeId(v int64) *UpdateNodeRunModeRequest
	GetNodeId() *int64
	SetProjectEnv(v string) *UpdateNodeRunModeRequest
	GetProjectEnv() *string
	SetSchedulerType(v int32) *UpdateNodeRunModeRequest
	GetSchedulerType() *int32
}

type UpdateNodeRunModeRequest struct {
	// The node ID. You can call the [ListNodes](https://help.aliyun.com/document_detail/173979.html) operation to query the node ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1234
	NodeId *int64 `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	// The environment in which the node runs. Valid values: DEV and PROD. The value DEV indicates the development environment, and the value PROD indicates the production environment.
	//
	// 	- PROD
	//
	// 	- DEV
	//
	// This parameter is required.
	//
	// example:
	//
	// PROD
	ProjectEnv *string `json:"ProjectEnv,omitempty" xml:"ProjectEnv,omitempty"`
	// The operation that you want to perform on the node. Valid values:
	//
	// 	- 0: indicates that you want to unfreeze the node.
	//
	// 	- 2: indicates that you want to freeze the node.
	//
	// This parameter is required.
	//
	// example:
	//
	// 0
	SchedulerType *int32 `json:"SchedulerType,omitempty" xml:"SchedulerType,omitempty"`
}

func (s UpdateNodeRunModeRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateNodeRunModeRequest) GoString() string {
	return s.String()
}

func (s *UpdateNodeRunModeRequest) GetNodeId() *int64 {
	return s.NodeId
}

func (s *UpdateNodeRunModeRequest) GetProjectEnv() *string {
	return s.ProjectEnv
}

func (s *UpdateNodeRunModeRequest) GetSchedulerType() *int32 {
	return s.SchedulerType
}

func (s *UpdateNodeRunModeRequest) SetNodeId(v int64) *UpdateNodeRunModeRequest {
	s.NodeId = &v
	return s
}

func (s *UpdateNodeRunModeRequest) SetProjectEnv(v string) *UpdateNodeRunModeRequest {
	s.ProjectEnv = &v
	return s
}

func (s *UpdateNodeRunModeRequest) SetSchedulerType(v int32) *UpdateNodeRunModeRequest {
	s.SchedulerType = &v
	return s
}

func (s *UpdateNodeRunModeRequest) Validate() error {
	return dara.Validate(s)
}
