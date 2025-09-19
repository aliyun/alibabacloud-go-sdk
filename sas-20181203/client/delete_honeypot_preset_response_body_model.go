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
	// 1A975D03-5F49-5354-B2CB-3918D5DA431A
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
