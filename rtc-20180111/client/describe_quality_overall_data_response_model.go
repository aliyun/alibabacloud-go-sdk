// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeQualityOverallDataResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeQualityOverallDataResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeQualityOverallDataResponse
	GetStatusCode() *int32
	SetBody(v *DescribeQualityOverallDataResponseBody) *DescribeQualityOverallDataResponse
	GetBody() *DescribeQualityOverallDataResponseBody
}

type DescribeQualityOverallDataResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeQualityOverallDataResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeQualityOverallDataResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeQualityOverallDataResponse) GoString() string {
	return s.String()
}

func (s *DescribeQualityOverallDataResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeQualityOverallDataResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeQualityOverallDataResponse) GetBody() *DescribeQualityOverallDataResponseBody {
	return s.Body
}

func (s *DescribeQualityOverallDataResponse) SetHeaders(v map[string]*string) *DescribeQualityOverallDataResponse {
	s.Headers = v
	return s
}

func (s *DescribeQualityOverallDataResponse) SetStatusCode(v int32) *DescribeQualityOverallDataResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeQualityOverallDataResponse) SetBody(v *DescribeQualityOverallDataResponseBody) *DescribeQualityOverallDataResponse {
	s.Body = v
	return s
}

func (s *DescribeQualityOverallDataResponse) Validate() error {
	return dara.Validate(s)
}
