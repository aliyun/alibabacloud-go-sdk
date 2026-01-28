// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAuthorizationRuleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteAuthorizationRuleResponseBody
	GetRequestId() *string
}

type DeleteAuthorizationRuleResponseBody struct {
	// example:
	//
	// 0441BD79-92F3-53AA-8657-F8CE4A2B912A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteAuthorizationRuleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteAuthorizationRuleResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteAuthorizationRuleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteAuthorizationRuleResponseBody) SetRequestId(v string) *DeleteAuthorizationRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteAuthorizationRuleResponseBody) Validate() error {
	return dara.Validate(s)
}
