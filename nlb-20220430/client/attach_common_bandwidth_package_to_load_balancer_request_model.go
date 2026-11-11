// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAttachCommonBandwidthPackageToLoadBalancerRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBandwidthPackageId(v string) *AttachCommonBandwidthPackageToLoadBalancerRequest
	GetBandwidthPackageId() *string
	SetClientToken(v string) *AttachCommonBandwidthPackageToLoadBalancerRequest
	GetClientToken() *string
	SetDryRun(v bool) *AttachCommonBandwidthPackageToLoadBalancerRequest
	GetDryRun() *bool
	SetLoadBalancerId(v string) *AttachCommonBandwidthPackageToLoadBalancerRequest
	GetLoadBalancerId() *string
	SetRegionId(v string) *AttachCommonBandwidthPackageToLoadBalancerRequest
	GetRegionId() *string
}

type AttachCommonBandwidthPackageToLoadBalancerRequest struct {
	// The ID of the Internet Shared Bandwidth to be associated.
	//
	// This parameter is required.
	//
	// example:
	//
	// cbwp-bp1pzf0ym72pu3y76****
	BandwidthPackageId *string `json:"BandwidthPackageId,omitempty" xml:"BandwidthPackageId,omitempty"`
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can generate the value from your client to ensure that it is unique among different requests. The client token only supports ASCII characters.
	//
	// > If you do not specify this parameter, the system automatically uses the **RequestId*	- of the API request as the **ClientToken**. The **RequestId*	- of each API request is unique.
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-426655440000
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies whether to perform a dry run. Valid values:
	//
	// - **true**: performs a dry run and does not associate the Internet Shared Bandwidth. The system checks whether required parameters are specified, the request format is valid, and service limits are met. If the request fails the check, an error is returned. If the request passes the check, the `DryRunOperation` error code is returned.
	//
	// - **false*	- (default): sends a normal request. If the request passes the check, an HTTP 2xx status code is returned and the operation is performed.
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

func (s AttachCommonBandwidthPackageToLoadBalancerRequest) String() string {
	return dara.Prettify(s)
}

func (s AttachCommonBandwidthPackageToLoadBalancerRequest) GoString() string {
	return s.String()
}

func (s *AttachCommonBandwidthPackageToLoadBalancerRequest) GetBandwidthPackageId() *string {
	return s.BandwidthPackageId
}

func (s *AttachCommonBandwidthPackageToLoadBalancerRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *AttachCommonBandwidthPackageToLoadBalancerRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *AttachCommonBandwidthPackageToLoadBalancerRequest) GetLoadBalancerId() *string {
	return s.LoadBalancerId
}

func (s *AttachCommonBandwidthPackageToLoadBalancerRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *AttachCommonBandwidthPackageToLoadBalancerRequest) SetBandwidthPackageId(v string) *AttachCommonBandwidthPackageToLoadBalancerRequest {
	s.BandwidthPackageId = &v
	return s
}

func (s *AttachCommonBandwidthPackageToLoadBalancerRequest) SetClientToken(v string) *AttachCommonBandwidthPackageToLoadBalancerRequest {
	s.ClientToken = &v
	return s
}

func (s *AttachCommonBandwidthPackageToLoadBalancerRequest) SetDryRun(v bool) *AttachCommonBandwidthPackageToLoadBalancerRequest {
	s.DryRun = &v
	return s
}

func (s *AttachCommonBandwidthPackageToLoadBalancerRequest) SetLoadBalancerId(v string) *AttachCommonBandwidthPackageToLoadBalancerRequest {
	s.LoadBalancerId = &v
	return s
}

func (s *AttachCommonBandwidthPackageToLoadBalancerRequest) SetRegionId(v string) *AttachCommonBandwidthPackageToLoadBalancerRequest {
	s.RegionId = &v
	return s
}

func (s *AttachCommonBandwidthPackageToLoadBalancerRequest) Validate() error {
	return dara.Validate(s)
}
