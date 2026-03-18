// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryModifyDiskTypePriceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *QueryModifyDiskTypePriceRequest
	GetInstanceId() *string
	SetNodeGroupId(v string) *QueryModifyDiskTypePriceRequest
	GetNodeGroupId() *string
	SetPromotionOptionNo(v string) *QueryModifyDiskTypePriceRequest
	GetPromotionOptionNo() *string
	SetTargetDiskType(v string) *QueryModifyDiskTypePriceRequest
	GetTargetDiskType() *string
	SetTargetPerformanceLevel(v string) *QueryModifyDiskTypePriceRequest
	GetTargetPerformanceLevel() *string
}

type QueryModifyDiskTypePriceRequest struct {
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
	// ng-d332aa8bca48****
	NodeGroupId *string `json:"NodeGroupId,omitempty" xml:"NodeGroupId,omitempty"`
	// example:
	//
	// youhuiquan_12378dfj6
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

func (s QueryModifyDiskTypePriceRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryModifyDiskTypePriceRequest) GoString() string {
	return s.String()
}

func (s *QueryModifyDiskTypePriceRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *QueryModifyDiskTypePriceRequest) GetNodeGroupId() *string {
	return s.NodeGroupId
}

func (s *QueryModifyDiskTypePriceRequest) GetPromotionOptionNo() *string {
	return s.PromotionOptionNo
}

func (s *QueryModifyDiskTypePriceRequest) GetTargetDiskType() *string {
	return s.TargetDiskType
}

func (s *QueryModifyDiskTypePriceRequest) GetTargetPerformanceLevel() *string {
	return s.TargetPerformanceLevel
}

func (s *QueryModifyDiskTypePriceRequest) SetInstanceId(v string) *QueryModifyDiskTypePriceRequest {
	s.InstanceId = &v
	return s
}

func (s *QueryModifyDiskTypePriceRequest) SetNodeGroupId(v string) *QueryModifyDiskTypePriceRequest {
	s.NodeGroupId = &v
	return s
}

func (s *QueryModifyDiskTypePriceRequest) SetPromotionOptionNo(v string) *QueryModifyDiskTypePriceRequest {
	s.PromotionOptionNo = &v
	return s
}

func (s *QueryModifyDiskTypePriceRequest) SetTargetDiskType(v string) *QueryModifyDiskTypePriceRequest {
	s.TargetDiskType = &v
	return s
}

func (s *QueryModifyDiskTypePriceRequest) SetTargetPerformanceLevel(v string) *QueryModifyDiskTypePriceRequest {
	s.TargetPerformanceLevel = &v
	return s
}

func (s *QueryModifyDiskTypePriceRequest) Validate() error {
	return dara.Validate(s)
}
