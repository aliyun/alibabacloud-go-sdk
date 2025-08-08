// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModelProviderSchema interface {
	dara.Model
	String() string
	GoString() string
	SetDetail(v string) *ModelProviderSchema
	GetDetail() *string
	SetType(v string) *ModelProviderSchema
	GetType() *string
}

type ModelProviderSchema struct {
	Detail *string `json:"detail,omitempty" xml:"detail,omitempty"`
	// example:
	//
	// OpenAPI
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s ModelProviderSchema) String() string {
	return dara.Prettify(s)
}

func (s ModelProviderSchema) GoString() string {
	return s.String()
}

func (s *ModelProviderSchema) GetDetail() *string {
	return s.Detail
}

func (s *ModelProviderSchema) GetType() *string {
	return s.Type
}

func (s *ModelProviderSchema) SetDetail(v string) *ModelProviderSchema {
	s.Detail = &v
	return s
}

func (s *ModelProviderSchema) SetType(v string) *ModelProviderSchema {
	s.Type = &v
	return s
}

func (s *ModelProviderSchema) Validate() error {
	return dara.Validate(s)
}
