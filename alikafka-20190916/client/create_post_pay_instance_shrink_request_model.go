// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreatePostPayInstanceShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDeployType(v int32) *CreatePostPayInstanceShrinkRequest
	GetDeployType() *int32
	SetDiskSize(v int32) *CreatePostPayInstanceShrinkRequest
	GetDiskSize() *int32
	SetDiskType(v string) *CreatePostPayInstanceShrinkRequest
	GetDiskType() *string
	SetEipMax(v int32) *CreatePostPayInstanceShrinkRequest
	GetEipMax() *int32
	SetIoMaxSpec(v string) *CreatePostPayInstanceShrinkRequest
	GetIoMaxSpec() *string
	SetPaidType(v int32) *CreatePostPayInstanceShrinkRequest
	GetPaidType() *int32
	SetPartitionNum(v int32) *CreatePostPayInstanceShrinkRequest
	GetPartitionNum() *int32
	SetRegionId(v string) *CreatePostPayInstanceShrinkRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *CreatePostPayInstanceShrinkRequest
	GetResourceGroupId() *string
	SetServerlessConfigShrink(v string) *CreatePostPayInstanceShrinkRequest
	GetServerlessConfigShrink() *string
	SetSpecType(v string) *CreatePostPayInstanceShrinkRequest
	GetSpecType() *string
	SetTag(v []*CreatePostPayInstanceShrinkRequestTag) *CreatePostPayInstanceShrinkRequest
	GetTag() []*CreatePostPayInstanceShrinkRequestTag
}

type CreatePostPayInstanceShrinkRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 4
	DeployType *int32 `json:"DeployType,omitempty" xml:"DeployType,omitempty"`
	// example:
	//
	// 1500
	DiskSize *int32 `json:"DiskSize,omitempty" xml:"DiskSize,omitempty"`
	// example:
	//
	// 0
	DiskType *string `json:"DiskType,omitempty" xml:"DiskType,omitempty"`
	// example:
	//
	// 3
	EipMax *int32 `json:"EipMax,omitempty" xml:"EipMax,omitempty"`
	// example:
	//
	// alikafka.hw.2xlarge
	IoMaxSpec *string `json:"IoMaxSpec,omitempty" xml:"IoMaxSpec,omitempty"`
	// example:
	//
	// 0
	PaidType *int32 `json:"PaidType,omitempty" xml:"PaidType,omitempty"`
	// example:
	//
	// 100
	PartitionNum *int32 `json:"PartitionNum,omitempty" xml:"PartitionNum,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// rg-ac***********7q
	ResourceGroupId        *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ServerlessConfigShrink *string `json:"ServerlessConfig,omitempty" xml:"ServerlessConfig,omitempty"`
	// example:
	//
	// professional
	SpecType *string                                  `json:"SpecType,omitempty" xml:"SpecType,omitempty"`
	Tag      []*CreatePostPayInstanceShrinkRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s CreatePostPayInstanceShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreatePostPayInstanceShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreatePostPayInstanceShrinkRequest) GetDeployType() *int32 {
	return s.DeployType
}

func (s *CreatePostPayInstanceShrinkRequest) GetDiskSize() *int32 {
	return s.DiskSize
}

func (s *CreatePostPayInstanceShrinkRequest) GetDiskType() *string {
	return s.DiskType
}

func (s *CreatePostPayInstanceShrinkRequest) GetEipMax() *int32 {
	return s.EipMax
}

func (s *CreatePostPayInstanceShrinkRequest) GetIoMaxSpec() *string {
	return s.IoMaxSpec
}

func (s *CreatePostPayInstanceShrinkRequest) GetPaidType() *int32 {
	return s.PaidType
}

func (s *CreatePostPayInstanceShrinkRequest) GetPartitionNum() *int32 {
	return s.PartitionNum
}

func (s *CreatePostPayInstanceShrinkRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreatePostPayInstanceShrinkRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *CreatePostPayInstanceShrinkRequest) GetServerlessConfigShrink() *string {
	return s.ServerlessConfigShrink
}

func (s *CreatePostPayInstanceShrinkRequest) GetSpecType() *string {
	return s.SpecType
}

func (s *CreatePostPayInstanceShrinkRequest) GetTag() []*CreatePostPayInstanceShrinkRequestTag {
	return s.Tag
}

func (s *CreatePostPayInstanceShrinkRequest) SetDeployType(v int32) *CreatePostPayInstanceShrinkRequest {
	s.DeployType = &v
	return s
}

func (s *CreatePostPayInstanceShrinkRequest) SetDiskSize(v int32) *CreatePostPayInstanceShrinkRequest {
	s.DiskSize = &v
	return s
}

func (s *CreatePostPayInstanceShrinkRequest) SetDiskType(v string) *CreatePostPayInstanceShrinkRequest {
	s.DiskType = &v
	return s
}

func (s *CreatePostPayInstanceShrinkRequest) SetEipMax(v int32) *CreatePostPayInstanceShrinkRequest {
	s.EipMax = &v
	return s
}

func (s *CreatePostPayInstanceShrinkRequest) SetIoMaxSpec(v string) *CreatePostPayInstanceShrinkRequest {
	s.IoMaxSpec = &v
	return s
}

func (s *CreatePostPayInstanceShrinkRequest) SetPaidType(v int32) *CreatePostPayInstanceShrinkRequest {
	s.PaidType = &v
	return s
}

func (s *CreatePostPayInstanceShrinkRequest) SetPartitionNum(v int32) *CreatePostPayInstanceShrinkRequest {
	s.PartitionNum = &v
	return s
}

func (s *CreatePostPayInstanceShrinkRequest) SetRegionId(v string) *CreatePostPayInstanceShrinkRequest {
	s.RegionId = &v
	return s
}

func (s *CreatePostPayInstanceShrinkRequest) SetResourceGroupId(v string) *CreatePostPayInstanceShrinkRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *CreatePostPayInstanceShrinkRequest) SetServerlessConfigShrink(v string) *CreatePostPayInstanceShrinkRequest {
	s.ServerlessConfigShrink = &v
	return s
}

func (s *CreatePostPayInstanceShrinkRequest) SetSpecType(v string) *CreatePostPayInstanceShrinkRequest {
	s.SpecType = &v
	return s
}

func (s *CreatePostPayInstanceShrinkRequest) SetTag(v []*CreatePostPayInstanceShrinkRequestTag) *CreatePostPayInstanceShrinkRequest {
	s.Tag = v
	return s
}

func (s *CreatePostPayInstanceShrinkRequest) Validate() error {
	if s.Tag != nil {
		for _, item := range s.Tag {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type CreatePostPayInstanceShrinkRequestTag struct {
	// This parameter is required.
	//
	// example:
	//
	// FinanceDept
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// example:
	//
	// test
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreatePostPayInstanceShrinkRequestTag) String() string {
	return dara.Prettify(s)
}

func (s CreatePostPayInstanceShrinkRequestTag) GoString() string {
	return s.String()
}

func (s *CreatePostPayInstanceShrinkRequestTag) GetKey() *string {
	return s.Key
}

func (s *CreatePostPayInstanceShrinkRequestTag) GetValue() *string {
	return s.Value
}

func (s *CreatePostPayInstanceShrinkRequestTag) SetKey(v string) *CreatePostPayInstanceShrinkRequestTag {
	s.Key = &v
	return s
}

func (s *CreatePostPayInstanceShrinkRequestTag) SetValue(v string) *CreatePostPayInstanceShrinkRequestTag {
	s.Value = &v
	return s
}

func (s *CreatePostPayInstanceShrinkRequestTag) Validate() error {
	return dara.Validate(s)
}
