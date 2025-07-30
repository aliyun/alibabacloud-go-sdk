// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateConsumerChannelResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetConsumerGroupID(v string) *CreateConsumerChannelResponseBody
	GetConsumerGroupID() *string
	SetErrCode(v string) *CreateConsumerChannelResponseBody
	GetErrCode() *string
	SetErrMessage(v string) *CreateConsumerChannelResponseBody
	GetErrMessage() *string
	SetHttpStatusCode(v string) *CreateConsumerChannelResponseBody
	GetHttpStatusCode() *string
	SetRequestId(v string) *CreateConsumerChannelResponseBody
	GetRequestId() *string
	SetSuccess(v string) *CreateConsumerChannelResponseBody
	GetSuccess() *string
}

type CreateConsumerChannelResponseBody struct {
	// The ID of the consumer group. You can specify this parameter on a downstream client when you consume tracked data.
	//
	// example:
	//
	// dtsor2y66j4219****
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
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *string `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 5F566C5B-E5B0-4020-A531-FC6F5005****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the call was successful.
	//
	// example:
	//
	// true
	Success *string `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateConsumerChannelResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateConsumerChannelResponseBody) GoString() string {
	return s.String()
}

func (s *CreateConsumerChannelResponseBody) GetConsumerGroupID() *string {
	return s.ConsumerGroupID
}

func (s *CreateConsumerChannelResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *CreateConsumerChannelResponseBody) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *CreateConsumerChannelResponseBody) GetHttpStatusCode() *string {
	return s.HttpStatusCode
}

func (s *CreateConsumerChannelResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateConsumerChannelResponseBody) GetSuccess() *string {
	return s.Success
}

func (s *CreateConsumerChannelResponseBody) SetConsumerGroupID(v string) *CreateConsumerChannelResponseBody {
	s.ConsumerGroupID = &v
	return s
}

func (s *CreateConsumerChannelResponseBody) SetErrCode(v string) *CreateConsumerChannelResponseBody {
	s.ErrCode = &v
	return s
}

func (s *CreateConsumerChannelResponseBody) SetErrMessage(v string) *CreateConsumerChannelResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *CreateConsumerChannelResponseBody) SetHttpStatusCode(v string) *CreateConsumerChannelResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *CreateConsumerChannelResponseBody) SetRequestId(v string) *CreateConsumerChannelResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateConsumerChannelResponseBody) SetSuccess(v string) *CreateConsumerChannelResponseBody {
	s.Success = &v
	return s
}

func (s *CreateConsumerChannelResponseBody) Validate() error {
	return dara.Validate(s)
}
