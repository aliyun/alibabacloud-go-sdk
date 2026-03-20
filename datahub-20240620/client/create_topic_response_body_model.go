// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateTopicResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CreateTopicResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreateTopicResponseBody
	GetSuccess() *bool
	SetTopicName(v string) *CreateTopicResponseBody
	GetTopicName() *string
}

type CreateTopicResponseBody struct {
	// example:
	//
	// 2026020210000022c53d0b06900170
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// example:
	//
	// test01
	TopicName *string `json:"TopicName,omitempty" xml:"TopicName,omitempty"`
}

func (s CreateTopicResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateTopicResponseBody) GoString() string {
	return s.String()
}

func (s *CreateTopicResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateTopicResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateTopicResponseBody) GetTopicName() *string {
	return s.TopicName
}

func (s *CreateTopicResponseBody) SetRequestId(v string) *CreateTopicResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateTopicResponseBody) SetSuccess(v bool) *CreateTopicResponseBody {
	s.Success = &v
	return s
}

func (s *CreateTopicResponseBody) SetTopicName(v string) *CreateTopicResponseBody {
	s.TopicName = &v
	return s
}

func (s *CreateTopicResponseBody) Validate() error {
	return dara.Validate(s)
}
