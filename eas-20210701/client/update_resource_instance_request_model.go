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
	// The operation that updates the scheduling state of the instance in a dedicated resource group. Valid values:
	//
	// 	- Uncordon: allows scheduling the service to this instance.
	//
	// 	- Cordon: prohibits scheduling the service to this instance.
	//
	// 	- Drain: evicts the service that has been scheduled to this instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// Cordon
	Action      *string `json:"Action,omitempty" xml:"Action,omitempty"`
	NewDiskSize *string `json:"NewDiskSize,omitempty" xml:"NewDiskSize,omitempty"`
	Reason      *string `json:"Reason,omitempty" xml:"Reason,omitempty"`
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
