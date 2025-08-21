// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAttackAnalysisMaxQpsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeAttackAnalysisMaxQpsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeAttackAnalysisMaxQpsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeAttackAnalysisMaxQpsResponseBody) *DescribeAttackAnalysisMaxQpsResponse
	GetBody() *DescribeAttackAnalysisMaxQpsResponseBody
}

type DescribeAttackAnalysisMaxQpsResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeAttackAnalysisMaxQpsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeAttackAnalysisMaxQpsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeAttackAnalysisMaxQpsResponse) GoString() string {
	return s.String()
}

func (s *DescribeAttackAnalysisMaxQpsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeAttackAnalysisMaxQpsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeAttackAnalysisMaxQpsResponse) GetBody() *DescribeAttackAnalysisMaxQpsResponseBody {
	return s.Body
}

func (s *DescribeAttackAnalysisMaxQpsResponse) SetHeaders(v map[string]*string) *DescribeAttackAnalysisMaxQpsResponse {
	s.Headers = v
	return s
}

func (s *DescribeAttackAnalysisMaxQpsResponse) SetStatusCode(v int32) *DescribeAttackAnalysisMaxQpsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeAttackAnalysisMaxQpsResponse) SetBody(v *DescribeAttackAnalysisMaxQpsResponseBody) *DescribeAttackAnalysisMaxQpsResponse {
	s.Body = v
	return s
}

func (s *DescribeAttackAnalysisMaxQpsResponse) Validate() error {
	return dara.Validate(s)
}
