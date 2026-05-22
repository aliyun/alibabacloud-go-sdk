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
}

type DescribeRatePlanInstanceStatusRequest struct {
	// The instance ID, which can be obtained by calling the [ListUserRatePlanInstances](~~ListUserRatePlanInstances~~) operation.
	//
	// example:
	//
	// xcdn-91fknmb80f0g***
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
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

func (s *DescribeRatePlanInstanceStatusRequest) SetInstanceId(v string) *DescribeRatePlanInstanceStatusRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeRatePlanInstanceStatusRequest) Validate() error {
	return dara.Validate(s)
}
