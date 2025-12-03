// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetLoadBalancerModificationProtectionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *SetLoadBalancerModificationProtectionResponseBody
	GetRequestId() *string
}

type SetLoadBalancerModificationProtectionResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 791D8B68-AE0F-4174-AF54-088C8B3C5D54
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SetLoadBalancerModificationProtectionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SetLoadBalancerModificationProtectionResponseBody) GoString() string {
	return s.String()
}

func (s *SetLoadBalancerModificationProtectionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SetLoadBalancerModificationProtectionResponseBody) SetRequestId(v string) *SetLoadBalancerModificationProtectionResponseBody {
	s.RequestId = &v
	return s
}

func (s *SetLoadBalancerModificationProtectionResponseBody) Validate() error {
	return dara.Validate(s)
}
