// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteWarmPoolRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *DeleteWarmPoolRequest
	GetClientToken() *string
	SetRegionId(v string) *DeleteWarmPoolRequest
	GetRegionId() *string
	SetScalingGroupId(v string) *DeleteWarmPoolRequest
	GetScalingGroupId() *string
}

type DeleteWarmPoolRequest struct {
	// Ensures the idempotence of the request. You can generate a parameter value on your client to ensure that the value is unique across different requests. The value can contain only ASCII characters and cannot exceed 64 characters in length. For more information, see How to ensure idempotence.
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-42665544****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The region ID of the scaling group.
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
	// asg-bp1igpak5ft1flyp****
	ScalingGroupId *string `json:"ScalingGroupId,omitempty" xml:"ScalingGroupId,omitempty"`
}

func (s DeleteWarmPoolRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteWarmPoolRequest) GoString() string {
	return s.String()
}

func (s *DeleteWarmPoolRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *DeleteWarmPoolRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteWarmPoolRequest) GetScalingGroupId() *string {
	return s.ScalingGroupId
}

func (s *DeleteWarmPoolRequest) SetClientToken(v string) *DeleteWarmPoolRequest {
	s.ClientToken = &v
	return s
}

func (s *DeleteWarmPoolRequest) SetRegionId(v string) *DeleteWarmPoolRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteWarmPoolRequest) SetScalingGroupId(v string) *DeleteWarmPoolRequest {
	s.ScalingGroupId = &v
	return s
}

func (s *DeleteWarmPoolRequest) Validate() error {
	return dara.Validate(s)
}
