// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeIPv6TranslatorsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeIPv6TranslatorsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeIPv6TranslatorsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeIPv6TranslatorsResponseBody) *DescribeIPv6TranslatorsResponse
	GetBody() *DescribeIPv6TranslatorsResponseBody
}

type DescribeIPv6TranslatorsResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeIPv6TranslatorsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeIPv6TranslatorsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeIPv6TranslatorsResponse) GoString() string {
	return s.String()
}

func (s *DescribeIPv6TranslatorsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeIPv6TranslatorsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeIPv6TranslatorsResponse) GetBody() *DescribeIPv6TranslatorsResponseBody {
	return s.Body
}

func (s *DescribeIPv6TranslatorsResponse) SetHeaders(v map[string]*string) *DescribeIPv6TranslatorsResponse {
	s.Headers = v
	return s
}

func (s *DescribeIPv6TranslatorsResponse) SetStatusCode(v int32) *DescribeIPv6TranslatorsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeIPv6TranslatorsResponse) SetBody(v *DescribeIPv6TranslatorsResponseBody) *DescribeIPv6TranslatorsResponse {
	s.Body = v
	return s
}

func (s *DescribeIPv6TranslatorsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
