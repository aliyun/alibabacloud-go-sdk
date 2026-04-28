// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateConsumerGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetConsumerGroupId(v string) *CreateConsumerGroupResponseBody
	GetConsumerGroupId() *string
	SetRequestId(v string) *CreateConsumerGroupResponseBody
	GetRequestId() *string
}

type CreateConsumerGroupResponseBody struct {
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

func (s CreateConsumerGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateConsumerGroupResponseBody) GoString() string {
	return s.String()
}

func (s *CreateConsumerGroupResponseBody) GetConsumerGroupId() *string {
	return s.ConsumerGroupId
}

func (s *CreateConsumerGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateConsumerGroupResponseBody) SetConsumerGroupId(v string) *CreateConsumerGroupResponseBody {
	s.ConsumerGroupId = &v
	return s
}

func (s *CreateConsumerGroupResponseBody) SetRequestId(v string) *CreateConsumerGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateConsumerGroupResponseBody) Validate() error {
	return dara.Validate(s)
}
