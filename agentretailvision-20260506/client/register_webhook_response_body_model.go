// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRegisterWebhookResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *RegisterWebhookResponseBody
	GetCode() *string
	SetMessage(v string) *RegisterWebhookResponseBody
	GetMessage() *string
	SetRequestId(v string) *RegisterWebhookResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *RegisterWebhookResponseBody
	GetSuccess() *bool
}

type RegisterWebhookResponseBody struct {
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// E1AD60F1-BAC7-546B-9533-E7AD02B16E3F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s RegisterWebhookResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RegisterWebhookResponseBody) GoString() string {
	return s.String()
}

func (s *RegisterWebhookResponseBody) GetCode() *string {
	return s.Code
}

func (s *RegisterWebhookResponseBody) GetMessage() *string {
	return s.Message
}

func (s *RegisterWebhookResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RegisterWebhookResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *RegisterWebhookResponseBody) SetCode(v string) *RegisterWebhookResponseBody {
	s.Code = &v
	return s
}

func (s *RegisterWebhookResponseBody) SetMessage(v string) *RegisterWebhookResponseBody {
	s.Message = &v
	return s
}

func (s *RegisterWebhookResponseBody) SetRequestId(v string) *RegisterWebhookResponseBody {
	s.RequestId = &v
	return s
}

func (s *RegisterWebhookResponseBody) SetSuccess(v bool) *RegisterWebhookResponseBody {
	s.Success = &v
	return s
}

func (s *RegisterWebhookResponseBody) Validate() error {
	return dara.Validate(s)
}
