// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartMigrationResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *StartMigrationResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *StartMigrationResponse
	GetStatusCode() *int32
	SetBody(v *StartMigrationResponseBody) *StartMigrationResponse
	GetBody() *StartMigrationResponseBody
}

type StartMigrationResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *StartMigrationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s StartMigrationResponse) String() string {
	return dara.Prettify(s)
}

func (s StartMigrationResponse) GoString() string {
	return s.String()
}

func (s *StartMigrationResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *StartMigrationResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *StartMigrationResponse) GetBody() *StartMigrationResponseBody {
	return s.Body
}

func (s *StartMigrationResponse) SetHeaders(v map[string]*string) *StartMigrationResponse {
	s.Headers = v
	return s
}

func (s *StartMigrationResponse) SetStatusCode(v int32) *StartMigrationResponse {
	s.StatusCode = &v
	return s
}

func (s *StartMigrationResponse) SetBody(v *StartMigrationResponseBody) *StartMigrationResponse {
	s.Body = v
	return s
}

func (s *StartMigrationResponse) Validate() error {
	return dara.Validate(s)
}
