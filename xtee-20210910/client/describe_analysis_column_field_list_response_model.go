// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAnalysisColumnFieldListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeAnalysisColumnFieldListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeAnalysisColumnFieldListResponse
	GetStatusCode() *int32
	SetBody(v *DescribeAnalysisColumnFieldListResponseBody) *DescribeAnalysisColumnFieldListResponse
	GetBody() *DescribeAnalysisColumnFieldListResponseBody
}

type DescribeAnalysisColumnFieldListResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeAnalysisColumnFieldListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeAnalysisColumnFieldListResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeAnalysisColumnFieldListResponse) GoString() string {
	return s.String()
}

func (s *DescribeAnalysisColumnFieldListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeAnalysisColumnFieldListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeAnalysisColumnFieldListResponse) GetBody() *DescribeAnalysisColumnFieldListResponseBody {
	return s.Body
}

func (s *DescribeAnalysisColumnFieldListResponse) SetHeaders(v map[string]*string) *DescribeAnalysisColumnFieldListResponse {
	s.Headers = v
	return s
}

func (s *DescribeAnalysisColumnFieldListResponse) SetStatusCode(v int32) *DescribeAnalysisColumnFieldListResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeAnalysisColumnFieldListResponse) SetBody(v *DescribeAnalysisColumnFieldListResponseBody) *DescribeAnalysisColumnFieldListResponse {
	s.Body = v
	return s
}

func (s *DescribeAnalysisColumnFieldListResponse) Validate() error {
	return dara.Validate(s)
}
