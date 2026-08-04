// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateRealNameCertificationResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CreateRealNameCertificationResponseBody
	GetCode() *string
	SetMessage(v string) *CreateRealNameCertificationResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreateRealNameCertificationResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreateRealNameCertificationResponseBody
	GetSuccess() *bool
}

type CreateRealNameCertificationResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateRealNameCertificationResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateRealNameCertificationResponseBody) GoString() string {
	return s.String()
}

func (s *CreateRealNameCertificationResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreateRealNameCertificationResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateRealNameCertificationResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateRealNameCertificationResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateRealNameCertificationResponseBody) SetCode(v string) *CreateRealNameCertificationResponseBody {
	s.Code = &v
	return s
}

func (s *CreateRealNameCertificationResponseBody) SetMessage(v string) *CreateRealNameCertificationResponseBody {
	s.Message = &v
	return s
}

func (s *CreateRealNameCertificationResponseBody) SetRequestId(v string) *CreateRealNameCertificationResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateRealNameCertificationResponseBody) SetSuccess(v bool) *CreateRealNameCertificationResponseBody {
	s.Success = &v
	return s
}

func (s *CreateRealNameCertificationResponseBody) Validate() error {
	return dara.Validate(s)
}
