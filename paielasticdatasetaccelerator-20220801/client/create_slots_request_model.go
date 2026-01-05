// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateSlotsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDryRun(v bool) *CreateSlotsRequest
	GetDryRun() *bool
	SetSlots(v []*CreateSlotsRequestSlots) *CreateSlotsRequest
	GetSlots() []*CreateSlotsRequestSlots
}

type CreateSlotsRequest struct {
	// example:
	//
	// true
	DryRun *bool                      `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	Slots  []*CreateSlotsRequestSlots `json:"Slots,omitempty" xml:"Slots,omitempty" type:"Repeated"`
}

func (s CreateSlotsRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateSlotsRequest) GoString() string {
	return s.String()
}

func (s *CreateSlotsRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *CreateSlotsRequest) GetSlots() []*CreateSlotsRequestSlots {
	return s.Slots
}

func (s *CreateSlotsRequest) SetDryRun(v bool) *CreateSlotsRequest {
	s.DryRun = &v
	return s
}

func (s *CreateSlotsRequest) SetSlots(v []*CreateSlotsRequestSlots) *CreateSlotsRequest {
	s.Slots = v
	return s
}

func (s *CreateSlotsRequest) Validate() error {
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

type CreateSlotsRequestSlots struct {
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
	// end-my1tk3jggooi5uwks5,end-n69468yvjz8d12as7d,end-tga4omjrepklkg1onn
	EndpointIds *string `json:"EndpointIds,omitempty" xml:"EndpointIds,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// inst-my1tk3jggooi5uwks5
	InstanceId *string        `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	IoType     *string        `json:"IoType,omitempty" xml:"IoType,omitempty"`
	LifeCycle  *SlotLifeCycle `json:"LifeCycle,omitempty" xml:"LifeCycle,omitempty"`
	// example:
	//
	// slot_1
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// OSS
	StorageType *string `json:"StorageType,omitempty" xml:"StorageType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// oss://pai-vision-data-hz2.oss-cn-hangzhou-internal.aliyuncs.com/data/VOCdevkit/VOC2007/ImageSets/Main/val.txt
	StorageUri *string                        `json:"StorageUri,omitempty" xml:"StorageUri,omitempty"`
	Tags       []*CreateSlotsRequestSlotsTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
}

func (s CreateSlotsRequestSlots) String() string {
	return dara.Prettify(s)
}

func (s CreateSlotsRequestSlots) GoString() string {
	return s.String()
}

func (s *CreateSlotsRequestSlots) GetCapacity() *string {
	return s.Capacity
}

func (s *CreateSlotsRequestSlots) GetDescription() *string {
	return s.Description
}

func (s *CreateSlotsRequestSlots) GetEndpointIds() *string {
	return s.EndpointIds
}

func (s *CreateSlotsRequestSlots) GetInstanceId() *string {
	return s.InstanceId
}

func (s *CreateSlotsRequestSlots) GetIoType() *string {
	return s.IoType
}

func (s *CreateSlotsRequestSlots) GetLifeCycle() *SlotLifeCycle {
	return s.LifeCycle
}

func (s *CreateSlotsRequestSlots) GetName() *string {
	return s.Name
}

func (s *CreateSlotsRequestSlots) GetStorageType() *string {
	return s.StorageType
}

func (s *CreateSlotsRequestSlots) GetStorageUri() *string {
	return s.StorageUri
}

func (s *CreateSlotsRequestSlots) GetTags() []*CreateSlotsRequestSlotsTags {
	return s.Tags
}

func (s *CreateSlotsRequestSlots) SetCapacity(v string) *CreateSlotsRequestSlots {
	s.Capacity = &v
	return s
}

func (s *CreateSlotsRequestSlots) SetDescription(v string) *CreateSlotsRequestSlots {
	s.Description = &v
	return s
}

func (s *CreateSlotsRequestSlots) SetEndpointIds(v string) *CreateSlotsRequestSlots {
	s.EndpointIds = &v
	return s
}

func (s *CreateSlotsRequestSlots) SetInstanceId(v string) *CreateSlotsRequestSlots {
	s.InstanceId = &v
	return s
}

func (s *CreateSlotsRequestSlots) SetIoType(v string) *CreateSlotsRequestSlots {
	s.IoType = &v
	return s
}

func (s *CreateSlotsRequestSlots) SetLifeCycle(v *SlotLifeCycle) *CreateSlotsRequestSlots {
	s.LifeCycle = v
	return s
}

func (s *CreateSlotsRequestSlots) SetName(v string) *CreateSlotsRequestSlots {
	s.Name = &v
	return s
}

func (s *CreateSlotsRequestSlots) SetStorageType(v string) *CreateSlotsRequestSlots {
	s.StorageType = &v
	return s
}

func (s *CreateSlotsRequestSlots) SetStorageUri(v string) *CreateSlotsRequestSlots {
	s.StorageUri = &v
	return s
}

func (s *CreateSlotsRequestSlots) SetTags(v []*CreateSlotsRequestSlotsTags) *CreateSlotsRequestSlots {
	s.Tags = v
	return s
}

func (s *CreateSlotsRequestSlots) Validate() error {
	if s.LifeCycle != nil {
		if err := s.LifeCycle.Validate(); err != nil {
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

type CreateSlotsRequestSlotsTags struct {
	// example:
	//
	// dataset_version
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// example:
	//
	// v0.1.0
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateSlotsRequestSlotsTags) String() string {
	return dara.Prettify(s)
}

func (s CreateSlotsRequestSlotsTags) GoString() string {
	return s.String()
}

func (s *CreateSlotsRequestSlotsTags) GetKey() *string {
	return s.Key
}

func (s *CreateSlotsRequestSlotsTags) GetValue() *string {
	return s.Value
}

func (s *CreateSlotsRequestSlotsTags) SetKey(v string) *CreateSlotsRequestSlotsTags {
	s.Key = &v
	return s
}

func (s *CreateSlotsRequestSlotsTags) SetValue(v string) *CreateSlotsRequestSlotsTags {
	s.Value = &v
	return s
}

func (s *CreateSlotsRequestSlotsTags) Validate() error {
	return dara.Validate(s)
}
