// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetManagedScalingPolicyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClusterId(v string) *GetManagedScalingPolicyRequest
	GetClusterId() *string
	SetRegionId(v string) *GetManagedScalingPolicyRequest
	GetRegionId() *string
}

type GetManagedScalingPolicyRequest struct {
	// 集群ID。
	//
	// This parameter is required.
	//
	// example:
	//
	// c-b933c5aac8fe****
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// 区域ID。
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s GetManagedScalingPolicyRequest) String() string {
	return dara.Prettify(s)
}

func (s GetManagedScalingPolicyRequest) GoString() string {
	return s.String()
}

func (s *GetManagedScalingPolicyRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *GetManagedScalingPolicyRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *GetManagedScalingPolicyRequest) SetClusterId(v string) *GetManagedScalingPolicyRequest {
	s.ClusterId = &v
	return s
}

func (s *GetManagedScalingPolicyRequest) SetRegionId(v string) *GetManagedScalingPolicyRequest {
	s.RegionId = &v
	return s
}

func (s *GetManagedScalingPolicyRequest) Validate() error {
	return dara.Validate(s)
}
