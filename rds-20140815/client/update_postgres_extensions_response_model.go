// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdatePostgresExtensionsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdatePostgresExtensionsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdatePostgresExtensionsResponse
	GetStatusCode() *int32
	SetBody(v *UpdatePostgresExtensionsResponseBody) *UpdatePostgresExtensionsResponse
	GetBody() *UpdatePostgresExtensionsResponseBody
}

type UpdatePostgresExtensionsResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdatePostgresExtensionsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdatePostgresExtensionsResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdatePostgresExtensionsResponse) GoString() string {
	return s.String()
}

func (s *UpdatePostgresExtensionsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdatePostgresExtensionsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdatePostgresExtensionsResponse) GetBody() *UpdatePostgresExtensionsResponseBody {
	return s.Body
}

func (s *UpdatePostgresExtensionsResponse) SetHeaders(v map[string]*string) *UpdatePostgresExtensionsResponse {
	s.Headers = v
	return s
}

func (s *UpdatePostgresExtensionsResponse) SetStatusCode(v int32) *UpdatePostgresExtensionsResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdatePostgresExtensionsResponse) SetBody(v *UpdatePostgresExtensionsResponseBody) *UpdatePostgresExtensionsResponse {
	s.Body = v
	return s
}

func (s *UpdatePostgresExtensionsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
