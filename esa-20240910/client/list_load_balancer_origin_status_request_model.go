// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListLoadBalancerOriginStatusRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLoadBalancerIds(v string) *ListLoadBalancerOriginStatusRequest
	GetLoadBalancerIds() *string
	SetPoolType(v string) *ListLoadBalancerOriginStatusRequest
	GetPoolType() *string
	SetSiteId(v int64) *ListLoadBalancerOriginStatusRequest
	GetSiteId() *int64
}

type ListLoadBalancerOriginStatusRequest struct {
	// Load balancer ID. When querying multiple load balancers, separate the IDs with commas. A maximum of 100 load balancer IDs can be passed at once. Load balancer IDs can be obtained by calling the [ListLoadBalancers](https://help.aliyun.com/document_detail/2868897.html) interface.
	//
	// This parameter is required.
	//
	// example:
	//
	// 99874066052****,100892832360****
	LoadBalancerIds *string `json:"LoadBalancerIds,omitempty" xml:"LoadBalancerIds,omitempty"`
	// Source address pool type. Various source address pools are configured under the load balancer, including default pools, fallback pools, and primary region pools. Only the status of origins in the default pool affects the status of the load balancer itself. Passing `default_pool` means only querying the status of origins in the default source address pool under the load balancer.
	//
	// example:
	//
	// default_pool
	PoolType *string `json:"PoolType,omitempty" xml:"PoolType,omitempty"`
	// Site ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1159101787****
	SiteId *int64 `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
}

func (s ListLoadBalancerOriginStatusRequest) String() string {
	return dara.Prettify(s)
}

func (s ListLoadBalancerOriginStatusRequest) GoString() string {
	return s.String()
}

func (s *ListLoadBalancerOriginStatusRequest) GetLoadBalancerIds() *string {
	return s.LoadBalancerIds
}

func (s *ListLoadBalancerOriginStatusRequest) GetPoolType() *string {
	return s.PoolType
}

func (s *ListLoadBalancerOriginStatusRequest) GetSiteId() *int64 {
	return s.SiteId
}

func (s *ListLoadBalancerOriginStatusRequest) SetLoadBalancerIds(v string) *ListLoadBalancerOriginStatusRequest {
	s.LoadBalancerIds = &v
	return s
}

func (s *ListLoadBalancerOriginStatusRequest) SetPoolType(v string) *ListLoadBalancerOriginStatusRequest {
	s.PoolType = &v
	return s
}

func (s *ListLoadBalancerOriginStatusRequest) SetSiteId(v int64) *ListLoadBalancerOriginStatusRequest {
	s.SiteId = &v
	return s
}

func (s *ListLoadBalancerOriginStatusRequest) Validate() error {
	return dara.Validate(s)
}
