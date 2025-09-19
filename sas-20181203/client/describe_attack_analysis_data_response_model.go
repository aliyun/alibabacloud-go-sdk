// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAttackAnalysisDataResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeAttackAnalysisDataResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeAttackAnalysisDataResponse
	GetStatusCode() *int32
	SetBody(v *DescribeAttackAnalysisDataResponseBody) *DescribeAttackAnalysisDataResponse
	GetBody() *DescribeAttackAnalysisDataResponseBody
}

type DescribeAttackAnalysisDataResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeAttackAnalysisDataResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeAttackAnalysisDataResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeAttackAnalysisDataResponse) GoString() string {
	return s.String()
}

func (s *DescribeAttackAnalysisDataResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeAttackAnalysisDataResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeAttackAnalysisDataResponse) GetBody() *DescribeAttackAnalysisDataResponseBody {
	return s.Body
}

func (s *DescribeAttackAnalysisDataResponse) SetHeaders(v map[string]*string) *DescribeAttackAnalysisDataResponse {
	s.Headers = v
	return s
}

func (s *DescribeAttackAnalysisDataResponse) SetStatusCode(v int32) *DescribeAttackAnalysisDataResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeAttackAnalysisDataResponse) SetBody(v *DescribeAttackAnalysisDataResponseBody) *DescribeAttackAnalysisDataResponse {
	s.Body = v
	return s
}

func (s *DescribeAttackAnalysisDataResponse) Validate() error {
	return dara.Validate(s)
}
