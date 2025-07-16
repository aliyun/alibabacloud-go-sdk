// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAccountResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMessage(v string) *CreateAccountResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreateAccountResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreateAccountResponseBody
	GetSuccess() *bool
}

type CreateAccountResponseBody struct {
	// example:
	//
	// ****
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 9B2F3840-5C98-475C-B269-2D5C3A31797C
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateAccountResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateAccountResponseBody) GoString() string {
	return s.String()
}

func (s *CreateAccountResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateAccountResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateAccountResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateAccountResponseBody) SetMessage(v string) *CreateAccountResponseBody {
	s.Message = &v
	return s
}

func (s *CreateAccountResponseBody) SetRequestId(v string) *CreateAccountResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateAccountResponseBody) SetSuccess(v bool) *CreateAccountResponseBody {
	s.Success = &v
	return s
}

func (s *CreateAccountResponseBody) Validate() error {
	return dara.Validate(s)
}
