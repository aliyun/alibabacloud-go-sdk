// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTableSummaryResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetTableSummaryResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetTableSummaryResponse
	GetStatusCode() *int32
	SetBody(v *TableSummary) *GetTableSummaryResponse
	GetBody() *TableSummary
}

type GetTableSummaryResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *TableSummary      `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetTableSummaryResponse) String() string {
	return dara.Prettify(s)
}

func (s GetTableSummaryResponse) GoString() string {
	return s.String()
}

func (s *GetTableSummaryResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetTableSummaryResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetTableSummaryResponse) GetBody() *TableSummary {
	return s.Body
}

func (s *GetTableSummaryResponse) SetHeaders(v map[string]*string) *GetTableSummaryResponse {
	s.Headers = v
	return s
}

func (s *GetTableSummaryResponse) SetStatusCode(v int32) *GetTableSummaryResponse {
	s.StatusCode = &v
	return s
}

func (s *GetTableSummaryResponse) SetBody(v *TableSummary) *GetTableSummaryResponse {
	s.Body = v
	return s
}

func (s *GetTableSummaryResponse) Validate() error {
	return dara.Validate(s)
}
