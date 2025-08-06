// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVodUserQuotaResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeVodUserQuotaResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeVodUserQuotaResponse
	GetStatusCode() *int32
	SetBody(v *DescribeVodUserQuotaResponseBody) *DescribeVodUserQuotaResponse
	GetBody() *DescribeVodUserQuotaResponseBody
}

type DescribeVodUserQuotaResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeVodUserQuotaResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeVodUserQuotaResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeVodUserQuotaResponse) GoString() string {
	return s.String()
}

func (s *DescribeVodUserQuotaResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeVodUserQuotaResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeVodUserQuotaResponse) GetBody() *DescribeVodUserQuotaResponseBody {
	return s.Body
}

func (s *DescribeVodUserQuotaResponse) SetHeaders(v map[string]*string) *DescribeVodUserQuotaResponse {
	s.Headers = v
	return s
}

func (s *DescribeVodUserQuotaResponse) SetStatusCode(v int32) *DescribeVodUserQuotaResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeVodUserQuotaResponse) SetBody(v *DescribeVodUserQuotaResponseBody) *DescribeVodUserQuotaResponse {
	s.Body = v
	return s
}

func (s *DescribeVodUserQuotaResponse) Validate() error {
	return dara.Validate(s)
}
