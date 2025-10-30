// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeResourceLogFieldConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeResourceLogFieldConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeResourceLogFieldConfigResponse
	GetStatusCode() *int32
	SetBody(v *DescribeResourceLogFieldConfigResponseBody) *DescribeResourceLogFieldConfigResponse
	GetBody() *DescribeResourceLogFieldConfigResponseBody
}

type DescribeResourceLogFieldConfigResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeResourceLogFieldConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeResourceLogFieldConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeResourceLogFieldConfigResponse) GoString() string {
	return s.String()
}

func (s *DescribeResourceLogFieldConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeResourceLogFieldConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeResourceLogFieldConfigResponse) GetBody() *DescribeResourceLogFieldConfigResponseBody {
	return s.Body
}

func (s *DescribeResourceLogFieldConfigResponse) SetHeaders(v map[string]*string) *DescribeResourceLogFieldConfigResponse {
	s.Headers = v
	return s
}

func (s *DescribeResourceLogFieldConfigResponse) SetStatusCode(v int32) *DescribeResourceLogFieldConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeResourceLogFieldConfigResponse) SetBody(v *DescribeResourceLogFieldConfigResponseBody) *DescribeResourceLogFieldConfigResponse {
	s.Body = v
	return s
}

func (s *DescribeResourceLogFieldConfigResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
