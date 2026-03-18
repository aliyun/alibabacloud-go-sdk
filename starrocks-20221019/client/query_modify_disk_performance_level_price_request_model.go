// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryModifyDiskPerformanceLevelPriceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *QueryModifyDiskPerformanceLevelPriceRequest
	GetInstanceId() *string
	SetNodeGroupId(v string) *QueryModifyDiskPerformanceLevelPriceRequest
	GetNodeGroupId() *string
	SetPromotionOptionNo(v string) *QueryModifyDiskPerformanceLevelPriceRequest
	GetPromotionOptionNo() *string
	SetTarget(v string) *QueryModifyDiskPerformanceLevelPriceRequest
	GetTarget() *string
}

type QueryModifyDiskPerformanceLevelPriceRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// c-a7fa8b85ff6bced6
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// ng-204ea5d711860b4d
	NodeGroupId *string `json:"NodeGroupId,omitempty" xml:"NodeGroupId,omitempty"`
	// example:
	//
	// youhuiquan_12378dfj6
	PromotionOptionNo *string `json:"PromotionOptionNo,omitempty" xml:"PromotionOptionNo,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// pl2
	Target *string `json:"Target,omitempty" xml:"Target,omitempty"`
}

func (s QueryModifyDiskPerformanceLevelPriceRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryModifyDiskPerformanceLevelPriceRequest) GoString() string {
	return s.String()
}

func (s *QueryModifyDiskPerformanceLevelPriceRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *QueryModifyDiskPerformanceLevelPriceRequest) GetNodeGroupId() *string {
	return s.NodeGroupId
}

func (s *QueryModifyDiskPerformanceLevelPriceRequest) GetPromotionOptionNo() *string {
	return s.PromotionOptionNo
}

func (s *QueryModifyDiskPerformanceLevelPriceRequest) GetTarget() *string {
	return s.Target
}

func (s *QueryModifyDiskPerformanceLevelPriceRequest) SetInstanceId(v string) *QueryModifyDiskPerformanceLevelPriceRequest {
	s.InstanceId = &v
	return s
}

func (s *QueryModifyDiskPerformanceLevelPriceRequest) SetNodeGroupId(v string) *QueryModifyDiskPerformanceLevelPriceRequest {
	s.NodeGroupId = &v
	return s
}

func (s *QueryModifyDiskPerformanceLevelPriceRequest) SetPromotionOptionNo(v string) *QueryModifyDiskPerformanceLevelPriceRequest {
	s.PromotionOptionNo = &v
	return s
}

func (s *QueryModifyDiskPerformanceLevelPriceRequest) SetTarget(v string) *QueryModifyDiskPerformanceLevelPriceRequest {
	s.Target = &v
	return s
}

func (s *QueryModifyDiskPerformanceLevelPriceRequest) Validate() error {
	return dara.Validate(s)
}
