// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeResourceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeResourceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeResourceResponse
	GetStatusCode() *int32
	SetBody(v *DescribeResourceResponseBody) *DescribeResourceResponse
	GetBody() *DescribeResourceResponseBody
}

type DescribeResourceResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeResourceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeResourceResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeResourceResponse) GoString() string {
	return s.String()
}

func (s *DescribeResourceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeResourceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeResourceResponse) GetBody() *DescribeResourceResponseBody {
	return s.Body
}

func (s *DescribeResourceResponse) SetHeaders(v map[string]*string) *DescribeResourceResponse {
	s.Headers = v
	return s
}

func (s *DescribeResourceResponse) SetStatusCode(v int32) *DescribeResourceResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeResourceResponse) SetBody(v *DescribeResourceResponseBody) *DescribeResourceResponse {
	s.Body = v
	return s
}

func (s *DescribeResourceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
