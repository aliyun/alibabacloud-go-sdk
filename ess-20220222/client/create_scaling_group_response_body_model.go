// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateScalingGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CreateScalingGroupResponseBody
	GetRequestId() *string
	SetScalingGroupId(v string) *CreateScalingGroupResponseBody
	GetScalingGroupId() *string
}

type CreateScalingGroupResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The ID of the scaling group.
	//
	// example:
	//
	// asg-bp14wlu85wrpchm0****
	ScalingGroupId *string `json:"ScalingGroupId,omitempty" xml:"ScalingGroupId,omitempty"`
}

func (s CreateScalingGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateScalingGroupResponseBody) GoString() string {
	return s.String()
}

func (s *CreateScalingGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateScalingGroupResponseBody) GetScalingGroupId() *string {
	return s.ScalingGroupId
}

func (s *CreateScalingGroupResponseBody) SetRequestId(v string) *CreateScalingGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateScalingGroupResponseBody) SetScalingGroupId(v string) *CreateScalingGroupResponseBody {
	s.ScalingGroupId = &v
	return s
}

func (s *CreateScalingGroupResponseBody) Validate() error {
	return dara.Validate(s)
}
