// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateConsumerGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *UpdateConsumerGroupResponseBody
	GetCode() *string
	SetData(v bool) *UpdateConsumerGroupResponseBody
	GetData() *bool
	SetDynamicCode(v string) *UpdateConsumerGroupResponseBody
	GetDynamicCode() *string
	SetDynamicMessage(v string) *UpdateConsumerGroupResponseBody
	GetDynamicMessage() *string
	SetHttpStatusCode(v int32) *UpdateConsumerGroupResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *UpdateConsumerGroupResponseBody
	GetMessage() *string
	SetRequestId(v string) *UpdateConsumerGroupResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateConsumerGroupResponseBody
	GetSuccess() *bool
}

type UpdateConsumerGroupResponseBody struct {
	// Error code.
	//
	// example:
	//
	// InvalidDeliveryOrderType
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The result returned.
	//
	// example:
	//
	// true
	Data *bool `json:"data,omitempty" xml:"data,omitempty"`
	// Dynamic error code.
	//
	// example:
	//
	// xxx
	DynamicCode *string `json:"dynamicCode,omitempty" xml:"dynamicCode,omitempty"`
	// Dynamic error message.
	//
	// example:
	//
	// xxx
	DynamicMessage *string `json:"dynamicMessage,omitempty" xml:"dynamicMessage,omitempty"`
	// HTTP status code.
	//
	// example:
	//
	// 400
	HttpStatusCode *int32 `json:"httpStatusCode,omitempty" xml:"httpStatusCode,omitempty"`
	// Error message.
	//
	// example:
	//
	// Parameter deliveryOrderType is invalid.
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// The request ID, which is unique for each request and can be used for troubleshooting and problem localization.
	//
	// example:
	//
	// C7F94090-3358-506A-97DC-34BC803C****
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// Indicates whether the execution was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s UpdateConsumerGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateConsumerGroupResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateConsumerGroupResponseBody) GetCode() *string {
	return s.Code
}

func (s *UpdateConsumerGroupResponseBody) GetData() *bool {
	return s.Data
}

func (s *UpdateConsumerGroupResponseBody) GetDynamicCode() *string {
	return s.DynamicCode
}

func (s *UpdateConsumerGroupResponseBody) GetDynamicMessage() *string {
	return s.DynamicMessage
}

func (s *UpdateConsumerGroupResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *UpdateConsumerGroupResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateConsumerGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateConsumerGroupResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateConsumerGroupResponseBody) SetCode(v string) *UpdateConsumerGroupResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateConsumerGroupResponseBody) SetData(v bool) *UpdateConsumerGroupResponseBody {
	s.Data = &v
	return s
}

func (s *UpdateConsumerGroupResponseBody) SetDynamicCode(v string) *UpdateConsumerGroupResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *UpdateConsumerGroupResponseBody) SetDynamicMessage(v string) *UpdateConsumerGroupResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *UpdateConsumerGroupResponseBody) SetHttpStatusCode(v int32) *UpdateConsumerGroupResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *UpdateConsumerGroupResponseBody) SetMessage(v string) *UpdateConsumerGroupResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateConsumerGroupResponseBody) SetRequestId(v string) *UpdateConsumerGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateConsumerGroupResponseBody) SetSuccess(v bool) *UpdateConsumerGroupResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateConsumerGroupResponseBody) Validate() error {
	return dara.Validate(s)
}
