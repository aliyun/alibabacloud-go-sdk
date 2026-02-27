// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateDataServiceAppMemberRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOpTenantId(v int64) *UpdateDataServiceAppMemberRequest
	GetOpTenantId() *int64
	SetUpdateCommand(v *UpdateDataServiceAppMemberRequestUpdateCommand) *UpdateDataServiceAppMemberRequest
	GetUpdateCommand() *UpdateDataServiceAppMemberRequestUpdateCommand
}

type UpdateDataServiceAppMemberRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
	// This parameter is required.
	UpdateCommand *UpdateDataServiceAppMemberRequestUpdateCommand `json:"UpdateCommand,omitempty" xml:"UpdateCommand,omitempty" type:"Struct"`
}

func (s UpdateDataServiceAppMemberRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateDataServiceAppMemberRequest) GoString() string {
	return s.String()
}

func (s *UpdateDataServiceAppMemberRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *UpdateDataServiceAppMemberRequest) GetUpdateCommand() *UpdateDataServiceAppMemberRequestUpdateCommand {
	return s.UpdateCommand
}

func (s *UpdateDataServiceAppMemberRequest) SetOpTenantId(v int64) *UpdateDataServiceAppMemberRequest {
	s.OpTenantId = &v
	return s
}

func (s *UpdateDataServiceAppMemberRequest) SetUpdateCommand(v *UpdateDataServiceAppMemberRequestUpdateCommand) *UpdateDataServiceAppMemberRequest {
	s.UpdateCommand = v
	return s
}

func (s *UpdateDataServiceAppMemberRequest) Validate() error {
	if s.UpdateCommand != nil {
		if err := s.UpdateCommand.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateDataServiceAppMemberRequestUpdateCommand struct {
	// This parameter is required.
	//
	// example:
	//
	// 200000000
	AppId *int32 `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// This parameter is required.
	MemberList []*UpdateDataServiceAppMemberRequestUpdateCommandMemberList `json:"MemberList,omitempty" xml:"MemberList,omitempty" type:"Repeated"`
}

func (s UpdateDataServiceAppMemberRequestUpdateCommand) String() string {
	return dara.Prettify(s)
}

func (s UpdateDataServiceAppMemberRequestUpdateCommand) GoString() string {
	return s.String()
}

func (s *UpdateDataServiceAppMemberRequestUpdateCommand) GetAppId() *int32 {
	return s.AppId
}

func (s *UpdateDataServiceAppMemberRequestUpdateCommand) GetMemberList() []*UpdateDataServiceAppMemberRequestUpdateCommandMemberList {
	return s.MemberList
}

func (s *UpdateDataServiceAppMemberRequestUpdateCommand) SetAppId(v int32) *UpdateDataServiceAppMemberRequestUpdateCommand {
	s.AppId = &v
	return s
}

func (s *UpdateDataServiceAppMemberRequestUpdateCommand) SetMemberList(v []*UpdateDataServiceAppMemberRequestUpdateCommandMemberList) *UpdateDataServiceAppMemberRequestUpdateCommand {
	s.MemberList = v
	return s
}

func (s *UpdateDataServiceAppMemberRequestUpdateCommand) Validate() error {
	if s.MemberList != nil {
		for _, item := range s.MemberList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type UpdateDataServiceAppMemberRequestUpdateCommandMemberList struct {
	// This parameter is required.
	//
	// example:
	//
	// 2026-12-12
	EffectiveEnd *string `json:"EffectiveEnd,omitempty" xml:"EffectiveEnd,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 12345
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s UpdateDataServiceAppMemberRequestUpdateCommandMemberList) String() string {
	return dara.Prettify(s)
}

func (s UpdateDataServiceAppMemberRequestUpdateCommandMemberList) GoString() string {
	return s.String()
}

func (s *UpdateDataServiceAppMemberRequestUpdateCommandMemberList) GetEffectiveEnd() *string {
	return s.EffectiveEnd
}

func (s *UpdateDataServiceAppMemberRequestUpdateCommandMemberList) GetUserId() *string {
	return s.UserId
}

func (s *UpdateDataServiceAppMemberRequestUpdateCommandMemberList) SetEffectiveEnd(v string) *UpdateDataServiceAppMemberRequestUpdateCommandMemberList {
	s.EffectiveEnd = &v
	return s
}

func (s *UpdateDataServiceAppMemberRequestUpdateCommandMemberList) SetUserId(v string) *UpdateDataServiceAppMemberRequestUpdateCommandMemberList {
	s.UserId = &v
	return s
}

func (s *UpdateDataServiceAppMemberRequestUpdateCommandMemberList) Validate() error {
	return dara.Validate(s)
}
