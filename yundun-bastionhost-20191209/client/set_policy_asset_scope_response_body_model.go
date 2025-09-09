// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetPolicyAssetScopeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *SetPolicyAssetScopeResponseBody
	GetRequestId() *string
}

type SetPolicyAssetScopeResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 5EAB922E-F476-5DFA-9290-313C608E724B
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SetPolicyAssetScopeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SetPolicyAssetScopeResponseBody) GoString() string {
	return s.String()
}

func (s *SetPolicyAssetScopeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SetPolicyAssetScopeResponseBody) SetRequestId(v string) *SetPolicyAssetScopeResponseBody {
	s.RequestId = &v
	return s
}

func (s *SetPolicyAssetScopeResponseBody) Validate() error {
	return dara.Validate(s)
}
