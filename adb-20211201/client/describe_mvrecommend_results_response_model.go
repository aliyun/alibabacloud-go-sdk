// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeMVRecommendResultsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeMVRecommendResultsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeMVRecommendResultsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeMVRecommendResultsResponseBody) *DescribeMVRecommendResultsResponse
	GetBody() *DescribeMVRecommendResultsResponseBody
}

type DescribeMVRecommendResultsResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeMVRecommendResultsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeMVRecommendResultsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeMVRecommendResultsResponse) GoString() string {
	return s.String()
}

func (s *DescribeMVRecommendResultsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeMVRecommendResultsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeMVRecommendResultsResponse) GetBody() *DescribeMVRecommendResultsResponseBody {
	return s.Body
}

func (s *DescribeMVRecommendResultsResponse) SetHeaders(v map[string]*string) *DescribeMVRecommendResultsResponse {
	s.Headers = v
	return s
}

func (s *DescribeMVRecommendResultsResponse) SetStatusCode(v int32) *DescribeMVRecommendResultsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeMVRecommendResultsResponse) SetBody(v *DescribeMVRecommendResultsResponseBody) *DescribeMVRecommendResultsResponse {
	s.Body = v
	return s
}

func (s *DescribeMVRecommendResultsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
