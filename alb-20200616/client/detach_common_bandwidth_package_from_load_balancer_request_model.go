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
	// The EIP bandwidth plan ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cbwp-bp1pzf0ym72pu3y76****
	BandwidthPackageId *string `json:"BandwidthPackageId,omitempty" xml:"BandwidthPackageId,omitempty"`
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters.
	//
	// >  If you do not specify this parameter, the system automatically uses the **request ID*	- as the **client token**. The **request ID*	- may be different for each request.
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-426655440000
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies whether to perform only a dry run, without performing the actual request. Valid values:
	//
	// 	- **true**: performs only a dry run. The system checks the request for potential issues, including missing parameter values, incorrect request syntax, and service limits. If the request fails the dry run, an error message is returned. If the request passes the dry run, the `DryRunOperation` error code is returned.
	//
	// 	- **false*	- (default): performs a dry run and performs the actual request. If the request passes the dry run, a 2xx HTTP status code is returned and the operation is performed.
	//
	// example:
	//
	// true
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The ALB instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// alb-d676fho813rmu3****
	LoadBalancerId *string `json:"LoadBalancerId,omitempty" xml:"LoadBalancerId,omitempty"`
	// The region ID of the ALB instance.
	//
	// You can call the [DescribeRegions](https://help.aliyun.com/document_detail/36063.html) operation to query the most recent region list.
	//
	// This parameter is required.
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
