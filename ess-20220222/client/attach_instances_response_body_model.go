// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAttachInstancesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *AttachInstancesResponseBody
	GetRequestId() *string
	SetScalingActivityId(v string) *AttachInstancesResponseBody
	GetScalingActivityId() *string
}

type AttachInstancesResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3C83E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The ID of the scaling activity.
	//
	// example:
	//
	// asa-bp1crxor24s28xf1****
	ScalingActivityId *string `json:"ScalingActivityId,omitempty" xml:"ScalingActivityId,omitempty"`
}

func (s AttachInstancesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AttachInstancesResponseBody) GoString() string {
	return s.String()
}

func (s *AttachInstancesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AttachInstancesResponseBody) GetScalingActivityId() *string {
	return s.ScalingActivityId
}

func (s *AttachInstancesResponseBody) SetRequestId(v string) *AttachInstancesResponseBody {
	s.RequestId = &v
	return s
}

func (s *AttachInstancesResponseBody) SetScalingActivityId(v string) *AttachInstancesResponseBody {
	s.ScalingActivityId = &v
	return s
}

func (s *AttachInstancesResponseBody) Validate() error {
	return dara.Validate(s)
}
