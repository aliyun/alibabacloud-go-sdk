// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDcdnKvNamespaceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetNamespace(v string) *DeleteDcdnKvNamespaceRequest
	GetNamespace() *string
}

type DeleteDcdnKvNamespaceRequest struct {
	// The name of the namespace. You can call the [PutDcdnKvNamespace](~~PutDcdnKvNamespace~~) operation to query the name of a namespace.
	//
	// This parameter is required.
	//
	// example:
	//
	// ns1
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
}

func (s DeleteDcdnKvNamespaceRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteDcdnKvNamespaceRequest) GoString() string {
	return s.String()
}

func (s *DeleteDcdnKvNamespaceRequest) GetNamespace() *string {
	return s.Namespace
}

func (s *DeleteDcdnKvNamespaceRequest) SetNamespace(v string) *DeleteDcdnKvNamespaceRequest {
	s.Namespace = &v
	return s
}

func (s *DeleteDcdnKvNamespaceRequest) Validate() error {
	return dara.Validate(s)
}
