// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iVerifySendMessageResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *VerifySendMessageResponseBody
	GetCode() *string
	SetData(v string) *VerifySendMessageResponseBody
	GetData() *string
	SetDynamicCode(v string) *VerifySendMessageResponseBody
	GetDynamicCode() *string
	SetDynamicMessage(v string) *VerifySendMessageResponseBody
	GetDynamicMessage() *string
	SetHttpStatusCode(v int32) *VerifySendMessageResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *VerifySendMessageResponseBody
	GetMessage() *string
	SetRequestId(v string) *VerifySendMessageResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *VerifySendMessageResponseBody
	GetSuccess() *bool
}

type VerifySendMessageResponseBody struct {
	// The error code.
	//
	// example:
	//
	// InvalidConsumerGroupId
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The returned data.
	//
	// example:
	//
	// 0A64228900207A4F0F2931A4E0D40BE5
	Data *string `json:"data,omitempty" xml:"data,omitempty"`
	// The dynamic error code.
	//
	// example:
	//
	// ConsumerGroupId
	DynamicCode *string `json:"dynamicCode,omitempty" xml:"dynamicCode,omitempty"`
	// The dynamic error message.
	//
	// example:
	//
	// consumerGroupId
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
	// The instance cannot be found.
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 3BD2C19B-66DE-59C7-B2F6-FD1BE21DC8C1
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s VerifySendMessageResponseBody) String() string {
	return dara.Prettify(s)
}

func (s VerifySendMessageResponseBody) GoString() string {
	return s.String()
}

func (s *VerifySendMessageResponseBody) GetCode() *string {
	return s.Code
}

func (s *VerifySendMessageResponseBody) GetData() *string {
	return s.Data
}

func (s *VerifySendMessageResponseBody) GetDynamicCode() *string {
	return s.DynamicCode
}

func (s *VerifySendMessageResponseBody) GetDynamicMessage() *string {
	return s.DynamicMessage
}

func (s *VerifySendMessageResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *VerifySendMessageResponseBody) GetMessage() *string {
	return s.Message
}

func (s *VerifySendMessageResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *VerifySendMessageResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *VerifySendMessageResponseBody) SetCode(v string) *VerifySendMessageResponseBody {
	s.Code = &v
	return s
}

func (s *VerifySendMessageResponseBody) SetData(v string) *VerifySendMessageResponseBody {
	s.Data = &v
	return s
}

func (s *VerifySendMessageResponseBody) SetDynamicCode(v string) *VerifySendMessageResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *VerifySendMessageResponseBody) SetDynamicMessage(v string) *VerifySendMessageResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *VerifySendMessageResponseBody) SetHttpStatusCode(v int32) *VerifySendMessageResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *VerifySendMessageResponseBody) SetMessage(v string) *VerifySendMessageResponseBody {
	s.Message = &v
	return s
}

func (s *VerifySendMessageResponseBody) SetRequestId(v string) *VerifySendMessageResponseBody {
	s.RequestId = &v
	return s
}

func (s *VerifySendMessageResponseBody) SetSuccess(v bool) *VerifySendMessageResponseBody {
	s.Success = &v
	return s
}

func (s *VerifySendMessageResponseBody) Validate() error {
	return dara.Validate(s)
}
