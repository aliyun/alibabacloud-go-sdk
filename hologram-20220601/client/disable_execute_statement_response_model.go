// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDisableExecuteStatementResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DisableExecuteStatementResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DisableExecuteStatementResponse
	GetStatusCode() *int32
	SetBody(v *DisableExecuteStatementResponseBody) *DisableExecuteStatementResponse
	GetBody() *DisableExecuteStatementResponseBody
}

type DisableExecuteStatementResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DisableExecuteStatementResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DisableExecuteStatementResponse) String() string {
	return dara.Prettify(s)
}

func (s DisableExecuteStatementResponse) GoString() string {
	return s.String()
}

func (s *DisableExecuteStatementResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DisableExecuteStatementResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DisableExecuteStatementResponse) GetBody() *DisableExecuteStatementResponseBody {
	return s.Body
}

func (s *DisableExecuteStatementResponse) SetHeaders(v map[string]*string) *DisableExecuteStatementResponse {
	s.Headers = v
	return s
}

func (s *DisableExecuteStatementResponse) SetStatusCode(v int32) *DisableExecuteStatementResponse {
	s.StatusCode = &v
	return s
}

func (s *DisableExecuteStatementResponse) SetBody(v *DisableExecuteStatementResponseBody) *DisableExecuteStatementResponse {
	s.Body = v
	return s
}

func (s *DisableExecuteStatementResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
