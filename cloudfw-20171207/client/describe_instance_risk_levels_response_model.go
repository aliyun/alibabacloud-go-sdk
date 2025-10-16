// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeInstanceRiskLevelsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeInstanceRiskLevelsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeInstanceRiskLevelsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeInstanceRiskLevelsResponseBody) *DescribeInstanceRiskLevelsResponse
	GetBody() *DescribeInstanceRiskLevelsResponseBody
}

type DescribeInstanceRiskLevelsResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeInstanceRiskLevelsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeInstanceRiskLevelsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstanceRiskLevelsResponse) GoString() string {
	return s.String()
}

func (s *DescribeInstanceRiskLevelsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeInstanceRiskLevelsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeInstanceRiskLevelsResponse) GetBody() *DescribeInstanceRiskLevelsResponseBody {
	return s.Body
}

func (s *DescribeInstanceRiskLevelsResponse) SetHeaders(v map[string]*string) *DescribeInstanceRiskLevelsResponse {
	s.Headers = v
	return s
}

func (s *DescribeInstanceRiskLevelsResponse) SetStatusCode(v int32) *DescribeInstanceRiskLevelsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeInstanceRiskLevelsResponse) SetBody(v *DescribeInstanceRiskLevelsResponseBody) *DescribeInstanceRiskLevelsResponse {
	s.Body = v
	return s
}

func (s *DescribeInstanceRiskLevelsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
