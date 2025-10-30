// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDetailStartResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDetailStartResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDetailStartResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDetailStartResponseBody) *DescribeDetailStartResponse
	GetBody() *DescribeDetailStartResponseBody
}

type DescribeDetailStartResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDetailStartResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDetailStartResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDetailStartResponse) GoString() string {
	return s.String()
}

func (s *DescribeDetailStartResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDetailStartResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDetailStartResponse) GetBody() *DescribeDetailStartResponseBody {
	return s.Body
}

func (s *DescribeDetailStartResponse) SetHeaders(v map[string]*string) *DescribeDetailStartResponse {
	s.Headers = v
	return s
}

func (s *DescribeDetailStartResponse) SetStatusCode(v int32) *DescribeDetailStartResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDetailStartResponse) SetBody(v *DescribeDetailStartResponseBody) *DescribeDetailStartResponse {
	s.Body = v
	return s
}

func (s *DescribeDetailStartResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
