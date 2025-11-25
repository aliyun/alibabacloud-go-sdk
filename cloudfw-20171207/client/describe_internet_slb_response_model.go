// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeInternetSlbResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeInternetSlbResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeInternetSlbResponse
	GetStatusCode() *int32
	SetBody(v *DescribeInternetSlbResponseBody) *DescribeInternetSlbResponse
	GetBody() *DescribeInternetSlbResponseBody
}

type DescribeInternetSlbResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeInternetSlbResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeInternetSlbResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeInternetSlbResponse) GoString() string {
	return s.String()
}

func (s *DescribeInternetSlbResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeInternetSlbResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeInternetSlbResponse) GetBody() *DescribeInternetSlbResponseBody {
	return s.Body
}

func (s *DescribeInternetSlbResponse) SetHeaders(v map[string]*string) *DescribeInternetSlbResponse {
	s.Headers = v
	return s
}

func (s *DescribeInternetSlbResponse) SetStatusCode(v int32) *DescribeInternetSlbResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeInternetSlbResponse) SetBody(v *DescribeInternetSlbResponseBody) *DescribeInternetSlbResponse {
	s.Body = v
	return s
}

func (s *DescribeInternetSlbResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
