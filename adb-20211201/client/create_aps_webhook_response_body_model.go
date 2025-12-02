// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateApsWebhookResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CreateApsWebhookResponseBody
	GetCode() *string
	SetData(v string) *CreateApsWebhookResponseBody
	GetData() *string
	SetHttpStatusCode(v int32) *CreateApsWebhookResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *CreateApsWebhookResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreateApsWebhookResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreateApsWebhookResponseBody
	GetSuccess() *bool
}

type CreateApsWebhookResponseBody struct {
	// The code returned for the request.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned result.
	//
	// example:
	//
	// []
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The returned message.
	//
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The unique ID of the request.
	//
	// example:
	//
	// 1234567890abcdef
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateApsWebhookResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateApsWebhookResponseBody) GoString() string {
	return s.String()
}

func (s *CreateApsWebhookResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreateApsWebhookResponseBody) GetData() *string {
	return s.Data
}

func (s *CreateApsWebhookResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *CreateApsWebhookResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateApsWebhookResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateApsWebhookResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateApsWebhookResponseBody) SetCode(v string) *CreateApsWebhookResponseBody {
	s.Code = &v
	return s
}

func (s *CreateApsWebhookResponseBody) SetData(v string) *CreateApsWebhookResponseBody {
	s.Data = &v
	return s
}

func (s *CreateApsWebhookResponseBody) SetHttpStatusCode(v int32) *CreateApsWebhookResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *CreateApsWebhookResponseBody) SetMessage(v string) *CreateApsWebhookResponseBody {
	s.Message = &v
	return s
}

func (s *CreateApsWebhookResponseBody) SetRequestId(v string) *CreateApsWebhookResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateApsWebhookResponseBody) SetSuccess(v bool) *CreateApsWebhookResponseBody {
	s.Success = &v
	return s
}

func (s *CreateApsWebhookResponseBody) Validate() error {
	return dara.Validate(s)
}
