// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLogConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeLogConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeLogConfigResponse
	GetStatusCode() *int32
	SetBody(v *DescribeLogConfigResponseBody) *DescribeLogConfigResponse
	GetBody() *DescribeLogConfigResponseBody
}

type DescribeLogConfigResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeLogConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeLogConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeLogConfigResponse) GoString() string {
	return s.String()
}

func (s *DescribeLogConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeLogConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeLogConfigResponse) GetBody() *DescribeLogConfigResponseBody {
	return s.Body
}

func (s *DescribeLogConfigResponse) SetHeaders(v map[string]*string) *DescribeLogConfigResponse {
	s.Headers = v
	return s
}

func (s *DescribeLogConfigResponse) SetStatusCode(v int32) *DescribeLogConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeLogConfigResponse) SetBody(v *DescribeLogConfigResponseBody) *DescribeLogConfigResponse {
	s.Body = v
	return s
}

func (s *DescribeLogConfigResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
