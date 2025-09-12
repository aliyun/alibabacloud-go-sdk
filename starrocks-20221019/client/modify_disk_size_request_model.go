// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDiskSizeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFastMode(v bool) *ModifyDiskSizeRequest
	GetFastMode() *bool
	SetInstanceId(v string) *ModifyDiskSizeRequest
	GetInstanceId() *string
	SetNodeGroupId(v string) *ModifyDiskSizeRequest
	GetNodeGroupId() *string
	SetPromotionOptionNo(v string) *ModifyDiskSizeRequest
	GetPromotionOptionNo() *string
	SetTarget(v int32) *ModifyDiskSizeRequest
	GetTarget() *int32
}

type ModifyDiskSizeRequest struct {
	FastMode *bool `json:"FastMode,omitempty" xml:"FastMode,omitempty"`
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
	// The disk size to which you want to change to. Unit: GB.
	//
	// This parameter is required.
	//
	// example:
	//
	// 500
	Target *int32 `json:"Target,omitempty" xml:"Target,omitempty"`
}

func (s ModifyDiskSizeRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyDiskSizeRequest) GoString() string {
	return s.String()
}

func (s *ModifyDiskSizeRequest) GetFastMode() *bool {
	return s.FastMode
}

func (s *ModifyDiskSizeRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ModifyDiskSizeRequest) GetNodeGroupId() *string {
	return s.NodeGroupId
}

func (s *ModifyDiskSizeRequest) GetPromotionOptionNo() *string {
	return s.PromotionOptionNo
}

func (s *ModifyDiskSizeRequest) GetTarget() *int32 {
	return s.Target
}

func (s *ModifyDiskSizeRequest) SetFastMode(v bool) *ModifyDiskSizeRequest {
	s.FastMode = &v
	return s
}

func (s *ModifyDiskSizeRequest) SetInstanceId(v string) *ModifyDiskSizeRequest {
	s.InstanceId = &v
	return s
}

func (s *ModifyDiskSizeRequest) SetNodeGroupId(v string) *ModifyDiskSizeRequest {
	s.NodeGroupId = &v
	return s
}

func (s *ModifyDiskSizeRequest) SetPromotionOptionNo(v string) *ModifyDiskSizeRequest {
	s.PromotionOptionNo = &v
	return s
}

func (s *ModifyDiskSizeRequest) SetTarget(v int32) *ModifyDiskSizeRequest {
	s.Target = &v
	return s
}

func (s *ModifyDiskSizeRequest) Validate() error {
	return dara.Validate(s)
}
