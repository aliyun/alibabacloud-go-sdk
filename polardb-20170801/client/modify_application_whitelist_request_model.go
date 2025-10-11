// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyApplicationWhitelistRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApplicationId(v string) *ModifyApplicationWhitelistRequest
	GetApplicationId() *string
	SetComponentId(v string) *ModifyApplicationWhitelistRequest
	GetComponentId() *string
	SetModifyMode(v string) *ModifyApplicationWhitelistRequest
	GetModifyMode() *string
	SetSecurityGroups(v string) *ModifyApplicationWhitelistRequest
	GetSecurityGroups() *string
	SetSecurityIPArrayName(v string) *ModifyApplicationWhitelistRequest
	GetSecurityIPArrayName() *string
	SetSecurityIPList(v string) *ModifyApplicationWhitelistRequest
	GetSecurityIPList() *string
}

type ModifyApplicationWhitelistRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// pa-**************
	ApplicationId *string `json:"ApplicationId,omitempty" xml:"ApplicationId,omitempty"`
	// example:
	//
	// pac-*******************
	ComponentId *string `json:"ComponentId,omitempty" xml:"ComponentId,omitempty"`
	// example:
	//
	// Append
	ModifyMode *string `json:"ModifyMode,omitempty" xml:"ModifyMode,omitempty"`
	// example:
	//
	// sg-**************
	SecurityGroups *string `json:"SecurityGroups,omitempty" xml:"SecurityGroups,omitempty"`
	// example:
	//
	// default
	SecurityIPArrayName *string `json:"SecurityIPArrayName,omitempty" xml:"SecurityIPArrayName,omitempty"`
	// example:
	//
	// 127.0.0.1,172.17.0.0/24
	SecurityIPList *string `json:"SecurityIPList,omitempty" xml:"SecurityIPList,omitempty"`
}

func (s ModifyApplicationWhitelistRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyApplicationWhitelistRequest) GoString() string {
	return s.String()
}

func (s *ModifyApplicationWhitelistRequest) GetApplicationId() *string {
	return s.ApplicationId
}

func (s *ModifyApplicationWhitelistRequest) GetComponentId() *string {
	return s.ComponentId
}

func (s *ModifyApplicationWhitelistRequest) GetModifyMode() *string {
	return s.ModifyMode
}

func (s *ModifyApplicationWhitelistRequest) GetSecurityGroups() *string {
	return s.SecurityGroups
}

func (s *ModifyApplicationWhitelistRequest) GetSecurityIPArrayName() *string {
	return s.SecurityIPArrayName
}

func (s *ModifyApplicationWhitelistRequest) GetSecurityIPList() *string {
	return s.SecurityIPList
}

func (s *ModifyApplicationWhitelistRequest) SetApplicationId(v string) *ModifyApplicationWhitelistRequest {
	s.ApplicationId = &v
	return s
}

func (s *ModifyApplicationWhitelistRequest) SetComponentId(v string) *ModifyApplicationWhitelistRequest {
	s.ComponentId = &v
	return s
}

func (s *ModifyApplicationWhitelistRequest) SetModifyMode(v string) *ModifyApplicationWhitelistRequest {
	s.ModifyMode = &v
	return s
}

func (s *ModifyApplicationWhitelistRequest) SetSecurityGroups(v string) *ModifyApplicationWhitelistRequest {
	s.SecurityGroups = &v
	return s
}

func (s *ModifyApplicationWhitelistRequest) SetSecurityIPArrayName(v string) *ModifyApplicationWhitelistRequest {
	s.SecurityIPArrayName = &v
	return s
}

func (s *ModifyApplicationWhitelistRequest) SetSecurityIPList(v string) *ModifyApplicationWhitelistRequest {
	s.SecurityIPList = &v
	return s
}

func (s *ModifyApplicationWhitelistRequest) Validate() error {
	return dara.Validate(s)
}
