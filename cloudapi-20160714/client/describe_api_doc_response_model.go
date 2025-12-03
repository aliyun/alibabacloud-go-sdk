// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeApiDocResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeApiDocResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeApiDocResponse
	GetStatusCode() *int32
	SetBody(v *DescribeApiDocResponseBody) *DescribeApiDocResponse
	GetBody() *DescribeApiDocResponseBody
}

type DescribeApiDocResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeApiDocResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeApiDocResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeApiDocResponse) GoString() string {
	return s.String()
}

func (s *DescribeApiDocResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeApiDocResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeApiDocResponse) GetBody() *DescribeApiDocResponseBody {
	return s.Body
}

func (s *DescribeApiDocResponse) SetHeaders(v map[string]*string) *DescribeApiDocResponse {
	s.Headers = v
	return s
}

func (s *DescribeApiDocResponse) SetStatusCode(v int32) *DescribeApiDocResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeApiDocResponse) SetBody(v *DescribeApiDocResponseBody) *DescribeApiDocResponse {
	s.Body = v
	return s
}

func (s *DescribeApiDocResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
