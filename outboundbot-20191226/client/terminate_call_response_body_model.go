// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTerminateCallResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *TerminateCallResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *TerminateCallResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *TerminateCallResponseBody
	GetMessage() *string
	SetRequestId(v string) *TerminateCallResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *TerminateCallResponseBody
	GetSuccess() *bool
}

type TerminateCallResponseBody struct {
	// API status code
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
	// Indicates whether the request succeeded
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s TerminateCallResponseBody) String() string {
	return dara.Prettify(s)
}

func (s TerminateCallResponseBody) GoString() string {
	return s.String()
}

func (s *TerminateCallResponseBody) GetCode() *string {
	return s.Code
}

func (s *TerminateCallResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *TerminateCallResponseBody) GetMessage() *string {
	return s.Message
}

func (s *TerminateCallResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *TerminateCallResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *TerminateCallResponseBody) SetCode(v string) *TerminateCallResponseBody {
	s.Code = &v
	return s
}

func (s *TerminateCallResponseBody) SetHttpStatusCode(v int32) *TerminateCallResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *TerminateCallResponseBody) SetMessage(v string) *TerminateCallResponseBody {
	s.Message = &v
	return s
}

func (s *TerminateCallResponseBody) SetRequestId(v string) *TerminateCallResponseBody {
	s.RequestId = &v
	return s
}

func (s *TerminateCallResponseBody) SetSuccess(v bool) *TerminateCallResponseBody {
	s.Success = &v
	return s
}

func (s *TerminateCallResponseBody) Validate() error {
	return dara.Validate(s)
}
