// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateRegistryModuleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAcl(v string) *CreateRegistryModuleRequest
	GetAcl() *string
	SetClientToken(v string) *CreateRegistryModuleRequest
	GetClientToken() *string
	SetDescription(v string) *CreateRegistryModuleRequest
	GetDescription() *string
	SetModuleName(v string) *CreateRegistryModuleRequest
	GetModuleName() *string
	SetNamespaceName(v string) *CreateRegistryModuleRequest
	GetNamespaceName() *string
	SetProvider(v string) *CreateRegistryModuleRequest
	GetProvider() *string
	SetType(v string) *CreateRegistryModuleRequest
	GetType() *string
}

type CreateRegistryModuleRequest struct {
	// example:
	//
	// private
	Acl *string `json:"acl,omitempty" xml:"acl,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// a65451293e64979ba7a4b573950217fe
	ClientToken *string `json:"clientToken,omitempty" xml:"clientToken,omitempty"`
	// example:
	//
	// test
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// vpc
	ModuleName *string `json:"moduleName,omitempty" xml:"moduleName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// test
	NamespaceName *string `json:"namespaceName,omitempty" xml:"namespaceName,omitempty"`
	// example:
	//
	// alicloud
	Provider *string `json:"provider,omitempty" xml:"provider,omitempty"`
	// example:
	//
	// self
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s CreateRegistryModuleRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateRegistryModuleRequest) GoString() string {
	return s.String()
}

func (s *CreateRegistryModuleRequest) GetAcl() *string {
	return s.Acl
}

func (s *CreateRegistryModuleRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateRegistryModuleRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateRegistryModuleRequest) GetModuleName() *string {
	return s.ModuleName
}

func (s *CreateRegistryModuleRequest) GetNamespaceName() *string {
	return s.NamespaceName
}

func (s *CreateRegistryModuleRequest) GetProvider() *string {
	return s.Provider
}

func (s *CreateRegistryModuleRequest) GetType() *string {
	return s.Type
}

func (s *CreateRegistryModuleRequest) SetAcl(v string) *CreateRegistryModuleRequest {
	s.Acl = &v
	return s
}

func (s *CreateRegistryModuleRequest) SetClientToken(v string) *CreateRegistryModuleRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateRegistryModuleRequest) SetDescription(v string) *CreateRegistryModuleRequest {
	s.Description = &v
	return s
}

func (s *CreateRegistryModuleRequest) SetModuleName(v string) *CreateRegistryModuleRequest {
	s.ModuleName = &v
	return s
}

func (s *CreateRegistryModuleRequest) SetNamespaceName(v string) *CreateRegistryModuleRequest {
	s.NamespaceName = &v
	return s
}

func (s *CreateRegistryModuleRequest) SetProvider(v string) *CreateRegistryModuleRequest {
	s.Provider = &v
	return s
}

func (s *CreateRegistryModuleRequest) SetType(v string) *CreateRegistryModuleRequest {
	s.Type = &v
	return s
}

func (s *CreateRegistryModuleRequest) Validate() error {
	return dara.Validate(s)
}
