// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteTopicResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteTopicResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeleteTopicResponseBody
	GetSuccess() *bool
}

type DeleteTopicResponseBody struct {
	// example:
	//
	// 20250623101207d2a3770b026dd321
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteTopicResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteTopicResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteTopicResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteTopicResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteTopicResponseBody) SetRequestId(v string) *DeleteTopicResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteTopicResponseBody) SetSuccess(v bool) *DeleteTopicResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteTopicResponseBody) Validate() error {
	return dara.Validate(s)
}
