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
	// This parameter is required.
	CreateCommand *CreateSecurityClassifyCatalogRequestCreateCommand `json:"CreateCommand,omitempty" xml:"CreateCommand,omitempty" type:"Struct"`
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
	// This parameter is required.
	//
	// example:
	//
	// test
	Name      *string   `json:"Name,omitempty" xml:"Name,omitempty"`
	OwnerList []*string `json:"OwnerList,omitempty" xml:"OwnerList,omitempty" type:"Repeated"`
	// example:
	//
	// /d1/
	ParentPath *string `json:"ParentPath,omitempty" xml:"ParentPath,omitempty"`
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
