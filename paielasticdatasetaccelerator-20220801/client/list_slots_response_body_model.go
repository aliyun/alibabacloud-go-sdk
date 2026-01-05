// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSlotsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ListSlotsResponseBody
	GetRequestId() *string
	SetSlots(v []*ListSlotsResponseBodySlots) *ListSlotsResponseBody
	GetSlots() []*ListSlotsResponseBodySlots
	SetTotalCount(v int32) *ListSlotsResponseBody
	GetTotalCount() *int32
}

type ListSlotsResponseBody struct {
	// example:
	//
	// A731A84D-55C9-44F7-99BB-E1CF0CF19197
	RequestId *string                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Slots     []*ListSlotsResponseBodySlots `json:"Slots,omitempty" xml:"Slots,omitempty" type:"Repeated"`
	// example:
	//
	// 100
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListSlotsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListSlotsResponseBody) GoString() string {
	return s.String()
}

func (s *ListSlotsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListSlotsResponseBody) GetSlots() []*ListSlotsResponseBodySlots {
	return s.Slots
}

func (s *ListSlotsResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListSlotsResponseBody) SetRequestId(v string) *ListSlotsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListSlotsResponseBody) SetSlots(v []*ListSlotsResponseBodySlots) *ListSlotsResponseBody {
	s.Slots = v
	return s
}

func (s *ListSlotsResponseBody) SetTotalCount(v int32) *ListSlotsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListSlotsResponseBody) Validate() error {
	if s.Slots != nil {
		for _, item := range s.Slots {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListSlotsResponseBodySlots struct {
	// example:
	//
	// 30.0G
	Capacity *string `json:"Capacity,omitempty" xml:"Capacity,omitempty"`
	// example:
	//
	// xgboost数据集加速槽
	Description *string                                `json:"Description,omitempty" xml:"Description,omitempty"`
	Endpoints   []*ListSlotsResponseBodySlotsEndpoints `json:"Endpoints,omitempty" xml:"Endpoints,omitempty" type:"Repeated"`
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
	OwnerId *string     `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	Status  *SlotStatus `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// OSS
	StorageType *string `json:"StorageType,omitempty" xml:"StorageType,omitempty"`
	// example:
	//
	// oss://pai-vision-data-hz2.oss-cn-hangzhou-internal.aliyuncs.com/data/VOCdevkit/VOC2007/ImageSets/Main/val.txt
	StorageUri *string                           `json:"StorageUri,omitempty" xml:"StorageUri,omitempty"`
	Tags       []*ListSlotsResponseBodySlotsTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	// example:
	//
	// 276065346797410278
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	// example:
	//
	// slot-5zk866779me51jgu3w
	Uuid *string `json:"Uuid,omitempty" xml:"Uuid,omitempty"`
}

func (s ListSlotsResponseBodySlots) String() string {
	return dara.Prettify(s)
}

func (s ListSlotsResponseBodySlots) GoString() string {
	return s.String()
}

func (s *ListSlotsResponseBodySlots) GetCapacity() *string {
	return s.Capacity
}

func (s *ListSlotsResponseBodySlots) GetDescription() *string {
	return s.Description
}

func (s *ListSlotsResponseBodySlots) GetEndpoints() []*ListSlotsResponseBodySlotsEndpoints {
	return s.Endpoints
}

func (s *ListSlotsResponseBodySlots) GetGmtCreateTime() *string {
	return s.GmtCreateTime
}

func (s *ListSlotsResponseBodySlots) GetGmtModifiedTime() *string {
	return s.GmtModifiedTime
}

func (s *ListSlotsResponseBodySlots) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListSlotsResponseBodySlots) GetIoType() *string {
	return s.IoType
}

func (s *ListSlotsResponseBodySlots) GetLifeCycle() *SlotLifeCycle {
	return s.LifeCycle
}

func (s *ListSlotsResponseBodySlots) GetName() *string {
	return s.Name
}

func (s *ListSlotsResponseBodySlots) GetOwnerId() *string {
	return s.OwnerId
}

func (s *ListSlotsResponseBodySlots) GetStatus() *SlotStatus {
	return s.Status
}

func (s *ListSlotsResponseBodySlots) GetStorageType() *string {
	return s.StorageType
}

func (s *ListSlotsResponseBodySlots) GetStorageUri() *string {
	return s.StorageUri
}

func (s *ListSlotsResponseBodySlots) GetTags() []*ListSlotsResponseBodySlotsTags {
	return s.Tags
}

func (s *ListSlotsResponseBodySlots) GetUserId() *string {
	return s.UserId
}

func (s *ListSlotsResponseBodySlots) GetUuid() *string {
	return s.Uuid
}

func (s *ListSlotsResponseBodySlots) SetCapacity(v string) *ListSlotsResponseBodySlots {
	s.Capacity = &v
	return s
}

func (s *ListSlotsResponseBodySlots) SetDescription(v string) *ListSlotsResponseBodySlots {
	s.Description = &v
	return s
}

func (s *ListSlotsResponseBodySlots) SetEndpoints(v []*ListSlotsResponseBodySlotsEndpoints) *ListSlotsResponseBodySlots {
	s.Endpoints = v
	return s
}

func (s *ListSlotsResponseBodySlots) SetGmtCreateTime(v string) *ListSlotsResponseBodySlots {
	s.GmtCreateTime = &v
	return s
}

func (s *ListSlotsResponseBodySlots) SetGmtModifiedTime(v string) *ListSlotsResponseBodySlots {
	s.GmtModifiedTime = &v
	return s
}

func (s *ListSlotsResponseBodySlots) SetInstanceId(v string) *ListSlotsResponseBodySlots {
	s.InstanceId = &v
	return s
}

func (s *ListSlotsResponseBodySlots) SetIoType(v string) *ListSlotsResponseBodySlots {
	s.IoType = &v
	return s
}

func (s *ListSlotsResponseBodySlots) SetLifeCycle(v *SlotLifeCycle) *ListSlotsResponseBodySlots {
	s.LifeCycle = v
	return s
}

func (s *ListSlotsResponseBodySlots) SetName(v string) *ListSlotsResponseBodySlots {
	s.Name = &v
	return s
}

func (s *ListSlotsResponseBodySlots) SetOwnerId(v string) *ListSlotsResponseBodySlots {
	s.OwnerId = &v
	return s
}

func (s *ListSlotsResponseBodySlots) SetStatus(v *SlotStatus) *ListSlotsResponseBodySlots {
	s.Status = v
	return s
}

func (s *ListSlotsResponseBodySlots) SetStorageType(v string) *ListSlotsResponseBodySlots {
	s.StorageType = &v
	return s
}

func (s *ListSlotsResponseBodySlots) SetStorageUri(v string) *ListSlotsResponseBodySlots {
	s.StorageUri = &v
	return s
}

func (s *ListSlotsResponseBodySlots) SetTags(v []*ListSlotsResponseBodySlotsTags) *ListSlotsResponseBodySlots {
	s.Tags = v
	return s
}

func (s *ListSlotsResponseBodySlots) SetUserId(v string) *ListSlotsResponseBodySlots {
	s.UserId = &v
	return s
}

func (s *ListSlotsResponseBodySlots) SetUuid(v string) *ListSlotsResponseBodySlots {
	s.Uuid = &v
	return s
}

func (s *ListSlotsResponseBodySlots) Validate() error {
	if s.Endpoints != nil {
		for _, item := range s.Endpoints {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
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

type ListSlotsResponseBodySlotsEndpoints struct {
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
	// endpoint-1
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// 1557702098194904
	OwnerId *string         `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	Status  *EndpointStatus `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// VPC
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// example:
	//
	// 276065346797410278
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	// example:
	//
	// end-ivrq92qhbyrg4jctih
	Uuid *string `json:"Uuid,omitempty" xml:"Uuid,omitempty"`
	// example:
	//
	// vpc-j6co2fcdsl1q0gnuc3ym3
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	// example:
	//
	// vsw-j6cmr00qjyrft6jo2mg7g
	VswitchId *string `json:"VswitchId,omitempty" xml:"VswitchId,omitempty"`
}

func (s ListSlotsResponseBodySlotsEndpoints) String() string {
	return dara.Prettify(s)
}

func (s ListSlotsResponseBodySlotsEndpoints) GoString() string {
	return s.String()
}

func (s *ListSlotsResponseBodySlotsEndpoints) GetGmtCreateTime() *string {
	return s.GmtCreateTime
}

func (s *ListSlotsResponseBodySlotsEndpoints) GetGmtModifiedTime() *string {
	return s.GmtModifiedTime
}

func (s *ListSlotsResponseBodySlotsEndpoints) GetName() *string {
	return s.Name
}

func (s *ListSlotsResponseBodySlotsEndpoints) GetOwnerId() *string {
	return s.OwnerId
}

func (s *ListSlotsResponseBodySlotsEndpoints) GetStatus() *EndpointStatus {
	return s.Status
}

func (s *ListSlotsResponseBodySlotsEndpoints) GetType() *string {
	return s.Type
}

func (s *ListSlotsResponseBodySlotsEndpoints) GetUserId() *string {
	return s.UserId
}

func (s *ListSlotsResponseBodySlotsEndpoints) GetUuid() *string {
	return s.Uuid
}

func (s *ListSlotsResponseBodySlotsEndpoints) GetVpcId() *string {
	return s.VpcId
}

func (s *ListSlotsResponseBodySlotsEndpoints) GetVswitchId() *string {
	return s.VswitchId
}

func (s *ListSlotsResponseBodySlotsEndpoints) SetGmtCreateTime(v string) *ListSlotsResponseBodySlotsEndpoints {
	s.GmtCreateTime = &v
	return s
}

func (s *ListSlotsResponseBodySlotsEndpoints) SetGmtModifiedTime(v string) *ListSlotsResponseBodySlotsEndpoints {
	s.GmtModifiedTime = &v
	return s
}

func (s *ListSlotsResponseBodySlotsEndpoints) SetName(v string) *ListSlotsResponseBodySlotsEndpoints {
	s.Name = &v
	return s
}

func (s *ListSlotsResponseBodySlotsEndpoints) SetOwnerId(v string) *ListSlotsResponseBodySlotsEndpoints {
	s.OwnerId = &v
	return s
}

func (s *ListSlotsResponseBodySlotsEndpoints) SetStatus(v *EndpointStatus) *ListSlotsResponseBodySlotsEndpoints {
	s.Status = v
	return s
}

func (s *ListSlotsResponseBodySlotsEndpoints) SetType(v string) *ListSlotsResponseBodySlotsEndpoints {
	s.Type = &v
	return s
}

func (s *ListSlotsResponseBodySlotsEndpoints) SetUserId(v string) *ListSlotsResponseBodySlotsEndpoints {
	s.UserId = &v
	return s
}

func (s *ListSlotsResponseBodySlotsEndpoints) SetUuid(v string) *ListSlotsResponseBodySlotsEndpoints {
	s.Uuid = &v
	return s
}

func (s *ListSlotsResponseBodySlotsEndpoints) SetVpcId(v string) *ListSlotsResponseBodySlotsEndpoints {
	s.VpcId = &v
	return s
}

func (s *ListSlotsResponseBodySlotsEndpoints) SetVswitchId(v string) *ListSlotsResponseBodySlotsEndpoints {
	s.VswitchId = &v
	return s
}

func (s *ListSlotsResponseBodySlotsEndpoints) Validate() error {
	if s.Status != nil {
		if err := s.Status.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListSlotsResponseBodySlotsTags struct {
	// example:
	//
	// dataset_version
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// example:
	//
	// v0.1.0
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListSlotsResponseBodySlotsTags) String() string {
	return dara.Prettify(s)
}

func (s ListSlotsResponseBodySlotsTags) GoString() string {
	return s.String()
}

func (s *ListSlotsResponseBodySlotsTags) GetKey() *string {
	return s.Key
}

func (s *ListSlotsResponseBodySlotsTags) GetValue() *string {
	return s.Value
}

func (s *ListSlotsResponseBodySlotsTags) SetKey(v string) *ListSlotsResponseBodySlotsTags {
	s.Key = &v
	return s
}

func (s *ListSlotsResponseBodySlotsTags) SetValue(v string) *ListSlotsResponseBodySlotsTags {
	s.Value = &v
	return s
}

func (s *ListSlotsResponseBodySlotsTags) Validate() error {
	return dara.Validate(s)
}
