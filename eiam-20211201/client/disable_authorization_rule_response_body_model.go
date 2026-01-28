// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDisableAuthorizationRuleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DisableAuthorizationRuleResponseBody
	GetRequestId() *string
}

type DisableAuthorizationRuleResponseBody struct {
	// example:
	//
	// 0441BD79-92F3-53AA-8657-F8CE4A2B912A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DisableAuthorizationRuleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DisableAuthorizationRuleResponseBody) GoString() string {
	return s.String()
}

func (s *DisableAuthorizationRuleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DisableAuthorizationRuleResponseBody) SetRequestId(v string) *DisableAuthorizationRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *DisableAuthorizationRuleResponseBody) Validate() error {
	return dara.Validate(s)
}
