// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteMigrationJobResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteMigrationJobResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteMigrationJobResponse
	GetStatusCode() *int32
	SetBody(v *DeleteMigrationJobResponseBody) *DeleteMigrationJobResponse
	GetBody() *DeleteMigrationJobResponseBody
}

type DeleteMigrationJobResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteMigrationJobResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteMigrationJobResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteMigrationJobResponse) GoString() string {
	return s.String()
}

func (s *DeleteMigrationJobResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteMigrationJobResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteMigrationJobResponse) GetBody() *DeleteMigrationJobResponseBody {
	return s.Body
}

func (s *DeleteMigrationJobResponse) SetHeaders(v map[string]*string) *DeleteMigrationJobResponse {
	s.Headers = v
	return s
}

func (s *DeleteMigrationJobResponse) SetStatusCode(v int32) *DeleteMigrationJobResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteMigrationJobResponse) SetBody(v *DeleteMigrationJobResponseBody) *DeleteMigrationJobResponse {
	s.Body = v
	return s
}

func (s *DeleteMigrationJobResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
