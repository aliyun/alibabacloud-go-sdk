// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeTemplateResourcesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeTemplateResourcesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeTemplateResourcesResponse
	GetStatusCode() *int32
	SetBody(v *DescribeTemplateResourcesResponseBody) *DescribeTemplateResourcesResponse
	GetBody() *DescribeTemplateResourcesResponseBody
}

type DescribeTemplateResourcesResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeTemplateResourcesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeTemplateResourcesResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeTemplateResourcesResponse) GoString() string {
	return s.String()
}

func (s *DescribeTemplateResourcesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeTemplateResourcesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeTemplateResourcesResponse) GetBody() *DescribeTemplateResourcesResponseBody {
	return s.Body
}

func (s *DescribeTemplateResourcesResponse) SetHeaders(v map[string]*string) *DescribeTemplateResourcesResponse {
	s.Headers = v
	return s
}

func (s *DescribeTemplateResourcesResponse) SetStatusCode(v int32) *DescribeTemplateResourcesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeTemplateResourcesResponse) SetBody(v *DescribeTemplateResourcesResponseBody) *DescribeTemplateResourcesResponse {
	s.Body = v
	return s
}

func (s *DescribeTemplateResourcesResponse) Validate() error {
	return dara.Validate(s)
}
