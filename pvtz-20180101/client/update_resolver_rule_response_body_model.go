// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateResolverRuleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateResolverRuleResponseBody
	GetRequestId() *string
}

type UpdateResolverRuleResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 0C9959BE-3A6A-4803-8DCE-973B42ACD599
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateResolverRuleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateResolverRuleResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateResolverRuleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateResolverRuleResponseBody) SetRequestId(v string) *UpdateResolverRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateResolverRuleResponseBody) Validate() error {
	return dara.Validate(s)
}
