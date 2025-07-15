// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddFeedbackResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *AddFeedbackResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *AddFeedbackResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *AddFeedbackResponseBody
	GetMessage() *string
	SetRequestId(v string) *AddFeedbackResponseBody
	GetRequestId() *string
}

type AddFeedbackResponseBody struct {
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
	// C81FD1A5-4B99-470A-A527-D80150228784
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AddFeedbackResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AddFeedbackResponseBody) GoString() string {
	return s.String()
}

func (s *AddFeedbackResponseBody) GetCode() *string {
	return s.Code
}

func (s *AddFeedbackResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *AddFeedbackResponseBody) GetMessage() *string {
	return s.Message
}

func (s *AddFeedbackResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AddFeedbackResponseBody) SetCode(v string) *AddFeedbackResponseBody {
	s.Code = &v
	return s
}

func (s *AddFeedbackResponseBody) SetHttpStatusCode(v int32) *AddFeedbackResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *AddFeedbackResponseBody) SetMessage(v string) *AddFeedbackResponseBody {
	s.Message = &v
	return s
}

func (s *AddFeedbackResponseBody) SetRequestId(v string) *AddFeedbackResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddFeedbackResponseBody) Validate() error {
	return dara.Validate(s)
}
