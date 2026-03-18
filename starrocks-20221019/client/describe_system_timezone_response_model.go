// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSystemTimezoneResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeSystemTimezoneResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeSystemTimezoneResponse
	GetStatusCode() *int32
	SetBody(v *DescribeSystemTimezoneResponseBody) *DescribeSystemTimezoneResponse
	GetBody() *DescribeSystemTimezoneResponseBody
}

type DescribeSystemTimezoneResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeSystemTimezoneResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeSystemTimezoneResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeSystemTimezoneResponse) GoString() string {
	return s.String()
}

func (s *DescribeSystemTimezoneResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeSystemTimezoneResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeSystemTimezoneResponse) GetBody() *DescribeSystemTimezoneResponseBody {
	return s.Body
}

func (s *DescribeSystemTimezoneResponse) SetHeaders(v map[string]*string) *DescribeSystemTimezoneResponse {
	s.Headers = v
	return s
}

func (s *DescribeSystemTimezoneResponse) SetStatusCode(v int32) *DescribeSystemTimezoneResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeSystemTimezoneResponse) SetBody(v *DescribeSystemTimezoneResponseBody) *DescribeSystemTimezoneResponse {
	s.Body = v
	return s
}

func (s *DescribeSystemTimezoneResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
