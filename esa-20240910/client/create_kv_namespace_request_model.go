// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateKvNamespaceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *CreateKvNamespaceRequest
	GetDescription() *string
	SetNamespace(v string) *CreateKvNamespaceRequest
	GetNamespace() *string
}

type CreateKvNamespaceRequest struct {
	// The description of the namespace.
	//
	// example:
	//
	// this is a test namespace.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The name of the namespace.
	//
	// This parameter is required.
	//
	// example:
	//
	// test_namespace
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
}

func (s CreateKvNamespaceRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateKvNamespaceRequest) GoString() string {
	return s.String()
}

func (s *CreateKvNamespaceRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateKvNamespaceRequest) GetNamespace() *string {
	return s.Namespace
}

func (s *CreateKvNamespaceRequest) SetDescription(v string) *CreateKvNamespaceRequest {
	s.Description = &v
	return s
}

func (s *CreateKvNamespaceRequest) SetNamespace(v string) *CreateKvNamespaceRequest {
	s.Namespace = &v
	return s
}

func (s *CreateKvNamespaceRequest) Validate() error {
	return dara.Validate(s)
}
