// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeOutgoingDomainResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeOutgoingDomainResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeOutgoingDomainResponse
	GetStatusCode() *int32
	SetBody(v *DescribeOutgoingDomainResponseBody) *DescribeOutgoingDomainResponse
	GetBody() *DescribeOutgoingDomainResponseBody
}

type DescribeOutgoingDomainResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeOutgoingDomainResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeOutgoingDomainResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeOutgoingDomainResponse) GoString() string {
	return s.String()
}

func (s *DescribeOutgoingDomainResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeOutgoingDomainResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeOutgoingDomainResponse) GetBody() *DescribeOutgoingDomainResponseBody {
	return s.Body
}

func (s *DescribeOutgoingDomainResponse) SetHeaders(v map[string]*string) *DescribeOutgoingDomainResponse {
	s.Headers = v
	return s
}

func (s *DescribeOutgoingDomainResponse) SetStatusCode(v int32) *DescribeOutgoingDomainResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeOutgoingDomainResponse) SetBody(v *DescribeOutgoingDomainResponseBody) *DescribeOutgoingDomainResponse {
	s.Body = v
	return s
}

func (s *DescribeOutgoingDomainResponse) Validate() error {
	return dara.Validate(s)
}
