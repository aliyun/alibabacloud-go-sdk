// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDisableLoadBalancerIpv6InternetRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *DisableLoadBalancerIpv6InternetRequest
	GetClientToken() *string
	SetDryRun(v bool) *DisableLoadBalancerIpv6InternetRequest
	GetDryRun() *bool
	SetLoadBalancerId(v string) *DisableLoadBalancerIpv6InternetRequest
	GetLoadBalancerId() *string
	SetRegionId(v string) *DisableLoadBalancerIpv6InternetRequest
	GetRegionId() *string
}

type DisableLoadBalancerIpv6InternetRequest struct {
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the value, but you must ensure that it is unique among different requests. The client token can contain only ASCII characters.
	//
	// > If you do not specify this parameter, the system automatically uses the **RequestId*	- of the API request as the **ClientToken**. The **RequestId*	- of each API request may be different.
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-426655440000
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies whether to perform a dry run. Valid values:
	//
	// - **true**: performs a dry run without changing the network type of the IPv6 instance. The system checks whether required parameters are specified, the request format is valid, and service limits are met. If the request fails the check, the corresponding error is returned. If the request passes the check, the `DryRunOperation` error code is returned.
	//
	// - **false*	- (default): sends a normal request. After the request passes the check, an HTTP 2xx status code is returned and the operation is performed.
	//
	// example:
	//
	// false
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The Network Load Balancer (NLB) instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// nlb-83ckzc8d4xlp8o****
	LoadBalancerId *string `json:"LoadBalancerId,omitempty" xml:"LoadBalancerId,omitempty"`
	// The region ID of the Network Load Balancer (NLB) instance.
	//
	// You can call the [DescribeRegions](https://help.aliyun.com/document_detail/443657.html) operation to query region IDs.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DisableLoadBalancerIpv6InternetRequest) String() string {
	return dara.Prettify(s)
}

func (s DisableLoadBalancerIpv6InternetRequest) GoString() string {
	return s.String()
}

func (s *DisableLoadBalancerIpv6InternetRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *DisableLoadBalancerIpv6InternetRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *DisableLoadBalancerIpv6InternetRequest) GetLoadBalancerId() *string {
	return s.LoadBalancerId
}

func (s *DisableLoadBalancerIpv6InternetRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DisableLoadBalancerIpv6InternetRequest) SetClientToken(v string) *DisableLoadBalancerIpv6InternetRequest {
	s.ClientToken = &v
	return s
}

func (s *DisableLoadBalancerIpv6InternetRequest) SetDryRun(v bool) *DisableLoadBalancerIpv6InternetRequest {
	s.DryRun = &v
	return s
}

func (s *DisableLoadBalancerIpv6InternetRequest) SetLoadBalancerId(v string) *DisableLoadBalancerIpv6InternetRequest {
	s.LoadBalancerId = &v
	return s
}

func (s *DisableLoadBalancerIpv6InternetRequest) SetRegionId(v string) *DisableLoadBalancerIpv6InternetRequest {
	s.RegionId = &v
	return s
}

func (s *DisableLoadBalancerIpv6InternetRequest) Validate() error {
	return dara.Validate(s)
}
