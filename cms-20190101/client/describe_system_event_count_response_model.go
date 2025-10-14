// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSystemEventCountResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeSystemEventCountResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeSystemEventCountResponse
	GetStatusCode() *int32
	SetBody(v *DescribeSystemEventCountResponseBody) *DescribeSystemEventCountResponse
	GetBody() *DescribeSystemEventCountResponseBody
}

type DescribeSystemEventCountResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeSystemEventCountResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeSystemEventCountResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeSystemEventCountResponse) GoString() string {
	return s.String()
}

func (s *DescribeSystemEventCountResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeSystemEventCountResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeSystemEventCountResponse) GetBody() *DescribeSystemEventCountResponseBody {
	return s.Body
}

func (s *DescribeSystemEventCountResponse) SetHeaders(v map[string]*string) *DescribeSystemEventCountResponse {
	s.Headers = v
	return s
}

func (s *DescribeSystemEventCountResponse) SetStatusCode(v int32) *DescribeSystemEventCountResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeSystemEventCountResponse) SetBody(v *DescribeSystemEventCountResponseBody) *DescribeSystemEventCountResponse {
	s.Body = v
	return s
}

func (s *DescribeSystemEventCountResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
