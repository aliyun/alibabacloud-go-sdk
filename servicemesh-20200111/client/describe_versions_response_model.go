// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVersionsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeVersionsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeVersionsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeVersionsResponseBody) *DescribeVersionsResponse
	GetBody() *DescribeVersionsResponseBody
}

type DescribeVersionsResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeVersionsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeVersionsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeVersionsResponse) GoString() string {
	return s.String()
}

func (s *DescribeVersionsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeVersionsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeVersionsResponse) GetBody() *DescribeVersionsResponseBody {
	return s.Body
}

func (s *DescribeVersionsResponse) SetHeaders(v map[string]*string) *DescribeVersionsResponse {
	s.Headers = v
	return s
}

func (s *DescribeVersionsResponse) SetStatusCode(v int32) *DescribeVersionsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeVersionsResponse) SetBody(v *DescribeVersionsResponseBody) *DescribeVersionsResponse {
	s.Body = v
	return s
}

func (s *DescribeVersionsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
