// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSaveEffectiveDaysResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *SaveEffectiveDaysResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *SaveEffectiveDaysResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *SaveEffectiveDaysResponseBody
	GetMessage() *string
	SetRequestId(v string) *SaveEffectiveDaysResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *SaveEffectiveDaysResponseBody
	GetSuccess() *bool
}

type SaveEffectiveDaysResponseBody struct {
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

func (s SaveEffectiveDaysResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SaveEffectiveDaysResponseBody) GoString() string {
	return s.String()
}

func (s *SaveEffectiveDaysResponseBody) GetCode() *string {
	return s.Code
}

func (s *SaveEffectiveDaysResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *SaveEffectiveDaysResponseBody) GetMessage() *string {
	return s.Message
}

func (s *SaveEffectiveDaysResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SaveEffectiveDaysResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *SaveEffectiveDaysResponseBody) SetCode(v string) *SaveEffectiveDaysResponseBody {
	s.Code = &v
	return s
}

func (s *SaveEffectiveDaysResponseBody) SetHttpStatusCode(v int32) *SaveEffectiveDaysResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *SaveEffectiveDaysResponseBody) SetMessage(v string) *SaveEffectiveDaysResponseBody {
	s.Message = &v
	return s
}

func (s *SaveEffectiveDaysResponseBody) SetRequestId(v string) *SaveEffectiveDaysResponseBody {
	s.RequestId = &v
	return s
}

func (s *SaveEffectiveDaysResponseBody) SetSuccess(v bool) *SaveEffectiveDaysResponseBody {
	s.Success = &v
	return s
}

func (s *SaveEffectiveDaysResponseBody) Validate() error {
	return dara.Validate(s)
}
