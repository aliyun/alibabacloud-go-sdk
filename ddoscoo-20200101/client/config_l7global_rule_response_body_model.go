// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iConfigL7GlobalRuleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ConfigL7GlobalRuleResponseBody
	GetRequestId() *string
}

type ConfigL7GlobalRuleResponseBody struct {
	// example:
	//
	// CF33B4C3-196E-4015-AADD-5CAD00057B80
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ConfigL7GlobalRuleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ConfigL7GlobalRuleResponseBody) GoString() string {
	return s.String()
}

func (s *ConfigL7GlobalRuleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ConfigL7GlobalRuleResponseBody) SetRequestId(v string) *ConfigL7GlobalRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *ConfigL7GlobalRuleResponseBody) Validate() error {
	return dara.Validate(s)
}
