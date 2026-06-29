// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateSecurityClassifyCatalogRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCreateCommand(v *CreateSecurityClassifyCatalogRequestCreateCommand) *CreateSecurityClassifyCatalogRequest
	GetCreateCommand() *CreateSecurityClassifyCatalogRequestCreateCommand
	SetOpTenantId(v int64) *CreateSecurityClassifyCatalogRequest
	GetOpTenantId() *int64
}

type CreateSecurityClassifyCatalogRequest struct {
	// The create command.
	//
	// This parameter is required.
	CreateCommand *CreateSecurityClassifyCatalogRequestCreateCommand `json:"CreateCommand,omitempty" xml:"CreateCommand,omitempty" type:"Struct"`
	// The tenant ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
}

func (s CreateSecurityClassifyCatalogRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateSecurityClassifyCatalogRequest) GoString() string {
	return s.String()
}

func (s *CreateSecurityClassifyCatalogRequest) GetCreateCommand() *CreateSecurityClassifyCatalogRequestCreateCommand {
	return s.CreateCommand
}

func (s *CreateSecurityClassifyCatalogRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *CreateSecurityClassifyCatalogRequest) SetCreateCommand(v *CreateSecurityClassifyCatalogRequestCreateCommand) *CreateSecurityClassifyCatalogRequest {
	s.CreateCommand = v
	return s
}

func (s *CreateSecurityClassifyCatalogRequest) SetOpTenantId(v int64) *CreateSecurityClassifyCatalogRequest {
	s.OpTenantId = &v
	return s
}

func (s *CreateSecurityClassifyCatalogRequest) Validate() error {
	if s.CreateCommand != nil {
		if err := s.CreateCommand.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateSecurityClassifyCatalogRequestCreateCommand struct {
	// The name of the classification folder.
	//
	// This parameter is required.
	//
	// example:
	//
	// test
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The list of administrator IDs. This parameter takes effect only when the parent folder is the root folder.
	OwnerList []*string `json:"OwnerList,omitempty" xml:"OwnerList,omitempty" type:"Repeated"`
	// The full path of the parent folder. Default value: /.
	//
	// example:
	//
	// /d1/
	ParentPath *string `json:"ParentPath,omitempty" xml:"ParentPath,omitempty"`
	// The visibility scope of the classification folder. Valid values:
	//
	// - PUBLIC: visible to all users.
	//
	// - PRIVATE: visible only to administrators.
	//
	// Default value: PUBLIC. This parameter takes effect only when the parent folder is the root folder.
	//
	// example:
	//
	// PUBLIC
	VisibleType *string `json:"VisibleType,omitempty" xml:"VisibleType,omitempty"`
}

func (s CreateSecurityClassifyCatalogRequestCreateCommand) String() string {
	return dara.Prettify(s)
}

func (s CreateSecurityClassifyCatalogRequestCreateCommand) GoString() string {
	return s.String()
}

func (s *CreateSecurityClassifyCatalogRequestCreateCommand) GetName() *string {
	return s.Name
}

func (s *CreateSecurityClassifyCatalogRequestCreateCommand) GetOwnerList() []*string {
	return s.OwnerList
}

func (s *CreateSecurityClassifyCatalogRequestCreateCommand) GetParentPath() *string {
	return s.ParentPath
}

func (s *CreateSecurityClassifyCatalogRequestCreateCommand) GetVisibleType() *string {
	return s.VisibleType
}

func (s *CreateSecurityClassifyCatalogRequestCreateCommand) SetName(v string) *CreateSecurityClassifyCatalogRequestCreateCommand {
	s.Name = &v
	return s
}

func (s *CreateSecurityClassifyCatalogRequestCreateCommand) SetOwnerList(v []*string) *CreateSecurityClassifyCatalogRequestCreateCommand {
	s.OwnerList = v
	return s
}

func (s *CreateSecurityClassifyCatalogRequestCreateCommand) SetParentPath(v string) *CreateSecurityClassifyCatalogRequestCreateCommand {
	s.ParentPath = &v
	return s
}

func (s *CreateSecurityClassifyCatalogRequestCreateCommand) SetVisibleType(v string) *CreateSecurityClassifyCatalogRequestCreateCommand {
	s.VisibleType = &v
	return s
}

func (s *CreateSecurityClassifyCatalogRequestCreateCommand) Validate() error {
	return dara.Validate(s)
}
