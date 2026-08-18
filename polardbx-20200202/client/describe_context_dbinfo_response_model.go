// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeContextDBInfoResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeContextDBInfoResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeContextDBInfoResponse
	GetStatusCode() *int32
	SetBody(v *DescribeContextDBInfoResponseBody) *DescribeContextDBInfoResponse
	GetBody() *DescribeContextDBInfoResponseBody
}

type DescribeContextDBInfoResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeContextDBInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeContextDBInfoResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeContextDBInfoResponse) GoString() string {
	return s.String()
}

func (s *DescribeContextDBInfoResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeContextDBInfoResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeContextDBInfoResponse) GetBody() *DescribeContextDBInfoResponseBody {
	return s.Body
}

func (s *DescribeContextDBInfoResponse) SetHeaders(v map[string]*string) *DescribeContextDBInfoResponse {
	s.Headers = v
	return s
}

func (s *DescribeContextDBInfoResponse) SetStatusCode(v int32) *DescribeContextDBInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeContextDBInfoResponse) SetBody(v *DescribeContextDBInfoResponseBody) *DescribeContextDBInfoResponse {
	s.Body = v
	return s
}

func (s *DescribeContextDBInfoResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
