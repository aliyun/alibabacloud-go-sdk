// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAllAccountsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeAllAccountsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeAllAccountsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeAllAccountsResponseBody) *DescribeAllAccountsResponse
	GetBody() *DescribeAllAccountsResponseBody
}

type DescribeAllAccountsResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeAllAccountsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeAllAccountsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeAllAccountsResponse) GoString() string {
	return s.String()
}

func (s *DescribeAllAccountsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeAllAccountsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeAllAccountsResponse) GetBody() *DescribeAllAccountsResponseBody {
	return s.Body
}

func (s *DescribeAllAccountsResponse) SetHeaders(v map[string]*string) *DescribeAllAccountsResponse {
	s.Headers = v
	return s
}

func (s *DescribeAllAccountsResponse) SetStatusCode(v int32) *DescribeAllAccountsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeAllAccountsResponse) SetBody(v *DescribeAllAccountsResponseBody) *DescribeAllAccountsResponse {
	s.Body = v
	return s
}

func (s *DescribeAllAccountsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
