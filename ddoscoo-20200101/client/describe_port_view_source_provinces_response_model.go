// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePortViewSourceProvincesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribePortViewSourceProvincesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribePortViewSourceProvincesResponse
	GetStatusCode() *int32
	SetBody(v *DescribePortViewSourceProvincesResponseBody) *DescribePortViewSourceProvincesResponse
	GetBody() *DescribePortViewSourceProvincesResponseBody
}

type DescribePortViewSourceProvincesResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribePortViewSourceProvincesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribePortViewSourceProvincesResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribePortViewSourceProvincesResponse) GoString() string {
	return s.String()
}

func (s *DescribePortViewSourceProvincesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribePortViewSourceProvincesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribePortViewSourceProvincesResponse) GetBody() *DescribePortViewSourceProvincesResponseBody {
	return s.Body
}

func (s *DescribePortViewSourceProvincesResponse) SetHeaders(v map[string]*string) *DescribePortViewSourceProvincesResponse {
	s.Headers = v
	return s
}

func (s *DescribePortViewSourceProvincesResponse) SetStatusCode(v int32) *DescribePortViewSourceProvincesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribePortViewSourceProvincesResponse) SetBody(v *DescribePortViewSourceProvincesResponseBody) *DescribePortViewSourceProvincesResponse {
	s.Body = v
	return s
}

func (s *DescribePortViewSourceProvincesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
