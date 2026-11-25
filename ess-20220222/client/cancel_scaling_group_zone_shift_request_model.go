// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCancelScalingGroupZoneShiftRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *CancelScalingGroupZoneShiftRequest
	GetClientToken() *string
	SetRegionId(v string) *CancelScalingGroupZoneShiftRequest
	GetRegionId() *string
	SetScalingGroupId(v string) *CancelScalingGroupZoneShiftRequest
	GetScalingGroupId() *string
}

type CancelScalingGroupZoneShiftRequest struct {
	// The client token that is used to ensure the idempotence of the request. You can generate a unique value for this parameter on your client to ensure that the value is unique across different requests. The value can contain only ASCII characters and cannot exceed 64 characters in length. For more information, see How to ensure idempotence.
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-42665544****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The ID of the region to which the scaling group belongs.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the scaling group.
	//
	// This parameter is required.
	//
	// example:
	//
	// asg-bp14wlu85wrpchm0****
	ScalingGroupId *string `json:"ScalingGroupId,omitempty" xml:"ScalingGroupId,omitempty"`
}

func (s CancelScalingGroupZoneShiftRequest) String() string {
	return dara.Prettify(s)
}

func (s CancelScalingGroupZoneShiftRequest) GoString() string {
	return s.String()
}

func (s *CancelScalingGroupZoneShiftRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CancelScalingGroupZoneShiftRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CancelScalingGroupZoneShiftRequest) GetScalingGroupId() *string {
	return s.ScalingGroupId
}

func (s *CancelScalingGroupZoneShiftRequest) SetClientToken(v string) *CancelScalingGroupZoneShiftRequest {
	s.ClientToken = &v
	return s
}

func (s *CancelScalingGroupZoneShiftRequest) SetRegionId(v string) *CancelScalingGroupZoneShiftRequest {
	s.RegionId = &v
	return s
}

func (s *CancelScalingGroupZoneShiftRequest) SetScalingGroupId(v string) *CancelScalingGroupZoneShiftRequest {
	s.ScalingGroupId = &v
	return s
}

func (s *CancelScalingGroupZoneShiftRequest) Validate() error {
	return dara.Validate(s)
}
