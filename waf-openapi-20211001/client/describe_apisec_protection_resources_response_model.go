// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeApisecProtectionResourcesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeApisecProtectionResourcesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeApisecProtectionResourcesResponse
	GetStatusCode() *int32
	SetBody(v *DescribeApisecProtectionResourcesResponseBody) *DescribeApisecProtectionResourcesResponse
	GetBody() *DescribeApisecProtectionResourcesResponseBody
}

type DescribeApisecProtectionResourcesResponse struct {
	Headers    map[string]*string                             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeApisecProtectionResourcesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeApisecProtectionResourcesResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeApisecProtectionResourcesResponse) GoString() string {
	return s.String()
}

func (s *DescribeApisecProtectionResourcesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeApisecProtectionResourcesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeApisecProtectionResourcesResponse) GetBody() *DescribeApisecProtectionResourcesResponseBody {
	return s.Body
}

func (s *DescribeApisecProtectionResourcesResponse) SetHeaders(v map[string]*string) *DescribeApisecProtectionResourcesResponse {
	s.Headers = v
	return s
}

func (s *DescribeApisecProtectionResourcesResponse) SetStatusCode(v int32) *DescribeApisecProtectionResourcesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeApisecProtectionResourcesResponse) SetBody(v *DescribeApisecProtectionResourcesResponseBody) *DescribeApisecProtectionResourcesResponse {
	s.Body = v
	return s
}

func (s *DescribeApisecProtectionResourcesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
