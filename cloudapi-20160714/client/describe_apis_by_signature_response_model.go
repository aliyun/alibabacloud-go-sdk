// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeApisBySignatureResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeApisBySignatureResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeApisBySignatureResponse
	GetStatusCode() *int32
	SetBody(v *DescribeApisBySignatureResponseBody) *DescribeApisBySignatureResponse
	GetBody() *DescribeApisBySignatureResponseBody
}

type DescribeApisBySignatureResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeApisBySignatureResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeApisBySignatureResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeApisBySignatureResponse) GoString() string {
	return s.String()
}

func (s *DescribeApisBySignatureResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeApisBySignatureResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeApisBySignatureResponse) GetBody() *DescribeApisBySignatureResponseBody {
	return s.Body
}

func (s *DescribeApisBySignatureResponse) SetHeaders(v map[string]*string) *DescribeApisBySignatureResponse {
	s.Headers = v
	return s
}

func (s *DescribeApisBySignatureResponse) SetStatusCode(v int32) *DescribeApisBySignatureResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeApisBySignatureResponse) SetBody(v *DescribeApisBySignatureResponseBody) *DescribeApisBySignatureResponse {
	s.Body = v
	return s
}

func (s *DescribeApisBySignatureResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
