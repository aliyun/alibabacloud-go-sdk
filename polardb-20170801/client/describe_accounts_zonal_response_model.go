// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAccountsZonalResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeAccountsZonalResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeAccountsZonalResponse
	GetStatusCode() *int32
	SetBody(v *DescribeAccountsZonalResponseBody) *DescribeAccountsZonalResponse
	GetBody() *DescribeAccountsZonalResponseBody
}

type DescribeAccountsZonalResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeAccountsZonalResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeAccountsZonalResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeAccountsZonalResponse) GoString() string {
	return s.String()
}

func (s *DescribeAccountsZonalResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeAccountsZonalResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeAccountsZonalResponse) GetBody() *DescribeAccountsZonalResponseBody {
	return s.Body
}

func (s *DescribeAccountsZonalResponse) SetHeaders(v map[string]*string) *DescribeAccountsZonalResponse {
	s.Headers = v
	return s
}

func (s *DescribeAccountsZonalResponse) SetStatusCode(v int32) *DescribeAccountsZonalResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeAccountsZonalResponse) SetBody(v *DescribeAccountsZonalResponseBody) *DescribeAccountsZonalResponse {
	s.Body = v
	return s
}

func (s *DescribeAccountsZonalResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
