// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateSqlExecJobResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateSqlExecJobResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateSqlExecJobResponse
	GetStatusCode() *int32
	SetBody(v *CreateSqlExecJobResponseBody) *CreateSqlExecJobResponse
	GetBody() *CreateSqlExecJobResponseBody
}

type CreateSqlExecJobResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateSqlExecJobResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateSqlExecJobResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateSqlExecJobResponse) GoString() string {
	return s.String()
}

func (s *CreateSqlExecJobResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateSqlExecJobResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateSqlExecJobResponse) GetBody() *CreateSqlExecJobResponseBody {
	return s.Body
}

func (s *CreateSqlExecJobResponse) SetHeaders(v map[string]*string) *CreateSqlExecJobResponse {
	s.Headers = v
	return s
}

func (s *CreateSqlExecJobResponse) SetStatusCode(v int32) *CreateSqlExecJobResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateSqlExecJobResponse) SetBody(v *CreateSqlExecJobResponseBody) *CreateSqlExecJobResponse {
	s.Body = v
	return s
}

func (s *CreateSqlExecJobResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
