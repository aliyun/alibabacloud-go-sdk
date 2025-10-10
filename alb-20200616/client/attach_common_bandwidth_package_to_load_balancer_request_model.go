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
	// 	- **false**(default): performs a dry run and performs the actual request. If the request passes the dry run, a 2xx HTTP status code is returned and the operation is performed.
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
