// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyLoadBalancerAttributeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyLoadBalancerAttributeResponseBody
	GetRequestId() *string
}

type ModifyLoadBalancerAttributeResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3C83E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyLoadBalancerAttributeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyLoadBalancerAttributeResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyLoadBalancerAttributeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyLoadBalancerAttributeResponseBody) SetRequestId(v string) *ModifyLoadBalancerAttributeResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyLoadBalancerAttributeResponseBody) Validate() error {
	return dara.Validate(s)
}
