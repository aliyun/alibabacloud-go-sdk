// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateBrandRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBrandName(v string) *CreateBrandRequest
	GetBrandName() *string
	SetInstanceId(v string) *CreateBrandRequest
	GetInstanceId() *string
}

type CreateBrandRequest struct {
	// 品牌化名称
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

func (s CreateBrandRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateBrandRequest) GoString() string {
	return s.String()
}

func (s *CreateBrandRequest) GetBrandName() *string {
	return s.BrandName
}

func (s *CreateBrandRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *CreateBrandRequest) SetBrandName(v string) *CreateBrandRequest {
	s.BrandName = &v
	return s
}

func (s *CreateBrandRequest) SetInstanceId(v string) *CreateBrandRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateBrandRequest) Validate() error {
	return dara.Validate(s)
}
