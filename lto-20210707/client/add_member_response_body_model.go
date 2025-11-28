// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddMemberResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *AddMemberResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *AddMemberResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *AddMemberResponseBody
	GetMessage() *string
	SetRequestId(v string) *AddMemberResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *AddMemberResponseBody
	GetSuccess() *bool
}

type AddMemberResponseBody struct {
	Code           *string `json:"Code,omitempty" xml:"Code,omitempty"`
	HttpStatusCode *int32  `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s AddMemberResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AddMemberResponseBody) GoString() string {
	return s.String()
}

func (s *AddMemberResponseBody) GetCode() *string {
	return s.Code
}

func (s *AddMemberResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *AddMemberResponseBody) GetMessage() *string {
	return s.Message
}

func (s *AddMemberResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AddMemberResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *AddMemberResponseBody) SetCode(v string) *AddMemberResponseBody {
	s.Code = &v
	return s
}

func (s *AddMemberResponseBody) SetHttpStatusCode(v int32) *AddMemberResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *AddMemberResponseBody) SetMessage(v string) *AddMemberResponseBody {
	s.Message = &v
	return s
}

func (s *AddMemberResponseBody) SetRequestId(v string) *AddMemberResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddMemberResponseBody) SetSuccess(v bool) *AddMemberResponseBody {
	s.Success = &v
	return s
}

func (s *AddMemberResponseBody) Validate() error {
	return dara.Validate(s)
}
