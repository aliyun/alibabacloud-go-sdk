// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetLoadBalancerNameResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *SetLoadBalancerNameResponseBody
	GetRequestId() *string
}

type SetLoadBalancerNameResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// CEF72CEB-54B6-4AE8-B225-F876FF7BA984
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SetLoadBalancerNameResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SetLoadBalancerNameResponseBody) GoString() string {
	return s.String()
}

func (s *SetLoadBalancerNameResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SetLoadBalancerNameResponseBody) SetRequestId(v string) *SetLoadBalancerNameResponseBody {
	s.RequestId = &v
	return s
}

func (s *SetLoadBalancerNameResponseBody) Validate() error {
	return dara.Validate(s)
}
