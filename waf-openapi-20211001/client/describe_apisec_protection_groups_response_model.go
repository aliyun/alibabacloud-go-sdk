// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeApisecProtectionGroupsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeApisecProtectionGroupsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeApisecProtectionGroupsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeApisecProtectionGroupsResponseBody) *DescribeApisecProtectionGroupsResponse
	GetBody() *DescribeApisecProtectionGroupsResponseBody
}

type DescribeApisecProtectionGroupsResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeApisecProtectionGroupsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeApisecProtectionGroupsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeApisecProtectionGroupsResponse) GoString() string {
	return s.String()
}

func (s *DescribeApisecProtectionGroupsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeApisecProtectionGroupsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeApisecProtectionGroupsResponse) GetBody() *DescribeApisecProtectionGroupsResponseBody {
	return s.Body
}

func (s *DescribeApisecProtectionGroupsResponse) SetHeaders(v map[string]*string) *DescribeApisecProtectionGroupsResponse {
	s.Headers = v
	return s
}

func (s *DescribeApisecProtectionGroupsResponse) SetStatusCode(v int32) *DescribeApisecProtectionGroupsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeApisecProtectionGroupsResponse) SetBody(v *DescribeApisecProtectionGroupsResponseBody) *DescribeApisecProtectionGroupsResponse {
	s.Body = v
	return s
}

func (s *DescribeApisecProtectionGroupsResponse) Validate() error {
	return dara.Validate(s)
}
