// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRestartNodeGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFastMode(v bool) *RestartNodeGroupRequest
	GetFastMode() *bool
	SetInstanceId(v string) *RestartNodeGroupRequest
	GetInstanceId() *string
	SetNodeGroupId(v string) *RestartNodeGroupRequest
	GetNodeGroupId() *string
}

type RestartNodeGroupRequest struct {
	// example:
	//
	// true
	FastMode *bool `json:"FastMode,omitempty" xml:"FastMode,omitempty"`
	// example:
	//
	// c-b25e21e24388****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// ng-3d5ce6454354****
	NodeGroupId *string `json:"NodeGroupId,omitempty" xml:"NodeGroupId,omitempty"`
}

func (s RestartNodeGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s RestartNodeGroupRequest) GoString() string {
	return s.String()
}

func (s *RestartNodeGroupRequest) GetFastMode() *bool {
	return s.FastMode
}

func (s *RestartNodeGroupRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *RestartNodeGroupRequest) GetNodeGroupId() *string {
	return s.NodeGroupId
}

func (s *RestartNodeGroupRequest) SetFastMode(v bool) *RestartNodeGroupRequest {
	s.FastMode = &v
	return s
}

func (s *RestartNodeGroupRequest) SetInstanceId(v string) *RestartNodeGroupRequest {
	s.InstanceId = &v
	return s
}

func (s *RestartNodeGroupRequest) SetNodeGroupId(v string) *RestartNodeGroupRequest {
	s.NodeGroupId = &v
	return s
}

func (s *RestartNodeGroupRequest) Validate() error {
	return dara.Validate(s)
}
