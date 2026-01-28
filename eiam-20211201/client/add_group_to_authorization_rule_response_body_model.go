// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddGroupToAuthorizationRuleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *AddGroupToAuthorizationRuleResponseBody
	GetRequestId() *string
}

type AddGroupToAuthorizationRuleResponseBody struct {
	// example:
	//
	// 0441BD79-92F3-53AA-8657-F8CE4A2B912A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AddGroupToAuthorizationRuleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AddGroupToAuthorizationRuleResponseBody) GoString() string {
	return s.String()
}

func (s *AddGroupToAuthorizationRuleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AddGroupToAuthorizationRuleResponseBody) SetRequestId(v string) *AddGroupToAuthorizationRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddGroupToAuthorizationRuleResponseBody) Validate() error {
	return dara.Validate(s)
}
