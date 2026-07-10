// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGroupDepartSaveResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GroupDepartSaveResponseBody
	GetCode() *string
	SetMessage(v string) *GroupDepartSaveResponseBody
	GetMessage() *string
	SetRequestId(v string) *GroupDepartSaveResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GroupDepartSaveResponseBody
	GetSuccess() *bool
	SetTraceId(v string) *GroupDepartSaveResponseBody
	GetTraceId() *string
}

type GroupDepartSaveResponseBody struct {
	Code      *string `json:"code,omitempty" xml:"code,omitempty"`
	Message   *string `json:"message,omitempty" xml:"message,omitempty"`
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Success   *bool   `json:"success,omitempty" xml:"success,omitempty"`
	TraceId   *string `json:"traceId,omitempty" xml:"traceId,omitempty"`
}

func (s GroupDepartSaveResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GroupDepartSaveResponseBody) GoString() string {
	return s.String()
}

func (s *GroupDepartSaveResponseBody) GetCode() *string {
	return s.Code
}

func (s *GroupDepartSaveResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GroupDepartSaveResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GroupDepartSaveResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GroupDepartSaveResponseBody) GetTraceId() *string {
	return s.TraceId
}

func (s *GroupDepartSaveResponseBody) SetCode(v string) *GroupDepartSaveResponseBody {
	s.Code = &v
	return s
}

func (s *GroupDepartSaveResponseBody) SetMessage(v string) *GroupDepartSaveResponseBody {
	s.Message = &v
	return s
}

func (s *GroupDepartSaveResponseBody) SetRequestId(v string) *GroupDepartSaveResponseBody {
	s.RequestId = &v
	return s
}

func (s *GroupDepartSaveResponseBody) SetSuccess(v bool) *GroupDepartSaveResponseBody {
	s.Success = &v
	return s
}

func (s *GroupDepartSaveResponseBody) SetTraceId(v string) *GroupDepartSaveResponseBody {
	s.TraceId = &v
	return s
}

func (s *GroupDepartSaveResponseBody) Validate() error {
	return dara.Validate(s)
}
