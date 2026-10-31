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
	SetOpUserId(v string) *CheckResourcePermissionRequest
	GetOpUserId() *string
}

type CheckResourcePermissionRequest struct {
	// The user resource permission check command.
	//
	// This parameter is required.
	CheckCommand *CheckResourcePermissionRequestCheckCommand `json:"CheckCommand,omitempty" xml:"CheckCommand,omitempty" type:"Struct"`
	// The tenant ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
	// The ID of the operator.
	//
	// example:
	//
	// 30001011
	OpUserId *string `json:"OpUserId,omitempty" xml:"OpUserId,omitempty"`
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

func (s *CheckResourcePermissionRequest) GetOpUserId() *string {
	return s.OpUserId
}

func (s *CheckResourcePermissionRequest) SetCheckCommand(v *CheckResourcePermissionRequestCheckCommand) *CheckResourcePermissionRequest {
	s.CheckCommand = v
	return s
}

func (s *CheckResourcePermissionRequest) SetOpTenantId(v int64) *CheckResourcePermissionRequest {
	s.OpTenantId = &v
	return s
}

func (s *CheckResourcePermissionRequest) SetOpUserId(v string) *CheckResourcePermissionRequest {
	s.OpUserId = &v
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
	// The operation type.
	//
	// This parameter is required.
	//
	// example:
	//
	// UPDATE
	Operate *string `json:"Operate,omitempty" xml:"Operate,omitempty"`
	// The resources.
	//
	// This parameter is required.
	ResourceList []*CheckResourcePermissionRequestCheckCommandResourceList `json:"ResourceList,omitempty" xml:"ResourceList,omitempty" type:"Repeated"`
	// The resource type. Valid values:
	//
	// - PHYSICAL_TABLE: Physical table.
	//
	// - PHYSICAL_FIELD: Physical table field.
	//
	// - LOGICAL_TABLE: Logical fact table.
	//
	// - LOGICAL_FIELD: Logical fact table field.
	//
	// - LABEL_TABLE: Logical label table.
	//
	// - LABEL_FIELD: Logical label table field.
	//
	// - DATASOURCE: Data source.
	//
	// - GLOBAL_PARAM: Global parameter.
	//
	// - REALTIME_LOGICAL_TABLE: Real-time metadata table.
	//
	// - REALTIME_LOGICAL_FIELD: Real-time metadata table field.
	//
	// - REALTIME_MIRROR_TABLE: Mirror table.
	//
	// - REALTIME_MIRROR_FIELD: Real-time mirror table field.
	//
	// - FUNCTION: Function.
	//
	// - FEATURE: Feature permission.
	//
	// - PHYSICAL_VIEW: Physical view.
	//
	// - LOGICAL_VIEW: Logical view.
	//
	// - QD_SERVICE_FEATURE: Service label.
	//
	// - QD_ADVANCED_FEATURE: Advanced label.
	//
	// - QD_CLUSTER: Group.
	//
	// - QD_EVENT: Event.
	//
	// - QD_OFFLINE_SERVICE: Offline label service task.
	//
	// - PHYSICAL_MATERIALIZED_VIEW: Materialized view.
	//
	// This parameter is required.
	//
	// example:
	//
	// PHYSICAL_TABLE
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// The user ID.
	//
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
	// The resource ID.
	//
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
