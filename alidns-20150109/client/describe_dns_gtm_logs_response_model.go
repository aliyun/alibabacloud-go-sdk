// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDnsGtmLogsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDnsGtmLogsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDnsGtmLogsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDnsGtmLogsResponseBody) *DescribeDnsGtmLogsResponse
	GetBody() *DescribeDnsGtmLogsResponseBody
}

type DescribeDnsGtmLogsResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDnsGtmLogsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDnsGtmLogsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDnsGtmLogsResponse) GoString() string {
	return s.String()
}

func (s *DescribeDnsGtmLogsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDnsGtmLogsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDnsGtmLogsResponse) GetBody() *DescribeDnsGtmLogsResponseBody {
	return s.Body
}

func (s *DescribeDnsGtmLogsResponse) SetHeaders(v map[string]*string) *DescribeDnsGtmLogsResponse {
	s.Headers = v
	return s
}

func (s *DescribeDnsGtmLogsResponse) SetStatusCode(v int32) *DescribeDnsGtmLogsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDnsGtmLogsResponse) SetBody(v *DescribeDnsGtmLogsResponseBody) *DescribeDnsGtmLogsResponse {
	s.Body = v
	return s
}

func (s *DescribeDnsGtmLogsResponse) Validate() error {
	return dara.Validate(s)
}
