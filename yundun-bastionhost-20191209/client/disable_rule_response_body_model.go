// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDisableRuleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DisableRuleResponseBody
	GetRequestId() *string
}

type DisableRuleResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 98DBE5C2-7D7A-5393-9E5A-71074336D33B
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DisableRuleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DisableRuleResponseBody) GoString() string {
	return s.String()
}

func (s *DisableRuleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DisableRuleResponseBody) SetRequestId(v string) *DisableRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *DisableRuleResponseBody) Validate() error {
	return dara.Validate(s)
}
