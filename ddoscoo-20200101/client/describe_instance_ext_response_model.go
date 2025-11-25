// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeInstanceExtResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeInstanceExtResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeInstanceExtResponse
	GetStatusCode() *int32
	SetBody(v *DescribeInstanceExtResponseBody) *DescribeInstanceExtResponse
	GetBody() *DescribeInstanceExtResponseBody
}

type DescribeInstanceExtResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeInstanceExtResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeInstanceExtResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstanceExtResponse) GoString() string {
	return s.String()
}

func (s *DescribeInstanceExtResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeInstanceExtResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeInstanceExtResponse) GetBody() *DescribeInstanceExtResponseBody {
	return s.Body
}

func (s *DescribeInstanceExtResponse) SetHeaders(v map[string]*string) *DescribeInstanceExtResponse {
	s.Headers = v
	return s
}

func (s *DescribeInstanceExtResponse) SetStatusCode(v int32) *DescribeInstanceExtResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeInstanceExtResponse) SetBody(v *DescribeInstanceExtResponseBody) *DescribeInstanceExtResponse {
	s.Body = v
	return s
}

func (s *DescribeInstanceExtResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
