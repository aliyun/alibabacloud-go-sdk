// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeInstanceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCapacity(v string) *DescribeInstanceResponseBody
	GetCapacity() *string
	SetDescription(v string) *DescribeInstanceResponseBody
	GetDescription() *string
	SetGmtCreateTime(v string) *DescribeInstanceResponseBody
	GetGmtCreateTime() *string
	SetGmtModifiedTime(v string) *DescribeInstanceResponseBody
	GetGmtModifiedTime() *string
	SetIoType(v string) *DescribeInstanceResponseBody
	GetIoType() *string
	SetMaxEndpoint(v int32) *DescribeInstanceResponseBody
	GetMaxEndpoint() *int32
	SetMaxSlot(v int32) *DescribeInstanceResponseBody
	GetMaxSlot() *int32
	SetName(v string) *DescribeInstanceResponseBody
	GetName() *string
	SetOwnerId(v string) *DescribeInstanceResponseBody
	GetOwnerId() *string
	SetPaymentType(v string) *DescribeInstanceResponseBody
	GetPaymentType() *string
	SetProviderType(v string) *DescribeInstanceResponseBody
	GetProviderType() *string
	SetRequestId(v string) *DescribeInstanceResponseBody
	GetRequestId() *string
	SetStatus(v *InstanceStatus) *DescribeInstanceResponseBody
	GetStatus() *InstanceStatus
	SetStorageType(v string) *DescribeInstanceResponseBody
	GetStorageType() *string
	SetTags(v []*DescribeInstanceResponseBodyTags) *DescribeInstanceResponseBody
	GetTags() []*DescribeInstanceResponseBodyTags
	SetType(v string) *DescribeInstanceResponseBody
	GetType() *string
	SetUserId(v string) *DescribeInstanceResponseBody
	GetUserId() *string
	SetUuid(v string) *DescribeInstanceResponseBody
	GetUuid() *string
}

type DescribeInstanceResponseBody struct {
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
	ProviderType *string `json:"ProviderType,omitempty" xml:"ProviderType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// A731A84D-55C9-44F7-99BB-E1CF0CF19197
	RequestId *string         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Status    *InstanceStatus `json:"Status,omitempty" xml:"Status,omitempty"`
	// 数据集加速实例的存储类型。
	//
	// example:
	//
	// OSS
	StorageType *string                             `json:"StorageType,omitempty" xml:"StorageType,omitempty"`
	Tags        []*DescribeInstanceResponseBodyTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	// example:
	//
	// basic
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

func (s DescribeInstanceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeInstanceResponseBody) GetCapacity() *string {
	return s.Capacity
}

func (s *DescribeInstanceResponseBody) GetDescription() *string {
	return s.Description
}

func (s *DescribeInstanceResponseBody) GetGmtCreateTime() *string {
	return s.GmtCreateTime
}

func (s *DescribeInstanceResponseBody) GetGmtModifiedTime() *string {
	return s.GmtModifiedTime
}

func (s *DescribeInstanceResponseBody) GetIoType() *string {
	return s.IoType
}

func (s *DescribeInstanceResponseBody) GetMaxEndpoint() *int32 {
	return s.MaxEndpoint
}

func (s *DescribeInstanceResponseBody) GetMaxSlot() *int32 {
	return s.MaxSlot
}

func (s *DescribeInstanceResponseBody) GetName() *string {
	return s.Name
}

func (s *DescribeInstanceResponseBody) GetOwnerId() *string {
	return s.OwnerId
}

func (s *DescribeInstanceResponseBody) GetPaymentType() *string {
	return s.PaymentType
}

func (s *DescribeInstanceResponseBody) GetProviderType() *string {
	return s.ProviderType
}

func (s *DescribeInstanceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeInstanceResponseBody) GetStatus() *InstanceStatus {
	return s.Status
}

func (s *DescribeInstanceResponseBody) GetStorageType() *string {
	return s.StorageType
}

func (s *DescribeInstanceResponseBody) GetTags() []*DescribeInstanceResponseBodyTags {
	return s.Tags
}

func (s *DescribeInstanceResponseBody) GetType() *string {
	return s.Type
}

func (s *DescribeInstanceResponseBody) GetUserId() *string {
	return s.UserId
}

func (s *DescribeInstanceResponseBody) GetUuid() *string {
	return s.Uuid
}

func (s *DescribeInstanceResponseBody) SetCapacity(v string) *DescribeInstanceResponseBody {
	s.Capacity = &v
	return s
}

func (s *DescribeInstanceResponseBody) SetDescription(v string) *DescribeInstanceResponseBody {
	s.Description = &v
	return s
}

func (s *DescribeInstanceResponseBody) SetGmtCreateTime(v string) *DescribeInstanceResponseBody {
	s.GmtCreateTime = &v
	return s
}

func (s *DescribeInstanceResponseBody) SetGmtModifiedTime(v string) *DescribeInstanceResponseBody {
	s.GmtModifiedTime = &v
	return s
}

func (s *DescribeInstanceResponseBody) SetIoType(v string) *DescribeInstanceResponseBody {
	s.IoType = &v
	return s
}

func (s *DescribeInstanceResponseBody) SetMaxEndpoint(v int32) *DescribeInstanceResponseBody {
	s.MaxEndpoint = &v
	return s
}

func (s *DescribeInstanceResponseBody) SetMaxSlot(v int32) *DescribeInstanceResponseBody {
	s.MaxSlot = &v
	return s
}

func (s *DescribeInstanceResponseBody) SetName(v string) *DescribeInstanceResponseBody {
	s.Name = &v
	return s
}

func (s *DescribeInstanceResponseBody) SetOwnerId(v string) *DescribeInstanceResponseBody {
	s.OwnerId = &v
	return s
}

func (s *DescribeInstanceResponseBody) SetPaymentType(v string) *DescribeInstanceResponseBody {
	s.PaymentType = &v
	return s
}

func (s *DescribeInstanceResponseBody) SetProviderType(v string) *DescribeInstanceResponseBody {
	s.ProviderType = &v
	return s
}

func (s *DescribeInstanceResponseBody) SetRequestId(v string) *DescribeInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeInstanceResponseBody) SetStatus(v *InstanceStatus) *DescribeInstanceResponseBody {
	s.Status = v
	return s
}

func (s *DescribeInstanceResponseBody) SetStorageType(v string) *DescribeInstanceResponseBody {
	s.StorageType = &v
	return s
}

func (s *DescribeInstanceResponseBody) SetTags(v []*DescribeInstanceResponseBodyTags) *DescribeInstanceResponseBody {
	s.Tags = v
	return s
}

func (s *DescribeInstanceResponseBody) SetType(v string) *DescribeInstanceResponseBody {
	s.Type = &v
	return s
}

func (s *DescribeInstanceResponseBody) SetUserId(v string) *DescribeInstanceResponseBody {
	s.UserId = &v
	return s
}

func (s *DescribeInstanceResponseBody) SetUuid(v string) *DescribeInstanceResponseBody {
	s.Uuid = &v
	return s
}

func (s *DescribeInstanceResponseBody) Validate() error {
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

type DescribeInstanceResponseBodyTags struct {
	// example:
	//
	// dataset_version
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// example:
	//
	// v0.1.0
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeInstanceResponseBodyTags) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstanceResponseBodyTags) GoString() string {
	return s.String()
}

func (s *DescribeInstanceResponseBodyTags) GetKey() *string {
	return s.Key
}

func (s *DescribeInstanceResponseBodyTags) GetValue() *string {
	return s.Value
}

func (s *DescribeInstanceResponseBodyTags) SetKey(v string) *DescribeInstanceResponseBodyTags {
	s.Key = &v
	return s
}

func (s *DescribeInstanceResponseBodyTags) SetValue(v string) *DescribeInstanceResponseBodyTags {
	s.Value = &v
	return s
}

func (s *DescribeInstanceResponseBodyTags) Validate() error {
	return dara.Validate(s)
}
