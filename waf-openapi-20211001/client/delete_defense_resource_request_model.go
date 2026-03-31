// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDefenseResourceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *DeleteDefenseResourceRequest
	GetInstanceId() *string
	SetRegionId(v string) *DeleteDefenseResourceRequest
	GetRegionId() *string
	SetResource(v string) *DeleteDefenseResourceRequest
	GetResource() *string
	SetResourceManagerResourceGroupId(v string) *DeleteDefenseResourceRequest
	GetResourceManagerResourceGroupId() *string
}

type DeleteDefenseResourceRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// waf_v2_public_cn-9lb*******
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// zhh*****-2034.test.top-clb7
	Resource *string `json:"Resource,omitempty" xml:"Resource,omitempty"`
	// example:
	//
	// rg-acfm***q
	ResourceManagerResourceGroupId *string `json:"ResourceManagerResourceGroupId,omitempty" xml:"ResourceManagerResourceGroupId,omitempty"`
}

func (s DeleteDefenseResourceRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteDefenseResourceRequest) GoString() string {
	return s.String()
}

func (s *DeleteDefenseResourceRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DeleteDefenseResourceRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteDefenseResourceRequest) GetResource() *string {
	return s.Resource
}

func (s *DeleteDefenseResourceRequest) GetResourceManagerResourceGroupId() *string {
	return s.ResourceManagerResourceGroupId
}

func (s *DeleteDefenseResourceRequest) SetInstanceId(v string) *DeleteDefenseResourceRequest {
	s.InstanceId = &v
	return s
}

func (s *DeleteDefenseResourceRequest) SetRegionId(v string) *DeleteDefenseResourceRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteDefenseResourceRequest) SetResource(v string) *DeleteDefenseResourceRequest {
	s.Resource = &v
	return s
}

func (s *DeleteDefenseResourceRequest) SetResourceManagerResourceGroupId(v string) *DeleteDefenseResourceRequest {
	s.ResourceManagerResourceGroupId = &v
	return s
}

func (s *DeleteDefenseResourceRequest) Validate() error {
	return dara.Validate(s)
}
