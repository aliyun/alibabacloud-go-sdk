// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePreCheckStatusResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribePreCheckStatusResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribePreCheckStatusResponse
	GetStatusCode() *int32
	SetBody(v *DescribePreCheckStatusResponseBody) *DescribePreCheckStatusResponse
	GetBody() *DescribePreCheckStatusResponseBody
}

type DescribePreCheckStatusResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribePreCheckStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribePreCheckStatusResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribePreCheckStatusResponse) GoString() string {
	return s.String()
}

func (s *DescribePreCheckStatusResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribePreCheckStatusResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribePreCheckStatusResponse) GetBody() *DescribePreCheckStatusResponseBody {
	return s.Body
}

func (s *DescribePreCheckStatusResponse) SetHeaders(v map[string]*string) *DescribePreCheckStatusResponse {
	s.Headers = v
	return s
}

func (s *DescribePreCheckStatusResponse) SetStatusCode(v int32) *DescribePreCheckStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribePreCheckStatusResponse) SetBody(v *DescribePreCheckStatusResponseBody) *DescribePreCheckStatusResponse {
	s.Body = v
	return s
}

func (s *DescribePreCheckStatusResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
