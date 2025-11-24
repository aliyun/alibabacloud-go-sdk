// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeReusableSlbResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeReusableSlbResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeReusableSlbResponse
	GetStatusCode() *int32
	SetBody(v *DescribeReusableSlbResponseBody) *DescribeReusableSlbResponse
	GetBody() *DescribeReusableSlbResponseBody
}

type DescribeReusableSlbResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeReusableSlbResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeReusableSlbResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeReusableSlbResponse) GoString() string {
	return s.String()
}

func (s *DescribeReusableSlbResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeReusableSlbResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeReusableSlbResponse) GetBody() *DescribeReusableSlbResponseBody {
	return s.Body
}

func (s *DescribeReusableSlbResponse) SetHeaders(v map[string]*string) *DescribeReusableSlbResponse {
	s.Headers = v
	return s
}

func (s *DescribeReusableSlbResponse) SetStatusCode(v int32) *DescribeReusableSlbResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeReusableSlbResponse) SetBody(v *DescribeReusableSlbResponseBody) *DescribeReusableSlbResponse {
	s.Body = v
	return s
}

func (s *DescribeReusableSlbResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
