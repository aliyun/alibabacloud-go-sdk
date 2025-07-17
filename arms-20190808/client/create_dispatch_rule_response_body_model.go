// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDispatchRuleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDispatchRuleId(v int64) *CreateDispatchRuleResponseBody
	GetDispatchRuleId() *int64
	SetRequestId(v string) *CreateDispatchRuleResponseBody
	GetRequestId() *string
}

type CreateDispatchRuleResponseBody struct {
	// The ID of the dispatch policy.
	//
	// example:
	//
	// 10413
	DispatchRuleId *int64 `json:"DispatchRuleId,omitempty" xml:"DispatchRuleId,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// A5EC8221-08F2-4C95-9AF1-49FD998C****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateDispatchRuleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateDispatchRuleResponseBody) GoString() string {
	return s.String()
}

func (s *CreateDispatchRuleResponseBody) GetDispatchRuleId() *int64 {
	return s.DispatchRuleId
}

func (s *CreateDispatchRuleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateDispatchRuleResponseBody) SetDispatchRuleId(v int64) *CreateDispatchRuleResponseBody {
	s.DispatchRuleId = &v
	return s
}

func (s *CreateDispatchRuleResponseBody) SetRequestId(v string) *CreateDispatchRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateDispatchRuleResponseBody) Validate() error {
	return dara.Validate(s)
}
