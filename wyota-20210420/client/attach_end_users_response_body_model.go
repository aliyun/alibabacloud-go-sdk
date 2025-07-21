// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAttachEndUsersResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *AttachEndUsersResponseBody
	GetCode() *string
	SetMessage(v string) *AttachEndUsersResponseBody
	GetMessage() *string
	SetRequestId(v string) *AttachEndUsersResponseBody
	GetRequestId() *string
}

type AttachEndUsersResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AttachEndUsersResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AttachEndUsersResponseBody) GoString() string {
	return s.String()
}

func (s *AttachEndUsersResponseBody) GetCode() *string {
	return s.Code
}

func (s *AttachEndUsersResponseBody) GetMessage() *string {
	return s.Message
}

func (s *AttachEndUsersResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AttachEndUsersResponseBody) SetCode(v string) *AttachEndUsersResponseBody {
	s.Code = &v
	return s
}

func (s *AttachEndUsersResponseBody) SetMessage(v string) *AttachEndUsersResponseBody {
	s.Message = &v
	return s
}

func (s *AttachEndUsersResponseBody) SetRequestId(v string) *AttachEndUsersResponseBody {
	s.RequestId = &v
	return s
}

func (s *AttachEndUsersResponseBody) Validate() error {
	return dara.Validate(s)
}
