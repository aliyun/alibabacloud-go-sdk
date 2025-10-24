// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddMediaWorkflowRequest interface {
	dara.Model
	String() string
	GoString() string
	SetName(v string) *AddMediaWorkflowRequest
	GetName() *string
	SetOwnerAccount(v string) *AddMediaWorkflowRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *AddMediaWorkflowRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *AddMediaWorkflowRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *AddMediaWorkflowRequest
	GetResourceOwnerId() *int64
	SetTopology(v string) *AddMediaWorkflowRequest
	GetTopology() *string
	SetTriggerMode(v string) *AddMediaWorkflowRequest
	GetTriggerMode() *string
}

type AddMediaWorkflowRequest struct {
	// The name of the media workflow.
	//
	// 	- The value cannot be empty.
	//
	// 	- The name cannot be the same as that of an existing media workflow within the current Alibaba Cloud account.
	//
	// 	- The name can be up to 64 characters in length.
	//
	// 	- The value must be encoded in the UTF-8 format.
	//
	// This parameter is required.
	//
	// example:
	//
	// mediaworkflow-example
	Name                 *string `json:"Name,omitempty" xml:"Name,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The topology of the media workflow. The value must be a JSON object that contains the activities and activity dependencies. For more information, see the **Sample topology*	- section of this topic.
	//
	// >  The Object Storage Service (OSS) bucket must reside in the same region as your MPS service.
	//
	// This parameter is required.
	Topology *string `json:"Topology,omitempty" xml:"Topology,omitempty"`
	// The triggering mode of the media workflow. Valid values:
	//
	// 	- **OssAutoTrigger**: The media workflow is automatically triggered.
	//
	// 	- **NotInAuto**: The media workflow is not automatically triggered.
	//
	// example:
	//
	// OssAutoTrigger
	TriggerMode *string `json:"TriggerMode,omitempty" xml:"TriggerMode,omitempty"`
}

func (s AddMediaWorkflowRequest) String() string {
	return dara.Prettify(s)
}

func (s AddMediaWorkflowRequest) GoString() string {
	return s.String()
}

func (s *AddMediaWorkflowRequest) GetName() *string {
	return s.Name
}

func (s *AddMediaWorkflowRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *AddMediaWorkflowRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *AddMediaWorkflowRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *AddMediaWorkflowRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *AddMediaWorkflowRequest) GetTopology() *string {
	return s.Topology
}

func (s *AddMediaWorkflowRequest) GetTriggerMode() *string {
	return s.TriggerMode
}

func (s *AddMediaWorkflowRequest) SetName(v string) *AddMediaWorkflowRequest {
	s.Name = &v
	return s
}

func (s *AddMediaWorkflowRequest) SetOwnerAccount(v string) *AddMediaWorkflowRequest {
	s.OwnerAccount = &v
	return s
}

func (s *AddMediaWorkflowRequest) SetOwnerId(v int64) *AddMediaWorkflowRequest {
	s.OwnerId = &v
	return s
}

func (s *AddMediaWorkflowRequest) SetResourceOwnerAccount(v string) *AddMediaWorkflowRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *AddMediaWorkflowRequest) SetResourceOwnerId(v int64) *AddMediaWorkflowRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *AddMediaWorkflowRequest) SetTopology(v string) *AddMediaWorkflowRequest {
	s.Topology = &v
	return s
}

func (s *AddMediaWorkflowRequest) SetTriggerMode(v string) *AddMediaWorkflowRequest {
	s.TriggerMode = &v
	return s
}

func (s *AddMediaWorkflowRequest) Validate() error {
	return dara.Validate(s)
}
