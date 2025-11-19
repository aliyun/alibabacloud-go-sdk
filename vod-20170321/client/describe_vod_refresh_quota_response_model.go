// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVodRefreshQuotaResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeVodRefreshQuotaResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeVodRefreshQuotaResponse
	GetStatusCode() *int32
	SetBody(v *DescribeVodRefreshQuotaResponseBody) *DescribeVodRefreshQuotaResponse
	GetBody() *DescribeVodRefreshQuotaResponseBody
}

type DescribeVodRefreshQuotaResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeVodRefreshQuotaResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeVodRefreshQuotaResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeVodRefreshQuotaResponse) GoString() string {
	return s.String()
}

func (s *DescribeVodRefreshQuotaResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeVodRefreshQuotaResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeVodRefreshQuotaResponse) GetBody() *DescribeVodRefreshQuotaResponseBody {
	return s.Body
}

func (s *DescribeVodRefreshQuotaResponse) SetHeaders(v map[string]*string) *DescribeVodRefreshQuotaResponse {
	s.Headers = v
	return s
}

func (s *DescribeVodRefreshQuotaResponse) SetStatusCode(v int32) *DescribeVodRefreshQuotaResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeVodRefreshQuotaResponse) SetBody(v *DescribeVodRefreshQuotaResponseBody) *DescribeVodRefreshQuotaResponse {
	s.Body = v
	return s
}

func (s *DescribeVodRefreshQuotaResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
