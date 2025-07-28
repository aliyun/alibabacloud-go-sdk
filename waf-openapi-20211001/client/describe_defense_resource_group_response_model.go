// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDefenseResourceGroupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDefenseResourceGroupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDefenseResourceGroupResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDefenseResourceGroupResponseBody) *DescribeDefenseResourceGroupResponse
	GetBody() *DescribeDefenseResourceGroupResponseBody
}

type DescribeDefenseResourceGroupResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDefenseResourceGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDefenseResourceGroupResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDefenseResourceGroupResponse) GoString() string {
	return s.String()
}

func (s *DescribeDefenseResourceGroupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDefenseResourceGroupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDefenseResourceGroupResponse) GetBody() *DescribeDefenseResourceGroupResponseBody {
	return s.Body
}

func (s *DescribeDefenseResourceGroupResponse) SetHeaders(v map[string]*string) *DescribeDefenseResourceGroupResponse {
	s.Headers = v
	return s
}

func (s *DescribeDefenseResourceGroupResponse) SetStatusCode(v int32) *DescribeDefenseResourceGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDefenseResourceGroupResponse) SetBody(v *DescribeDefenseResourceGroupResponseBody) *DescribeDefenseResourceGroupResponse {
	s.Body = v
	return s
}

func (s *DescribeDefenseResourceGroupResponse) Validate() error {
	return dara.Validate(s)
}
