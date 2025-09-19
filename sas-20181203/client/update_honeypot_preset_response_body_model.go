// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateHoneypotPresetResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *UpdateHoneypotPresetResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *UpdateHoneypotPresetResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *UpdateHoneypotPresetResponseBody
	GetMessage() *string
	SetRequestId(v string) *UpdateHoneypotPresetResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateHoneypotPresetResponseBody
	GetSuccess() *bool
}

type UpdateHoneypotPresetResponseBody struct {
	// The status code returned. The status code **200*	- indicates that the request was successful. Other status codes indicate that the request failed. You can identify the cause of the failure based on the status code.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The HTTP status code returned.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The message returned.
	//
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// 60922C83-6B19-5A57-8F13-4663C6D391F4
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

func (s UpdateHoneypotPresetResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateHoneypotPresetResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateHoneypotPresetResponseBody) GetCode() *string {
	return s.Code
}

func (s *UpdateHoneypotPresetResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *UpdateHoneypotPresetResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateHoneypotPresetResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateHoneypotPresetResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateHoneypotPresetResponseBody) SetCode(v string) *UpdateHoneypotPresetResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateHoneypotPresetResponseBody) SetHttpStatusCode(v int32) *UpdateHoneypotPresetResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *UpdateHoneypotPresetResponseBody) SetMessage(v string) *UpdateHoneypotPresetResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateHoneypotPresetResponseBody) SetRequestId(v string) *UpdateHoneypotPresetResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateHoneypotPresetResponseBody) SetSuccess(v bool) *UpdateHoneypotPresetResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateHoneypotPresetResponseBody) Validate() error {
	return dara.Validate(s)
}
