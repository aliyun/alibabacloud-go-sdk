// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeTrailsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeTrailsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeTrailsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeTrailsResponseBody) *DescribeTrailsResponse
	GetBody() *DescribeTrailsResponseBody
}

type DescribeTrailsResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeTrailsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeTrailsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeTrailsResponse) GoString() string {
	return s.String()
}

func (s *DescribeTrailsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeTrailsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeTrailsResponse) GetBody() *DescribeTrailsResponseBody {
	return s.Body
}

func (s *DescribeTrailsResponse) SetHeaders(v map[string]*string) *DescribeTrailsResponse {
	s.Headers = v
	return s
}

func (s *DescribeTrailsResponse) SetStatusCode(v int32) *DescribeTrailsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeTrailsResponse) SetBody(v *DescribeTrailsResponseBody) *DescribeTrailsResponse {
	s.Body = v
	return s
}

func (s *DescribeTrailsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
