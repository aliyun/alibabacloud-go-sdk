// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iFixDataRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEnv(v string) *FixDataRequest
	GetEnv() *string
	SetFixDataCommand(v *FixDataRequestFixDataCommand) *FixDataRequest
	GetFixDataCommand() *FixDataRequestFixDataCommand
	SetOpTenantId(v int64) *FixDataRequest
	GetOpTenantId() *int64
}

type FixDataRequest struct {
	// The environment identifier. Valid values:
	//
	// - DEV: development environment.
	//
	// - PROD (default): production environment.
	//
	// example:
	//
	// PROD
	Env *string `json:"Env,omitempty" xml:"Env,omitempty"`
	// The command to rerun downstream nodes to fix data link issues. You can choose to force a rerun.
	//
	// This parameter is required.
	FixDataCommand *FixDataRequestFixDataCommand `json:"FixDataCommand,omitempty" xml:"FixDataCommand,omitempty" type:"Struct"`
	// The tenant ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
}

func (s FixDataRequest) String() string {
	return dara.Prettify(s)
}

func (s FixDataRequest) GoString() string {
	return s.String()
}

func (s *FixDataRequest) GetEnv() *string {
	return s.Env
}

func (s *FixDataRequest) GetFixDataCommand() *FixDataRequestFixDataCommand {
	return s.FixDataCommand
}

func (s *FixDataRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *FixDataRequest) SetEnv(v string) *FixDataRequest {
	s.Env = &v
	return s
}

func (s *FixDataRequest) SetFixDataCommand(v *FixDataRequestFixDataCommand) *FixDataRequest {
	s.FixDataCommand = v
	return s
}

func (s *FixDataRequest) SetOpTenantId(v int64) *FixDataRequest {
	s.OpTenantId = &v
	return s
}

func (s *FixDataRequest) Validate() error {
	if s.FixDataCommand != nil {
		if err := s.FixDataCommand.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type FixDataRequestFixDataCommand struct {
	// Specifies whether to rerun the root instance. If you do not specify this parameter, the default value is true.
	//
	// example:
	//
	// false
	ContainRootInstance *bool `json:"ContainRootInstance,omitempty" xml:"ContainRootInstance,omitempty"`
	// The downstream instances. If you have specified a downstream range, you do not need to specify this parameter. Otherwise, you must specify the list of downstream instances.
	DownStreamInstanceIdList []*FixDataRequestFixDataCommandDownStreamInstanceIdList `json:"DownStreamInstanceIdList,omitempty" xml:"DownStreamInstanceIdList,omitempty" type:"Repeated"`
	// The downstream range. Valid values:
	//
	// - ALL_FAILED_INSTANCE: all failed instances.
	//
	// - ALL_INSTANCE: all instances.
	//
	// - ALL_FINAL_INSTANCE: all desired state instances.
	//
	// - If you do not specify this parameter, the rerun is performed based on the specified downstream instances.
	//
	// example:
	//
	// ALL_INSTANCE
	DownstreamRange *string `json:"DownstreamRange,omitempty" xml:"DownstreamRange,omitempty"`
	// Specifies whether to force a rerun.
	//
	// example:
	//
	// false
	ForceRerun *bool `json:"ForceRerun,omitempty" xml:"ForceRerun,omitempty"`
	// The project ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 132344
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The root instance.
	//
	// This parameter is required.
	RootInstanceId *FixDataRequestFixDataCommandRootInstanceId `json:"RootInstanceId,omitempty" xml:"RootInstanceId,omitempty" type:"Struct"`
}

func (s FixDataRequestFixDataCommand) String() string {
	return dara.Prettify(s)
}

func (s FixDataRequestFixDataCommand) GoString() string {
	return s.String()
}

func (s *FixDataRequestFixDataCommand) GetContainRootInstance() *bool {
	return s.ContainRootInstance
}

func (s *FixDataRequestFixDataCommand) GetDownStreamInstanceIdList() []*FixDataRequestFixDataCommandDownStreamInstanceIdList {
	return s.DownStreamInstanceIdList
}

func (s *FixDataRequestFixDataCommand) GetDownstreamRange() *string {
	return s.DownstreamRange
}

func (s *FixDataRequestFixDataCommand) GetForceRerun() *bool {
	return s.ForceRerun
}

func (s *FixDataRequestFixDataCommand) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *FixDataRequestFixDataCommand) GetRootInstanceId() *FixDataRequestFixDataCommandRootInstanceId {
	return s.RootInstanceId
}

func (s *FixDataRequestFixDataCommand) SetContainRootInstance(v bool) *FixDataRequestFixDataCommand {
	s.ContainRootInstance = &v
	return s
}

func (s *FixDataRequestFixDataCommand) SetDownStreamInstanceIdList(v []*FixDataRequestFixDataCommandDownStreamInstanceIdList) *FixDataRequestFixDataCommand {
	s.DownStreamInstanceIdList = v
	return s
}

func (s *FixDataRequestFixDataCommand) SetDownstreamRange(v string) *FixDataRequestFixDataCommand {
	s.DownstreamRange = &v
	return s
}

func (s *FixDataRequestFixDataCommand) SetForceRerun(v bool) *FixDataRequestFixDataCommand {
	s.ForceRerun = &v
	return s
}

func (s *FixDataRequestFixDataCommand) SetProjectId(v int64) *FixDataRequestFixDataCommand {
	s.ProjectId = &v
	return s
}

func (s *FixDataRequestFixDataCommand) SetRootInstanceId(v *FixDataRequestFixDataCommandRootInstanceId) *FixDataRequestFixDataCommand {
	s.RootInstanceId = v
	return s
}

func (s *FixDataRequestFixDataCommand) Validate() error {
	if s.DownStreamInstanceIdList != nil {
		for _, item := range s.DownStreamInstanceIdList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.RootInstanceId != nil {
		if err := s.RootInstanceId.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type FixDataRequestFixDataCommandDownStreamInstanceIdList struct {
	// The field instance ID.
	FieldInstanceIdList []*string `json:"FieldInstanceIdList,omitempty" xml:"FieldInstanceIdList,omitempty" type:"Repeated"`
	// The instance ID.
	//
	// example:
	//
	// t_2323421
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s FixDataRequestFixDataCommandDownStreamInstanceIdList) String() string {
	return dara.Prettify(s)
}

func (s FixDataRequestFixDataCommandDownStreamInstanceIdList) GoString() string {
	return s.String()
}

func (s *FixDataRequestFixDataCommandDownStreamInstanceIdList) GetFieldInstanceIdList() []*string {
	return s.FieldInstanceIdList
}

func (s *FixDataRequestFixDataCommandDownStreamInstanceIdList) GetId() *string {
	return s.Id
}

func (s *FixDataRequestFixDataCommandDownStreamInstanceIdList) SetFieldInstanceIdList(v []*string) *FixDataRequestFixDataCommandDownStreamInstanceIdList {
	s.FieldInstanceIdList = v
	return s
}

func (s *FixDataRequestFixDataCommandDownStreamInstanceIdList) SetId(v string) *FixDataRequestFixDataCommandDownStreamInstanceIdList {
	s.Id = &v
	return s
}

func (s *FixDataRequestFixDataCommandDownStreamInstanceIdList) Validate() error {
	return dara.Validate(s)
}

type FixDataRequestFixDataCommandRootInstanceId struct {
	// The field IDs. This parameter is available when the node is a logical table instance ID. If you do not specify this parameter, the full table is used by default.
	FieldInstanceIdList []*string `json:"FieldInstanceIdList,omitempty" xml:"FieldInstanceIdList,omitempty" type:"Repeated"`
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// t_2323111
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s FixDataRequestFixDataCommandRootInstanceId) String() string {
	return dara.Prettify(s)
}

func (s FixDataRequestFixDataCommandRootInstanceId) GoString() string {
	return s.String()
}

func (s *FixDataRequestFixDataCommandRootInstanceId) GetFieldInstanceIdList() []*string {
	return s.FieldInstanceIdList
}

func (s *FixDataRequestFixDataCommandRootInstanceId) GetId() *string {
	return s.Id
}

func (s *FixDataRequestFixDataCommandRootInstanceId) SetFieldInstanceIdList(v []*string) *FixDataRequestFixDataCommandRootInstanceId {
	s.FieldInstanceIdList = v
	return s
}

func (s *FixDataRequestFixDataCommandRootInstanceId) SetId(v string) *FixDataRequestFixDataCommandRootInstanceId {
	s.Id = &v
	return s
}

func (s *FixDataRequestFixDataCommandRootInstanceId) Validate() error {
	return dara.Validate(s)
}
