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
	// This parameter is required.
	LoadBalancerIds *string `json:"LoadBalancerIds,omitempty" xml:"LoadBalancerIds,omitempty"`
	PoolType        *string `json:"PoolType,omitempty" xml:"PoolType,omitempty"`
	// This parameter is required.
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
