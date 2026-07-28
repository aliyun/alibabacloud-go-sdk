// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateRegistryModuleAttributeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAcl(v string) *UpdateRegistryModuleAttributeRequest
	GetAcl() *string
	SetClientToken(v string) *UpdateRegistryModuleAttributeRequest
	GetClientToken() *string
	SetDescription(v string) *UpdateRegistryModuleAttributeRequest
	GetDescription() *string
}

type UpdateRegistryModuleAttributeRequest struct {
	// The access permission. Valid values:
	//
	// - private: private.
	//
	// example:
	//
	// private
	Acl *string `json:"acl,omitempty" xml:"acl,omitempty"`
	// The idempotence token. Format: [0-9a-zA-Z-]{1,64}. Use a UUID.
	//
	// This parameter is required.
	//
	// example:
	//
	// a65451293e64979ba7a4b573950217fe
	ClientToken *string `json:"clientToken,omitempty" xml:"clientToken,omitempty"`
	// The description of the Registry template.
	//
	// example:
	//
	// description
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
}

func (s UpdateRegistryModuleAttributeRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateRegistryModuleAttributeRequest) GoString() string {
	return s.String()
}

func (s *UpdateRegistryModuleAttributeRequest) GetAcl() *string {
	return s.Acl
}

func (s *UpdateRegistryModuleAttributeRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *UpdateRegistryModuleAttributeRequest) GetDescription() *string {
	return s.Description
}

func (s *UpdateRegistryModuleAttributeRequest) SetAcl(v string) *UpdateRegistryModuleAttributeRequest {
	s.Acl = &v
	return s
}

func (s *UpdateRegistryModuleAttributeRequest) SetClientToken(v string) *UpdateRegistryModuleAttributeRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateRegistryModuleAttributeRequest) SetDescription(v string) *UpdateRegistryModuleAttributeRequest {
	s.Description = &v
	return s
}

func (s *UpdateRegistryModuleAttributeRequest) Validate() error {
	return dara.Validate(s)
}
