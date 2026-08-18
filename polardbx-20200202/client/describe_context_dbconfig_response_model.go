// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeContextDBConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeContextDBConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeContextDBConfigResponse
	GetStatusCode() *int32
	SetBody(v *DescribeContextDBConfigResponseBody) *DescribeContextDBConfigResponse
	GetBody() *DescribeContextDBConfigResponseBody
}

type DescribeContextDBConfigResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeContextDBConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeContextDBConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeContextDBConfigResponse) GoString() string {
	return s.String()
}

func (s *DescribeContextDBConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeContextDBConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeContextDBConfigResponse) GetBody() *DescribeContextDBConfigResponseBody {
	return s.Body
}

func (s *DescribeContextDBConfigResponse) SetHeaders(v map[string]*string) *DescribeContextDBConfigResponse {
	s.Headers = v
	return s
}

func (s *DescribeContextDBConfigResponse) SetStatusCode(v int32) *DescribeContextDBConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeContextDBConfigResponse) SetBody(v *DescribeContextDBConfigResponseBody) *DescribeContextDBConfigResponse {
	s.Body = v
	return s
}

func (s *DescribeContextDBConfigResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
