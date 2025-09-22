// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetGwpBenchmarkSummaryResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetGwpBenchmarkSummaryResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetGwpBenchmarkSummaryResponse
	GetStatusCode() *int32
	SetBody(v *GetGwpBenchmarkSummaryResponseBody) *GetGwpBenchmarkSummaryResponse
	GetBody() *GetGwpBenchmarkSummaryResponseBody
}

type GetGwpBenchmarkSummaryResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetGwpBenchmarkSummaryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetGwpBenchmarkSummaryResponse) String() string {
	return dara.Prettify(s)
}

func (s GetGwpBenchmarkSummaryResponse) GoString() string {
	return s.String()
}

func (s *GetGwpBenchmarkSummaryResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetGwpBenchmarkSummaryResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetGwpBenchmarkSummaryResponse) GetBody() *GetGwpBenchmarkSummaryResponseBody {
	return s.Body
}

func (s *GetGwpBenchmarkSummaryResponse) SetHeaders(v map[string]*string) *GetGwpBenchmarkSummaryResponse {
	s.Headers = v
	return s
}

func (s *GetGwpBenchmarkSummaryResponse) SetStatusCode(v int32) *GetGwpBenchmarkSummaryResponse {
	s.StatusCode = &v
	return s
}

func (s *GetGwpBenchmarkSummaryResponse) SetBody(v *GetGwpBenchmarkSummaryResponseBody) *GetGwpBenchmarkSummaryResponse {
	s.Body = v
	return s
}

func (s *GetGwpBenchmarkSummaryResponse) Validate() error {
	return dara.Validate(s)
}
