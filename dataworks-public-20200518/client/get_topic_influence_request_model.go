// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTopicInfluenceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetTopicId(v int64) *GetTopicInfluenceRequest
	GetTopicId() *int64
}

type GetTopicInfluenceRequest struct {
	// The ID of the event.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1234
	TopicId *int64 `json:"TopicId,omitempty" xml:"TopicId,omitempty"`
}

func (s GetTopicInfluenceRequest) String() string {
	return dara.Prettify(s)
}

func (s GetTopicInfluenceRequest) GoString() string {
	return s.String()
}

func (s *GetTopicInfluenceRequest) GetTopicId() *int64 {
	return s.TopicId
}

func (s *GetTopicInfluenceRequest) SetTopicId(v int64) *GetTopicInfluenceRequest {
	s.TopicId = &v
	return s
}

func (s *GetTopicInfluenceRequest) Validate() error {
	return dara.Validate(s)
}
