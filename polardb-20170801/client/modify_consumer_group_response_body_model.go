// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyConsumerGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetConsumerGroupId(v string) *ModifyConsumerGroupResponseBody
	GetConsumerGroupId() *string
	SetRequestId(v string) *ModifyConsumerGroupResponseBody
	GetRequestId() *string
}

type ModifyConsumerGroupResponseBody struct {
	// example:
	//
	// cg-xxxxxx
	ConsumerGroupId *string `json:"ConsumerGroupId,omitempty" xml:"ConsumerGroupId,omitempty"`
	// Id of the request
	//
	// example:
	//
	// A7E6A8FD-C50B-46B2-BA85-D8B8D3******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyConsumerGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyConsumerGroupResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyConsumerGroupResponseBody) GetConsumerGroupId() *string {
	return s.ConsumerGroupId
}

func (s *ModifyConsumerGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyConsumerGroupResponseBody) SetConsumerGroupId(v string) *ModifyConsumerGroupResponseBody {
	s.ConsumerGroupId = &v
	return s
}

func (s *ModifyConsumerGroupResponseBody) SetRequestId(v string) *ModifyConsumerGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyConsumerGroupResponseBody) Validate() error {
	return dara.Validate(s)
}
