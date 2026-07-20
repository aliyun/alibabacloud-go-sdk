// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateTableRequest interface {
	dara.Model
	String() string
	GoString() string
	SetIdentifier(v *Identifier) *CreateTableRequest
	GetIdentifier() *Identifier
	SetSchema(v *Schema) *CreateTableRequest
	GetSchema() *Schema
}

type CreateTableRequest struct {
	Identifier *Identifier `json:"identifier,omitempty" xml:"identifier,omitempty"`
	Schema     *Schema     `json:"schema,omitempty" xml:"schema,omitempty"`
}

func (s CreateTableRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateTableRequest) GoString() string {
	return s.String()
}

func (s *CreateTableRequest) GetIdentifier() *Identifier {
	return s.Identifier
}

func (s *CreateTableRequest) GetSchema() *Schema {
	return s.Schema
}

func (s *CreateTableRequest) SetIdentifier(v *Identifier) *CreateTableRequest {
	s.Identifier = v
	return s
}

func (s *CreateTableRequest) SetSchema(v *Schema) *CreateTableRequest {
	s.Schema = v
	return s
}

func (s *CreateTableRequest) Validate() error {
	if s.Identifier != nil {
		if err := s.Identifier.Validate(); err != nil {
			return err
		}
	}
	if s.Schema != nil {
		if err := s.Schema.Validate(); err != nil {
			return err
		}
	}
	return nil
}
