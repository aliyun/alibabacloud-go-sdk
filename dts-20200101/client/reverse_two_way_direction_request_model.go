// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReverseTwoWayDirectionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDtsInstanceId(v string) *ReverseTwoWayDirectionRequest
	GetDtsInstanceId() *string
	SetIgnoreErrorSubJob(v bool) *ReverseTwoWayDirectionRequest
	GetIgnoreErrorSubJob() *bool
	SetRegionId(v string) *ReverseTwoWayDirectionRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *ReverseTwoWayDirectionRequest
	GetResourceGroupId() *string
}

type ReverseTwoWayDirectionRequest struct {
	// example:
	//
	// dtsldy114cy24f****
	DtsInstanceId *string `json:"DtsInstanceId,omitempty" xml:"DtsInstanceId,omitempty"`
	// example:
	//
	// false
	IgnoreErrorSubJob *bool `json:"IgnoreErrorSubJob,omitempty" xml:"IgnoreErrorSubJob,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId        *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
}

func (s ReverseTwoWayDirectionRequest) String() string {
	return dara.Prettify(s)
}

func (s ReverseTwoWayDirectionRequest) GoString() string {
	return s.String()
}

func (s *ReverseTwoWayDirectionRequest) GetDtsInstanceId() *string {
	return s.DtsInstanceId
}

func (s *ReverseTwoWayDirectionRequest) GetIgnoreErrorSubJob() *bool {
	return s.IgnoreErrorSubJob
}

func (s *ReverseTwoWayDirectionRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ReverseTwoWayDirectionRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *ReverseTwoWayDirectionRequest) SetDtsInstanceId(v string) *ReverseTwoWayDirectionRequest {
	s.DtsInstanceId = &v
	return s
}

func (s *ReverseTwoWayDirectionRequest) SetIgnoreErrorSubJob(v bool) *ReverseTwoWayDirectionRequest {
	s.IgnoreErrorSubJob = &v
	return s
}

func (s *ReverseTwoWayDirectionRequest) SetRegionId(v string) *ReverseTwoWayDirectionRequest {
	s.RegionId = &v
	return s
}

func (s *ReverseTwoWayDirectionRequest) SetResourceGroupId(v string) *ReverseTwoWayDirectionRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *ReverseTwoWayDirectionRequest) Validate() error {
	return dara.Validate(s)
}
