// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeUserTrafficResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeUserTrafficResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeUserTrafficResponse
	GetStatusCode() *int32
	SetBody(v *DescribeUserTrafficResponseBody) *DescribeUserTrafficResponse
	GetBody() *DescribeUserTrafficResponseBody
}

type DescribeUserTrafficResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeUserTrafficResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeUserTrafficResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeUserTrafficResponse) GoString() string {
	return s.String()
}

func (s *DescribeUserTrafficResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeUserTrafficResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeUserTrafficResponse) GetBody() *DescribeUserTrafficResponseBody {
	return s.Body
}

func (s *DescribeUserTrafficResponse) SetHeaders(v map[string]*string) *DescribeUserTrafficResponse {
	s.Headers = v
	return s
}

func (s *DescribeUserTrafficResponse) SetStatusCode(v int32) *DescribeUserTrafficResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeUserTrafficResponse) SetBody(v *DescribeUserTrafficResponseBody) *DescribeUserTrafficResponse {
	s.Body = v
	return s
}

func (s *DescribeUserTrafficResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
