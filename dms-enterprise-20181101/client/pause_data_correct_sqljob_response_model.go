// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPauseDataCorrectSQLJobResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *PauseDataCorrectSQLJobResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *PauseDataCorrectSQLJobResponse
	GetStatusCode() *int32
	SetBody(v *PauseDataCorrectSQLJobResponseBody) *PauseDataCorrectSQLJobResponse
	GetBody() *PauseDataCorrectSQLJobResponseBody
}

type PauseDataCorrectSQLJobResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *PauseDataCorrectSQLJobResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s PauseDataCorrectSQLJobResponse) String() string {
	return dara.Prettify(s)
}

func (s PauseDataCorrectSQLJobResponse) GoString() string {
	return s.String()
}

func (s *PauseDataCorrectSQLJobResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *PauseDataCorrectSQLJobResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *PauseDataCorrectSQLJobResponse) GetBody() *PauseDataCorrectSQLJobResponseBody {
	return s.Body
}

func (s *PauseDataCorrectSQLJobResponse) SetHeaders(v map[string]*string) *PauseDataCorrectSQLJobResponse {
	s.Headers = v
	return s
}

func (s *PauseDataCorrectSQLJobResponse) SetStatusCode(v int32) *PauseDataCorrectSQLJobResponse {
	s.StatusCode = &v
	return s
}

func (s *PauseDataCorrectSQLJobResponse) SetBody(v *PauseDataCorrectSQLJobResponseBody) *PauseDataCorrectSQLJobResponse {
	s.Body = v
	return s
}

func (s *PauseDataCorrectSQLJobResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
