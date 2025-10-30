// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAnalysisColumnListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeAnalysisColumnListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeAnalysisColumnListResponse
	GetStatusCode() *int32
	SetBody(v *DescribeAnalysisColumnListResponseBody) *DescribeAnalysisColumnListResponse
	GetBody() *DescribeAnalysisColumnListResponseBody
}

type DescribeAnalysisColumnListResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeAnalysisColumnListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeAnalysisColumnListResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeAnalysisColumnListResponse) GoString() string {
	return s.String()
}

func (s *DescribeAnalysisColumnListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeAnalysisColumnListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeAnalysisColumnListResponse) GetBody() *DescribeAnalysisColumnListResponseBody {
	return s.Body
}

func (s *DescribeAnalysisColumnListResponse) SetHeaders(v map[string]*string) *DescribeAnalysisColumnListResponse {
	s.Headers = v
	return s
}

func (s *DescribeAnalysisColumnListResponse) SetStatusCode(v int32) *DescribeAnalysisColumnListResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeAnalysisColumnListResponse) SetBody(v *DescribeAnalysisColumnListResponseBody) *DescribeAnalysisColumnListResponse {
	s.Body = v
	return s
}

func (s *DescribeAnalysisColumnListResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
