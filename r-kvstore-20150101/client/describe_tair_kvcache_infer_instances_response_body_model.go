// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeTairKVCacheInferInstancesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetInstances(v *DescribeTairKVCacheInferInstancesResponseBodyInstances) *DescribeTairKVCacheInferInstancesResponseBody
	GetInstances() *DescribeTairKVCacheInferInstancesResponseBodyInstances
	SetPageNumber(v int32) *DescribeTairKVCacheInferInstancesResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeTairKVCacheInferInstancesResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *DescribeTairKVCacheInferInstancesResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *DescribeTairKVCacheInferInstancesResponseBody
	GetTotalCount() *int32
}

type DescribeTairKVCacheInferInstancesResponseBody struct {
	Instances  *DescribeTairKVCacheInferInstancesResponseBodyInstances `json:"Instances,omitempty" xml:"Instances,omitempty" type:"Struct"`
	PageNumber *int32                                                  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32                                                  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId  *string                                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount *int32                                                  `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeTairKVCacheInferInstancesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeTairKVCacheInferInstancesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeTairKVCacheInferInstancesResponseBody) GetInstances() *DescribeTairKVCacheInferInstancesResponseBodyInstances {
	return s.Instances
}

func (s *DescribeTairKVCacheInferInstancesResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeTairKVCacheInferInstancesResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeTairKVCacheInferInstancesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeTairKVCacheInferInstancesResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeTairKVCacheInferInstancesResponseBody) SetInstances(v *DescribeTairKVCacheInferInstancesResponseBodyInstances) *DescribeTairKVCacheInferInstancesResponseBody {
	s.Instances = v
	return s
}

func (s *DescribeTairKVCacheInferInstancesResponseBody) SetPageNumber(v int32) *DescribeTairKVCacheInferInstancesResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeTairKVCacheInferInstancesResponseBody) SetPageSize(v int32) *DescribeTairKVCacheInferInstancesResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeTairKVCacheInferInstancesResponseBody) SetRequestId(v string) *DescribeTairKVCacheInferInstancesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeTairKVCacheInferInstancesResponseBody) SetTotalCount(v int32) *DescribeTairKVCacheInferInstancesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeTairKVCacheInferInstancesResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeTairKVCacheInferInstancesResponseBodyInstances struct {
	TairInferInstanceDTO []*DescribeTairKVCacheInferInstancesResponseBodyInstancesTairInferInstanceDTO `json:"TairInferInstanceDTO,omitempty" xml:"TairInferInstanceDTO,omitempty" type:"Repeated"`
}

func (s DescribeTairKVCacheInferInstancesResponseBodyInstances) String() string {
	return dara.Prettify(s)
}

func (s DescribeTairKVCacheInferInstancesResponseBodyInstances) GoString() string {
	return s.String()
}

func (s *DescribeTairKVCacheInferInstancesResponseBodyInstances) GetTairInferInstanceDTO() []*DescribeTairKVCacheInferInstancesResponseBodyInstancesTairInferInstanceDTO {
	return s.TairInferInstanceDTO
}

func (s *DescribeTairKVCacheInferInstancesResponseBodyInstances) SetTairInferInstanceDTO(v []*DescribeTairKVCacheInferInstancesResponseBodyInstancesTairInferInstanceDTO) *DescribeTairKVCacheInferInstancesResponseBodyInstances {
	s.TairInferInstanceDTO = v
	return s
}

func (s *DescribeTairKVCacheInferInstancesResponseBodyInstances) Validate() error {
	return dara.Validate(s)
}

type DescribeTairKVCacheInferInstancesResponseBodyInstancesTairInferInstanceDTO struct {
	AckId           *string                                                                         `json:"AckId,omitempty" xml:"AckId,omitempty"`
	Capacity        *int64                                                                          `json:"Capacity,omitempty" xml:"Capacity,omitempty"`
	ChargeType      *string                                                                         `json:"ChargeType,omitempty" xml:"ChargeType,omitempty"`
	ComputeUnitNum  *int32                                                                          `json:"ComputeUnitNum,omitempty" xml:"ComputeUnitNum,omitempty"`
	CreateTime      *string                                                                         `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	DestroyTime     *string                                                                         `json:"DestroyTime,omitempty" xml:"DestroyTime,omitempty"`
	EndTime         *string                                                                         `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	InstanceClass   *string                                                                         `json:"InstanceClass,omitempty" xml:"InstanceClass,omitempty"`
	InstanceId      *string                                                                         `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	InstanceName    *string                                                                         `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	InstanceStatus  *string                                                                         `json:"InstanceStatus,omitempty" xml:"InstanceStatus,omitempty"`
	InstanceType    *string                                                                         `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	Model           *string                                                                         `json:"Model,omitempty" xml:"Model,omitempty"`
	ModelServiceNum *int32                                                                          `json:"ModelServiceNum,omitempty" xml:"ModelServiceNum,omitempty"`
	NetworkType     *string                                                                         `json:"NetworkType,omitempty" xml:"NetworkType,omitempty"`
	PrivateIp       *string                                                                         `json:"PrivateIp,omitempty" xml:"PrivateIp,omitempty"`
	RegionId        *string                                                                         `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceGroupId *string                                                                         `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	Tags            *DescribeTairKVCacheInferInstancesResponseBodyInstancesTairInferInstanceDTOTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Struct"`
	VNodeCount      *int32                                                                          `json:"VNodeCount,omitempty" xml:"VNodeCount,omitempty"`
	VNodeName       *string                                                                         `json:"VNodeName,omitempty" xml:"VNodeName,omitempty"`
	VSwitchId       *string                                                                         `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	VpcId           *string                                                                         `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	ZoneId          *string                                                                         `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s DescribeTairKVCacheInferInstancesResponseBodyInstancesTairInferInstanceDTO) String() string {
	return dara.Prettify(s)
}

func (s DescribeTairKVCacheInferInstancesResponseBodyInstancesTairInferInstanceDTO) GoString() string {
	return s.String()
}

func (s *DescribeTairKVCacheInferInstancesResponseBodyInstancesTairInferInstanceDTO) GetAckId() *string {
	return s.AckId
}

func (s *DescribeTairKVCacheInferInstancesResponseBodyInstancesTairInferInstanceDTO) GetCapacity() *int64 {
	return s.Capacity
}

func (s *DescribeTairKVCacheInferInstancesResponseBodyInstancesTairInferInstanceDTO) GetChargeType() *string {
	return s.ChargeType
}

func (s *DescribeTairKVCacheInferInstancesResponseBodyInstancesTairInferInstanceDTO) GetComputeUnitNum() *int32 {
	return s.ComputeUnitNum
}

func (s *DescribeTairKVCacheInferInstancesResponseBodyInstancesTairInferInstanceDTO) GetCreateTime() *string {
	return s.CreateTime
}

func (s *DescribeTairKVCacheInferInstancesResponseBodyInstancesTairInferInstanceDTO) GetDestroyTime() *string {
	return s.DestroyTime
}

func (s *DescribeTairKVCacheInferInstancesResponseBodyInstancesTairInferInstanceDTO) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeTairKVCacheInferInstancesResponseBodyInstancesTairInferInstanceDTO) GetInstanceClass() *string {
	return s.InstanceClass
}

func (s *DescribeTairKVCacheInferInstancesResponseBodyInstancesTairInferInstanceDTO) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeTairKVCacheInferInstancesResponseBodyInstancesTairInferInstanceDTO) GetInstanceName() *string {
	return s.InstanceName
}

func (s *DescribeTairKVCacheInferInstancesResponseBodyInstancesTairInferInstanceDTO) GetInstanceStatus() *string {
	return s.InstanceStatus
}

func (s *DescribeTairKVCacheInferInstancesResponseBodyInstancesTairInferInstanceDTO) GetInstanceType() *string {
	return s.InstanceType
}

func (s *DescribeTairKVCacheInferInstancesResponseBodyInstancesTairInferInstanceDTO) GetModel() *string {
	return s.Model
}

func (s *DescribeTairKVCacheInferInstancesResponseBodyInstancesTairInferInstanceDTO) GetModelServiceNum() *int32 {
	return s.ModelServiceNum
}

func (s *DescribeTairKVCacheInferInstancesResponseBodyInstancesTairInferInstanceDTO) GetNetworkType() *string {
	return s.NetworkType
}

func (s *DescribeTairKVCacheInferInstancesResponseBodyInstancesTairInferInstanceDTO) GetPrivateIp() *string {
	return s.PrivateIp
}

func (s *DescribeTairKVCacheInferInstancesResponseBodyInstancesTairInferInstanceDTO) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeTairKVCacheInferInstancesResponseBodyInstancesTairInferInstanceDTO) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeTairKVCacheInferInstancesResponseBodyInstancesTairInferInstanceDTO) GetTags() *DescribeTairKVCacheInferInstancesResponseBodyInstancesTairInferInstanceDTOTags {
	return s.Tags
}

func (s *DescribeTairKVCacheInferInstancesResponseBodyInstancesTairInferInstanceDTO) GetVNodeCount() *int32 {
	return s.VNodeCount
}

func (s *DescribeTairKVCacheInferInstancesResponseBodyInstancesTairInferInstanceDTO) GetVNodeName() *string {
	return s.VNodeName
}

func (s *DescribeTairKVCacheInferInstancesResponseBodyInstancesTairInferInstanceDTO) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *DescribeTairKVCacheInferInstancesResponseBodyInstancesTairInferInstanceDTO) GetVpcId() *string {
	return s.VpcId
}

func (s *DescribeTairKVCacheInferInstancesResponseBodyInstancesTairInferInstanceDTO) GetZoneId() *string {
	return s.ZoneId
}

func (s *DescribeTairKVCacheInferInstancesResponseBodyInstancesTairInferInstanceDTO) SetAckId(v string) *DescribeTairKVCacheInferInstancesResponseBodyInstancesTairInferInstanceDTO {
	s.AckId = &v
	return s
}

func (s *DescribeTairKVCacheInferInstancesResponseBodyInstancesTairInferInstanceDTO) SetCapacity(v int64) *DescribeTairKVCacheInferInstancesResponseBodyInstancesTairInferInstanceDTO {
	s.Capacity = &v
	return s
}

func (s *DescribeTairKVCacheInferInstancesResponseBodyInstancesTairInferInstanceDTO) SetChargeType(v string) *DescribeTairKVCacheInferInstancesResponseBodyInstancesTairInferInstanceDTO {
	s.ChargeType = &v
	return s
}

func (s *DescribeTairKVCacheInferInstancesResponseBodyInstancesTairInferInstanceDTO) SetComputeUnitNum(v int32) *DescribeTairKVCacheInferInstancesResponseBodyInstancesTairInferInstanceDTO {
	s.ComputeUnitNum = &v
	return s
}

func (s *DescribeTairKVCacheInferInstancesResponseBodyInstancesTairInferInstanceDTO) SetCreateTime(v string) *DescribeTairKVCacheInferInstancesResponseBodyInstancesTairInferInstanceDTO {
	s.CreateTime = &v
	return s
}

func (s *DescribeTairKVCacheInferInstancesResponseBodyInstancesTairInferInstanceDTO) SetDestroyTime(v string) *DescribeTairKVCacheInferInstancesResponseBodyInstancesTairInferInstanceDTO {
	s.DestroyTime = &v
	return s
}

func (s *DescribeTairKVCacheInferInstancesResponseBodyInstancesTairInferInstanceDTO) SetEndTime(v string) *DescribeTairKVCacheInferInstancesResponseBodyInstancesTairInferInstanceDTO {
	s.EndTime = &v
	return s
}

func (s *DescribeTairKVCacheInferInstancesResponseBodyInstancesTairInferInstanceDTO) SetInstanceClass(v string) *DescribeTairKVCacheInferInstancesResponseBodyInstancesTairInferInstanceDTO {
	s.InstanceClass = &v
	return s
}

func (s *DescribeTairKVCacheInferInstancesResponseBodyInstancesTairInferInstanceDTO) SetInstanceId(v string) *DescribeTairKVCacheInferInstancesResponseBodyInstancesTairInferInstanceDTO {
	s.InstanceId = &v
	return s
}

func (s *DescribeTairKVCacheInferInstancesResponseBodyInstancesTairInferInstanceDTO) SetInstanceName(v string) *DescribeTairKVCacheInferInstancesResponseBodyInstancesTairInferInstanceDTO {
	s.InstanceName = &v
	return s
}

func (s *DescribeTairKVCacheInferInstancesResponseBodyInstancesTairInferInstanceDTO) SetInstanceStatus(v string) *DescribeTairKVCacheInferInstancesResponseBodyInstancesTairInferInstanceDTO {
	s.InstanceStatus = &v
	return s
}

func (s *DescribeTairKVCacheInferInstancesResponseBodyInstancesTairInferInstanceDTO) SetInstanceType(v string) *DescribeTairKVCacheInferInstancesResponseBodyInstancesTairInferInstanceDTO {
	s.InstanceType = &v
	return s
}

func (s *DescribeTairKVCacheInferInstancesResponseBodyInstancesTairInferInstanceDTO) SetModel(v string) *DescribeTairKVCacheInferInstancesResponseBodyInstancesTairInferInstanceDTO {
	s.Model = &v
	return s
}

func (s *DescribeTairKVCacheInferInstancesResponseBodyInstancesTairInferInstanceDTO) SetModelServiceNum(v int32) *DescribeTairKVCacheInferInstancesResponseBodyInstancesTairInferInstanceDTO {
	s.ModelServiceNum = &v
	return s
}

func (s *DescribeTairKVCacheInferInstancesResponseBodyInstancesTairInferInstanceDTO) SetNetworkType(v string) *DescribeTairKVCacheInferInstancesResponseBodyInstancesTairInferInstanceDTO {
	s.NetworkType = &v
	return s
}

func (s *DescribeTairKVCacheInferInstancesResponseBodyInstancesTairInferInstanceDTO) SetPrivateIp(v string) *DescribeTairKVCacheInferInstancesResponseBodyInstancesTairInferInstanceDTO {
	s.PrivateIp = &v
	return s
}

func (s *DescribeTairKVCacheInferInstancesResponseBodyInstancesTairInferInstanceDTO) SetRegionId(v string) *DescribeTairKVCacheInferInstancesResponseBodyInstancesTairInferInstanceDTO {
	s.RegionId = &v
	return s
}

func (s *DescribeTairKVCacheInferInstancesResponseBodyInstancesTairInferInstanceDTO) SetResourceGroupId(v string) *DescribeTairKVCacheInferInstancesResponseBodyInstancesTairInferInstanceDTO {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeTairKVCacheInferInstancesResponseBodyInstancesTairInferInstanceDTO) SetTags(v *DescribeTairKVCacheInferInstancesResponseBodyInstancesTairInferInstanceDTOTags) *DescribeTairKVCacheInferInstancesResponseBodyInstancesTairInferInstanceDTO {
	s.Tags = v
	return s
}

func (s *DescribeTairKVCacheInferInstancesResponseBodyInstancesTairInferInstanceDTO) SetVNodeCount(v int32) *DescribeTairKVCacheInferInstancesResponseBodyInstancesTairInferInstanceDTO {
	s.VNodeCount = &v
	return s
}

func (s *DescribeTairKVCacheInferInstancesResponseBodyInstancesTairInferInstanceDTO) SetVNodeName(v string) *DescribeTairKVCacheInferInstancesResponseBodyInstancesTairInferInstanceDTO {
	s.VNodeName = &v
	return s
}

func (s *DescribeTairKVCacheInferInstancesResponseBodyInstancesTairInferInstanceDTO) SetVSwitchId(v string) *DescribeTairKVCacheInferInstancesResponseBodyInstancesTairInferInstanceDTO {
	s.VSwitchId = &v
	return s
}

func (s *DescribeTairKVCacheInferInstancesResponseBodyInstancesTairInferInstanceDTO) SetVpcId(v string) *DescribeTairKVCacheInferInstancesResponseBodyInstancesTairInferInstanceDTO {
	s.VpcId = &v
	return s
}

func (s *DescribeTairKVCacheInferInstancesResponseBodyInstancesTairInferInstanceDTO) SetZoneId(v string) *DescribeTairKVCacheInferInstancesResponseBodyInstancesTairInferInstanceDTO {
	s.ZoneId = &v
	return s
}

func (s *DescribeTairKVCacheInferInstancesResponseBodyInstancesTairInferInstanceDTO) Validate() error {
	return dara.Validate(s)
}

type DescribeTairKVCacheInferInstancesResponseBodyInstancesTairInferInstanceDTOTags struct {
	Tag []*DescribeTairKVCacheInferInstancesResponseBodyInstancesTairInferInstanceDTOTagsTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s DescribeTairKVCacheInferInstancesResponseBodyInstancesTairInferInstanceDTOTags) String() string {
	return dara.Prettify(s)
}

func (s DescribeTairKVCacheInferInstancesResponseBodyInstancesTairInferInstanceDTOTags) GoString() string {
	return s.String()
}

func (s *DescribeTairKVCacheInferInstancesResponseBodyInstancesTairInferInstanceDTOTags) GetTag() []*DescribeTairKVCacheInferInstancesResponseBodyInstancesTairInferInstanceDTOTagsTag {
	return s.Tag
}

func (s *DescribeTairKVCacheInferInstancesResponseBodyInstancesTairInferInstanceDTOTags) SetTag(v []*DescribeTairKVCacheInferInstancesResponseBodyInstancesTairInferInstanceDTOTagsTag) *DescribeTairKVCacheInferInstancesResponseBodyInstancesTairInferInstanceDTOTags {
	s.Tag = v
	return s
}

func (s *DescribeTairKVCacheInferInstancesResponseBodyInstancesTairInferInstanceDTOTags) Validate() error {
	return dara.Validate(s)
}

type DescribeTairKVCacheInferInstancesResponseBodyInstancesTairInferInstanceDTOTagsTag struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeTairKVCacheInferInstancesResponseBodyInstancesTairInferInstanceDTOTagsTag) String() string {
	return dara.Prettify(s)
}

func (s DescribeTairKVCacheInferInstancesResponseBodyInstancesTairInferInstanceDTOTagsTag) GoString() string {
	return s.String()
}

func (s *DescribeTairKVCacheInferInstancesResponseBodyInstancesTairInferInstanceDTOTagsTag) GetKey() *string {
	return s.Key
}

func (s *DescribeTairKVCacheInferInstancesResponseBodyInstancesTairInferInstanceDTOTagsTag) GetValue() *string {
	return s.Value
}

func (s *DescribeTairKVCacheInferInstancesResponseBodyInstancesTairInferInstanceDTOTagsTag) SetKey(v string) *DescribeTairKVCacheInferInstancesResponseBodyInstancesTairInferInstanceDTOTagsTag {
	s.Key = &v
	return s
}

func (s *DescribeTairKVCacheInferInstancesResponseBodyInstancesTairInferInstanceDTOTagsTag) SetValue(v string) *DescribeTairKVCacheInferInstancesResponseBodyInstancesTairInferInstanceDTOTagsTag {
	s.Value = &v
	return s
}

func (s *DescribeTairKVCacheInferInstancesResponseBodyInstancesTairInferInstanceDTOTagsTag) Validate() error {
	return dara.Validate(s)
}
