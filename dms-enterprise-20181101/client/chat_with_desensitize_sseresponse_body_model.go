// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iChatWithDesensitizeSSEResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v string) *ChatWithDesensitizeSSEResponseBody
	GetData() *string
	SetErrorCode(v string) *ChatWithDesensitizeSSEResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *ChatWithDesensitizeSSEResponseBody
	GetErrorMessage() *string
	SetRequestId(v string) *ChatWithDesensitizeSSEResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ChatWithDesensitizeSSEResponseBody
	GetSuccess() *bool
}

type ChatWithDesensitizeSSEResponseBody struct {
	// The data returned.
	//
	// example:
	//
	// true
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// The error code.
	//
	// example:
	//
	// UnknownError
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The error message.
	//
	// example:
	//
	// UnknownError
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// Request ID.
	//
	// example:
	//
	// 283C461F-11D8-48AA-B695-DF092DA32AF3
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- **true**: The request was successful.
	//
	// 	- **false**: The request failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ChatWithDesensitizeSSEResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ChatWithDesensitizeSSEResponseBody) GoString() string {
	return s.String()
}

func (s *ChatWithDesensitizeSSEResponseBody) GetData() *string {
	return s.Data
}

func (s *ChatWithDesensitizeSSEResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *ChatWithDesensitizeSSEResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *ChatWithDesensitizeSSEResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ChatWithDesensitizeSSEResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ChatWithDesensitizeSSEResponseBody) SetData(v string) *ChatWithDesensitizeSSEResponseBody {
	s.Data = &v
	return s
}

func (s *ChatWithDesensitizeSSEResponseBody) SetErrorCode(v string) *ChatWithDesensitizeSSEResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ChatWithDesensitizeSSEResponseBody) SetErrorMessage(v string) *ChatWithDesensitizeSSEResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *ChatWithDesensitizeSSEResponseBody) SetRequestId(v string) *ChatWithDesensitizeSSEResponseBody {
	s.RequestId = &v
	return s
}

func (s *ChatWithDesensitizeSSEResponseBody) SetSuccess(v bool) *ChatWithDesensitizeSSEResponseBody {
	s.Success = &v
	return s
}

func (s *ChatWithDesensitizeSSEResponseBody) Validate() error {
	return dara.Validate(s)
}
