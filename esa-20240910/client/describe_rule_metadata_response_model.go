// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRuleMetadataResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeRuleMetadataResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeRuleMetadataResponse
	GetStatusCode() *int32
	SetBody(v *DescribeRuleMetadataResponseBody) *DescribeRuleMetadataResponse
	GetBody() *DescribeRuleMetadataResponseBody
}

type DescribeRuleMetadataResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeRuleMetadataResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeRuleMetadataResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeRuleMetadataResponse) GoString() string {
	return s.String()
}

func (s *DescribeRuleMetadataResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeRuleMetadataResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeRuleMetadataResponse) GetBody() *DescribeRuleMetadataResponseBody {
	return s.Body
}

func (s *DescribeRuleMetadataResponse) SetHeaders(v map[string]*string) *DescribeRuleMetadataResponse {
	s.Headers = v
	return s
}

func (s *DescribeRuleMetadataResponse) SetStatusCode(v int32) *DescribeRuleMetadataResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeRuleMetadataResponse) SetBody(v *DescribeRuleMetadataResponseBody) *DescribeRuleMetadataResponse {
	s.Body = v
	return s
}

func (s *DescribeRuleMetadataResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
