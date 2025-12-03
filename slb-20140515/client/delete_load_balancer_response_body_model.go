// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteLoadBalancerResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteLoadBalancerResponseBody
	GetRequestId() *string
}

type DeleteLoadBalancerResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// CEF72CEB-54B6-4AE8-B225-F876FF7BA984
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteLoadBalancerResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteLoadBalancerResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteLoadBalancerResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteLoadBalancerResponseBody) SetRequestId(v string) *DeleteLoadBalancerResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteLoadBalancerResponseBody) Validate() error {
	return dara.Validate(s)
}
