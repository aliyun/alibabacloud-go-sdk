// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAgreeMemberAccessResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *AgreeMemberAccessResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *AgreeMemberAccessResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *AgreeMemberAccessResponseBody
	GetMessage() *string
	SetRequestId(v string) *AgreeMemberAccessResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *AgreeMemberAccessResponseBody
	GetSuccess() *bool
}

type AgreeMemberAccessResponseBody struct {
	Code           *string `json:"Code,omitempty" xml:"Code,omitempty"`
	HttpStatusCode *int32  `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s AgreeMemberAccessResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AgreeMemberAccessResponseBody) GoString() string {
	return s.String()
}

func (s *AgreeMemberAccessResponseBody) GetCode() *string {
	return s.Code
}

func (s *AgreeMemberAccessResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *AgreeMemberAccessResponseBody) GetMessage() *string {
	return s.Message
}

func (s *AgreeMemberAccessResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AgreeMemberAccessResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *AgreeMemberAccessResponseBody) SetCode(v string) *AgreeMemberAccessResponseBody {
	s.Code = &v
	return s
}

func (s *AgreeMemberAccessResponseBody) SetHttpStatusCode(v int32) *AgreeMemberAccessResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *AgreeMemberAccessResponseBody) SetMessage(v string) *AgreeMemberAccessResponseBody {
	s.Message = &v
	return s
}

func (s *AgreeMemberAccessResponseBody) SetRequestId(v string) *AgreeMemberAccessResponseBody {
	s.RequestId = &v
	return s
}

func (s *AgreeMemberAccessResponseBody) SetSuccess(v bool) *AgreeMemberAccessResponseBody {
	s.Success = &v
	return s
}

func (s *AgreeMemberAccessResponseBody) Validate() error {
	return dara.Validate(s)
}
