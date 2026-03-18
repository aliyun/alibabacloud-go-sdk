// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryModifyDiskNumberPriceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *QueryModifyDiskNumberPriceRequest
	GetInstanceId() *string
	SetNodeGroupId(v string) *QueryModifyDiskNumberPriceRequest
	GetNodeGroupId() *string
	SetPromotionOptionNo(v string) *QueryModifyDiskNumberPriceRequest
	GetPromotionOptionNo() *string
	SetTarget(v int32) *QueryModifyDiskNumberPriceRequest
	GetTarget() *int32
}

type QueryModifyDiskNumberPriceRequest struct {
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
	// 500
	Target *int32 `json:"Target,omitempty" xml:"Target,omitempty"`
}

func (s QueryModifyDiskNumberPriceRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryModifyDiskNumberPriceRequest) GoString() string {
	return s.String()
}

func (s *QueryModifyDiskNumberPriceRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *QueryModifyDiskNumberPriceRequest) GetNodeGroupId() *string {
	return s.NodeGroupId
}

func (s *QueryModifyDiskNumberPriceRequest) GetPromotionOptionNo() *string {
	return s.PromotionOptionNo
}

func (s *QueryModifyDiskNumberPriceRequest) GetTarget() *int32 {
	return s.Target
}

func (s *QueryModifyDiskNumberPriceRequest) SetInstanceId(v string) *QueryModifyDiskNumberPriceRequest {
	s.InstanceId = &v
	return s
}

func (s *QueryModifyDiskNumberPriceRequest) SetNodeGroupId(v string) *QueryModifyDiskNumberPriceRequest {
	s.NodeGroupId = &v
	return s
}

func (s *QueryModifyDiskNumberPriceRequest) SetPromotionOptionNo(v string) *QueryModifyDiskNumberPriceRequest {
	s.PromotionOptionNo = &v
	return s
}

func (s *QueryModifyDiskNumberPriceRequest) SetTarget(v int32) *QueryModifyDiskNumberPriceRequest {
	s.Target = &v
	return s
}

func (s *QueryModifyDiskNumberPriceRequest) Validate() error {
	return dara.Validate(s)
}
