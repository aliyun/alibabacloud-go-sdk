// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReleaseAgAccountResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ReleaseAgAccountResponseBody
	GetCode() *string
	SetMessage(v string) *ReleaseAgAccountResponseBody
	GetMessage() *string
	SetRequestId(v string) *ReleaseAgAccountResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ReleaseAgAccountResponseBody
	GetSuccess() *bool
}

type ReleaseAgAccountResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ReleaseAgAccountResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ReleaseAgAccountResponseBody) GoString() string {
	return s.String()
}

func (s *ReleaseAgAccountResponseBody) GetCode() *string {
	return s.Code
}

func (s *ReleaseAgAccountResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ReleaseAgAccountResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ReleaseAgAccountResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ReleaseAgAccountResponseBody) SetCode(v string) *ReleaseAgAccountResponseBody {
	s.Code = &v
	return s
}

func (s *ReleaseAgAccountResponseBody) SetMessage(v string) *ReleaseAgAccountResponseBody {
	s.Message = &v
	return s
}

func (s *ReleaseAgAccountResponseBody) SetRequestId(v string) *ReleaseAgAccountResponseBody {
	s.RequestId = &v
	return s
}

func (s *ReleaseAgAccountResponseBody) SetSuccess(v bool) *ReleaseAgAccountResponseBody {
	s.Success = &v
	return s
}

func (s *ReleaseAgAccountResponseBody) Validate() error {
	return dara.Validate(s)
}
