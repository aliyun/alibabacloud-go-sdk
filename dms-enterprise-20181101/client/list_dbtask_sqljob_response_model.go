// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDBTaskSQLJobResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListDBTaskSQLJobResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListDBTaskSQLJobResponse
	GetStatusCode() *int32
	SetBody(v *ListDBTaskSQLJobResponseBody) *ListDBTaskSQLJobResponse
	GetBody() *ListDBTaskSQLJobResponseBody
}

type ListDBTaskSQLJobResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListDBTaskSQLJobResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListDBTaskSQLJobResponse) String() string {
	return dara.Prettify(s)
}

func (s ListDBTaskSQLJobResponse) GoString() string {
	return s.String()
}

func (s *ListDBTaskSQLJobResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListDBTaskSQLJobResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListDBTaskSQLJobResponse) GetBody() *ListDBTaskSQLJobResponseBody {
	return s.Body
}

func (s *ListDBTaskSQLJobResponse) SetHeaders(v map[string]*string) *ListDBTaskSQLJobResponse {
	s.Headers = v
	return s
}

func (s *ListDBTaskSQLJobResponse) SetStatusCode(v int32) *ListDBTaskSQLJobResponse {
	s.StatusCode = &v
	return s
}

func (s *ListDBTaskSQLJobResponse) SetBody(v *ListDBTaskSQLJobResponseBody) *ListDBTaskSQLJobResponse {
	s.Body = v
	return s
}

func (s *ListDBTaskSQLJobResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
