// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeTransferDomainsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeTransferDomainsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeTransferDomainsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeTransferDomainsResponseBody) *DescribeTransferDomainsResponse
	GetBody() *DescribeTransferDomainsResponseBody
}

type DescribeTransferDomainsResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeTransferDomainsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeTransferDomainsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeTransferDomainsResponse) GoString() string {
	return s.String()
}

func (s *DescribeTransferDomainsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeTransferDomainsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeTransferDomainsResponse) GetBody() *DescribeTransferDomainsResponseBody {
	return s.Body
}

func (s *DescribeTransferDomainsResponse) SetHeaders(v map[string]*string) *DescribeTransferDomainsResponse {
	s.Headers = v
	return s
}

func (s *DescribeTransferDomainsResponse) SetStatusCode(v int32) *DescribeTransferDomainsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeTransferDomainsResponse) SetBody(v *DescribeTransferDomainsResponseBody) *DescribeTransferDomainsResponse {
	s.Body = v
	return s
}

func (s *DescribeTransferDomainsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
