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
	// The request ID.
	//
	// example:
	//
	// CEF72CEB-54B6-4AE8-B225-F876FF7BA984
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
