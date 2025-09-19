// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteHoneypotProbeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DeleteHoneypotProbeResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *DeleteHoneypotProbeResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *DeleteHoneypotProbeResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeleteHoneypotProbeResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeleteHoneypotProbeResponseBody
	GetSuccess() *bool
}

type DeleteHoneypotProbeResponseBody struct {
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
	// 7FD1C1DC-AA67-510A-A022-5D23310C3658
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

func (s DeleteHoneypotProbeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteHoneypotProbeResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteHoneypotProbeResponseBody) GetCode() *string {
	return s.Code
}

func (s *DeleteHoneypotProbeResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *DeleteHoneypotProbeResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteHoneypotProbeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteHoneypotProbeResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteHoneypotProbeResponseBody) SetCode(v string) *DeleteHoneypotProbeResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteHoneypotProbeResponseBody) SetHttpStatusCode(v int32) *DeleteHoneypotProbeResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DeleteHoneypotProbeResponseBody) SetMessage(v string) *DeleteHoneypotProbeResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteHoneypotProbeResponseBody) SetRequestId(v string) *DeleteHoneypotProbeResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteHoneypotProbeResponseBody) SetSuccess(v bool) *DeleteHoneypotProbeResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteHoneypotProbeResponseBody) Validate() error {
	return dara.Validate(s)
}
