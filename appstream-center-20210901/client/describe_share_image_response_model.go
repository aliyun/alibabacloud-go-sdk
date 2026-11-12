// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeShareImageResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeShareImageResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeShareImageResponse
	GetStatusCode() *int32
	SetBody(v *DescribeShareImageResponseBody) *DescribeShareImageResponse
	GetBody() *DescribeShareImageResponseBody
}

type DescribeShareImageResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeShareImageResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeShareImageResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeShareImageResponse) GoString() string {
	return s.String()
}

func (s *DescribeShareImageResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeShareImageResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeShareImageResponse) GetBody() *DescribeShareImageResponseBody {
	return s.Body
}

func (s *DescribeShareImageResponse) SetHeaders(v map[string]*string) *DescribeShareImageResponse {
	s.Headers = v
	return s
}

func (s *DescribeShareImageResponse) SetStatusCode(v int32) *DescribeShareImageResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeShareImageResponse) SetBody(v *DescribeShareImageResponseBody) *DescribeShareImageResponse {
	s.Body = v
	return s
}

func (s *DescribeShareImageResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
