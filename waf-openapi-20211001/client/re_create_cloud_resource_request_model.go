// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReCreateCloudResourceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *ReCreateCloudResourceRequest
	GetInstanceId() *string
	SetPort(v int32) *ReCreateCloudResourceRequest
	GetPort() *int32
	SetRegionId(v string) *ReCreateCloudResourceRequest
	GetRegionId() *string
	SetResourceInstanceId(v string) *ReCreateCloudResourceRequest
	GetResourceInstanceId() *string
	SetResourceManagerResourceGroupId(v string) *ReCreateCloudResourceRequest
	GetResourceManagerResourceGroupId() *string
	SetResourceProduct(v string) *ReCreateCloudResourceRequest
	GetResourceProduct() *string
}

type ReCreateCloudResourceRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// waf_v2_public_cn-***
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 443
	Port *int32 `json:"Port,omitempty" xml:"Port,omitempty"`
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// lb-bp1*****jqnnqk5uj2p
	ResourceInstanceId *string `json:"ResourceInstanceId,omitempty" xml:"ResourceInstanceId,omitempty"`
	// example:
	//
	// rg-acfm***q
	ResourceManagerResourceGroupId *string `json:"ResourceManagerResourceGroupId,omitempty" xml:"ResourceManagerResourceGroupId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// clb7
	ResourceProduct *string `json:"ResourceProduct,omitempty" xml:"ResourceProduct,omitempty"`
}

func (s ReCreateCloudResourceRequest) String() string {
	return dara.Prettify(s)
}

func (s ReCreateCloudResourceRequest) GoString() string {
	return s.String()
}

func (s *ReCreateCloudResourceRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ReCreateCloudResourceRequest) GetPort() *int32 {
	return s.Port
}

func (s *ReCreateCloudResourceRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ReCreateCloudResourceRequest) GetResourceInstanceId() *string {
	return s.ResourceInstanceId
}

func (s *ReCreateCloudResourceRequest) GetResourceManagerResourceGroupId() *string {
	return s.ResourceManagerResourceGroupId
}

func (s *ReCreateCloudResourceRequest) GetResourceProduct() *string {
	return s.ResourceProduct
}

func (s *ReCreateCloudResourceRequest) SetInstanceId(v string) *ReCreateCloudResourceRequest {
	s.InstanceId = &v
	return s
}

func (s *ReCreateCloudResourceRequest) SetPort(v int32) *ReCreateCloudResourceRequest {
	s.Port = &v
	return s
}

func (s *ReCreateCloudResourceRequest) SetRegionId(v string) *ReCreateCloudResourceRequest {
	s.RegionId = &v
	return s
}

func (s *ReCreateCloudResourceRequest) SetResourceInstanceId(v string) *ReCreateCloudResourceRequest {
	s.ResourceInstanceId = &v
	return s
}

func (s *ReCreateCloudResourceRequest) SetResourceManagerResourceGroupId(v string) *ReCreateCloudResourceRequest {
	s.ResourceManagerResourceGroupId = &v
	return s
}

func (s *ReCreateCloudResourceRequest) SetResourceProduct(v string) *ReCreateCloudResourceRequest {
	s.ResourceProduct = &v
	return s
}

func (s *ReCreateCloudResourceRequest) Validate() error {
	return dara.Validate(s)
}
