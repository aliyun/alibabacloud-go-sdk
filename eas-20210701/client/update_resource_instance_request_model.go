// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateResourceInstanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAction(v string) *UpdateResourceInstanceRequest
	GetAction() *string
	SetNewDiskSize(v string) *UpdateResourceInstanceRequest
	GetNewDiskSize() *string
	SetReason(v string) *UpdateResourceInstanceRequest
	GetReason() *string
}

type UpdateResourceInstanceRequest struct {
	// Updates the schedulable behavior of the machine instance in the dedicated resource group. Valid values:
	//
	// - Uncordon: Allows services to be scheduled to this instance.
	//
	// - Cordon: Prevents services from being scheduled to this instance.
	//
	// - Drain: Evicts service instances that are already scheduled to this instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// Cordon
	Action *string `json:"Action,omitempty" xml:"Action,omitempty"`
	// The size of the system cloud disk after scaling, in GiB.
	//
	// example:
	//
	// 500
	NewDiskSize *string `json:"NewDiskSize,omitempty" xml:"NewDiskSize,omitempty"`
	// Takes effect when Action is set to Cordon or Drain. Records the reason for the operation.
	//
	// example:
	//
	// Node maintenance
	Reason *string `json:"Reason,omitempty" xml:"Reason,omitempty"`
}

func (s UpdateResourceInstanceRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateResourceInstanceRequest) GoString() string {
	return s.String()
}

func (s *UpdateResourceInstanceRequest) GetAction() *string {
	return s.Action
}

func (s *UpdateResourceInstanceRequest) GetNewDiskSize() *string {
	return s.NewDiskSize
}

func (s *UpdateResourceInstanceRequest) GetReason() *string {
	return s.Reason
}

func (s *UpdateResourceInstanceRequest) SetAction(v string) *UpdateResourceInstanceRequest {
	s.Action = &v
	return s
}

func (s *UpdateResourceInstanceRequest) SetNewDiskSize(v string) *UpdateResourceInstanceRequest {
	s.NewDiskSize = &v
	return s
}

func (s *UpdateResourceInstanceRequest) SetReason(v string) *UpdateResourceInstanceRequest {
	s.Reason = &v
	return s
}

func (s *UpdateResourceInstanceRequest) Validate() error {
	return dara.Validate(s)
}
