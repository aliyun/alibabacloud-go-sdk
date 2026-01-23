// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteSecurityClassifyCatalogRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDeleteCommand(v *DeleteSecurityClassifyCatalogRequestDeleteCommand) *DeleteSecurityClassifyCatalogRequest
	GetDeleteCommand() *DeleteSecurityClassifyCatalogRequestDeleteCommand
	SetOpTenantId(v int64) *DeleteSecurityClassifyCatalogRequest
	GetOpTenantId() *int64
}

type DeleteSecurityClassifyCatalogRequest struct {
	// This parameter is required.
	DeleteCommand *DeleteSecurityClassifyCatalogRequestDeleteCommand `json:"DeleteCommand,omitempty" xml:"DeleteCommand,omitempty" type:"Struct"`
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
}

func (s DeleteSecurityClassifyCatalogRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteSecurityClassifyCatalogRequest) GoString() string {
	return s.String()
}

func (s *DeleteSecurityClassifyCatalogRequest) GetDeleteCommand() *DeleteSecurityClassifyCatalogRequestDeleteCommand {
	return s.DeleteCommand
}

func (s *DeleteSecurityClassifyCatalogRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *DeleteSecurityClassifyCatalogRequest) SetDeleteCommand(v *DeleteSecurityClassifyCatalogRequestDeleteCommand) *DeleteSecurityClassifyCatalogRequest {
	s.DeleteCommand = v
	return s
}

func (s *DeleteSecurityClassifyCatalogRequest) SetOpTenantId(v int64) *DeleteSecurityClassifyCatalogRequest {
	s.OpTenantId = &v
	return s
}

func (s *DeleteSecurityClassifyCatalogRequest) Validate() error {
	if s.DeleteCommand != nil {
		if err := s.DeleteCommand.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DeleteSecurityClassifyCatalogRequestDeleteCommand struct {
	// This parameter is required.
	//
	// example:
	//
	// /d1/d2/
	Path                 *string `json:"Path,omitempty" xml:"Path,omitempty"`
	ReturnRemovedDetails *bool   `json:"ReturnRemovedDetails,omitempty" xml:"ReturnRemovedDetails,omitempty"`
}

func (s DeleteSecurityClassifyCatalogRequestDeleteCommand) String() string {
	return dara.Prettify(s)
}

func (s DeleteSecurityClassifyCatalogRequestDeleteCommand) GoString() string {
	return s.String()
}

func (s *DeleteSecurityClassifyCatalogRequestDeleteCommand) GetPath() *string {
	return s.Path
}

func (s *DeleteSecurityClassifyCatalogRequestDeleteCommand) GetReturnRemovedDetails() *bool {
	return s.ReturnRemovedDetails
}

func (s *DeleteSecurityClassifyCatalogRequestDeleteCommand) SetPath(v string) *DeleteSecurityClassifyCatalogRequestDeleteCommand {
	s.Path = &v
	return s
}

func (s *DeleteSecurityClassifyCatalogRequestDeleteCommand) SetReturnRemovedDetails(v bool) *DeleteSecurityClassifyCatalogRequestDeleteCommand {
	s.ReturnRemovedDetails = &v
	return s
}

func (s *DeleteSecurityClassifyCatalogRequestDeleteCommand) Validate() error {
	return dara.Validate(s)
}
