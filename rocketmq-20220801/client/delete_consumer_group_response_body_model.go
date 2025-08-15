// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteConsumerGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DeleteConsumerGroupResponseBody
	GetCode() *string
	SetData(v bool) *DeleteConsumerGroupResponseBody
	GetData() *bool
	SetDynamicCode(v string) *DeleteConsumerGroupResponseBody
	GetDynamicCode() *string
	SetDynamicMessage(v string) *DeleteConsumerGroupResponseBody
	GetDynamicMessage() *string
	SetHttpStatusCode(v int32) *DeleteConsumerGroupResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *DeleteConsumerGroupResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeleteConsumerGroupResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeleteConsumerGroupResponseBody
	GetSuccess() *bool
}

type DeleteConsumerGroupResponseBody struct {
	// The error code.
	//
	// example:
	//
	// InvalidConsumerGroupId
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The result data that is returned.
	//
	// example:
	//
	// true
	Data *bool `json:"data,omitempty" xml:"data,omitempty"`
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
	// 400
	HttpStatusCode *int32 `json:"httpStatusCode,omitempty" xml:"httpStatusCode,omitempty"`
	// The error message.
	//
	// example:
	//
	// Parameter consumerGroupId is invalid.
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// The ID of the request. The system generates a unique ID for each request. You can troubleshoot issues based on the request ID.
	//
	// example:
	//
	// C7F94090-3358-506A-97DC-34BC803C****
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// Indicates whether the call is successful.
	//
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s DeleteConsumerGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteConsumerGroupResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteConsumerGroupResponseBody) GetCode() *string {
	return s.Code
}

func (s *DeleteConsumerGroupResponseBody) GetData() *bool {
	return s.Data
}

func (s *DeleteConsumerGroupResponseBody) GetDynamicCode() *string {
	return s.DynamicCode
}

func (s *DeleteConsumerGroupResponseBody) GetDynamicMessage() *string {
	return s.DynamicMessage
}

func (s *DeleteConsumerGroupResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *DeleteConsumerGroupResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteConsumerGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteConsumerGroupResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteConsumerGroupResponseBody) SetCode(v string) *DeleteConsumerGroupResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteConsumerGroupResponseBody) SetData(v bool) *DeleteConsumerGroupResponseBody {
	s.Data = &v
	return s
}

func (s *DeleteConsumerGroupResponseBody) SetDynamicCode(v string) *DeleteConsumerGroupResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *DeleteConsumerGroupResponseBody) SetDynamicMessage(v string) *DeleteConsumerGroupResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *DeleteConsumerGroupResponseBody) SetHttpStatusCode(v int32) *DeleteConsumerGroupResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DeleteConsumerGroupResponseBody) SetMessage(v string) *DeleteConsumerGroupResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteConsumerGroupResponseBody) SetRequestId(v string) *DeleteConsumerGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteConsumerGroupResponseBody) SetSuccess(v bool) *DeleteConsumerGroupResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteConsumerGroupResponseBody) Validate() error {
	return dara.Validate(s)
}
