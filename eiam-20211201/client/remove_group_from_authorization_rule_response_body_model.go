// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemoveGroupFromAuthorizationRuleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *RemoveGroupFromAuthorizationRuleResponseBody
	GetRequestId() *string
}

type RemoveGroupFromAuthorizationRuleResponseBody struct {
	// example:
	//
	// 0441BD79-92F3-53AA-8657-F8CE4A2B912A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RemoveGroupFromAuthorizationRuleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RemoveGroupFromAuthorizationRuleResponseBody) GoString() string {
	return s.String()
}

func (s *RemoveGroupFromAuthorizationRuleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RemoveGroupFromAuthorizationRuleResponseBody) SetRequestId(v string) *RemoveGroupFromAuthorizationRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *RemoveGroupFromAuthorizationRuleResponseBody) Validate() error {
	return dara.Validate(s)
}
