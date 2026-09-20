// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTopicRequest interface {
	dara.Model
	String() string
	GoString() string
	SetTopicId(v int64) *GetTopicRequest
	GetTopicId() *int64
}

type GetTopicRequest struct {
	// The ID of the event. You can call [listTopics](https://help.aliyun.com/document_detail/173973.html) to obtain the ID.
	//
	// The documentation example is for format demonstration only. Valid TopicId values can be obtained from Data.Topics[].TopicId in the ListTopics response.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1234
	TopicId *int64 `json:"TopicId,omitempty" xml:"TopicId,omitempty"`
}

func (s GetTopicRequest) String() string {
	return dara.Prettify(s)
}

func (s GetTopicRequest) GoString() string {
	return s.String()
}

func (s *GetTopicRequest) GetTopicId() *int64 {
	return s.TopicId
}

func (s *GetTopicRequest) SetTopicId(v int64) *GetTopicRequest {
	s.TopicId = &v
	return s
}

func (s *GetTopicRequest) Validate() error {
	return dara.Validate(s)
}
