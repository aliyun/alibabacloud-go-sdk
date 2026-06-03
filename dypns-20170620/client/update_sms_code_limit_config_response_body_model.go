// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateSmsCodeLimitConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *UpdateSmsCodeLimitConfigResponseBody
	GetCode() *string
	SetData(v string) *UpdateSmsCodeLimitConfigResponseBody
	GetData() *string
	SetMessage(v string) *UpdateSmsCodeLimitConfigResponseBody
	GetMessage() *string
	SetRequestId(v string) *UpdateSmsCodeLimitConfigResponseBody
	GetRequestId() *string
}

type UpdateSmsCodeLimitConfigResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateSmsCodeLimitConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateSmsCodeLimitConfigResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateSmsCodeLimitConfigResponseBody) GetCode() *string {
	return s.Code
}

func (s *UpdateSmsCodeLimitConfigResponseBody) GetData() *string {
	return s.Data
}

func (s *UpdateSmsCodeLimitConfigResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateSmsCodeLimitConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateSmsCodeLimitConfigResponseBody) SetCode(v string) *UpdateSmsCodeLimitConfigResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateSmsCodeLimitConfigResponseBody) SetData(v string) *UpdateSmsCodeLimitConfigResponseBody {
	s.Data = &v
	return s
}

func (s *UpdateSmsCodeLimitConfigResponseBody) SetMessage(v string) *UpdateSmsCodeLimitConfigResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateSmsCodeLimitConfigResponseBody) SetRequestId(v string) *UpdateSmsCodeLimitConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateSmsCodeLimitConfigResponseBody) Validate() error {
	return dara.Validate(s)
}
