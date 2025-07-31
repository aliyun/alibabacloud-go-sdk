// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGrantResourcePermissionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetGrantCommand(v *GrantResourcePermissionRequestGrantCommand) *GrantResourcePermissionRequest
	GetGrantCommand() *GrantResourcePermissionRequestGrantCommand
	SetOpTenantId(v int64) *GrantResourcePermissionRequest
	GetOpTenantId() *int64
}

type GrantResourcePermissionRequest struct {
	// This parameter is required.
	GrantCommand *GrantResourcePermissionRequestGrantCommand `json:"GrantCommand,omitempty" xml:"GrantCommand,omitempty" type:"Struct"`
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
}

func (s GrantResourcePermissionRequest) String() string {
	return dara.Prettify(s)
}

func (s GrantResourcePermissionRequest) GoString() string {
	return s.String()
}

func (s *GrantResourcePermissionRequest) GetGrantCommand() *GrantResourcePermissionRequestGrantCommand {
	return s.GrantCommand
}

func (s *GrantResourcePermissionRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *GrantResourcePermissionRequest) SetGrantCommand(v *GrantResourcePermissionRequestGrantCommand) *GrantResourcePermissionRequest {
	s.GrantCommand = v
	return s
}

func (s *GrantResourcePermissionRequest) SetOpTenantId(v int64) *GrantResourcePermissionRequest {
	s.OpTenantId = &v
	return s
}

func (s *GrantResourcePermissionRequest) Validate() error {
	return dara.Validate(s)
}

type GrantResourcePermissionRequestGrantCommand struct {
	// This parameter is required.
	//
	// example:
	//
	// 1717343597000
	EffectiveEnd *string `json:"EffectiveEnd,omitempty" xml:"EffectiveEnd,omitempty"`
	// This parameter is required.
	OperateList []*string `json:"OperateList,omitempty" xml:"OperateList,omitempty" type:"Repeated"`
	// example:
	//
	// xx
	Reason *string `json:"Reason,omitempty" xml:"Reason,omitempty"`
	// This parameter is required.
	ResourceList []*GrantResourcePermissionRequestGrantCommandResourceList `json:"ResourceList,omitempty" xml:"ResourceList,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// PHYSICAL_TABLE
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// This parameter is required.
	UserIdList []*string `json:"UserIdList,omitempty" xml:"UserIdList,omitempty" type:"Repeated"`
}

func (s GrantResourcePermissionRequestGrantCommand) String() string {
	return dara.Prettify(s)
}

func (s GrantResourcePermissionRequestGrantCommand) GoString() string {
	return s.String()
}

func (s *GrantResourcePermissionRequestGrantCommand) GetEffectiveEnd() *string {
	return s.EffectiveEnd
}

func (s *GrantResourcePermissionRequestGrantCommand) GetOperateList() []*string {
	return s.OperateList
}

func (s *GrantResourcePermissionRequestGrantCommand) GetReason() *string {
	return s.Reason
}

func (s *GrantResourcePermissionRequestGrantCommand) GetResourceList() []*GrantResourcePermissionRequestGrantCommandResourceList {
	return s.ResourceList
}

func (s *GrantResourcePermissionRequestGrantCommand) GetResourceType() *string {
	return s.ResourceType
}

func (s *GrantResourcePermissionRequestGrantCommand) GetUserIdList() []*string {
	return s.UserIdList
}

func (s *GrantResourcePermissionRequestGrantCommand) SetEffectiveEnd(v string) *GrantResourcePermissionRequestGrantCommand {
	s.EffectiveEnd = &v
	return s
}

func (s *GrantResourcePermissionRequestGrantCommand) SetOperateList(v []*string) *GrantResourcePermissionRequestGrantCommand {
	s.OperateList = v
	return s
}

func (s *GrantResourcePermissionRequestGrantCommand) SetReason(v string) *GrantResourcePermissionRequestGrantCommand {
	s.Reason = &v
	return s
}

func (s *GrantResourcePermissionRequestGrantCommand) SetResourceList(v []*GrantResourcePermissionRequestGrantCommandResourceList) *GrantResourcePermissionRequestGrantCommand {
	s.ResourceList = v
	return s
}

func (s *GrantResourcePermissionRequestGrantCommand) SetResourceType(v string) *GrantResourcePermissionRequestGrantCommand {
	s.ResourceType = &v
	return s
}

func (s *GrantResourcePermissionRequestGrantCommand) SetUserIdList(v []*string) *GrantResourcePermissionRequestGrantCommand {
	s.UserIdList = v
	return s
}

func (s *GrantResourcePermissionRequestGrantCommand) Validate() error {
	return dara.Validate(s)
}

type GrantResourcePermissionRequestGrantCommandResourceList struct {
	// example:
	//
	// hadoop.300000806.data_distill.behavior_gameinfor_01
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
}

func (s GrantResourcePermissionRequestGrantCommandResourceList) String() string {
	return dara.Prettify(s)
}

func (s GrantResourcePermissionRequestGrantCommandResourceList) GoString() string {
	return s.String()
}

func (s *GrantResourcePermissionRequestGrantCommandResourceList) GetResourceId() *string {
	return s.ResourceId
}

func (s *GrantResourcePermissionRequestGrantCommandResourceList) SetResourceId(v string) *GrantResourcePermissionRequestGrantCommandResourceList {
	s.ResourceId = &v
	return s
}

func (s *GrantResourcePermissionRequestGrantCommandResourceList) Validate() error {
	return dara.Validate(s)
}
