// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSmartAccessGatewayClientUsersResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeSmartAccessGatewayClientUsersResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeSmartAccessGatewayClientUsersResponse
	GetStatusCode() *int32
	SetBody(v *DescribeSmartAccessGatewayClientUsersResponseBody) *DescribeSmartAccessGatewayClientUsersResponse
	GetBody() *DescribeSmartAccessGatewayClientUsersResponseBody
}

type DescribeSmartAccessGatewayClientUsersResponse struct {
	Headers    map[string]*string                                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeSmartAccessGatewayClientUsersResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeSmartAccessGatewayClientUsersResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeSmartAccessGatewayClientUsersResponse) GoString() string {
	return s.String()
}

func (s *DescribeSmartAccessGatewayClientUsersResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeSmartAccessGatewayClientUsersResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeSmartAccessGatewayClientUsersResponse) GetBody() *DescribeSmartAccessGatewayClientUsersResponseBody {
	return s.Body
}

func (s *DescribeSmartAccessGatewayClientUsersResponse) SetHeaders(v map[string]*string) *DescribeSmartAccessGatewayClientUsersResponse {
	s.Headers = v
	return s
}

func (s *DescribeSmartAccessGatewayClientUsersResponse) SetStatusCode(v int32) *DescribeSmartAccessGatewayClientUsersResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeSmartAccessGatewayClientUsersResponse) SetBody(v *DescribeSmartAccessGatewayClientUsersResponseBody) *DescribeSmartAccessGatewayClientUsersResponse {
	s.Body = v
	return s
}

func (s *DescribeSmartAccessGatewayClientUsersResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
