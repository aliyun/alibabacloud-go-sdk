// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRiskLevelsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeRiskLevelsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeRiskLevelsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeRiskLevelsResponseBody) *DescribeRiskLevelsResponse
	GetBody() *DescribeRiskLevelsResponseBody
}

type DescribeRiskLevelsResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeRiskLevelsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeRiskLevelsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeRiskLevelsResponse) GoString() string {
	return s.String()
}

func (s *DescribeRiskLevelsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeRiskLevelsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeRiskLevelsResponse) GetBody() *DescribeRiskLevelsResponseBody {
	return s.Body
}

func (s *DescribeRiskLevelsResponse) SetHeaders(v map[string]*string) *DescribeRiskLevelsResponse {
	s.Headers = v
	return s
}

func (s *DescribeRiskLevelsResponse) SetStatusCode(v int32) *DescribeRiskLevelsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeRiskLevelsResponse) SetBody(v *DescribeRiskLevelsResponseBody) *DescribeRiskLevelsResponse {
	s.Body = v
	return s
}

func (s *DescribeRiskLevelsResponse) Validate() error {
	return dara.Validate(s)
}
