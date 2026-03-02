// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateLibrarySchemaRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBody(v *LibrarySchema) *CreateLibrarySchemaRequest
	GetBody() *LibrarySchema
}

type CreateLibrarySchemaRequest struct {
	Body *LibrarySchema `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateLibrarySchemaRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateLibrarySchemaRequest) GoString() string {
	return s.String()
}

func (s *CreateLibrarySchemaRequest) GetBody() *LibrarySchema {
	return s.Body
}

func (s *CreateLibrarySchemaRequest) SetBody(v *LibrarySchema) *CreateLibrarySchemaRequest {
	s.Body = v
	return s
}

func (s *CreateLibrarySchemaRequest) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
