// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeApisecApiResourcesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeApisecApiResourcesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeApisecApiResourcesResponse
	GetStatusCode() *int32
	SetBody(v *DescribeApisecApiResourcesResponseBody) *DescribeApisecApiResourcesResponse
	GetBody() *DescribeApisecApiResourcesResponseBody
}

type DescribeApisecApiResourcesResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeApisecApiResourcesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeApisecApiResourcesResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeApisecApiResourcesResponse) GoString() string {
	return s.String()
}

func (s *DescribeApisecApiResourcesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeApisecApiResourcesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeApisecApiResourcesResponse) GetBody() *DescribeApisecApiResourcesResponseBody {
	return s.Body
}

func (s *DescribeApisecApiResourcesResponse) SetHeaders(v map[string]*string) *DescribeApisecApiResourcesResponse {
	s.Headers = v
	return s
}

func (s *DescribeApisecApiResourcesResponse) SetStatusCode(v int32) *DescribeApisecApiResourcesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeApisecApiResourcesResponse) SetBody(v *DescribeApisecApiResourcesResponseBody) *DescribeApisecApiResourcesResponse {
	s.Body = v
	return s
}

func (s *DescribeApisecApiResourcesResponse) Validate() error {
	return dara.Validate(s)
}
