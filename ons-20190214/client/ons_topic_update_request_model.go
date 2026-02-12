// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOnsTopicUpdateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *OnsTopicUpdateRequest
	GetInstanceId() *string
	SetPerm(v int32) *OnsTopicUpdateRequest
	GetPerm() *int32
	SetTopic(v string) *OnsTopicUpdateRequest
	GetTopic() *string
}

type OnsTopicUpdateRequest struct {
	// The ID of the instance to which the topic belongs.
	//
	// example:
	//
	// MQ_INST_111111111111_DOxxxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The read/write mode that you want to configure for the topic. Valid values:
	//
	// 	- **6**: Both read and write operations are allowed.
	//
	// 	- **4**: Write operations are forbidden.
	//
	// 	- **2**: Read operations are forbidden.
	//
	// This parameter is required.
	//
	// example:
	//
	// 6
	Perm *int32 `json:"Perm,omitempty" xml:"Perm,omitempty"`
	// The name of the topic that you want to manage.
	//
	// This parameter is required.
	//
	// example:
	//
	// test
	Topic *string `json:"Topic,omitempty" xml:"Topic,omitempty"`
}

func (s OnsTopicUpdateRequest) String() string {
	return dara.Prettify(s)
}

func (s OnsTopicUpdateRequest) GoString() string {
	return s.String()
}

func (s *OnsTopicUpdateRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *OnsTopicUpdateRequest) GetPerm() *int32 {
	return s.Perm
}

func (s *OnsTopicUpdateRequest) GetTopic() *string {
	return s.Topic
}

func (s *OnsTopicUpdateRequest) SetInstanceId(v string) *OnsTopicUpdateRequest {
	s.InstanceId = &v
	return s
}

func (s *OnsTopicUpdateRequest) SetPerm(v int32) *OnsTopicUpdateRequest {
	s.Perm = &v
	return s
}

func (s *OnsTopicUpdateRequest) SetTopic(v string) *OnsTopicUpdateRequest {
	s.Topic = &v
	return s
}

func (s *OnsTopicUpdateRequest) Validate() error {
	return dara.Validate(s)
}
