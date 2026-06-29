// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSimpleTenant interface {
	dara.Model
	String() string
	GoString() string
	SetCreator(v *SimpleUser) *SimpleTenant
	GetCreator() *SimpleUser
	SetDescription(v string) *SimpleTenant
	GetDescription() *string
	SetGmtCreateTime(v string) *SimpleTenant
	GetGmtCreateTime() *string
	SetGmtModifiedTime(v string) *SimpleTenant
	GetGmtModifiedTime() *string
	SetModifier(v *SimpleUser) *SimpleTenant
	GetModifier() *SimpleUser
	SetRole(v string) *SimpleTenant
	GetRole() *string
	SetTenantId(v string) *SimpleTenant
	GetTenantId() *string
	SetTenantName(v string) *SimpleTenant
	GetTenantName() *string
	SetUUID(v string) *SimpleTenant
	GetUUID() *string
}

type SimpleTenant struct {
	// Creator
	Creator *SimpleUser `json:"Creator,omitempty" xml:"Creator,omitempty"`
	// Description
	//
	// example:
	//
	// 这是一个测试租户
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// Creation Time
	//
	// example:
	//
	// 2021-07-07 16:09:20
	GmtCreateTime *string `json:"GmtCreateTime,omitempty" xml:"GmtCreateTime,omitempty"`
	// Updated At
	//
	// example:
	//
	// 2021-07-07 16:09:20
	GmtModifiedTime *string `json:"GmtModifiedTime,omitempty" xml:"GmtModifiedTime,omitempty"`
	// Updated By Information
	Modifier *SimpleUser `json:"Modifier,omitempty" xml:"Modifier,omitempty"`
	// Role Information
	//
	// example:
	//
	// ADMIN
	Role *string `json:"Role,omitempty" xml:"Role,omitempty"`
	// Tenant ID
	//
	// example:
	//
	// GA***W134
	TenantId *string `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
	// Tenant Name
	//
	// example:
	//
	// demo
	TenantName *string `json:"TenantName,omitempty" xml:"TenantName,omitempty"`
	// Unique Identifier
	//
	// example:
	//
	// paiworkspace-0001
	UUID *string `json:"UUID,omitempty" xml:"UUID,omitempty"`
}

func (s SimpleTenant) String() string {
	return dara.Prettify(s)
}

func (s SimpleTenant) GoString() string {
	return s.String()
}

func (s *SimpleTenant) GetCreator() *SimpleUser {
	return s.Creator
}

func (s *SimpleTenant) GetDescription() *string {
	return s.Description
}

func (s *SimpleTenant) GetGmtCreateTime() *string {
	return s.GmtCreateTime
}

func (s *SimpleTenant) GetGmtModifiedTime() *string {
	return s.GmtModifiedTime
}

func (s *SimpleTenant) GetModifier() *SimpleUser {
	return s.Modifier
}

func (s *SimpleTenant) GetRole() *string {
	return s.Role
}

func (s *SimpleTenant) GetTenantId() *string {
	return s.TenantId
}

func (s *SimpleTenant) GetTenantName() *string {
	return s.TenantName
}

func (s *SimpleTenant) GetUUID() *string {
	return s.UUID
}

func (s *SimpleTenant) SetCreator(v *SimpleUser) *SimpleTenant {
	s.Creator = v
	return s
}

func (s *SimpleTenant) SetDescription(v string) *SimpleTenant {
	s.Description = &v
	return s
}

func (s *SimpleTenant) SetGmtCreateTime(v string) *SimpleTenant {
	s.GmtCreateTime = &v
	return s
}

func (s *SimpleTenant) SetGmtModifiedTime(v string) *SimpleTenant {
	s.GmtModifiedTime = &v
	return s
}

func (s *SimpleTenant) SetModifier(v *SimpleUser) *SimpleTenant {
	s.Modifier = v
	return s
}

func (s *SimpleTenant) SetRole(v string) *SimpleTenant {
	s.Role = &v
	return s
}

func (s *SimpleTenant) SetTenantId(v string) *SimpleTenant {
	s.TenantId = &v
	return s
}

func (s *SimpleTenant) SetTenantName(v string) *SimpleTenant {
	s.TenantName = &v
	return s
}

func (s *SimpleTenant) SetUUID(v string) *SimpleTenant {
	s.UUID = &v
	return s
}

func (s *SimpleTenant) Validate() error {
	if s.Creator != nil {
		if err := s.Creator.Validate(); err != nil {
			return err
		}
	}
	if s.Modifier != nil {
		if err := s.Modifier.Validate(); err != nil {
			return err
		}
	}
	return nil
}
