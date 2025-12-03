// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSignaturesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeSignaturesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeSignaturesResponse
	GetStatusCode() *int32
	SetBody(v *DescribeSignaturesResponseBody) *DescribeSignaturesResponse
	GetBody() *DescribeSignaturesResponseBody
}

type DescribeSignaturesResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeSignaturesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeSignaturesResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeSignaturesResponse) GoString() string {
	return s.String()
}

func (s *DescribeSignaturesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeSignaturesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeSignaturesResponse) GetBody() *DescribeSignaturesResponseBody {
	return s.Body
}

func (s *DescribeSignaturesResponse) SetHeaders(v map[string]*string) *DescribeSignaturesResponse {
	s.Headers = v
	return s
}

func (s *DescribeSignaturesResponse) SetStatusCode(v int32) *DescribeSignaturesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeSignaturesResponse) SetBody(v *DescribeSignaturesResponseBody) *DescribeSignaturesResponse {
	s.Body = v
	return s
}

func (s *DescribeSignaturesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
