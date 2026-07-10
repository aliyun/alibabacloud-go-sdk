// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartSqlExecutionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBody(v *StartSqlExecutionBody) *StartSqlExecutionRequest
	GetBody() *StartSqlExecutionBody
}

type StartSqlExecutionRequest struct {
	// The request body, which contains the SQL content to execute and description information.
	Body *StartSqlExecutionBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s StartSqlExecutionRequest) String() string {
	return dara.Prettify(s)
}

func (s StartSqlExecutionRequest) GoString() string {
	return s.String()
}

func (s *StartSqlExecutionRequest) GetBody() *StartSqlExecutionBody {
	return s.Body
}

func (s *StartSqlExecutionRequest) SetBody(v *StartSqlExecutionBody) *StartSqlExecutionRequest {
	s.Body = v
	return s
}

func (s *StartSqlExecutionRequest) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
