// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iConfigLayer4RuleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ConfigLayer4RuleResponseBody
	GetRequestId() *string
}

type ConfigLayer4RuleResponseBody struct {
	// example:
	//
	// 0bcf28g5-d57c-11e7-9bs0-d89d6717dxbc
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ConfigLayer4RuleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ConfigLayer4RuleResponseBody) GoString() string {
	return s.String()
}

func (s *ConfigLayer4RuleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ConfigLayer4RuleResponseBody) SetRequestId(v string) *ConfigLayer4RuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *ConfigLayer4RuleResponseBody) Validate() error {
	return dara.Validate(s)
}
