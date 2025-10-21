// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTopicRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRegionId(v string) *GetTopicRequest
	GetRegionId() *string
	SetTopicId(v int64) *GetTopicRequest
	GetTopicId() *int64
}

type GetTopicRequest struct {
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// 216
	TopicId *int64 `json:"TopicId,omitempty" xml:"TopicId,omitempty"`
}

func (s GetTopicRequest) String() string {
	return dara.Prettify(s)
}

func (s GetTopicRequest) GoString() string {
	return s.String()
}

func (s *GetTopicRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *GetTopicRequest) GetTopicId() *int64 {
	return s.TopicId
}

func (s *GetTopicRequest) SetRegionId(v string) *GetTopicRequest {
	s.RegionId = &v
	return s
}

func (s *GetTopicRequest) SetTopicId(v int64) *GetTopicRequest {
	s.TopicId = &v
	return s
}

func (s *GetTopicRequest) Validate() error {
	return dara.Validate(s)
}
