// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateSlotRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCapacity(v string) *CreateSlotRequest
	GetCapacity() *string
	SetDescription(v string) *CreateSlotRequest
	GetDescription() *string
	SetEndpointIds(v string) *CreateSlotRequest
	GetEndpointIds() *string
	SetEndpoints(v []*CreateSlotRequestEndpoints) *CreateSlotRequest
	GetEndpoints() []*CreateSlotRequestEndpoints
	SetInstanceId(v string) *CreateSlotRequest
	GetInstanceId() *string
	SetIoType(v string) *CreateSlotRequest
	GetIoType() *string
	SetLifeCycle(v *SlotLifeCycle) *CreateSlotRequest
	GetLifeCycle() *SlotLifeCycle
	SetName(v string) *CreateSlotRequest
	GetName() *string
	SetStorageType(v string) *CreateSlotRequest
	GetStorageType() *string
	SetStorageUri(v string) *CreateSlotRequest
	GetStorageUri() *string
	SetTags(v []*CreateSlotRequestTags) *CreateSlotRequest
	GetTags() []*CreateSlotRequestTags
}

type CreateSlotRequest struct {
	// This parameter is required.
	//
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
	EndpointIds *string                       `json:"EndpointIds,omitempty" xml:"EndpointIds,omitempty"`
	Endpoints   []*CreateSlotRequestEndpoints `json:"Endpoints,omitempty" xml:"Endpoints,omitempty" type:"Repeated"`
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
	StorageUri *string                  `json:"StorageUri,omitempty" xml:"StorageUri,omitempty"`
	Tags       []*CreateSlotRequestTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
}

func (s CreateSlotRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateSlotRequest) GoString() string {
	return s.String()
}

func (s *CreateSlotRequest) GetCapacity() *string {
	return s.Capacity
}

func (s *CreateSlotRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateSlotRequest) GetEndpointIds() *string {
	return s.EndpointIds
}

func (s *CreateSlotRequest) GetEndpoints() []*CreateSlotRequestEndpoints {
	return s.Endpoints
}

func (s *CreateSlotRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *CreateSlotRequest) GetIoType() *string {
	return s.IoType
}

func (s *CreateSlotRequest) GetLifeCycle() *SlotLifeCycle {
	return s.LifeCycle
}

func (s *CreateSlotRequest) GetName() *string {
	return s.Name
}

func (s *CreateSlotRequest) GetStorageType() *string {
	return s.StorageType
}

func (s *CreateSlotRequest) GetStorageUri() *string {
	return s.StorageUri
}

func (s *CreateSlotRequest) GetTags() []*CreateSlotRequestTags {
	return s.Tags
}

func (s *CreateSlotRequest) SetCapacity(v string) *CreateSlotRequest {
	s.Capacity = &v
	return s
}

func (s *CreateSlotRequest) SetDescription(v string) *CreateSlotRequest {
	s.Description = &v
	return s
}

func (s *CreateSlotRequest) SetEndpointIds(v string) *CreateSlotRequest {
	s.EndpointIds = &v
	return s
}

func (s *CreateSlotRequest) SetEndpoints(v []*CreateSlotRequestEndpoints) *CreateSlotRequest {
	s.Endpoints = v
	return s
}

func (s *CreateSlotRequest) SetInstanceId(v string) *CreateSlotRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateSlotRequest) SetIoType(v string) *CreateSlotRequest {
	s.IoType = &v
	return s
}

func (s *CreateSlotRequest) SetLifeCycle(v *SlotLifeCycle) *CreateSlotRequest {
	s.LifeCycle = v
	return s
}

func (s *CreateSlotRequest) SetName(v string) *CreateSlotRequest {
	s.Name = &v
	return s
}

func (s *CreateSlotRequest) SetStorageType(v string) *CreateSlotRequest {
	s.StorageType = &v
	return s
}

func (s *CreateSlotRequest) SetStorageUri(v string) *CreateSlotRequest {
	s.StorageUri = &v
	return s
}

func (s *CreateSlotRequest) SetTags(v []*CreateSlotRequestTags) *CreateSlotRequest {
	s.Tags = v
	return s
}

func (s *CreateSlotRequest) Validate() error {
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

type CreateSlotRequestEndpoints struct {
	// example:
	//
	// endpoint-1
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// VPC
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// example:
	//
	// vpc-j6co2fcdsl1q0gnuc3ym3
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	// example:
	//
	// vsw-j6cmr00qjyrft6jo2mg7g
	VswitchId *string `json:"VswitchId,omitempty" xml:"VswitchId,omitempty"`
}

func (s CreateSlotRequestEndpoints) String() string {
	return dara.Prettify(s)
}

func (s CreateSlotRequestEndpoints) GoString() string {
	return s.String()
}

func (s *CreateSlotRequestEndpoints) GetName() *string {
	return s.Name
}

func (s *CreateSlotRequestEndpoints) GetType() *string {
	return s.Type
}

func (s *CreateSlotRequestEndpoints) GetVpcId() *string {
	return s.VpcId
}

func (s *CreateSlotRequestEndpoints) GetVswitchId() *string {
	return s.VswitchId
}

func (s *CreateSlotRequestEndpoints) SetName(v string) *CreateSlotRequestEndpoints {
	s.Name = &v
	return s
}

func (s *CreateSlotRequestEndpoints) SetType(v string) *CreateSlotRequestEndpoints {
	s.Type = &v
	return s
}

func (s *CreateSlotRequestEndpoints) SetVpcId(v string) *CreateSlotRequestEndpoints {
	s.VpcId = &v
	return s
}

func (s *CreateSlotRequestEndpoints) SetVswitchId(v string) *CreateSlotRequestEndpoints {
	s.VswitchId = &v
	return s
}

func (s *CreateSlotRequestEndpoints) Validate() error {
	return dara.Validate(s)
}

type CreateSlotRequestTags struct {
	// example:
	//
	// dataset_version
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// example:
	//
	// v0.1.0
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateSlotRequestTags) String() string {
	return dara.Prettify(s)
}

func (s CreateSlotRequestTags) GoString() string {
	return s.String()
}

func (s *CreateSlotRequestTags) GetKey() *string {
	return s.Key
}

func (s *CreateSlotRequestTags) GetValue() *string {
	return s.Value
}

func (s *CreateSlotRequestTags) SetKey(v string) *CreateSlotRequestTags {
	s.Key = &v
	return s
}

func (s *CreateSlotRequestTags) SetValue(v string) *CreateSlotRequestTags {
	s.Value = &v
	return s
}

func (s *CreateSlotRequestTags) Validate() error {
	return dara.Validate(s)
}
