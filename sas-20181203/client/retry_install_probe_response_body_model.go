// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRetryInstallProbeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *RetryInstallProbeResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *RetryInstallProbeResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *RetryInstallProbeResponseBody
	GetMessage() *string
	SetRequestId(v string) *RetryInstallProbeResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *RetryInstallProbeResponseBody
	GetSuccess() *bool
}

type RetryInstallProbeResponseBody struct {
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
	// A4247271-7C31-5A54-9EA1-658D96ED****
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

func (s RetryInstallProbeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RetryInstallProbeResponseBody) GoString() string {
	return s.String()
}

func (s *RetryInstallProbeResponseBody) GetCode() *string {
	return s.Code
}

func (s *RetryInstallProbeResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *RetryInstallProbeResponseBody) GetMessage() *string {
	return s.Message
}

func (s *RetryInstallProbeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RetryInstallProbeResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *RetryInstallProbeResponseBody) SetCode(v string) *RetryInstallProbeResponseBody {
	s.Code = &v
	return s
}

func (s *RetryInstallProbeResponseBody) SetHttpStatusCode(v int32) *RetryInstallProbeResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *RetryInstallProbeResponseBody) SetMessage(v string) *RetryInstallProbeResponseBody {
	s.Message = &v
	return s
}

func (s *RetryInstallProbeResponseBody) SetRequestId(v string) *RetryInstallProbeResponseBody {
	s.RequestId = &v
	return s
}

func (s *RetryInstallProbeResponseBody) SetSuccess(v bool) *RetryInstallProbeResponseBody {
	s.Success = &v
	return s
}

func (s *RetryInstallProbeResponseBody) Validate() error {
	return dara.Validate(s)
}
