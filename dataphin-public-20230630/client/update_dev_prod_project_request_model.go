// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateDevProdProjectRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOpTenantId(v int64) *UpdateDevProdProjectRequest
	GetOpTenantId() *int64
	SetUpdateCommand(v *UpdateDevProdProjectRequestUpdateCommand) *UpdateDevProdProjectRequest
	GetUpdateCommand() *UpdateDevProdProjectRequestUpdateCommand
}

type UpdateDevProdProjectRequest struct {
	// The tenant ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
	// The update command.
	//
	// This parameter is required.
	UpdateCommand *UpdateDevProdProjectRequestUpdateCommand `json:"UpdateCommand,omitempty" xml:"UpdateCommand,omitempty" type:"Struct"`
}

func (s UpdateDevProdProjectRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateDevProdProjectRequest) GoString() string {
	return s.String()
}

func (s *UpdateDevProdProjectRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *UpdateDevProdProjectRequest) GetUpdateCommand() *UpdateDevProdProjectRequestUpdateCommand {
	return s.UpdateCommand
}

func (s *UpdateDevProdProjectRequest) SetOpTenantId(v int64) *UpdateDevProdProjectRequest {
	s.OpTenantId = &v
	return s
}

func (s *UpdateDevProdProjectRequest) SetUpdateCommand(v *UpdateDevProdProjectRequestUpdateCommand) *UpdateDevProdProjectRequest {
	s.UpdateCommand = v
	return s
}

func (s *UpdateDevProdProjectRequest) Validate() error {
	if s.UpdateCommand != nil {
		if err := s.UpdateCommand.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateDevProdProjectRequestUpdateCommand struct {
	// The business unit ID.
	//
	// example:
	//
	// 1001
	BizUnitId *int64 `json:"BizUnitId,omitempty" xml:"BizUnitId,omitempty"`
	// The ID of the offline compute source in the development environment.
	//
	// example:
	//
	// 2001
	DevComputeSourceId *int64 `json:"DevComputeSourceId,omitempty" xml:"DevComputeSourceId,omitempty"`
	// The description of the development environment.
	//
	// example:
	//
	// dev desc
	DevDescription *string `json:"DevDescription,omitempty" xml:"DevDescription,omitempty"`
	// The ID of the real-time compute source in the development environment.
	//
	// example:
	//
	// 2002
	DevStreamComputeSourceId *int64 `json:"DevStreamComputeSourceId,omitempty" xml:"DevStreamComputeSourceId,omitempty"`
	// The display name of the project.
	//
	// example:
	//
	// My project.
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	// The project ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 102311
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The project name.
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
	// The ID of the offline compute source in the production environment.
	//
	// example:
	//
	// 2003
	ProdComputeSourceId *int64 `json:"ProdComputeSourceId,omitempty" xml:"ProdComputeSourceId,omitempty"`
	// The description of the production environment.
	//
	// example:
	//
	// prod desc
	ProdDescription *string `json:"ProdDescription,omitempty" xml:"ProdDescription,omitempty"`
	// The ID of the real-time compute source in the production environment.
	//
	// example:
	//
	// 2004
	ProdStreamComputeSourceId *int64 `json:"ProdStreamComputeSourceId,omitempty" xml:"ProdStreamComputeSourceId,omitempty"`
	// The sandbox whitelist.
	WhiteLists []*UpdateDevProdProjectRequestUpdateCommandWhiteLists `json:"WhiteLists,omitempty" xml:"WhiteLists,omitempty" type:"Repeated"`
}

func (s UpdateDevProdProjectRequestUpdateCommand) String() string {
	return dara.Prettify(s)
}

func (s UpdateDevProdProjectRequestUpdateCommand) GoString() string {
	return s.String()
}

func (s *UpdateDevProdProjectRequestUpdateCommand) GetBizUnitId() *int64 {
	return s.BizUnitId
}

func (s *UpdateDevProdProjectRequestUpdateCommand) GetDevComputeSourceId() *int64 {
	return s.DevComputeSourceId
}

func (s *UpdateDevProdProjectRequestUpdateCommand) GetDevDescription() *string {
	return s.DevDescription
}

func (s *UpdateDevProdProjectRequestUpdateCommand) GetDevStreamComputeSourceId() *int64 {
	return s.DevStreamComputeSourceId
}

func (s *UpdateDevProdProjectRequestUpdateCommand) GetDisplayName() *string {
	return s.DisplayName
}

func (s *UpdateDevProdProjectRequestUpdateCommand) GetId() *int64 {
	return s.Id
}

func (s *UpdateDevProdProjectRequestUpdateCommand) GetName() *string {
	return s.Name
}

func (s *UpdateDevProdProjectRequestUpdateCommand) GetNameSpaceTag() *string {
	return s.NameSpaceTag
}

func (s *UpdateDevProdProjectRequestUpdateCommand) GetProdComputeSourceId() *int64 {
	return s.ProdComputeSourceId
}

func (s *UpdateDevProdProjectRequestUpdateCommand) GetProdDescription() *string {
	return s.ProdDescription
}

func (s *UpdateDevProdProjectRequestUpdateCommand) GetProdStreamComputeSourceId() *int64 {
	return s.ProdStreamComputeSourceId
}

func (s *UpdateDevProdProjectRequestUpdateCommand) GetWhiteLists() []*UpdateDevProdProjectRequestUpdateCommandWhiteLists {
	return s.WhiteLists
}

func (s *UpdateDevProdProjectRequestUpdateCommand) SetBizUnitId(v int64) *UpdateDevProdProjectRequestUpdateCommand {
	s.BizUnitId = &v
	return s
}

func (s *UpdateDevProdProjectRequestUpdateCommand) SetDevComputeSourceId(v int64) *UpdateDevProdProjectRequestUpdateCommand {
	s.DevComputeSourceId = &v
	return s
}

func (s *UpdateDevProdProjectRequestUpdateCommand) SetDevDescription(v string) *UpdateDevProdProjectRequestUpdateCommand {
	s.DevDescription = &v
	return s
}

func (s *UpdateDevProdProjectRequestUpdateCommand) SetDevStreamComputeSourceId(v int64) *UpdateDevProdProjectRequestUpdateCommand {
	s.DevStreamComputeSourceId = &v
	return s
}

func (s *UpdateDevProdProjectRequestUpdateCommand) SetDisplayName(v string) *UpdateDevProdProjectRequestUpdateCommand {
	s.DisplayName = &v
	return s
}

func (s *UpdateDevProdProjectRequestUpdateCommand) SetId(v int64) *UpdateDevProdProjectRequestUpdateCommand {
	s.Id = &v
	return s
}

func (s *UpdateDevProdProjectRequestUpdateCommand) SetName(v string) *UpdateDevProdProjectRequestUpdateCommand {
	s.Name = &v
	return s
}

func (s *UpdateDevProdProjectRequestUpdateCommand) SetNameSpaceTag(v string) *UpdateDevProdProjectRequestUpdateCommand {
	s.NameSpaceTag = &v
	return s
}

func (s *UpdateDevProdProjectRequestUpdateCommand) SetProdComputeSourceId(v int64) *UpdateDevProdProjectRequestUpdateCommand {
	s.ProdComputeSourceId = &v
	return s
}

func (s *UpdateDevProdProjectRequestUpdateCommand) SetProdDescription(v string) *UpdateDevProdProjectRequestUpdateCommand {
	s.ProdDescription = &v
	return s
}

func (s *UpdateDevProdProjectRequestUpdateCommand) SetProdStreamComputeSourceId(v int64) *UpdateDevProdProjectRequestUpdateCommand {
	s.ProdStreamComputeSourceId = &v
	return s
}

func (s *UpdateDevProdProjectRequestUpdateCommand) SetWhiteLists(v []*UpdateDevProdProjectRequestUpdateCommandWhiteLists) *UpdateDevProdProjectRequestUpdateCommand {
	s.WhiteLists = v
	return s
}

func (s *UpdateDevProdProjectRequestUpdateCommand) Validate() error {
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

type UpdateDevProdProjectRequestUpdateCommandWhiteLists struct {
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

func (s UpdateDevProdProjectRequestUpdateCommandWhiteLists) String() string {
	return dara.Prettify(s)
}

func (s UpdateDevProdProjectRequestUpdateCommandWhiteLists) GoString() string {
	return s.String()
}

func (s *UpdateDevProdProjectRequestUpdateCommandWhiteLists) GetDescription() *string {
	return s.Description
}

func (s *UpdateDevProdProjectRequestUpdateCommandWhiteLists) GetIp() *string {
	return s.Ip
}

func (s *UpdateDevProdProjectRequestUpdateCommandWhiteLists) GetPort() *string {
	return s.Port
}

func (s *UpdateDevProdProjectRequestUpdateCommandWhiteLists) SetDescription(v string) *UpdateDevProdProjectRequestUpdateCommandWhiteLists {
	s.Description = &v
	return s
}

func (s *UpdateDevProdProjectRequestUpdateCommandWhiteLists) SetIp(v string) *UpdateDevProdProjectRequestUpdateCommandWhiteLists {
	s.Ip = &v
	return s
}

func (s *UpdateDevProdProjectRequestUpdateCommandWhiteLists) SetPort(v string) *UpdateDevProdProjectRequestUpdateCommandWhiteLists {
	s.Port = &v
	return s
}

func (s *UpdateDevProdProjectRequestUpdateCommandWhiteLists) Validate() error {
	return dara.Validate(s)
}
