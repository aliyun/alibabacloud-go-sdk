// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateLoadBalancerHTTPSListenerResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CreateLoadBalancerHTTPSListenerResponseBody
	GetRequestId() *string
}

type CreateLoadBalancerHTTPSListenerResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3C83E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateLoadBalancerHTTPSListenerResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateLoadBalancerHTTPSListenerResponseBody) GoString() string {
	return s.String()
}

func (s *CreateLoadBalancerHTTPSListenerResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateLoadBalancerHTTPSListenerResponseBody) SetRequestId(v string) *CreateLoadBalancerHTTPSListenerResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateLoadBalancerHTTPSListenerResponseBody) Validate() error {
	return dara.Validate(s)
}
