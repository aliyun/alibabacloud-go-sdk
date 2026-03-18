// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeTrafficResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeTrafficResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeTrafficResponse
	GetStatusCode() *int32
	SetBody(v *DescribeTrafficResponseBody) *DescribeTrafficResponse
	GetBody() *DescribeTrafficResponseBody
}

type DescribeTrafficResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeTrafficResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeTrafficResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeTrafficResponse) GoString() string {
	return s.String()
}

func (s *DescribeTrafficResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeTrafficResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeTrafficResponse) GetBody() *DescribeTrafficResponseBody {
	return s.Body
}

func (s *DescribeTrafficResponse) SetHeaders(v map[string]*string) *DescribeTrafficResponse {
	s.Headers = v
	return s
}

func (s *DescribeTrafficResponse) SetStatusCode(v int32) *DescribeTrafficResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeTrafficResponse) SetBody(v *DescribeTrafficResponseBody) *DescribeTrafficResponse {
	s.Body = v
	return s
}

func (s *DescribeTrafficResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
