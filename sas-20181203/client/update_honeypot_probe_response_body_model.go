// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateHoneypotProbeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *UpdateHoneypotProbeResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *UpdateHoneypotProbeResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *UpdateHoneypotProbeResponseBody
	GetMessage() *string
	SetRequestId(v string) *UpdateHoneypotProbeResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateHoneypotProbeResponseBody
	GetSuccess() *bool
}

type UpdateHoneypotProbeResponseBody struct {
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
	// AB62FFAA-E1A5-5D7C-8D97-0F16C6A6F520
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

func (s UpdateHoneypotProbeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateHoneypotProbeResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateHoneypotProbeResponseBody) GetCode() *string {
	return s.Code
}

func (s *UpdateHoneypotProbeResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *UpdateHoneypotProbeResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateHoneypotProbeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateHoneypotProbeResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateHoneypotProbeResponseBody) SetCode(v string) *UpdateHoneypotProbeResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateHoneypotProbeResponseBody) SetHttpStatusCode(v int32) *UpdateHoneypotProbeResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *UpdateHoneypotProbeResponseBody) SetMessage(v string) *UpdateHoneypotProbeResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateHoneypotProbeResponseBody) SetRequestId(v string) *UpdateHoneypotProbeResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateHoneypotProbeResponseBody) SetSuccess(v bool) *UpdateHoneypotProbeResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateHoneypotProbeResponseBody) Validate() error {
	return dara.Validate(s)
}
