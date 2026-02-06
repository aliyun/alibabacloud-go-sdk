// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateGlobalQuestionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CreateGlobalQuestionResponseBody
	GetCode() *string
	SetGlobalQuestionId(v string) *CreateGlobalQuestionResponseBody
	GetGlobalQuestionId() *string
	SetHttpStatusCode(v int32) *CreateGlobalQuestionResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *CreateGlobalQuestionResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreateGlobalQuestionResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreateGlobalQuestionResponseBody
	GetSuccess() *bool
}

type CreateGlobalQuestionResponseBody struct {
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// a546b616-724b-437f-bdb3-629a30c98567
	GlobalQuestionId *string `json:"GlobalQuestionId,omitempty" xml:"GlobalQuestionId,omitempty"`
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

func (s CreateGlobalQuestionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateGlobalQuestionResponseBody) GoString() string {
	return s.String()
}

func (s *CreateGlobalQuestionResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreateGlobalQuestionResponseBody) GetGlobalQuestionId() *string {
	return s.GlobalQuestionId
}

func (s *CreateGlobalQuestionResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *CreateGlobalQuestionResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateGlobalQuestionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateGlobalQuestionResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateGlobalQuestionResponseBody) SetCode(v string) *CreateGlobalQuestionResponseBody {
	s.Code = &v
	return s
}

func (s *CreateGlobalQuestionResponseBody) SetGlobalQuestionId(v string) *CreateGlobalQuestionResponseBody {
	s.GlobalQuestionId = &v
	return s
}

func (s *CreateGlobalQuestionResponseBody) SetHttpStatusCode(v int32) *CreateGlobalQuestionResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *CreateGlobalQuestionResponseBody) SetMessage(v string) *CreateGlobalQuestionResponseBody {
	s.Message = &v
	return s
}

func (s *CreateGlobalQuestionResponseBody) SetRequestId(v string) *CreateGlobalQuestionResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateGlobalQuestionResponseBody) SetSuccess(v bool) *CreateGlobalQuestionResponseBody {
	s.Success = &v
	return s
}

func (s *CreateGlobalQuestionResponseBody) Validate() error {
	return dara.Validate(s)
}
