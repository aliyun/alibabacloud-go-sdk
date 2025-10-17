// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartMigrationJobResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *StartMigrationJobResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *StartMigrationJobResponse
	GetStatusCode() *int32
	SetBody(v *StartMigrationJobResponseBody) *StartMigrationJobResponse
	GetBody() *StartMigrationJobResponseBody
}

type StartMigrationJobResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *StartMigrationJobResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s StartMigrationJobResponse) String() string {
	return dara.Prettify(s)
}

func (s StartMigrationJobResponse) GoString() string {
	return s.String()
}

func (s *StartMigrationJobResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *StartMigrationJobResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *StartMigrationJobResponse) GetBody() *StartMigrationJobResponseBody {
	return s.Body
}

func (s *StartMigrationJobResponse) SetHeaders(v map[string]*string) *StartMigrationJobResponse {
	s.Headers = v
	return s
}

func (s *StartMigrationJobResponse) SetStatusCode(v int32) *StartMigrationJobResponse {
	s.StatusCode = &v
	return s
}

func (s *StartMigrationJobResponse) SetBody(v *StartMigrationJobResponseBody) *StartMigrationJobResponse {
	s.Body = v
	return s
}

func (s *StartMigrationJobResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
