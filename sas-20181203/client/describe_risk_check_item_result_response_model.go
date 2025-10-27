// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRiskCheckItemResultResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeRiskCheckItemResultResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeRiskCheckItemResultResponse
	GetStatusCode() *int32
	SetBody(v *DescribeRiskCheckItemResultResponseBody) *DescribeRiskCheckItemResultResponse
	GetBody() *DescribeRiskCheckItemResultResponseBody
}

type DescribeRiskCheckItemResultResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeRiskCheckItemResultResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeRiskCheckItemResultResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeRiskCheckItemResultResponse) GoString() string {
	return s.String()
}

func (s *DescribeRiskCheckItemResultResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeRiskCheckItemResultResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeRiskCheckItemResultResponse) GetBody() *DescribeRiskCheckItemResultResponseBody {
	return s.Body
}

func (s *DescribeRiskCheckItemResultResponse) SetHeaders(v map[string]*string) *DescribeRiskCheckItemResultResponse {
	s.Headers = v
	return s
}

func (s *DescribeRiskCheckItemResultResponse) SetStatusCode(v int32) *DescribeRiskCheckItemResultResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeRiskCheckItemResultResponse) SetBody(v *DescribeRiskCheckItemResultResponseBody) *DescribeRiskCheckItemResultResponse {
	s.Body = v
	return s
}

func (s *DescribeRiskCheckItemResultResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
