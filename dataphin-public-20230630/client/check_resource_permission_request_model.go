// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCheckResourcePermissionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCheckCommand(v *CheckResourcePermissionRequestCheckCommand) *CheckResourcePermissionRequest
	GetCheckCommand() *CheckResourcePermissionRequestCheckCommand
	SetOpTenantId(v int64) *CheckResourcePermissionRequest
	GetOpTenantId() *int64
}

type CheckResourcePermissionRequest struct {
	// This parameter is required.
	CheckCommand *CheckResourcePermissionRequestCheckCommand `json:"CheckCommand,omitempty" xml:"CheckCommand,omitempty" type:"Struct"`
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
}

func (s CheckResourcePermissionRequest) String() string {
	return dara.Prettify(s)
}

func (s CheckResourcePermissionRequest) GoString() string {
	return s.String()
}

func (s *CheckResourcePermissionRequest) GetCheckCommand() *CheckResourcePermissionRequestCheckCommand {
	return s.CheckCommand
}

func (s *CheckResourcePermissionRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *CheckResourcePermissionRequest) SetCheckCommand(v *CheckResourcePermissionRequestCheckCommand) *CheckResourcePermissionRequest {
	s.CheckCommand = v
	return s
}

func (s *CheckResourcePermissionRequest) SetOpTenantId(v int64) *CheckResourcePermissionRequest {
	s.OpTenantId = &v
	return s
}

func (s *CheckResourcePermissionRequest) Validate() error {
	if s.CheckCommand != nil {
		if err := s.CheckCommand.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CheckResourcePermissionRequestCheckCommand struct {
	// This parameter is required.
	//
	// example:
	//
	// UPDATE
	Operate *string `json:"Operate,omitempty" xml:"Operate,omitempty"`
	// This parameter is required.
	ResourceList []*CheckResourcePermissionRequestCheckCommandResourceList `json:"ResourceList,omitempty" xml:"ResourceList,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// PHYSICAL_TABLE
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 323231
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s CheckResourcePermissionRequestCheckCommand) String() string {
	return dara.Prettify(s)
}

func (s CheckResourcePermissionRequestCheckCommand) GoString() string {
	return s.String()
}

func (s *CheckResourcePermissionRequestCheckCommand) GetOperate() *string {
	return s.Operate
}

func (s *CheckResourcePermissionRequestCheckCommand) GetResourceList() []*CheckResourcePermissionRequestCheckCommandResourceList {
	return s.ResourceList
}

func (s *CheckResourcePermissionRequestCheckCommand) GetResourceType() *string {
	return s.ResourceType
}

func (s *CheckResourcePermissionRequestCheckCommand) GetUserId() *string {
	return s.UserId
}

func (s *CheckResourcePermissionRequestCheckCommand) SetOperate(v string) *CheckResourcePermissionRequestCheckCommand {
	s.Operate = &v
	return s
}

func (s *CheckResourcePermissionRequestCheckCommand) SetResourceList(v []*CheckResourcePermissionRequestCheckCommandResourceList) *CheckResourcePermissionRequestCheckCommand {
	s.ResourceList = v
	return s
}

func (s *CheckResourcePermissionRequestCheckCommand) SetResourceType(v string) *CheckResourcePermissionRequestCheckCommand {
	s.ResourceType = &v
	return s
}

func (s *CheckResourcePermissionRequestCheckCommand) SetUserId(v string) *CheckResourcePermissionRequestCheckCommand {
	s.UserId = &v
	return s
}

func (s *CheckResourcePermissionRequestCheckCommand) Validate() error {
	if s.ResourceList != nil {
		for _, item := range s.ResourceList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type CheckResourcePermissionRequestCheckCommandResourceList struct {
	// This parameter is required.
	//
	// example:
	//
	// hadoop.300000806.data_distill.behavior_gameinfor_01
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
}

func (s CheckResourcePermissionRequestCheckCommandResourceList) String() string {
	return dara.Prettify(s)
}

func (s CheckResourcePermissionRequestCheckCommandResourceList) GoString() string {
	return s.String()
}

func (s *CheckResourcePermissionRequestCheckCommandResourceList) GetResourceId() *string {
	return s.ResourceId
}

func (s *CheckResourcePermissionRequestCheckCommandResourceList) SetResourceId(v string) *CheckResourcePermissionRequestCheckCommandResourceList {
	s.ResourceId = &v
	return s
}

func (s *CheckResourcePermissionRequestCheckCommandResourceList) Validate() error {
	return dara.Validate(s)
}
