// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRCSecurityGroupPermissionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeRCSecurityGroupPermissionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeRCSecurityGroupPermissionResponse
	GetStatusCode() *int32
	SetBody(v *DescribeRCSecurityGroupPermissionResponseBody) *DescribeRCSecurityGroupPermissionResponse
	GetBody() *DescribeRCSecurityGroupPermissionResponseBody
}

type DescribeRCSecurityGroupPermissionResponse struct {
	Headers    map[string]*string                             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeRCSecurityGroupPermissionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeRCSecurityGroupPermissionResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeRCSecurityGroupPermissionResponse) GoString() string {
	return s.String()
}

func (s *DescribeRCSecurityGroupPermissionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeRCSecurityGroupPermissionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeRCSecurityGroupPermissionResponse) GetBody() *DescribeRCSecurityGroupPermissionResponseBody {
	return s.Body
}

func (s *DescribeRCSecurityGroupPermissionResponse) SetHeaders(v map[string]*string) *DescribeRCSecurityGroupPermissionResponse {
	s.Headers = v
	return s
}

func (s *DescribeRCSecurityGroupPermissionResponse) SetStatusCode(v int32) *DescribeRCSecurityGroupPermissionResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeRCSecurityGroupPermissionResponse) SetBody(v *DescribeRCSecurityGroupPermissionResponseBody) *DescribeRCSecurityGroupPermissionResponse {
	s.Body = v
	return s
}

func (s *DescribeRCSecurityGroupPermissionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
