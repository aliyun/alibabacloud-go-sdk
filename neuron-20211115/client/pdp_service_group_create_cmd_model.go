// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPdpServiceGroupCreateCmd interface {
	dara.Model
	String() string
	GoString() string
	SetAccountId(v string) *PdpServiceGroupCreateCmd
	GetAccountId() *string
	SetAlias(v string) *PdpServiceGroupCreateCmd
	GetAlias() *string
	SetDescription(v string) *PdpServiceGroupCreateCmd
	GetDescription() *string
	SetEnv(v string) *PdpServiceGroupCreateCmd
	GetEnv() *string
	SetGroupType(v string) *PdpServiceGroupCreateCmd
	GetGroupType() *string
	SetName(v string) *PdpServiceGroupCreateCmd
	GetName() *string
	SetServiceId(v int64) *PdpServiceGroupCreateCmd
	GetServiceId() *int64
}

type PdpServiceGroupCreateCmd struct {
	// example:
	//
	// 121321412341
	AccountId *string `json:"accountId,omitempty" xml:"accountId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 管理
	Alias *string `json:"alias,omitempty" xml:"alias,omitempty"`
	// example:
	//
	// 管理
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// TEST
	Env *string `json:"env,omitempty" xml:"env,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// CUSTOM
	GroupType *string `json:"groupType,omitempty" xml:"groupType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// employee
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	ServiceId *int64 `json:"serviceId,omitempty" xml:"serviceId,omitempty"`
}

func (s PdpServiceGroupCreateCmd) String() string {
	return dara.Prettify(s)
}

func (s PdpServiceGroupCreateCmd) GoString() string {
	return s.String()
}

func (s *PdpServiceGroupCreateCmd) GetAccountId() *string {
	return s.AccountId
}

func (s *PdpServiceGroupCreateCmd) GetAlias() *string {
	return s.Alias
}

func (s *PdpServiceGroupCreateCmd) GetDescription() *string {
	return s.Description
}

func (s *PdpServiceGroupCreateCmd) GetEnv() *string {
	return s.Env
}

func (s *PdpServiceGroupCreateCmd) GetGroupType() *string {
	return s.GroupType
}

func (s *PdpServiceGroupCreateCmd) GetName() *string {
	return s.Name
}

func (s *PdpServiceGroupCreateCmd) GetServiceId() *int64 {
	return s.ServiceId
}

func (s *PdpServiceGroupCreateCmd) SetAccountId(v string) *PdpServiceGroupCreateCmd {
	s.AccountId = &v
	return s
}

func (s *PdpServiceGroupCreateCmd) SetAlias(v string) *PdpServiceGroupCreateCmd {
	s.Alias = &v
	return s
}

func (s *PdpServiceGroupCreateCmd) SetDescription(v string) *PdpServiceGroupCreateCmd {
	s.Description = &v
	return s
}

func (s *PdpServiceGroupCreateCmd) SetEnv(v string) *PdpServiceGroupCreateCmd {
	s.Env = &v
	return s
}

func (s *PdpServiceGroupCreateCmd) SetGroupType(v string) *PdpServiceGroupCreateCmd {
	s.GroupType = &v
	return s
}

func (s *PdpServiceGroupCreateCmd) SetName(v string) *PdpServiceGroupCreateCmd {
	s.Name = &v
	return s
}

func (s *PdpServiceGroupCreateCmd) SetServiceId(v int64) *PdpServiceGroupCreateCmd {
	s.ServiceId = &v
	return s
}

func (s *PdpServiceGroupCreateCmd) Validate() error {
	return dara.Validate(s)
}
