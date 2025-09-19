// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateHoneypotProbeBindResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *UpdateHoneypotProbeBindResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *UpdateHoneypotProbeBindResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *UpdateHoneypotProbeBindResponseBody
	GetMessage() *string
	SetRequestId(v string) *UpdateHoneypotProbeBindResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateHoneypotProbeBindResponseBody
	GetSuccess() *bool
}

type UpdateHoneypotProbeBindResponseBody struct {
	// The response code. The status code **200*	- indicates that the request was successful. Other status codes indicate that the request failed. You can identify the cause of the failure based on the status code.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
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
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 425D9617-4F4F-571E-A9CF-0245C9FC****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateHoneypotProbeBindResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateHoneypotProbeBindResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateHoneypotProbeBindResponseBody) GetCode() *string {
	return s.Code
}

func (s *UpdateHoneypotProbeBindResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *UpdateHoneypotProbeBindResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateHoneypotProbeBindResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateHoneypotProbeBindResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateHoneypotProbeBindResponseBody) SetCode(v string) *UpdateHoneypotProbeBindResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateHoneypotProbeBindResponseBody) SetHttpStatusCode(v int32) *UpdateHoneypotProbeBindResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *UpdateHoneypotProbeBindResponseBody) SetMessage(v string) *UpdateHoneypotProbeBindResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateHoneypotProbeBindResponseBody) SetRequestId(v string) *UpdateHoneypotProbeBindResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateHoneypotProbeBindResponseBody) SetSuccess(v bool) *UpdateHoneypotProbeBindResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateHoneypotProbeBindResponseBody) Validate() error {
	return dara.Validate(s)
}
