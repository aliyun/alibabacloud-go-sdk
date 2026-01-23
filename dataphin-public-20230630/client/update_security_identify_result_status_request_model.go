// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateSecurityIdentifyResultStatusRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOpTenantId(v int64) *UpdateSecurityIdentifyResultStatusRequest
	GetOpTenantId() *int64
	SetUpdateCommand(v *UpdateSecurityIdentifyResultStatusRequestUpdateCommand) *UpdateSecurityIdentifyResultStatusRequest
	GetUpdateCommand() *UpdateSecurityIdentifyResultStatusRequestUpdateCommand
}

type UpdateSecurityIdentifyResultStatusRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
	// This parameter is required.
	UpdateCommand *UpdateSecurityIdentifyResultStatusRequestUpdateCommand `json:"UpdateCommand,omitempty" xml:"UpdateCommand,omitempty" type:"Struct"`
}

func (s UpdateSecurityIdentifyResultStatusRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateSecurityIdentifyResultStatusRequest) GoString() string {
	return s.String()
}

func (s *UpdateSecurityIdentifyResultStatusRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *UpdateSecurityIdentifyResultStatusRequest) GetUpdateCommand() *UpdateSecurityIdentifyResultStatusRequestUpdateCommand {
	return s.UpdateCommand
}

func (s *UpdateSecurityIdentifyResultStatusRequest) SetOpTenantId(v int64) *UpdateSecurityIdentifyResultStatusRequest {
	s.OpTenantId = &v
	return s
}

func (s *UpdateSecurityIdentifyResultStatusRequest) SetUpdateCommand(v *UpdateSecurityIdentifyResultStatusRequestUpdateCommand) *UpdateSecurityIdentifyResultStatusRequest {
	s.UpdateCommand = v
	return s
}

func (s *UpdateSecurityIdentifyResultStatusRequest) Validate() error {
	if s.UpdateCommand != nil {
		if err := s.UpdateCommand.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateSecurityIdentifyResultStatusRequestUpdateCommand struct {
	// This parameter is required.
	Enable *bool `json:"Enable,omitempty" xml:"Enable,omitempty"`
	// This parameter is required.
	IdentifyResultIdList []*int64 `json:"IdentifyResultIdList,omitempty" xml:"IdentifyResultIdList,omitempty" type:"Repeated"`
}

func (s UpdateSecurityIdentifyResultStatusRequestUpdateCommand) String() string {
	return dara.Prettify(s)
}

func (s UpdateSecurityIdentifyResultStatusRequestUpdateCommand) GoString() string {
	return s.String()
}

func (s *UpdateSecurityIdentifyResultStatusRequestUpdateCommand) GetEnable() *bool {
	return s.Enable
}

func (s *UpdateSecurityIdentifyResultStatusRequestUpdateCommand) GetIdentifyResultIdList() []*int64 {
	return s.IdentifyResultIdList
}

func (s *UpdateSecurityIdentifyResultStatusRequestUpdateCommand) SetEnable(v bool) *UpdateSecurityIdentifyResultStatusRequestUpdateCommand {
	s.Enable = &v
	return s
}

func (s *UpdateSecurityIdentifyResultStatusRequestUpdateCommand) SetIdentifyResultIdList(v []*int64) *UpdateSecurityIdentifyResultStatusRequestUpdateCommand {
	s.IdentifyResultIdList = v
	return s
}

func (s *UpdateSecurityIdentifyResultStatusRequestUpdateCommand) Validate() error {
	return dara.Validate(s)
}
