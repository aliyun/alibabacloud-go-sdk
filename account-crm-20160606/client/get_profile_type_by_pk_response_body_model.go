// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetProfileTypeByPkResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccountType(v string) *GetProfileTypeByPkResponseBody
	GetAccountType() *string
	SetCode(v string) *GetProfileTypeByPkResponseBody
	GetCode() *string
	SetMessage(v string) *GetProfileTypeByPkResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetProfileTypeByPkResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetProfileTypeByPkResponseBody
	GetSuccess() *bool
}

type GetProfileTypeByPkResponseBody struct {
	AccountType *string `json:"AccountType,omitempty" xml:"AccountType,omitempty"`
	Code        *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message     *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId   *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success     *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetProfileTypeByPkResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetProfileTypeByPkResponseBody) GoString() string {
	return s.String()
}

func (s *GetProfileTypeByPkResponseBody) GetAccountType() *string {
	return s.AccountType
}

func (s *GetProfileTypeByPkResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetProfileTypeByPkResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetProfileTypeByPkResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetProfileTypeByPkResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetProfileTypeByPkResponseBody) SetAccountType(v string) *GetProfileTypeByPkResponseBody {
	s.AccountType = &v
	return s
}

func (s *GetProfileTypeByPkResponseBody) SetCode(v string) *GetProfileTypeByPkResponseBody {
	s.Code = &v
	return s
}

func (s *GetProfileTypeByPkResponseBody) SetMessage(v string) *GetProfileTypeByPkResponseBody {
	s.Message = &v
	return s
}

func (s *GetProfileTypeByPkResponseBody) SetRequestId(v string) *GetProfileTypeByPkResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetProfileTypeByPkResponseBody) SetSuccess(v bool) *GetProfileTypeByPkResponseBody {
	s.Success = &v
	return s
}

func (s *GetProfileTypeByPkResponseBody) Validate() error {
	return dara.Validate(s)
}
