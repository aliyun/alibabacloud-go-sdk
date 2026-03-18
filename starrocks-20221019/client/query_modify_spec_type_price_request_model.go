// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryModifySpecTypePriceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *QueryModifySpecTypePriceRequest
	GetInstanceId() *string
	SetNodeGroupId(v string) *QueryModifySpecTypePriceRequest
	GetNodeGroupId() *string
	SetPromotionOptionNo(v string) *QueryModifySpecTypePriceRequest
	GetPromotionOptionNo() *string
	SetTargetSpecType(v string) *QueryModifySpecTypePriceRequest
	GetTargetSpecType() *string
}

type QueryModifySpecTypePriceRequest struct {
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
	// standard
	TargetSpecType *string `json:"TargetSpecType,omitempty" xml:"TargetSpecType,omitempty"`
}

func (s QueryModifySpecTypePriceRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryModifySpecTypePriceRequest) GoString() string {
	return s.String()
}

func (s *QueryModifySpecTypePriceRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *QueryModifySpecTypePriceRequest) GetNodeGroupId() *string {
	return s.NodeGroupId
}

func (s *QueryModifySpecTypePriceRequest) GetPromotionOptionNo() *string {
	return s.PromotionOptionNo
}

func (s *QueryModifySpecTypePriceRequest) GetTargetSpecType() *string {
	return s.TargetSpecType
}

func (s *QueryModifySpecTypePriceRequest) SetInstanceId(v string) *QueryModifySpecTypePriceRequest {
	s.InstanceId = &v
	return s
}

func (s *QueryModifySpecTypePriceRequest) SetNodeGroupId(v string) *QueryModifySpecTypePriceRequest {
	s.NodeGroupId = &v
	return s
}

func (s *QueryModifySpecTypePriceRequest) SetPromotionOptionNo(v string) *QueryModifySpecTypePriceRequest {
	s.PromotionOptionNo = &v
	return s
}

func (s *QueryModifySpecTypePriceRequest) SetTargetSpecType(v string) *QueryModifySpecTypePriceRequest {
	s.TargetSpecType = &v
	return s
}

func (s *QueryModifySpecTypePriceRequest) Validate() error {
	return dara.Validate(s)
}
