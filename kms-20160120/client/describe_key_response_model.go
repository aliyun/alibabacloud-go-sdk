// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeKeyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeKeyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeKeyResponse
	GetStatusCode() *int32
	SetBody(v *DescribeKeyResponseBody) *DescribeKeyResponse
	GetBody() *DescribeKeyResponseBody
}

type DescribeKeyResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeKeyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeKeyResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeKeyResponse) GoString() string {
	return s.String()
}

func (s *DescribeKeyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeKeyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeKeyResponse) GetBody() *DescribeKeyResponseBody {
	return s.Body
}

func (s *DescribeKeyResponse) SetHeaders(v map[string]*string) *DescribeKeyResponse {
	s.Headers = v
	return s
}

func (s *DescribeKeyResponse) SetStatusCode(v int32) *DescribeKeyResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeKeyResponse) SetBody(v *DescribeKeyResponseBody) *DescribeKeyResponse {
	s.Body = v
	return s
}

func (s *DescribeKeyResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
