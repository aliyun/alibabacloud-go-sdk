// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetPolicyUserScopeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *SetPolicyUserScopeResponseBody
	GetRequestId() *string
}

type SetPolicyUserScopeResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// EC9BF0F4-8983-491A-BC8C-1B4DD94976DE
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SetPolicyUserScopeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SetPolicyUserScopeResponseBody) GoString() string {
	return s.String()
}

func (s *SetPolicyUserScopeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SetPolicyUserScopeResponseBody) SetRequestId(v string) *SetPolicyUserScopeResponseBody {
	s.RequestId = &v
	return s
}

func (s *SetPolicyUserScopeResponseBody) Validate() error {
	return dara.Validate(s)
}
