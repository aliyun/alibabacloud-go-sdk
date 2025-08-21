// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iConfigLayer4RuleBakModeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ConfigLayer4RuleBakModeResponseBody
	GetRequestId() *string
}

type ConfigLayer4RuleBakModeResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// CC042262-15A3-4A49-ADF0-130968EA47BC
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ConfigLayer4RuleBakModeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ConfigLayer4RuleBakModeResponseBody) GoString() string {
	return s.String()
}

func (s *ConfigLayer4RuleBakModeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ConfigLayer4RuleBakModeResponseBody) SetRequestId(v string) *ConfigLayer4RuleBakModeResponseBody {
	s.RequestId = &v
	return s
}

func (s *ConfigLayer4RuleBakModeResponseBody) Validate() error {
	return dara.Validate(s)
}
