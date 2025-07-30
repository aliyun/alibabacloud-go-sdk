// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyCuRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFastMode(v bool) *ModifyCuRequest
	GetFastMode() *bool
	SetInstanceId(v string) *ModifyCuRequest
	GetInstanceId() *string
	SetNodeGroupId(v string) *ModifyCuRequest
	GetNodeGroupId() *string
	SetPromotionOptionNo(v string) *ModifyCuRequest
	GetPromotionOptionNo() *string
	SetTarget(v int32) *ModifyCuRequest
	GetTarget() *int32
}

type ModifyCuRequest struct {
	// Specifies whether to restart compute nodes in quick restart mode. Default value: false. Valid values:
	//
	// 	- true: Compute nodes are restarted in quick restart mode in multiple batches. The batches are executed in parallel, and the nodes in each batch are restarted at the same time.
	//
	// 	- false: Compute nodes are restarted in rolling restart mode.
	//
	// example:
	//
	// true
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
	// The number of CUs to which you want to change.
	//
	// Valid values:
	//
	// 	- 4
	//
	// 	- 8
	//
	// 	- 16
	//
	// 	- 32
	//
	// 	- 64
	//
	// This parameter is required.
	//
	// example:
	//
	// 4
	Target *int32 `json:"Target,omitempty" xml:"Target,omitempty"`
}

func (s ModifyCuRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyCuRequest) GoString() string {
	return s.String()
}

func (s *ModifyCuRequest) GetFastMode() *bool {
	return s.FastMode
}

func (s *ModifyCuRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ModifyCuRequest) GetNodeGroupId() *string {
	return s.NodeGroupId
}

func (s *ModifyCuRequest) GetPromotionOptionNo() *string {
	return s.PromotionOptionNo
}

func (s *ModifyCuRequest) GetTarget() *int32 {
	return s.Target
}

func (s *ModifyCuRequest) SetFastMode(v bool) *ModifyCuRequest {
	s.FastMode = &v
	return s
}

func (s *ModifyCuRequest) SetInstanceId(v string) *ModifyCuRequest {
	s.InstanceId = &v
	return s
}

func (s *ModifyCuRequest) SetNodeGroupId(v string) *ModifyCuRequest {
	s.NodeGroupId = &v
	return s
}

func (s *ModifyCuRequest) SetPromotionOptionNo(v string) *ModifyCuRequest {
	s.PromotionOptionNo = &v
	return s
}

func (s *ModifyCuRequest) SetTarget(v int32) *ModifyCuRequest {
	s.Target = &v
	return s
}

func (s *ModifyCuRequest) Validate() error {
	return dara.Validate(s)
}
