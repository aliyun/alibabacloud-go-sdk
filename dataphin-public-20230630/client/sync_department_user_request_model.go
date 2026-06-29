// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSyncDepartmentUserRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOpTenantId(v int64) *SyncDepartmentUserRequest
	GetOpTenantId() *int64
	SetSyncDepartmentUserCommand(v *SyncDepartmentUserRequestSyncDepartmentUserCommand) *SyncDepartmentUserRequest
	GetSyncDepartmentUserCommand() *SyncDepartmentUserRequestSyncDepartmentUserCommand
}

type SyncDepartmentUserRequest struct {
	// The tenant ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
	// The request command.
	//
	// This parameter is required.
	SyncDepartmentUserCommand *SyncDepartmentUserRequestSyncDepartmentUserCommand `json:"SyncDepartmentUserCommand,omitempty" xml:"SyncDepartmentUserCommand,omitempty" type:"Struct"`
}

func (s SyncDepartmentUserRequest) String() string {
	return dara.Prettify(s)
}

func (s SyncDepartmentUserRequest) GoString() string {
	return s.String()
}

func (s *SyncDepartmentUserRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *SyncDepartmentUserRequest) GetSyncDepartmentUserCommand() *SyncDepartmentUserRequestSyncDepartmentUserCommand {
	return s.SyncDepartmentUserCommand
}

func (s *SyncDepartmentUserRequest) SetOpTenantId(v int64) *SyncDepartmentUserRequest {
	s.OpTenantId = &v
	return s
}

func (s *SyncDepartmentUserRequest) SetSyncDepartmentUserCommand(v *SyncDepartmentUserRequestSyncDepartmentUserCommand) *SyncDepartmentUserRequest {
	s.SyncDepartmentUserCommand = v
	return s
}

func (s *SyncDepartmentUserRequest) Validate() error {
	if s.SyncDepartmentUserCommand != nil {
		if err := s.SyncDepartmentUserCommand.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type SyncDepartmentUserRequestSyncDepartmentUserCommand struct {
	// The mapping between users and their affiliated departments.
	//
	// This parameter is required.
	DeptUserMapping []*SyncDepartmentUserRequestSyncDepartmentUserCommandDeptUserMapping `json:"DeptUserMapping,omitempty" xml:"DeptUserMapping,omitempty" type:"Repeated"`
}

func (s SyncDepartmentUserRequestSyncDepartmentUserCommand) String() string {
	return dara.Prettify(s)
}

func (s SyncDepartmentUserRequestSyncDepartmentUserCommand) GoString() string {
	return s.String()
}

func (s *SyncDepartmentUserRequestSyncDepartmentUserCommand) GetDeptUserMapping() []*SyncDepartmentUserRequestSyncDepartmentUserCommandDeptUserMapping {
	return s.DeptUserMapping
}

func (s *SyncDepartmentUserRequestSyncDepartmentUserCommand) SetDeptUserMapping(v []*SyncDepartmentUserRequestSyncDepartmentUserCommandDeptUserMapping) *SyncDepartmentUserRequestSyncDepartmentUserCommand {
	s.DeptUserMapping = v
	return s
}

func (s *SyncDepartmentUserRequestSyncDepartmentUserCommand) Validate() error {
	if s.DeptUserMapping != nil {
		for _, item := range s.DeptUserMapping {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type SyncDepartmentUserRequestSyncDepartmentUserCommandDeptUserMapping struct {
	// The list of department IDs to which the user belongs. If this parameter is left empty, the user-department affiliation is deleted.
	DepartmentIdList []*string `json:"DepartmentIdList,omitempty" xml:"DepartmentIdList,omitempty" type:"Repeated"`
	// The user ID in the user system. This value is the unique identifier of the user.
	//
	// This parameter is required.
	//
	// example:
	//
	// 30000001
	SourceUserId *string `json:"SourceUserId,omitempty" xml:"SourceUserId,omitempty"`
}

func (s SyncDepartmentUserRequestSyncDepartmentUserCommandDeptUserMapping) String() string {
	return dara.Prettify(s)
}

func (s SyncDepartmentUserRequestSyncDepartmentUserCommandDeptUserMapping) GoString() string {
	return s.String()
}

func (s *SyncDepartmentUserRequestSyncDepartmentUserCommandDeptUserMapping) GetDepartmentIdList() []*string {
	return s.DepartmentIdList
}

func (s *SyncDepartmentUserRequestSyncDepartmentUserCommandDeptUserMapping) GetSourceUserId() *string {
	return s.SourceUserId
}

func (s *SyncDepartmentUserRequestSyncDepartmentUserCommandDeptUserMapping) SetDepartmentIdList(v []*string) *SyncDepartmentUserRequestSyncDepartmentUserCommandDeptUserMapping {
	s.DepartmentIdList = v
	return s
}

func (s *SyncDepartmentUserRequestSyncDepartmentUserCommandDeptUserMapping) SetSourceUserId(v string) *SyncDepartmentUserRequestSyncDepartmentUserCommandDeptUserMapping {
	s.SourceUserId = &v
	return s
}

func (s *SyncDepartmentUserRequestSyncDepartmentUserCommandDeptUserMapping) Validate() error {
	return dara.Validate(s)
}
