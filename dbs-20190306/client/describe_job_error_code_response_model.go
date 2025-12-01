// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeJobErrorCodeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeJobErrorCodeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeJobErrorCodeResponse
	GetStatusCode() *int32
	SetBody(v *DescribeJobErrorCodeResponseBody) *DescribeJobErrorCodeResponse
	GetBody() *DescribeJobErrorCodeResponseBody
}

type DescribeJobErrorCodeResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeJobErrorCodeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeJobErrorCodeResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeJobErrorCodeResponse) GoString() string {
	return s.String()
}

func (s *DescribeJobErrorCodeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeJobErrorCodeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeJobErrorCodeResponse) GetBody() *DescribeJobErrorCodeResponseBody {
	return s.Body
}

func (s *DescribeJobErrorCodeResponse) SetHeaders(v map[string]*string) *DescribeJobErrorCodeResponse {
	s.Headers = v
	return s
}

func (s *DescribeJobErrorCodeResponse) SetStatusCode(v int32) *DescribeJobErrorCodeResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeJobErrorCodeResponse) SetBody(v *DescribeJobErrorCodeResponseBody) *DescribeJobErrorCodeResponse {
	s.Body = v
	return s
}

func (s *DescribeJobErrorCodeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
