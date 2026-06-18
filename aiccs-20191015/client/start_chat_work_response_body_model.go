// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartChatWorkResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *StartChatWorkResponseBody
	GetCode() *string
	SetData(v string) *StartChatWorkResponseBody
	GetData() *string
	SetHttpStatusCode(v int32) *StartChatWorkResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *StartChatWorkResponseBody
	GetMessage() *string
	SetRequestId(v string) *StartChatWorkResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *StartChatWorkResponseBody
	GetSuccess() *bool
}

type StartChatWorkResponseBody struct {
	// Status code. A return value of 200 indicates that the request succeeded.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// A value of **1*	- indicates that the user has switched to the working status.
	//
	// example:
	//
	// 1
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// Description of the status code.
	//
	// example:
	//
	// xxxx
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Request ID.
	//
	// example:
	//
	// EE338D98-9BD3-4413-B165
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the API call succeeded.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s StartChatWorkResponseBody) String() string {
	return dara.Prettify(s)
}

func (s StartChatWorkResponseBody) GoString() string {
	return s.String()
}

func (s *StartChatWorkResponseBody) GetCode() *string {
	return s.Code
}

func (s *StartChatWorkResponseBody) GetData() *string {
	return s.Data
}

func (s *StartChatWorkResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *StartChatWorkResponseBody) GetMessage() *string {
	return s.Message
}

func (s *StartChatWorkResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *StartChatWorkResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *StartChatWorkResponseBody) SetCode(v string) *StartChatWorkResponseBody {
	s.Code = &v
	return s
}

func (s *StartChatWorkResponseBody) SetData(v string) *StartChatWorkResponseBody {
	s.Data = &v
	return s
}

func (s *StartChatWorkResponseBody) SetHttpStatusCode(v int32) *StartChatWorkResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *StartChatWorkResponseBody) SetMessage(v string) *StartChatWorkResponseBody {
	s.Message = &v
	return s
}

func (s *StartChatWorkResponseBody) SetRequestId(v string) *StartChatWorkResponseBody {
	s.RequestId = &v
	return s
}

func (s *StartChatWorkResponseBody) SetSuccess(v bool) *StartChatWorkResponseBody {
	s.Success = &v
	return s
}

func (s *StartChatWorkResponseBody) Validate() error {
	return dara.Validate(s)
}
