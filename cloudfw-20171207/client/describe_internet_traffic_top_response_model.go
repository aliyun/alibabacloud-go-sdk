// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeInternetTrafficTopResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeInternetTrafficTopResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeInternetTrafficTopResponse
	GetStatusCode() *int32
	SetBody(v *DescribeInternetTrafficTopResponseBody) *DescribeInternetTrafficTopResponse
	GetBody() *DescribeInternetTrafficTopResponseBody
}

type DescribeInternetTrafficTopResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeInternetTrafficTopResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeInternetTrafficTopResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeInternetTrafficTopResponse) GoString() string {
	return s.String()
}

func (s *DescribeInternetTrafficTopResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeInternetTrafficTopResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeInternetTrafficTopResponse) GetBody() *DescribeInternetTrafficTopResponseBody {
	return s.Body
}

func (s *DescribeInternetTrafficTopResponse) SetHeaders(v map[string]*string) *DescribeInternetTrafficTopResponse {
	s.Headers = v
	return s
}

func (s *DescribeInternetTrafficTopResponse) SetStatusCode(v int32) *DescribeInternetTrafficTopResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeInternetTrafficTopResponse) SetBody(v *DescribeInternetTrafficTopResponseBody) *DescribeInternetTrafficTopResponse {
	s.Body = v
	return s
}

func (s *DescribeInternetTrafficTopResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
