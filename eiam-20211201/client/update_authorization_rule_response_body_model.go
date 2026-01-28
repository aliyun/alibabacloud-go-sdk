// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateAuthorizationRuleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateAuthorizationRuleResponseBody
	GetRequestId() *string
}

type UpdateAuthorizationRuleResponseBody struct {
	// example:
	//
	// 0441BD79-92F3-53AA-8657-F8CE4A2B912A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateAuthorizationRuleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateAuthorizationRuleResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateAuthorizationRuleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateAuthorizationRuleResponseBody) SetRequestId(v string) *UpdateAuthorizationRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateAuthorizationRuleResponseBody) Validate() error {
	return dara.Validate(s)
}
