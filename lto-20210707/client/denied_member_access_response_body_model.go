// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeniedMemberAccessResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DeniedMemberAccessResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *DeniedMemberAccessResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *DeniedMemberAccessResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeniedMemberAccessResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeniedMemberAccessResponseBody
	GetSuccess() *bool
}

type DeniedMemberAccessResponseBody struct {
	Code           *string `json:"Code,omitempty" xml:"Code,omitempty"`
	HttpStatusCode *int32  `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeniedMemberAccessResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeniedMemberAccessResponseBody) GoString() string {
	return s.String()
}

func (s *DeniedMemberAccessResponseBody) GetCode() *string {
	return s.Code
}

func (s *DeniedMemberAccessResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *DeniedMemberAccessResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeniedMemberAccessResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeniedMemberAccessResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeniedMemberAccessResponseBody) SetCode(v string) *DeniedMemberAccessResponseBody {
	s.Code = &v
	return s
}

func (s *DeniedMemberAccessResponseBody) SetHttpStatusCode(v int32) *DeniedMemberAccessResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DeniedMemberAccessResponseBody) SetMessage(v string) *DeniedMemberAccessResponseBody {
	s.Message = &v
	return s
}

func (s *DeniedMemberAccessResponseBody) SetRequestId(v string) *DeniedMemberAccessResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeniedMemberAccessResponseBody) SetSuccess(v bool) *DeniedMemberAccessResponseBody {
	s.Success = &v
	return s
}

func (s *DeniedMemberAccessResponseBody) Validate() error {
	return dara.Validate(s)
}
