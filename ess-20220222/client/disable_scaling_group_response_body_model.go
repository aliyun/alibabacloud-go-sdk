// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDisableScalingGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DisableScalingGroupResponseBody
	GetRequestId() *string
}

type DisableScalingGroupResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3C83E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DisableScalingGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DisableScalingGroupResponseBody) GoString() string {
	return s.String()
}

func (s *DisableScalingGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DisableScalingGroupResponseBody) SetRequestId(v string) *DisableScalingGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *DisableScalingGroupResponseBody) Validate() error {
	return dara.Validate(s)
}
