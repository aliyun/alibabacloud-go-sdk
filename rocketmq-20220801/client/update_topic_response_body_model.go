// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateTopicResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *UpdateTopicResponseBody
	GetCode() *string
	SetData(v bool) *UpdateTopicResponseBody
	GetData() *bool
	SetDynamicCode(v string) *UpdateTopicResponseBody
	GetDynamicCode() *string
	SetDynamicMessage(v string) *UpdateTopicResponseBody
	GetDynamicMessage() *string
	SetHttpStatusCode(v int32) *UpdateTopicResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *UpdateTopicResponseBody
	GetMessage() *string
	SetRequestId(v string) *UpdateTopicResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateTopicResponseBody
	GetSuccess() *bool
}

type UpdateTopicResponseBody struct {
	// Error code.
	//
	// example:
	//
	// Topic.NotFound
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// Return result.
	//
	// example:
	//
	// true
	Data *bool `json:"data,omitempty" xml:"data,omitempty"`
	// Dynamic error code
	//
	// example:
	//
	// TopicName
	DynamicCode *string `json:"dynamicCode,omitempty" xml:"dynamicCode,omitempty"`
	// 动态错误信息
	//
	// example:
	//
	// topicName
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
	// The topic cannot be found.
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// Request ID, each request has a unique ID that can be used for troubleshooting and problem localization.
	//
	// example:
	//
	// AF9A8B10-C426-530F-A0DD-96320B39****
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// Whether the execution result is successful.
	//
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s UpdateTopicResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateTopicResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateTopicResponseBody) GetCode() *string {
	return s.Code
}

func (s *UpdateTopicResponseBody) GetData() *bool {
	return s.Data
}

func (s *UpdateTopicResponseBody) GetDynamicCode() *string {
	return s.DynamicCode
}

func (s *UpdateTopicResponseBody) GetDynamicMessage() *string {
	return s.DynamicMessage
}

func (s *UpdateTopicResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *UpdateTopicResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateTopicResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateTopicResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateTopicResponseBody) SetCode(v string) *UpdateTopicResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateTopicResponseBody) SetData(v bool) *UpdateTopicResponseBody {
	s.Data = &v
	return s
}

func (s *UpdateTopicResponseBody) SetDynamicCode(v string) *UpdateTopicResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *UpdateTopicResponseBody) SetDynamicMessage(v string) *UpdateTopicResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *UpdateTopicResponseBody) SetHttpStatusCode(v int32) *UpdateTopicResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *UpdateTopicResponseBody) SetMessage(v string) *UpdateTopicResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateTopicResponseBody) SetRequestId(v string) *UpdateTopicResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateTopicResponseBody) SetSuccess(v bool) *UpdateTopicResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateTopicResponseBody) Validate() error {
	return dara.Validate(s)
}
