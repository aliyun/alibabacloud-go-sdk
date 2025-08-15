// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateTopicResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CreateTopicResponseBody
	GetCode() *string
	SetData(v bool) *CreateTopicResponseBody
	GetData() *bool
	SetDynamicCode(v string) *CreateTopicResponseBody
	GetDynamicCode() *string
	SetDynamicMessage(v string) *CreateTopicResponseBody
	GetDynamicMessage() *string
	SetHttpStatusCode(v int32) *CreateTopicResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *CreateTopicResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreateTopicResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreateTopicResponseBody
	GetSuccess() *bool
}

type CreateTopicResponseBody struct {
	// Error code.
	//
	// example:
	//
	// Topic.Existed
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// Return result.
	//
	// example:
	//
	// true
	Data *bool `json:"data,omitempty" xml:"data,omitempty"`
	// Dynamic error code.
	//
	// example:
	//
	// TopicName
	DynamicCode *string `json:"dynamicCode,omitempty" xml:"dynamicCode,omitempty"`
	// Dynamic error message.
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
	// The topic already exists.
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// Request ID, each request\\"s ID is unique and can be used for troubleshooting and problem localization.
	//
	// example:
	//
	// AF9A8B10-C426-530F-A0DD-96320B39****
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// Indicates whether the execution was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s CreateTopicResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateTopicResponseBody) GoString() string {
	return s.String()
}

func (s *CreateTopicResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreateTopicResponseBody) GetData() *bool {
	return s.Data
}

func (s *CreateTopicResponseBody) GetDynamicCode() *string {
	return s.DynamicCode
}

func (s *CreateTopicResponseBody) GetDynamicMessage() *string {
	return s.DynamicMessage
}

func (s *CreateTopicResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *CreateTopicResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateTopicResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateTopicResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateTopicResponseBody) SetCode(v string) *CreateTopicResponseBody {
	s.Code = &v
	return s
}

func (s *CreateTopicResponseBody) SetData(v bool) *CreateTopicResponseBody {
	s.Data = &v
	return s
}

func (s *CreateTopicResponseBody) SetDynamicCode(v string) *CreateTopicResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *CreateTopicResponseBody) SetDynamicMessage(v string) *CreateTopicResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *CreateTopicResponseBody) SetHttpStatusCode(v int32) *CreateTopicResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *CreateTopicResponseBody) SetMessage(v string) *CreateTopicResponseBody {
	s.Message = &v
	return s
}

func (s *CreateTopicResponseBody) SetRequestId(v string) *CreateTopicResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateTopicResponseBody) SetSuccess(v bool) *CreateTopicResponseBody {
	s.Success = &v
	return s
}

func (s *CreateTopicResponseBody) Validate() error {
	return dara.Validate(s)
}
