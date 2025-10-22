// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDataLakeDatabaseResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteDataLakeDatabaseResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteDataLakeDatabaseResponse
	GetStatusCode() *int32
	SetBody(v *DeleteDataLakeDatabaseResponseBody) *DeleteDataLakeDatabaseResponse
	GetBody() *DeleteDataLakeDatabaseResponseBody
}

type DeleteDataLakeDatabaseResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteDataLakeDatabaseResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteDataLakeDatabaseResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteDataLakeDatabaseResponse) GoString() string {
	return s.String()
}

func (s *DeleteDataLakeDatabaseResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteDataLakeDatabaseResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteDataLakeDatabaseResponse) GetBody() *DeleteDataLakeDatabaseResponseBody {
	return s.Body
}

func (s *DeleteDataLakeDatabaseResponse) SetHeaders(v map[string]*string) *DeleteDataLakeDatabaseResponse {
	s.Headers = v
	return s
}

func (s *DeleteDataLakeDatabaseResponse) SetStatusCode(v int32) *DeleteDataLakeDatabaseResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteDataLakeDatabaseResponse) SetBody(v *DeleteDataLakeDatabaseResponseBody) *DeleteDataLakeDatabaseResponse {
	s.Body = v
	return s
}

func (s *DeleteDataLakeDatabaseResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
