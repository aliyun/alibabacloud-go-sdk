// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetBrandRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBrandId(v string) *GetBrandRequest
	GetBrandId() *string
	SetInstanceId(v string) *GetBrandRequest
	GetInstanceId() *string
}

type GetBrandRequest struct {
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

func (s GetBrandRequest) String() string {
	return dara.Prettify(s)
}

func (s GetBrandRequest) GoString() string {
	return s.String()
}

func (s *GetBrandRequest) GetBrandId() *string {
	return s.BrandId
}

func (s *GetBrandRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetBrandRequest) SetBrandId(v string) *GetBrandRequest {
	s.BrandId = &v
	return s
}

func (s *GetBrandRequest) SetInstanceId(v string) *GetBrandRequest {
	s.InstanceId = &v
	return s
}

func (s *GetBrandRequest) Validate() error {
	return dara.Validate(s)
}
