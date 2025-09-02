// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateImportMigrationResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateImportMigrationResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateImportMigrationResponse
	GetStatusCode() *int32
	SetBody(v *CreateImportMigrationResponseBody) *CreateImportMigrationResponse
	GetBody() *CreateImportMigrationResponseBody
}

type CreateImportMigrationResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateImportMigrationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateImportMigrationResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateImportMigrationResponse) GoString() string {
	return s.String()
}

func (s *CreateImportMigrationResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateImportMigrationResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateImportMigrationResponse) GetBody() *CreateImportMigrationResponseBody {
	return s.Body
}

func (s *CreateImportMigrationResponse) SetHeaders(v map[string]*string) *CreateImportMigrationResponse {
	s.Headers = v
	return s
}

func (s *CreateImportMigrationResponse) SetStatusCode(v int32) *CreateImportMigrationResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateImportMigrationResponse) SetBody(v *CreateImportMigrationResponseBody) *CreateImportMigrationResponse {
	s.Body = v
	return s
}

func (s *CreateImportMigrationResponse) Validate() error {
	return dara.Validate(s)
}
