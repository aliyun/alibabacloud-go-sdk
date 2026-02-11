// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCheckServiceLinkedRoleForDeletingRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDeletionTaskId(v string) *CheckServiceLinkedRoleForDeletingRequest
	GetDeletionTaskId() *string
	SetRegionId(v string) *CheckServiceLinkedRoleForDeletingRequest
	GetRegionId() *string
	SetRoleArn(v string) *CheckServiceLinkedRoleForDeletingRequest
	GetRoleArn() *string
	SetSPIRegionId(v string) *CheckServiceLinkedRoleForDeletingRequest
	GetSPIRegionId() *string
	SetServiceName(v string) *CheckServiceLinkedRoleForDeletingRequest
	GetServiceName() *string
}

type CheckServiceLinkedRoleForDeletingRequest struct {
	// This parameter is required.
	DeletionTaskId *string `json:"DeletionTaskId,omitempty" xml:"DeletionTaskId,omitempty"`
	RegionId       *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// This parameter is required.
	RoleArn *string `json:"RoleArn,omitempty" xml:"RoleArn,omitempty"`
	// This parameter is required.
	SPIRegionId *string `json:"SPIRegionId,omitempty" xml:"SPIRegionId,omitempty"`
	// This parameter is required.
	ServiceName *string `json:"ServiceName,omitempty" xml:"ServiceName,omitempty"`
}

func (s CheckServiceLinkedRoleForDeletingRequest) String() string {
	return dara.Prettify(s)
}

func (s CheckServiceLinkedRoleForDeletingRequest) GoString() string {
	return s.String()
}

func (s *CheckServiceLinkedRoleForDeletingRequest) GetDeletionTaskId() *string {
	return s.DeletionTaskId
}

func (s *CheckServiceLinkedRoleForDeletingRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CheckServiceLinkedRoleForDeletingRequest) GetRoleArn() *string {
	return s.RoleArn
}

func (s *CheckServiceLinkedRoleForDeletingRequest) GetSPIRegionId() *string {
	return s.SPIRegionId
}

func (s *CheckServiceLinkedRoleForDeletingRequest) GetServiceName() *string {
	return s.ServiceName
}

func (s *CheckServiceLinkedRoleForDeletingRequest) SetDeletionTaskId(v string) *CheckServiceLinkedRoleForDeletingRequest {
	s.DeletionTaskId = &v
	return s
}

func (s *CheckServiceLinkedRoleForDeletingRequest) SetRegionId(v string) *CheckServiceLinkedRoleForDeletingRequest {
	s.RegionId = &v
	return s
}

func (s *CheckServiceLinkedRoleForDeletingRequest) SetRoleArn(v string) *CheckServiceLinkedRoleForDeletingRequest {
	s.RoleArn = &v
	return s
}

func (s *CheckServiceLinkedRoleForDeletingRequest) SetSPIRegionId(v string) *CheckServiceLinkedRoleForDeletingRequest {
	s.SPIRegionId = &v
	return s
}

func (s *CheckServiceLinkedRoleForDeletingRequest) SetServiceName(v string) *CheckServiceLinkedRoleForDeletingRequest {
	s.ServiceName = &v
	return s
}

func (s *CheckServiceLinkedRoleForDeletingRequest) Validate() error {
	return dara.Validate(s)
}
