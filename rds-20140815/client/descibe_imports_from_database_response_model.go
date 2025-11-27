// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescibeImportsFromDatabaseResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescibeImportsFromDatabaseResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescibeImportsFromDatabaseResponse
	GetStatusCode() *int32
	SetBody(v *DescibeImportsFromDatabaseResponseBody) *DescibeImportsFromDatabaseResponse
	GetBody() *DescibeImportsFromDatabaseResponseBody
}

type DescibeImportsFromDatabaseResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescibeImportsFromDatabaseResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescibeImportsFromDatabaseResponse) String() string {
	return dara.Prettify(s)
}

func (s DescibeImportsFromDatabaseResponse) GoString() string {
	return s.String()
}

func (s *DescibeImportsFromDatabaseResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescibeImportsFromDatabaseResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescibeImportsFromDatabaseResponse) GetBody() *DescibeImportsFromDatabaseResponseBody {
	return s.Body
}

func (s *DescibeImportsFromDatabaseResponse) SetHeaders(v map[string]*string) *DescibeImportsFromDatabaseResponse {
	s.Headers = v
	return s
}

func (s *DescibeImportsFromDatabaseResponse) SetStatusCode(v int32) *DescibeImportsFromDatabaseResponse {
	s.StatusCode = &v
	return s
}

func (s *DescibeImportsFromDatabaseResponse) SetBody(v *DescibeImportsFromDatabaseResponseBody) *DescibeImportsFromDatabaseResponse {
	s.Body = v
	return s
}

func (s *DescibeImportsFromDatabaseResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
