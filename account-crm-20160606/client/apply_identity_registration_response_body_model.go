// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iApplyIdentityRegistrationResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ApplyIdentityRegistrationResponseBody
	GetCode() *string
	SetData(v int64) *ApplyIdentityRegistrationResponseBody
	GetData() *int64
	SetMessage(v string) *ApplyIdentityRegistrationResponseBody
	GetMessage() *string
	SetRequestId(v string) *ApplyIdentityRegistrationResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ApplyIdentityRegistrationResponseBody
	GetSuccess() *bool
}

type ApplyIdentityRegistrationResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *int64  `json:"Data,omitempty" xml:"Data,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ApplyIdentityRegistrationResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ApplyIdentityRegistrationResponseBody) GoString() string {
	return s.String()
}

func (s *ApplyIdentityRegistrationResponseBody) GetCode() *string {
	return s.Code
}

func (s *ApplyIdentityRegistrationResponseBody) GetData() *int64 {
	return s.Data
}

func (s *ApplyIdentityRegistrationResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ApplyIdentityRegistrationResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ApplyIdentityRegistrationResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ApplyIdentityRegistrationResponseBody) SetCode(v string) *ApplyIdentityRegistrationResponseBody {
	s.Code = &v
	return s
}

func (s *ApplyIdentityRegistrationResponseBody) SetData(v int64) *ApplyIdentityRegistrationResponseBody {
	s.Data = &v
	return s
}

func (s *ApplyIdentityRegistrationResponseBody) SetMessage(v string) *ApplyIdentityRegistrationResponseBody {
	s.Message = &v
	return s
}

func (s *ApplyIdentityRegistrationResponseBody) SetRequestId(v string) *ApplyIdentityRegistrationResponseBody {
	s.RequestId = &v
	return s
}

func (s *ApplyIdentityRegistrationResponseBody) SetSuccess(v bool) *ApplyIdentityRegistrationResponseBody {
	s.Success = &v
	return s
}

func (s *ApplyIdentityRegistrationResponseBody) Validate() error {
	return dara.Validate(s)
}
