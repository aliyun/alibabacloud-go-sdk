// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeNisInspectionRecommendationResourcesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeNisInspectionRecommendationResourcesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeNisInspectionRecommendationResourcesResponse
	GetStatusCode() *int32
	SetBody(v *DescribeNisInspectionRecommendationResourcesResponseBody) *DescribeNisInspectionRecommendationResourcesResponse
	GetBody() *DescribeNisInspectionRecommendationResourcesResponseBody
}

type DescribeNisInspectionRecommendationResourcesResponse struct {
	Headers    map[string]*string                                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeNisInspectionRecommendationResourcesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeNisInspectionRecommendationResourcesResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeNisInspectionRecommendationResourcesResponse) GoString() string {
	return s.String()
}

func (s *DescribeNisInspectionRecommendationResourcesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeNisInspectionRecommendationResourcesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeNisInspectionRecommendationResourcesResponse) GetBody() *DescribeNisInspectionRecommendationResourcesResponseBody {
	return s.Body
}

func (s *DescribeNisInspectionRecommendationResourcesResponse) SetHeaders(v map[string]*string) *DescribeNisInspectionRecommendationResourcesResponse {
	s.Headers = v
	return s
}

func (s *DescribeNisInspectionRecommendationResourcesResponse) SetStatusCode(v int32) *DescribeNisInspectionRecommendationResourcesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeNisInspectionRecommendationResourcesResponse) SetBody(v *DescribeNisInspectionRecommendationResourcesResponseBody) *DescribeNisInspectionRecommendationResourcesResponse {
	s.Body = v
	return s
}

func (s *DescribeNisInspectionRecommendationResourcesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
