// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeResourceSupportRegionsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeResourceSupportRegionsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeResourceSupportRegionsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeResourceSupportRegionsResponseBody) *DescribeResourceSupportRegionsResponse
	GetBody() *DescribeResourceSupportRegionsResponseBody
}

type DescribeResourceSupportRegionsResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeResourceSupportRegionsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeResourceSupportRegionsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeResourceSupportRegionsResponse) GoString() string {
	return s.String()
}

func (s *DescribeResourceSupportRegionsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeResourceSupportRegionsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeResourceSupportRegionsResponse) GetBody() *DescribeResourceSupportRegionsResponseBody {
	return s.Body
}

func (s *DescribeResourceSupportRegionsResponse) SetHeaders(v map[string]*string) *DescribeResourceSupportRegionsResponse {
	s.Headers = v
	return s
}

func (s *DescribeResourceSupportRegionsResponse) SetStatusCode(v int32) *DescribeResourceSupportRegionsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeResourceSupportRegionsResponse) SetBody(v *DescribeResourceSupportRegionsResponseBody) *DescribeResourceSupportRegionsResponse {
	s.Body = v
	return s
}

func (s *DescribeResourceSupportRegionsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
