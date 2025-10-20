// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteBrandRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBrandId(v string) *DeleteBrandRequest
	GetBrandId() *string
	SetInstanceId(v string) *DeleteBrandRequest
	GetInstanceId() *string
}

type DeleteBrandRequest struct {
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

func (s DeleteBrandRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteBrandRequest) GoString() string {
	return s.String()
}

func (s *DeleteBrandRequest) GetBrandId() *string {
	return s.BrandId
}

func (s *DeleteBrandRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DeleteBrandRequest) SetBrandId(v string) *DeleteBrandRequest {
	s.BrandId = &v
	return s
}

func (s *DeleteBrandRequest) SetInstanceId(v string) *DeleteBrandRequest {
	s.InstanceId = &v
	return s
}

func (s *DeleteBrandRequest) Validate() error {
	return dara.Validate(s)
}
