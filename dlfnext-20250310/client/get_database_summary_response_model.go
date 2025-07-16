// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDatabaseSummaryResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetDatabaseSummaryResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetDatabaseSummaryResponse
	GetStatusCode() *int32
	SetBody(v *DatabaseSummary) *GetDatabaseSummaryResponse
	GetBody() *DatabaseSummary
}

type GetDatabaseSummaryResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DatabaseSummary   `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetDatabaseSummaryResponse) String() string {
	return dara.Prettify(s)
}

func (s GetDatabaseSummaryResponse) GoString() string {
	return s.String()
}

func (s *GetDatabaseSummaryResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetDatabaseSummaryResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetDatabaseSummaryResponse) GetBody() *DatabaseSummary {
	return s.Body
}

func (s *GetDatabaseSummaryResponse) SetHeaders(v map[string]*string) *GetDatabaseSummaryResponse {
	s.Headers = v
	return s
}

func (s *GetDatabaseSummaryResponse) SetStatusCode(v int32) *GetDatabaseSummaryResponse {
	s.StatusCode = &v
	return s
}

func (s *GetDatabaseSummaryResponse) SetBody(v *DatabaseSummary) *GetDatabaseSummaryResponse {
	s.Body = v
	return s
}

func (s *GetDatabaseSummaryResponse) Validate() error {
	return dara.Validate(s)
}
