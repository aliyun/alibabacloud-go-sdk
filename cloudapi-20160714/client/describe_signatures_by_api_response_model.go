// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSignaturesByApiResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeSignaturesByApiResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeSignaturesByApiResponse
	GetStatusCode() *int32
	SetBody(v *DescribeSignaturesByApiResponseBody) *DescribeSignaturesByApiResponse
	GetBody() *DescribeSignaturesByApiResponseBody
}

type DescribeSignaturesByApiResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeSignaturesByApiResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeSignaturesByApiResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeSignaturesByApiResponse) GoString() string {
	return s.String()
}

func (s *DescribeSignaturesByApiResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeSignaturesByApiResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeSignaturesByApiResponse) GetBody() *DescribeSignaturesByApiResponseBody {
	return s.Body
}

func (s *DescribeSignaturesByApiResponse) SetHeaders(v map[string]*string) *DescribeSignaturesByApiResponse {
	s.Headers = v
	return s
}

func (s *DescribeSignaturesByApiResponse) SetStatusCode(v int32) *DescribeSignaturesByApiResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeSignaturesByApiResponse) SetBody(v *DescribeSignaturesByApiResponseBody) *DescribeSignaturesByApiResponse {
	s.Body = v
	return s
}

func (s *DescribeSignaturesByApiResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
