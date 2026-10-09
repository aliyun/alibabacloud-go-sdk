// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMetaSchemaValue interface {
	dara.Model
	String() string
	GoString() string
	SetType(v string) *MetaSchemaValue
	GetType() *string
}

type MetaSchemaValue struct {
	// The dataset field types. Valid values: text, long, double, and json.
	//
	// example:
	//
	// text
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s MetaSchemaValue) String() string {
	return dara.Prettify(s)
}

func (s MetaSchemaValue) GoString() string {
	return s.String()
}

func (s *MetaSchemaValue) GetType() *string {
	return s.Type
}

func (s *MetaSchemaValue) SetType(v string) *MetaSchemaValue {
	s.Type = &v
	return s
}

func (s *MetaSchemaValue) Validate() error {
	return dara.Validate(s)
}
