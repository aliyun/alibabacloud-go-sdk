// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStopLoadBalancerListenerResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *StopLoadBalancerListenerResponseBody
	GetRequestId() *string
}

type StopLoadBalancerListenerResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// CEF72CEB-54B6-4AE8-B225-F876FF7BA984
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s StopLoadBalancerListenerResponseBody) String() string {
	return dara.Prettify(s)
}

func (s StopLoadBalancerListenerResponseBody) GoString() string {
	return s.String()
}

func (s *StopLoadBalancerListenerResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *StopLoadBalancerListenerResponseBody) SetRequestId(v string) *StopLoadBalancerListenerResponseBody {
	s.RequestId = &v
	return s
}

func (s *StopLoadBalancerListenerResponseBody) Validate() error {
	return dara.Validate(s)
}
