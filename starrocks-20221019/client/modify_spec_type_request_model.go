// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifySpecTypeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFastMode(v bool) *ModifySpecTypeRequest
	GetFastMode() *bool
	SetInstanceId(v string) *ModifySpecTypeRequest
	GetInstanceId() *string
	SetNodeGroupId(v string) *ModifySpecTypeRequest
	GetNodeGroupId() *string
	SetPromotionOptionNo(v string) *ModifySpecTypeRequest
	GetPromotionOptionNo() *string
	SetTargetSpecType(v string) *ModifySpecTypeRequest
	GetTargetSpecType() *string
}

type ModifySpecTypeRequest struct {
	// example:
	//
	// true
	FastMode *bool `json:"FastMode,omitempty" xml:"FastMode,omitempty"`
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
	// youhuiquan_promotion_option_id_for_blank
	PromotionOptionNo *string `json:"PromotionOptionNo,omitempty" xml:"PromotionOptionNo,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// standard
	TargetSpecType *string `json:"TargetSpecType,omitempty" xml:"TargetSpecType,omitempty"`
}

func (s ModifySpecTypeRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifySpecTypeRequest) GoString() string {
	return s.String()
}

func (s *ModifySpecTypeRequest) GetFastMode() *bool {
	return s.FastMode
}

func (s *ModifySpecTypeRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ModifySpecTypeRequest) GetNodeGroupId() *string {
	return s.NodeGroupId
}

func (s *ModifySpecTypeRequest) GetPromotionOptionNo() *string {
	return s.PromotionOptionNo
}

func (s *ModifySpecTypeRequest) GetTargetSpecType() *string {
	return s.TargetSpecType
}

func (s *ModifySpecTypeRequest) SetFastMode(v bool) *ModifySpecTypeRequest {
	s.FastMode = &v
	return s
}

func (s *ModifySpecTypeRequest) SetInstanceId(v string) *ModifySpecTypeRequest {
	s.InstanceId = &v
	return s
}

func (s *ModifySpecTypeRequest) SetNodeGroupId(v string) *ModifySpecTypeRequest {
	s.NodeGroupId = &v
	return s
}

func (s *ModifySpecTypeRequest) SetPromotionOptionNo(v string) *ModifySpecTypeRequest {
	s.PromotionOptionNo = &v
	return s
}

func (s *ModifySpecTypeRequest) SetTargetSpecType(v string) *ModifySpecTypeRequest {
	s.TargetSpecType = &v
	return s
}

func (s *ModifySpecTypeRequest) Validate() error {
	return dara.Validate(s)
}
