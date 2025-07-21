// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDetachEndUsersResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DetachEndUsersResponseBody
	GetCode() *string
	SetMessage(v string) *DetachEndUsersResponseBody
	GetMessage() *string
	SetRequestId(v string) *DetachEndUsersResponseBody
	GetRequestId() *string
}

type DetachEndUsersResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DetachEndUsersResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DetachEndUsersResponseBody) GoString() string {
	return s.String()
}

func (s *DetachEndUsersResponseBody) GetCode() *string {
	return s.Code
}

func (s *DetachEndUsersResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DetachEndUsersResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DetachEndUsersResponseBody) SetCode(v string) *DetachEndUsersResponseBody {
	s.Code = &v
	return s
}

func (s *DetachEndUsersResponseBody) SetMessage(v string) *DetachEndUsersResponseBody {
	s.Message = &v
	return s
}

func (s *DetachEndUsersResponseBody) SetRequestId(v string) *DetachEndUsersResponseBody {
	s.RequestId = &v
	return s
}

func (s *DetachEndUsersResponseBody) Validate() error {
	return dara.Validate(s)
}
