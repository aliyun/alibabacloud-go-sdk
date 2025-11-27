// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeStreamURLResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeStreamURLResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeStreamURLResponse
	GetStatusCode() *int32
	SetBody(v *DescribeStreamURLResponseBody) *DescribeStreamURLResponse
	GetBody() *DescribeStreamURLResponseBody
}

type DescribeStreamURLResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeStreamURLResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeStreamURLResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeStreamURLResponse) GoString() string {
	return s.String()
}

func (s *DescribeStreamURLResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeStreamURLResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeStreamURLResponse) GetBody() *DescribeStreamURLResponseBody {
	return s.Body
}

func (s *DescribeStreamURLResponse) SetHeaders(v map[string]*string) *DescribeStreamURLResponse {
	s.Headers = v
	return s
}

func (s *DescribeStreamURLResponse) SetStatusCode(v int32) *DescribeStreamURLResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeStreamURLResponse) SetBody(v *DescribeStreamURLResponseBody) *DescribeStreamURLResponse {
	s.Body = v
	return s
}

func (s *DescribeStreamURLResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
