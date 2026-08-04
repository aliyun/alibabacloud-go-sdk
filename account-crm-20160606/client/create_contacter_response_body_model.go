// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateContacterResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CreateContacterResponseBody
	GetCode() *string
	SetContacterId(v string) *CreateContacterResponseBody
	GetContacterId() *string
	SetMessage(v string) *CreateContacterResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreateContacterResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreateContacterResponseBody
	GetSuccess() *bool
}

type CreateContacterResponseBody struct {
	Code        *string `json:"Code,omitempty" xml:"Code,omitempty"`
	ContacterId *string `json:"ContacterId,omitempty" xml:"ContacterId,omitempty"`
	Message     *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId   *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success     *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateContacterResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateContacterResponseBody) GoString() string {
	return s.String()
}

func (s *CreateContacterResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreateContacterResponseBody) GetContacterId() *string {
	return s.ContacterId
}

func (s *CreateContacterResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateContacterResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateContacterResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateContacterResponseBody) SetCode(v string) *CreateContacterResponseBody {
	s.Code = &v
	return s
}

func (s *CreateContacterResponseBody) SetContacterId(v string) *CreateContacterResponseBody {
	s.ContacterId = &v
	return s
}

func (s *CreateContacterResponseBody) SetMessage(v string) *CreateContacterResponseBody {
	s.Message = &v
	return s
}

func (s *CreateContacterResponseBody) SetRequestId(v string) *CreateContacterResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateContacterResponseBody) SetSuccess(v bool) *CreateContacterResponseBody {
	s.Success = &v
	return s
}

func (s *CreateContacterResponseBody) Validate() error {
	return dara.Validate(s)
}
