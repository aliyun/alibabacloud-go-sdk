// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateChatSessionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *CreateChatSessionResponseBodyData) *CreateChatSessionResponseBody
	GetData() *CreateChatSessionResponseBodyData
	SetRequestId(v string) *CreateChatSessionResponseBody
	GetRequestId() *string
}

type CreateChatSessionResponseBody struct {
	// Returned data structure.
	Data *CreateChatSessionResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// ID of the request
	//
	// example:
	//
	// 9bc20a5a-b26b-4c28-922a-7cd10b61f96f
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s CreateChatSessionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateChatSessionResponseBody) GoString() string {
	return s.String()
}

func (s *CreateChatSessionResponseBody) GetData() *CreateChatSessionResponseBodyData {
	return s.Data
}

func (s *CreateChatSessionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateChatSessionResponseBody) SetData(v *CreateChatSessionResponseBodyData) *CreateChatSessionResponseBody {
	s.Data = v
	return s
}

func (s *CreateChatSessionResponseBody) SetRequestId(v string) *CreateChatSessionResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateChatSessionResponseBody) Validate() error {
	return dara.Validate(s)
}

type CreateChatSessionResponseBodyData struct {
	// Q&A session ID, used to record multiple Q&A sessions of the same user.
	//
	// example:
	//
	// 596ac39c-8855-4128-bad7-78aebeff48fc
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
