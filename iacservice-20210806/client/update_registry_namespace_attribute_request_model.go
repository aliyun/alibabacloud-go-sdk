// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateRegistryNamespaceAttributeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAcl(v string) *UpdateRegistryNamespaceAttributeRequest
	GetAcl() *string
	SetClientToken(v string) *UpdateRegistryNamespaceAttributeRequest
	GetClientToken() *string
	SetDescription(v string) *UpdateRegistryNamespaceAttributeRequest
	GetDescription() *string
}

type UpdateRegistryNamespaceAttributeRequest struct {
	// example:
	//
	// private
	Acl *string `json:"acl,omitempty" xml:"acl,omitempty"`
	// This parameter is required.
	ClientToken *string `json:"clientToken,omitempty" xml:"clientToken,omitempty"`
	// example:
	//
	// 123
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
}

func (s UpdateRegistryNamespaceAttributeRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateRegistryNamespaceAttributeRequest) GoString() string {
	return s.String()
}

func (s *UpdateRegistryNamespaceAttributeRequest) GetAcl() *string {
	return s.Acl
}

func (s *UpdateRegistryNamespaceAttributeRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *UpdateRegistryNamespaceAttributeRequest) GetDescription() *string {
	return s.Description
}

func (s *UpdateRegistryNamespaceAttributeRequest) SetAcl(v string) *UpdateRegistryNamespaceAttributeRequest {
	s.Acl = &v
	return s
}

func (s *UpdateRegistryNamespaceAttributeRequest) SetClientToken(v string) *UpdateRegistryNamespaceAttributeRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateRegistryNamespaceAttributeRequest) SetDescription(v string) *UpdateRegistryNamespaceAttributeRequest {
	s.Description = &v
	return s
}

func (s *UpdateRegistryNamespaceAttributeRequest) Validate() error {
	return dara.Validate(s)
}
