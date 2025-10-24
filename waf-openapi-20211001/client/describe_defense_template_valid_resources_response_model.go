// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDefenseTemplateValidResourcesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDefenseTemplateValidResourcesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDefenseTemplateValidResourcesResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDefenseTemplateValidResourcesResponseBody) *DescribeDefenseTemplateValidResourcesResponse
	GetBody() *DescribeDefenseTemplateValidResourcesResponseBody
}

type DescribeDefenseTemplateValidResourcesResponse struct {
	Headers    map[string]*string                                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDefenseTemplateValidResourcesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDefenseTemplateValidResourcesResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDefenseTemplateValidResourcesResponse) GoString() string {
	return s.String()
}

func (s *DescribeDefenseTemplateValidResourcesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDefenseTemplateValidResourcesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDefenseTemplateValidResourcesResponse) GetBody() *DescribeDefenseTemplateValidResourcesResponseBody {
	return s.Body
}

func (s *DescribeDefenseTemplateValidResourcesResponse) SetHeaders(v map[string]*string) *DescribeDefenseTemplateValidResourcesResponse {
	s.Headers = v
	return s
}

func (s *DescribeDefenseTemplateValidResourcesResponse) SetStatusCode(v int32) *DescribeDefenseTemplateValidResourcesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDefenseTemplateValidResourcesResponse) SetBody(v *DescribeDefenseTemplateValidResourcesResponseBody) *DescribeDefenseTemplateValidResourcesResponse {
	s.Body = v
	return s
}

func (s *DescribeDefenseTemplateValidResourcesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
