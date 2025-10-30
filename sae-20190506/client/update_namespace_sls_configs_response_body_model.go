// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateNamespaceSlsConfigsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *UpdateNamespaceSlsConfigsResponseBody
	GetCode() *string
	SetErrorCode(v string) *UpdateNamespaceSlsConfigsResponseBody
	GetErrorCode() *string
	SetMessage(v string) *UpdateNamespaceSlsConfigsResponseBody
	GetMessage() *string
	SetSuccess(v bool) *UpdateNamespaceSlsConfigsResponseBody
	GetSuccess() *bool
	SetTraceId(v string) *UpdateNamespaceSlsConfigsResponseBody
	GetTraceId() *string
	SetRequestId(v string) *UpdateNamespaceSlsConfigsResponseBody
	GetRequestId() *string
}

type UpdateNamespaceSlsConfigsResponseBody struct {
	// example:
	//
	// 200
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// example:
	//
	// 0a98a02315955564772843261e****
	TraceId *string `json:"TraceId,omitempty" xml:"TraceId,omitempty"`
	// Id of the request
	//
	// example:
	//
	// AD0286B2-F4A4-5D43-9329-97DEF1019F06
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s UpdateNamespaceSlsConfigsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateNamespaceSlsConfigsResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateNamespaceSlsConfigsResponseBody) GetCode() *string {
	return s.Code
}

func (s *UpdateNamespaceSlsConfigsResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *UpdateNamespaceSlsConfigsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateNamespaceSlsConfigsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateNamespaceSlsConfigsResponseBody) GetTraceId() *string {
	return s.TraceId
}

func (s *UpdateNamespaceSlsConfigsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateNamespaceSlsConfigsResponseBody) SetCode(v string) *UpdateNamespaceSlsConfigsResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateNamespaceSlsConfigsResponseBody) SetErrorCode(v string) *UpdateNamespaceSlsConfigsResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *UpdateNamespaceSlsConfigsResponseBody) SetMessage(v string) *UpdateNamespaceSlsConfigsResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateNamespaceSlsConfigsResponseBody) SetSuccess(v bool) *UpdateNamespaceSlsConfigsResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateNamespaceSlsConfigsResponseBody) SetTraceId(v string) *UpdateNamespaceSlsConfigsResponseBody {
	s.TraceId = &v
	return s
}

func (s *UpdateNamespaceSlsConfigsResponseBody) SetRequestId(v string) *UpdateNamespaceSlsConfigsResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateNamespaceSlsConfigsResponseBody) Validate() error {
	return dara.Validate(s)
}
