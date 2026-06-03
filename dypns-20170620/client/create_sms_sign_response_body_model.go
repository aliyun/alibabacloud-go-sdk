// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateSmsSignResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CreateSmsSignResponseBody
	GetCode() *string
	SetData(v string) *CreateSmsSignResponseBody
	GetData() *string
	SetMessage(v string) *CreateSmsSignResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreateSmsSignResponseBody
	GetRequestId() *string
}

type CreateSmsSignResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateSmsSignResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateSmsSignResponseBody) GoString() string {
	return s.String()
}

func (s *CreateSmsSignResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreateSmsSignResponseBody) GetData() *string {
	return s.Data
}

func (s *CreateSmsSignResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateSmsSignResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateSmsSignResponseBody) SetCode(v string) *CreateSmsSignResponseBody {
	s.Code = &v
	return s
}

func (s *CreateSmsSignResponseBody) SetData(v string) *CreateSmsSignResponseBody {
	s.Data = &v
	return s
}

func (s *CreateSmsSignResponseBody) SetMessage(v string) *CreateSmsSignResponseBody {
	s.Message = &v
	return s
}

func (s *CreateSmsSignResponseBody) SetRequestId(v string) *CreateSmsSignResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateSmsSignResponseBody) Validate() error {
	return dara.Validate(s)
}
