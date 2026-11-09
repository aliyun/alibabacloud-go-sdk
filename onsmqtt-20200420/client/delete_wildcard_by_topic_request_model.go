// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteWildcardByTopicRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFirstTopic(v string) *DeleteWildcardByTopicRequest
	GetFirstTopic() *string
	SetInstanceId(v string) *DeleteWildcardByTopicRequest
	GetInstanceId() *string
	SetWildcards(v string) *DeleteWildcardByTopicRequest
	GetWildcards() *string
}

type DeleteWildcardByTopicRequest struct {
	// The first-level topic.
	//
	// This parameter is required.
	//
	// example:
	//
	// test
	FirstTopic *string `json:"FirstTopic,omitempty" xml:"FirstTopic,omitempty"`
	// The ID of the instance to which the group ID that you want to delete belongs.
	//
	// This parameter is required.
	//
	// example:
	//
	// mqtt-cn-0pp1ldu****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The wildcard subscriptions to delete.
	//
	// This parameter is required.
	//
	// example:
	//
	// test/#
	Wildcards *string `json:"Wildcards,omitempty" xml:"Wildcards,omitempty"`
}

func (s DeleteWildcardByTopicRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteWildcardByTopicRequest) GoString() string {
	return s.String()
}

func (s *DeleteWildcardByTopicRequest) GetFirstTopic() *string {
	return s.FirstTopic
}

func (s *DeleteWildcardByTopicRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DeleteWildcardByTopicRequest) GetWildcards() *string {
	return s.Wildcards
}

func (s *DeleteWildcardByTopicRequest) SetFirstTopic(v string) *DeleteWildcardByTopicRequest {
	s.FirstTopic = &v
	return s
}

func (s *DeleteWildcardByTopicRequest) SetInstanceId(v string) *DeleteWildcardByTopicRequest {
	s.InstanceId = &v
	return s
}

func (s *DeleteWildcardByTopicRequest) SetWildcards(v string) *DeleteWildcardByTopicRequest {
	s.Wildcards = &v
	return s
}

func (s *DeleteWildcardByTopicRequest) Validate() error {
	return dara.Validate(s)
}
