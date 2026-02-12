// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOnsTopicListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *OnsTopicListRequest
	GetInstanceId() *string
	SetTag(v []*OnsTopicListRequestTag) *OnsTopicListRequest
	GetTag() []*OnsTopicListRequestTag
	SetTopic(v string) *OnsTopicListRequest
	GetTopic() *string
}

type OnsTopicListRequest struct {
	// The ID of the instance that contains the topics you want to query.
	//
	// example:
	//
	// MQ_INST_188077086902****_BXSuW61e
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The list of tags that are attached to the topic. A maximum of 20 tags can be included in the list.
	Tag []*OnsTopicListRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	// The name of the topic that you want to query. This parameter is required if you want to query a specific topic. If you do not include this parameter in a request, all topics that you can access are queried.
	//
	// example:
	//
	// test
	Topic *string `json:"Topic,omitempty" xml:"Topic,omitempty"`
}

func (s OnsTopicListRequest) String() string {
	return dara.Prettify(s)
}

func (s OnsTopicListRequest) GoString() string {
	return s.String()
}

func (s *OnsTopicListRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *OnsTopicListRequest) GetTag() []*OnsTopicListRequestTag {
	return s.Tag
}

func (s *OnsTopicListRequest) GetTopic() *string {
	return s.Topic
}

func (s *OnsTopicListRequest) SetInstanceId(v string) *OnsTopicListRequest {
	s.InstanceId = &v
	return s
}

func (s *OnsTopicListRequest) SetTag(v []*OnsTopicListRequestTag) *OnsTopicListRequest {
	s.Tag = v
	return s
}

func (s *OnsTopicListRequest) SetTopic(v string) *OnsTopicListRequest {
	s.Topic = &v
	return s
}

func (s *OnsTopicListRequest) Validate() error {
	if s.Tag != nil {
		for _, item := range s.Tag {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type OnsTopicListRequestTag struct {
	// The key of the tag that is attached to the topics you want to query. This parameter is not required. If you configure this parameter, you must also configure the **Value*	- parameter.***	- If you include the Key and Value parameters in a request, this operation queries only the topics that use the specified tag. If you do not include these parameters in a request, this operation queries all topics that you can access.
	//
	// 	- The value of this parameter cannot be an empty string.
	//
	// 	- A tag value can be up to 128 characters in length and cannot start with `acs:` or `aliyun`. It cannot contain `http://` or `https://`.
	//
	// This parameter is required.
	//
	// example:
	//
	// CartService
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of the tag that is attached to the topics you want to query. This parameter is not required. If you configure this parameter, you must also configure the **Key*	- parameter.***	- If you include the Key and Value parameters in a request, this operation queries only the topics that use the specified tag. If you do not include these parameters in a request, this operation queries all topics that you can access.
	//
	// 	- The value of this parameter can be an empty string.
	//
	// 	- A tag key can be up to 128 characters in length and cannot contain `http://` or `https://`. It cannot start with `acs:` or `aliyun`.
	//
	// This parameter is required.
	//
	// example:
	//
	// ServiceA
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s OnsTopicListRequestTag) String() string {
	return dara.Prettify(s)
}

func (s OnsTopicListRequestTag) GoString() string {
	return s.String()
}

func (s *OnsTopicListRequestTag) GetKey() *string {
	return s.Key
}

func (s *OnsTopicListRequestTag) GetValue() *string {
	return s.Value
}

func (s *OnsTopicListRequestTag) SetKey(v string) *OnsTopicListRequestTag {
	s.Key = &v
	return s
}

func (s *OnsTopicListRequestTag) SetValue(v string) *OnsTopicListRequestTag {
	s.Value = &v
	return s
}

func (s *OnsTopicListRequestTag) Validate() error {
	return dara.Validate(s)
}
