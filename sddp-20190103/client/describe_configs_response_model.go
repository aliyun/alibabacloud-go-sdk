// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeConfigsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeConfigsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeConfigsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeConfigsResponseBody) *DescribeConfigsResponse
	GetBody() *DescribeConfigsResponseBody
}

type DescribeConfigsResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeConfigsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeConfigsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeConfigsResponse) GoString() string {
	return s.String()
}

func (s *DescribeConfigsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeConfigsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeConfigsResponse) GetBody() *DescribeConfigsResponseBody {
	return s.Body
}

func (s *DescribeConfigsResponse) SetHeaders(v map[string]*string) *DescribeConfigsResponse {
	s.Headers = v
	return s
}

func (s *DescribeConfigsResponse) SetStatusCode(v int32) *DescribeConfigsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeConfigsResponse) SetBody(v *DescribeConfigsResponseBody) *DescribeConfigsResponse {
	s.Body = v
	return s
}

func (s *DescribeConfigsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
