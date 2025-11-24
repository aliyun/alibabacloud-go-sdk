// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeUsersWithPermissionsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeUsersWithPermissionsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeUsersWithPermissionsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeUsersWithPermissionsResponseBody) *DescribeUsersWithPermissionsResponse
	GetBody() *DescribeUsersWithPermissionsResponseBody
}

type DescribeUsersWithPermissionsResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeUsersWithPermissionsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeUsersWithPermissionsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeUsersWithPermissionsResponse) GoString() string {
	return s.String()
}

func (s *DescribeUsersWithPermissionsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeUsersWithPermissionsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeUsersWithPermissionsResponse) GetBody() *DescribeUsersWithPermissionsResponseBody {
	return s.Body
}

func (s *DescribeUsersWithPermissionsResponse) SetHeaders(v map[string]*string) *DescribeUsersWithPermissionsResponse {
	s.Headers = v
	return s
}

func (s *DescribeUsersWithPermissionsResponse) SetStatusCode(v int32) *DescribeUsersWithPermissionsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeUsersWithPermissionsResponse) SetBody(v *DescribeUsersWithPermissionsResponseBody) *DescribeUsersWithPermissionsResponse {
	s.Body = v
	return s
}

func (s *DescribeUsersWithPermissionsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
