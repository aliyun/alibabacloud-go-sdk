// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteHoneypotProbeBindResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DeleteHoneypotProbeBindResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *DeleteHoneypotProbeBindResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *DeleteHoneypotProbeBindResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeleteHoneypotProbeBindResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeleteHoneypotProbeBindResponseBody
	GetSuccess() *bool
}

type DeleteHoneypotProbeBindResponseBody struct {
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
	// 571B2642-BF51-5BDD-906B-D2340DB9****
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

func (s DeleteHoneypotProbeBindResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteHoneypotProbeBindResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteHoneypotProbeBindResponseBody) GetCode() *string {
	return s.Code
}

func (s *DeleteHoneypotProbeBindResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *DeleteHoneypotProbeBindResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteHoneypotProbeBindResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteHoneypotProbeBindResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteHoneypotProbeBindResponseBody) SetCode(v string) *DeleteHoneypotProbeBindResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteHoneypotProbeBindResponseBody) SetHttpStatusCode(v int32) *DeleteHoneypotProbeBindResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DeleteHoneypotProbeBindResponseBody) SetMessage(v string) *DeleteHoneypotProbeBindResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteHoneypotProbeBindResponseBody) SetRequestId(v string) *DeleteHoneypotProbeBindResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteHoneypotProbeBindResponseBody) SetSuccess(v bool) *DeleteHoneypotProbeBindResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteHoneypotProbeBindResponseBody) Validate() error {
	return dara.Validate(s)
}
