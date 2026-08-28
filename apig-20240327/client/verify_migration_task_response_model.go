// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iVerifyMigrationTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *VerifyMigrationTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *VerifyMigrationTaskResponse
	GetStatusCode() *int32
	SetBody(v *VerifyMigrationTaskResponseBody) *VerifyMigrationTaskResponse
	GetBody() *VerifyMigrationTaskResponseBody
}

type VerifyMigrationTaskResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *VerifyMigrationTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s VerifyMigrationTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s VerifyMigrationTaskResponse) GoString() string {
	return s.String()
}

func (s *VerifyMigrationTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *VerifyMigrationTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *VerifyMigrationTaskResponse) GetBody() *VerifyMigrationTaskResponseBody {
	return s.Body
}

func (s *VerifyMigrationTaskResponse) SetHeaders(v map[string]*string) *VerifyMigrationTaskResponse {
	s.Headers = v
	return s
}

func (s *VerifyMigrationTaskResponse) SetStatusCode(v int32) *VerifyMigrationTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *VerifyMigrationTaskResponse) SetBody(v *VerifyMigrationTaskResponseBody) *VerifyMigrationTaskResponse {
	s.Body = v
	return s
}

func (s *VerifyMigrationTaskResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
