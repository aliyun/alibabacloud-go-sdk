// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDefenseTemplateValidGroupsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDefenseTemplateValidGroupsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDefenseTemplateValidGroupsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDefenseTemplateValidGroupsResponseBody) *DescribeDefenseTemplateValidGroupsResponse
	GetBody() *DescribeDefenseTemplateValidGroupsResponseBody
}

type DescribeDefenseTemplateValidGroupsResponse struct {
	Headers    map[string]*string                              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDefenseTemplateValidGroupsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDefenseTemplateValidGroupsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDefenseTemplateValidGroupsResponse) GoString() string {
	return s.String()
}

func (s *DescribeDefenseTemplateValidGroupsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDefenseTemplateValidGroupsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDefenseTemplateValidGroupsResponse) GetBody() *DescribeDefenseTemplateValidGroupsResponseBody {
	return s.Body
}

func (s *DescribeDefenseTemplateValidGroupsResponse) SetHeaders(v map[string]*string) *DescribeDefenseTemplateValidGroupsResponse {
	s.Headers = v
	return s
}

func (s *DescribeDefenseTemplateValidGroupsResponse) SetStatusCode(v int32) *DescribeDefenseTemplateValidGroupsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDefenseTemplateValidGroupsResponse) SetBody(v *DescribeDefenseTemplateValidGroupsResponseBody) *DescribeDefenseTemplateValidGroupsResponse {
	s.Body = v
	return s
}

func (s *DescribeDefenseTemplateValidGroupsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
