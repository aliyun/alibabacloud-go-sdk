// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iApplyScalingGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ApplyScalingGroupResponseBody
	GetRequestId() *string
	SetScalingGroupId(v string) *ApplyScalingGroupResponseBody
	GetScalingGroupId() *string
}

type ApplyScalingGroupResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// CC107349-57B7-4405-B1BF-9BF5AF7F****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The ID of the enabled scaling group.
	//
	// example:
	//
	// asg-bp1igpak5ft1flyp****
	ScalingGroupId *string `json:"ScalingGroupId,omitempty" xml:"ScalingGroupId,omitempty"`
}

func (s ApplyScalingGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ApplyScalingGroupResponseBody) GoString() string {
	return s.String()
}

func (s *ApplyScalingGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ApplyScalingGroupResponseBody) GetScalingGroupId() *string {
	return s.ScalingGroupId
}

func (s *ApplyScalingGroupResponseBody) SetRequestId(v string) *ApplyScalingGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *ApplyScalingGroupResponseBody) SetScalingGroupId(v string) *ApplyScalingGroupResponseBody {
	s.ScalingGroupId = &v
	return s
}

func (s *ApplyScalingGroupResponseBody) Validate() error {
	return dara.Validate(s)
}
