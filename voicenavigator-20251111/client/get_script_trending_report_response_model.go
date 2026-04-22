// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetScriptTrendingReportResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetScriptTrendingReportResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetScriptTrendingReportResponse
	GetStatusCode() *int32
	SetBody(v *GetScriptTrendingReportResponseBody) *GetScriptTrendingReportResponse
	GetBody() *GetScriptTrendingReportResponseBody
}

type GetScriptTrendingReportResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetScriptTrendingReportResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetScriptTrendingReportResponse) String() string {
	return dara.Prettify(s)
}

func (s GetScriptTrendingReportResponse) GoString() string {
	return s.String()
}

func (s *GetScriptTrendingReportResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetScriptTrendingReportResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetScriptTrendingReportResponse) GetBody() *GetScriptTrendingReportResponseBody {
	return s.Body
}

func (s *GetScriptTrendingReportResponse) SetHeaders(v map[string]*string) *GetScriptTrendingReportResponse {
	s.Headers = v
	return s
}

func (s *GetScriptTrendingReportResponse) SetStatusCode(v int32) *GetScriptTrendingReportResponse {
	s.StatusCode = &v
	return s
}

func (s *GetScriptTrendingReportResponse) SetBody(v *GetScriptTrendingReportResponseBody) *GetScriptTrendingReportResponse {
	s.Body = v
	return s
}

func (s *GetScriptTrendingReportResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
