// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDcdnKvNamespaceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetNamespace(v string) *DescribeDcdnKvNamespaceRequest
	GetNamespace() *string
}

type DescribeDcdnKvNamespaceRequest struct {
	// The name of the namespace.
	//
	// This parameter is required.
	//
	// example:
	//
	// ns1
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
}

func (s DescribeDcdnKvNamespaceRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnKvNamespaceRequest) GoString() string {
	return s.String()
}

func (s *DescribeDcdnKvNamespaceRequest) GetNamespace() *string {
	return s.Namespace
}

func (s *DescribeDcdnKvNamespaceRequest) SetNamespace(v string) *DescribeDcdnKvNamespaceRequest {
	s.Namespace = &v
	return s
}

func (s *DescribeDcdnKvNamespaceRequest) Validate() error {
	return dara.Validate(s)
}
