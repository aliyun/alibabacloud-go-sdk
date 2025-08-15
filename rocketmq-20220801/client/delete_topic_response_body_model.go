// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteTopicResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DeleteTopicResponseBody
	GetCode() *string
	SetData(v bool) *DeleteTopicResponseBody
	GetData() *bool
	SetDynamicCode(v string) *DeleteTopicResponseBody
	GetDynamicCode() *string
	SetDynamicMessage(v string) *DeleteTopicResponseBody
	GetDynamicMessage() *string
	SetHttpStatusCode(v int32) *DeleteTopicResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *DeleteTopicResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeleteTopicResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeleteTopicResponseBody
	GetSuccess() *bool
}

type DeleteTopicResponseBody struct {
	// The error code.
	//
	// example:
	//
	// Topic.NotFound
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
	// TopicName
	DynamicCode *string `json:"dynamicCode,omitempty" xml:"dynamicCode,omitempty"`
	// The dynamic error message.
	//
	// example:
	//
	// topicName
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
	// The topic cannot be found.
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// The ID of the request. The system generates a unique ID for each request. You can troubleshoot issues based on the request ID.
	//
	// example:
	//
	// AF9A8B10-C426-530F-A0DD-96320B39****
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// Indicates whether the call is successful.
	//
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s DeleteTopicResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteTopicResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteTopicResponseBody) GetCode() *string {
	return s.Code
}

func (s *DeleteTopicResponseBody) GetData() *bool {
	return s.Data
}

func (s *DeleteTopicResponseBody) GetDynamicCode() *string {
	return s.DynamicCode
}

func (s *DeleteTopicResponseBody) GetDynamicMessage() *string {
	return s.DynamicMessage
}

func (s *DeleteTopicResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *DeleteTopicResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteTopicResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteTopicResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteTopicResponseBody) SetCode(v string) *DeleteTopicResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteTopicResponseBody) SetData(v bool) *DeleteTopicResponseBody {
	s.Data = &v
	return s
}

func (s *DeleteTopicResponseBody) SetDynamicCode(v string) *DeleteTopicResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *DeleteTopicResponseBody) SetDynamicMessage(v string) *DeleteTopicResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *DeleteTopicResponseBody) SetHttpStatusCode(v int32) *DeleteTopicResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DeleteTopicResponseBody) SetMessage(v string) *DeleteTopicResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteTopicResponseBody) SetRequestId(v string) *DeleteTopicResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteTopicResponseBody) SetSuccess(v bool) *DeleteTopicResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteTopicResponseBody) Validate() error {
	return dara.Validate(s)
}
