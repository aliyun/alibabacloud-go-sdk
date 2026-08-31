// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateBasicProjectRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOpTenantId(v int64) *UpdateBasicProjectRequest
	GetOpTenantId() *int64
	SetOpUserId(v string) *UpdateBasicProjectRequest
	GetOpUserId() *string
	SetUpdateCommand(v *UpdateBasicProjectRequestUpdateCommand) *UpdateBasicProjectRequest
	GetUpdateCommand() *UpdateBasicProjectRequestUpdateCommand
}

type UpdateBasicProjectRequest struct {
	// The tenant ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
	// example:
	//
	// 30001011
	OpUserId *string `json:"OpUserId,omitempty" xml:"OpUserId,omitempty"`
	// The update command.
	//
	// This parameter is required.
	UpdateCommand *UpdateBasicProjectRequestUpdateCommand `json:"UpdateCommand,omitempty" xml:"UpdateCommand,omitempty" type:"Struct"`
}

func (s UpdateBasicProjectRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateBasicProjectRequest) GoString() string {
	return s.String()
}

func (s *UpdateBasicProjectRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *UpdateBasicProjectRequest) GetOpUserId() *string {
	return s.OpUserId
}

func (s *UpdateBasicProjectRequest) GetUpdateCommand() *UpdateBasicProjectRequestUpdateCommand {
	return s.UpdateCommand
}

func (s *UpdateBasicProjectRequest) SetOpTenantId(v int64) *UpdateBasicProjectRequest {
	s.OpTenantId = &v
	return s
}

func (s *UpdateBasicProjectRequest) SetOpUserId(v string) *UpdateBasicProjectRequest {
	s.OpUserId = &v
	return s
}

func (s *UpdateBasicProjectRequest) SetUpdateCommand(v *UpdateBasicProjectRequestUpdateCommand) *UpdateBasicProjectRequest {
	s.UpdateCommand = v
	return s
}

func (s *UpdateBasicProjectRequest) Validate() error {
	if s.UpdateCommand != nil {
		if err := s.UpdateCommand.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateBasicProjectRequestUpdateCommand struct {
	// The business unit ID.
	//
	// example:
	//
	// 1001
	BizUnitId *int64 `json:"BizUnitId,omitempty" xml:"BizUnitId,omitempty"`
	// The ID of the offline compute source.
	//
	// example:
	//
	// 2001
	ComputeSourceId *int64 `json:"ComputeSourceId,omitempty" xml:"ComputeSourceId,omitempty"`
	// The project description.
	//
	// example:
	//
	// test project
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The display name of the project.
	//
	// example:
	//
	// MyProject.
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	// The project ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 102311
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The project name. This value cannot be modified. Pass in the current project name.
	//
	// This parameter is required.
	//
	// example:
	//
	// my_project
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The namespace identifier.
	//
	// example:
	//
	// dev
	NameSpaceTag *string `json:"NameSpaceTag,omitempty" xml:"NameSpaceTag,omitempty"`
	// The ID of the real-time compute source.
	//
	// example:
	//
	// 2002
	StreamComputeSourceId *int64 `json:"StreamComputeSourceId,omitempty" xml:"StreamComputeSourceId,omitempty"`
	// The sandbox whitelist.
	WhiteLists []*UpdateBasicProjectRequestUpdateCommandWhiteLists `json:"WhiteLists,omitempty" xml:"WhiteLists,omitempty" type:"Repeated"`
}

func (s UpdateBasicProjectRequestUpdateCommand) String() string {
	return dara.Prettify(s)
}

func (s UpdateBasicProjectRequestUpdateCommand) GoString() string {
	return s.String()
}

func (s *UpdateBasicProjectRequestUpdateCommand) GetBizUnitId() *int64 {
	return s.BizUnitId
}

func (s *UpdateBasicProjectRequestUpdateCommand) GetComputeSourceId() *int64 {
	return s.ComputeSourceId
}

func (s *UpdateBasicProjectRequestUpdateCommand) GetDescription() *string {
	return s.Description
}

func (s *UpdateBasicProjectRequestUpdateCommand) GetDisplayName() *string {
	return s.DisplayName
}

func (s *UpdateBasicProjectRequestUpdateCommand) GetId() *int64 {
	return s.Id
}

func (s *UpdateBasicProjectRequestUpdateCommand) GetName() *string {
	return s.Name
}

func (s *UpdateBasicProjectRequestUpdateCommand) GetNameSpaceTag() *string {
	return s.NameSpaceTag
}

func (s *UpdateBasicProjectRequestUpdateCommand) GetStreamComputeSourceId() *int64 {
	return s.StreamComputeSourceId
}

func (s *UpdateBasicProjectRequestUpdateCommand) GetWhiteLists() []*UpdateBasicProjectRequestUpdateCommandWhiteLists {
	return s.WhiteLists
}

func (s *UpdateBasicProjectRequestUpdateCommand) SetBizUnitId(v int64) *UpdateBasicProjectRequestUpdateCommand {
	s.BizUnitId = &v
	return s
}

func (s *UpdateBasicProjectRequestUpdateCommand) SetComputeSourceId(v int64) *UpdateBasicProjectRequestUpdateCommand {
	s.ComputeSourceId = &v
	return s
}

func (s *UpdateBasicProjectRequestUpdateCommand) SetDescription(v string) *UpdateBasicProjectRequestUpdateCommand {
	s.Description = &v
	return s
}

func (s *UpdateBasicProjectRequestUpdateCommand) SetDisplayName(v string) *UpdateBasicProjectRequestUpdateCommand {
	s.DisplayName = &v
	return s
}

func (s *UpdateBasicProjectRequestUpdateCommand) SetId(v int64) *UpdateBasicProjectRequestUpdateCommand {
	s.Id = &v
	return s
}

func (s *UpdateBasicProjectRequestUpdateCommand) SetName(v string) *UpdateBasicProjectRequestUpdateCommand {
	s.Name = &v
	return s
}

func (s *UpdateBasicProjectRequestUpdateCommand) SetNameSpaceTag(v string) *UpdateBasicProjectRequestUpdateCommand {
	s.NameSpaceTag = &v
	return s
}

func (s *UpdateBasicProjectRequestUpdateCommand) SetStreamComputeSourceId(v int64) *UpdateBasicProjectRequestUpdateCommand {
	s.StreamComputeSourceId = &v
	return s
}

func (s *UpdateBasicProjectRequestUpdateCommand) SetWhiteLists(v []*UpdateBasicProjectRequestUpdateCommandWhiteLists) *UpdateBasicProjectRequestUpdateCommand {
	s.WhiteLists = v
	return s
}

func (s *UpdateBasicProjectRequestUpdateCommand) Validate() error {
	if s.WhiteLists != nil {
		for _, item := range s.WhiteLists {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type UpdateBasicProjectRequestUpdateCommandWhiteLists struct {
	// The description.
	//
	// example:
	//
	// test
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// IP
	//
	// example:
	//
	// 127.0.0.1
	Ip *string `json:"Ip,omitempty" xml:"Ip,omitempty"`
	// The port.
	//
	// example:
	//
	// 8080
	Port *string `json:"Port,omitempty" xml:"Port,omitempty"`
}

func (s UpdateBasicProjectRequestUpdateCommandWhiteLists) String() string {
	return dara.Prettify(s)
}

func (s UpdateBasicProjectRequestUpdateCommandWhiteLists) GoString() string {
	return s.String()
}

func (s *UpdateBasicProjectRequestUpdateCommandWhiteLists) GetDescription() *string {
	return s.Description
}

func (s *UpdateBasicProjectRequestUpdateCommandWhiteLists) GetIp() *string {
	return s.Ip
}

func (s *UpdateBasicProjectRequestUpdateCommandWhiteLists) GetPort() *string {
	return s.Port
}

func (s *UpdateBasicProjectRequestUpdateCommandWhiteLists) SetDescription(v string) *UpdateBasicProjectRequestUpdateCommandWhiteLists {
	s.Description = &v
	return s
}

func (s *UpdateBasicProjectRequestUpdateCommandWhiteLists) SetIp(v string) *UpdateBasicProjectRequestUpdateCommandWhiteLists {
	s.Ip = &v
	return s
}

func (s *UpdateBasicProjectRequestUpdateCommandWhiteLists) SetPort(v string) *UpdateBasicProjectRequestUpdateCommandWhiteLists {
	s.Port = &v
	return s
}

func (s *UpdateBasicProjectRequestUpdateCommandWhiteLists) Validate() error {
	return dara.Validate(s)
}
