// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateLoadBalancerResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetId(v int64) *CreateLoadBalancerResponseBody
	GetId() *int64
	SetRequestId(v string) *CreateLoadBalancerResponseBody
	GetRequestId() *string
}

type CreateLoadBalancerResponseBody struct {
	// Load balancer ID.
	//
	// example:
	//
	// 99867648760****
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// Request ID.
	//
	// example:
	//
	// EEEBE525-F576-1196-8DAF-2D70CA3F4D2F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateLoadBalancerResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateLoadBalancerResponseBody) GoString() string {
	return s.String()
}

func (s *CreateLoadBalancerResponseBody) GetId() *int64 {
	return s.Id
}

func (s *CreateLoadBalancerResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateLoadBalancerResponseBody) SetId(v int64) *CreateLoadBalancerResponseBody {
	s.Id = &v
	return s
}

func (s *CreateLoadBalancerResponseBody) SetRequestId(v string) *CreateLoadBalancerResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateLoadBalancerResponseBody) Validate() error {
	return dara.Validate(s)
}
