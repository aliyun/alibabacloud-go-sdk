// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRobotCallResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *RobotCallResponseBody
	GetCode() *string
	SetData(v string) *RobotCallResponseBody
	GetData() *string
	SetMessage(v string) *RobotCallResponseBody
	GetMessage() *string
	SetRequestId(v string) *RobotCallResponseBody
	GetRequestId() *string
}

type RobotCallResponseBody struct {
	// Request status code. A return value of "OK" indicates that the request succeeded.
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The unique receipt ID of this call, which is also the unique call ID.
	//
	// example:
	//
	// 1160128*****^10281427*****
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// Description of the status code.
	//
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Request ID.
	//
	// example:
	//
	// A90E4451-FED7-49D2-87C8-00700A8C4D0D
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RobotCallResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RobotCallResponseBody) GoString() string {
	return s.String()
}

func (s *RobotCallResponseBody) GetCode() *string {
	return s.Code
}

func (s *RobotCallResponseBody) GetData() *string {
	return s.Data
}

func (s *RobotCallResponseBody) GetMessage() *string {
	return s.Message
}

func (s *RobotCallResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RobotCallResponseBody) SetCode(v string) *RobotCallResponseBody {
	s.Code = &v
	return s
}

func (s *RobotCallResponseBody) SetData(v string) *RobotCallResponseBody {
	s.Data = &v
	return s
}

func (s *RobotCallResponseBody) SetMessage(v string) *RobotCallResponseBody {
	s.Message = &v
	return s
}

func (s *RobotCallResponseBody) SetRequestId(v string) *RobotCallResponseBody {
	s.RequestId = &v
	return s
}

func (s *RobotCallResponseBody) Validate() error {
	return dara.Validate(s)
}
