// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateOriginRuleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateOriginRuleResponseBody
	GetRequestId() *string
}

type UpdateOriginRuleResponseBody struct {
	// Request ID.
	//
	// example:
	//
	// 0AEDAF20-4DDF-4165-8750-47FF9C1929C9
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateOriginRuleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateOriginRuleResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateOriginRuleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateOriginRuleResponseBody) SetRequestId(v string) *UpdateOriginRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateOriginRuleResponseBody) Validate() error {
	return dara.Validate(s)
}
