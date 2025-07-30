// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyConsumerGroupPasswordResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrCode(v string) *ModifyConsumerGroupPasswordResponseBody
	GetErrCode() *string
	SetErrMessage(v string) *ModifyConsumerGroupPasswordResponseBody
	GetErrMessage() *string
	SetRequestId(v string) *ModifyConsumerGroupPasswordResponseBody
	GetRequestId() *string
	SetSuccess(v string) *ModifyConsumerGroupPasswordResponseBody
	GetSuccess() *string
}

type ModifyConsumerGroupPasswordResponseBody struct {
	// The current password of the consumer group.
	//
	// example:
	//
	// InternalError
	ErrCode *string `json:"ErrCode,omitempty" xml:"ErrCode,omitempty"`
	// The ID of the Alibaba Cloud account. You do not need to specify this parameter because this parameter will be removed in the future.
	//
	// example:
	//
	// The request processing has failed due to some unknown error.
	ErrMessage *string `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty"`
	// The username of the consumer group. You can call the [DescribeConsumerGroup](https://help.aliyun.com/document_detail/122886.html) operation to query the username.
	//
	// example:
	//
	// A06B5CFF-9576-4BC1-BE62-A3D43E1F****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The new password of the consumer group.
	//
	// 	- A password must contain two or more of the following characters: uppercase letters, lowercase letters, digits, and special characters.
	//
	// 	- A password must be 8 to 32 characters in length.
	//
	// example:
	//
	// true
	Success *string `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ModifyConsumerGroupPasswordResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyConsumerGroupPasswordResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyConsumerGroupPasswordResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *ModifyConsumerGroupPasswordResponseBody) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *ModifyConsumerGroupPasswordResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyConsumerGroupPasswordResponseBody) GetSuccess() *string {
	return s.Success
}

func (s *ModifyConsumerGroupPasswordResponseBody) SetErrCode(v string) *ModifyConsumerGroupPasswordResponseBody {
	s.ErrCode = &v
	return s
}

func (s *ModifyConsumerGroupPasswordResponseBody) SetErrMessage(v string) *ModifyConsumerGroupPasswordResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *ModifyConsumerGroupPasswordResponseBody) SetRequestId(v string) *ModifyConsumerGroupPasswordResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyConsumerGroupPasswordResponseBody) SetSuccess(v string) *ModifyConsumerGroupPasswordResponseBody {
	s.Success = &v
	return s
}

func (s *ModifyConsumerGroupPasswordResponseBody) Validate() error {
	return dara.Validate(s)
}
