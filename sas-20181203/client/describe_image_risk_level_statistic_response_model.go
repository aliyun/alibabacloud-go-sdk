// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeImageRiskLevelStatisticResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeImageRiskLevelStatisticResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeImageRiskLevelStatisticResponse
	GetStatusCode() *int32
	SetBody(v *DescribeImageRiskLevelStatisticResponseBody) *DescribeImageRiskLevelStatisticResponse
	GetBody() *DescribeImageRiskLevelStatisticResponseBody
}

type DescribeImageRiskLevelStatisticResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeImageRiskLevelStatisticResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeImageRiskLevelStatisticResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeImageRiskLevelStatisticResponse) GoString() string {
	return s.String()
}

func (s *DescribeImageRiskLevelStatisticResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeImageRiskLevelStatisticResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeImageRiskLevelStatisticResponse) GetBody() *DescribeImageRiskLevelStatisticResponseBody {
	return s.Body
}

func (s *DescribeImageRiskLevelStatisticResponse) SetHeaders(v map[string]*string) *DescribeImageRiskLevelStatisticResponse {
	s.Headers = v
	return s
}

func (s *DescribeImageRiskLevelStatisticResponse) SetStatusCode(v int32) *DescribeImageRiskLevelStatisticResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeImageRiskLevelStatisticResponse) SetBody(v *DescribeImageRiskLevelStatisticResponseBody) *DescribeImageRiskLevelStatisticResponse {
	s.Body = v
	return s
}

func (s *DescribeImageRiskLevelStatisticResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
