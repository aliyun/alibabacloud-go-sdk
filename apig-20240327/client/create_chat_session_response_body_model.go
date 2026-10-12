// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateChatSessionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CreateChatSessionResponseBody
	GetCode() *string
	SetData(v *CreateChatSessionResponseBodyData) *CreateChatSessionResponseBody
	GetData() *CreateChatSessionResponseBodyData
	SetMessage(v string) *CreateChatSessionResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreateChatSessionResponseBody
	GetRequestId() *string
}

type CreateChatSessionResponseBody struct {
	// The response status code.
	//
	// example:
	//
	// Ok
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The response data.
	Data *CreateChatSessionResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// The response message.
	//
	// example:
	//
	// success
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 393E2630-DBE7-5221-AB35-9E740675491A
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s CreateChatSessionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateChatSessionResponseBody) GoString() string {
	return s.String()
}

func (s *CreateChatSessionResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreateChatSessionResponseBody) GetData() *CreateChatSessionResponseBodyData {
	return s.Data
}

func (s *CreateChatSessionResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateChatSessionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateChatSessionResponseBody) SetCode(v string) *CreateChatSessionResponseBody {
	s.Code = &v
	return s
}

func (s *CreateChatSessionResponseBody) SetData(v *CreateChatSessionResponseBodyData) *CreateChatSessionResponseBody {
	s.Data = v
	return s
}

func (s *CreateChatSessionResponseBody) SetMessage(v string) *CreateChatSessionResponseBody {
	s.Message = &v
	return s
}

func (s *CreateChatSessionResponseBody) SetRequestId(v string) *CreateChatSessionResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateChatSessionResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateChatSessionResponseBodyData struct {
	// The chat session ID. You can use this ID to locate the session in API operations such as UpdateChatSession and DeleteChatSession.
	//
	// example:
	//
	// ses-xxxxxxxx
	SessionId *string `json:"sessionId,omitempty" xml:"sessionId,omitempty"`
}

func (s CreateChatSessionResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s CreateChatSessionResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateChatSessionResponseBodyData) GetSessionId() *string {
	return s.SessionId
}

func (s *CreateChatSessionResponseBodyData) SetSessionId(v string) *CreateChatSessionResponseBodyData {
	s.SessionId = &v
	return s
}

func (s *CreateChatSessionResponseBodyData) Validate() error {
	return dara.Validate(s)
}
