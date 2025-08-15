// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iVerifyConsumeMessageResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *VerifyConsumeMessageResponseBody
	GetCode() *string
	SetData(v bool) *VerifyConsumeMessageResponseBody
	GetData() *bool
	SetDynamicCode(v string) *VerifyConsumeMessageResponseBody
	GetDynamicCode() *string
	SetDynamicMessage(v string) *VerifyConsumeMessageResponseBody
	GetDynamicMessage() *string
	SetHttpStatusCode(v int32) *VerifyConsumeMessageResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *VerifyConsumeMessageResponseBody
	GetMessage() *string
	SetRequestId(v string) *VerifyConsumeMessageResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *VerifyConsumeMessageResponseBody
	GetSuccess() *bool
}

type VerifyConsumeMessageResponseBody struct {
	// The error code.
	//
	// example:
	//
	// Topic.NotFound
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The returned data.
	//
	// example:
	//
	// true
	Data *bool `json:"data,omitempty" xml:"data,omitempty"`
	// The dynamic error code.
	//
	// example:
	//
	// InstanceId
	DynamicCode *string `json:"dynamicCode,omitempty" xml:"dynamicCode,omitempty"`
	// The dynamic error message.
	//
	// example:
	//
	// instanceId
	DynamicMessage *string `json:"dynamicMessage,omitempty" xml:"dynamicMessage,omitempty"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"httpStatusCode,omitempty" xml:"httpStatusCode,omitempty"`
	// The error message.
	//
	// example:
	//
	// Parameter instanceId is mandatory for this action .
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 5304143F-AB0E-5AB4-A227-7C5489216FD5
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s VerifyConsumeMessageResponseBody) String() string {
	return dara.Prettify(s)
}

func (s VerifyConsumeMessageResponseBody) GoString() string {
	return s.String()
}

func (s *VerifyConsumeMessageResponseBody) GetCode() *string {
	return s.Code
}

func (s *VerifyConsumeMessageResponseBody) GetData() *bool {
	return s.Data
}

func (s *VerifyConsumeMessageResponseBody) GetDynamicCode() *string {
	return s.DynamicCode
}

func (s *VerifyConsumeMessageResponseBody) GetDynamicMessage() *string {
	return s.DynamicMessage
}

func (s *VerifyConsumeMessageResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *VerifyConsumeMessageResponseBody) GetMessage() *string {
	return s.Message
}

func (s *VerifyConsumeMessageResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *VerifyConsumeMessageResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *VerifyConsumeMessageResponseBody) SetCode(v string) *VerifyConsumeMessageResponseBody {
	s.Code = &v
	return s
}

func (s *VerifyConsumeMessageResponseBody) SetData(v bool) *VerifyConsumeMessageResponseBody {
	s.Data = &v
	return s
}

func (s *VerifyConsumeMessageResponseBody) SetDynamicCode(v string) *VerifyConsumeMessageResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *VerifyConsumeMessageResponseBody) SetDynamicMessage(v string) *VerifyConsumeMessageResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *VerifyConsumeMessageResponseBody) SetHttpStatusCode(v int32) *VerifyConsumeMessageResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *VerifyConsumeMessageResponseBody) SetMessage(v string) *VerifyConsumeMessageResponseBody {
	s.Message = &v
	return s
}

func (s *VerifyConsumeMessageResponseBody) SetRequestId(v string) *VerifyConsumeMessageResponseBody {
	s.RequestId = &v
	return s
}

func (s *VerifyConsumeMessageResponseBody) SetSuccess(v bool) *VerifyConsumeMessageResponseBody {
	s.Success = &v
	return s
}

func (s *VerifyConsumeMessageResponseBody) Validate() error {
	return dara.Validate(s)
}
