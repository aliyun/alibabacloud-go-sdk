// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iToolsetSchema interface {
	dara.Model
	String() string
	GoString() string
	SetDetail(v string) *ToolsetSchema
	GetDetail() *string
	SetType(v string) *ToolsetSchema
	GetType() *string
}

type ToolsetSchema struct {
	Detail *string `json:"detail,omitempty" xml:"detail,omitempty"`
	// example:
	//
	// OpenAPI
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s ToolsetSchema) String() string {
	return dara.Prettify(s)
}

func (s ToolsetSchema) GoString() string {
	return s.String()
}

func (s *ToolsetSchema) GetDetail() *string {
	return s.Detail
}

func (s *ToolsetSchema) GetType() *string {
	return s.Type
}

func (s *ToolsetSchema) SetDetail(v string) *ToolsetSchema {
	s.Detail = &v
	return s
}

func (s *ToolsetSchema) SetType(v string) *ToolsetSchema {
	s.Type = &v
	return s
}

func (s *ToolsetSchema) Validate() error {
	return dara.Validate(s)
}
