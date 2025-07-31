// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRevokeResourcePermissionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOpTenantId(v int64) *RevokeResourcePermissionRequest
	GetOpTenantId() *int64
	SetRevokeCommand(v *RevokeResourcePermissionRequestRevokeCommand) *RevokeResourcePermissionRequest
	GetRevokeCommand() *RevokeResourcePermissionRequestRevokeCommand
}

type RevokeResourcePermissionRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
	// This parameter is required.
	RevokeCommand *RevokeResourcePermissionRequestRevokeCommand `json:"RevokeCommand,omitempty" xml:"RevokeCommand,omitempty" type:"Struct"`
}

func (s RevokeResourcePermissionRequest) String() string {
	return dara.Prettify(s)
}

func (s RevokeResourcePermissionRequest) GoString() string {
	return s.String()
}

func (s *RevokeResourcePermissionRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *RevokeResourcePermissionRequest) GetRevokeCommand() *RevokeResourcePermissionRequestRevokeCommand {
	return s.RevokeCommand
}

func (s *RevokeResourcePermissionRequest) SetOpTenantId(v int64) *RevokeResourcePermissionRequest {
	s.OpTenantId = &v
	return s
}

func (s *RevokeResourcePermissionRequest) SetRevokeCommand(v *RevokeResourcePermissionRequestRevokeCommand) *RevokeResourcePermissionRequest {
	s.RevokeCommand = v
	return s
}

func (s *RevokeResourcePermissionRequest) Validate() error {
	return dara.Validate(s)
}

type RevokeResourcePermissionRequestRevokeCommand struct {
	OperateList []*string `json:"OperateList,omitempty" xml:"OperateList,omitempty" type:"Repeated"`
	// example:
	//
	// xx
	Reason *string `json:"Reason,omitempty" xml:"Reason,omitempty"`
	// This parameter is required.
	ResourceList []*RevokeResourcePermissionRequestRevokeCommandResourceList `json:"ResourceList,omitempty" xml:"ResourceList,omitempty" type:"Repeated"`
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
	// 13131
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s RevokeResourcePermissionRequestRevokeCommand) String() string {
	return dara.Prettify(s)
}

func (s RevokeResourcePermissionRequestRevokeCommand) GoString() string {
	return s.String()
}

func (s *RevokeResourcePermissionRequestRevokeCommand) GetOperateList() []*string {
	return s.OperateList
}

func (s *RevokeResourcePermissionRequestRevokeCommand) GetReason() *string {
	return s.Reason
}

func (s *RevokeResourcePermissionRequestRevokeCommand) GetResourceList() []*RevokeResourcePermissionRequestRevokeCommandResourceList {
	return s.ResourceList
}

func (s *RevokeResourcePermissionRequestRevokeCommand) GetResourceType() *string {
	return s.ResourceType
}

func (s *RevokeResourcePermissionRequestRevokeCommand) GetUserId() *string {
	return s.UserId
}

func (s *RevokeResourcePermissionRequestRevokeCommand) SetOperateList(v []*string) *RevokeResourcePermissionRequestRevokeCommand {
	s.OperateList = v
	return s
}

func (s *RevokeResourcePermissionRequestRevokeCommand) SetReason(v string) *RevokeResourcePermissionRequestRevokeCommand {
	s.Reason = &v
	return s
}

func (s *RevokeResourcePermissionRequestRevokeCommand) SetResourceList(v []*RevokeResourcePermissionRequestRevokeCommandResourceList) *RevokeResourcePermissionRequestRevokeCommand {
	s.ResourceList = v
	return s
}

func (s *RevokeResourcePermissionRequestRevokeCommand) SetResourceType(v string) *RevokeResourcePermissionRequestRevokeCommand {
	s.ResourceType = &v
	return s
}

func (s *RevokeResourcePermissionRequestRevokeCommand) SetUserId(v string) *RevokeResourcePermissionRequestRevokeCommand {
	s.UserId = &v
	return s
}

func (s *RevokeResourcePermissionRequestRevokeCommand) Validate() error {
	return dara.Validate(s)
}

type RevokeResourcePermissionRequestRevokeCommandResourceList struct {
	// example:
	//
	// odps.300002102.beginner_test.amin_table
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
}

func (s RevokeResourcePermissionRequestRevokeCommandResourceList) String() string {
	return dara.Prettify(s)
}

func (s RevokeResourcePermissionRequestRevokeCommandResourceList) GoString() string {
	return s.String()
}

func (s *RevokeResourcePermissionRequestRevokeCommandResourceList) GetResourceId() *string {
	return s.ResourceId
}

func (s *RevokeResourcePermissionRequestRevokeCommandResourceList) SetResourceId(v string) *RevokeResourcePermissionRequestRevokeCommandResourceList {
	s.ResourceId = &v
	return s
}

func (s *RevokeResourcePermissionRequestRevokeCommandResourceList) Validate() error {
	return dara.Validate(s)
}
