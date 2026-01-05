// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListInstancesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetInstances(v []*ListInstancesResponseBodyInstances) *ListInstancesResponseBody
	GetInstances() []*ListInstancesResponseBodyInstances
	SetRequestId(v string) *ListInstancesResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *ListInstancesResponseBody
	GetTotalCount() *int32
}

type ListInstancesResponseBody struct {
	Instances []*ListInstancesResponseBodyInstances `json:"Instances,omitempty" xml:"Instances,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// A731A84D-55C9-44F7-99BB-E1CF0CF19197
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 100
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListInstancesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListInstancesResponseBody) GoString() string {
	return s.String()
}

func (s *ListInstancesResponseBody) GetInstances() []*ListInstancesResponseBodyInstances {
	return s.Instances
}

func (s *ListInstancesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListInstancesResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListInstancesResponseBody) SetInstances(v []*ListInstancesResponseBodyInstances) *ListInstancesResponseBody {
	s.Instances = v
	return s
}

func (s *ListInstancesResponseBody) SetRequestId(v string) *ListInstancesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListInstancesResponseBody) SetTotalCount(v int32) *ListInstancesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListInstancesResponseBody) Validate() error {
	if s.Instances != nil {
		for _, item := range s.Instances {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListInstancesResponseBodyInstances struct {
	// example:
	//
	// 30.0G
	Capacity *string `json:"Capacity,omitempty" xml:"Capacity,omitempty"`
	// example:
	//
	// xgboost数据集加速实例
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// 2014-10-02T15:01:23Z
	GmtCreateTime *string `json:"GmtCreateTime,omitempty" xml:"GmtCreateTime,omitempty"`
	// example:
	//
	// 2014-10-02T15:01:23Z
	GmtModifiedTime *string `json:"GmtModifiedTime,omitempty" xml:"GmtModifiedTime,omitempty"`
	// example:
	//
	// readonly
	IoType *string `json:"IoType,omitempty" xml:"IoType,omitempty"`
	// 数据集加速实例的最大挂载点个数。
	MaxEndpoint *int32 `json:"MaxEndpoint,omitempty" xml:"MaxEndpoint,omitempty"`
	// example:
	//
	// 20
	MaxSlot *int32 `json:"MaxSlot,omitempty" xml:"MaxSlot,omitempty"`
	// example:
	//
	// acc_instance_1
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// 1557702098194904
	OwnerId *string `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// example:
	//
	// PayAsYouGo
	PaymentType *string `json:"PaymentType,omitempty" xml:"PaymentType,omitempty"`
	// 数据集加速实例的资源提供者类型。
	//
	// example:
	//
	// Ecs
	ProviderType *string         `json:"ProviderType,omitempty" xml:"ProviderType,omitempty"`
	Status       *InstanceStatus `json:"Status,omitempty" xml:"Status,omitempty"`
	// 数据集加速实例的存储类型。
	//
	// example:
	//
	// OSS
	StorageType *string                                   `json:"StorageType,omitempty" xml:"StorageType,omitempty"`
	Tags        []*ListInstancesResponseBodyInstancesTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	// example:
	//
	// Basic
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// example:
	//
	// 276065346797410278
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	// example:
	//
	// inst-my1tk3jggooi5uwks5
	Uuid *string `json:"Uuid,omitempty" xml:"Uuid,omitempty"`
}

func (s ListInstancesResponseBodyInstances) String() string {
	return dara.Prettify(s)
}

func (s ListInstancesResponseBodyInstances) GoString() string {
	return s.String()
}

func (s *ListInstancesResponseBodyInstances) GetCapacity() *string {
	return s.Capacity
}

func (s *ListInstancesResponseBodyInstances) GetDescription() *string {
	return s.Description
}

func (s *ListInstancesResponseBodyInstances) GetGmtCreateTime() *string {
	return s.GmtCreateTime
}

func (s *ListInstancesResponseBodyInstances) GetGmtModifiedTime() *string {
	return s.GmtModifiedTime
}

func (s *ListInstancesResponseBodyInstances) GetIoType() *string {
	return s.IoType
}

func (s *ListInstancesResponseBodyInstances) GetMaxEndpoint() *int32 {
	return s.MaxEndpoint
}

func (s *ListInstancesResponseBodyInstances) GetMaxSlot() *int32 {
	return s.MaxSlot
}

func (s *ListInstancesResponseBodyInstances) GetName() *string {
	return s.Name
}

func (s *ListInstancesResponseBodyInstances) GetOwnerId() *string {
	return s.OwnerId
}

func (s *ListInstancesResponseBodyInstances) GetPaymentType() *string {
	return s.PaymentType
}

func (s *ListInstancesResponseBodyInstances) GetProviderType() *string {
	return s.ProviderType
}

func (s *ListInstancesResponseBodyInstances) GetStatus() *InstanceStatus {
	return s.Status
}

func (s *ListInstancesResponseBodyInstances) GetStorageType() *string {
	return s.StorageType
}

func (s *ListInstancesResponseBodyInstances) GetTags() []*ListInstancesResponseBodyInstancesTags {
	return s.Tags
}

func (s *ListInstancesResponseBodyInstances) GetType() *string {
	return s.Type
}

func (s *ListInstancesResponseBodyInstances) GetUserId() *string {
	return s.UserId
}

func (s *ListInstancesResponseBodyInstances) GetUuid() *string {
	return s.Uuid
}

func (s *ListInstancesResponseBodyInstances) SetCapacity(v string) *ListInstancesResponseBodyInstances {
	s.Capacity = &v
	return s
}

func (s *ListInstancesResponseBodyInstances) SetDescription(v string) *ListInstancesResponseBodyInstances {
	s.Description = &v
	return s
}

func (s *ListInstancesResponseBodyInstances) SetGmtCreateTime(v string) *ListInstancesResponseBodyInstances {
	s.GmtCreateTime = &v
	return s
}

func (s *ListInstancesResponseBodyInstances) SetGmtModifiedTime(v string) *ListInstancesResponseBodyInstances {
	s.GmtModifiedTime = &v
	return s
}

func (s *ListInstancesResponseBodyInstances) SetIoType(v string) *ListInstancesResponseBodyInstances {
	s.IoType = &v
	return s
}

func (s *ListInstancesResponseBodyInstances) SetMaxEndpoint(v int32) *ListInstancesResponseBodyInstances {
	s.MaxEndpoint = &v
	return s
}

func (s *ListInstancesResponseBodyInstances) SetMaxSlot(v int32) *ListInstancesResponseBodyInstances {
	s.MaxSlot = &v
	return s
}

func (s *ListInstancesResponseBodyInstances) SetName(v string) *ListInstancesResponseBodyInstances {
	s.Name = &v
	return s
}

func (s *ListInstancesResponseBodyInstances) SetOwnerId(v string) *ListInstancesResponseBodyInstances {
	s.OwnerId = &v
	return s
}

func (s *ListInstancesResponseBodyInstances) SetPaymentType(v string) *ListInstancesResponseBodyInstances {
	s.PaymentType = &v
	return s
}

func (s *ListInstancesResponseBodyInstances) SetProviderType(v string) *ListInstancesResponseBodyInstances {
	s.ProviderType = &v
	return s
}

func (s *ListInstancesResponseBodyInstances) SetStatus(v *InstanceStatus) *ListInstancesResponseBodyInstances {
	s.Status = v
	return s
}

func (s *ListInstancesResponseBodyInstances) SetStorageType(v string) *ListInstancesResponseBodyInstances {
	s.StorageType = &v
	return s
}

func (s *ListInstancesResponseBodyInstances) SetTags(v []*ListInstancesResponseBodyInstancesTags) *ListInstancesResponseBodyInstances {
	s.Tags = v
	return s
}

func (s *ListInstancesResponseBodyInstances) SetType(v string) *ListInstancesResponseBodyInstances {
	s.Type = &v
	return s
}

func (s *ListInstancesResponseBodyInstances) SetUserId(v string) *ListInstancesResponseBodyInstances {
	s.UserId = &v
	return s
}

func (s *ListInstancesResponseBodyInstances) SetUuid(v string) *ListInstancesResponseBodyInstances {
	s.Uuid = &v
	return s
}

func (s *ListInstancesResponseBodyInstances) Validate() error {
	if s.Status != nil {
		if err := s.Status.Validate(); err != nil {
			return err
		}
	}
	if s.Tags != nil {
		for _, item := range s.Tags {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListInstancesResponseBodyInstancesTags struct {
	// example:
	//
	// dataset_version
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// example:
	//
	// v0.1.0
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListInstancesResponseBodyInstancesTags) String() string {
	return dara.Prettify(s)
}

func (s ListInstancesResponseBodyInstancesTags) GoString() string {
	return s.String()
}

func (s *ListInstancesResponseBodyInstancesTags) GetKey() *string {
	return s.Key
}

func (s *ListInstancesResponseBodyInstancesTags) GetValue() *string {
	return s.Value
}

func (s *ListInstancesResponseBodyInstancesTags) SetKey(v string) *ListInstancesResponseBodyInstancesTags {
	s.Key = &v
	return s
}

func (s *ListInstancesResponseBodyInstancesTags) SetValue(v string) *ListInstancesResponseBodyInstancesTags {
	s.Value = &v
	return s
}

func (s *ListInstancesResponseBodyInstancesTags) Validate() error {
	return dara.Validate(s)
}
