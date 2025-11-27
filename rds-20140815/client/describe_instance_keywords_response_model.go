// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeInstanceKeywordsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeInstanceKeywordsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeInstanceKeywordsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeInstanceKeywordsResponseBody) *DescribeInstanceKeywordsResponse
	GetBody() *DescribeInstanceKeywordsResponseBody
}

type DescribeInstanceKeywordsResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeInstanceKeywordsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeInstanceKeywordsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstanceKeywordsResponse) GoString() string {
	return s.String()
}

func (s *DescribeInstanceKeywordsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeInstanceKeywordsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeInstanceKeywordsResponse) GetBody() *DescribeInstanceKeywordsResponseBody {
	return s.Body
}

func (s *DescribeInstanceKeywordsResponse) SetHeaders(v map[string]*string) *DescribeInstanceKeywordsResponse {
	s.Headers = v
	return s
}

func (s *DescribeInstanceKeywordsResponse) SetStatusCode(v int32) *DescribeInstanceKeywordsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeInstanceKeywordsResponse) SetBody(v *DescribeInstanceKeywordsResponseBody) *DescribeInstanceKeywordsResponse {
	s.Body = v
	return s
}

func (s *DescribeInstanceKeywordsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
