// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iLoginPolarClawChannelResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetApplicationId(v string) *LoginPolarClawChannelResponseBody
	GetApplicationId() *string
	SetCode(v int32) *LoginPolarClawChannelResponseBody
	GetCode() *int32
	SetMessage(v string) *LoginPolarClawChannelResponseBody
	GetMessage() *string
	SetOperation(v string) *LoginPolarClawChannelResponseBody
	GetOperation() *string
	SetRequestId(v string) *LoginPolarClawChannelResponseBody
	GetRequestId() *string
	SetState(v string) *LoginPolarClawChannelResponseBody
	GetState() *string
	SetTaskId(v string) *LoginPolarClawChannelResponseBody
	GetTaskId() *string
}

type LoginPolarClawChannelResponseBody struct {
	// example:
	//
	// pa-**************
	ApplicationId *string `json:"ApplicationId,omitempty" xml:"ApplicationId,omitempty"`
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// LoginPolarClawChannel
	Operation *string `json:"Operation,omitempty" xml:"Operation,omitempty"`
	// Id of the request
	//
	// example:
	//
	// CD3FA5F3-FAF3-44CA-AFFF-BAF869******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// running
	State *string `json:"State,omitempty" xml:"State,omitempty"`
	// example:
	//
	// dc6762fb-20ad-4796-84fe-5c5d0dc413ce
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s LoginPolarClawChannelResponseBody) String() string {
	return dara.Prettify(s)
}

func (s LoginPolarClawChannelResponseBody) GoString() string {
	return s.String()
}

func (s *LoginPolarClawChannelResponseBody) GetApplicationId() *string {
	return s.ApplicationId
}

func (s *LoginPolarClawChannelResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *LoginPolarClawChannelResponseBody) GetMessage() *string {
	return s.Message
}

func (s *LoginPolarClawChannelResponseBody) GetOperation() *string {
	return s.Operation
}

func (s *LoginPolarClawChannelResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *LoginPolarClawChannelResponseBody) GetState() *string {
	return s.State
}

func (s *LoginPolarClawChannelResponseBody) GetTaskId() *string {
	return s.TaskId
}

func (s *LoginPolarClawChannelResponseBody) SetApplicationId(v string) *LoginPolarClawChannelResponseBody {
	s.ApplicationId = &v
	return s
}

func (s *LoginPolarClawChannelResponseBody) SetCode(v int32) *LoginPolarClawChannelResponseBody {
	s.Code = &v
	return s
}

func (s *LoginPolarClawChannelResponseBody) SetMessage(v string) *LoginPolarClawChannelResponseBody {
	s.Message = &v
	return s
}

func (s *LoginPolarClawChannelResponseBody) SetOperation(v string) *LoginPolarClawChannelResponseBody {
	s.Operation = &v
	return s
}

func (s *LoginPolarClawChannelResponseBody) SetRequestId(v string) *LoginPolarClawChannelResponseBody {
	s.RequestId = &v
	return s
}

func (s *LoginPolarClawChannelResponseBody) SetState(v string) *LoginPolarClawChannelResponseBody {
	s.State = &v
	return s
}

func (s *LoginPolarClawChannelResponseBody) SetTaskId(v string) *LoginPolarClawChannelResponseBody {
	s.TaskId = &v
	return s
}

func (s *LoginPolarClawChannelResponseBody) Validate() error {
	return dara.Validate(s)
}
