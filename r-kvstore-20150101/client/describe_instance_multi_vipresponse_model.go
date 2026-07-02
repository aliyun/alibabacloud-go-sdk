// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeInstanceMultiVIPResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeInstanceMultiVIPResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeInstanceMultiVIPResponse
	GetStatusCode() *int32
	SetBody(v *DescribeInstanceMultiVIPResponseBody) *DescribeInstanceMultiVIPResponse
	GetBody() *DescribeInstanceMultiVIPResponseBody
}

type DescribeInstanceMultiVIPResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeInstanceMultiVIPResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeInstanceMultiVIPResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstanceMultiVIPResponse) GoString() string {
	return s.String()
}

func (s *DescribeInstanceMultiVIPResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeInstanceMultiVIPResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeInstanceMultiVIPResponse) GetBody() *DescribeInstanceMultiVIPResponseBody {
	return s.Body
}

func (s *DescribeInstanceMultiVIPResponse) SetHeaders(v map[string]*string) *DescribeInstanceMultiVIPResponse {
	s.Headers = v
	return s
}

func (s *DescribeInstanceMultiVIPResponse) SetStatusCode(v int32) *DescribeInstanceMultiVIPResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeInstanceMultiVIPResponse) SetBody(v *DescribeInstanceMultiVIPResponseBody) *DescribeInstanceMultiVIPResponse {
	s.Body = v
	return s
}

func (s *DescribeInstanceMultiVIPResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
