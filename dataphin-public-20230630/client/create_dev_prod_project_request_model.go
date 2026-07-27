// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDevProdProjectRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCreateCommand(v *CreateDevProdProjectRequestCreateCommand) *CreateDevProdProjectRequest
	GetCreateCommand() *CreateDevProdProjectRequestCreateCommand
	SetOpTenantId(v int64) *CreateDevProdProjectRequest
	GetOpTenantId() *int64
}

type CreateDevProdProjectRequest struct {
	// The create command.
	//
	// This parameter is required.
	CreateCommand *CreateDevProdProjectRequestCreateCommand `json:"CreateCommand,omitempty" xml:"CreateCommand,omitempty" type:"Struct"`
	// The tenant ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
}

func (s CreateDevProdProjectRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateDevProdProjectRequest) GoString() string {
	return s.String()
}

func (s *CreateDevProdProjectRequest) GetCreateCommand() *CreateDevProdProjectRequestCreateCommand {
	return s.CreateCommand
}

func (s *CreateDevProdProjectRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *CreateDevProdProjectRequest) SetCreateCommand(v *CreateDevProdProjectRequestCreateCommand) *CreateDevProdProjectRequest {
	s.CreateCommand = v
	return s
}

func (s *CreateDevProdProjectRequest) SetOpTenantId(v int64) *CreateDevProdProjectRequest {
	s.OpTenantId = &v
	return s
}

func (s *CreateDevProdProjectRequest) Validate() error {
	if s.CreateCommand != nil {
		if err := s.CreateCommand.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateDevProdProjectRequestCreateCommand struct {
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
	WhiteLists []*CreateDevProdProjectRequestCreateCommandWhiteLists `json:"WhiteLists,omitempty" xml:"WhiteLists,omitempty" type:"Repeated"`
}

func (s CreateDevProdProjectRequestCreateCommand) String() string {
	return dara.Prettify(s)
}

func (s CreateDevProdProjectRequestCreateCommand) GoString() string {
	return s.String()
}

func (s *CreateDevProdProjectRequestCreateCommand) GetBizUnitId() *int64 {
	return s.BizUnitId
}

func (s *CreateDevProdProjectRequestCreateCommand) GetDevComputeSourceId() *int64 {
	return s.DevComputeSourceId
}

func (s *CreateDevProdProjectRequestCreateCommand) GetDevDescription() *string {
	return s.DevDescription
}

func (s *CreateDevProdProjectRequestCreateCommand) GetDevStreamComputeSourceId() *int64 {
	return s.DevStreamComputeSourceId
}

func (s *CreateDevProdProjectRequestCreateCommand) GetDisplayName() *string {
	return s.DisplayName
}

func (s *CreateDevProdProjectRequestCreateCommand) GetName() *string {
	return s.Name
}

func (s *CreateDevProdProjectRequestCreateCommand) GetNameSpaceTag() *string {
	return s.NameSpaceTag
}

func (s *CreateDevProdProjectRequestCreateCommand) GetProdComputeSourceId() *int64 {
	return s.ProdComputeSourceId
}

func (s *CreateDevProdProjectRequestCreateCommand) GetProdDescription() *string {
	return s.ProdDescription
}

func (s *CreateDevProdProjectRequestCreateCommand) GetProdStreamComputeSourceId() *int64 {
	return s.ProdStreamComputeSourceId
}

func (s *CreateDevProdProjectRequestCreateCommand) GetWhiteLists() []*CreateDevProdProjectRequestCreateCommandWhiteLists {
	return s.WhiteLists
}

func (s *CreateDevProdProjectRequestCreateCommand) SetBizUnitId(v int64) *CreateDevProdProjectRequestCreateCommand {
	s.BizUnitId = &v
	return s
}

func (s *CreateDevProdProjectRequestCreateCommand) SetDevComputeSourceId(v int64) *CreateDevProdProjectRequestCreateCommand {
	s.DevComputeSourceId = &v
	return s
}

func (s *CreateDevProdProjectRequestCreateCommand) SetDevDescription(v string) *CreateDevProdProjectRequestCreateCommand {
	s.DevDescription = &v
	return s
}

func (s *CreateDevProdProjectRequestCreateCommand) SetDevStreamComputeSourceId(v int64) *CreateDevProdProjectRequestCreateCommand {
	s.DevStreamComputeSourceId = &v
	return s
}

func (s *CreateDevProdProjectRequestCreateCommand) SetDisplayName(v string) *CreateDevProdProjectRequestCreateCommand {
	s.DisplayName = &v
	return s
}

func (s *CreateDevProdProjectRequestCreateCommand) SetName(v string) *CreateDevProdProjectRequestCreateCommand {
	s.Name = &v
	return s
}

func (s *CreateDevProdProjectRequestCreateCommand) SetNameSpaceTag(v string) *CreateDevProdProjectRequestCreateCommand {
	s.NameSpaceTag = &v
	return s
}

func (s *CreateDevProdProjectRequestCreateCommand) SetProdComputeSourceId(v int64) *CreateDevProdProjectRequestCreateCommand {
	s.ProdComputeSourceId = &v
	return s
}

func (s *CreateDevProdProjectRequestCreateCommand) SetProdDescription(v string) *CreateDevProdProjectRequestCreateCommand {
	s.ProdDescription = &v
	return s
}

func (s *CreateDevProdProjectRequestCreateCommand) SetProdStreamComputeSourceId(v int64) *CreateDevProdProjectRequestCreateCommand {
	s.ProdStreamComputeSourceId = &v
	return s
}

func (s *CreateDevProdProjectRequestCreateCommand) SetWhiteLists(v []*CreateDevProdProjectRequestCreateCommandWhiteLists) *CreateDevProdProjectRequestCreateCommand {
	s.WhiteLists = v
	return s
}

func (s *CreateDevProdProjectRequestCreateCommand) Validate() error {
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

type CreateDevProdProjectRequestCreateCommandWhiteLists struct {
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

func (s CreateDevProdProjectRequestCreateCommandWhiteLists) String() string {
	return dara.Prettify(s)
}

func (s CreateDevProdProjectRequestCreateCommandWhiteLists) GoString() string {
	return s.String()
}

func (s *CreateDevProdProjectRequestCreateCommandWhiteLists) GetDescription() *string {
	return s.Description
}

func (s *CreateDevProdProjectRequestCreateCommandWhiteLists) GetIp() *string {
	return s.Ip
}

func (s *CreateDevProdProjectRequestCreateCommandWhiteLists) GetPort() *string {
	return s.Port
}

func (s *CreateDevProdProjectRequestCreateCommandWhiteLists) SetDescription(v string) *CreateDevProdProjectRequestCreateCommandWhiteLists {
	s.Description = &v
	return s
}

func (s *CreateDevProdProjectRequestCreateCommandWhiteLists) SetIp(v string) *CreateDevProdProjectRequestCreateCommandWhiteLists {
	s.Ip = &v
	return s
}

func (s *CreateDevProdProjectRequestCreateCommandWhiteLists) SetPort(v string) *CreateDevProdProjectRequestCreateCommandWhiteLists {
	s.Port = &v
	return s
}

func (s *CreateDevProdProjectRequestCreateCommandWhiteLists) Validate() error {
	return dara.Validate(s)
}
