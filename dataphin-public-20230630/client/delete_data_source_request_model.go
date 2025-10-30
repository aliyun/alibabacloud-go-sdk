// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDataSourceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDeleteCommand(v *DeleteDataSourceRequestDeleteCommand) *DeleteDataSourceRequest
	GetDeleteCommand() *DeleteDataSourceRequestDeleteCommand
	SetOpTenantId(v int64) *DeleteDataSourceRequest
	GetOpTenantId() *int64
}

type DeleteDataSourceRequest struct {
	// This parameter is required.
	DeleteCommand *DeleteDataSourceRequestDeleteCommand `json:"DeleteCommand,omitempty" xml:"DeleteCommand,omitempty" type:"Struct"`
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
}

func (s DeleteDataSourceRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteDataSourceRequest) GoString() string {
	return s.String()
}

func (s *DeleteDataSourceRequest) GetDeleteCommand() *DeleteDataSourceRequestDeleteCommand {
	return s.DeleteCommand
}

func (s *DeleteDataSourceRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *DeleteDataSourceRequest) SetDeleteCommand(v *DeleteDataSourceRequestDeleteCommand) *DeleteDataSourceRequest {
	s.DeleteCommand = v
	return s
}

func (s *DeleteDataSourceRequest) SetOpTenantId(v int64) *DeleteDataSourceRequest {
	s.OpTenantId = &v
	return s
}

func (s *DeleteDataSourceRequest) Validate() error {
	if s.DeleteCommand != nil {
		if err := s.DeleteCommand.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DeleteDataSourceRequestDeleteCommand struct {
	// This parameter is required.
	//
	// example:
	//
	// DEV
	Mode *string `json:"Mode,omitempty" xml:"Mode,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 13121
	ProdDataSourceId *int64 `json:"ProdDataSourceId,omitempty" xml:"ProdDataSourceId,omitempty"`
}

func (s DeleteDataSourceRequestDeleteCommand) String() string {
	return dara.Prettify(s)
}

func (s DeleteDataSourceRequestDeleteCommand) GoString() string {
	return s.String()
}

func (s *DeleteDataSourceRequestDeleteCommand) GetMode() *string {
	return s.Mode
}

func (s *DeleteDataSourceRequestDeleteCommand) GetProdDataSourceId() *int64 {
	return s.ProdDataSourceId
}

func (s *DeleteDataSourceRequestDeleteCommand) SetMode(v string) *DeleteDataSourceRequestDeleteCommand {
	s.Mode = &v
	return s
}

func (s *DeleteDataSourceRequestDeleteCommand) SetProdDataSourceId(v int64) *DeleteDataSourceRequestDeleteCommand {
	s.ProdDataSourceId = &v
	return s
}

func (s *DeleteDataSourceRequestDeleteCommand) Validate() error {
	return dara.Validate(s)
}
