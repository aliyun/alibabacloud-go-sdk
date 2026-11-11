// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDetachCommonBandwidthPackageFromLoadBalancerRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBandwidthPackageId(v string) *DetachCommonBandwidthPackageFromLoadBalancerRequest
	GetBandwidthPackageId() *string
	SetClientToken(v string) *DetachCommonBandwidthPackageFromLoadBalancerRequest
	GetClientToken() *string
	SetDryRun(v bool) *DetachCommonBandwidthPackageFromLoadBalancerRequest
	GetDryRun() *bool
	SetLoadBalancerId(v string) *DetachCommonBandwidthPackageFromLoadBalancerRequest
	GetLoadBalancerId() *string
	SetRegionId(v string) *DetachCommonBandwidthPackageFromLoadBalancerRequest
	GetRegionId() *string
}

type DetachCommonBandwidthPackageFromLoadBalancerRequest struct {
	// The ID of the Internet Shared Bandwidth plan to remove.
	//
	// This parameter is required.
	//
	// example:
	//
	// cbwp-bp1pzf0ym72pu3y76****
	BandwidthPackageId *string `json:"BandwidthPackageId,omitempty" xml:"BandwidthPackageId,omitempty"`
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can generate a value for this parameter on your client to ensure that the value is unique across different requests. The client token can contain only ASCII characters.
	//
	// > If you do not specify this parameter, the system automatically uses the **RequestId*	- of the API request as the **ClientToken**. The **RequestId*	- of each API request is different.
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-426655440000
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies whether to only precheck the request. Valid values:
	//
	// - **true**: sends a check request and does not disassociate the bandwidth plan. The system checks whether the required parameters are specified, the request format is valid, and the service limits are met. If the request fails the check, an error is returned. If the request passes the check, the `DryRunOperation` error code is returned.
	//
	// - **false*	- (default): sends a regular request. If the request passes the check, an HTTP 2xx status code is returned and the operation is performed.
	//
	// example:
	//
	// false
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The ID of the Network Load Balancer (NLB) instance.
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

func (s DetachCommonBandwidthPackageFromLoadBalancerRequest) String() string {
	return dara.Prettify(s)
}

func (s DetachCommonBandwidthPackageFromLoadBalancerRequest) GoString() string {
	return s.String()
}

func (s *DetachCommonBandwidthPackageFromLoadBalancerRequest) GetBandwidthPackageId() *string {
	return s.BandwidthPackageId
}

func (s *DetachCommonBandwidthPackageFromLoadBalancerRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *DetachCommonBandwidthPackageFromLoadBalancerRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *DetachCommonBandwidthPackageFromLoadBalancerRequest) GetLoadBalancerId() *string {
	return s.LoadBalancerId
}

func (s *DetachCommonBandwidthPackageFromLoadBalancerRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DetachCommonBandwidthPackageFromLoadBalancerRequest) SetBandwidthPackageId(v string) *DetachCommonBandwidthPackageFromLoadBalancerRequest {
	s.BandwidthPackageId = &v
	return s
}

func (s *DetachCommonBandwidthPackageFromLoadBalancerRequest) SetClientToken(v string) *DetachCommonBandwidthPackageFromLoadBalancerRequest {
	s.ClientToken = &v
	return s
}

func (s *DetachCommonBandwidthPackageFromLoadBalancerRequest) SetDryRun(v bool) *DetachCommonBandwidthPackageFromLoadBalancerRequest {
	s.DryRun = &v
	return s
}

func (s *DetachCommonBandwidthPackageFromLoadBalancerRequest) SetLoadBalancerId(v string) *DetachCommonBandwidthPackageFromLoadBalancerRequest {
	s.LoadBalancerId = &v
	return s
}

func (s *DetachCommonBandwidthPackageFromLoadBalancerRequest) SetRegionId(v string) *DetachCommonBandwidthPackageFromLoadBalancerRequest {
	s.RegionId = &v
	return s
}

func (s *DetachCommonBandwidthPackageFromLoadBalancerRequest) Validate() error {
	return dara.Validate(s)
}
