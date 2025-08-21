// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iConfigLayer4RulePolicyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ConfigLayer4RulePolicyResponseBody
	GetRequestId() *string
}

type ConfigLayer4RulePolicyResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// CC042262-15A3-4A49-ADF0-130968EA47BC
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ConfigLayer4RulePolicyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ConfigLayer4RulePolicyResponseBody) GoString() string {
	return s.String()
}

func (s *ConfigLayer4RulePolicyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ConfigLayer4RulePolicyResponseBody) SetRequestId(v string) *ConfigLayer4RulePolicyResponseBody {
	s.RequestId = &v
	return s
}

func (s *ConfigLayer4RulePolicyResponseBody) Validate() error {
	return dara.Validate(s)
}
