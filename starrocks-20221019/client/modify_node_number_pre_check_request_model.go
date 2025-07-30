// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyNodeNumberPreCheckRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *ModifyNodeNumberPreCheckRequest
	GetInstanceId() *string
	SetNodeGroupId(v string) *ModifyNodeNumberPreCheckRequest
	GetNodeGroupId() *string
	SetTarget(v int32) *ModifyNodeNumberPreCheckRequest
	GetTarget() *int32
}

type ModifyNodeNumberPreCheckRequest struct {
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// c-b25e21e24388****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The warehouse ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// ng-3d5ce6454354****
	NodeGroupId *string `json:"NodeGroupId,omitempty" xml:"NodeGroupId,omitempty"`
	// The number of nodes to which you want to change to.
	//
	// This parameter is required.
	//
	// example:
	//
	// 3
	Target *int32 `json:"Target,omitempty" xml:"Target,omitempty"`
}

func (s ModifyNodeNumberPreCheckRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyNodeNumberPreCheckRequest) GoString() string {
	return s.String()
}

func (s *ModifyNodeNumberPreCheckRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ModifyNodeNumberPreCheckRequest) GetNodeGroupId() *string {
	return s.NodeGroupId
}

func (s *ModifyNodeNumberPreCheckRequest) GetTarget() *int32 {
	return s.Target
}

func (s *ModifyNodeNumberPreCheckRequest) SetInstanceId(v string) *ModifyNodeNumberPreCheckRequest {
	s.InstanceId = &v
	return s
}

func (s *ModifyNodeNumberPreCheckRequest) SetNodeGroupId(v string) *ModifyNodeNumberPreCheckRequest {
	s.NodeGroupId = &v
	return s
}

func (s *ModifyNodeNumberPreCheckRequest) SetTarget(v int32) *ModifyNodeNumberPreCheckRequest {
	s.Target = &v
	return s
}

func (s *ModifyNodeNumberPreCheckRequest) Validate() error {
	return dara.Validate(s)
}
