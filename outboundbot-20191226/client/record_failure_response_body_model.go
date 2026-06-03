// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRecordFailureResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *RecordFailureResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *RecordFailureResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *RecordFailureResponseBody
	GetMessage() *string
	SetRequestId(v string) *RecordFailureResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *RecordFailureResponseBody
	GetSuccess() *bool
}

type RecordFailureResponseBody struct {
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

func (s RecordFailureResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RecordFailureResponseBody) GoString() string {
	return s.String()
}

func (s *RecordFailureResponseBody) GetCode() *string {
	return s.Code
}

func (s *RecordFailureResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *RecordFailureResponseBody) GetMessage() *string {
	return s.Message
}

func (s *RecordFailureResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RecordFailureResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *RecordFailureResponseBody) SetCode(v string) *RecordFailureResponseBody {
	s.Code = &v
	return s
}

func (s *RecordFailureResponseBody) SetHttpStatusCode(v int32) *RecordFailureResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *RecordFailureResponseBody) SetMessage(v string) *RecordFailureResponseBody {
	s.Message = &v
	return s
}

func (s *RecordFailureResponseBody) SetRequestId(v string) *RecordFailureResponseBody {
	s.RequestId = &v
	return s
}

func (s *RecordFailureResponseBody) SetSuccess(v bool) *RecordFailureResponseBody {
	s.Success = &v
	return s
}

func (s *RecordFailureResponseBody) Validate() error {
	return dara.Validate(s)
}
