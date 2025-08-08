// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRollbackCurrentProjectNodeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *RollbackCurrentProjectNodeRequest
	GetInstanceId() *string
	SetNodeId(v int64) *RollbackCurrentProjectNodeRequest
	GetNodeId() *int64
	SetRemark(v string) *RollbackCurrentProjectNodeRequest
	GetRemark() *string
}

type RollbackCurrentProjectNodeRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 4****89
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1925
	NodeId *int64  `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	Remark *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
}

func (s RollbackCurrentProjectNodeRequest) String() string {
	return dara.Prettify(s)
}

func (s RollbackCurrentProjectNodeRequest) GoString() string {
	return s.String()
}

func (s *RollbackCurrentProjectNodeRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *RollbackCurrentProjectNodeRequest) GetNodeId() *int64 {
	return s.NodeId
}

func (s *RollbackCurrentProjectNodeRequest) GetRemark() *string {
	return s.Remark
}

func (s *RollbackCurrentProjectNodeRequest) SetInstanceId(v string) *RollbackCurrentProjectNodeRequest {
	s.InstanceId = &v
	return s
}

func (s *RollbackCurrentProjectNodeRequest) SetNodeId(v int64) *RollbackCurrentProjectNodeRequest {
	s.NodeId = &v
	return s
}

func (s *RollbackCurrentProjectNodeRequest) SetRemark(v string) *RollbackCurrentProjectNodeRequest {
	s.Remark = &v
	return s
}

func (s *RollbackCurrentProjectNodeRequest) Validate() error {
	return dara.Validate(s)
}
