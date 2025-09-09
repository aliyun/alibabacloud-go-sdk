// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDatabaseResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteDatabaseResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteDatabaseResponse
	GetStatusCode() *int32
	SetBody(v *DeleteDatabaseResponseBody) *DeleteDatabaseResponse
	GetBody() *DeleteDatabaseResponseBody
}

type DeleteDatabaseResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteDatabaseResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteDatabaseResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteDatabaseResponse) GoString() string {
	return s.String()
}

func (s *DeleteDatabaseResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteDatabaseResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteDatabaseResponse) GetBody() *DeleteDatabaseResponseBody {
	return s.Body
}

func (s *DeleteDatabaseResponse) SetHeaders(v map[string]*string) *DeleteDatabaseResponse {
	s.Headers = v
	return s
}

func (s *DeleteDatabaseResponse) SetStatusCode(v int32) *DeleteDatabaseResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteDatabaseResponse) SetBody(v *DeleteDatabaseResponseBody) *DeleteDatabaseResponse {
	s.Body = v
	return s
}

func (s *DeleteDatabaseResponse) Validate() error {
	return dara.Validate(s)
}
