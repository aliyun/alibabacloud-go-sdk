// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAccessConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeAccessConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeAccessConfigResponse
	GetStatusCode() *int32
	SetBody(v *DescribeAccessConfigResponseBody) *DescribeAccessConfigResponse
	GetBody() *DescribeAccessConfigResponseBody
}

type DescribeAccessConfigResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeAccessConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeAccessConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeAccessConfigResponse) GoString() string {
	return s.String()
}

func (s *DescribeAccessConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeAccessConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeAccessConfigResponse) GetBody() *DescribeAccessConfigResponseBody {
	return s.Body
}

func (s *DescribeAccessConfigResponse) SetHeaders(v map[string]*string) *DescribeAccessConfigResponse {
	s.Headers = v
	return s
}

func (s *DescribeAccessConfigResponse) SetStatusCode(v int32) *DescribeAccessConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeAccessConfigResponse) SetBody(v *DescribeAccessConfigResponseBody) *DescribeAccessConfigResponse {
	s.Body = v
	return s
}

func (s *DescribeAccessConfigResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
