// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRevertAssignedSessionGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *RevertAssignedSessionGroupResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *RevertAssignedSessionGroupResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *RevertAssignedSessionGroupResponseBody
	GetMessage() *string
	SetMessages(v *RevertAssignedSessionGroupResponseBodyMessages) *RevertAssignedSessionGroupResponseBody
	GetMessages() *RevertAssignedSessionGroupResponseBodyMessages
	SetRequestId(v string) *RevertAssignedSessionGroupResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *RevertAssignedSessionGroupResponseBody
	GetSuccess() *bool
}

type RevertAssignedSessionGroupResponseBody struct {
	// The result code. A value of **200*	- indicates success. Other values indicate failure. Use this code to identify the cause of a failed request.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// If the request fails, this parameter provides error details. If the request succeeds, the value is \\`successful\\`.
	//
	// example:
	//
	// successful
	Message  *string                                         `json:"Message,omitempty" xml:"Message,omitempty"`
	Messages *RevertAssignedSessionGroupResponseBodyMessages `json:"Messages,omitempty" xml:"Messages,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 106C6CA0-282D-4AF7-85F0-D2D24F4CE647
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// The caller can use this field to determine whether the request succeeded.
	//
	// - **true**: The request was successful.
	//
	// - false or **null**: The request failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s RevertAssignedSessionGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RevertAssignedSessionGroupResponseBody) GoString() string {
	return s.String()
}

func (s *RevertAssignedSessionGroupResponseBody) GetCode() *string {
	return s.Code
}

func (s *RevertAssignedSessionGroupResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *RevertAssignedSessionGroupResponseBody) GetMessage() *string {
	return s.Message
}

func (s *RevertAssignedSessionGroupResponseBody) GetMessages() *RevertAssignedSessionGroupResponseBodyMessages {
	return s.Messages
}

func (s *RevertAssignedSessionGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RevertAssignedSessionGroupResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *RevertAssignedSessionGroupResponseBody) SetCode(v string) *RevertAssignedSessionGroupResponseBody {
	s.Code = &v
	return s
}

func (s *RevertAssignedSessionGroupResponseBody) SetHttpStatusCode(v int32) *RevertAssignedSessionGroupResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *RevertAssignedSessionGroupResponseBody) SetMessage(v string) *RevertAssignedSessionGroupResponseBody {
	s.Message = &v
	return s
}

func (s *RevertAssignedSessionGroupResponseBody) SetMessages(v *RevertAssignedSessionGroupResponseBodyMessages) *RevertAssignedSessionGroupResponseBody {
	s.Messages = v
	return s
}

func (s *RevertAssignedSessionGroupResponseBody) SetRequestId(v string) *RevertAssignedSessionGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *RevertAssignedSessionGroupResponseBody) SetSuccess(v bool) *RevertAssignedSessionGroupResponseBody {
	s.Success = &v
	return s
}

func (s *RevertAssignedSessionGroupResponseBody) Validate() error {
	if s.Messages != nil {
		if err := s.Messages.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type RevertAssignedSessionGroupResponseBodyMessages struct {
	Message []*string `json:"Message,omitempty" xml:"Message,omitempty" type:"Repeated"`
}

func (s RevertAssignedSessionGroupResponseBodyMessages) String() string {
	return dara.Prettify(s)
}

func (s RevertAssignedSessionGroupResponseBodyMessages) GoString() string {
	return s.String()
}

func (s *RevertAssignedSessionGroupResponseBodyMessages) GetMessage() []*string {
	return s.Message
}

func (s *RevertAssignedSessionGroupResponseBodyMessages) SetMessage(v []*string) *RevertAssignedSessionGroupResponseBodyMessages {
	s.Message = v
	return s
}

func (s *RevertAssignedSessionGroupResponseBodyMessages) Validate() error {
	return dara.Validate(s)
}
