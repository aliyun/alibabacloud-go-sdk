// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMigrationTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetMigrationTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetMigrationTaskResponse
	GetStatusCode() *int32
	SetBody(v *GetMigrationTaskResponseBody) *GetMigrationTaskResponse
	GetBody() *GetMigrationTaskResponseBody
}

type GetMigrationTaskResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetMigrationTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetMigrationTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s GetMigrationTaskResponse) GoString() string {
	return s.String()
}

func (s *GetMigrationTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetMigrationTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetMigrationTaskResponse) GetBody() *GetMigrationTaskResponseBody {
	return s.Body
}

func (s *GetMigrationTaskResponse) SetHeaders(v map[string]*string) *GetMigrationTaskResponse {
	s.Headers = v
	return s
}

func (s *GetMigrationTaskResponse) SetStatusCode(v int32) *GetMigrationTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *GetMigrationTaskResponse) SetBody(v *GetMigrationTaskResponseBody) *GetMigrationTaskResponse {
	s.Body = v
	return s
}

func (s *GetMigrationTaskResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
