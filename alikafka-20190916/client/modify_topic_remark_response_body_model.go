// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyTopicRemarkResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *ModifyTopicRemarkResponseBody
	GetCode() *int32
	SetMessage(v string) *ModifyTopicRemarkResponseBody
	GetMessage() *string
	SetRequestId(v string) *ModifyTopicRemarkResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ModifyTopicRemarkResponseBody
	GetSuccess() *bool
}

type ModifyTopicRemarkResponseBody struct {
	// Return code. A value of 200 indicates success.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// Return message.
	//
	// example:
	//
	// operation success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Request ID.
	//
	// example:
	//
	// DB6F1BEA-903B-4FD8-8809-46E7E9CE***
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the call was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ModifyTopicRemarkResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyTopicRemarkResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyTopicRemarkResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *ModifyTopicRemarkResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ModifyTopicRemarkResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyTopicRemarkResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ModifyTopicRemarkResponseBody) SetCode(v int32) *ModifyTopicRemarkResponseBody {
	s.Code = &v
	return s
}

func (s *ModifyTopicRemarkResponseBody) SetMessage(v string) *ModifyTopicRemarkResponseBody {
	s.Message = &v
	return s
}

func (s *ModifyTopicRemarkResponseBody) SetRequestId(v string) *ModifyTopicRemarkResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyTopicRemarkResponseBody) SetSuccess(v bool) *ModifyTopicRemarkResponseBody {
	s.Success = &v
	return s
}

func (s *ModifyTopicRemarkResponseBody) Validate() error {
	return dara.Validate(s)
}
