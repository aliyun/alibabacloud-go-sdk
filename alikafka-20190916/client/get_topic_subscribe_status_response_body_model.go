// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTopicSubscribeStatusResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *GetTopicSubscribeStatusResponseBody
	GetCode() *int32
	SetMessage(v string) *GetTopicSubscribeStatusResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetTopicSubscribeStatusResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetTopicSubscribeStatusResponseBody
	GetSuccess() *bool
	SetTopicSubscribeStatus(v *GetTopicSubscribeStatusResponseBodyTopicSubscribeStatus) *GetTopicSubscribeStatusResponseBody
	GetTopicSubscribeStatus() *GetTopicSubscribeStatusResponseBodyTopicSubscribeStatus
}

type GetTopicSubscribeStatusResponseBody struct {
	// The HTTP status code.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned message.
	//
	// example:
	//
	// operation success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 06084011-E093-46F3-A51F-4B19A8AD****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request is successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// The subscription details.
	TopicSubscribeStatus *GetTopicSubscribeStatusResponseBodyTopicSubscribeStatus `json:"TopicSubscribeStatus,omitempty" xml:"TopicSubscribeStatus,omitempty" type:"Struct"`
}

func (s GetTopicSubscribeStatusResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetTopicSubscribeStatusResponseBody) GoString() string {
	return s.String()
}

func (s *GetTopicSubscribeStatusResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *GetTopicSubscribeStatusResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetTopicSubscribeStatusResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetTopicSubscribeStatusResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetTopicSubscribeStatusResponseBody) GetTopicSubscribeStatus() *GetTopicSubscribeStatusResponseBodyTopicSubscribeStatus {
	return s.TopicSubscribeStatus
}

func (s *GetTopicSubscribeStatusResponseBody) SetCode(v int32) *GetTopicSubscribeStatusResponseBody {
	s.Code = &v
	return s
}

func (s *GetTopicSubscribeStatusResponseBody) SetMessage(v string) *GetTopicSubscribeStatusResponseBody {
	s.Message = &v
	return s
}

func (s *GetTopicSubscribeStatusResponseBody) SetRequestId(v string) *GetTopicSubscribeStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetTopicSubscribeStatusResponseBody) SetSuccess(v bool) *GetTopicSubscribeStatusResponseBody {
	s.Success = &v
	return s
}

func (s *GetTopicSubscribeStatusResponseBody) SetTopicSubscribeStatus(v *GetTopicSubscribeStatusResponseBodyTopicSubscribeStatus) *GetTopicSubscribeStatusResponseBody {
	s.TopicSubscribeStatus = v
	return s
}

func (s *GetTopicSubscribeStatusResponseBody) Validate() error {
	if s.TopicSubscribeStatus != nil {
		if err := s.TopicSubscribeStatus.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetTopicSubscribeStatusResponseBodyTopicSubscribeStatus struct {
	// The groups that subscribe to the topic.
	ConsumerGroups []*string `json:"ConsumerGroups,omitempty" xml:"ConsumerGroups,omitempty" type:"Repeated"`
	// The topic name.
	//
	// example:
	//
	// topic_api_1681624879908
	Topic *string `json:"Topic,omitempty" xml:"Topic,omitempty"`
}

func (s GetTopicSubscribeStatusResponseBodyTopicSubscribeStatus) String() string {
	return dara.Prettify(s)
}

func (s GetTopicSubscribeStatusResponseBodyTopicSubscribeStatus) GoString() string {
	return s.String()
}

func (s *GetTopicSubscribeStatusResponseBodyTopicSubscribeStatus) GetConsumerGroups() []*string {
	return s.ConsumerGroups
}

func (s *GetTopicSubscribeStatusResponseBodyTopicSubscribeStatus) GetTopic() *string {
	return s.Topic
}

func (s *GetTopicSubscribeStatusResponseBodyTopicSubscribeStatus) SetConsumerGroups(v []*string) *GetTopicSubscribeStatusResponseBodyTopicSubscribeStatus {
	s.ConsumerGroups = v
	return s
}

func (s *GetTopicSubscribeStatusResponseBodyTopicSubscribeStatus) SetTopic(v string) *GetTopicSubscribeStatusResponseBodyTopicSubscribeStatus {
	s.Topic = &v
	return s
}

func (s *GetTopicSubscribeStatusResponseBodyTopicSubscribeStatus) Validate() error {
	return dara.Validate(s)
}
