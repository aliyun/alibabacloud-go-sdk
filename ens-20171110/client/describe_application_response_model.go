// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeApplicationResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeApplicationResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeApplicationResponse
	GetStatusCode() *int32
	SetBody(v *DescribeApplicationResponseBody) *DescribeApplicationResponse
	GetBody() *DescribeApplicationResponseBody
}

type DescribeApplicationResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeApplicationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeApplicationResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeApplicationResponse) GoString() string {
	return s.String()
}

func (s *DescribeApplicationResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeApplicationResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeApplicationResponse) GetBody() *DescribeApplicationResponseBody {
	return s.Body
}

func (s *DescribeApplicationResponse) SetHeaders(v map[string]*string) *DescribeApplicationResponse {
	s.Headers = v
	return s
}

func (s *DescribeApplicationResponse) SetStatusCode(v int32) *DescribeApplicationResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeApplicationResponse) SetBody(v *DescribeApplicationResponseBody) *DescribeApplicationResponse {
	s.Body = v
	return s
}

func (s *DescribeApplicationResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
