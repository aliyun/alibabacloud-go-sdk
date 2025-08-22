// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPutDcdnKvNamespaceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *PutDcdnKvNamespaceRequest
	GetDescription() *string
	SetNamespace(v string) *PutDcdnKvNamespaceRequest
	GetNamespace() *string
}

type PutDcdnKvNamespaceRequest struct {
	// The description of the namespace.
	//
	// example:
	//
	// the first namespace
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The name of the namespace. The name can contain letters, digits, hyphens (-), and underscores (_).
	//
	// This parameter is required.
	//
	// example:
	//
	// ns1
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
}

func (s PutDcdnKvNamespaceRequest) String() string {
	return dara.Prettify(s)
}

func (s PutDcdnKvNamespaceRequest) GoString() string {
	return s.String()
}

func (s *PutDcdnKvNamespaceRequest) GetDescription() *string {
	return s.Description
}

func (s *PutDcdnKvNamespaceRequest) GetNamespace() *string {
	return s.Namespace
}

func (s *PutDcdnKvNamespaceRequest) SetDescription(v string) *PutDcdnKvNamespaceRequest {
	s.Description = &v
	return s
}

func (s *PutDcdnKvNamespaceRequest) SetNamespace(v string) *PutDcdnKvNamespaceRequest {
	s.Namespace = &v
	return s
}

func (s *PutDcdnKvNamespaceRequest) Validate() error {
	return dara.Validate(s)
}
