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
	// This parameter is required.
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// This parameter is required.
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
