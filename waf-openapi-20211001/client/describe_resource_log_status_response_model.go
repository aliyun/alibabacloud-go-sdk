// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeResourceLogStatusResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeResourceLogStatusResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeResourceLogStatusResponse
	GetStatusCode() *int32
	SetBody(v *DescribeResourceLogStatusResponseBody) *DescribeResourceLogStatusResponse
	GetBody() *DescribeResourceLogStatusResponseBody
}

type DescribeResourceLogStatusResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeResourceLogStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeResourceLogStatusResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeResourceLogStatusResponse) GoString() string {
	return s.String()
}

func (s *DescribeResourceLogStatusResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeResourceLogStatusResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeResourceLogStatusResponse) GetBody() *DescribeResourceLogStatusResponseBody {
	return s.Body
}

func (s *DescribeResourceLogStatusResponse) SetHeaders(v map[string]*string) *DescribeResourceLogStatusResponse {
	s.Headers = v
	return s
}

func (s *DescribeResourceLogStatusResponse) SetStatusCode(v int32) *DescribeResourceLogStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeResourceLogStatusResponse) SetBody(v *DescribeResourceLogStatusResponseBody) *DescribeResourceLogStatusResponse {
	s.Body = v
	return s
}

func (s *DescribeResourceLogStatusResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
