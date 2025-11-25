// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeInternetOpenStatisticResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeInternetOpenStatisticResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeInternetOpenStatisticResponse
	GetStatusCode() *int32
	SetBody(v *DescribeInternetOpenStatisticResponseBody) *DescribeInternetOpenStatisticResponse
	GetBody() *DescribeInternetOpenStatisticResponseBody
}

type DescribeInternetOpenStatisticResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeInternetOpenStatisticResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeInternetOpenStatisticResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeInternetOpenStatisticResponse) GoString() string {
	return s.String()
}

func (s *DescribeInternetOpenStatisticResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeInternetOpenStatisticResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeInternetOpenStatisticResponse) GetBody() *DescribeInternetOpenStatisticResponseBody {
	return s.Body
}

func (s *DescribeInternetOpenStatisticResponse) SetHeaders(v map[string]*string) *DescribeInternetOpenStatisticResponse {
	s.Headers = v
	return s
}

func (s *DescribeInternetOpenStatisticResponse) SetStatusCode(v int32) *DescribeInternetOpenStatisticResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeInternetOpenStatisticResponse) SetBody(v *DescribeInternetOpenStatisticResponseBody) *DescribeInternetOpenStatisticResponse {
	s.Body = v
	return s
}

func (s *DescribeInternetOpenStatisticResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
