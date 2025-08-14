// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartRtcRobotInstanceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *StartRtcRobotInstanceResponseBody
	GetInstanceId() *string
	SetRequestId(v string) *StartRtcRobotInstanceResponseBody
	GetRequestId() *string
}

type StartRtcRobotInstanceResponseBody struct {
	// example:
	//
	// 727dc0e296014bb58670940a3da95592
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 11DE0AB3-603B-5055-8A72-9C424854F983
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s StartRtcRobotInstanceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s StartRtcRobotInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *StartRtcRobotInstanceResponseBody) GetInstanceId() *string {
	return s.InstanceId
}

func (s *StartRtcRobotInstanceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *StartRtcRobotInstanceResponseBody) SetInstanceId(v string) *StartRtcRobotInstanceResponseBody {
	s.InstanceId = &v
	return s
}

func (s *StartRtcRobotInstanceResponseBody) SetRequestId(v string) *StartRtcRobotInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *StartRtcRobotInstanceResponseBody) Validate() error {
	return dara.Validate(s)
}
