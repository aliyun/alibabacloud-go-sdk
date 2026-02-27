// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddDataServiceAppMemberRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAddCommand(v *AddDataServiceAppMemberRequestAddCommand) *AddDataServiceAppMemberRequest
	GetAddCommand() *AddDataServiceAppMemberRequestAddCommand
	SetOpTenantId(v int64) *AddDataServiceAppMemberRequest
	GetOpTenantId() *int64
}

type AddDataServiceAppMemberRequest struct {
	// This parameter is required.
	AddCommand *AddDataServiceAppMemberRequestAddCommand `json:"AddCommand,omitempty" xml:"AddCommand,omitempty" type:"Struct"`
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
}

func (s AddDataServiceAppMemberRequest) String() string {
	return dara.Prettify(s)
}

func (s AddDataServiceAppMemberRequest) GoString() string {
	return s.String()
}

func (s *AddDataServiceAppMemberRequest) GetAddCommand() *AddDataServiceAppMemberRequestAddCommand {
	return s.AddCommand
}

func (s *AddDataServiceAppMemberRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *AddDataServiceAppMemberRequest) SetAddCommand(v *AddDataServiceAppMemberRequestAddCommand) *AddDataServiceAppMemberRequest {
	s.AddCommand = v
	return s
}

func (s *AddDataServiceAppMemberRequest) SetOpTenantId(v int64) *AddDataServiceAppMemberRequest {
	s.OpTenantId = &v
	return s
}

func (s *AddDataServiceAppMemberRequest) Validate() error {
	if s.AddCommand != nil {
		if err := s.AddCommand.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type AddDataServiceAppMemberRequestAddCommand struct {
	// This parameter is required.
	//
	// example:
	//
	// 200000000
	AppId *int32 `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// This parameter is required.
	MemberList []*AddDataServiceAppMemberRequestAddCommandMemberList `json:"MemberList,omitempty" xml:"MemberList,omitempty" type:"Repeated"`
}

func (s AddDataServiceAppMemberRequestAddCommand) String() string {
	return dara.Prettify(s)
}

func (s AddDataServiceAppMemberRequestAddCommand) GoString() string {
	return s.String()
}

func (s *AddDataServiceAppMemberRequestAddCommand) GetAppId() *int32 {
	return s.AppId
}

func (s *AddDataServiceAppMemberRequestAddCommand) GetMemberList() []*AddDataServiceAppMemberRequestAddCommandMemberList {
	return s.MemberList
}

func (s *AddDataServiceAppMemberRequestAddCommand) SetAppId(v int32) *AddDataServiceAppMemberRequestAddCommand {
	s.AppId = &v
	return s
}

func (s *AddDataServiceAppMemberRequestAddCommand) SetMemberList(v []*AddDataServiceAppMemberRequestAddCommandMemberList) *AddDataServiceAppMemberRequestAddCommand {
	s.MemberList = v
	return s
}

func (s *AddDataServiceAppMemberRequestAddCommand) Validate() error {
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

type AddDataServiceAppMemberRequestAddCommandMemberList struct {
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

func (s AddDataServiceAppMemberRequestAddCommandMemberList) String() string {
	return dara.Prettify(s)
}

func (s AddDataServiceAppMemberRequestAddCommandMemberList) GoString() string {
	return s.String()
}

func (s *AddDataServiceAppMemberRequestAddCommandMemberList) GetEffectiveEnd() *string {
	return s.EffectiveEnd
}

func (s *AddDataServiceAppMemberRequestAddCommandMemberList) GetUserId() *string {
	return s.UserId
}

func (s *AddDataServiceAppMemberRequestAddCommandMemberList) SetEffectiveEnd(v string) *AddDataServiceAppMemberRequestAddCommandMemberList {
	s.EffectiveEnd = &v
	return s
}

func (s *AddDataServiceAppMemberRequestAddCommandMemberList) SetUserId(v string) *AddDataServiceAppMemberRequestAddCommandMemberList {
	s.UserId = &v
	return s
}

func (s *AddDataServiceAppMemberRequestAddCommandMemberList) Validate() error {
	return dara.Validate(s)
}
