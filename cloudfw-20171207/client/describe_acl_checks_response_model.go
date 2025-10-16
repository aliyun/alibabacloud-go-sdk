// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAclChecksResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeAclChecksResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeAclChecksResponse
	GetStatusCode() *int32
	SetBody(v *DescribeAclChecksResponseBody) *DescribeAclChecksResponse
	GetBody() *DescribeAclChecksResponseBody
}

type DescribeAclChecksResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeAclChecksResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeAclChecksResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeAclChecksResponse) GoString() string {
	return s.String()
}

func (s *DescribeAclChecksResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeAclChecksResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeAclChecksResponse) GetBody() *DescribeAclChecksResponseBody {
	return s.Body
}

func (s *DescribeAclChecksResponse) SetHeaders(v map[string]*string) *DescribeAclChecksResponse {
	s.Headers = v
	return s
}

func (s *DescribeAclChecksResponse) SetStatusCode(v int32) *DescribeAclChecksResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeAclChecksResponse) SetBody(v *DescribeAclChecksResponseBody) *DescribeAclChecksResponse {
	s.Body = v
	return s
}

func (s *DescribeAclChecksResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
