// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteTopicRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRegionId(v string) *DeleteTopicRequest
	GetRegionId() *string
	SetTopicIdList(v []*int64) *DeleteTopicRequest
	GetTopicIdList() []*int64
}

type DeleteTopicRequest struct {
	// example:
	//
	// cn-shanghai
	RegionId    *string  `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	TopicIdList []*int64 `json:"TopicIdList,omitempty" xml:"TopicIdList,omitempty" type:"Repeated"`
}

func (s DeleteTopicRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteTopicRequest) GoString() string {
	return s.String()
}

func (s *DeleteTopicRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteTopicRequest) GetTopicIdList() []*int64 {
	return s.TopicIdList
}

func (s *DeleteTopicRequest) SetRegionId(v string) *DeleteTopicRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteTopicRequest) SetTopicIdList(v []*int64) *DeleteTopicRequest {
	s.TopicIdList = v
	return s
}

func (s *DeleteTopicRequest) Validate() error {
	return dara.Validate(s)
}
