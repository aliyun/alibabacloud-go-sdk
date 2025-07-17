// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAlterDatabaseResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AlterDatabaseResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AlterDatabaseResponse
	GetStatusCode() *int32
	SetBody(v *AlterDatabaseResponseBody) *AlterDatabaseResponse
	GetBody() *AlterDatabaseResponseBody
}

type AlterDatabaseResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AlterDatabaseResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AlterDatabaseResponse) String() string {
	return dara.Prettify(s)
}

func (s AlterDatabaseResponse) GoString() string {
	return s.String()
}

func (s *AlterDatabaseResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AlterDatabaseResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AlterDatabaseResponse) GetBody() *AlterDatabaseResponseBody {
	return s.Body
}

func (s *AlterDatabaseResponse) SetHeaders(v map[string]*string) *AlterDatabaseResponse {
	s.Headers = v
	return s
}

func (s *AlterDatabaseResponse) SetStatusCode(v int32) *AlterDatabaseResponse {
	s.StatusCode = &v
	return s
}

func (s *AlterDatabaseResponse) SetBody(v *AlterDatabaseResponseBody) *AlterDatabaseResponse {
	s.Body = v
	return s
}

func (s *AlterDatabaseResponse) Validate() error {
	return dara.Validate(s)
}
