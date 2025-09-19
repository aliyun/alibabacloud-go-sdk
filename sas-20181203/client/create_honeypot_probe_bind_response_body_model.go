// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateHoneypotProbeBindResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CreateHoneypotProbeBindResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *CreateHoneypotProbeBindResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *CreateHoneypotProbeBindResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreateHoneypotProbeBindResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreateHoneypotProbeBindResponseBody
	GetSuccess() *bool
}

type CreateHoneypotProbeBindResponseBody struct {
	// The response code. The status code **200*	- indicates that the request was successful. Other status codes indicate that the request failed. You can identify the cause of the failure based on the status code.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The HTTP status code that is returned.
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
	// E10BAF1C-A6C5-51E2-866C-76D5922E****
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

func (s CreateHoneypotProbeBindResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateHoneypotProbeBindResponseBody) GoString() string {
	return s.String()
}

func (s *CreateHoneypotProbeBindResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreateHoneypotProbeBindResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *CreateHoneypotProbeBindResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateHoneypotProbeBindResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateHoneypotProbeBindResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateHoneypotProbeBindResponseBody) SetCode(v string) *CreateHoneypotProbeBindResponseBody {
	s.Code = &v
	return s
}

func (s *CreateHoneypotProbeBindResponseBody) SetHttpStatusCode(v int32) *CreateHoneypotProbeBindResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *CreateHoneypotProbeBindResponseBody) SetMessage(v string) *CreateHoneypotProbeBindResponseBody {
	s.Message = &v
	return s
}

func (s *CreateHoneypotProbeBindResponseBody) SetRequestId(v string) *CreateHoneypotProbeBindResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateHoneypotProbeBindResponseBody) SetSuccess(v bool) *CreateHoneypotProbeBindResponseBody {
	s.Success = &v
	return s
}

func (s *CreateHoneypotProbeBindResponseBody) Validate() error {
	return dara.Validate(s)
}
