// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeGtmLogsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeGtmLogsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeGtmLogsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeGtmLogsResponseBody) *DescribeGtmLogsResponse
	GetBody() *DescribeGtmLogsResponseBody
}

type DescribeGtmLogsResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeGtmLogsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeGtmLogsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeGtmLogsResponse) GoString() string {
	return s.String()
}

func (s *DescribeGtmLogsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeGtmLogsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeGtmLogsResponse) GetBody() *DescribeGtmLogsResponseBody {
	return s.Body
}

func (s *DescribeGtmLogsResponse) SetHeaders(v map[string]*string) *DescribeGtmLogsResponse {
	s.Headers = v
	return s
}

func (s *DescribeGtmLogsResponse) SetStatusCode(v int32) *DescribeGtmLogsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeGtmLogsResponse) SetBody(v *DescribeGtmLogsResponseBody) *DescribeGtmLogsResponse {
	s.Body = v
	return s
}

func (s *DescribeGtmLogsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
