// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddUserToAuthorizationRuleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *AddUserToAuthorizationRuleResponseBody
	GetRequestId() *string
}

type AddUserToAuthorizationRuleResponseBody struct {
	// example:
	//
	// 0441BD79-92F3-53AA-8657-F8CE4A2B912A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AddUserToAuthorizationRuleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AddUserToAuthorizationRuleResponseBody) GoString() string {
	return s.String()
}

func (s *AddUserToAuthorizationRuleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AddUserToAuthorizationRuleResponseBody) SetRequestId(v string) *AddUserToAuthorizationRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddUserToAuthorizationRuleResponseBody) Validate() error {
	return dara.Validate(s)
}
