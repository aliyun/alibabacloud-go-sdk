// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeNetworkPackagesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeNetworkPackagesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeNetworkPackagesResponse
	GetStatusCode() *int32
	SetBody(v *DescribeNetworkPackagesResponseBody) *DescribeNetworkPackagesResponse
	GetBody() *DescribeNetworkPackagesResponseBody
}

type DescribeNetworkPackagesResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeNetworkPackagesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeNetworkPackagesResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeNetworkPackagesResponse) GoString() string {
	return s.String()
}

func (s *DescribeNetworkPackagesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeNetworkPackagesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeNetworkPackagesResponse) GetBody() *DescribeNetworkPackagesResponseBody {
	return s.Body
}

func (s *DescribeNetworkPackagesResponse) SetHeaders(v map[string]*string) *DescribeNetworkPackagesResponse {
	s.Headers = v
	return s
}

func (s *DescribeNetworkPackagesResponse) SetStatusCode(v int32) *DescribeNetworkPackagesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeNetworkPackagesResponse) SetBody(v *DescribeNetworkPackagesResponseBody) *DescribeNetworkPackagesResponse {
	s.Body = v
	return s
}

func (s *DescribeNetworkPackagesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
