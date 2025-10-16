// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAclAppsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeAclAppsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeAclAppsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeAclAppsResponseBody) *DescribeAclAppsResponse
	GetBody() *DescribeAclAppsResponseBody
}

type DescribeAclAppsResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeAclAppsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeAclAppsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeAclAppsResponse) GoString() string {
	return s.String()
}

func (s *DescribeAclAppsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeAclAppsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeAclAppsResponse) GetBody() *DescribeAclAppsResponseBody {
	return s.Body
}

func (s *DescribeAclAppsResponse) SetHeaders(v map[string]*string) *DescribeAclAppsResponse {
	s.Headers = v
	return s
}

func (s *DescribeAclAppsResponse) SetStatusCode(v int32) *DescribeAclAppsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeAclAppsResponse) SetBody(v *DescribeAclAppsResponseBody) *DescribeAclAppsResponse {
	s.Body = v
	return s
}

func (s *DescribeAclAppsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
