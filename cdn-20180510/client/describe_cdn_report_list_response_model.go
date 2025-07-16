// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCdnReportListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeCdnReportListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeCdnReportListResponse
	GetStatusCode() *int32
	SetBody(v *DescribeCdnReportListResponseBody) *DescribeCdnReportListResponse
	GetBody() *DescribeCdnReportListResponseBody
}

type DescribeCdnReportListResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeCdnReportListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeCdnReportListResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeCdnReportListResponse) GoString() string {
	return s.String()
}

func (s *DescribeCdnReportListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeCdnReportListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeCdnReportListResponse) GetBody() *DescribeCdnReportListResponseBody {
	return s.Body
}

func (s *DescribeCdnReportListResponse) SetHeaders(v map[string]*string) *DescribeCdnReportListResponse {
	s.Headers = v
	return s
}

func (s *DescribeCdnReportListResponse) SetStatusCode(v int32) *DescribeCdnReportListResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeCdnReportListResponse) SetBody(v *DescribeCdnReportListResponseBody) *DescribeCdnReportListResponse {
	s.Body = v
	return s
}

func (s *DescribeCdnReportListResponse) Validate() error {
	return dara.Validate(s)
}
