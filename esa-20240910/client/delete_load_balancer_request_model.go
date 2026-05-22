// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteLoadBalancerRequest interface {
	dara.Model
	String() string
	GoString() string
	SetId(v int64) *DeleteLoadBalancerRequest
	GetId() *int64
	SetSiteId(v int64) *DeleteLoadBalancerRequest
	GetSiteId() *int64
}

type DeleteLoadBalancerRequest struct {
	// The ID of the load balancer, used to uniquely identify the load balancer to be queried. This ID is returned directly upon creation of the load balancer and can also be obtained through the [ListLoadBalancers](https://help.aliyun.com/document_detail/2868897.html) interface for querying all load balancers under a site.
	//
	// This parameter is required.
	//
	// example:
	//
	// 99867648760****
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The ID of the site, which can be obtained by calling the [ListSites](~~ListSites~~) interface.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1159101787****
	SiteId *int64 `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
}

func (s DeleteLoadBalancerRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteLoadBalancerRequest) GoString() string {
	return s.String()
}

func (s *DeleteLoadBalancerRequest) GetId() *int64 {
	return s.Id
}

func (s *DeleteLoadBalancerRequest) GetSiteId() *int64 {
	return s.SiteId
}

func (s *DeleteLoadBalancerRequest) SetId(v int64) *DeleteLoadBalancerRequest {
	s.Id = &v
	return s
}

func (s *DeleteLoadBalancerRequest) SetSiteId(v int64) *DeleteLoadBalancerRequest {
	s.SiteId = &v
	return s
}

func (s *DeleteLoadBalancerRequest) Validate() error {
	return dara.Validate(s)
}
