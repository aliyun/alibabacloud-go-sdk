// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeResourceTypeAutoEnableResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeResourceTypeAutoEnableResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeResourceTypeAutoEnableResponse
	GetStatusCode() *int32
	SetBody(v *DescribeResourceTypeAutoEnableResponseBody) *DescribeResourceTypeAutoEnableResponse
	GetBody() *DescribeResourceTypeAutoEnableResponseBody
}

type DescribeResourceTypeAutoEnableResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeResourceTypeAutoEnableResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeResourceTypeAutoEnableResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeResourceTypeAutoEnableResponse) GoString() string {
	return s.String()
}

func (s *DescribeResourceTypeAutoEnableResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeResourceTypeAutoEnableResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeResourceTypeAutoEnableResponse) GetBody() *DescribeResourceTypeAutoEnableResponseBody {
	return s.Body
}

func (s *DescribeResourceTypeAutoEnableResponse) SetHeaders(v map[string]*string) *DescribeResourceTypeAutoEnableResponse {
	s.Headers = v
	return s
}

func (s *DescribeResourceTypeAutoEnableResponse) SetStatusCode(v int32) *DescribeResourceTypeAutoEnableResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeResourceTypeAutoEnableResponse) SetBody(v *DescribeResourceTypeAutoEnableResponseBody) *DescribeResourceTypeAutoEnableResponse {
	s.Body = v
	return s
}

func (s *DescribeResourceTypeAutoEnableResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
