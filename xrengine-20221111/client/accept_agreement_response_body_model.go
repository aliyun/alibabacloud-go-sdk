// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAcceptAgreementResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *AcceptAgreementResponseBody
	GetCode() *string
	SetErrorName(v string) *AcceptAgreementResponseBody
	GetErrorName() *string
	SetHttpCode(v int32) *AcceptAgreementResponseBody
	GetHttpCode() *int32
	SetMessage(v string) *AcceptAgreementResponseBody
	GetMessage() *string
	SetRequestId(v string) *AcceptAgreementResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *AcceptAgreementResponseBody
	GetSuccess() *bool
}

type AcceptAgreementResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	ErrorName *string `json:"ErrorName,omitempty" xml:"ErrorName,omitempty"`
	HttpCode  *int32  `json:"HttpCode,omitempty" xml:"HttpCode,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s AcceptAgreementResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AcceptAgreementResponseBody) GoString() string {
	return s.String()
}

func (s *AcceptAgreementResponseBody) GetCode() *string {
	return s.Code
}

func (s *AcceptAgreementResponseBody) GetErrorName() *string {
	return s.ErrorName
}

func (s *AcceptAgreementResponseBody) GetHttpCode() *int32 {
	return s.HttpCode
}

func (s *AcceptAgreementResponseBody) GetMessage() *string {
	return s.Message
}

func (s *AcceptAgreementResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AcceptAgreementResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *AcceptAgreementResponseBody) SetCode(v string) *AcceptAgreementResponseBody {
	s.Code = &v
	return s
}

func (s *AcceptAgreementResponseBody) SetErrorName(v string) *AcceptAgreementResponseBody {
	s.ErrorName = &v
	return s
}

func (s *AcceptAgreementResponseBody) SetHttpCode(v int32) *AcceptAgreementResponseBody {
	s.HttpCode = &v
	return s
}

func (s *AcceptAgreementResponseBody) SetMessage(v string) *AcceptAgreementResponseBody {
	s.Message = &v
	return s
}

func (s *AcceptAgreementResponseBody) SetRequestId(v string) *AcceptAgreementResponseBody {
	s.RequestId = &v
	return s
}

func (s *AcceptAgreementResponseBody) SetSuccess(v bool) *AcceptAgreementResponseBody {
	s.Success = &v
	return s
}

func (s *AcceptAgreementResponseBody) Validate() error {
	return dara.Validate(s)
}
