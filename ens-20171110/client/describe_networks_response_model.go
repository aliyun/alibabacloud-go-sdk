// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeNetworksResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeNetworksResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeNetworksResponse
	GetStatusCode() *int32
	SetBody(v *DescribeNetworksResponseBody) *DescribeNetworksResponse
	GetBody() *DescribeNetworksResponseBody
}

type DescribeNetworksResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeNetworksResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeNetworksResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeNetworksResponse) GoString() string {
	return s.String()
}

func (s *DescribeNetworksResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeNetworksResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeNetworksResponse) GetBody() *DescribeNetworksResponseBody {
	return s.Body
}

func (s *DescribeNetworksResponse) SetHeaders(v map[string]*string) *DescribeNetworksResponse {
	s.Headers = v
	return s
}

func (s *DescribeNetworksResponse) SetStatusCode(v int32) *DescribeNetworksResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeNetworksResponse) SetBody(v *DescribeNetworksResponseBody) *DescribeNetworksResponse {
	s.Body = v
	return s
}

func (s *DescribeNetworksResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
