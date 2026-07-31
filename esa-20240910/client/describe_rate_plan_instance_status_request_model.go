// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRatePlanInstanceStatusRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *DescribeRatePlanInstanceStatusRequest
	GetInstanceId() *string
	SetResourceOwner(v int64) *DescribeRatePlanInstanceStatusRequest
	GetResourceOwner() *int64
}

type DescribeRatePlanInstanceStatusRequest struct {
	// The instance ID. You can call the [ListUserRatePlanInstances](~~ListUserRatePlanInstances~~) operation to obtain the instance ID.
	//
	// example:
	//
	// xcdn-91fknmb80f0g***
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The resource owner account.
	//
	// example:
	//
	// 1700594193617909
	ResourceOwner *int64 `json:"ResourceOwner,omitempty" xml:"ResourceOwner,omitempty"`
}

func (s DescribeRatePlanInstanceStatusRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeRatePlanInstanceStatusRequest) GoString() string {
	return s.String()
}

func (s *DescribeRatePlanInstanceStatusRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeRatePlanInstanceStatusRequest) GetResourceOwner() *int64 {
	return s.ResourceOwner
}

func (s *DescribeRatePlanInstanceStatusRequest) SetInstanceId(v string) *DescribeRatePlanInstanceStatusRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeRatePlanInstanceStatusRequest) SetResourceOwner(v int64) *DescribeRatePlanInstanceStatusRequest {
	s.ResourceOwner = &v
	return s
}

func (s *DescribeRatePlanInstanceStatusRequest) Validate() error {
	return dara.Validate(s)
}
