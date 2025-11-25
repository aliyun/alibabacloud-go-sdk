// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeInstanceRdAccountsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeInstanceRdAccountsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeInstanceRdAccountsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeInstanceRdAccountsResponseBody) *DescribeInstanceRdAccountsResponse
	GetBody() *DescribeInstanceRdAccountsResponseBody
}

type DescribeInstanceRdAccountsResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeInstanceRdAccountsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeInstanceRdAccountsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstanceRdAccountsResponse) GoString() string {
	return s.String()
}

func (s *DescribeInstanceRdAccountsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeInstanceRdAccountsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeInstanceRdAccountsResponse) GetBody() *DescribeInstanceRdAccountsResponseBody {
	return s.Body
}

func (s *DescribeInstanceRdAccountsResponse) SetHeaders(v map[string]*string) *DescribeInstanceRdAccountsResponse {
	s.Headers = v
	return s
}

func (s *DescribeInstanceRdAccountsResponse) SetStatusCode(v int32) *DescribeInstanceRdAccountsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeInstanceRdAccountsResponse) SetBody(v *DescribeInstanceRdAccountsResponseBody) *DescribeInstanceRdAccountsResponse {
	s.Body = v
	return s
}

func (s *DescribeInstanceRdAccountsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
