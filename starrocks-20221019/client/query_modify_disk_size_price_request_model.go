// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryModifyDiskSizePriceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *QueryModifyDiskSizePriceRequest
	GetInstanceId() *string
	SetNodeGroupId(v string) *QueryModifyDiskSizePriceRequest
	GetNodeGroupId() *string
	SetPromotionOptionNo(v string) *QueryModifyDiskSizePriceRequest
	GetPromotionOptionNo() *string
	SetTarget(v int32) *QueryModifyDiskSizePriceRequest
	GetTarget() *int32
}

type QueryModifyDiskSizePriceRequest struct {
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
	// 1000
	Target *int32 `json:"Target,omitempty" xml:"Target,omitempty"`
}

func (s QueryModifyDiskSizePriceRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryModifyDiskSizePriceRequest) GoString() string {
	return s.String()
}

func (s *QueryModifyDiskSizePriceRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *QueryModifyDiskSizePriceRequest) GetNodeGroupId() *string {
	return s.NodeGroupId
}

func (s *QueryModifyDiskSizePriceRequest) GetPromotionOptionNo() *string {
	return s.PromotionOptionNo
}

func (s *QueryModifyDiskSizePriceRequest) GetTarget() *int32 {
	return s.Target
}

func (s *QueryModifyDiskSizePriceRequest) SetInstanceId(v string) *QueryModifyDiskSizePriceRequest {
	s.InstanceId = &v
	return s
}

func (s *QueryModifyDiskSizePriceRequest) SetNodeGroupId(v string) *QueryModifyDiskSizePriceRequest {
	s.NodeGroupId = &v
	return s
}

func (s *QueryModifyDiskSizePriceRequest) SetPromotionOptionNo(v string) *QueryModifyDiskSizePriceRequest {
	s.PromotionOptionNo = &v
	return s
}

func (s *QueryModifyDiskSizePriceRequest) SetTarget(v int32) *QueryModifyDiskSizePriceRequest {
	s.Target = &v
	return s
}

func (s *QueryModifyDiskSizePriceRequest) Validate() error {
	return dara.Validate(s)
}
