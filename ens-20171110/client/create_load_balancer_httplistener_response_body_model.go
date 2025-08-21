// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateLoadBalancerHTTPListenerResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CreateLoadBalancerHTTPListenerResponseBody
	GetRequestId() *string
}

type CreateLoadBalancerHTTPListenerResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3C83E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateLoadBalancerHTTPListenerResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateLoadBalancerHTTPListenerResponseBody) GoString() string {
	return s.String()
}

func (s *CreateLoadBalancerHTTPListenerResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateLoadBalancerHTTPListenerResponseBody) SetRequestId(v string) *CreateLoadBalancerHTTPListenerResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateLoadBalancerHTTPListenerResponseBody) Validate() error {
	return dara.Validate(s)
}
