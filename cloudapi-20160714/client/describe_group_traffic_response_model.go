// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeGroupTrafficResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeGroupTrafficResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeGroupTrafficResponse
	GetStatusCode() *int32
	SetBody(v *DescribeGroupTrafficResponseBody) *DescribeGroupTrafficResponse
	GetBody() *DescribeGroupTrafficResponseBody
}

type DescribeGroupTrafficResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeGroupTrafficResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeGroupTrafficResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeGroupTrafficResponse) GoString() string {
	return s.String()
}

func (s *DescribeGroupTrafficResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeGroupTrafficResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeGroupTrafficResponse) GetBody() *DescribeGroupTrafficResponseBody {
	return s.Body
}

func (s *DescribeGroupTrafficResponse) SetHeaders(v map[string]*string) *DescribeGroupTrafficResponse {
	s.Headers = v
	return s
}

func (s *DescribeGroupTrafficResponse) SetStatusCode(v int32) *DescribeGroupTrafficResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeGroupTrafficResponse) SetBody(v *DescribeGroupTrafficResponseBody) *DescribeGroupTrafficResponse {
	s.Body = v
	return s
}

func (s *DescribeGroupTrafficResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
