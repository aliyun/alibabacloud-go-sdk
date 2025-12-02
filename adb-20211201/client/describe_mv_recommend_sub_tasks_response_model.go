// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeMvRecommendSubTasksResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeMvRecommendSubTasksResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeMvRecommendSubTasksResponse
	GetStatusCode() *int32
	SetBody(v *DescribeMvRecommendSubTasksResponseBody) *DescribeMvRecommendSubTasksResponse
	GetBody() *DescribeMvRecommendSubTasksResponseBody
}

type DescribeMvRecommendSubTasksResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeMvRecommendSubTasksResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeMvRecommendSubTasksResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeMvRecommendSubTasksResponse) GoString() string {
	return s.String()
}

func (s *DescribeMvRecommendSubTasksResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeMvRecommendSubTasksResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeMvRecommendSubTasksResponse) GetBody() *DescribeMvRecommendSubTasksResponseBody {
	return s.Body
}

func (s *DescribeMvRecommendSubTasksResponse) SetHeaders(v map[string]*string) *DescribeMvRecommendSubTasksResponse {
	s.Headers = v
	return s
}

func (s *DescribeMvRecommendSubTasksResponse) SetStatusCode(v int32) *DescribeMvRecommendSubTasksResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeMvRecommendSubTasksResponse) SetBody(v *DescribeMvRecommendSubTasksResponseBody) *DescribeMvRecommendSubTasksResponse {
	s.Body = v
	return s
}

func (s *DescribeMvRecommendSubTasksResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
