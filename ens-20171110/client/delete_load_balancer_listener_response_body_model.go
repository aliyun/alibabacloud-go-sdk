// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteLoadBalancerListenerResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteLoadBalancerListenerResponseBody
	GetRequestId() *string
}

type DeleteLoadBalancerListenerResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3C83E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteLoadBalancerListenerResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteLoadBalancerListenerResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteLoadBalancerListenerResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteLoadBalancerListenerResponseBody) SetRequestId(v string) *DeleteLoadBalancerListenerResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteLoadBalancerListenerResponseBody) Validate() error {
	return dara.Validate(s)
}
