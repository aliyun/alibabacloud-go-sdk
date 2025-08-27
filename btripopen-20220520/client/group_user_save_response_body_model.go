// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGroupUserSaveResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GroupUserSaveResponseBody
	GetCode() *string
	SetMessage(v string) *GroupUserSaveResponseBody
	GetMessage() *string
	SetRequestId(v string) *GroupUserSaveResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GroupUserSaveResponseBody
	GetSuccess() *bool
	SetTraceId(v string) *GroupUserSaveResponseBody
	GetTraceId() *string
}

type GroupUserSaveResponseBody struct {
	// example:
	//
	// 0
	Code    *string `json:"code,omitempty" xml:"code,omitempty"`
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// example:
	//
	// C61ECFF6-606B-5F66-B81D-D77369043A5F
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
	// example:
	//
	// 21041ce316577904808056433edbb2
	TraceId *string `json:"traceId,omitempty" xml:"traceId,omitempty"`
}

func (s GroupUserSaveResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GroupUserSaveResponseBody) GoString() string {
	return s.String()
}

func (s *GroupUserSaveResponseBody) GetCode() *string {
	return s.Code
}

func (s *GroupUserSaveResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GroupUserSaveResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GroupUserSaveResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GroupUserSaveResponseBody) GetTraceId() *string {
	return s.TraceId
}

func (s *GroupUserSaveResponseBody) SetCode(v string) *GroupUserSaveResponseBody {
	s.Code = &v
	return s
}

func (s *GroupUserSaveResponseBody) SetMessage(v string) *GroupUserSaveResponseBody {
	s.Message = &v
	return s
}

func (s *GroupUserSaveResponseBody) SetRequestId(v string) *GroupUserSaveResponseBody {
	s.RequestId = &v
	return s
}

func (s *GroupUserSaveResponseBody) SetSuccess(v bool) *GroupUserSaveResponseBody {
	s.Success = &v
	return s
}

func (s *GroupUserSaveResponseBody) SetTraceId(v string) *GroupUserSaveResponseBody {
	s.TraceId = &v
	return s
}

func (s *GroupUserSaveResponseBody) Validate() error {
	return dara.Validate(s)
}
