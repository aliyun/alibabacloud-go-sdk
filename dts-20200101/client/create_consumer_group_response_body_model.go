// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateConsumerGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetConsumerGroupID(v string) *CreateConsumerGroupResponseBody
	GetConsumerGroupID() *string
	SetErrCode(v string) *CreateConsumerGroupResponseBody
	GetErrCode() *string
	SetErrMessage(v string) *CreateConsumerGroupResponseBody
	GetErrMessage() *string
	SetRequestId(v string) *CreateConsumerGroupResponseBody
	GetRequestId() *string
	SetSuccess(v string) *CreateConsumerGroupResponseBody
	GetSuccess() *string
}

type CreateConsumerGroupResponseBody struct {
	// The ID of the consumer group.
	//
	// example:
	//
	// dtswc411cg617p****
	ConsumerGroupID *string `json:"ConsumerGroupID,omitempty" xml:"ConsumerGroupID,omitempty"`
	// The error code returned if the call failed.
	//
	// example:
	//
	// InternalError
	ErrCode *string `json:"ErrCode,omitempty" xml:"ErrCode,omitempty"`
	// The error message returned if the call failed.
	//
	// example:
	//
	// The request processing has failed due to some unknown error.
	ErrMessage *string `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 6063641E-BAD1-4BA7-B70B-26FFFD18****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the call was successful.
	//
	// example:
	//
	// true
	Success *string `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateConsumerGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateConsumerGroupResponseBody) GoString() string {
	return s.String()
}

func (s *CreateConsumerGroupResponseBody) GetConsumerGroupID() *string {
	return s.ConsumerGroupID
}

func (s *CreateConsumerGroupResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *CreateConsumerGroupResponseBody) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *CreateConsumerGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateConsumerGroupResponseBody) GetSuccess() *string {
	return s.Success
}

func (s *CreateConsumerGroupResponseBody) SetConsumerGroupID(v string) *CreateConsumerGroupResponseBody {
	s.ConsumerGroupID = &v
	return s
}

func (s *CreateConsumerGroupResponseBody) SetErrCode(v string) *CreateConsumerGroupResponseBody {
	s.ErrCode = &v
	return s
}

func (s *CreateConsumerGroupResponseBody) SetErrMessage(v string) *CreateConsumerGroupResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *CreateConsumerGroupResponseBody) SetRequestId(v string) *CreateConsumerGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateConsumerGroupResponseBody) SetSuccess(v string) *CreateConsumerGroupResponseBody {
	s.Success = &v
	return s
}

func (s *CreateConsumerGroupResponseBody) Validate() error {
	return dara.Validate(s)
}
