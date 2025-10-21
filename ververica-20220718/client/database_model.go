// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDatabase interface {
	dara.Model
	String() string
	GoString() string
	SetComment(v string) *Database
	GetComment() *string
	SetName(v string) *Database
	GetName() *string
	SetProperties(v map[string]interface{}) *Database
	GetProperties() map[string]interface{}
}

type Database struct {
	Comment    *string                `json:"comment,omitempty" xml:"comment,omitempty"`
	Name       *string                `json:"name,omitempty" xml:"name,omitempty"`
	Properties map[string]interface{} `json:"properties,omitempty" xml:"properties,omitempty"`
}

func (s Database) String() string {
	return dara.Prettify(s)
}

func (s Database) GoString() string {
	return s.String()
}

func (s *Database) GetComment() *string {
	return s.Comment
}

func (s *Database) GetName() *string {
	return s.Name
}

func (s *Database) GetProperties() map[string]interface{} {
	return s.Properties
}

func (s *Database) SetComment(v string) *Database {
	s.Comment = &v
	return s
}

func (s *Database) SetName(v string) *Database {
	s.Name = &v
	return s
}

func (s *Database) SetProperties(v map[string]interface{}) *Database {
	s.Properties = v
	return s
}

func (s *Database) Validate() error {
	return dara.Validate(s)
}
