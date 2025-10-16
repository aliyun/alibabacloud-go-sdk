// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePrivateDnsDomainNameListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribePrivateDnsDomainNameListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribePrivateDnsDomainNameListResponse
	GetStatusCode() *int32
	SetBody(v *DescribePrivateDnsDomainNameListResponseBody) *DescribePrivateDnsDomainNameListResponse
	GetBody() *DescribePrivateDnsDomainNameListResponseBody
}

type DescribePrivateDnsDomainNameListResponse struct {
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribePrivateDnsDomainNameListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribePrivateDnsDomainNameListResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribePrivateDnsDomainNameListResponse) GoString() string {
	return s.String()
}

func (s *DescribePrivateDnsDomainNameListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribePrivateDnsDomainNameListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribePrivateDnsDomainNameListResponse) GetBody() *DescribePrivateDnsDomainNameListResponseBody {
	return s.Body
}

func (s *DescribePrivateDnsDomainNameListResponse) SetHeaders(v map[string]*string) *DescribePrivateDnsDomainNameListResponse {
	s.Headers = v
	return s
}

func (s *DescribePrivateDnsDomainNameListResponse) SetStatusCode(v int32) *DescribePrivateDnsDomainNameListResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribePrivateDnsDomainNameListResponse) SetBody(v *DescribePrivateDnsDomainNameListResponseBody) *DescribePrivateDnsDomainNameListResponse {
	s.Body = v
	return s
}

func (s *DescribePrivateDnsDomainNameListResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
