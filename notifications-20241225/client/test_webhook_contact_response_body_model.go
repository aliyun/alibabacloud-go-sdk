// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTestWebhookContactResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *TestWebhookContactResponseBody
	GetCode() *string
	SetData(v string) *TestWebhookContactResponseBody
	GetData() *string
	SetHttpCode(v int32) *TestWebhookContactResponseBody
	GetHttpCode() *int32
	SetMessage(v string) *TestWebhookContactResponseBody
	GetMessage() *string
	SetRequestId(v string) *TestWebhookContactResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *TestWebhookContactResponseBody
	GetSuccess() *bool
}

type TestWebhookContactResponseBody struct {
	// The business status code.
	//
	// example:
	//
	// SUCCESS
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The query result.
	//
	// example:
	//
	// /
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpCode *int32 `json:"HttpCode,omitempty" xml:"HttpCode,omitempty"`
	// The business message.
	//
	// example:
	//
	// Successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// A5F62766-1C2F-1F56-A39D-63E3D30F0633
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the call was successful. Valid values:
	//
	// - true: The call was successful.
	//
	// - false: The call failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s TestWebhookContactResponseBody) String() string {
	return dara.Prettify(s)
}

func (s TestWebhookContactResponseBody) GoString() string {
	return s.String()
}

func (s *TestWebhookContactResponseBody) GetCode() *string {
	return s.Code
}

func (s *TestWebhookContactResponseBody) GetData() *string {
	return s.Data
}

func (s *TestWebhookContactResponseBody) GetHttpCode() *int32 {
	return s.HttpCode
}

func (s *TestWebhookContactResponseBody) GetMessage() *string {
	return s.Message
}

func (s *TestWebhookContactResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *TestWebhookContactResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *TestWebhookContactResponseBody) SetCode(v string) *TestWebhookContactResponseBody {
	s.Code = &v
	return s
}

func (s *TestWebhookContactResponseBody) SetData(v string) *TestWebhookContactResponseBody {
	s.Data = &v
	return s
}

func (s *TestWebhookContactResponseBody) SetHttpCode(v int32) *TestWebhookContactResponseBody {
	s.HttpCode = &v
	return s
}

func (s *TestWebhookContactResponseBody) SetMessage(v string) *TestWebhookContactResponseBody {
	s.Message = &v
	return s
}

func (s *TestWebhookContactResponseBody) SetRequestId(v string) *TestWebhookContactResponseBody {
	s.RequestId = &v
	return s
}

func (s *TestWebhookContactResponseBody) SetSuccess(v bool) *TestWebhookContactResponseBody {
	s.Success = &v
	return s
}

func (s *TestWebhookContactResponseBody) Validate() error {
	return dara.Validate(s)
}
