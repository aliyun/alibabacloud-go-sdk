// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeOpenSearchWhitelistsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeOpenSearchWhitelistsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeOpenSearchWhitelistsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeOpenSearchWhitelistsResponseBody) *DescribeOpenSearchWhitelistsResponse
	GetBody() *DescribeOpenSearchWhitelistsResponseBody
}

type DescribeOpenSearchWhitelistsResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeOpenSearchWhitelistsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeOpenSearchWhitelistsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeOpenSearchWhitelistsResponse) GoString() string {
	return s.String()
}

func (s *DescribeOpenSearchWhitelistsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeOpenSearchWhitelistsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeOpenSearchWhitelistsResponse) GetBody() *DescribeOpenSearchWhitelistsResponseBody {
	return s.Body
}

func (s *DescribeOpenSearchWhitelistsResponse) SetHeaders(v map[string]*string) *DescribeOpenSearchWhitelistsResponse {
	s.Headers = v
	return s
}

func (s *DescribeOpenSearchWhitelistsResponse) SetStatusCode(v int32) *DescribeOpenSearchWhitelistsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeOpenSearchWhitelistsResponse) SetBody(v *DescribeOpenSearchWhitelistsResponseBody) *DescribeOpenSearchWhitelistsResponse {
	s.Body = v
	return s
}

func (s *DescribeOpenSearchWhitelistsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
