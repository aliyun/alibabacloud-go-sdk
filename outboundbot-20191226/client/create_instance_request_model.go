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
	// The list of calling numbers.
	//
	// > This parameter is left empty by default.
	CallingNumber []*string `json:"CallingNumber,omitempty" xml:"CallingNumber,omitempty" type:"Repeated"`
	// The instance description.
	//
	// example:
	//
	// 这个是第一的实例
	InstanceDescription *string `json:"InstanceDescription,omitempty" xml:"InstanceDescription,omitempty"`
	// The instance name.
	//
	// This parameter is required.
	//
	// example:
	//
	// 第一个实例
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// The maximum concurrent conversations for the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// 4
	MaxConcurrentConversation *int32 `json:"MaxConcurrentConversation,omitempty" xml:"MaxConcurrentConversation,omitempty"`
	// The service type.
	//
	// > If this parameter is not specified, the default value is Managed.
	//
	// - DialogStudio: Conversation Engine 3.0.
	//
	// - Authorized: Public Cloud Conversation Bot. This refers to the Intelligent Speech Interaction product service purchased by public cloud customers (a private customer service). Authorization is completed via Scenario Management > Edit > Invoke Service > Custom Service.
	//
	// - Provided: On Premises Deployment Conversation Bot.
	//
	// - Managed: Legacy outbound calling canvas. This is the default Intelligent Speech Interaction product service (public service) for Intelligent Outbound Calling products.
	//
	// example:
	//
	// DialogStudio
	NluServiceType *string `json:"NluServiceType,omitempty" xml:"NluServiceType,omitempty"`
	// The resource group ID.
	//
	// > You can obtain this from the Resource Management document. For more information, see the following link: https://api.aliyun.com/document/ResourceManager/2020-03-31/ListResourceGroups
	//
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
