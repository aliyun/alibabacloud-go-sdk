// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeUserPermissionsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeUserPermissionsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeUserPermissionsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeUserPermissionsResponseBody) *DescribeUserPermissionsResponse
	GetBody() *DescribeUserPermissionsResponseBody
}

type DescribeUserPermissionsResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeUserPermissionsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeUserPermissionsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeUserPermissionsResponse) GoString() string {
	return s.String()
}

func (s *DescribeUserPermissionsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeUserPermissionsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeUserPermissionsResponse) GetBody() *DescribeUserPermissionsResponseBody {
	return s.Body
}

func (s *DescribeUserPermissionsResponse) SetHeaders(v map[string]*string) *DescribeUserPermissionsResponse {
	s.Headers = v
	return s
}

func (s *DescribeUserPermissionsResponse) SetStatusCode(v int32) *DescribeUserPermissionsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeUserPermissionsResponse) SetBody(v *DescribeUserPermissionsResponseBody) *DescribeUserPermissionsResponse {
	s.Body = v
	return s
}

func (s *DescribeUserPermissionsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
