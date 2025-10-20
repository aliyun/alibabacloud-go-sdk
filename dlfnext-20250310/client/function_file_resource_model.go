// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iFunctionFileResource interface {
	dara.Model
	String() string
	GoString() string
	SetResourceType(v string) *FunctionFileResource
	GetResourceType() *string
	SetUri(v string) *FunctionFileResource
	GetUri() *string
}

type FunctionFileResource struct {
	ResourceType *string `json:"resourceType,omitempty" xml:"resourceType,omitempty"`
	Uri          *string `json:"uri,omitempty" xml:"uri,omitempty"`
}

func (s FunctionFileResource) String() string {
	return dara.Prettify(s)
}

func (s FunctionFileResource) GoString() string {
	return s.String()
}

func (s *FunctionFileResource) GetResourceType() *string {
	return s.ResourceType
}

func (s *FunctionFileResource) GetUri() *string {
	return s.Uri
}

func (s *FunctionFileResource) SetResourceType(v string) *FunctionFileResource {
	s.ResourceType = &v
	return s
}

func (s *FunctionFileResource) SetUri(v string) *FunctionFileResource {
	s.Uri = &v
	return s
}

func (s *FunctionFileResource) Validate() error {
	return dara.Validate(s)
}
