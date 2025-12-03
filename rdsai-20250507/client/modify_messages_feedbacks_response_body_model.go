// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyMessagesFeedbacksResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyMessagesFeedbacksResponseBody
	GetRequestId() *string
	SetResult(v string) *ModifyMessagesFeedbacksResponseBody
	GetResult() *string
}

type ModifyMessagesFeedbacksResponseBody struct {
	// example:
	//
	// FE9C65D7-930F-57A5-A207-8C396329****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// success
	Result *string `json:"Result,omitempty" xml:"Result,omitempty"`
}

func (s ModifyMessagesFeedbacksResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyMessagesFeedbacksResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyMessagesFeedbacksResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyMessagesFeedbacksResponseBody) GetResult() *string {
	return s.Result
}

func (s *ModifyMessagesFeedbacksResponseBody) SetRequestId(v string) *ModifyMessagesFeedbacksResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyMessagesFeedbacksResponseBody) SetResult(v string) *ModifyMessagesFeedbacksResponseBody {
	s.Result = &v
	return s
}

func (s *ModifyMessagesFeedbacksResponseBody) Validate() error {
	return dara.Validate(s)
}
