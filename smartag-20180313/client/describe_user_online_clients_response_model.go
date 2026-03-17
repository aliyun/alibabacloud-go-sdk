// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeUserOnlineClientsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeUserOnlineClientsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeUserOnlineClientsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeUserOnlineClientsResponseBody) *DescribeUserOnlineClientsResponse
	GetBody() *DescribeUserOnlineClientsResponseBody
}

type DescribeUserOnlineClientsResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeUserOnlineClientsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeUserOnlineClientsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeUserOnlineClientsResponse) GoString() string {
	return s.String()
}

func (s *DescribeUserOnlineClientsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeUserOnlineClientsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeUserOnlineClientsResponse) GetBody() *DescribeUserOnlineClientsResponseBody {
	return s.Body
}

func (s *DescribeUserOnlineClientsResponse) SetHeaders(v map[string]*string) *DescribeUserOnlineClientsResponse {
	s.Headers = v
	return s
}

func (s *DescribeUserOnlineClientsResponse) SetStatusCode(v int32) *DescribeUserOnlineClientsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeUserOnlineClientsResponse) SetBody(v *DescribeUserOnlineClientsResponseBody) *DescribeUserOnlineClientsResponse {
	s.Body = v
	return s
}

func (s *DescribeUserOnlineClientsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
