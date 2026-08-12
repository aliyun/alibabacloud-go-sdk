// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryModifyNodeNumberPriceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *QueryModifyNodeNumberPriceRequest
	GetInstanceId() *string
	SetNodeGroupId(v string) *QueryModifyNodeNumberPriceRequest
	GetNodeGroupId() *string
	SetPromotionOptionNo(v string) *QueryModifyNodeNumberPriceRequest
	GetPromotionOptionNo() *string
	SetTarget(v int32) *QueryModifyNodeNumberPriceRequest
	GetTarget() *int32
}

type QueryModifyNodeNumberPriceRequest struct {
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// c-b25e21e24388****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The compute group ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// ng-3d5ce6454354****
	NodeGroupId *string `json:"NodeGroupId,omitempty" xml:"NodeGroupId,omitempty"`
	// The coupon ID.
	//
	// example:
	//
	// youhuiquan_12378dfj6
	PromotionOptionNo *string `json:"PromotionOptionNo,omitempty" xml:"PromotionOptionNo,omitempty"`
	// The target number of nodes.
	//
	// This parameter is required.
	//
	// example:
	//
	// 4
	Target *int32 `json:"Target,omitempty" xml:"Target,omitempty"`
}

func (s QueryModifyNodeNumberPriceRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryModifyNodeNumberPriceRequest) GoString() string {
	return s.String()
}

func (s *QueryModifyNodeNumberPriceRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *QueryModifyNodeNumberPriceRequest) GetNodeGroupId() *string {
	return s.NodeGroupId
}

func (s *QueryModifyNodeNumberPriceRequest) GetPromotionOptionNo() *string {
	return s.PromotionOptionNo
}

func (s *QueryModifyNodeNumberPriceRequest) GetTarget() *int32 {
	return s.Target
}

func (s *QueryModifyNodeNumberPriceRequest) SetInstanceId(v string) *QueryModifyNodeNumberPriceRequest {
	s.InstanceId = &v
	return s
}

func (s *QueryModifyNodeNumberPriceRequest) SetNodeGroupId(v string) *QueryModifyNodeNumberPriceRequest {
	s.NodeGroupId = &v
	return s
}

func (s *QueryModifyNodeNumberPriceRequest) SetPromotionOptionNo(v string) *QueryModifyNodeNumberPriceRequest {
	s.PromotionOptionNo = &v
	return s
}

func (s *QueryModifyNodeNumberPriceRequest) SetTarget(v int32) *QueryModifyNodeNumberPriceRequest {
	s.Target = &v
	return s
}

func (s *QueryModifyNodeNumberPriceRequest) Validate() error {
	return dara.Validate(s)
}
