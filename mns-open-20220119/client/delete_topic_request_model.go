// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteTopicRequest interface {
	dara.Model
	String() string
	GoString() string
	SetTopicName(v string) *DeleteTopicRequest
	GetTopicName() *string
}

type DeleteTopicRequest struct {
	// The name of the topic that you want to delete.
	//
	// This parameter is required.
	//
	// example:
	//
	// tf-testAccMNSTopic-112965059402264645
	TopicName *string `json:"TopicName,omitempty" xml:"TopicName,omitempty"`
}

func (s DeleteTopicRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteTopicRequest) GoString() string {
	return s.String()
}

func (s *DeleteTopicRequest) GetTopicName() *string {
	return s.TopicName
}

func (s *DeleteTopicRequest) SetTopicName(v string) *DeleteTopicRequest {
	s.TopicName = &v
	return s
}

func (s *DeleteTopicRequest) Validate() error {
	return dara.Validate(s)
}
