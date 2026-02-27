// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateDataServiceAppGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOpTenantId(v int64) *UpdateDataServiceAppGroupRequest
	GetOpTenantId() *int64
	SetUpdateCommand(v *UpdateDataServiceAppGroupRequestUpdateCommand) *UpdateDataServiceAppGroupRequest
	GetUpdateCommand() *UpdateDataServiceAppGroupRequestUpdateCommand
}

type UpdateDataServiceAppGroupRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
	// This parameter is required.
	UpdateCommand *UpdateDataServiceAppGroupRequestUpdateCommand `json:"UpdateCommand,omitempty" xml:"UpdateCommand,omitempty" type:"Struct"`
}

func (s UpdateDataServiceAppGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateDataServiceAppGroupRequest) GoString() string {
	return s.String()
}

func (s *UpdateDataServiceAppGroupRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *UpdateDataServiceAppGroupRequest) GetUpdateCommand() *UpdateDataServiceAppGroupRequestUpdateCommand {
	return s.UpdateCommand
}

func (s *UpdateDataServiceAppGroupRequest) SetOpTenantId(v int64) *UpdateDataServiceAppGroupRequest {
	s.OpTenantId = &v
	return s
}

func (s *UpdateDataServiceAppGroupRequest) SetUpdateCommand(v *UpdateDataServiceAppGroupRequestUpdateCommand) *UpdateDataServiceAppGroupRequest {
	s.UpdateCommand = v
	return s
}

func (s *UpdateDataServiceAppGroupRequest) Validate() error {
	if s.UpdateCommand != nil {
		if err := s.UpdateCommand.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateDataServiceAppGroupRequestUpdateCommand struct {
	// This parameter is required.
	//
	// example:
	//
	// 12345
	GroupId *int32 `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 默认应用分组
	GroupName *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
}

func (s UpdateDataServiceAppGroupRequestUpdateCommand) String() string {
	return dara.Prettify(s)
}

func (s UpdateDataServiceAppGroupRequestUpdateCommand) GoString() string {
	return s.String()
}

func (s *UpdateDataServiceAppGroupRequestUpdateCommand) GetGroupId() *int32 {
	return s.GroupId
}

func (s *UpdateDataServiceAppGroupRequestUpdateCommand) GetGroupName() *string {
	return s.GroupName
}

func (s *UpdateDataServiceAppGroupRequestUpdateCommand) SetGroupId(v int32) *UpdateDataServiceAppGroupRequestUpdateCommand {
	s.GroupId = &v
	return s
}

func (s *UpdateDataServiceAppGroupRequestUpdateCommand) SetGroupName(v string) *UpdateDataServiceAppGroupRequestUpdateCommand {
	s.GroupName = &v
	return s
}

func (s *UpdateDataServiceAppGroupRequestUpdateCommand) Validate() error {
	return dara.Validate(s)
}
