// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAnalysisConditionFavoriteListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeAnalysisConditionFavoriteListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeAnalysisConditionFavoriteListResponse
	GetStatusCode() *int32
	SetBody(v *DescribeAnalysisConditionFavoriteListResponseBody) *DescribeAnalysisConditionFavoriteListResponse
	GetBody() *DescribeAnalysisConditionFavoriteListResponseBody
}

type DescribeAnalysisConditionFavoriteListResponse struct {
	Headers    map[string]*string                                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeAnalysisConditionFavoriteListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeAnalysisConditionFavoriteListResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeAnalysisConditionFavoriteListResponse) GoString() string {
	return s.String()
}

func (s *DescribeAnalysisConditionFavoriteListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeAnalysisConditionFavoriteListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeAnalysisConditionFavoriteListResponse) GetBody() *DescribeAnalysisConditionFavoriteListResponseBody {
	return s.Body
}

func (s *DescribeAnalysisConditionFavoriteListResponse) SetHeaders(v map[string]*string) *DescribeAnalysisConditionFavoriteListResponse {
	s.Headers = v
	return s
}

func (s *DescribeAnalysisConditionFavoriteListResponse) SetStatusCode(v int32) *DescribeAnalysisConditionFavoriteListResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeAnalysisConditionFavoriteListResponse) SetBody(v *DescribeAnalysisConditionFavoriteListResponseBody) *DescribeAnalysisConditionFavoriteListResponse {
	s.Body = v
	return s
}

func (s *DescribeAnalysisConditionFavoriteListResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
