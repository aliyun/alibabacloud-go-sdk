// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateMemberResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *UpdateMemberResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *UpdateMemberResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *UpdateMemberResponseBody
	GetMessage() *string
	SetRequestId(v string) *UpdateMemberResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateMemberResponseBody
	GetSuccess() *bool
}

type UpdateMemberResponseBody struct {
	Code           *string `json:"Code,omitempty" xml:"Code,omitempty"`
	HttpStatusCode *int32  `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateMemberResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateMemberResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateMemberResponseBody) GetCode() *string {
	return s.Code
}

func (s *UpdateMemberResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *UpdateMemberResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateMemberResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateMemberResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateMemberResponseBody) SetCode(v string) *UpdateMemberResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateMemberResponseBody) SetHttpStatusCode(v int32) *UpdateMemberResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *UpdateMemberResponseBody) SetMessage(v string) *UpdateMemberResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateMemberResponseBody) SetRequestId(v string) *UpdateMemberResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateMemberResponseBody) SetSuccess(v bool) *UpdateMemberResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateMemberResponseBody) Validate() error {
	return dara.Validate(s)
}
