// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetLoadBalancerDeleteProtectionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *SetLoadBalancerDeleteProtectionResponseBody
	GetRequestId() *string
}

type SetLoadBalancerDeleteProtectionResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 791D8B68-AE0F-4174-AF54-088C8B3C5D54
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SetLoadBalancerDeleteProtectionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SetLoadBalancerDeleteProtectionResponseBody) GoString() string {
	return s.String()
}

func (s *SetLoadBalancerDeleteProtectionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SetLoadBalancerDeleteProtectionResponseBody) SetRequestId(v string) *SetLoadBalancerDeleteProtectionResponseBody {
	s.RequestId = &v
	return s
}

func (s *SetLoadBalancerDeleteProtectionResponseBody) Validate() error {
	return dara.Validate(s)
}
