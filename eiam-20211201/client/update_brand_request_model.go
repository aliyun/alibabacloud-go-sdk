// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateBrandRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBrandId(v string) *UpdateBrandRequest
	GetBrandId() *string
	SetBrandName(v string) *UpdateBrandRequest
	GetBrandName() *string
	SetInstanceId(v string) *UpdateBrandRequest
	GetInstanceId() *string
}

type UpdateBrandRequest struct {
	// 品牌化Id
	//
	// This parameter is required.
	//
	// example:
	//
	// brand_xxxx
	BrandId *string `json:"BrandId,omitempty" xml:"BrandId,omitempty"`
	// 品牌名称
	//
	// This parameter is required.
	//
	// example:
	//
	// Custom Brand
	BrandName *string `json:"BrandName,omitempty" xml:"BrandName,omitempty"`
	// IDaaS EIAM实例的ID。
	//
	// This parameter is required.
	//
	// example:
	//
	// idaas_ue2jvisn35ea5lmthk267xxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s UpdateBrandRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateBrandRequest) GoString() string {
	return s.String()
}

func (s *UpdateBrandRequest) GetBrandId() *string {
	return s.BrandId
}

func (s *UpdateBrandRequest) GetBrandName() *string {
	return s.BrandName
}

func (s *UpdateBrandRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *UpdateBrandRequest) SetBrandId(v string) *UpdateBrandRequest {
	s.BrandId = &v
	return s
}

func (s *UpdateBrandRequest) SetBrandName(v string) *UpdateBrandRequest {
	s.BrandName = &v
	return s
}

func (s *UpdateBrandRequest) SetInstanceId(v string) *UpdateBrandRequest {
	s.InstanceId = &v
	return s
}

func (s *UpdateBrandRequest) Validate() error {
	return dara.Validate(s)
}
