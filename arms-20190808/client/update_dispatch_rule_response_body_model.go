// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateDispatchRuleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateDispatchRuleResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateDispatchRuleResponseBody
	GetSuccess() *bool
}

type UpdateDispatchRuleResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// A5EC8221-08F2-4C95-9AF1-49FD998C****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// 	- `true`: The request was successful.
	//
	// 	- `false`: The request failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateDispatchRuleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateDispatchRuleResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateDispatchRuleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateDispatchRuleResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateDispatchRuleResponseBody) SetRequestId(v string) *UpdateDispatchRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateDispatchRuleResponseBody) SetSuccess(v bool) *UpdateDispatchRuleResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateDispatchRuleResponseBody) Validate() error {
	return dara.Validate(s)
}
