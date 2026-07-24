// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateSqlFileRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBody(v *SqlFile) *CreateSqlFileRequest
	GetBody() *SqlFile
}

type CreateSqlFileRequest struct {
	// The request body, which contains the SQL file content and related metadata to be created.
	Body *SqlFile `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateSqlFileRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateSqlFileRequest) GoString() string {
	return s.String()
}

func (s *CreateSqlFileRequest) GetBody() *SqlFile {
	return s.Body
}

func (s *CreateSqlFileRequest) SetBody(v *SqlFile) *CreateSqlFileRequest {
	s.Body = v
	return s
}

func (s *CreateSqlFileRequest) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
