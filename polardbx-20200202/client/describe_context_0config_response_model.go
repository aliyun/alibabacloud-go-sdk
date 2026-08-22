// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeContext0ConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeContext0ConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeContext0ConfigResponse
	GetStatusCode() *int32
	SetBody(v *DescribeContext0ConfigResponseBody) *DescribeContext0ConfigResponse
	GetBody() *DescribeContext0ConfigResponseBody
}

type DescribeContext0ConfigResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeContext0ConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeContext0ConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeContext0ConfigResponse) GoString() string {
	return s.String()
}

func (s *DescribeContext0ConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeContext0ConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeContext0ConfigResponse) GetBody() *DescribeContext0ConfigResponseBody {
	return s.Body
}

func (s *DescribeContext0ConfigResponse) SetHeaders(v map[string]*string) *DescribeContext0ConfigResponse {
	s.Headers = v
	return s
}

func (s *DescribeContext0ConfigResponse) SetStatusCode(v int32) *DescribeContext0ConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeContext0ConfigResponse) SetBody(v *DescribeContext0ConfigResponseBody) *DescribeContext0ConfigResponse {
	s.Body = v
	return s
}

func (s *DescribeContext0ConfigResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
