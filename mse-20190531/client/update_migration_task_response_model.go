// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateMigrationTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateMigrationTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateMigrationTaskResponse
	GetStatusCode() *int32
	SetBody(v *UpdateMigrationTaskResponseBody) *UpdateMigrationTaskResponse
	GetBody() *UpdateMigrationTaskResponseBody
}

type UpdateMigrationTaskResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateMigrationTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateMigrationTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateMigrationTaskResponse) GoString() string {
	return s.String()
}

func (s *UpdateMigrationTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateMigrationTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateMigrationTaskResponse) GetBody() *UpdateMigrationTaskResponseBody {
	return s.Body
}

func (s *UpdateMigrationTaskResponse) SetHeaders(v map[string]*string) *UpdateMigrationTaskResponse {
	s.Headers = v
	return s
}

func (s *UpdateMigrationTaskResponse) SetStatusCode(v int32) *UpdateMigrationTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateMigrationTaskResponse) SetBody(v *UpdateMigrationTaskResponseBody) *UpdateMigrationTaskResponse {
	s.Body = v
	return s
}

func (s *UpdateMigrationTaskResponse) Validate() error {
	return dara.Validate(s)
}
