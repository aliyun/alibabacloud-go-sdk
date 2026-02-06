// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMaxAttemptsPerDayResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetMaxAttemptsPerDayResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *GetMaxAttemptsPerDayResponseBody
	GetHttpStatusCode() *int32
	SetMaxAttemptsPerDay(v int32) *GetMaxAttemptsPerDayResponseBody
	GetMaxAttemptsPerDay() *int32
	SetMessage(v string) *GetMaxAttemptsPerDayResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetMaxAttemptsPerDayResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetMaxAttemptsPerDayResponseBody
	GetSuccess() *bool
}

type GetMaxAttemptsPerDayResponseBody struct {
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
	// 2
	MaxAttemptsPerDay *int32 `json:"MaxAttemptsPerDay,omitempty" xml:"MaxAttemptsPerDay,omitempty"`
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

func (s GetMaxAttemptsPerDayResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetMaxAttemptsPerDayResponseBody) GoString() string {
	return s.String()
}

func (s *GetMaxAttemptsPerDayResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetMaxAttemptsPerDayResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GetMaxAttemptsPerDayResponseBody) GetMaxAttemptsPerDay() *int32 {
	return s.MaxAttemptsPerDay
}

func (s *GetMaxAttemptsPerDayResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetMaxAttemptsPerDayResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetMaxAttemptsPerDayResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetMaxAttemptsPerDayResponseBody) SetCode(v string) *GetMaxAttemptsPerDayResponseBody {
	s.Code = &v
	return s
}

func (s *GetMaxAttemptsPerDayResponseBody) SetHttpStatusCode(v int32) *GetMaxAttemptsPerDayResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetMaxAttemptsPerDayResponseBody) SetMaxAttemptsPerDay(v int32) *GetMaxAttemptsPerDayResponseBody {
	s.MaxAttemptsPerDay = &v
	return s
}

func (s *GetMaxAttemptsPerDayResponseBody) SetMessage(v string) *GetMaxAttemptsPerDayResponseBody {
	s.Message = &v
	return s
}

func (s *GetMaxAttemptsPerDayResponseBody) SetRequestId(v string) *GetMaxAttemptsPerDayResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetMaxAttemptsPerDayResponseBody) SetSuccess(v bool) *GetMaxAttemptsPerDayResponseBody {
	s.Success = &v
	return s
}

func (s *GetMaxAttemptsPerDayResponseBody) Validate() error {
	return dara.Validate(s)
}
