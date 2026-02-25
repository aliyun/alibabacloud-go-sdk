// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDiskTypeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *ModifyDiskTypeRequest
	GetInstanceId() *string
	SetNodeGroupId(v string) *ModifyDiskTypeRequest
	GetNodeGroupId() *string
	SetPromotionOptionNo(v string) *ModifyDiskTypeRequest
	GetPromotionOptionNo() *string
	SetTargetDiskType(v string) *ModifyDiskTypeRequest
	GetTargetDiskType() *string
	SetTargetPerformanceLevel(v string) *ModifyDiskTypeRequest
	GetTargetPerformanceLevel() *string
}

type ModifyDiskTypeRequest struct {
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
	// essd
	TargetDiskType *string `json:"TargetDiskType,omitempty" xml:"TargetDiskType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// pl2
	TargetPerformanceLevel *string `json:"TargetPerformanceLevel,omitempty" xml:"TargetPerformanceLevel,omitempty"`
}

func (s ModifyDiskTypeRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyDiskTypeRequest) GoString() string {
	return s.String()
}

func (s *ModifyDiskTypeRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ModifyDiskTypeRequest) GetNodeGroupId() *string {
	return s.NodeGroupId
}

func (s *ModifyDiskTypeRequest) GetPromotionOptionNo() *string {
	return s.PromotionOptionNo
}

func (s *ModifyDiskTypeRequest) GetTargetDiskType() *string {
	return s.TargetDiskType
}

func (s *ModifyDiskTypeRequest) GetTargetPerformanceLevel() *string {
	return s.TargetPerformanceLevel
}

func (s *ModifyDiskTypeRequest) SetInstanceId(v string) *ModifyDiskTypeRequest {
	s.InstanceId = &v
	return s
}

func (s *ModifyDiskTypeRequest) SetNodeGroupId(v string) *ModifyDiskTypeRequest {
	s.NodeGroupId = &v
	return s
}

func (s *ModifyDiskTypeRequest) SetPromotionOptionNo(v string) *ModifyDiskTypeRequest {
	s.PromotionOptionNo = &v
	return s
}

func (s *ModifyDiskTypeRequest) SetTargetDiskType(v string) *ModifyDiskTypeRequest {
	s.TargetDiskType = &v
	return s
}

func (s *ModifyDiskTypeRequest) SetTargetPerformanceLevel(v string) *ModifyDiskTypeRequest {
	s.TargetPerformanceLevel = &v
	return s
}

func (s *ModifyDiskTypeRequest) Validate() error {
	return dara.Validate(s)
}
