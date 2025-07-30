// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCacheAnalysisReportListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeCacheAnalysisReportListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeCacheAnalysisReportListResponse
	GetStatusCode() *int32
	SetBody(v *DescribeCacheAnalysisReportListResponseBody) *DescribeCacheAnalysisReportListResponse
	GetBody() *DescribeCacheAnalysisReportListResponseBody
}

type DescribeCacheAnalysisReportListResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeCacheAnalysisReportListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeCacheAnalysisReportListResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeCacheAnalysisReportListResponse) GoString() string {
	return s.String()
}

func (s *DescribeCacheAnalysisReportListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeCacheAnalysisReportListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeCacheAnalysisReportListResponse) GetBody() *DescribeCacheAnalysisReportListResponseBody {
	return s.Body
}

func (s *DescribeCacheAnalysisReportListResponse) SetHeaders(v map[string]*string) *DescribeCacheAnalysisReportListResponse {
	s.Headers = v
	return s
}

func (s *DescribeCacheAnalysisReportListResponse) SetStatusCode(v int32) *DescribeCacheAnalysisReportListResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeCacheAnalysisReportListResponse) SetBody(v *DescribeCacheAnalysisReportListResponseBody) *DescribeCacheAnalysisReportListResponse {
	s.Body = v
	return s
}

func (s *DescribeCacheAnalysisReportListResponse) Validate() error {
	return dara.Validate(s)
}
