// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddApplicationToAuthorizationRuleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *AddApplicationToAuthorizationRuleResponseBody
	GetRequestId() *string
}

type AddApplicationToAuthorizationRuleResponseBody struct {
	// example:
	//
	// 0441BD79-92F3-53AA-8657-F8CE4A2B912A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AddApplicationToAuthorizationRuleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AddApplicationToAuthorizationRuleResponseBody) GoString() string {
	return s.String()
}

func (s *AddApplicationToAuthorizationRuleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AddApplicationToAuthorizationRuleResponseBody) SetRequestId(v string) *AddApplicationToAuthorizationRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddApplicationToAuthorizationRuleResponseBody) Validate() error {
	return dara.Validate(s)
}
