// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteMigrationTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteMigrationTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteMigrationTaskResponse
	GetStatusCode() *int32
	SetBody(v *DeleteMigrationTaskResponseBody) *DeleteMigrationTaskResponse
	GetBody() *DeleteMigrationTaskResponseBody
}

type DeleteMigrationTaskResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteMigrationTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteMigrationTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteMigrationTaskResponse) GoString() string {
	return s.String()
}

func (s *DeleteMigrationTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteMigrationTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteMigrationTaskResponse) GetBody() *DeleteMigrationTaskResponseBody {
	return s.Body
}

func (s *DeleteMigrationTaskResponse) SetHeaders(v map[string]*string) *DeleteMigrationTaskResponse {
	s.Headers = v
	return s
}

func (s *DeleteMigrationTaskResponse) SetStatusCode(v int32) *DeleteMigrationTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteMigrationTaskResponse) SetBody(v *DeleteMigrationTaskResponseBody) *DeleteMigrationTaskResponse {
	s.Body = v
	return s
}

func (s *DeleteMigrationTaskResponse) Validate() error {
	return dara.Validate(s)
}
