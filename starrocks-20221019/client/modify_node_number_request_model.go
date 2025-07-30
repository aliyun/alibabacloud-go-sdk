// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyNodeNumberRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *ModifyNodeNumberRequest
	GetInstanceId() *string
	SetNodeGroupId(v string) *ModifyNodeNumberRequest
	GetNodeGroupId() *string
	SetPromotionOptionNo(v string) *ModifyNodeNumberRequest
	GetPromotionOptionNo() *string
	SetTarget(v int32) *ModifyNodeNumberRequest
	GetTarget() *int32
}

type ModifyNodeNumberRequest struct {
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
	NodeGroupId       *string `json:"NodeGroupId,omitempty" xml:"NodeGroupId,omitempty"`
	PromotionOptionNo *string `json:"PromotionOptionNo,omitempty" xml:"PromotionOptionNo,omitempty"`
	// The number of nodes to which you want to change to.
	//
	// This parameter is required.
	//
	// example:
	//
	// 3
	Target *int32 `json:"Target,omitempty" xml:"Target,omitempty"`
}

func (s ModifyNodeNumberRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyNodeNumberRequest) GoString() string {
	return s.String()
}

func (s *ModifyNodeNumberRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ModifyNodeNumberRequest) GetNodeGroupId() *string {
	return s.NodeGroupId
}

func (s *ModifyNodeNumberRequest) GetPromotionOptionNo() *string {
	return s.PromotionOptionNo
}

func (s *ModifyNodeNumberRequest) GetTarget() *int32 {
	return s.Target
}

func (s *ModifyNodeNumberRequest) SetInstanceId(v string) *ModifyNodeNumberRequest {
	s.InstanceId = &v
	return s
}

func (s *ModifyNodeNumberRequest) SetNodeGroupId(v string) *ModifyNodeNumberRequest {
	s.NodeGroupId = &v
	return s
}

func (s *ModifyNodeNumberRequest) SetPromotionOptionNo(v string) *ModifyNodeNumberRequest {
	s.PromotionOptionNo = &v
	return s
}

func (s *ModifyNodeNumberRequest) SetTarget(v int32) *ModifyNodeNumberRequest {
	s.Target = &v
	return s
}

func (s *ModifyNodeNumberRequest) Validate() error {
	return dara.Validate(s)
}
