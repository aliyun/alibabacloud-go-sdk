// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSuspendMigrationJobResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SuspendMigrationJobResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SuspendMigrationJobResponse
	GetStatusCode() *int32
	SetBody(v *SuspendMigrationJobResponseBody) *SuspendMigrationJobResponse
	GetBody() *SuspendMigrationJobResponseBody
}

type SuspendMigrationJobResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SuspendMigrationJobResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SuspendMigrationJobResponse) String() string {
	return dara.Prettify(s)
}

func (s SuspendMigrationJobResponse) GoString() string {
	return s.String()
}

func (s *SuspendMigrationJobResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SuspendMigrationJobResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SuspendMigrationJobResponse) GetBody() *SuspendMigrationJobResponseBody {
	return s.Body
}

func (s *SuspendMigrationJobResponse) SetHeaders(v map[string]*string) *SuspendMigrationJobResponse {
	s.Headers = v
	return s
}

func (s *SuspendMigrationJobResponse) SetStatusCode(v int32) *SuspendMigrationJobResponse {
	s.StatusCode = &v
	return s
}

func (s *SuspendMigrationJobResponse) SetBody(v *SuspendMigrationJobResponseBody) *SuspendMigrationJobResponse {
	s.Body = v
	return s
}

func (s *SuspendMigrationJobResponse) Validate() error {
	return dara.Validate(s)
}
