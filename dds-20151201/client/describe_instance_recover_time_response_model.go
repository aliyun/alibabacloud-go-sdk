// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeInstanceRecoverTimeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeInstanceRecoverTimeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeInstanceRecoverTimeResponse
	GetStatusCode() *int32
	SetBody(v *DescribeInstanceRecoverTimeResponseBody) *DescribeInstanceRecoverTimeResponse
	GetBody() *DescribeInstanceRecoverTimeResponseBody
}

type DescribeInstanceRecoverTimeResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeInstanceRecoverTimeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeInstanceRecoverTimeResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstanceRecoverTimeResponse) GoString() string {
	return s.String()
}

func (s *DescribeInstanceRecoverTimeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeInstanceRecoverTimeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeInstanceRecoverTimeResponse) GetBody() *DescribeInstanceRecoverTimeResponseBody {
	return s.Body
}

func (s *DescribeInstanceRecoverTimeResponse) SetHeaders(v map[string]*string) *DescribeInstanceRecoverTimeResponse {
	s.Headers = v
	return s
}

func (s *DescribeInstanceRecoverTimeResponse) SetStatusCode(v int32) *DescribeInstanceRecoverTimeResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeInstanceRecoverTimeResponse) SetBody(v *DescribeInstanceRecoverTimeResponseBody) *DescribeInstanceRecoverTimeResponse {
	s.Body = v
	return s
}

func (s *DescribeInstanceRecoverTimeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
