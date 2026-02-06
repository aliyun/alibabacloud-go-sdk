// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSaveMaxAttemptsPerDayResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *SaveMaxAttemptsPerDayResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *SaveMaxAttemptsPerDayResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *SaveMaxAttemptsPerDayResponseBody
	GetMessage() *string
	SetRequestId(v string) *SaveMaxAttemptsPerDayResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *SaveMaxAttemptsPerDayResponseBody
	GetSuccess() *bool
}

type SaveMaxAttemptsPerDayResponseBody struct {
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

func (s SaveMaxAttemptsPerDayResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SaveMaxAttemptsPerDayResponseBody) GoString() string {
	return s.String()
}

func (s *SaveMaxAttemptsPerDayResponseBody) GetCode() *string {
	return s.Code
}

func (s *SaveMaxAttemptsPerDayResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *SaveMaxAttemptsPerDayResponseBody) GetMessage() *string {
	return s.Message
}

func (s *SaveMaxAttemptsPerDayResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SaveMaxAttemptsPerDayResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *SaveMaxAttemptsPerDayResponseBody) SetCode(v string) *SaveMaxAttemptsPerDayResponseBody {
	s.Code = &v
	return s
}

func (s *SaveMaxAttemptsPerDayResponseBody) SetHttpStatusCode(v int32) *SaveMaxAttemptsPerDayResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *SaveMaxAttemptsPerDayResponseBody) SetMessage(v string) *SaveMaxAttemptsPerDayResponseBody {
	s.Message = &v
	return s
}

func (s *SaveMaxAttemptsPerDayResponseBody) SetRequestId(v string) *SaveMaxAttemptsPerDayResponseBody {
	s.RequestId = &v
	return s
}

func (s *SaveMaxAttemptsPerDayResponseBody) SetSuccess(v bool) *SaveMaxAttemptsPerDayResponseBody {
	s.Success = &v
	return s
}

func (s *SaveMaxAttemptsPerDayResponseBody) Validate() error {
	return dara.Validate(s)
}
