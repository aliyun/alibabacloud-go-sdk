// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAdvicesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeAdvicesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeAdvicesResponse
	GetStatusCode() *int32
	SetBody(v *DescribeAdvicesResponseBody) *DescribeAdvicesResponse
	GetBody() *DescribeAdvicesResponseBody
}

type DescribeAdvicesResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeAdvicesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeAdvicesResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeAdvicesResponse) GoString() string {
	return s.String()
}

func (s *DescribeAdvicesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeAdvicesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeAdvicesResponse) GetBody() *DescribeAdvicesResponseBody {
	return s.Body
}

func (s *DescribeAdvicesResponse) SetHeaders(v map[string]*string) *DescribeAdvicesResponse {
	s.Headers = v
	return s
}

func (s *DescribeAdvicesResponse) SetStatusCode(v int32) *DescribeAdvicesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeAdvicesResponse) SetBody(v *DescribeAdvicesResponseBody) *DescribeAdvicesResponse {
	s.Body = v
	return s
}

func (s *DescribeAdvicesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
