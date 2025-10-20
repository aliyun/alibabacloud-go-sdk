// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDisableBrandRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBrandId(v string) *DisableBrandRequest
	GetBrandId() *string
	SetInstanceId(v string) *DisableBrandRequest
	GetInstanceId() *string
}

type DisableBrandRequest struct {
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

func (s DisableBrandRequest) String() string {
	return dara.Prettify(s)
}

func (s DisableBrandRequest) GoString() string {
	return s.String()
}

func (s *DisableBrandRequest) GetBrandId() *string {
	return s.BrandId
}

func (s *DisableBrandRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DisableBrandRequest) SetBrandId(v string) *DisableBrandRequest {
	s.BrandId = &v
	return s
}

func (s *DisableBrandRequest) SetInstanceId(v string) *DisableBrandRequest {
	s.InstanceId = &v
	return s
}

func (s *DisableBrandRequest) Validate() error {
	return dara.Validate(s)
}
