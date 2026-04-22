// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeIntranetUserCanAnalysisVpcsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeIntranetUserCanAnalysisVpcsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeIntranetUserCanAnalysisVpcsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeIntranetUserCanAnalysisVpcsResponseBody) *DescribeIntranetUserCanAnalysisVpcsResponse
	GetBody() *DescribeIntranetUserCanAnalysisVpcsResponseBody
}

type DescribeIntranetUserCanAnalysisVpcsResponse struct {
	Headers    map[string]*string                               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeIntranetUserCanAnalysisVpcsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeIntranetUserCanAnalysisVpcsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeIntranetUserCanAnalysisVpcsResponse) GoString() string {
	return s.String()
}

func (s *DescribeIntranetUserCanAnalysisVpcsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeIntranetUserCanAnalysisVpcsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeIntranetUserCanAnalysisVpcsResponse) GetBody() *DescribeIntranetUserCanAnalysisVpcsResponseBody {
	return s.Body
}

func (s *DescribeIntranetUserCanAnalysisVpcsResponse) SetHeaders(v map[string]*string) *DescribeIntranetUserCanAnalysisVpcsResponse {
	s.Headers = v
	return s
}

func (s *DescribeIntranetUserCanAnalysisVpcsResponse) SetStatusCode(v int32) *DescribeIntranetUserCanAnalysisVpcsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeIntranetUserCanAnalysisVpcsResponse) SetBody(v *DescribeIntranetUserCanAnalysisVpcsResponseBody) *DescribeIntranetUserCanAnalysisVpcsResponse {
	s.Body = v
	return s
}

func (s *DescribeIntranetUserCanAnalysisVpcsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
