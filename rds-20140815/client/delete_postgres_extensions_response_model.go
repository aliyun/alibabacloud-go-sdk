// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeletePostgresExtensionsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeletePostgresExtensionsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeletePostgresExtensionsResponse
	GetStatusCode() *int32
	SetBody(v *DeletePostgresExtensionsResponseBody) *DeletePostgresExtensionsResponse
	GetBody() *DeletePostgresExtensionsResponseBody
}

type DeletePostgresExtensionsResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeletePostgresExtensionsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeletePostgresExtensionsResponse) String() string {
	return dara.Prettify(s)
}

func (s DeletePostgresExtensionsResponse) GoString() string {
	return s.String()
}

func (s *DeletePostgresExtensionsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeletePostgresExtensionsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeletePostgresExtensionsResponse) GetBody() *DeletePostgresExtensionsResponseBody {
	return s.Body
}

func (s *DeletePostgresExtensionsResponse) SetHeaders(v map[string]*string) *DeletePostgresExtensionsResponse {
	s.Headers = v
	return s
}

func (s *DeletePostgresExtensionsResponse) SetStatusCode(v int32) *DeletePostgresExtensionsResponse {
	s.StatusCode = &v
	return s
}

func (s *DeletePostgresExtensionsResponse) SetBody(v *DeletePostgresExtensionsResponseBody) *DeletePostgresExtensionsResponse {
	s.Body = v
	return s
}

func (s *DeletePostgresExtensionsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
