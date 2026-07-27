// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateBasicProjectRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCreateCommand(v *CreateBasicProjectRequestCreateCommand) *CreateBasicProjectRequest
	GetCreateCommand() *CreateBasicProjectRequestCreateCommand
	SetOpTenantId(v int64) *CreateBasicProjectRequest
	GetOpTenantId() *int64
}

type CreateBasicProjectRequest struct {
	// The create command.
	//
	// This parameter is required.
	CreateCommand *CreateBasicProjectRequestCreateCommand `json:"CreateCommand,omitempty" xml:"CreateCommand,omitempty" type:"Struct"`
	// The tenant ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
}

func (s CreateBasicProjectRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateBasicProjectRequest) GoString() string {
	return s.String()
}

func (s *CreateBasicProjectRequest) GetCreateCommand() *CreateBasicProjectRequestCreateCommand {
	return s.CreateCommand
}

func (s *CreateBasicProjectRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *CreateBasicProjectRequest) SetCreateCommand(v *CreateBasicProjectRequestCreateCommand) *CreateBasicProjectRequest {
	s.CreateCommand = v
	return s
}

func (s *CreateBasicProjectRequest) SetOpTenantId(v int64) *CreateBasicProjectRequest {
	s.OpTenantId = &v
	return s
}

func (s *CreateBasicProjectRequest) Validate() error {
	if s.CreateCommand != nil {
		if err := s.CreateCommand.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateBasicProjectRequestCreateCommand struct {
	// The business unit ID.
	//
	// example:
	//
	// 1001
	BizUnitId *int64 `json:"BizUnitId,omitempty" xml:"BizUnitId,omitempty"`
	// The offline compute source ID.
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
	// The project display name.
	//
	// example:
	//
	// My project.
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	// The project name.
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
	// The real-time compute source ID.
	//
	// example:
	//
	// 2002
	StreamComputeSourceId *int64 `json:"StreamComputeSourceId,omitempty" xml:"StreamComputeSourceId,omitempty"`
	// The project type. If this parameter is left empty, the default value GENERAL is used.
	//
	// example:
	//
	// GENERAL
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The sandbox whitelist.
	WhiteLists []*CreateBasicProjectRequestCreateCommandWhiteLists `json:"WhiteLists,omitempty" xml:"WhiteLists,omitempty" type:"Repeated"`
}

func (s CreateBasicProjectRequestCreateCommand) String() string {
	return dara.Prettify(s)
}

func (s CreateBasicProjectRequestCreateCommand) GoString() string {
	return s.String()
}

func (s *CreateBasicProjectRequestCreateCommand) GetBizUnitId() *int64 {
	return s.BizUnitId
}

func (s *CreateBasicProjectRequestCreateCommand) GetComputeSourceId() *int64 {
	return s.ComputeSourceId
}

func (s *CreateBasicProjectRequestCreateCommand) GetDescription() *string {
	return s.Description
}

func (s *CreateBasicProjectRequestCreateCommand) GetDisplayName() *string {
	return s.DisplayName
}

func (s *CreateBasicProjectRequestCreateCommand) GetName() *string {
	return s.Name
}

func (s *CreateBasicProjectRequestCreateCommand) GetNameSpaceTag() *string {
	return s.NameSpaceTag
}

func (s *CreateBasicProjectRequestCreateCommand) GetStreamComputeSourceId() *int64 {
	return s.StreamComputeSourceId
}

func (s *CreateBasicProjectRequestCreateCommand) GetType() *string {
	return s.Type
}

func (s *CreateBasicProjectRequestCreateCommand) GetWhiteLists() []*CreateBasicProjectRequestCreateCommandWhiteLists {
	return s.WhiteLists
}

func (s *CreateBasicProjectRequestCreateCommand) SetBizUnitId(v int64) *CreateBasicProjectRequestCreateCommand {
	s.BizUnitId = &v
	return s
}

func (s *CreateBasicProjectRequestCreateCommand) SetComputeSourceId(v int64) *CreateBasicProjectRequestCreateCommand {
	s.ComputeSourceId = &v
	return s
}

func (s *CreateBasicProjectRequestCreateCommand) SetDescription(v string) *CreateBasicProjectRequestCreateCommand {
	s.Description = &v
	return s
}

func (s *CreateBasicProjectRequestCreateCommand) SetDisplayName(v string) *CreateBasicProjectRequestCreateCommand {
	s.DisplayName = &v
	return s
}

func (s *CreateBasicProjectRequestCreateCommand) SetName(v string) *CreateBasicProjectRequestCreateCommand {
	s.Name = &v
	return s
}

func (s *CreateBasicProjectRequestCreateCommand) SetNameSpaceTag(v string) *CreateBasicProjectRequestCreateCommand {
	s.NameSpaceTag = &v
	return s
}

func (s *CreateBasicProjectRequestCreateCommand) SetStreamComputeSourceId(v int64) *CreateBasicProjectRequestCreateCommand {
	s.StreamComputeSourceId = &v
	return s
}

func (s *CreateBasicProjectRequestCreateCommand) SetType(v string) *CreateBasicProjectRequestCreateCommand {
	s.Type = &v
	return s
}

func (s *CreateBasicProjectRequestCreateCommand) SetWhiteLists(v []*CreateBasicProjectRequestCreateCommandWhiteLists) *CreateBasicProjectRequestCreateCommand {
	s.WhiteLists = v
	return s
}

func (s *CreateBasicProjectRequestCreateCommand) Validate() error {
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

type CreateBasicProjectRequestCreateCommandWhiteLists struct {
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

func (s CreateBasicProjectRequestCreateCommandWhiteLists) String() string {
	return dara.Prettify(s)
}

func (s CreateBasicProjectRequestCreateCommandWhiteLists) GoString() string {
	return s.String()
}

func (s *CreateBasicProjectRequestCreateCommandWhiteLists) GetDescription() *string {
	return s.Description
}

func (s *CreateBasicProjectRequestCreateCommandWhiteLists) GetIp() *string {
	return s.Ip
}

func (s *CreateBasicProjectRequestCreateCommandWhiteLists) GetPort() *string {
	return s.Port
}

func (s *CreateBasicProjectRequestCreateCommandWhiteLists) SetDescription(v string) *CreateBasicProjectRequestCreateCommandWhiteLists {
	s.Description = &v
	return s
}

func (s *CreateBasicProjectRequestCreateCommandWhiteLists) SetIp(v string) *CreateBasicProjectRequestCreateCommandWhiteLists {
	s.Ip = &v
	return s
}

func (s *CreateBasicProjectRequestCreateCommandWhiteLists) SetPort(v string) *CreateBasicProjectRequestCreateCommandWhiteLists {
	s.Port = &v
	return s
}

func (s *CreateBasicProjectRequestCreateCommandWhiteLists) Validate() error {
	return dara.Validate(s)
}
