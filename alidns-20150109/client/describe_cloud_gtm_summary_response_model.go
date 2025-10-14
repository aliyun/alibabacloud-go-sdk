// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCloudGtmSummaryResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeCloudGtmSummaryResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeCloudGtmSummaryResponse
	GetStatusCode() *int32
	SetBody(v *DescribeCloudGtmSummaryResponseBody) *DescribeCloudGtmSummaryResponse
	GetBody() *DescribeCloudGtmSummaryResponseBody
}

type DescribeCloudGtmSummaryResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeCloudGtmSummaryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeCloudGtmSummaryResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeCloudGtmSummaryResponse) GoString() string {
	return s.String()
}

func (s *DescribeCloudGtmSummaryResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeCloudGtmSummaryResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeCloudGtmSummaryResponse) GetBody() *DescribeCloudGtmSummaryResponseBody {
	return s.Body
}

func (s *DescribeCloudGtmSummaryResponse) SetHeaders(v map[string]*string) *DescribeCloudGtmSummaryResponse {
	s.Headers = v
	return s
}

func (s *DescribeCloudGtmSummaryResponse) SetStatusCode(v int32) *DescribeCloudGtmSummaryResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeCloudGtmSummaryResponse) SetBody(v *DescribeCloudGtmSummaryResponseBody) *DescribeCloudGtmSummaryResponse {
	s.Body = v
	return s
}

func (s *DescribeCloudGtmSummaryResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
