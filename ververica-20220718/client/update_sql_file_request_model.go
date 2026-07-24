// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateSqlFileRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBody(v *SqlFile) *UpdateSqlFileRequest
	GetBody() *SqlFile
}

type UpdateSqlFileRequest struct {
	// The SQL script information to update.
	Body *SqlFile `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateSqlFileRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateSqlFileRequest) GoString() string {
	return s.String()
}

func (s *UpdateSqlFileRequest) GetBody() *SqlFile {
	return s.Body
}

func (s *UpdateSqlFileRequest) SetBody(v *SqlFile) *UpdateSqlFileRequest {
	s.Body = v
	return s
}

func (s *UpdateSqlFileRequest) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
