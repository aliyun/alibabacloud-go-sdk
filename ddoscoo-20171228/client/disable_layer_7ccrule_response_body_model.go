// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDisableLayer7CCRuleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DisableLayer7CCRuleResponseBody
	GetRequestId() *string
}

type DisableLayer7CCRuleResponseBody struct {
	// example:
	//
	// CF33B4C3-196E-4015-AADD-5CAD00057B80
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DisableLayer7CCRuleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DisableLayer7CCRuleResponseBody) GoString() string {
	return s.String()
}

func (s *DisableLayer7CCRuleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DisableLayer7CCRuleResponseBody) SetRequestId(v string) *DisableLayer7CCRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *DisableLayer7CCRuleResponseBody) Validate() error {
	return dara.Validate(s)
}
