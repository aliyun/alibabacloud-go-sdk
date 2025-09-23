// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iConfigLayer4RuleAttributeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ConfigLayer4RuleAttributeResponseBody
	GetRequestId() *string
}

type ConfigLayer4RuleAttributeResponseBody struct {
	// example:
	//
	// CF33B4C3-196E-4015-AADD-5CAD00057B80
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ConfigLayer4RuleAttributeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ConfigLayer4RuleAttributeResponseBody) GoString() string {
	return s.String()
}

func (s *ConfigLayer4RuleAttributeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ConfigLayer4RuleAttributeResponseBody) SetRequestId(v string) *ConfigLayer4RuleAttributeResponseBody {
	s.RequestId = &v
	return s
}

func (s *ConfigLayer4RuleAttributeResponseBody) Validate() error {
	return dara.Validate(s)
}
