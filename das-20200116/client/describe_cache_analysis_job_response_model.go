// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCacheAnalysisJobResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeCacheAnalysisJobResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeCacheAnalysisJobResponse
	GetStatusCode() *int32
	SetBody(v *DescribeCacheAnalysisJobResponseBody) *DescribeCacheAnalysisJobResponse
	GetBody() *DescribeCacheAnalysisJobResponseBody
}

type DescribeCacheAnalysisJobResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeCacheAnalysisJobResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeCacheAnalysisJobResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeCacheAnalysisJobResponse) GoString() string {
	return s.String()
}

func (s *DescribeCacheAnalysisJobResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeCacheAnalysisJobResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeCacheAnalysisJobResponse) GetBody() *DescribeCacheAnalysisJobResponseBody {
	return s.Body
}

func (s *DescribeCacheAnalysisJobResponse) SetHeaders(v map[string]*string) *DescribeCacheAnalysisJobResponse {
	s.Headers = v
	return s
}

func (s *DescribeCacheAnalysisJobResponse) SetStatusCode(v int32) *DescribeCacheAnalysisJobResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeCacheAnalysisJobResponse) SetBody(v *DescribeCacheAnalysisJobResponseBody) *DescribeCacheAnalysisJobResponse {
	s.Body = v
	return s
}

func (s *DescribeCacheAnalysisJobResponse) Validate() error {
	return dara.Validate(s)
}
