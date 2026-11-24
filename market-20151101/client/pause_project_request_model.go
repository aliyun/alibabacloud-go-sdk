// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPauseProjectRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *PauseProjectRequest
	GetInstanceId() *string
	SetNodeId(v int64) *PauseProjectRequest
	GetNodeId() *int64
	SetRemark(v string) *PauseProjectRequest
	GetRemark() *string
}

type PauseProjectRequest struct {
	// The instance ID of the Alibaba Cloud Marketplace order.
	//
	// This parameter is required.
	//
	// example:
	//
	// 4****89
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The ID of the current operation flow node.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1922
	NodeId *int64 `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	// The reason for the remark.
	//
	// example:
	//
	// Waiting for the warehouse to confirm the inventory. The inventory is displayed as abnormal.
	Remark *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
}

func (s PauseProjectRequest) String() string {
	return dara.Prettify(s)
}

func (s PauseProjectRequest) GoString() string {
	return s.String()
}

func (s *PauseProjectRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *PauseProjectRequest) GetNodeId() *int64 {
	return s.NodeId
}

func (s *PauseProjectRequest) GetRemark() *string {
	return s.Remark
}

func (s *PauseProjectRequest) SetInstanceId(v string) *PauseProjectRequest {
	s.InstanceId = &v
	return s
}

func (s *PauseProjectRequest) SetNodeId(v int64) *PauseProjectRequest {
	s.NodeId = &v
	return s
}

func (s *PauseProjectRequest) SetRemark(v string) *PauseProjectRequest {
	s.Remark = &v
	return s
}

func (s *PauseProjectRequest) Validate() error {
	return dara.Validate(s)
}
