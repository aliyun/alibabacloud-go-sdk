// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryModifyCuPriceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *QueryModifyCuPriceRequest
	GetInstanceId() *string
	SetNodeGroupId(v string) *QueryModifyCuPriceRequest
	GetNodeGroupId() *string
	SetPromotionOptionNo(v string) *QueryModifyCuPriceRequest
	GetPromotionOptionNo() *string
	SetTarget(v int32) *QueryModifyCuPriceRequest
	GetTarget() *int32
}

type QueryModifyCuPriceRequest struct {
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
	// youhuiquan_12378dfj6
	PromotionOptionNo *string `json:"PromotionOptionNo,omitempty" xml:"PromotionOptionNo,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 2
	Target *int32 `json:"Target,omitempty" xml:"Target,omitempty"`
}

func (s QueryModifyCuPriceRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryModifyCuPriceRequest) GoString() string {
	return s.String()
}

func (s *QueryModifyCuPriceRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *QueryModifyCuPriceRequest) GetNodeGroupId() *string {
	return s.NodeGroupId
}

func (s *QueryModifyCuPriceRequest) GetPromotionOptionNo() *string {
	return s.PromotionOptionNo
}

func (s *QueryModifyCuPriceRequest) GetTarget() *int32 {
	return s.Target
}

func (s *QueryModifyCuPriceRequest) SetInstanceId(v string) *QueryModifyCuPriceRequest {
	s.InstanceId = &v
	return s
}

func (s *QueryModifyCuPriceRequest) SetNodeGroupId(v string) *QueryModifyCuPriceRequest {
	s.NodeGroupId = &v
	return s
}

func (s *QueryModifyCuPriceRequest) SetPromotionOptionNo(v string) *QueryModifyCuPriceRequest {
	s.PromotionOptionNo = &v
	return s
}

func (s *QueryModifyCuPriceRequest) SetTarget(v int32) *QueryModifyCuPriceRequest {
	s.Target = &v
	return s
}

func (s *QueryModifyCuPriceRequest) Validate() error {
	return dara.Validate(s)
}
