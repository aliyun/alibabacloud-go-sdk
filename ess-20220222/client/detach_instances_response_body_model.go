// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDetachInstancesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DetachInstancesResponseBody
	GetRequestId() *string
	SetScalingActivityId(v string) *DetachInstancesResponseBody
	GetScalingActivityId() *string
}

type DetachInstancesResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The ID of the scaling activity.
	//
	// example:
	//
	// asa-bp1gbswjhjrw8tko****
	ScalingActivityId *string `json:"ScalingActivityId,omitempty" xml:"ScalingActivityId,omitempty"`
}

func (s DetachInstancesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DetachInstancesResponseBody) GoString() string {
	return s.String()
}

func (s *DetachInstancesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DetachInstancesResponseBody) GetScalingActivityId() *string {
	return s.ScalingActivityId
}

func (s *DetachInstancesResponseBody) SetRequestId(v string) *DetachInstancesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DetachInstancesResponseBody) SetScalingActivityId(v string) *DetachInstancesResponseBody {
	s.ScalingActivityId = &v
	return s
}

func (s *DetachInstancesResponseBody) Validate() error {
	return dara.Validate(s)
}
