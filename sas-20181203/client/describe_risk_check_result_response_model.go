// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRiskCheckResultResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeRiskCheckResultResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeRiskCheckResultResponse
	GetStatusCode() *int32
	SetBody(v *DescribeRiskCheckResultResponseBody) *DescribeRiskCheckResultResponse
	GetBody() *DescribeRiskCheckResultResponseBody
}

type DescribeRiskCheckResultResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeRiskCheckResultResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeRiskCheckResultResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeRiskCheckResultResponse) GoString() string {
	return s.String()
}

func (s *DescribeRiskCheckResultResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeRiskCheckResultResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeRiskCheckResultResponse) GetBody() *DescribeRiskCheckResultResponseBody {
	return s.Body
}

func (s *DescribeRiskCheckResultResponse) SetHeaders(v map[string]*string) *DescribeRiskCheckResultResponse {
	s.Headers = v
	return s
}

func (s *DescribeRiskCheckResultResponse) SetStatusCode(v int32) *DescribeRiskCheckResultResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeRiskCheckResultResponse) SetBody(v *DescribeRiskCheckResultResponseBody) *DescribeRiskCheckResultResponse {
	s.Body = v
	return s
}

func (s *DescribeRiskCheckResultResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
