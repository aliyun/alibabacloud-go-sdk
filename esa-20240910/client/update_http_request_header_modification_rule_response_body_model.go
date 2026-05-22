// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateHttpRequestHeaderModificationRuleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateHttpRequestHeaderModificationRuleResponseBody
	GetRequestId() *string
}

type UpdateHttpRequestHeaderModificationRuleResponseBody struct {
	// Request ID.
	//
	// example:
	//
	// 2430E05E-1340-5773-B5E1-B743929F46F2
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateHttpRequestHeaderModificationRuleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateHttpRequestHeaderModificationRuleResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateHttpRequestHeaderModificationRuleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateHttpRequestHeaderModificationRuleResponseBody) SetRequestId(v string) *UpdateHttpRequestHeaderModificationRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateHttpRequestHeaderModificationRuleResponseBody) Validate() error {
	return dara.Validate(s)
}
