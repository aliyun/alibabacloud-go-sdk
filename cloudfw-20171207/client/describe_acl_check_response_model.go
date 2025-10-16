// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAclCheckResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeAclCheckResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeAclCheckResponse
	GetStatusCode() *int32
	SetBody(v *DescribeAclCheckResponseBody) *DescribeAclCheckResponse
	GetBody() *DescribeAclCheckResponseBody
}

type DescribeAclCheckResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeAclCheckResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeAclCheckResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeAclCheckResponse) GoString() string {
	return s.String()
}

func (s *DescribeAclCheckResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeAclCheckResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeAclCheckResponse) GetBody() *DescribeAclCheckResponseBody {
	return s.Body
}

func (s *DescribeAclCheckResponse) SetHeaders(v map[string]*string) *DescribeAclCheckResponse {
	s.Headers = v
	return s
}

func (s *DescribeAclCheckResponse) SetStatusCode(v int32) *DescribeAclCheckResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeAclCheckResponse) SetBody(v *DescribeAclCheckResponseBody) *DescribeAclCheckResponse {
	s.Body = v
	return s
}

func (s *DescribeAclCheckResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
