// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeInstanceKeyPairResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeInstanceKeyPairResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeInstanceKeyPairResponse
	GetStatusCode() *int32
	SetBody(v *DescribeInstanceKeyPairResponseBody) *DescribeInstanceKeyPairResponse
	GetBody() *DescribeInstanceKeyPairResponseBody
}

type DescribeInstanceKeyPairResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeInstanceKeyPairResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeInstanceKeyPairResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstanceKeyPairResponse) GoString() string {
	return s.String()
}

func (s *DescribeInstanceKeyPairResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeInstanceKeyPairResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeInstanceKeyPairResponse) GetBody() *DescribeInstanceKeyPairResponseBody {
	return s.Body
}

func (s *DescribeInstanceKeyPairResponse) SetHeaders(v map[string]*string) *DescribeInstanceKeyPairResponse {
	s.Headers = v
	return s
}

func (s *DescribeInstanceKeyPairResponse) SetStatusCode(v int32) *DescribeInstanceKeyPairResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeInstanceKeyPairResponse) SetBody(v *DescribeInstanceKeyPairResponseBody) *DescribeInstanceKeyPairResponse {
	s.Body = v
	return s
}

func (s *DescribeInstanceKeyPairResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
