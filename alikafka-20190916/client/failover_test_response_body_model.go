// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iFailoverTestResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *FailoverTestResponseBody
	GetCode() *int32
	SetMessage(v string) *FailoverTestResponseBody
	GetMessage() *string
	SetRequestId(v string) *FailoverTestResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *FailoverTestResponseBody
	GetSuccess() *bool
}

type FailoverTestResponseBody struct {
	Code      *int32  `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s FailoverTestResponseBody) String() string {
	return dara.Prettify(s)
}

func (s FailoverTestResponseBody) GoString() string {
	return s.String()
}

func (s *FailoverTestResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *FailoverTestResponseBody) GetMessage() *string {
	return s.Message
}

func (s *FailoverTestResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *FailoverTestResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *FailoverTestResponseBody) SetCode(v int32) *FailoverTestResponseBody {
	s.Code = &v
	return s
}

func (s *FailoverTestResponseBody) SetMessage(v string) *FailoverTestResponseBody {
	s.Message = &v
	return s
}

func (s *FailoverTestResponseBody) SetRequestId(v string) *FailoverTestResponseBody {
	s.RequestId = &v
	return s
}

func (s *FailoverTestResponseBody) SetSuccess(v bool) *FailoverTestResponseBody {
	s.Success = &v
	return s
}

func (s *FailoverTestResponseBody) Validate() error {
	return dara.Validate(s)
}
