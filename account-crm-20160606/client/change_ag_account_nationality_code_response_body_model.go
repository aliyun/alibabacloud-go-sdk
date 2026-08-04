// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iChangeAgAccountNationalityCodeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ChangeAgAccountNationalityCodeResponseBody
	GetCode() *string
	SetMessage(v string) *ChangeAgAccountNationalityCodeResponseBody
	GetMessage() *string
	SetRequestId(v string) *ChangeAgAccountNationalityCodeResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ChangeAgAccountNationalityCodeResponseBody
	GetSuccess() *bool
}

type ChangeAgAccountNationalityCodeResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ChangeAgAccountNationalityCodeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ChangeAgAccountNationalityCodeResponseBody) GoString() string {
	return s.String()
}

func (s *ChangeAgAccountNationalityCodeResponseBody) GetCode() *string {
	return s.Code
}

func (s *ChangeAgAccountNationalityCodeResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ChangeAgAccountNationalityCodeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ChangeAgAccountNationalityCodeResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ChangeAgAccountNationalityCodeResponseBody) SetCode(v string) *ChangeAgAccountNationalityCodeResponseBody {
	s.Code = &v
	return s
}

func (s *ChangeAgAccountNationalityCodeResponseBody) SetMessage(v string) *ChangeAgAccountNationalityCodeResponseBody {
	s.Message = &v
	return s
}

func (s *ChangeAgAccountNationalityCodeResponseBody) SetRequestId(v string) *ChangeAgAccountNationalityCodeResponseBody {
	s.RequestId = &v
	return s
}

func (s *ChangeAgAccountNationalityCodeResponseBody) SetSuccess(v bool) *ChangeAgAccountNationalityCodeResponseBody {
	s.Success = &v
	return s
}

func (s *ChangeAgAccountNationalityCodeResponseBody) Validate() error {
	return dara.Validate(s)
}
