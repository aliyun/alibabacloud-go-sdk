// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRefreshAdvisorResourceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetProduct(v string) *RefreshAdvisorResourceRequest
	GetProduct() *string
	SetResourceId(v string) *RefreshAdvisorResourceRequest
	GetResourceId() *string
}

type RefreshAdvisorResourceRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// ecs
	Product *string `json:"Product,omitempty" xml:"Product,omitempty"`
	// example:
	//
	// i-bp67acfmxazb4p****
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
}

func (s RefreshAdvisorResourceRequest) String() string {
	return dara.Prettify(s)
}

func (s RefreshAdvisorResourceRequest) GoString() string {
	return s.String()
}

func (s *RefreshAdvisorResourceRequest) GetProduct() *string {
	return s.Product
}

func (s *RefreshAdvisorResourceRequest) GetResourceId() *string {
	return s.ResourceId
}

func (s *RefreshAdvisorResourceRequest) SetProduct(v string) *RefreshAdvisorResourceRequest {
	s.Product = &v
	return s
}

func (s *RefreshAdvisorResourceRequest) SetResourceId(v string) *RefreshAdvisorResourceRequest {
	s.ResourceId = &v
	return s
}

func (s *RefreshAdvisorResourceRequest) Validate() error {
	return dara.Validate(s)
}
