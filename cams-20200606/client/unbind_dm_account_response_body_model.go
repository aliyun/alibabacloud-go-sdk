// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUnbindDmAccountResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *UnbindDmAccountResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v string) *UnbindDmAccountResponseBody
	GetCode() *string
	SetData(v string) *UnbindDmAccountResponseBody
	GetData() *string
	SetMessage(v string) *UnbindDmAccountResponseBody
	GetMessage() *string
	SetRequestId(v string) *UnbindDmAccountResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UnbindDmAccountResponseBody
	GetSuccess() *bool
}

type UnbindDmAccountResponseBody struct {
	// example:
	//
	// None
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// NONE
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// xx-xx***
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UnbindDmAccountResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UnbindDmAccountResponseBody) GoString() string {
	return s.String()
}

func (s *UnbindDmAccountResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *UnbindDmAccountResponseBody) GetCode() *string {
	return s.Code
}

func (s *UnbindDmAccountResponseBody) GetData() *string {
	return s.Data
}

func (s *UnbindDmAccountResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UnbindDmAccountResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UnbindDmAccountResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UnbindDmAccountResponseBody) SetAccessDeniedDetail(v string) *UnbindDmAccountResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *UnbindDmAccountResponseBody) SetCode(v string) *UnbindDmAccountResponseBody {
	s.Code = &v
	return s
}

func (s *UnbindDmAccountResponseBody) SetData(v string) *UnbindDmAccountResponseBody {
	s.Data = &v
	return s
}

func (s *UnbindDmAccountResponseBody) SetMessage(v string) *UnbindDmAccountResponseBody {
	s.Message = &v
	return s
}

func (s *UnbindDmAccountResponseBody) SetRequestId(v string) *UnbindDmAccountResponseBody {
	s.RequestId = &v
	return s
}

func (s *UnbindDmAccountResponseBody) SetSuccess(v bool) *UnbindDmAccountResponseBody {
	s.Success = &v
	return s
}

func (s *UnbindDmAccountResponseBody) Validate() error {
	return dara.Validate(s)
}
