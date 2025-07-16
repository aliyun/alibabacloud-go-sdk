// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRefreshQuotaResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeRefreshQuotaResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeRefreshQuotaResponse
	GetStatusCode() *int32
	SetBody(v *DescribeRefreshQuotaResponseBody) *DescribeRefreshQuotaResponse
	GetBody() *DescribeRefreshQuotaResponseBody
}

type DescribeRefreshQuotaResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeRefreshQuotaResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeRefreshQuotaResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeRefreshQuotaResponse) GoString() string {
	return s.String()
}

func (s *DescribeRefreshQuotaResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeRefreshQuotaResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeRefreshQuotaResponse) GetBody() *DescribeRefreshQuotaResponseBody {
	return s.Body
}

func (s *DescribeRefreshQuotaResponse) SetHeaders(v map[string]*string) *DescribeRefreshQuotaResponse {
	s.Headers = v
	return s
}

func (s *DescribeRefreshQuotaResponse) SetStatusCode(v int32) *DescribeRefreshQuotaResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeRefreshQuotaResponse) SetBody(v *DescribeRefreshQuotaResponseBody) *DescribeRefreshQuotaResponse {
	s.Body = v
	return s
}

func (s *DescribeRefreshQuotaResponse) Validate() error {
	return dara.Validate(s)
}
