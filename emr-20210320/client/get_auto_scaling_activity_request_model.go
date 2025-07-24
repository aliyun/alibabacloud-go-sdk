// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAutoScalingActivityRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClusterId(v string) *GetAutoScalingActivityRequest
	GetClusterId() *string
	SetRegionId(v string) *GetAutoScalingActivityRequest
	GetRegionId() *string
	SetScalingActivityId(v string) *GetAutoScalingActivityRequest
	GetScalingActivityId() *string
}

type GetAutoScalingActivityRequest struct {
	// The cluster ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// c-b933c5aac8fe****
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the scaling activity.
	//
	// This parameter is required.
	//
	// example:
	//
	// asa-36373b084d6b4b13aa50f4129a9e****
	ScalingActivityId *string `json:"ScalingActivityId,omitempty" xml:"ScalingActivityId,omitempty"`
}

func (s GetAutoScalingActivityRequest) String() string {
	return dara.Prettify(s)
}

func (s GetAutoScalingActivityRequest) GoString() string {
	return s.String()
}

func (s *GetAutoScalingActivityRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *GetAutoScalingActivityRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *GetAutoScalingActivityRequest) GetScalingActivityId() *string {
	return s.ScalingActivityId
}

func (s *GetAutoScalingActivityRequest) SetClusterId(v string) *GetAutoScalingActivityRequest {
	s.ClusterId = &v
	return s
}

func (s *GetAutoScalingActivityRequest) SetRegionId(v string) *GetAutoScalingActivityRequest {
	s.RegionId = &v
	return s
}

func (s *GetAutoScalingActivityRequest) SetScalingActivityId(v string) *GetAutoScalingActivityRequest {
	s.ScalingActivityId = &v
	return s
}

func (s *GetAutoScalingActivityRequest) Validate() error {
	return dara.Validate(s)
}
