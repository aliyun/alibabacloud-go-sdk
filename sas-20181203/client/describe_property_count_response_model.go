// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePropertyCountResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribePropertyCountResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribePropertyCountResponse
	GetStatusCode() *int32
	SetBody(v *DescribePropertyCountResponseBody) *DescribePropertyCountResponse
	GetBody() *DescribePropertyCountResponseBody
}

type DescribePropertyCountResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribePropertyCountResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribePropertyCountResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribePropertyCountResponse) GoString() string {
	return s.String()
}

func (s *DescribePropertyCountResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribePropertyCountResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribePropertyCountResponse) GetBody() *DescribePropertyCountResponseBody {
	return s.Body
}

func (s *DescribePropertyCountResponse) SetHeaders(v map[string]*string) *DescribePropertyCountResponse {
	s.Headers = v
	return s
}

func (s *DescribePropertyCountResponse) SetStatusCode(v int32) *DescribePropertyCountResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribePropertyCountResponse) SetBody(v *DescribePropertyCountResponseBody) *DescribePropertyCountResponse {
	s.Body = v
	return s
}

func (s *DescribePropertyCountResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
