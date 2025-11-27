// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVsDomainBpsDataResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeVsDomainBpsDataResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeVsDomainBpsDataResponse
	GetStatusCode() *int32
	SetBody(v *DescribeVsDomainBpsDataResponseBody) *DescribeVsDomainBpsDataResponse
	GetBody() *DescribeVsDomainBpsDataResponseBody
}

type DescribeVsDomainBpsDataResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeVsDomainBpsDataResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeVsDomainBpsDataResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeVsDomainBpsDataResponse) GoString() string {
	return s.String()
}

func (s *DescribeVsDomainBpsDataResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeVsDomainBpsDataResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeVsDomainBpsDataResponse) GetBody() *DescribeVsDomainBpsDataResponseBody {
	return s.Body
}

func (s *DescribeVsDomainBpsDataResponse) SetHeaders(v map[string]*string) *DescribeVsDomainBpsDataResponse {
	s.Headers = v
	return s
}

func (s *DescribeVsDomainBpsDataResponse) SetStatusCode(v int32) *DescribeVsDomainBpsDataResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeVsDomainBpsDataResponse) SetBody(v *DescribeVsDomainBpsDataResponseBody) *DescribeVsDomainBpsDataResponse {
	s.Body = v
	return s
}

func (s *DescribeVsDomainBpsDataResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
