// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateRtcRobotInstanceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateRtcRobotInstanceResponseBody
	GetRequestId() *string
}

type UpdateRtcRobotInstanceResponseBody struct {
	// Id of the request
	//
	// example:
	//
	// 7707F0A2-C6FD-5959-87EB-7C4D02384FD4
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateRtcRobotInstanceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateRtcRobotInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateRtcRobotInstanceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateRtcRobotInstanceResponseBody) SetRequestId(v string) *UpdateRtcRobotInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateRtcRobotInstanceResponseBody) Validate() error {
	return dara.Validate(s)
}
