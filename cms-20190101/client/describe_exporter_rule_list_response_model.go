// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeExporterRuleListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeExporterRuleListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeExporterRuleListResponse
	GetStatusCode() *int32
	SetBody(v *DescribeExporterRuleListResponseBody) *DescribeExporterRuleListResponse
	GetBody() *DescribeExporterRuleListResponseBody
}

type DescribeExporterRuleListResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeExporterRuleListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeExporterRuleListResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeExporterRuleListResponse) GoString() string {
	return s.String()
}

func (s *DescribeExporterRuleListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeExporterRuleListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeExporterRuleListResponse) GetBody() *DescribeExporterRuleListResponseBody {
	return s.Body
}

func (s *DescribeExporterRuleListResponse) SetHeaders(v map[string]*string) *DescribeExporterRuleListResponse {
	s.Headers = v
	return s
}

func (s *DescribeExporterRuleListResponse) SetStatusCode(v int32) *DescribeExporterRuleListResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeExporterRuleListResponse) SetBody(v *DescribeExporterRuleListResponseBody) *DescribeExporterRuleListResponse {
	s.Body = v
	return s
}

func (s *DescribeExporterRuleListResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
