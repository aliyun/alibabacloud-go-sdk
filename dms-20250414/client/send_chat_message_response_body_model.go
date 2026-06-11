// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSendChatMessageResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *SendChatMessageResponseBodyData) *SendChatMessageResponseBody
	GetData() *SendChatMessageResponseBodyData
	SetErrorCode(v string) *SendChatMessageResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *SendChatMessageResponseBody
	GetErrorMessage() *string
	SetRequestId(v string) *SendChatMessageResponseBody
	GetRequestId() *string
	SetSuccess(v string) *SendChatMessageResponseBody
	GetSuccess() *string
}

type SendChatMessageResponseBody struct {
	// The response data.
	Data *SendChatMessageResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The error code. A value of `Success` indicates that the request was successful.
	//
	// example:
	//
	// Success
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The error message. This field is empty if the request is successful.
	//
	// example:
	//
	// ""
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The request ID.
	//
	// example:
	//
	// FE65CE1F-****-****-****-******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *string `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s SendChatMessageResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SendChatMessageResponseBody) GoString() string {
	return s.String()
}

func (s *SendChatMessageResponseBody) GetData() *SendChatMessageResponseBodyData {
	return s.Data
}

func (s *SendChatMessageResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *SendChatMessageResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *SendChatMessageResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SendChatMessageResponseBody) GetSuccess() *string {
	return s.Success
}

func (s *SendChatMessageResponseBody) SetData(v *SendChatMessageResponseBodyData) *SendChatMessageResponseBody {
	s.Data = v
	return s
}

func (s *SendChatMessageResponseBody) SetErrorCode(v string) *SendChatMessageResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *SendChatMessageResponseBody) SetErrorMessage(v string) *SendChatMessageResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *SendChatMessageResponseBody) SetRequestId(v string) *SendChatMessageResponseBody {
	s.RequestId = &v
	return s
}

func (s *SendChatMessageResponseBody) SetSuccess(v string) *SendChatMessageResponseBody {
	s.Success = &v
	return s
}

func (s *SendChatMessageResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type SendChatMessageResponseBodyData struct {
	// The agent ID.
	//
	// example:
	//
	// 3jqqdiuxun******
	AgentId *string `json:"AgentId,omitempty" xml:"AgentId,omitempty"`
	// Describes the result of the request.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The session ID.
	//
	// example:
	//
	// 6zbqbho********
	SessionId *string `json:"SessionId,omitempty" xml:"SessionId,omitempty"`
}

func (s SendChatMessageResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s SendChatMessageResponseBodyData) GoString() string {
	return s.String()
}

func (s *SendChatMessageResponseBodyData) GetAgentId() *string {
	return s.AgentId
}

func (s *SendChatMessageResponseBodyData) GetMessage() *string {
	return s.Message
}

func (s *SendChatMessageResponseBodyData) GetSessionId() *string {
	return s.SessionId
}

func (s *SendChatMessageResponseBodyData) SetAgentId(v string) *SendChatMessageResponseBodyData {
	s.AgentId = &v
	return s
}

func (s *SendChatMessageResponseBodyData) SetMessage(v string) *SendChatMessageResponseBodyData {
	s.Message = &v
	return s
}

func (s *SendChatMessageResponseBodyData) SetSessionId(v string) *SendChatMessageResponseBodyData {
	s.SessionId = &v
	return s
}

func (s *SendChatMessageResponseBodyData) Validate() error {
	return dara.Validate(s)
}
