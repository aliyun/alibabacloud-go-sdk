// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeNetworkTrafficTopRatioResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeNetworkTrafficTopRatioResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeNetworkTrafficTopRatioResponse
	GetStatusCode() *int32
	SetBody(v *DescribeNetworkTrafficTopRatioResponseBody) *DescribeNetworkTrafficTopRatioResponse
	GetBody() *DescribeNetworkTrafficTopRatioResponseBody
}

type DescribeNetworkTrafficTopRatioResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeNetworkTrafficTopRatioResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeNetworkTrafficTopRatioResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeNetworkTrafficTopRatioResponse) GoString() string {
	return s.String()
}

func (s *DescribeNetworkTrafficTopRatioResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeNetworkTrafficTopRatioResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeNetworkTrafficTopRatioResponse) GetBody() *DescribeNetworkTrafficTopRatioResponseBody {
	return s.Body
}

func (s *DescribeNetworkTrafficTopRatioResponse) SetHeaders(v map[string]*string) *DescribeNetworkTrafficTopRatioResponse {
	s.Headers = v
	return s
}

func (s *DescribeNetworkTrafficTopRatioResponse) SetStatusCode(v int32) *DescribeNetworkTrafficTopRatioResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeNetworkTrafficTopRatioResponse) SetBody(v *DescribeNetworkTrafficTopRatioResponseBody) *DescribeNetworkTrafficTopRatioResponse {
	s.Body = v
	return s
}

func (s *DescribeNetworkTrafficTopRatioResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
