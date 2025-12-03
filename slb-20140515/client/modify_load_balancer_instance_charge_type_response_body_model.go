// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyLoadBalancerInstanceChargeTypeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyLoadBalancerInstanceChargeTypeResponseBody
	GetRequestId() *string
}

type ModifyLoadBalancerInstanceChargeTypeResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 365F4154-92F6-4AE4-92F8-7FF34B540710
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyLoadBalancerInstanceChargeTypeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyLoadBalancerInstanceChargeTypeResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyLoadBalancerInstanceChargeTypeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyLoadBalancerInstanceChargeTypeResponseBody) SetRequestId(v string) *ModifyLoadBalancerInstanceChargeTypeResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyLoadBalancerInstanceChargeTypeResponseBody) Validate() error {
	return dara.Validate(s)
}
