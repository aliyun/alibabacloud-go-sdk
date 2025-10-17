// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateMigrationJobResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateMigrationJobResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateMigrationJobResponse
	GetStatusCode() *int32
	SetBody(v *CreateMigrationJobResponseBody) *CreateMigrationJobResponse
	GetBody() *CreateMigrationJobResponseBody
}

type CreateMigrationJobResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateMigrationJobResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateMigrationJobResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateMigrationJobResponse) GoString() string {
	return s.String()
}

func (s *CreateMigrationJobResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateMigrationJobResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateMigrationJobResponse) GetBody() *CreateMigrationJobResponseBody {
	return s.Body
}

func (s *CreateMigrationJobResponse) SetHeaders(v map[string]*string) *CreateMigrationJobResponse {
	s.Headers = v
	return s
}

func (s *CreateMigrationJobResponse) SetStatusCode(v int32) *CreateMigrationJobResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateMigrationJobResponse) SetBody(v *CreateMigrationJobResponseBody) *CreateMigrationJobResponse {
	s.Body = v
	return s
}

func (s *CreateMigrationJobResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
