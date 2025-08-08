// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyInstancePropertiesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *ModifyInstancePropertiesRequest
	GetInstanceId() *string
	SetInstanceIds(v []*string) *ModifyInstancePropertiesRequest
	GetInstanceIds() []*string
	SetKey(v string) *ModifyInstancePropertiesRequest
	GetKey() *string
	SetResourceType(v string) *ModifyInstancePropertiesRequest
	GetResourceType() *string
	SetValue(v string) *ModifyInstancePropertiesRequest
	GetValue() *string
}

type ModifyInstancePropertiesRequest struct {
	// example:
	//
	// mdp-0c62ayep0nk4v****
	InstanceId  *string   `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	InstanceIds []*string `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty" type:"Repeated"`
	// example:
	//
	// PackageUsedUpStrategy
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// DurationPackage
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// example:
	//
	// Postpaid
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ModifyInstancePropertiesRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyInstancePropertiesRequest) GoString() string {
	return s.String()
}

func (s *ModifyInstancePropertiesRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ModifyInstancePropertiesRequest) GetInstanceIds() []*string {
	return s.InstanceIds
}

func (s *ModifyInstancePropertiesRequest) GetKey() *string {
	return s.Key
}

func (s *ModifyInstancePropertiesRequest) GetResourceType() *string {
	return s.ResourceType
}

func (s *ModifyInstancePropertiesRequest) GetValue() *string {
	return s.Value
}

func (s *ModifyInstancePropertiesRequest) SetInstanceId(v string) *ModifyInstancePropertiesRequest {
	s.InstanceId = &v
	return s
}

func (s *ModifyInstancePropertiesRequest) SetInstanceIds(v []*string) *ModifyInstancePropertiesRequest {
	s.InstanceIds = v
	return s
}

func (s *ModifyInstancePropertiesRequest) SetKey(v string) *ModifyInstancePropertiesRequest {
	s.Key = &v
	return s
}

func (s *ModifyInstancePropertiesRequest) SetResourceType(v string) *ModifyInstancePropertiesRequest {
	s.ResourceType = &v
	return s
}

func (s *ModifyInstancePropertiesRequest) SetValue(v string) *ModifyInstancePropertiesRequest {
	s.Value = &v
	return s
}

func (s *ModifyInstancePropertiesRequest) Validate() error {
	return dara.Validate(s)
}
