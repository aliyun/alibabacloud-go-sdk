// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPublishRegistryModuleVersionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *PublishRegistryModuleVersionRequest
	GetClientToken() *string
	SetModuleName(v string) *PublishRegistryModuleVersionRequest
	GetModuleName() *string
	SetNamespaceName(v string) *PublishRegistryModuleVersionRequest
	GetNamespaceName() *string
	SetVersion(v string) *PublishRegistryModuleVersionRequest
	GetVersion() *string
}

type PublishRegistryModuleVersionRequest struct {
	// The idempotence token. Format: [0-9a-zA-Z-]{1,64}. Use a UUID.
	//
	// This parameter is required.
	//
	// example:
	//
	// a65451293e64979ba7a4b573950217fe
	ClientToken *string `json:"clientToken,omitempty" xml:"clientToken,omitempty"`
	// The template name.
	//
	// This parameter is required.
	//
	// example:
	//
	// ModuleName
	ModuleName *string `json:"moduleName,omitempty" xml:"moduleName,omitempty"`
	// The workspace name.
	//
	// This parameter is required.
	//
	// example:
	//
	// NamespaceName
	NamespaceName *string `json:"namespaceName,omitempty" xml:"namespaceName,omitempty"`
	// The version number. The value must conform to the [semantic version](http://semver.org/) specification, such as 1.0.1. The initial version is 1.0.0.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1.2.0
	Version *string `json:"version,omitempty" xml:"version,omitempty"`
}

func (s PublishRegistryModuleVersionRequest) String() string {
	return dara.Prettify(s)
}

func (s PublishRegistryModuleVersionRequest) GoString() string {
	return s.String()
}

func (s *PublishRegistryModuleVersionRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *PublishRegistryModuleVersionRequest) GetModuleName() *string {
	return s.ModuleName
}

func (s *PublishRegistryModuleVersionRequest) GetNamespaceName() *string {
	return s.NamespaceName
}

func (s *PublishRegistryModuleVersionRequest) GetVersion() *string {
	return s.Version
}

func (s *PublishRegistryModuleVersionRequest) SetClientToken(v string) *PublishRegistryModuleVersionRequest {
	s.ClientToken = &v
	return s
}

func (s *PublishRegistryModuleVersionRequest) SetModuleName(v string) *PublishRegistryModuleVersionRequest {
	s.ModuleName = &v
	return s
}

func (s *PublishRegistryModuleVersionRequest) SetNamespaceName(v string) *PublishRegistryModuleVersionRequest {
	s.NamespaceName = &v
	return s
}

func (s *PublishRegistryModuleVersionRequest) SetVersion(v string) *PublishRegistryModuleVersionRequest {
	s.Version = &v
	return s
}

func (s *PublishRegistryModuleVersionRequest) Validate() error {
	return dara.Validate(s)
}
