// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReleaseSasInstanceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ReleaseSasInstanceResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *ReleaseSasInstanceResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *ReleaseSasInstanceResponseBody
	GetMessage() *string
	SetRequestId(v string) *ReleaseSasInstanceResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ReleaseSasInstanceResponseBody
	GetSuccess() *bool
}

type ReleaseSasInstanceResponseBody struct {
	Code           *string `json:"Code,omitempty" xml:"Code,omitempty"`
	HttpStatusCode *int32  `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ReleaseSasInstanceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ReleaseSasInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *ReleaseSasInstanceResponseBody) GetCode() *string {
	return s.Code
}

func (s *ReleaseSasInstanceResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ReleaseSasInstanceResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ReleaseSasInstanceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ReleaseSasInstanceResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ReleaseSasInstanceResponseBody) SetCode(v string) *ReleaseSasInstanceResponseBody {
	s.Code = &v
	return s
}

func (s *ReleaseSasInstanceResponseBody) SetHttpStatusCode(v int32) *ReleaseSasInstanceResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ReleaseSasInstanceResponseBody) SetMessage(v string) *ReleaseSasInstanceResponseBody {
	s.Message = &v
	return s
}

func (s *ReleaseSasInstanceResponseBody) SetRequestId(v string) *ReleaseSasInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *ReleaseSasInstanceResponseBody) SetSuccess(v bool) *ReleaseSasInstanceResponseBody {
	s.Success = &v
	return s
}

func (s *ReleaseSasInstanceResponseBody) Validate() error {
	return dara.Validate(s)
}
