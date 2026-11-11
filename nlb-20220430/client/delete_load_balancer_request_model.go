// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteLoadBalancerRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *DeleteLoadBalancerRequest
	GetClientToken() *string
	SetDryRun(v bool) *DeleteLoadBalancerRequest
	GetDryRun() *bool
	SetLoadBalancerId(v string) *DeleteLoadBalancerRequest
	GetLoadBalancerId() *string
	SetRegionId(v string) *DeleteLoadBalancerRequest
	GetRegionId() *string
}

type DeleteLoadBalancerRequest struct {
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can generate the value of this parameter on your client to ensure that the value is unique across different requests. The client token can contain only ASCII characters.
	//
	// > If you do not specify this parameter, the system automatically uses the **RequestId*	- of the API request as the **ClientToken**. The **RequestId*	- of each API request is unique.
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-426655440000
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies whether to perform only a dry run on the request. Valid values:
	//
	// - **true**: Performs a dry run without deleting the instance. The system checks whether required parameters are specified, the request format is valid, and service limits are met. If the request fails the dry run, an error is returned. If the request passes the dry run, the `DryRunOperation` error code is returned.
	//
	// - **false*	- (default): Sends a regular request. If the request passes the dry run, an HTTP 2xx status code is returned and the operation is performed.
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
	// nlb-wb7r6dlwetvt5j****
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

func (s DeleteLoadBalancerRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteLoadBalancerRequest) GoString() string {
	return s.String()
}

func (s *DeleteLoadBalancerRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *DeleteLoadBalancerRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *DeleteLoadBalancerRequest) GetLoadBalancerId() *string {
	return s.LoadBalancerId
}

func (s *DeleteLoadBalancerRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteLoadBalancerRequest) SetClientToken(v string) *DeleteLoadBalancerRequest {
	s.ClientToken = &v
	return s
}

func (s *DeleteLoadBalancerRequest) SetDryRun(v bool) *DeleteLoadBalancerRequest {
	s.DryRun = &v
	return s
}

func (s *DeleteLoadBalancerRequest) SetLoadBalancerId(v string) *DeleteLoadBalancerRequest {
	s.LoadBalancerId = &v
	return s
}

func (s *DeleteLoadBalancerRequest) SetRegionId(v string) *DeleteLoadBalancerRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteLoadBalancerRequest) Validate() error {
	return dara.Validate(s)
}
