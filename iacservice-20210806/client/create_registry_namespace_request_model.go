// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateRegistryNamespaceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAcl(v string) *CreateRegistryNamespaceRequest
	GetAcl() *string
	SetClientToken(v string) *CreateRegistryNamespaceRequest
	GetClientToken() *string
	SetDescription(v string) *CreateRegistryNamespaceRequest
	GetDescription() *string
	SetMaintainer(v string) *CreateRegistryNamespaceRequest
	GetMaintainer() *string
	SetNamespaceName(v string) *CreateRegistryNamespaceRequest
	GetNamespaceName() *string
}

type CreateRegistryNamespaceRequest struct {
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
	// example:
	//
	// admin
	Maintainer *string `json:"maintainer,omitempty" xml:"maintainer,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// iac
	NamespaceName *string `json:"namespaceName,omitempty" xml:"namespaceName,omitempty"`
}

func (s CreateRegistryNamespaceRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateRegistryNamespaceRequest) GoString() string {
	return s.String()
}

func (s *CreateRegistryNamespaceRequest) GetAcl() *string {
	return s.Acl
}

func (s *CreateRegistryNamespaceRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateRegistryNamespaceRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateRegistryNamespaceRequest) GetMaintainer() *string {
	return s.Maintainer
}

func (s *CreateRegistryNamespaceRequest) GetNamespaceName() *string {
	return s.NamespaceName
}

func (s *CreateRegistryNamespaceRequest) SetAcl(v string) *CreateRegistryNamespaceRequest {
	s.Acl = &v
	return s
}

func (s *CreateRegistryNamespaceRequest) SetClientToken(v string) *CreateRegistryNamespaceRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateRegistryNamespaceRequest) SetDescription(v string) *CreateRegistryNamespaceRequest {
	s.Description = &v
	return s
}

func (s *CreateRegistryNamespaceRequest) SetMaintainer(v string) *CreateRegistryNamespaceRequest {
	s.Maintainer = &v
	return s
}

func (s *CreateRegistryNamespaceRequest) SetNamespaceName(v string) *CreateRegistryNamespaceRequest {
	s.NamespaceName = &v
	return s
}

func (s *CreateRegistryNamespaceRequest) Validate() error {
	return dara.Validate(s)
}
