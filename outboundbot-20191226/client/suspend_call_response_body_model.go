// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSuspendCallResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *SuspendCallResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *SuspendCallResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *SuspendCallResponseBody
	GetMessage() *string
	SetRequestId(v string) *SuspendCallResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *SuspendCallResponseBody
	GetSuccess() *bool
}

type SuspendCallResponseBody struct {
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// Success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 254EB995-DEDF-48A4-9101-9CA5B72FFBCC
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s SuspendCallResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SuspendCallResponseBody) GoString() string {
	return s.String()
}

func (s *SuspendCallResponseBody) GetCode() *string {
	return s.Code
}

func (s *SuspendCallResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *SuspendCallResponseBody) GetMessage() *string {
	return s.Message
}

func (s *SuspendCallResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SuspendCallResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *SuspendCallResponseBody) SetCode(v string) *SuspendCallResponseBody {
	s.Code = &v
	return s
}

func (s *SuspendCallResponseBody) SetHttpStatusCode(v int32) *SuspendCallResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *SuspendCallResponseBody) SetMessage(v string) *SuspendCallResponseBody {
	s.Message = &v
	return s
}

func (s *SuspendCallResponseBody) SetRequestId(v string) *SuspendCallResponseBody {
	s.RequestId = &v
	return s
}

func (s *SuspendCallResponseBody) SetSuccess(v bool) *SuspendCallResponseBody {
	s.Success = &v
	return s
}

func (s *SuspendCallResponseBody) Validate() error {
	return dara.Validate(s)
}
