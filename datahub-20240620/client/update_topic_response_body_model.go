// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateTopicResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateTopicResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateTopicResponseBody
	GetSuccess() *bool
}

type UpdateTopicResponseBody struct {
	// example:
	//
	// 20260319152525d2a3770b00c232d4
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateTopicResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateTopicResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateTopicResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateTopicResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateTopicResponseBody) SetRequestId(v string) *UpdateTopicResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateTopicResponseBody) SetSuccess(v bool) *UpdateTopicResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateTopicResponseBody) Validate() error {
	return dara.Validate(s)
}
