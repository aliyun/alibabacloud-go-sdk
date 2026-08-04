// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAccountOneKeyDeleteResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *AccountOneKeyDeleteResponseBody
	GetCode() *string
	SetMessage(v string) *AccountOneKeyDeleteResponseBody
	GetMessage() *string
	SetRequestId(v string) *AccountOneKeyDeleteResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *AccountOneKeyDeleteResponseBody
	GetSuccess() *bool
}

type AccountOneKeyDeleteResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s AccountOneKeyDeleteResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AccountOneKeyDeleteResponseBody) GoString() string {
	return s.String()
}

func (s *AccountOneKeyDeleteResponseBody) GetCode() *string {
	return s.Code
}

func (s *AccountOneKeyDeleteResponseBody) GetMessage() *string {
	return s.Message
}

func (s *AccountOneKeyDeleteResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AccountOneKeyDeleteResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *AccountOneKeyDeleteResponseBody) SetCode(v string) *AccountOneKeyDeleteResponseBody {
	s.Code = &v
	return s
}

func (s *AccountOneKeyDeleteResponseBody) SetMessage(v string) *AccountOneKeyDeleteResponseBody {
	s.Message = &v
	return s
}

func (s *AccountOneKeyDeleteResponseBody) SetRequestId(v string) *AccountOneKeyDeleteResponseBody {
	s.RequestId = &v
	return s
}

func (s *AccountOneKeyDeleteResponseBody) SetSuccess(v bool) *AccountOneKeyDeleteResponseBody {
	s.Success = &v
	return s
}

func (s *AccountOneKeyDeleteResponseBody) Validate() error {
	return dara.Validate(s)
}
