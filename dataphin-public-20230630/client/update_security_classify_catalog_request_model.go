// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateSecurityClassifyCatalogRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOpTenantId(v int64) *UpdateSecurityClassifyCatalogRequest
	GetOpTenantId() *int64
	SetUpdateCommand(v *UpdateSecurityClassifyCatalogRequestUpdateCommand) *UpdateSecurityClassifyCatalogRequest
	GetUpdateCommand() *UpdateSecurityClassifyCatalogRequestUpdateCommand
}

type UpdateSecurityClassifyCatalogRequest struct {
	// The tenant ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
	// The update instruction.
	//
	// This parameter is required.
	UpdateCommand *UpdateSecurityClassifyCatalogRequestUpdateCommand `json:"UpdateCommand,omitempty" xml:"UpdateCommand,omitempty" type:"Struct"`
}

func (s UpdateSecurityClassifyCatalogRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateSecurityClassifyCatalogRequest) GoString() string {
	return s.String()
}

func (s *UpdateSecurityClassifyCatalogRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *UpdateSecurityClassifyCatalogRequest) GetUpdateCommand() *UpdateSecurityClassifyCatalogRequestUpdateCommand {
	return s.UpdateCommand
}

func (s *UpdateSecurityClassifyCatalogRequest) SetOpTenantId(v int64) *UpdateSecurityClassifyCatalogRequest {
	s.OpTenantId = &v
	return s
}

func (s *UpdateSecurityClassifyCatalogRequest) SetUpdateCommand(v *UpdateSecurityClassifyCatalogRequestUpdateCommand) *UpdateSecurityClassifyCatalogRequest {
	s.UpdateCommand = v
	return s
}

func (s *UpdateSecurityClassifyCatalogRequest) Validate() error {
	if s.UpdateCommand != nil {
		if err := s.UpdateCommand.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateSecurityClassifyCatalogRequestUpdateCommand struct {
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
	// The full path of the parent classification folder. Default value: /.
	//
	// example:
	//
	// /d1/
	ParentPath *string `json:"ParentPath,omitempty" xml:"ParentPath,omitempty"`
	// The original full path of the folder.
	//
	// This parameter is required.
	//
	// example:
	//
	// /d1/d2/
	Path *string `json:"Path,omitempty" xml:"Path,omitempty"`
	// The visibility scope of the classification folder. This parameter takes effect only when the parent folder is the root folder. Valid values:
	//
	// - PUBLIC: visible to all users.
	//
	// - PRIVATE: visible only to administrators.
	//
	// Default value: PUBLIC.
	//
	// example:
	//
	// PUBLIC
	VisibleType *string `json:"VisibleType,omitempty" xml:"VisibleType,omitempty"`
}

func (s UpdateSecurityClassifyCatalogRequestUpdateCommand) String() string {
	return dara.Prettify(s)
}

func (s UpdateSecurityClassifyCatalogRequestUpdateCommand) GoString() string {
	return s.String()
}

func (s *UpdateSecurityClassifyCatalogRequestUpdateCommand) GetName() *string {
	return s.Name
}

func (s *UpdateSecurityClassifyCatalogRequestUpdateCommand) GetOwnerList() []*string {
	return s.OwnerList
}

func (s *UpdateSecurityClassifyCatalogRequestUpdateCommand) GetParentPath() *string {
	return s.ParentPath
}

func (s *UpdateSecurityClassifyCatalogRequestUpdateCommand) GetPath() *string {
	return s.Path
}

func (s *UpdateSecurityClassifyCatalogRequestUpdateCommand) GetVisibleType() *string {
	return s.VisibleType
}

func (s *UpdateSecurityClassifyCatalogRequestUpdateCommand) SetName(v string) *UpdateSecurityClassifyCatalogRequestUpdateCommand {
	s.Name = &v
	return s
}

func (s *UpdateSecurityClassifyCatalogRequestUpdateCommand) SetOwnerList(v []*string) *UpdateSecurityClassifyCatalogRequestUpdateCommand {
	s.OwnerList = v
	return s
}

func (s *UpdateSecurityClassifyCatalogRequestUpdateCommand) SetParentPath(v string) *UpdateSecurityClassifyCatalogRequestUpdateCommand {
	s.ParentPath = &v
	return s
}

func (s *UpdateSecurityClassifyCatalogRequestUpdateCommand) SetPath(v string) *UpdateSecurityClassifyCatalogRequestUpdateCommand {
	s.Path = &v
	return s
}

func (s *UpdateSecurityClassifyCatalogRequestUpdateCommand) SetVisibleType(v string) *UpdateSecurityClassifyCatalogRequestUpdateCommand {
	s.VisibleType = &v
	return s
}

func (s *UpdateSecurityClassifyCatalogRequestUpdateCommand) Validate() error {
	return dara.Validate(s)
}
