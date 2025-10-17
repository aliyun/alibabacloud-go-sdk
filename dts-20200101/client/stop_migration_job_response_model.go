// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStopMigrationJobResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *StopMigrationJobResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *StopMigrationJobResponse
	GetStatusCode() *int32
	SetBody(v *StopMigrationJobResponseBody) *StopMigrationJobResponse
	GetBody() *StopMigrationJobResponseBody
}

type StopMigrationJobResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *StopMigrationJobResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s StopMigrationJobResponse) String() string {
	return dara.Prettify(s)
}

func (s StopMigrationJobResponse) GoString() string {
	return s.String()
}

func (s *StopMigrationJobResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *StopMigrationJobResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *StopMigrationJobResponse) GetBody() *StopMigrationJobResponseBody {
	return s.Body
}

func (s *StopMigrationJobResponse) SetHeaders(v map[string]*string) *StopMigrationJobResponse {
	s.Headers = v
	return s
}

func (s *StopMigrationJobResponse) SetStatusCode(v int32) *StopMigrationJobResponse {
	s.StatusCode = &v
	return s
}

func (s *StopMigrationJobResponse) SetBody(v *StopMigrationJobResponseBody) *StopMigrationJobResponse {
	s.Body = v
	return s
}

func (s *StopMigrationJobResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
