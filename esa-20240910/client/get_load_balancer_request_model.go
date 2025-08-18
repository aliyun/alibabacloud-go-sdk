// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetLoadBalancerRequest interface {
	dara.Model
	String() string
	GoString() string
	SetId(v int64) *GetLoadBalancerRequest
	GetId() *int64
	SetSiteId(v int64) *GetLoadBalancerRequest
	GetSiteId() *int64
}

type GetLoadBalancerRequest struct {
	// The ID of the load balancer, used to uniquely identify the load balancer to be queried. This ID is returned directly when the load balancer is created, or it can be obtained through the [ListLoadBalancers](https://help.aliyun.com/document_detail/2868897.html) interface for querying all load balancers under a site.
	//
	// This parameter is required.
	//
	// example:
	//
	// 99867648760****
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The site ID, which can be obtained by calling the [ListSites](~~ListSites~~) interface.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1159101787****
	SiteId *int64 `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
}

func (s GetLoadBalancerRequest) String() string {
	return dara.Prettify(s)
}

func (s GetLoadBalancerRequest) GoString() string {
	return s.String()
}

func (s *GetLoadBalancerRequest) GetId() *int64 {
	return s.Id
}

func (s *GetLoadBalancerRequest) GetSiteId() *int64 {
	return s.SiteId
}

func (s *GetLoadBalancerRequest) SetId(v int64) *GetLoadBalancerRequest {
	s.Id = &v
	return s
}

func (s *GetLoadBalancerRequest) SetSiteId(v int64) *GetLoadBalancerRequest {
	s.SiteId = &v
	return s
}

func (s *GetLoadBalancerRequest) Validate() error {
	return dara.Validate(s)
}
