// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateMediaWorkflowTriggerModeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMediaWorkflowId(v string) *UpdateMediaWorkflowTriggerModeRequest
	GetMediaWorkflowId() *string
	SetOwnerAccount(v string) *UpdateMediaWorkflowTriggerModeRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *UpdateMediaWorkflowTriggerModeRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *UpdateMediaWorkflowTriggerModeRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *UpdateMediaWorkflowTriggerModeRequest
	GetResourceOwnerId() *int64
	SetTriggerMode(v string) *UpdateMediaWorkflowTriggerModeRequest
	GetTriggerMode() *string
}

type UpdateMediaWorkflowTriggerModeRequest struct {
	// The ID of the media workflow that you want to update. To obtain the ID of the media workflow, you can log on to the **ApsaraVideo Media Processing (MPS) console*	- and choose **Workflows*	- > **Workflow Settings*	- in the left-side navigation pane.
	//
	// This parameter is required.
	//
	// example:
	//
	// e00732b977da427d9177a4dee646****
	MediaWorkflowId      *string `json:"MediaWorkflowId,omitempty" xml:"MediaWorkflowId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The trigger mode of the media workflow. Valid values:
	//
	// 	- **OssAutoTrigger**: automatically triggers the media workflow.
	//
	// 	- **NotInAuto**: does not automatically trigger the media workflow.
	//
	// This parameter is required.
	//
	// example:
	//
	// NotInAuto
	TriggerMode *string `json:"TriggerMode,omitempty" xml:"TriggerMode,omitempty"`
}

func (s UpdateMediaWorkflowTriggerModeRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateMediaWorkflowTriggerModeRequest) GoString() string {
	return s.String()
}

func (s *UpdateMediaWorkflowTriggerModeRequest) GetMediaWorkflowId() *string {
	return s.MediaWorkflowId
}

func (s *UpdateMediaWorkflowTriggerModeRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *UpdateMediaWorkflowTriggerModeRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *UpdateMediaWorkflowTriggerModeRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *UpdateMediaWorkflowTriggerModeRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *UpdateMediaWorkflowTriggerModeRequest) GetTriggerMode() *string {
	return s.TriggerMode
}

func (s *UpdateMediaWorkflowTriggerModeRequest) SetMediaWorkflowId(v string) *UpdateMediaWorkflowTriggerModeRequest {
	s.MediaWorkflowId = &v
	return s
}

func (s *UpdateMediaWorkflowTriggerModeRequest) SetOwnerAccount(v string) *UpdateMediaWorkflowTriggerModeRequest {
	s.OwnerAccount = &v
	return s
}

func (s *UpdateMediaWorkflowTriggerModeRequest) SetOwnerId(v int64) *UpdateMediaWorkflowTriggerModeRequest {
	s.OwnerId = &v
	return s
}

func (s *UpdateMediaWorkflowTriggerModeRequest) SetResourceOwnerAccount(v string) *UpdateMediaWorkflowTriggerModeRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *UpdateMediaWorkflowTriggerModeRequest) SetResourceOwnerId(v int64) *UpdateMediaWorkflowTriggerModeRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *UpdateMediaWorkflowTriggerModeRequest) SetTriggerMode(v string) *UpdateMediaWorkflowTriggerModeRequest {
	s.TriggerMode = &v
	return s
}

func (s *UpdateMediaWorkflowTriggerModeRequest) Validate() error {
	return dara.Validate(s)
}
