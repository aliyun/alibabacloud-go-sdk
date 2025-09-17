// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateCrossZoneMigrationJobResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateCrossZoneMigrationJobResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateCrossZoneMigrationJobResponse
	GetStatusCode() *int32
	SetBody(v *CreateCrossZoneMigrationJobResponseBody) *CreateCrossZoneMigrationJobResponse
	GetBody() *CreateCrossZoneMigrationJobResponseBody
}

type CreateCrossZoneMigrationJobResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateCrossZoneMigrationJobResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateCrossZoneMigrationJobResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateCrossZoneMigrationJobResponse) GoString() string {
	return s.String()
}

func (s *CreateCrossZoneMigrationJobResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateCrossZoneMigrationJobResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateCrossZoneMigrationJobResponse) GetBody() *CreateCrossZoneMigrationJobResponseBody {
	return s.Body
}

func (s *CreateCrossZoneMigrationJobResponse) SetHeaders(v map[string]*string) *CreateCrossZoneMigrationJobResponse {
	s.Headers = v
	return s
}

func (s *CreateCrossZoneMigrationJobResponse) SetStatusCode(v int32) *CreateCrossZoneMigrationJobResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateCrossZoneMigrationJobResponse) SetBody(v *CreateCrossZoneMigrationJobResponseBody) *CreateCrossZoneMigrationJobResponse {
	s.Body = v
	return s
}

func (s *CreateCrossZoneMigrationJobResponse) Validate() error {
	return dara.Validate(s)
}
