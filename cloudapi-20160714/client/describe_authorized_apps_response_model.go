// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAuthorizedAppsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeAuthorizedAppsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeAuthorizedAppsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeAuthorizedAppsResponseBody) *DescribeAuthorizedAppsResponse
	GetBody() *DescribeAuthorizedAppsResponseBody
}

type DescribeAuthorizedAppsResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeAuthorizedAppsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeAuthorizedAppsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeAuthorizedAppsResponse) GoString() string {
	return s.String()
}

func (s *DescribeAuthorizedAppsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeAuthorizedAppsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeAuthorizedAppsResponse) GetBody() *DescribeAuthorizedAppsResponseBody {
	return s.Body
}

func (s *DescribeAuthorizedAppsResponse) SetHeaders(v map[string]*string) *DescribeAuthorizedAppsResponse {
	s.Headers = v
	return s
}

func (s *DescribeAuthorizedAppsResponse) SetStatusCode(v int32) *DescribeAuthorizedAppsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeAuthorizedAppsResponse) SetBody(v *DescribeAuthorizedAppsResponseBody) *DescribeAuthorizedAppsResponse {
	s.Body = v
	return s
}

func (s *DescribeAuthorizedAppsResponse) Validate() error {
	return dara.Validate(s)
}
