// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateExportMigrationResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateExportMigrationResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateExportMigrationResponse
	GetStatusCode() *int32
	SetBody(v *CreateExportMigrationResponseBody) *CreateExportMigrationResponse
	GetBody() *CreateExportMigrationResponseBody
}

type CreateExportMigrationResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateExportMigrationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateExportMigrationResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateExportMigrationResponse) GoString() string {
	return s.String()
}

func (s *CreateExportMigrationResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateExportMigrationResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateExportMigrationResponse) GetBody() *CreateExportMigrationResponseBody {
	return s.Body
}

func (s *CreateExportMigrationResponse) SetHeaders(v map[string]*string) *CreateExportMigrationResponse {
	s.Headers = v
	return s
}

func (s *CreateExportMigrationResponse) SetStatusCode(v int32) *CreateExportMigrationResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateExportMigrationResponse) SetBody(v *CreateExportMigrationResponseBody) *CreateExportMigrationResponse {
	s.Body = v
	return s
}

func (s *CreateExportMigrationResponse) Validate() error {
	return dara.Validate(s)
}
