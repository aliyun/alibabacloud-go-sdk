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
	// The event ID. You can call the [ListTopics](https://help.aliyun.com/document_detail/173973.html) operation to query the ID.
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
