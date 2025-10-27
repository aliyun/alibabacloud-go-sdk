// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRdsAnalysisResourceQuotasResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeRdsAnalysisResourceQuotasResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeRdsAnalysisResourceQuotasResponse
	GetStatusCode() *int32
	SetBody(v *DescribeRdsAnalysisResourceQuotasResponseBody) *DescribeRdsAnalysisResourceQuotasResponse
	GetBody() *DescribeRdsAnalysisResourceQuotasResponseBody
}

type DescribeRdsAnalysisResourceQuotasResponse struct {
	Headers    map[string]*string                             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeRdsAnalysisResourceQuotasResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeRdsAnalysisResourceQuotasResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeRdsAnalysisResourceQuotasResponse) GoString() string {
	return s.String()
}

func (s *DescribeRdsAnalysisResourceQuotasResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeRdsAnalysisResourceQuotasResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeRdsAnalysisResourceQuotasResponse) GetBody() *DescribeRdsAnalysisResourceQuotasResponseBody {
	return s.Body
}

func (s *DescribeRdsAnalysisResourceQuotasResponse) SetHeaders(v map[string]*string) *DescribeRdsAnalysisResourceQuotasResponse {
	s.Headers = v
	return s
}

func (s *DescribeRdsAnalysisResourceQuotasResponse) SetStatusCode(v int32) *DescribeRdsAnalysisResourceQuotasResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeRdsAnalysisResourceQuotasResponse) SetBody(v *DescribeRdsAnalysisResourceQuotasResponseBody) *DescribeRdsAnalysisResourceQuotasResponse {
	s.Body = v
	return s
}

func (s *DescribeRdsAnalysisResourceQuotasResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
