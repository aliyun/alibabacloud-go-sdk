// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddMigrationTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AddMigrationTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AddMigrationTaskResponse
	GetStatusCode() *int32
	SetBody(v *AddMigrationTaskResponseBody) *AddMigrationTaskResponse
	GetBody() *AddMigrationTaskResponseBody
}

type AddMigrationTaskResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddMigrationTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddMigrationTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s AddMigrationTaskResponse) GoString() string {
	return s.String()
}

func (s *AddMigrationTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AddMigrationTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AddMigrationTaskResponse) GetBody() *AddMigrationTaskResponseBody {
	return s.Body
}

func (s *AddMigrationTaskResponse) SetHeaders(v map[string]*string) *AddMigrationTaskResponse {
	s.Headers = v
	return s
}

func (s *AddMigrationTaskResponse) SetStatusCode(v int32) *AddMigrationTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *AddMigrationTaskResponse) SetBody(v *AddMigrationTaskResponseBody) *AddMigrationTaskResponse {
	s.Body = v
	return s
}

func (s *AddMigrationTaskResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
