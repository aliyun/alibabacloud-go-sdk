// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeApplicationStatusResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeApplicationStatusResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeApplicationStatusResponse
	GetStatusCode() *int32
	SetBody(v *DescribeApplicationStatusResponseBody) *DescribeApplicationStatusResponse
	GetBody() *DescribeApplicationStatusResponseBody
}

type DescribeApplicationStatusResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeApplicationStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeApplicationStatusResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeApplicationStatusResponse) GoString() string {
	return s.String()
}

func (s *DescribeApplicationStatusResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeApplicationStatusResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeApplicationStatusResponse) GetBody() *DescribeApplicationStatusResponseBody {
	return s.Body
}

func (s *DescribeApplicationStatusResponse) SetHeaders(v map[string]*string) *DescribeApplicationStatusResponse {
	s.Headers = v
	return s
}

func (s *DescribeApplicationStatusResponse) SetStatusCode(v int32) *DescribeApplicationStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeApplicationStatusResponse) SetBody(v *DescribeApplicationStatusResponseBody) *DescribeApplicationStatusResponse {
	s.Body = v
	return s
}

func (s *DescribeApplicationStatusResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
