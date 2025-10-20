// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetLoginRedirectApplicationForBrandRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApplicationId(v string) *SetLoginRedirectApplicationForBrandRequest
	GetApplicationId() *string
	SetBrandId(v string) *SetLoginRedirectApplicationForBrandRequest
	GetBrandId() *string
	SetInstanceId(v string) *SetLoginRedirectApplicationForBrandRequest
	GetInstanceId() *string
}

type SetLoginRedirectApplicationForBrandRequest struct {
	// 应用ID
	//
	// This parameter is required.
	//
	// example:
	//
	// app_xxxx
	ApplicationId *string `json:"ApplicationId,omitempty" xml:"ApplicationId,omitempty"`
	// 品牌化Id
	//
	// This parameter is required.
	//
	// example:
	//
	// brand_xxxx
	BrandId *string `json:"BrandId,omitempty" xml:"BrandId,omitempty"`
	// IDaaS EIAM实例的ID。
	//
	// This parameter is required.
	//
	// example:
	//
	// idaas_ue2jvisn35ea5lmthk267xxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s SetLoginRedirectApplicationForBrandRequest) String() string {
	return dara.Prettify(s)
}

func (s SetLoginRedirectApplicationForBrandRequest) GoString() string {
	return s.String()
}

func (s *SetLoginRedirectApplicationForBrandRequest) GetApplicationId() *string {
	return s.ApplicationId
}

func (s *SetLoginRedirectApplicationForBrandRequest) GetBrandId() *string {
	return s.BrandId
}

func (s *SetLoginRedirectApplicationForBrandRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *SetLoginRedirectApplicationForBrandRequest) SetApplicationId(v string) *SetLoginRedirectApplicationForBrandRequest {
	s.ApplicationId = &v
	return s
}

func (s *SetLoginRedirectApplicationForBrandRequest) SetBrandId(v string) *SetLoginRedirectApplicationForBrandRequest {
	s.BrandId = &v
	return s
}

func (s *SetLoginRedirectApplicationForBrandRequest) SetInstanceId(v string) *SetLoginRedirectApplicationForBrandRequest {
	s.InstanceId = &v
	return s
}

func (s *SetLoginRedirectApplicationForBrandRequest) Validate() error {
	return dara.Validate(s)
}
