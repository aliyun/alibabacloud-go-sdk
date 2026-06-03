// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateInstanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCallingNumber(v []*string) *CreateInstanceRequest
	GetCallingNumber() []*string
	SetInstanceDescription(v string) *CreateInstanceRequest
	GetInstanceDescription() *string
	SetInstanceName(v string) *CreateInstanceRequest
	GetInstanceName() *string
	SetMaxConcurrentConversation(v int32) *CreateInstanceRequest
	GetMaxConcurrentConversation() *int32
	SetNluServiceType(v string) *CreateInstanceRequest
	GetNluServiceType() *string
	SetResourceGroupId(v string) *CreateInstanceRequest
	GetResourceGroupId() *string
}

type CreateInstanceRequest struct {
	CallingNumber       []*string `json:"CallingNumber,omitempty" xml:"CallingNumber,omitempty" type:"Repeated"`
	InstanceDescription *string   `json:"InstanceDescription,omitempty" xml:"InstanceDescription,omitempty"`
	// This parameter is required.
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 4
	MaxConcurrentConversation *int32 `json:"MaxConcurrentConversation,omitempty" xml:"MaxConcurrentConversation,omitempty"`
	// example:
	//
	// Provided
	NluServiceType *string `json:"NluServiceType,omitempty" xml:"NluServiceType,omitempty"`
	// example:
	//
	// rg-acfmwd4qr3z773y
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
}

func (s CreateInstanceRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateInstanceRequest) GoString() string {
	return s.String()
}

func (s *CreateInstanceRequest) GetCallingNumber() []*string {
	return s.CallingNumber
}

func (s *CreateInstanceRequest) GetInstanceDescription() *string {
	return s.InstanceDescription
}

func (s *CreateInstanceRequest) GetInstanceName() *string {
	return s.InstanceName
}

func (s *CreateInstanceRequest) GetMaxConcurrentConversation() *int32 {
	return s.MaxConcurrentConversation
}

func (s *CreateInstanceRequest) GetNluServiceType() *string {
	return s.NluServiceType
}

func (s *CreateInstanceRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *CreateInstanceRequest) SetCallingNumber(v []*string) *CreateInstanceRequest {
	s.CallingNumber = v
	return s
}

func (s *CreateInstanceRequest) SetInstanceDescription(v string) *CreateInstanceRequest {
	s.InstanceDescription = &v
	return s
}

func (s *CreateInstanceRequest) SetInstanceName(v string) *CreateInstanceRequest {
	s.InstanceName = &v
	return s
}

func (s *CreateInstanceRequest) SetMaxConcurrentConversation(v int32) *CreateInstanceRequest {
	s.MaxConcurrentConversation = &v
	return s
}

func (s *CreateInstanceRequest) SetNluServiceType(v string) *CreateInstanceRequest {
	s.NluServiceType = &v
	return s
}

func (s *CreateInstanceRequest) SetResourceGroupId(v string) *CreateInstanceRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *CreateInstanceRequest) Validate() error {
	return dara.Validate(s)
}
