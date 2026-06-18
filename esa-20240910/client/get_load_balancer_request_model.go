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
	// The ID of the load balancer. This ID is returned when the load balancer is created. You can also call [ListLoadBalancers](https://help.aliyun.com/document_detail/2868897.html) to get the IDs of all load balancers in a site.
	//
	// This parameter is required.
	//
	// example:
	//
	// 99867648760****
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The ID of the site. Call [ListSites](~~ListSites~~) to get this ID.
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
