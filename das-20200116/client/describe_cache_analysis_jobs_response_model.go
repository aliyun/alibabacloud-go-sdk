// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCacheAnalysisJobsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeCacheAnalysisJobsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeCacheAnalysisJobsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeCacheAnalysisJobsResponseBody) *DescribeCacheAnalysisJobsResponse
	GetBody() *DescribeCacheAnalysisJobsResponseBody
}

type DescribeCacheAnalysisJobsResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeCacheAnalysisJobsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeCacheAnalysisJobsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeCacheAnalysisJobsResponse) GoString() string {
	return s.String()
}

func (s *DescribeCacheAnalysisJobsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeCacheAnalysisJobsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeCacheAnalysisJobsResponse) GetBody() *DescribeCacheAnalysisJobsResponseBody {
	return s.Body
}

func (s *DescribeCacheAnalysisJobsResponse) SetHeaders(v map[string]*string) *DescribeCacheAnalysisJobsResponse {
	s.Headers = v
	return s
}

func (s *DescribeCacheAnalysisJobsResponse) SetStatusCode(v int32) *DescribeCacheAnalysisJobsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeCacheAnalysisJobsResponse) SetBody(v *DescribeCacheAnalysisJobsResponseBody) *DescribeCacheAnalysisJobsResponse {
	s.Body = v
	return s
}

func (s *DescribeCacheAnalysisJobsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
