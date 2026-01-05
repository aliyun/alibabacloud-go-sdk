// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateInstanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCapacity(v string) *CreateInstanceRequest
	GetCapacity() *string
	SetDescription(v string) *CreateInstanceRequest
	GetDescription() *string
	SetMaxEndpoint(v string) *CreateInstanceRequest
	GetMaxEndpoint() *string
	SetMaxSlot(v string) *CreateInstanceRequest
	GetMaxSlot() *string
	SetName(v string) *CreateInstanceRequest
	GetName() *string
	SetPaymentType(v string) *CreateInstanceRequest
	GetPaymentType() *string
	SetProviderType(v string) *CreateInstanceRequest
	GetProviderType() *string
	SetStorageType(v string) *CreateInstanceRequest
	GetStorageType() *string
	SetTags(v []*CreateInstanceRequestTags) *CreateInstanceRequest
	GetTags() []*CreateInstanceRequestTags
	SetType(v string) *CreateInstanceRequest
	GetType() *string
}

type CreateInstanceRequest struct {
	// This parameter is required.
	//
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
	// 20
	MaxEndpoint *string `json:"MaxEndpoint,omitempty" xml:"MaxEndpoint,omitempty"`
	// example:
	//
	// 20
	MaxSlot *string `json:"MaxSlot,omitempty" xml:"MaxSlot,omitempty"`
	// example:
	//
	// acc_instance_1
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// PayAsYouGo
	PaymentType *string `json:"PaymentType,omitempty" xml:"PaymentType,omitempty"`
	// example:
	//
	// Ecs
	ProviderType *string `json:"ProviderType,omitempty" xml:"ProviderType,omitempty"`
	// example:
	//
	// OSS
	StorageType *string                      `json:"StorageType,omitempty" xml:"StorageType,omitempty"`
	Tags        []*CreateInstanceRequestTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// basic
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s CreateInstanceRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateInstanceRequest) GoString() string {
	return s.String()
}

func (s *CreateInstanceRequest) GetCapacity() *string {
	return s.Capacity
}

func (s *CreateInstanceRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateInstanceRequest) GetMaxEndpoint() *string {
	return s.MaxEndpoint
}

func (s *CreateInstanceRequest) GetMaxSlot() *string {
	return s.MaxSlot
}

func (s *CreateInstanceRequest) GetName() *string {
	return s.Name
}

func (s *CreateInstanceRequest) GetPaymentType() *string {
	return s.PaymentType
}

func (s *CreateInstanceRequest) GetProviderType() *string {
	return s.ProviderType
}

func (s *CreateInstanceRequest) GetStorageType() *string {
	return s.StorageType
}

func (s *CreateInstanceRequest) GetTags() []*CreateInstanceRequestTags {
	return s.Tags
}

func (s *CreateInstanceRequest) GetType() *string {
	return s.Type
}

func (s *CreateInstanceRequest) SetCapacity(v string) *CreateInstanceRequest {
	s.Capacity = &v
	return s
}

func (s *CreateInstanceRequest) SetDescription(v string) *CreateInstanceRequest {
	s.Description = &v
	return s
}

func (s *CreateInstanceRequest) SetMaxEndpoint(v string) *CreateInstanceRequest {
	s.MaxEndpoint = &v
	return s
}

func (s *CreateInstanceRequest) SetMaxSlot(v string) *CreateInstanceRequest {
	s.MaxSlot = &v
	return s
}

func (s *CreateInstanceRequest) SetName(v string) *CreateInstanceRequest {
	s.Name = &v
	return s
}

func (s *CreateInstanceRequest) SetPaymentType(v string) *CreateInstanceRequest {
	s.PaymentType = &v
	return s
}

func (s *CreateInstanceRequest) SetProviderType(v string) *CreateInstanceRequest {
	s.ProviderType = &v
	return s
}

func (s *CreateInstanceRequest) SetStorageType(v string) *CreateInstanceRequest {
	s.StorageType = &v
	return s
}

func (s *CreateInstanceRequest) SetTags(v []*CreateInstanceRequestTags) *CreateInstanceRequest {
	s.Tags = v
	return s
}

func (s *CreateInstanceRequest) SetType(v string) *CreateInstanceRequest {
	s.Type = &v
	return s
}

func (s *CreateInstanceRequest) Validate() error {
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

type CreateInstanceRequestTags struct {
	// example:
	//
	// dataset_version
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// example:
	//
	// v0.1.0
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateInstanceRequestTags) String() string {
	return dara.Prettify(s)
}

func (s CreateInstanceRequestTags) GoString() string {
	return s.String()
}

func (s *CreateInstanceRequestTags) GetKey() *string {
	return s.Key
}

func (s *CreateInstanceRequestTags) GetValue() *string {
	return s.Value
}

func (s *CreateInstanceRequestTags) SetKey(v string) *CreateInstanceRequestTags {
	s.Key = &v
	return s
}

func (s *CreateInstanceRequestTags) SetValue(v string) *CreateInstanceRequestTags {
	s.Value = &v
	return s
}

func (s *CreateInstanceRequestTags) Validate() error {
	return dara.Validate(s)
}
