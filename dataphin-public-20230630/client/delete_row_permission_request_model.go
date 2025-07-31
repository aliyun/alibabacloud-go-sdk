// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteRowPermissionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDeleteRowPermissionCommand(v *DeleteRowPermissionRequestDeleteRowPermissionCommand) *DeleteRowPermissionRequest
	GetDeleteRowPermissionCommand() *DeleteRowPermissionRequestDeleteRowPermissionCommand
	SetOpTenantId(v int64) *DeleteRowPermissionRequest
	GetOpTenantId() *int64
}

type DeleteRowPermissionRequest struct {
	// This parameter is required.
	DeleteRowPermissionCommand *DeleteRowPermissionRequestDeleteRowPermissionCommand `json:"DeleteRowPermissionCommand,omitempty" xml:"DeleteRowPermissionCommand,omitempty" type:"Struct"`
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
}

func (s DeleteRowPermissionRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteRowPermissionRequest) GoString() string {
	return s.String()
}

func (s *DeleteRowPermissionRequest) GetDeleteRowPermissionCommand() *DeleteRowPermissionRequestDeleteRowPermissionCommand {
	return s.DeleteRowPermissionCommand
}

func (s *DeleteRowPermissionRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *DeleteRowPermissionRequest) SetDeleteRowPermissionCommand(v *DeleteRowPermissionRequestDeleteRowPermissionCommand) *DeleteRowPermissionRequest {
	s.DeleteRowPermissionCommand = v
	return s
}

func (s *DeleteRowPermissionRequest) SetOpTenantId(v int64) *DeleteRowPermissionRequest {
	s.OpTenantId = &v
	return s
}

func (s *DeleteRowPermissionRequest) Validate() error {
	return dara.Validate(s)
}

type DeleteRowPermissionRequestDeleteRowPermissionCommand struct {
	// This parameter is required.
	//
	// example:
	//
	// 300001234
	RowPermissionId *int64 `json:"RowPermissionId,omitempty" xml:"RowPermissionId,omitempty"`
}

func (s DeleteRowPermissionRequestDeleteRowPermissionCommand) String() string {
	return dara.Prettify(s)
}

func (s DeleteRowPermissionRequestDeleteRowPermissionCommand) GoString() string {
	return s.String()
}

func (s *DeleteRowPermissionRequestDeleteRowPermissionCommand) GetRowPermissionId() *int64 {
	return s.RowPermissionId
}

func (s *DeleteRowPermissionRequestDeleteRowPermissionCommand) SetRowPermissionId(v int64) *DeleteRowPermissionRequestDeleteRowPermissionCommand {
	s.RowPermissionId = &v
	return s
}

func (s *DeleteRowPermissionRequestDeleteRowPermissionCommand) Validate() error {
	return dara.Validate(s)
}
