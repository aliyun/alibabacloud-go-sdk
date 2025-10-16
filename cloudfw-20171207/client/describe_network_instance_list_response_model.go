// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeNetworkInstanceListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeNetworkInstanceListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeNetworkInstanceListResponse
	GetStatusCode() *int32
	SetBody(v *DescribeNetworkInstanceListResponseBody) *DescribeNetworkInstanceListResponse
	GetBody() *DescribeNetworkInstanceListResponseBody
}

type DescribeNetworkInstanceListResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeNetworkInstanceListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeNetworkInstanceListResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeNetworkInstanceListResponse) GoString() string {
	return s.String()
}

func (s *DescribeNetworkInstanceListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeNetworkInstanceListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeNetworkInstanceListResponse) GetBody() *DescribeNetworkInstanceListResponseBody {
	return s.Body
}

func (s *DescribeNetworkInstanceListResponse) SetHeaders(v map[string]*string) *DescribeNetworkInstanceListResponse {
	s.Headers = v
	return s
}

func (s *DescribeNetworkInstanceListResponse) SetStatusCode(v int32) *DescribeNetworkInstanceListResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeNetworkInstanceListResponse) SetBody(v *DescribeNetworkInstanceListResponseBody) *DescribeNetworkInstanceListResponse {
	s.Body = v
	return s
}

func (s *DescribeNetworkInstanceListResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
