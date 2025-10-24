// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeUserLogFieldConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeUserLogFieldConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeUserLogFieldConfigResponse
	GetStatusCode() *int32
	SetBody(v *DescribeUserLogFieldConfigResponseBody) *DescribeUserLogFieldConfigResponse
	GetBody() *DescribeUserLogFieldConfigResponseBody
}

type DescribeUserLogFieldConfigResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeUserLogFieldConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeUserLogFieldConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeUserLogFieldConfigResponse) GoString() string {
	return s.String()
}

func (s *DescribeUserLogFieldConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeUserLogFieldConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeUserLogFieldConfigResponse) GetBody() *DescribeUserLogFieldConfigResponseBody {
	return s.Body
}

func (s *DescribeUserLogFieldConfigResponse) SetHeaders(v map[string]*string) *DescribeUserLogFieldConfigResponse {
	s.Headers = v
	return s
}

func (s *DescribeUserLogFieldConfigResponse) SetStatusCode(v int32) *DescribeUserLogFieldConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeUserLogFieldConfigResponse) SetBody(v *DescribeUserLogFieldConfigResponseBody) *DescribeUserLogFieldConfigResponse {
	s.Body = v
	return s
}

func (s *DescribeUserLogFieldConfigResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
