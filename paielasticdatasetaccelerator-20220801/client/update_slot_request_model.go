// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateSlotRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCapacity(v string) *UpdateSlotRequest
	GetCapacity() *string
	SetDescription(v string) *UpdateSlotRequest
	GetDescription() *string
	SetLifeCycle(v *SlotLifeCycle) *UpdateSlotRequest
	GetLifeCycle() *SlotLifeCycle
	SetName(v string) *UpdateSlotRequest
	GetName() *string
	SetStorageType(v string) *UpdateSlotRequest
	GetStorageType() *string
	SetStorageUri(v string) *UpdateSlotRequest
	GetStorageUri() *string
	SetTags(v []*UpdateSlotRequestTags) *UpdateSlotRequest
	GetTags() []*UpdateSlotRequestTags
}

type UpdateSlotRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 30.0G
	Capacity *string `json:"Capacity,omitempty" xml:"Capacity,omitempty"`
	// example:
	//
	// xgboost数据集加速槽
	Description *string        `json:"Description,omitempty" xml:"Description,omitempty"`
	LifeCycle   *SlotLifeCycle `json:"LifeCycle,omitempty" xml:"LifeCycle,omitempty"`
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
	StorageUri *string                  `json:"StorageUri,omitempty" xml:"StorageUri,omitempty"`
	Tags       []*UpdateSlotRequestTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
}

func (s UpdateSlotRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateSlotRequest) GoString() string {
	return s.String()
}

func (s *UpdateSlotRequest) GetCapacity() *string {
	return s.Capacity
}

func (s *UpdateSlotRequest) GetDescription() *string {
	return s.Description
}

func (s *UpdateSlotRequest) GetLifeCycle() *SlotLifeCycle {
	return s.LifeCycle
}

func (s *UpdateSlotRequest) GetName() *string {
	return s.Name
}

func (s *UpdateSlotRequest) GetStorageType() *string {
	return s.StorageType
}

func (s *UpdateSlotRequest) GetStorageUri() *string {
	return s.StorageUri
}

func (s *UpdateSlotRequest) GetTags() []*UpdateSlotRequestTags {
	return s.Tags
}

func (s *UpdateSlotRequest) SetCapacity(v string) *UpdateSlotRequest {
	s.Capacity = &v
	return s
}

func (s *UpdateSlotRequest) SetDescription(v string) *UpdateSlotRequest {
	s.Description = &v
	return s
}

func (s *UpdateSlotRequest) SetLifeCycle(v *SlotLifeCycle) *UpdateSlotRequest {
	s.LifeCycle = v
	return s
}

func (s *UpdateSlotRequest) SetName(v string) *UpdateSlotRequest {
	s.Name = &v
	return s
}

func (s *UpdateSlotRequest) SetStorageType(v string) *UpdateSlotRequest {
	s.StorageType = &v
	return s
}

func (s *UpdateSlotRequest) SetStorageUri(v string) *UpdateSlotRequest {
	s.StorageUri = &v
	return s
}

func (s *UpdateSlotRequest) SetTags(v []*UpdateSlotRequestTags) *UpdateSlotRequest {
	s.Tags = v
	return s
}

func (s *UpdateSlotRequest) Validate() error {
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

type UpdateSlotRequestTags struct {
	// example:
	//
	// dataset_version
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// example:
	//
	// v0.1.0
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s UpdateSlotRequestTags) String() string {
	return dara.Prettify(s)
}

func (s UpdateSlotRequestTags) GoString() string {
	return s.String()
}

func (s *UpdateSlotRequestTags) GetKey() *string {
	return s.Key
}

func (s *UpdateSlotRequestTags) GetValue() *string {
	return s.Value
}

func (s *UpdateSlotRequestTags) SetKey(v string) *UpdateSlotRequestTags {
	s.Key = &v
	return s
}

func (s *UpdateSlotRequestTags) SetValue(v string) *UpdateSlotRequestTags {
	s.Value = &v
	return s
}

func (s *UpdateSlotRequestTags) Validate() error {
	return dara.Validate(s)
}
