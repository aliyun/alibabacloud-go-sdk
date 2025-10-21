// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iValidateSqlStatementRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBody(v *SqlStatementWithContext) *ValidateSqlStatementRequest
	GetBody() *SqlStatementWithContext
}

type ValidateSqlStatementRequest struct {
	// The content of the code that you want to verify.
	//
	// This parameter is required.
	Body *SqlStatementWithContext `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ValidateSqlStatementRequest) String() string {
	return dara.Prettify(s)
}

func (s ValidateSqlStatementRequest) GoString() string {
	return s.String()
}

func (s *ValidateSqlStatementRequest) GetBody() *SqlStatementWithContext {
	return s.Body
}

func (s *ValidateSqlStatementRequest) SetBody(v *SqlStatementWithContext) *ValidateSqlStatementRequest {
	s.Body = v
	return s
}

func (s *ValidateSqlStatementRequest) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
