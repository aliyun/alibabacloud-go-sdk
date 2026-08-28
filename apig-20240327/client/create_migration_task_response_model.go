// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateMigrationTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateMigrationTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateMigrationTaskResponse
	GetStatusCode() *int32
	SetBody(v *CreateMigrationTaskResponseBody) *CreateMigrationTaskResponse
	GetBody() *CreateMigrationTaskResponseBody
}

type CreateMigrationTaskResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateMigrationTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateMigrationTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateMigrationTaskResponse) GoString() string {
	return s.String()
}

func (s *CreateMigrationTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateMigrationTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateMigrationTaskResponse) GetBody() *CreateMigrationTaskResponseBody {
	return s.Body
}

func (s *CreateMigrationTaskResponse) SetHeaders(v map[string]*string) *CreateMigrationTaskResponse {
	s.Headers = v
	return s
}

func (s *CreateMigrationTaskResponse) SetStatusCode(v int32) *CreateMigrationTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateMigrationTaskResponse) SetBody(v *CreateMigrationTaskResponseBody) *CreateMigrationTaskResponse {
	s.Body = v
	return s
}

func (s *CreateMigrationTaskResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
