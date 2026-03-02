// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iResourceCreateCmd interface {
	dara.Model
	String() string
	GoString() string
	SetAddress(v string) *ResourceCreateCmd
	GetAddress() *string
	SetCompanyId(v int64) *ResourceCreateCmd
	GetCompanyId() *int64
	SetConfiguration(v string) *ResourceCreateCmd
	GetConfiguration() *string
	SetCredentials(v string) *ResourceCreateCmd
	GetCredentials() *string
	SetDescription(v string) *ResourceCreateCmd
	GetDescription() *string
	SetEnv(v string) *ResourceCreateCmd
	GetEnv() *string
	SetName(v string) *ResourceCreateCmd
	GetName() *string
	SetNamespace(v string) *ResourceCreateCmd
	GetNamespace() *string
	SetProductId(v int64) *ResourceCreateCmd
	GetProductId() *int64
	SetType(v string) *ResourceCreateCmd
	GetType() *string
	SetUseScope(v string) *ResourceCreateCmd
	GetUseScope() *string
	SetUseType(v string) *ResourceCreateCmd
	GetUseType() *string
}

type ResourceCreateCmd struct {
	// This parameter is required.
	//
	// example:
	//
	// http://
	Address *string `json:"address,omitempty" xml:"address,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	CompanyId     *int64  `json:"companyId,omitempty" xml:"companyId,omitempty"`
	Configuration *string `json:"configuration,omitempty" xml:"configuration,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// {ak:xxxx,sk:xxx}
	Credentials *string `json:"credentials,omitempty" xml:"credentials,omitempty"`
	// example:
	//
	// 用于统一管理资源
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// test
	Env *string `json:"env,omitempty" xml:"env,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// Redis
	Name      *string `json:"name,omitempty" xml:"name,omitempty"`
	Namespace *string `json:"namespace,omitempty" xml:"namespace,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	ProductId *int64 `json:"productId,omitempty" xml:"productId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// Redis
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// Inner
	UseScope *string `json:"useScope,omitempty" xml:"useScope,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// Share
	UseType *string `json:"useType,omitempty" xml:"useType,omitempty"`
}

func (s ResourceCreateCmd) String() string {
	return dara.Prettify(s)
}

func (s ResourceCreateCmd) GoString() string {
	return s.String()
}

func (s *ResourceCreateCmd) GetAddress() *string {
	return s.Address
}

func (s *ResourceCreateCmd) GetCompanyId() *int64 {
	return s.CompanyId
}

func (s *ResourceCreateCmd) GetConfiguration() *string {
	return s.Configuration
}

func (s *ResourceCreateCmd) GetCredentials() *string {
	return s.Credentials
}

func (s *ResourceCreateCmd) GetDescription() *string {
	return s.Description
}

func (s *ResourceCreateCmd) GetEnv() *string {
	return s.Env
}

func (s *ResourceCreateCmd) GetName() *string {
	return s.Name
}

func (s *ResourceCreateCmd) GetNamespace() *string {
	return s.Namespace
}

func (s *ResourceCreateCmd) GetProductId() *int64 {
	return s.ProductId
}

func (s *ResourceCreateCmd) GetType() *string {
	return s.Type
}

func (s *ResourceCreateCmd) GetUseScope() *string {
	return s.UseScope
}

func (s *ResourceCreateCmd) GetUseType() *string {
	return s.UseType
}

func (s *ResourceCreateCmd) SetAddress(v string) *ResourceCreateCmd {
	s.Address = &v
	return s
}

func (s *ResourceCreateCmd) SetCompanyId(v int64) *ResourceCreateCmd {
	s.CompanyId = &v
	return s
}

func (s *ResourceCreateCmd) SetConfiguration(v string) *ResourceCreateCmd {
	s.Configuration = &v
	return s
}

func (s *ResourceCreateCmd) SetCredentials(v string) *ResourceCreateCmd {
	s.Credentials = &v
	return s
}

func (s *ResourceCreateCmd) SetDescription(v string) *ResourceCreateCmd {
	s.Description = &v
	return s
}

func (s *ResourceCreateCmd) SetEnv(v string) *ResourceCreateCmd {
	s.Env = &v
	return s
}

func (s *ResourceCreateCmd) SetName(v string) *ResourceCreateCmd {
	s.Name = &v
	return s
}

func (s *ResourceCreateCmd) SetNamespace(v string) *ResourceCreateCmd {
	s.Namespace = &v
	return s
}

func (s *ResourceCreateCmd) SetProductId(v int64) *ResourceCreateCmd {
	s.ProductId = &v
	return s
}

func (s *ResourceCreateCmd) SetType(v string) *ResourceCreateCmd {
	s.Type = &v
	return s
}

func (s *ResourceCreateCmd) SetUseScope(v string) *ResourceCreateCmd {
	s.UseScope = &v
	return s
}

func (s *ResourceCreateCmd) SetUseType(v string) *ResourceCreateCmd {
	s.UseType = &v
	return s
}

func (s *ResourceCreateCmd) Validate() error {
	return dara.Validate(s)
}
