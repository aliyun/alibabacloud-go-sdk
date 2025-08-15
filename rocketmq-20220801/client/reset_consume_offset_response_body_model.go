// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iResetConsumeOffsetResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ResetConsumeOffsetResponseBody
	GetCode() *string
	SetDynamicCode(v string) *ResetConsumeOffsetResponseBody
	GetDynamicCode() *string
	SetDynamicMessage(v string) *ResetConsumeOffsetResponseBody
	GetDynamicMessage() *string
	SetHttpStatusCode(v int32) *ResetConsumeOffsetResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *ResetConsumeOffsetResponseBody
	GetMessage() *string
	SetRequestId(v string) *ResetConsumeOffsetResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ResetConsumeOffsetResponseBody
	GetSuccess() *bool
}

type ResetConsumeOffsetResponseBody struct {
	// The returned error code.
	//
	// example:
	//
	// MissingInstanceId
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The returned dynamic error code.
	//
	// example:
	//
	// InstanceId
	DynamicCode *string `json:"dynamicCode,omitempty" xml:"dynamicCode,omitempty"`
	// The returned dynamic error message.
	//
	// example:
	//
	// instanceId
	DynamicMessage *string `json:"dynamicMessage,omitempty" xml:"dynamicMessage,omitempty"`
	// The returned HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"httpStatusCode,omitempty" xml:"httpStatusCode,omitempty"`
	// The returned error message.
	//
	// example:
	//
	// The instance cannot be found.
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// F9A95891-EAD4-5A2B-8A30-676CD18921AF
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// Indicates whether the request is successful.
	//
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s ResetConsumeOffsetResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ResetConsumeOffsetResponseBody) GoString() string {
	return s.String()
}

func (s *ResetConsumeOffsetResponseBody) GetCode() *string {
	return s.Code
}

func (s *ResetConsumeOffsetResponseBody) GetDynamicCode() *string {
	return s.DynamicCode
}

func (s *ResetConsumeOffsetResponseBody) GetDynamicMessage() *string {
	return s.DynamicMessage
}

func (s *ResetConsumeOffsetResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ResetConsumeOffsetResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ResetConsumeOffsetResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ResetConsumeOffsetResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ResetConsumeOffsetResponseBody) SetCode(v string) *ResetConsumeOffsetResponseBody {
	s.Code = &v
	return s
}

func (s *ResetConsumeOffsetResponseBody) SetDynamicCode(v string) *ResetConsumeOffsetResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *ResetConsumeOffsetResponseBody) SetDynamicMessage(v string) *ResetConsumeOffsetResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *ResetConsumeOffsetResponseBody) SetHttpStatusCode(v int32) *ResetConsumeOffsetResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ResetConsumeOffsetResponseBody) SetMessage(v string) *ResetConsumeOffsetResponseBody {
	s.Message = &v
	return s
}

func (s *ResetConsumeOffsetResponseBody) SetRequestId(v string) *ResetConsumeOffsetResponseBody {
	s.RequestId = &v
	return s
}

func (s *ResetConsumeOffsetResponseBody) SetSuccess(v bool) *ResetConsumeOffsetResponseBody {
	s.Success = &v
	return s
}

func (s *ResetConsumeOffsetResponseBody) Validate() error {
	return dara.Validate(s)
}
