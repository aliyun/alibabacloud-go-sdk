// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartLoadBalancerListenerResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *StartLoadBalancerListenerResponseBody
	GetRequestId() *string
}

type StartLoadBalancerListenerResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3C83E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s StartLoadBalancerListenerResponseBody) String() string {
	return dara.Prettify(s)
}

func (s StartLoadBalancerListenerResponseBody) GoString() string {
	return s.String()
}

func (s *StartLoadBalancerListenerResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *StartLoadBalancerListenerResponseBody) SetRequestId(v string) *StartLoadBalancerListenerResponseBody {
	s.RequestId = &v
	return s
}

func (s *StartLoadBalancerListenerResponseBody) Validate() error {
	return dara.Validate(s)
}
