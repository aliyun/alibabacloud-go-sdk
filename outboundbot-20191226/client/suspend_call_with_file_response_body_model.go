// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSuspendCallWithFileResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *SuspendCallWithFileResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *SuspendCallWithFileResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *SuspendCallWithFileResponseBody
	GetMessage() *string
	SetRequestId(v string) *SuspendCallWithFileResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *SuspendCallWithFileResponseBody
	GetSuccess() *bool
}

type SuspendCallWithFileResponseBody struct {
	// Response code
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// HTTP status code
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// Response message
	//
	// example:
	//
	// Success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Request ID
	//
	// example:
	//
	// 254EB995-DEDF-48A4-9101-9CA5B72FFBCC
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the operation succeeded
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s SuspendCallWithFileResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SuspendCallWithFileResponseBody) GoString() string {
	return s.String()
}

func (s *SuspendCallWithFileResponseBody) GetCode() *string {
	return s.Code
}

func (s *SuspendCallWithFileResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *SuspendCallWithFileResponseBody) GetMessage() *string {
	return s.Message
}

func (s *SuspendCallWithFileResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SuspendCallWithFileResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *SuspendCallWithFileResponseBody) SetCode(v string) *SuspendCallWithFileResponseBody {
	s.Code = &v
	return s
}

func (s *SuspendCallWithFileResponseBody) SetHttpStatusCode(v int32) *SuspendCallWithFileResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *SuspendCallWithFileResponseBody) SetMessage(v string) *SuspendCallWithFileResponseBody {
	s.Message = &v
	return s
}

func (s *SuspendCallWithFileResponseBody) SetRequestId(v string) *SuspendCallWithFileResponseBody {
	s.RequestId = &v
	return s
}

func (s *SuspendCallWithFileResponseBody) SetSuccess(v bool) *SuspendCallWithFileResponseBody {
	s.Success = &v
	return s
}

func (s *SuspendCallWithFileResponseBody) Validate() error {
	return dara.Validate(s)
}
