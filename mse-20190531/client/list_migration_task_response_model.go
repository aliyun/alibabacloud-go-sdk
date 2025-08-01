// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListMigrationTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListMigrationTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListMigrationTaskResponse
	GetStatusCode() *int32
	SetBody(v *ListMigrationTaskResponseBody) *ListMigrationTaskResponse
	GetBody() *ListMigrationTaskResponseBody
}

type ListMigrationTaskResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListMigrationTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListMigrationTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s ListMigrationTaskResponse) GoString() string {
	return s.String()
}

func (s *ListMigrationTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListMigrationTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListMigrationTaskResponse) GetBody() *ListMigrationTaskResponseBody {
	return s.Body
}

func (s *ListMigrationTaskResponse) SetHeaders(v map[string]*string) *ListMigrationTaskResponse {
	s.Headers = v
	return s
}

func (s *ListMigrationTaskResponse) SetStatusCode(v int32) *ListMigrationTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *ListMigrationTaskResponse) SetBody(v *ListMigrationTaskResponseBody) *ListMigrationTaskResponse {
	s.Body = v
	return s
}

func (s *ListMigrationTaskResponse) Validate() error {
	return dara.Validate(s)
}
