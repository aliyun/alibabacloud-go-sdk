// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iResourceUpdateCmd interface {
	dara.Model
	String() string
	GoString() string
	SetAddress(v string) *ResourceUpdateCmd
	GetAddress() *string
	SetConfiguration(v string) *ResourceUpdateCmd
	GetConfiguration() *string
	SetCredentials(v string) *ResourceUpdateCmd
	GetCredentials() *string
	SetDescription(v string) *ResourceUpdateCmd
	GetDescription() *string
	SetId(v int64) *ResourceUpdateCmd
	GetId() *int64
	SetNamespace(v string) *ResourceUpdateCmd
	GetNamespace() *string
	SetUseScope(v string) *ResourceUpdateCmd
	GetUseScope() *string
	SetUseType(v string) *ResourceUpdateCmd
	GetUseType() *string
}

type ResourceUpdateCmd struct {
	// example:
	//
	// http://
	Address       *string `json:"address,omitempty" xml:"address,omitempty"`
	Configuration *string `json:"configuration,omitempty" xml:"configuration,omitempty"`
	// example:
	//
	// {ak:xxxx,sk:xxx}
	Credentials *string `json:"credentials,omitempty" xml:"credentials,omitempty"`
	// example:
	//
	// 用于统一管理资源
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// example:
	//
	// 1
	Id        *int64  `json:"id,omitempty" xml:"id,omitempty"`
	Namespace *string `json:"namespace,omitempty" xml:"namespace,omitempty"`
	// example:
	//
	// Inner
	UseScope *string `json:"useScope,omitempty" xml:"useScope,omitempty"`
	// example:
	//
	// Share
	UseType *string `json:"useType,omitempty" xml:"useType,omitempty"`
}

func (s ResourceUpdateCmd) String() string {
	return dara.Prettify(s)
}

func (s ResourceUpdateCmd) GoString() string {
	return s.String()
}

func (s *ResourceUpdateCmd) GetAddress() *string {
	return s.Address
}

func (s *ResourceUpdateCmd) GetConfiguration() *string {
	return s.Configuration
}

func (s *ResourceUpdateCmd) GetCredentials() *string {
	return s.Credentials
}

func (s *ResourceUpdateCmd) GetDescription() *string {
	return s.Description
}

func (s *ResourceUpdateCmd) GetId() *int64 {
	return s.Id
}

func (s *ResourceUpdateCmd) GetNamespace() *string {
	return s.Namespace
}

func (s *ResourceUpdateCmd) GetUseScope() *string {
	return s.UseScope
}

func (s *ResourceUpdateCmd) GetUseType() *string {
	return s.UseType
}

func (s *ResourceUpdateCmd) SetAddress(v string) *ResourceUpdateCmd {
	s.Address = &v
	return s
}

func (s *ResourceUpdateCmd) SetConfiguration(v string) *ResourceUpdateCmd {
	s.Configuration = &v
	return s
}

func (s *ResourceUpdateCmd) SetCredentials(v string) *ResourceUpdateCmd {
	s.Credentials = &v
	return s
}

func (s *ResourceUpdateCmd) SetDescription(v string) *ResourceUpdateCmd {
	s.Description = &v
	return s
}

func (s *ResourceUpdateCmd) SetId(v int64) *ResourceUpdateCmd {
	s.Id = &v
	return s
}

func (s *ResourceUpdateCmd) SetNamespace(v string) *ResourceUpdateCmd {
	s.Namespace = &v
	return s
}

func (s *ResourceUpdateCmd) SetUseScope(v string) *ResourceUpdateCmd {
	s.UseScope = &v
	return s
}

func (s *ResourceUpdateCmd) SetUseType(v string) *ResourceUpdateCmd {
	s.UseType = &v
	return s
}

func (s *ResourceUpdateCmd) Validate() error {
	return dara.Validate(s)
}
