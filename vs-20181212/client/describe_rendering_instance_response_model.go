// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRenderingInstanceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeRenderingInstanceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeRenderingInstanceResponse
	GetStatusCode() *int32
	SetBody(v *DescribeRenderingInstanceResponseBody) *DescribeRenderingInstanceResponse
	GetBody() *DescribeRenderingInstanceResponseBody
}

type DescribeRenderingInstanceResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeRenderingInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeRenderingInstanceResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeRenderingInstanceResponse) GoString() string {
	return s.String()
}

func (s *DescribeRenderingInstanceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeRenderingInstanceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeRenderingInstanceResponse) GetBody() *DescribeRenderingInstanceResponseBody {
	return s.Body
}

func (s *DescribeRenderingInstanceResponse) SetHeaders(v map[string]*string) *DescribeRenderingInstanceResponse {
	s.Headers = v
	return s
}

func (s *DescribeRenderingInstanceResponse) SetStatusCode(v int32) *DescribeRenderingInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeRenderingInstanceResponse) SetBody(v *DescribeRenderingInstanceResponseBody) *DescribeRenderingInstanceResponse {
	s.Body = v
	return s
}

func (s *DescribeRenderingInstanceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
