// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSlotResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCapacity(v string) *DescribeSlotResponseBody
	GetCapacity() *string
	SetDescription(v string) *DescribeSlotResponseBody
	GetDescription() *string
	SetGmtCreateTime(v string) *DescribeSlotResponseBody
	GetGmtCreateTime() *string
	SetGmtModifiedTime(v string) *DescribeSlotResponseBody
	GetGmtModifiedTime() *string
	SetInstanceId(v string) *DescribeSlotResponseBody
	GetInstanceId() *string
	SetIoType(v string) *DescribeSlotResponseBody
	GetIoType() *string
	SetLifeCycle(v *SlotLifeCycle) *DescribeSlotResponseBody
	GetLifeCycle() *SlotLifeCycle
	SetName(v string) *DescribeSlotResponseBody
	GetName() *string
	SetOwnerId(v string) *DescribeSlotResponseBody
	GetOwnerId() *string
	SetRequestId(v string) *DescribeSlotResponseBody
	GetRequestId() *string
	SetStatus(v *SlotStatus) *DescribeSlotResponseBody
	GetStatus() *SlotStatus
	SetStorageType(v string) *DescribeSlotResponseBody
	GetStorageType() *string
	SetStorageUri(v string) *DescribeSlotResponseBody
	GetStorageUri() *string
	SetTags(v []*DescribeSlotResponseBodyTags) *DescribeSlotResponseBody
	GetTags() []*DescribeSlotResponseBodyTags
	SetUserId(v string) *DescribeSlotResponseBody
	GetUserId() *string
	SetUuid(v string) *DescribeSlotResponseBody
	GetUuid() *string
}

type DescribeSlotResponseBody struct {
	// example:
	//
	// 30.0G
	Capacity *string `json:"Capacity,omitempty" xml:"Capacity,omitempty"`
	// example:
	//
	// xgboost数据集加速槽
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
	// inst-my1tk3jggooi5uwks5
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// 数据集加速槽的读写类型。
	//
	// example:
	//
	// readonly
	IoType    *string        `json:"IoType,omitempty" xml:"IoType,omitempty"`
	LifeCycle *SlotLifeCycle `json:"LifeCycle,omitempty" xml:"LifeCycle,omitempty"`
	// example:
	//
	// slot_1
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// 1557702098194904
	OwnerId *string `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// A731A84D-55C9-44F7-99BB-E1CF0CF19197
	RequestId *string     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Status    *SlotStatus `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// OSS
	StorageType *string `json:"StorageType,omitempty" xml:"StorageType,omitempty"`
	// example:
	//
	// oss://pai-vision-data-hz2.oss-cn-hangzhou-internal.aliyuncs.com/data/VOCdevkit/VOC2007/ImageSets/Main/val.txt
	StorageUri *string                         `json:"StorageUri,omitempty" xml:"StorageUri,omitempty"`
	Tags       []*DescribeSlotResponseBodyTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	// example:
	//
	// 276065346797410278
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	// example:
	//
	// slot-5zk866779me51jgu3w
	Uuid *string `json:"Uuid,omitempty" xml:"Uuid,omitempty"`
}

func (s DescribeSlotResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeSlotResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSlotResponseBody) GetCapacity() *string {
	return s.Capacity
}

func (s *DescribeSlotResponseBody) GetDescription() *string {
	return s.Description
}

func (s *DescribeSlotResponseBody) GetGmtCreateTime() *string {
	return s.GmtCreateTime
}

func (s *DescribeSlotResponseBody) GetGmtModifiedTime() *string {
	return s.GmtModifiedTime
}

func (s *DescribeSlotResponseBody) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeSlotResponseBody) GetIoType() *string {
	return s.IoType
}

func (s *DescribeSlotResponseBody) GetLifeCycle() *SlotLifeCycle {
	return s.LifeCycle
}

func (s *DescribeSlotResponseBody) GetName() *string {
	return s.Name
}

func (s *DescribeSlotResponseBody) GetOwnerId() *string {
	return s.OwnerId
}

func (s *DescribeSlotResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeSlotResponseBody) GetStatus() *SlotStatus {
	return s.Status
}

func (s *DescribeSlotResponseBody) GetStorageType() *string {
	return s.StorageType
}

func (s *DescribeSlotResponseBody) GetStorageUri() *string {
	return s.StorageUri
}

func (s *DescribeSlotResponseBody) GetTags() []*DescribeSlotResponseBodyTags {
	return s.Tags
}

func (s *DescribeSlotResponseBody) GetUserId() *string {
	return s.UserId
}

func (s *DescribeSlotResponseBody) GetUuid() *string {
	return s.Uuid
}

func (s *DescribeSlotResponseBody) SetCapacity(v string) *DescribeSlotResponseBody {
	s.Capacity = &v
	return s
}

func (s *DescribeSlotResponseBody) SetDescription(v string) *DescribeSlotResponseBody {
	s.Description = &v
	return s
}

func (s *DescribeSlotResponseBody) SetGmtCreateTime(v string) *DescribeSlotResponseBody {
	s.GmtCreateTime = &v
	return s
}

func (s *DescribeSlotResponseBody) SetGmtModifiedTime(v string) *DescribeSlotResponseBody {
	s.GmtModifiedTime = &v
	return s
}

func (s *DescribeSlotResponseBody) SetInstanceId(v string) *DescribeSlotResponseBody {
	s.InstanceId = &v
	return s
}

func (s *DescribeSlotResponseBody) SetIoType(v string) *DescribeSlotResponseBody {
	s.IoType = &v
	return s
}

func (s *DescribeSlotResponseBody) SetLifeCycle(v *SlotLifeCycle) *DescribeSlotResponseBody {
	s.LifeCycle = v
	return s
}

func (s *DescribeSlotResponseBody) SetName(v string) *DescribeSlotResponseBody {
	s.Name = &v
	return s
}

func (s *DescribeSlotResponseBody) SetOwnerId(v string) *DescribeSlotResponseBody {
	s.OwnerId = &v
	return s
}

func (s *DescribeSlotResponseBody) SetRequestId(v string) *DescribeSlotResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeSlotResponseBody) SetStatus(v *SlotStatus) *DescribeSlotResponseBody {
	s.Status = v
	return s
}

func (s *DescribeSlotResponseBody) SetStorageType(v string) *DescribeSlotResponseBody {
	s.StorageType = &v
	return s
}

func (s *DescribeSlotResponseBody) SetStorageUri(v string) *DescribeSlotResponseBody {
	s.StorageUri = &v
	return s
}

func (s *DescribeSlotResponseBody) SetTags(v []*DescribeSlotResponseBodyTags) *DescribeSlotResponseBody {
	s.Tags = v
	return s
}

func (s *DescribeSlotResponseBody) SetUserId(v string) *DescribeSlotResponseBody {
	s.UserId = &v
	return s
}

func (s *DescribeSlotResponseBody) SetUuid(v string) *DescribeSlotResponseBody {
	s.Uuid = &v
	return s
}

func (s *DescribeSlotResponseBody) Validate() error {
	if s.LifeCycle != nil {
		if err := s.LifeCycle.Validate(); err != nil {
			return err
		}
	}
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

type DescribeSlotResponseBodyTags struct {
	// example:
	//
	// dataset_version
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// example:
	//
	// v0.1.0
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeSlotResponseBodyTags) String() string {
	return dara.Prettify(s)
}

func (s DescribeSlotResponseBodyTags) GoString() string {
	return s.String()
}

func (s *DescribeSlotResponseBodyTags) GetKey() *string {
	return s.Key
}

func (s *DescribeSlotResponseBodyTags) GetValue() *string {
	return s.Value
}

func (s *DescribeSlotResponseBodyTags) SetKey(v string) *DescribeSlotResponseBodyTags {
	s.Key = &v
	return s
}

func (s *DescribeSlotResponseBodyTags) SetValue(v string) *DescribeSlotResponseBodyTags {
	s.Value = &v
	return s
}

func (s *DescribeSlotResponseBodyTags) Validate() error {
	return dara.Validate(s)
}
