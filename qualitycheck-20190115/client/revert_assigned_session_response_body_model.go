// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRevertAssignedSessionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *RevertAssignedSessionResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *RevertAssignedSessionResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *RevertAssignedSessionResponseBody
	GetMessage() *string
	SetMessages(v *RevertAssignedSessionResponseBodyMessages) *RevertAssignedSessionResponseBody
	GetMessages() *RevertAssignedSessionResponseBodyMessages
	SetRequestId(v string) *RevertAssignedSessionResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *RevertAssignedSessionResponseBody
	GetSuccess() *bool
}

type RevertAssignedSessionResponseBody struct {
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// successful
	Message  *string                                    `json:"Message,omitempty" xml:"Message,omitempty"`
	Messages *RevertAssignedSessionResponseBodyMessages `json:"Messages,omitempty" xml:"Messages,omitempty" type:"Struct"`
	// example:
	//
	// 82C91484-B2D5-4D2A-A21F-A6D73F4D55C6
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s RevertAssignedSessionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RevertAssignedSessionResponseBody) GoString() string {
	return s.String()
}

func (s *RevertAssignedSessionResponseBody) GetCode() *string {
	return s.Code
}

func (s *RevertAssignedSessionResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *RevertAssignedSessionResponseBody) GetMessage() *string {
	return s.Message
}

func (s *RevertAssignedSessionResponseBody) GetMessages() *RevertAssignedSessionResponseBodyMessages {
	return s.Messages
}

func (s *RevertAssignedSessionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RevertAssignedSessionResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *RevertAssignedSessionResponseBody) SetCode(v string) *RevertAssignedSessionResponseBody {
	s.Code = &v
	return s
}

func (s *RevertAssignedSessionResponseBody) SetHttpStatusCode(v int32) *RevertAssignedSessionResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *RevertAssignedSessionResponseBody) SetMessage(v string) *RevertAssignedSessionResponseBody {
	s.Message = &v
	return s
}

func (s *RevertAssignedSessionResponseBody) SetMessages(v *RevertAssignedSessionResponseBodyMessages) *RevertAssignedSessionResponseBody {
	s.Messages = v
	return s
}

func (s *RevertAssignedSessionResponseBody) SetRequestId(v string) *RevertAssignedSessionResponseBody {
	s.RequestId = &v
	return s
}

func (s *RevertAssignedSessionResponseBody) SetSuccess(v bool) *RevertAssignedSessionResponseBody {
	s.Success = &v
	return s
}

func (s *RevertAssignedSessionResponseBody) Validate() error {
	return dara.Validate(s)
}

type RevertAssignedSessionResponseBodyMessages struct {
	Message []*string `json:"Message,omitempty" xml:"Message,omitempty" type:"Repeated"`
}

func (s RevertAssignedSessionResponseBodyMessages) String() string {
	return dara.Prettify(s)
}

func (s RevertAssignedSessionResponseBodyMessages) GoString() string {
	return s.String()
}

func (s *RevertAssignedSessionResponseBodyMessages) GetMessage() []*string {
	return s.Message
}

func (s *RevertAssignedSessionResponseBodyMessages) SetMessage(v []*string) *RevertAssignedSessionResponseBodyMessages {
	s.Message = v
	return s
}

func (s *RevertAssignedSessionResponseBodyMessages) Validate() error {
	return dara.Validate(s)
}
