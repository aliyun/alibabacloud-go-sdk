// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iFreezeMemberResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *FreezeMemberResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *FreezeMemberResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *FreezeMemberResponseBody
	GetMessage() *string
	SetRequestId(v string) *FreezeMemberResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *FreezeMemberResponseBody
	GetSuccess() *bool
}

type FreezeMemberResponseBody struct {
	Code           *string `json:"Code,omitempty" xml:"Code,omitempty"`
	HttpStatusCode *int32  `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s FreezeMemberResponseBody) String() string {
	return dara.Prettify(s)
}

func (s FreezeMemberResponseBody) GoString() string {
	return s.String()
}

func (s *FreezeMemberResponseBody) GetCode() *string {
	return s.Code
}

func (s *FreezeMemberResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *FreezeMemberResponseBody) GetMessage() *string {
	return s.Message
}

func (s *FreezeMemberResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *FreezeMemberResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *FreezeMemberResponseBody) SetCode(v string) *FreezeMemberResponseBody {
	s.Code = &v
	return s
}

func (s *FreezeMemberResponseBody) SetHttpStatusCode(v int32) *FreezeMemberResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *FreezeMemberResponseBody) SetMessage(v string) *FreezeMemberResponseBody {
	s.Message = &v
	return s
}

func (s *FreezeMemberResponseBody) SetRequestId(v string) *FreezeMemberResponseBody {
	s.RequestId = &v
	return s
}

func (s *FreezeMemberResponseBody) SetSuccess(v bool) *FreezeMemberResponseBody {
	s.Success = &v
	return s
}

func (s *FreezeMemberResponseBody) Validate() error {
	return dara.Validate(s)
}
