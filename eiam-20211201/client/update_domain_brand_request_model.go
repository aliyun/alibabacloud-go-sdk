// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateDomainBrandRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBrandId(v string) *UpdateDomainBrandRequest
	GetBrandId() *string
	SetDomainId(v string) *UpdateDomainBrandRequest
	GetDomainId() *string
	SetInstanceId(v string) *UpdateDomainBrandRequest
	GetInstanceId() *string
}

type UpdateDomainBrandRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// brand_xxxx
	BrandId *string `json:"BrandId,omitempty" xml:"BrandId,omitempty"`
	// 域名ID。
	//
	// This parameter is required.
	//
	// example:
	//
	// dm_examplexxxxx
	DomainId *string `json:"DomainId,omitempty" xml:"DomainId,omitempty"`
	// IDaaS EIAM实例的ID。
	//
	// This parameter is required.
	//
	// example:
	//
	// idaas_ue2jvisn35ea5lmthk267xxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s UpdateDomainBrandRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateDomainBrandRequest) GoString() string {
	return s.String()
}

func (s *UpdateDomainBrandRequest) GetBrandId() *string {
	return s.BrandId
}

func (s *UpdateDomainBrandRequest) GetDomainId() *string {
	return s.DomainId
}

func (s *UpdateDomainBrandRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *UpdateDomainBrandRequest) SetBrandId(v string) *UpdateDomainBrandRequest {
	s.BrandId = &v
	return s
}

func (s *UpdateDomainBrandRequest) SetDomainId(v string) *UpdateDomainBrandRequest {
	s.DomainId = &v
	return s
}

func (s *UpdateDomainBrandRequest) SetInstanceId(v string) *UpdateDomainBrandRequest {
	s.InstanceId = &v
	return s
}

func (s *UpdateDomainBrandRequest) Validate() error {
	return dara.Validate(s)
}
