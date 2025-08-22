// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRDDomainConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeRDDomainConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeRDDomainConfigResponse
	GetStatusCode() *int32
	SetBody(v *DescribeRDDomainConfigResponseBody) *DescribeRDDomainConfigResponse
	GetBody() *DescribeRDDomainConfigResponseBody
}

type DescribeRDDomainConfigResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeRDDomainConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeRDDomainConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeRDDomainConfigResponse) GoString() string {
	return s.String()
}

func (s *DescribeRDDomainConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeRDDomainConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeRDDomainConfigResponse) GetBody() *DescribeRDDomainConfigResponseBody {
	return s.Body
}

func (s *DescribeRDDomainConfigResponse) SetHeaders(v map[string]*string) *DescribeRDDomainConfigResponse {
	s.Headers = v
	return s
}

func (s *DescribeRDDomainConfigResponse) SetStatusCode(v int32) *DescribeRDDomainConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeRDDomainConfigResponse) SetBody(v *DescribeRDDomainConfigResponseBody) *DescribeRDDomainConfigResponse {
	s.Body = v
	return s
}

func (s *DescribeRDDomainConfigResponse) Validate() error {
	return dara.Validate(s)
}
