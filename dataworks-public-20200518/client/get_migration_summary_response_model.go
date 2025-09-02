// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMigrationSummaryResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetMigrationSummaryResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetMigrationSummaryResponse
	GetStatusCode() *int32
	SetBody(v *GetMigrationSummaryResponseBody) *GetMigrationSummaryResponse
	GetBody() *GetMigrationSummaryResponseBody
}

type GetMigrationSummaryResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetMigrationSummaryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetMigrationSummaryResponse) String() string {
	return dara.Prettify(s)
}

func (s GetMigrationSummaryResponse) GoString() string {
	return s.String()
}

func (s *GetMigrationSummaryResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetMigrationSummaryResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetMigrationSummaryResponse) GetBody() *GetMigrationSummaryResponseBody {
	return s.Body
}

func (s *GetMigrationSummaryResponse) SetHeaders(v map[string]*string) *GetMigrationSummaryResponse {
	s.Headers = v
	return s
}

func (s *GetMigrationSummaryResponse) SetStatusCode(v int32) *GetMigrationSummaryResponse {
	s.StatusCode = &v
	return s
}

func (s *GetMigrationSummaryResponse) SetBody(v *GetMigrationSummaryResponseBody) *GetMigrationSummaryResponse {
	s.Body = v
	return s
}

func (s *GetMigrationSummaryResponse) Validate() error {
	return dara.Validate(s)
}
