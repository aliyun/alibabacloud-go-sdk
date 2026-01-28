// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateAuthorizationRuleDescriptionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateAuthorizationRuleDescriptionResponseBody
	GetRequestId() *string
}

type UpdateAuthorizationRuleDescriptionResponseBody struct {
	// example:
	//
	// 0441BD79-92F3-53AA-8657-F8CE4A2B912A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateAuthorizationRuleDescriptionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateAuthorizationRuleDescriptionResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateAuthorizationRuleDescriptionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateAuthorizationRuleDescriptionResponseBody) SetRequestId(v string) *UpdateAuthorizationRuleDescriptionResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateAuthorizationRuleDescriptionResponseBody) Validate() error {
	return dara.Validate(s)
}
