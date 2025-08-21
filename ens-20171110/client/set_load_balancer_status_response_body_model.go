// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetLoadBalancerStatusResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *SetLoadBalancerStatusResponseBody
	GetRequestId() *string
}

type SetLoadBalancerStatusResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3C83E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SetLoadBalancerStatusResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SetLoadBalancerStatusResponseBody) GoString() string {
	return s.String()
}

func (s *SetLoadBalancerStatusResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SetLoadBalancerStatusResponseBody) SetRequestId(v string) *SetLoadBalancerStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *SetLoadBalancerStatusResponseBody) Validate() error {
	return dara.Validate(s)
}
