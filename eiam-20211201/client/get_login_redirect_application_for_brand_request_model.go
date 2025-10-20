// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetLoginRedirectApplicationForBrandRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBrandId(v string) *GetLoginRedirectApplicationForBrandRequest
	GetBrandId() *string
	SetInstanceId(v string) *GetLoginRedirectApplicationForBrandRequest
	GetInstanceId() *string
}

type GetLoginRedirectApplicationForBrandRequest struct {
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

func (s GetLoginRedirectApplicationForBrandRequest) String() string {
	return dara.Prettify(s)
}

func (s GetLoginRedirectApplicationForBrandRequest) GoString() string {
	return s.String()
}

func (s *GetLoginRedirectApplicationForBrandRequest) GetBrandId() *string {
	return s.BrandId
}

func (s *GetLoginRedirectApplicationForBrandRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetLoginRedirectApplicationForBrandRequest) SetBrandId(v string) *GetLoginRedirectApplicationForBrandRequest {
	s.BrandId = &v
	return s
}

func (s *GetLoginRedirectApplicationForBrandRequest) SetInstanceId(v string) *GetLoginRedirectApplicationForBrandRequest {
	s.InstanceId = &v
	return s
}

func (s *GetLoginRedirectApplicationForBrandRequest) Validate() error {
	return dara.Validate(s)
}
