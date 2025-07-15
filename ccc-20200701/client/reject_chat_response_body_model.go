// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRejectChatResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *RejectChatResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *RejectChatResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *RejectChatResponseBody
	GetMessage() *string
	SetRequestId(v string) *RejectChatResponseBody
	GetRequestId() *string
}

type RejectChatResponseBody struct {
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int32  `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// B59382D2-5755-4C6D-861F-FA2AAD8F89F7
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RejectChatResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RejectChatResponseBody) GoString() string {
	return s.String()
}

func (s *RejectChatResponseBody) GetCode() *string {
	return s.Code
}

func (s *RejectChatResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *RejectChatResponseBody) GetMessage() *string {
	return s.Message
}

func (s *RejectChatResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RejectChatResponseBody) SetCode(v string) *RejectChatResponseBody {
	s.Code = &v
	return s
}

func (s *RejectChatResponseBody) SetHttpStatusCode(v int32) *RejectChatResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *RejectChatResponseBody) SetMessage(v string) *RejectChatResponseBody {
	s.Message = &v
	return s
}

func (s *RejectChatResponseBody) SetRequestId(v string) *RejectChatResponseBody {
	s.RequestId = &v
	return s
}

func (s *RejectChatResponseBody) Validate() error {
	return dara.Validate(s)
}
