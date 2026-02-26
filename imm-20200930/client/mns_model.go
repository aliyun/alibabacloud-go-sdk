// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMNS interface {
	dara.Model
	String() string
	GoString() string
	SetTopicName(v string) *MNS
	GetTopicName() *string
}

type MNS struct {
	// The SMQ topic. You can check topics within a region in the [SMQ console](https://mns.console.aliyun.com/). This parameter is required if you want to use SMQ for notifications.
	//
	// example:
	//
	// topic1
	TopicName *string `json:"TopicName,omitempty" xml:"TopicName,omitempty"`
}

func (s MNS) String() string {
	return dara.Prettify(s)
}

func (s MNS) GoString() string {
	return s.String()
}

func (s *MNS) GetTopicName() *string {
	return s.TopicName
}

func (s *MNS) SetTopicName(v string) *MNS {
	s.TopicName = &v
	return s
}

func (s *MNS) Validate() error {
	return dara.Validate(s)
}
