// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyTopicRemarkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *ModifyTopicRemarkRequest
	GetInstanceId() *string
	SetRegionId(v string) *ModifyTopicRemarkRequest
	GetRegionId() *string
	SetRemark(v string) *ModifyTopicRemarkRequest
	GetRemark() *string
	SetTopic(v string) *ModifyTopicRemarkRequest
	GetTopic() *string
}

type ModifyTopicRemarkRequest struct {
	// The ID of the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// alikafka_post-cn-0pp1l9z****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The region ID of the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The description of the topic.
	//
	// example:
	//
	// testremark
	Remark *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	// The name of the topic.
	//
	// This parameter is required.
	//
	// example:
	//
	// alikafka_post-cn-0pp1l9z8****
	Topic *string `json:"Topic,omitempty" xml:"Topic,omitempty"`
}

func (s ModifyTopicRemarkRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyTopicRemarkRequest) GoString() string {
	return s.String()
}

func (s *ModifyTopicRemarkRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ModifyTopicRemarkRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifyTopicRemarkRequest) GetRemark() *string {
	return s.Remark
}

func (s *ModifyTopicRemarkRequest) GetTopic() *string {
	return s.Topic
}

func (s *ModifyTopicRemarkRequest) SetInstanceId(v string) *ModifyTopicRemarkRequest {
	s.InstanceId = &v
	return s
}

func (s *ModifyTopicRemarkRequest) SetRegionId(v string) *ModifyTopicRemarkRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyTopicRemarkRequest) SetRemark(v string) *ModifyTopicRemarkRequest {
	s.Remark = &v
	return s
}

func (s *ModifyTopicRemarkRequest) SetTopic(v string) *ModifyTopicRemarkRequest {
	s.Topic = &v
	return s
}

func (s *ModifyTopicRemarkRequest) Validate() error {
	return dara.Validate(s)
}
