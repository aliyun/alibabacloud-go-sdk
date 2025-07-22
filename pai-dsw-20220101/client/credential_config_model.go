// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCredentialConfig interface {
	dara.Model
	String() string
	GoString() string
	SetAliyunEnvRoleKey(v string) *CredentialConfig
	GetAliyunEnvRoleKey() *string
	SetConfigs(v []*CredentialConfigConfigs) *CredentialConfig
	GetConfigs() []*CredentialConfigConfigs
	SetEnable(v bool) *CredentialConfig
	GetEnable() *bool
}

type CredentialConfig struct {
	// example:
	//
	// 0
	AliyunEnvRoleKey *string                    `json:"AliyunEnvRoleKey,omitempty" xml:"AliyunEnvRoleKey,omitempty"`
	Configs          []*CredentialConfigConfigs `json:"Configs,omitempty" xml:"Configs,omitempty" type:"Repeated"`
	// example:
	//
	// true
	Enable *bool `json:"Enable,omitempty" xml:"Enable,omitempty"`
}

func (s CredentialConfig) String() string {
	return dara.Prettify(s)
}

func (s CredentialConfig) GoString() string {
	return s.String()
}

func (s *CredentialConfig) GetAliyunEnvRoleKey() *string {
	return s.AliyunEnvRoleKey
}

func (s *CredentialConfig) GetConfigs() []*CredentialConfigConfigs {
	return s.Configs
}

func (s *CredentialConfig) GetEnable() *bool {
	return s.Enable
}

func (s *CredentialConfig) SetAliyunEnvRoleKey(v string) *CredentialConfig {
	s.AliyunEnvRoleKey = &v
	return s
}

func (s *CredentialConfig) SetConfigs(v []*CredentialConfigConfigs) *CredentialConfig {
	s.Configs = v
	return s
}

func (s *CredentialConfig) SetEnable(v bool) *CredentialConfig {
	s.Enable = &v
	return s
}

func (s *CredentialConfig) Validate() error {
	return dara.Validate(s)
}

type CredentialConfigConfigs struct {
	// This parameter is required.
	//
	// example:
	//
	// 0
	Key   *string                         `json:"Key,omitempty" xml:"Key,omitempty"`
	Roles []*CredentialConfigConfigsRoles `json:"Roles,omitempty" xml:"Roles,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// Role
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s CredentialConfigConfigs) String() string {
	return dara.Prettify(s)
}

func (s CredentialConfigConfigs) GoString() string {
	return s.String()
}

func (s *CredentialConfigConfigs) GetKey() *string {
	return s.Key
}

func (s *CredentialConfigConfigs) GetRoles() []*CredentialConfigConfigsRoles {
	return s.Roles
}

func (s *CredentialConfigConfigs) GetType() *string {
	return s.Type
}

func (s *CredentialConfigConfigs) SetKey(v string) *CredentialConfigConfigs {
	s.Key = &v
	return s
}

func (s *CredentialConfigConfigs) SetRoles(v []*CredentialConfigConfigsRoles) *CredentialConfigConfigs {
	s.Roles = v
	return s
}

func (s *CredentialConfigConfigs) SetType(v string) *CredentialConfigConfigs {
	s.Type = &v
	return s
}

func (s *CredentialConfigConfigs) Validate() error {
	return dara.Validate(s)
}

type CredentialConfigConfigsRoles struct {
	// example:
	//
	// 123******
	AssumeRoleFor *string `json:"AssumeRoleFor,omitempty" xml:"AssumeRoleFor,omitempty"`
	// example:
	//
	// {}
	Policy *string `json:"Policy,omitempty" xml:"Policy,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// acs:ram::123******:role/****
	RoleArn *string `json:"RoleArn,omitempty" xml:"RoleArn,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// service
	RoleType *string                               `json:"RoleType,omitempty" xml:"RoleType,omitempty"`
	UserInfo *CredentialConfigConfigsRolesUserInfo `json:"UserInfo,omitempty" xml:"UserInfo,omitempty" type:"Struct"`
}

func (s CredentialConfigConfigsRoles) String() string {
	return dara.Prettify(s)
}

func (s CredentialConfigConfigsRoles) GoString() string {
	return s.String()
}

func (s *CredentialConfigConfigsRoles) GetAssumeRoleFor() *string {
	return s.AssumeRoleFor
}

func (s *CredentialConfigConfigsRoles) GetPolicy() *string {
	return s.Policy
}

func (s *CredentialConfigConfigsRoles) GetRoleArn() *string {
	return s.RoleArn
}

func (s *CredentialConfigConfigsRoles) GetRoleType() *string {
	return s.RoleType
}

func (s *CredentialConfigConfigsRoles) GetUserInfo() *CredentialConfigConfigsRolesUserInfo {
	return s.UserInfo
}

func (s *CredentialConfigConfigsRoles) SetAssumeRoleFor(v string) *CredentialConfigConfigsRoles {
	s.AssumeRoleFor = &v
	return s
}

func (s *CredentialConfigConfigsRoles) SetPolicy(v string) *CredentialConfigConfigsRoles {
	s.Policy = &v
	return s
}

func (s *CredentialConfigConfigsRoles) SetRoleArn(v string) *CredentialConfigConfigsRoles {
	s.RoleArn = &v
	return s
}

func (s *CredentialConfigConfigsRoles) SetRoleType(v string) *CredentialConfigConfigsRoles {
	s.RoleType = &v
	return s
}

func (s *CredentialConfigConfigsRoles) SetUserInfo(v *CredentialConfigConfigsRolesUserInfo) *CredentialConfigConfigsRoles {
	s.UserInfo = v
	return s
}

func (s *CredentialConfigConfigsRoles) Validate() error {
	return dara.Validate(s)
}

type CredentialConfigConfigsRolesUserInfo struct {
	// example:
	//
	// LT********
	AccessKeyId *string `json:"AccessKeyId,omitempty" xml:"AccessKeyId,omitempty"`
	// example:
	//
	// 456******
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// ********
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	// example:
	//
	// sub
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s CredentialConfigConfigsRolesUserInfo) String() string {
	return dara.Prettify(s)
}

func (s CredentialConfigConfigsRolesUserInfo) GoString() string {
	return s.String()
}

func (s *CredentialConfigConfigsRolesUserInfo) GetAccessKeyId() *string {
	return s.AccessKeyId
}

func (s *CredentialConfigConfigsRolesUserInfo) GetId() *string {
	return s.Id
}

func (s *CredentialConfigConfigsRolesUserInfo) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *CredentialConfigConfigsRolesUserInfo) GetType() *string {
	return s.Type
}

func (s *CredentialConfigConfigsRolesUserInfo) SetAccessKeyId(v string) *CredentialConfigConfigsRolesUserInfo {
	s.AccessKeyId = &v
	return s
}

func (s *CredentialConfigConfigsRolesUserInfo) SetId(v string) *CredentialConfigConfigsRolesUserInfo {
	s.Id = &v
	return s
}

func (s *CredentialConfigConfigsRolesUserInfo) SetSecurityToken(v string) *CredentialConfigConfigsRolesUserInfo {
	s.SecurityToken = &v
	return s
}

func (s *CredentialConfigConfigsRolesUserInfo) SetType(v string) *CredentialConfigConfigsRolesUserInfo {
	s.Type = &v
	return s
}

func (s *CredentialConfigConfigsRolesUserInfo) Validate() error {
	return dara.Validate(s)
}
