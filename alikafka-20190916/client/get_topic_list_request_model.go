// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTopicListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCurrentPage(v string) *GetTopicListRequest
	GetCurrentPage() *string
	SetInstanceId(v string) *GetTopicListRequest
	GetInstanceId() *string
	SetPageSize(v string) *GetTopicListRequest
	GetPageSize() *string
	SetRegionId(v string) *GetTopicListRequest
	GetRegionId() *string
	SetTopic(v string) *GetTopicListRequest
	GetTopic() *string
}

type GetTopicListRequest struct {
	// The page number. Default value: 1
	//
	// example:
	//
	// 1
	CurrentPage *string `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The ID of the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// alikafka_pre-cn-0pp1954n****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The number of entries to return on each page. Default value: 10.
	//
	// example:
	//
	// 10
	PageSize *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The region ID of the instance to which the topics that you want to query belong.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The name of the topic that you want to query.
	//
	// example:
	//
	// topic_name
	Topic *string `json:"Topic,omitempty" xml:"Topic,omitempty"`
}

func (s GetTopicListRequest) String() string {
	return dara.Prettify(s)
}

func (s GetTopicListRequest) GoString() string {
	return s.String()
}

func (s *GetTopicListRequest) GetCurrentPage() *string {
	return s.CurrentPage
}

func (s *GetTopicListRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetTopicListRequest) GetPageSize() *string {
	return s.PageSize
}

func (s *GetTopicListRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *GetTopicListRequest) GetTopic() *string {
	return s.Topic
}

func (s *GetTopicListRequest) SetCurrentPage(v string) *GetTopicListRequest {
	s.CurrentPage = &v
	return s
}

func (s *GetTopicListRequest) SetInstanceId(v string) *GetTopicListRequest {
	s.InstanceId = &v
	return s
}

func (s *GetTopicListRequest) SetPageSize(v string) *GetTopicListRequest {
	s.PageSize = &v
	return s
}

func (s *GetTopicListRequest) SetRegionId(v string) *GetTopicListRequest {
	s.RegionId = &v
	return s
}

func (s *GetTopicListRequest) SetTopic(v string) *GetTopicListRequest {
	s.Topic = &v
	return s
}

func (s *GetTopicListRequest) Validate() error {
	return dara.Validate(s)
}
