// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDispatchRuleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteDispatchRuleResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeleteDispatchRuleResponseBody
	GetSuccess() *bool
}

type DeleteDispatchRuleResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 16AF921B-8187-489F-9913-43C808B4****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values: true and false.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteDispatchRuleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteDispatchRuleResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteDispatchRuleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteDispatchRuleResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteDispatchRuleResponseBody) SetRequestId(v string) *DeleteDispatchRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteDispatchRuleResponseBody) SetSuccess(v bool) *DeleteDispatchRuleResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteDispatchRuleResponseBody) Validate() error {
	return dara.Validate(s)
}
