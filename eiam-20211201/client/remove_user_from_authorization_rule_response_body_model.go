// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemoveUserFromAuthorizationRuleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *RemoveUserFromAuthorizationRuleResponseBody
	GetRequestId() *string
}

type RemoveUserFromAuthorizationRuleResponseBody struct {
	// example:
	//
	// 0441BD79-92F3-53AA-8657-F8CE4A2B912A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RemoveUserFromAuthorizationRuleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RemoveUserFromAuthorizationRuleResponseBody) GoString() string {
	return s.String()
}

func (s *RemoveUserFromAuthorizationRuleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RemoveUserFromAuthorizationRuleResponseBody) SetRequestId(v string) *RemoveUserFromAuthorizationRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *RemoveUserFromAuthorizationRuleResponseBody) Validate() error {
	return dara.Validate(s)
}
