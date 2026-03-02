// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateLibrarySchemaRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBody(v *LibrarySchemaUpdateCmd) *UpdateLibrarySchemaRequest
	GetBody() *LibrarySchemaUpdateCmd
}

type UpdateLibrarySchemaRequest struct {
	Body *LibrarySchemaUpdateCmd `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateLibrarySchemaRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateLibrarySchemaRequest) GoString() string {
	return s.String()
}

func (s *UpdateLibrarySchemaRequest) GetBody() *LibrarySchemaUpdateCmd {
	return s.Body
}

func (s *UpdateLibrarySchemaRequest) SetBody(v *LibrarySchemaUpdateCmd) *UpdateLibrarySchemaRequest {
	s.Body = v
	return s
}

func (s *UpdateLibrarySchemaRequest) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
