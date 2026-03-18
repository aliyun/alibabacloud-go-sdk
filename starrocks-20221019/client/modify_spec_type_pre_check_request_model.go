// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifySpecTypePreCheckRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *ModifySpecTypePreCheckRequest
	GetInstanceId() *string
	SetNodeGroupId(v string) *ModifySpecTypePreCheckRequest
	GetNodeGroupId() *string
	SetTargetSpecType(v string) *ModifySpecTypePreCheckRequest
	GetTargetSpecType() *string
}

type ModifySpecTypePreCheckRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// c-b25e21e24388****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// ng-3d5ce6454354****
	NodeGroupId *string `json:"NodeGroupId,omitempty" xml:"NodeGroupId,omitempty"`
	// example:
	//
	// standard
	TargetSpecType *string `json:"TargetSpecType,omitempty" xml:"TargetSpecType,omitempty"`
}

func (s ModifySpecTypePreCheckRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifySpecTypePreCheckRequest) GoString() string {
	return s.String()
}

func (s *ModifySpecTypePreCheckRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ModifySpecTypePreCheckRequest) GetNodeGroupId() *string {
	return s.NodeGroupId
}

func (s *ModifySpecTypePreCheckRequest) GetTargetSpecType() *string {
	return s.TargetSpecType
}

func (s *ModifySpecTypePreCheckRequest) SetInstanceId(v string) *ModifySpecTypePreCheckRequest {
	s.InstanceId = &v
	return s
}

func (s *ModifySpecTypePreCheckRequest) SetNodeGroupId(v string) *ModifySpecTypePreCheckRequest {
	s.NodeGroupId = &v
	return s
}

func (s *ModifySpecTypePreCheckRequest) SetTargetSpecType(v string) *ModifySpecTypePreCheckRequest {
	s.TargetSpecType = &v
	return s
}

func (s *ModifySpecTypePreCheckRequest) Validate() error {
	return dara.Validate(s)
}
