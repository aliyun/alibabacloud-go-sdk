// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStopRtcRobotInstanceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *StopRtcRobotInstanceResponseBody
	GetRequestId() *string
}

type StopRtcRobotInstanceResponseBody struct {
	// Id of the request
	//
	// example:
	//
	// AC84E5DD-AB56-56C0-A992-07ECB82008CA
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s StopRtcRobotInstanceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s StopRtcRobotInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *StopRtcRobotInstanceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *StopRtcRobotInstanceResponseBody) SetRequestId(v string) *StopRtcRobotInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *StopRtcRobotInstanceResponseBody) Validate() error {
	return dara.Validate(s)
}
