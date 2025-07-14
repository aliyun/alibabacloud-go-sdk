// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateNamespaceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEnableMicroRegistration(v bool) *CreateNamespaceRequest
	GetEnableMicroRegistration() *bool
	SetNameSpaceShortId(v string) *CreateNamespaceRequest
	GetNameSpaceShortId() *string
	SetNamespaceDescription(v string) *CreateNamespaceRequest
	GetNamespaceDescription() *string
	SetNamespaceId(v string) *CreateNamespaceRequest
	GetNamespaceId() *string
	SetNamespaceName(v string) *CreateNamespaceRequest
	GetNamespaceName() *string
}

type CreateNamespaceRequest struct {
	// Indicates whether to enable SAE built-in registry:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// Default value: true. If you do not use the built-in registry, you can set this parameter to false to accelerate the creation of a namespace.
	//
	// example:
	//
	// true
	EnableMicroRegistration *bool `json:"EnableMicroRegistration,omitempty" xml:"EnableMicroRegistration,omitempty"`
	// The trace ID that is used to query the details of the request.
	//
	// example:
	//
	// test
	NameSpaceShortId *string `json:"NameSpaceShortId,omitempty" xml:"NameSpaceShortId,omitempty"`
	// The message returned for the operation.
	//
	// example:
	//
	// desc
	NamespaceDescription *string `json:"NamespaceDescription,omitempty" xml:"NamespaceDescription,omitempty"`
	// The data returned.
	//
	// example:
	//
	// cn-beijing:test
	NamespaceId *string `json:"NamespaceId,omitempty" xml:"NamespaceId,omitempty"`
	// The ID of the request.
	//
	// This parameter is required.
	//
	// example:
	//
	// name
	NamespaceName *string `json:"NamespaceName,omitempty" xml:"NamespaceName,omitempty"`
}

func (s CreateNamespaceRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateNamespaceRequest) GoString() string {
	return s.String()
}

func (s *CreateNamespaceRequest) GetEnableMicroRegistration() *bool {
	return s.EnableMicroRegistration
}

func (s *CreateNamespaceRequest) GetNameSpaceShortId() *string {
	return s.NameSpaceShortId
}

func (s *CreateNamespaceRequest) GetNamespaceDescription() *string {
	return s.NamespaceDescription
}

func (s *CreateNamespaceRequest) GetNamespaceId() *string {
	return s.NamespaceId
}

func (s *CreateNamespaceRequest) GetNamespaceName() *string {
	return s.NamespaceName
}

func (s *CreateNamespaceRequest) SetEnableMicroRegistration(v bool) *CreateNamespaceRequest {
	s.EnableMicroRegistration = &v
	return s
}

func (s *CreateNamespaceRequest) SetNameSpaceShortId(v string) *CreateNamespaceRequest {
	s.NameSpaceShortId = &v
	return s
}

func (s *CreateNamespaceRequest) SetNamespaceDescription(v string) *CreateNamespaceRequest {
	s.NamespaceDescription = &v
	return s
}

func (s *CreateNamespaceRequest) SetNamespaceId(v string) *CreateNamespaceRequest {
	s.NamespaceId = &v
	return s
}

func (s *CreateNamespaceRequest) SetNamespaceName(v string) *CreateNamespaceRequest {
	s.NamespaceName = &v
	return s
}

func (s *CreateNamespaceRequest) Validate() error {
	return dara.Validate(s)
}
