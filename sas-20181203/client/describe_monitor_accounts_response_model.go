// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeMonitorAccountsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeMonitorAccountsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeMonitorAccountsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeMonitorAccountsResponseBody) *DescribeMonitorAccountsResponse
	GetBody() *DescribeMonitorAccountsResponseBody
}

type DescribeMonitorAccountsResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeMonitorAccountsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeMonitorAccountsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeMonitorAccountsResponse) GoString() string {
	return s.String()
}

func (s *DescribeMonitorAccountsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeMonitorAccountsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeMonitorAccountsResponse) GetBody() *DescribeMonitorAccountsResponseBody {
	return s.Body
}

func (s *DescribeMonitorAccountsResponse) SetHeaders(v map[string]*string) *DescribeMonitorAccountsResponse {
	s.Headers = v
	return s
}

func (s *DescribeMonitorAccountsResponse) SetStatusCode(v int32) *DescribeMonitorAccountsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeMonitorAccountsResponse) SetBody(v *DescribeMonitorAccountsResponseBody) *DescribeMonitorAccountsResponse {
	s.Body = v
	return s
}

func (s *DescribeMonitorAccountsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
