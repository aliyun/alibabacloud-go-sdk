// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDiskPerformanceLevelRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *ModifyDiskPerformanceLevelRequest
	GetInstanceId() *string
	SetNodeGroupId(v string) *ModifyDiskPerformanceLevelRequest
	GetNodeGroupId() *string
	SetPromotionOptionNo(v string) *ModifyDiskPerformanceLevelRequest
	GetPromotionOptionNo() *string
	SetTarget(v string) *ModifyDiskPerformanceLevelRequest
	GetTarget() *string
}

type ModifyDiskPerformanceLevelRequest struct {
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
	// The disk performance level to which you want to change.
	//
	// Valid values:
	//
	// 	- pl0
	//
	// 	- pl1
	//
	// 	- pl2
	//
	// 	- pl3
	//
	// This parameter is required.
	//
	// example:
	//
	// pl2
	Target *string `json:"Target,omitempty" xml:"Target,omitempty"`
}

func (s ModifyDiskPerformanceLevelRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyDiskPerformanceLevelRequest) GoString() string {
	return s.String()
}

func (s *ModifyDiskPerformanceLevelRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ModifyDiskPerformanceLevelRequest) GetNodeGroupId() *string {
	return s.NodeGroupId
}

func (s *ModifyDiskPerformanceLevelRequest) GetPromotionOptionNo() *string {
	return s.PromotionOptionNo
}

func (s *ModifyDiskPerformanceLevelRequest) GetTarget() *string {
	return s.Target
}

func (s *ModifyDiskPerformanceLevelRequest) SetInstanceId(v string) *ModifyDiskPerformanceLevelRequest {
	s.InstanceId = &v
	return s
}

func (s *ModifyDiskPerformanceLevelRequest) SetNodeGroupId(v string) *ModifyDiskPerformanceLevelRequest {
	s.NodeGroupId = &v
	return s
}

func (s *ModifyDiskPerformanceLevelRequest) SetPromotionOptionNo(v string) *ModifyDiskPerformanceLevelRequest {
	s.PromotionOptionNo = &v
	return s
}

func (s *ModifyDiskPerformanceLevelRequest) SetTarget(v string) *ModifyDiskPerformanceLevelRequest {
	s.Target = &v
	return s
}

func (s *ModifyDiskPerformanceLevelRequest) Validate() error {
	return dara.Validate(s)
}
