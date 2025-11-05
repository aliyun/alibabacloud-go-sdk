// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeUserConfigsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeUserConfigsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeUserConfigsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeUserConfigsResponseBody) *DescribeUserConfigsResponse
	GetBody() *DescribeUserConfigsResponseBody
}

type DescribeUserConfigsResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeUserConfigsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeUserConfigsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeUserConfigsResponse) GoString() string {
	return s.String()
}

func (s *DescribeUserConfigsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeUserConfigsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeUserConfigsResponse) GetBody() *DescribeUserConfigsResponseBody {
	return s.Body
}

func (s *DescribeUserConfigsResponse) SetHeaders(v map[string]*string) *DescribeUserConfigsResponse {
	s.Headers = v
	return s
}

func (s *DescribeUserConfigsResponse) SetStatusCode(v int32) *DescribeUserConfigsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeUserConfigsResponse) SetBody(v *DescribeUserConfigsResponseBody) *DescribeUserConfigsResponse {
	s.Body = v
	return s
}

func (s *DescribeUserConfigsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
