// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeInstanceConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeInstanceConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeInstanceConfigResponse
	GetStatusCode() *int32
	SetBody(v *DescribeInstanceConfigResponseBody) *DescribeInstanceConfigResponse
	GetBody() *DescribeInstanceConfigResponseBody
}

type DescribeInstanceConfigResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeInstanceConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeInstanceConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstanceConfigResponse) GoString() string {
	return s.String()
}

func (s *DescribeInstanceConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeInstanceConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeInstanceConfigResponse) GetBody() *DescribeInstanceConfigResponseBody {
	return s.Body
}

func (s *DescribeInstanceConfigResponse) SetHeaders(v map[string]*string) *DescribeInstanceConfigResponse {
	s.Headers = v
	return s
}

func (s *DescribeInstanceConfigResponse) SetStatusCode(v int32) *DescribeInstanceConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeInstanceConfigResponse) SetBody(v *DescribeInstanceConfigResponseBody) *DescribeInstanceConfigResponse {
	s.Body = v
	return s
}

func (s *DescribeInstanceConfigResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
