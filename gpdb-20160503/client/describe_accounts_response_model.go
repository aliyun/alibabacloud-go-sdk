// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAccountsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeAccountsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeAccountsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeAccountsResponseBody) *DescribeAccountsResponse
	GetBody() *DescribeAccountsResponseBody
}

type DescribeAccountsResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeAccountsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeAccountsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeAccountsResponse) GoString() string {
	return s.String()
}

func (s *DescribeAccountsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeAccountsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeAccountsResponse) GetBody() *DescribeAccountsResponseBody {
	return s.Body
}

func (s *DescribeAccountsResponse) SetHeaders(v map[string]*string) *DescribeAccountsResponse {
	s.Headers = v
	return s
}

func (s *DescribeAccountsResponse) SetStatusCode(v int32) *DescribeAccountsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeAccountsResponse) SetBody(v *DescribeAccountsResponseBody) *DescribeAccountsResponse {
	s.Body = v
	return s
}

func (s *DescribeAccountsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
