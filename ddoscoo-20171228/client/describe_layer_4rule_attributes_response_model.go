// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLayer4RuleAttributesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeLayer4RuleAttributesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeLayer4RuleAttributesResponse
	GetStatusCode() *int32
	SetBody(v *DescribeLayer4RuleAttributesResponseBody) *DescribeLayer4RuleAttributesResponse
	GetBody() *DescribeLayer4RuleAttributesResponseBody
}

type DescribeLayer4RuleAttributesResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeLayer4RuleAttributesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeLayer4RuleAttributesResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeLayer4RuleAttributesResponse) GoString() string {
	return s.String()
}

func (s *DescribeLayer4RuleAttributesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeLayer4RuleAttributesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeLayer4RuleAttributesResponse) GetBody() *DescribeLayer4RuleAttributesResponseBody {
	return s.Body
}

func (s *DescribeLayer4RuleAttributesResponse) SetHeaders(v map[string]*string) *DescribeLayer4RuleAttributesResponse {
	s.Headers = v
	return s
}

func (s *DescribeLayer4RuleAttributesResponse) SetStatusCode(v int32) *DescribeLayer4RuleAttributesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeLayer4RuleAttributesResponse) SetBody(v *DescribeLayer4RuleAttributesResponseBody) *DescribeLayer4RuleAttributesResponse {
	s.Body = v
	return s
}

func (s *DescribeLayer4RuleAttributesResponse) Validate() error {
	return dara.Validate(s)
}
