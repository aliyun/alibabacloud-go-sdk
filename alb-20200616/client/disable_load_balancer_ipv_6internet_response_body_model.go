// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDisableLoadBalancerIpv6InternetResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetJobId(v string) *DisableLoadBalancerIpv6InternetResponseBody
	GetJobId() *string
	SetRequestId(v string) *DisableLoadBalancerIpv6InternetResponseBody
	GetRequestId() *string
}

type DisableLoadBalancerIpv6InternetResponseBody struct {
	// The asynchronous task ID.
	//
	// example:
	//
	// d12871a6-ebb2-41f3-8d74-d9f452bb****
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 7D866E37-1123-5160-AFF1-BDAF5EB86A8A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DisableLoadBalancerIpv6InternetResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DisableLoadBalancerIpv6InternetResponseBody) GoString() string {
	return s.String()
}

func (s *DisableLoadBalancerIpv6InternetResponseBody) GetJobId() *string {
	return s.JobId
}

func (s *DisableLoadBalancerIpv6InternetResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DisableLoadBalancerIpv6InternetResponseBody) SetJobId(v string) *DisableLoadBalancerIpv6InternetResponseBody {
	s.JobId = &v
	return s
}

func (s *DisableLoadBalancerIpv6InternetResponseBody) SetRequestId(v string) *DisableLoadBalancerIpv6InternetResponseBody {
	s.RequestId = &v
	return s
}

func (s *DisableLoadBalancerIpv6InternetResponseBody) Validate() error {
	return dara.Validate(s)
}
