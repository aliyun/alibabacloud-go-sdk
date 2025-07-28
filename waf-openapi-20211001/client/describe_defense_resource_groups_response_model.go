// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDefenseResourceGroupsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDefenseResourceGroupsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDefenseResourceGroupsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDefenseResourceGroupsResponseBody) *DescribeDefenseResourceGroupsResponse
	GetBody() *DescribeDefenseResourceGroupsResponseBody
}

type DescribeDefenseResourceGroupsResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDefenseResourceGroupsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDefenseResourceGroupsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDefenseResourceGroupsResponse) GoString() string {
	return s.String()
}

func (s *DescribeDefenseResourceGroupsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDefenseResourceGroupsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDefenseResourceGroupsResponse) GetBody() *DescribeDefenseResourceGroupsResponseBody {
	return s.Body
}

func (s *DescribeDefenseResourceGroupsResponse) SetHeaders(v map[string]*string) *DescribeDefenseResourceGroupsResponse {
	s.Headers = v
	return s
}

func (s *DescribeDefenseResourceGroupsResponse) SetStatusCode(v int32) *DescribeDefenseResourceGroupsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDefenseResourceGroupsResponse) SetBody(v *DescribeDefenseResourceGroupsResponseBody) *DescribeDefenseResourceGroupsResponse {
	s.Body = v
	return s
}

func (s *DescribeDefenseResourceGroupsResponse) Validate() error {
	return dara.Validate(s)
}
