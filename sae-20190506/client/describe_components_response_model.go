// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeComponentsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeComponentsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeComponentsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeComponentsResponseBody) *DescribeComponentsResponse
	GetBody() *DescribeComponentsResponseBody
}

type DescribeComponentsResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeComponentsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeComponentsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeComponentsResponse) GoString() string {
	return s.String()
}

func (s *DescribeComponentsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeComponentsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeComponentsResponse) GetBody() *DescribeComponentsResponseBody {
	return s.Body
}

func (s *DescribeComponentsResponse) SetHeaders(v map[string]*string) *DescribeComponentsResponse {
	s.Headers = v
	return s
}

func (s *DescribeComponentsResponse) SetStatusCode(v int32) *DescribeComponentsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeComponentsResponse) SetBody(v *DescribeComponentsResponseBody) *DescribeComponentsResponse {
	s.Body = v
	return s
}

func (s *DescribeComponentsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
