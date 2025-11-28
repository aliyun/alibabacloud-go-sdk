// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUnFreezeMemberResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *UnFreezeMemberResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *UnFreezeMemberResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *UnFreezeMemberResponseBody
	GetMessage() *string
	SetRequestId(v string) *UnFreezeMemberResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UnFreezeMemberResponseBody
	GetSuccess() *bool
}

type UnFreezeMemberResponseBody struct {
	Code           *string `json:"Code,omitempty" xml:"Code,omitempty"`
	HttpStatusCode *int32  `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UnFreezeMemberResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UnFreezeMemberResponseBody) GoString() string {
	return s.String()
}

func (s *UnFreezeMemberResponseBody) GetCode() *string {
	return s.Code
}

func (s *UnFreezeMemberResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *UnFreezeMemberResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UnFreezeMemberResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UnFreezeMemberResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UnFreezeMemberResponseBody) SetCode(v string) *UnFreezeMemberResponseBody {
	s.Code = &v
	return s
}

func (s *UnFreezeMemberResponseBody) SetHttpStatusCode(v int32) *UnFreezeMemberResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *UnFreezeMemberResponseBody) SetMessage(v string) *UnFreezeMemberResponseBody {
	s.Message = &v
	return s
}

func (s *UnFreezeMemberResponseBody) SetRequestId(v string) *UnFreezeMemberResponseBody {
	s.RequestId = &v
	return s
}

func (s *UnFreezeMemberResponseBody) SetSuccess(v bool) *UnFreezeMemberResponseBody {
	s.Success = &v
	return s
}

func (s *UnFreezeMemberResponseBody) Validate() error {
	return dara.Validate(s)
}
