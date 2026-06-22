// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteHoneypotPresetResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DeleteHoneypotPresetResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *DeleteHoneypotPresetResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *DeleteHoneypotPresetResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeleteHoneypotPresetResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeleteHoneypotPresetResponseBody
	GetSuccess() *bool
}

type DeleteHoneypotPresetResponseBody struct {
	// The result code. A value of **200*	- indicates success. Any other value indicates failure. You can use this field to determine the cause of a failure.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The HTTP status code of the request.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The returned message.
	//
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request. The ID is a unique identifier that Alibaba Cloud generates for the request and can be used to troubleshoot issues.
	//
	// example:
	//
	// 1A975D03-5F49-5354-B2CB-3918D5DA431A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// - **true**: The request was successful.
	//
	// - **false**: The request failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteHoneypotPresetResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteHoneypotPresetResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteHoneypotPresetResponseBody) GetCode() *string {
	return s.Code
}

func (s *DeleteHoneypotPresetResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *DeleteHoneypotPresetResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteHoneypotPresetResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteHoneypotPresetResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteHoneypotPresetResponseBody) SetCode(v string) *DeleteHoneypotPresetResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteHoneypotPresetResponseBody) SetHttpStatusCode(v int32) *DeleteHoneypotPresetResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DeleteHoneypotPresetResponseBody) SetMessage(v string) *DeleteHoneypotPresetResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteHoneypotPresetResponseBody) SetRequestId(v string) *DeleteHoneypotPresetResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteHoneypotPresetResponseBody) SetSuccess(v bool) *DeleteHoneypotPresetResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteHoneypotPresetResponseBody) Validate() error {
	return dara.Validate(s)
}
