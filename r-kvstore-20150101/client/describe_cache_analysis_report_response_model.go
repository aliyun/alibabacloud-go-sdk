// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCacheAnalysisReportResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeCacheAnalysisReportResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeCacheAnalysisReportResponse
	GetStatusCode() *int32
	SetBody(v *DescribeCacheAnalysisReportResponseBody) *DescribeCacheAnalysisReportResponse
	GetBody() *DescribeCacheAnalysisReportResponseBody
}

type DescribeCacheAnalysisReportResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeCacheAnalysisReportResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeCacheAnalysisReportResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeCacheAnalysisReportResponse) GoString() string {
	return s.String()
}

func (s *DescribeCacheAnalysisReportResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeCacheAnalysisReportResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeCacheAnalysisReportResponse) GetBody() *DescribeCacheAnalysisReportResponseBody {
	return s.Body
}

func (s *DescribeCacheAnalysisReportResponse) SetHeaders(v map[string]*string) *DescribeCacheAnalysisReportResponse {
	s.Headers = v
	return s
}

func (s *DescribeCacheAnalysisReportResponse) SetStatusCode(v int32) *DescribeCacheAnalysisReportResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeCacheAnalysisReportResponse) SetBody(v *DescribeCacheAnalysisReportResponseBody) *DescribeCacheAnalysisReportResponse {
	s.Body = v
	return s
}

func (s *DescribeCacheAnalysisReportResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
