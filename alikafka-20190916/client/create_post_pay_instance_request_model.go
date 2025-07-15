// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreatePostPayInstanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDeployType(v int32) *CreatePostPayInstanceRequest
	GetDeployType() *int32
	SetDiskSize(v int32) *CreatePostPayInstanceRequest
	GetDiskSize() *int32
	SetDiskType(v string) *CreatePostPayInstanceRequest
	GetDiskType() *string
	SetEipMax(v int32) *CreatePostPayInstanceRequest
	GetEipMax() *int32
	SetIoMaxSpec(v string) *CreatePostPayInstanceRequest
	GetIoMaxSpec() *string
	SetPaidType(v int32) *CreatePostPayInstanceRequest
	GetPaidType() *int32
	SetPartitionNum(v int32) *CreatePostPayInstanceRequest
	GetPartitionNum() *int32
	SetRegionId(v string) *CreatePostPayInstanceRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *CreatePostPayInstanceRequest
	GetResourceGroupId() *string
	SetServerlessConfig(v *CreatePostPayInstanceRequestServerlessConfig) *CreatePostPayInstanceRequest
	GetServerlessConfig() *CreatePostPayInstanceRequestServerlessConfig
	SetSpecType(v string) *CreatePostPayInstanceRequest
	GetSpecType() *string
	SetTag(v []*CreatePostPayInstanceRequestTag) *CreatePostPayInstanceRequest
	GetTag() []*CreatePostPayInstanceRequestTag
}

type CreatePostPayInstanceRequest struct {
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
	ResourceGroupId  *string                                       `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ServerlessConfig *CreatePostPayInstanceRequestServerlessConfig `json:"ServerlessConfig,omitempty" xml:"ServerlessConfig,omitempty" type:"Struct"`
	// example:
	//
	// professional
	SpecType *string                            `json:"SpecType,omitempty" xml:"SpecType,omitempty"`
	Tag      []*CreatePostPayInstanceRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s CreatePostPayInstanceRequest) String() string {
	return dara.Prettify(s)
}

func (s CreatePostPayInstanceRequest) GoString() string {
	return s.String()
}

func (s *CreatePostPayInstanceRequest) GetDeployType() *int32 {
	return s.DeployType
}

func (s *CreatePostPayInstanceRequest) GetDiskSize() *int32 {
	return s.DiskSize
}

func (s *CreatePostPayInstanceRequest) GetDiskType() *string {
	return s.DiskType
}

func (s *CreatePostPayInstanceRequest) GetEipMax() *int32 {
	return s.EipMax
}

func (s *CreatePostPayInstanceRequest) GetIoMaxSpec() *string {
	return s.IoMaxSpec
}

func (s *CreatePostPayInstanceRequest) GetPaidType() *int32 {
	return s.PaidType
}

func (s *CreatePostPayInstanceRequest) GetPartitionNum() *int32 {
	return s.PartitionNum
}

func (s *CreatePostPayInstanceRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreatePostPayInstanceRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *CreatePostPayInstanceRequest) GetServerlessConfig() *CreatePostPayInstanceRequestServerlessConfig {
	return s.ServerlessConfig
}

func (s *CreatePostPayInstanceRequest) GetSpecType() *string {
	return s.SpecType
}

func (s *CreatePostPayInstanceRequest) GetTag() []*CreatePostPayInstanceRequestTag {
	return s.Tag
}

func (s *CreatePostPayInstanceRequest) SetDeployType(v int32) *CreatePostPayInstanceRequest {
	s.DeployType = &v
	return s
}

func (s *CreatePostPayInstanceRequest) SetDiskSize(v int32) *CreatePostPayInstanceRequest {
	s.DiskSize = &v
	return s
}

func (s *CreatePostPayInstanceRequest) SetDiskType(v string) *CreatePostPayInstanceRequest {
	s.DiskType = &v
	return s
}

func (s *CreatePostPayInstanceRequest) SetEipMax(v int32) *CreatePostPayInstanceRequest {
	s.EipMax = &v
	return s
}

func (s *CreatePostPayInstanceRequest) SetIoMaxSpec(v string) *CreatePostPayInstanceRequest {
	s.IoMaxSpec = &v
	return s
}

func (s *CreatePostPayInstanceRequest) SetPaidType(v int32) *CreatePostPayInstanceRequest {
	s.PaidType = &v
	return s
}

func (s *CreatePostPayInstanceRequest) SetPartitionNum(v int32) *CreatePostPayInstanceRequest {
	s.PartitionNum = &v
	return s
}

func (s *CreatePostPayInstanceRequest) SetRegionId(v string) *CreatePostPayInstanceRequest {
	s.RegionId = &v
	return s
}

func (s *CreatePostPayInstanceRequest) SetResourceGroupId(v string) *CreatePostPayInstanceRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *CreatePostPayInstanceRequest) SetServerlessConfig(v *CreatePostPayInstanceRequestServerlessConfig) *CreatePostPayInstanceRequest {
	s.ServerlessConfig = v
	return s
}

func (s *CreatePostPayInstanceRequest) SetSpecType(v string) *CreatePostPayInstanceRequest {
	s.SpecType = &v
	return s
}

func (s *CreatePostPayInstanceRequest) SetTag(v []*CreatePostPayInstanceRequestTag) *CreatePostPayInstanceRequest {
	s.Tag = v
	return s
}

func (s *CreatePostPayInstanceRequest) Validate() error {
	return dara.Validate(s)
}

type CreatePostPayInstanceRequestServerlessConfig struct {
	// example:
	//
	// 60
	ReservedPublishCapacity *int64 `json:"ReservedPublishCapacity,omitempty" xml:"ReservedPublishCapacity,omitempty"`
	// example:
	//
	// 50
	ReservedSubscribeCapacity *int64 `json:"ReservedSubscribeCapacity,omitempty" xml:"ReservedSubscribeCapacity,omitempty"`
}

func (s CreatePostPayInstanceRequestServerlessConfig) String() string {
	return dara.Prettify(s)
}

func (s CreatePostPayInstanceRequestServerlessConfig) GoString() string {
	return s.String()
}

func (s *CreatePostPayInstanceRequestServerlessConfig) GetReservedPublishCapacity() *int64 {
	return s.ReservedPublishCapacity
}

func (s *CreatePostPayInstanceRequestServerlessConfig) GetReservedSubscribeCapacity() *int64 {
	return s.ReservedSubscribeCapacity
}

func (s *CreatePostPayInstanceRequestServerlessConfig) SetReservedPublishCapacity(v int64) *CreatePostPayInstanceRequestServerlessConfig {
	s.ReservedPublishCapacity = &v
	return s
}

func (s *CreatePostPayInstanceRequestServerlessConfig) SetReservedSubscribeCapacity(v int64) *CreatePostPayInstanceRequestServerlessConfig {
	s.ReservedSubscribeCapacity = &v
	return s
}

func (s *CreatePostPayInstanceRequestServerlessConfig) Validate() error {
	return dara.Validate(s)
}

type CreatePostPayInstanceRequestTag struct {
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

func (s CreatePostPayInstanceRequestTag) String() string {
	return dara.Prettify(s)
}

func (s CreatePostPayInstanceRequestTag) GoString() string {
	return s.String()
}

func (s *CreatePostPayInstanceRequestTag) GetKey() *string {
	return s.Key
}

func (s *CreatePostPayInstanceRequestTag) GetValue() *string {
	return s.Value
}

func (s *CreatePostPayInstanceRequestTag) SetKey(v string) *CreatePostPayInstanceRequestTag {
	s.Key = &v
	return s
}

func (s *CreatePostPayInstanceRequestTag) SetValue(v string) *CreatePostPayInstanceRequestTag {
	s.Value = &v
	return s
}

func (s *CreatePostPayInstanceRequestTag) Validate() error {
	return dara.Validate(s)
}
