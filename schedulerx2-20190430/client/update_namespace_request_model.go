// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateNamespaceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *UpdateNamespaceRequest
	GetDescription() *string
	SetNamespace(v string) *UpdateNamespaceRequest
	GetNamespace() *string
	SetNamespaceName(v string) *UpdateNamespaceRequest
	GetNamespaceName() *string
	SetRegionId(v string) *UpdateNamespaceRequest
	GetRegionId() *string
}

type UpdateNamespaceRequest struct {
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// adcfc35d-e2fe-4fe9-bbaa-20e90ffc****
	Namespace     *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	NamespaceName *string `json:"NamespaceName,omitempty" xml:"NamespaceName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s UpdateNamespaceRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateNamespaceRequest) GoString() string {
	return s.String()
}

func (s *UpdateNamespaceRequest) GetDescription() *string {
	return s.Description
}

func (s *UpdateNamespaceRequest) GetNamespace() *string {
	return s.Namespace
}

func (s *UpdateNamespaceRequest) GetNamespaceName() *string {
	return s.NamespaceName
}

func (s *UpdateNamespaceRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *UpdateNamespaceRequest) SetDescription(v string) *UpdateNamespaceRequest {
	s.Description = &v
	return s
}

func (s *UpdateNamespaceRequest) SetNamespace(v string) *UpdateNamespaceRequest {
	s.Namespace = &v
	return s
}

func (s *UpdateNamespaceRequest) SetNamespaceName(v string) *UpdateNamespaceRequest {
	s.NamespaceName = &v
	return s
}

func (s *UpdateNamespaceRequest) SetRegionId(v string) *UpdateNamespaceRequest {
	s.RegionId = &v
	return s
}

func (s *UpdateNamespaceRequest) Validate() error {
	return dara.Validate(s)
}
