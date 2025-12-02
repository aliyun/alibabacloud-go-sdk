// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeMvRecommendTasksResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeMvRecommendTasksResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeMvRecommendTasksResponse
	GetStatusCode() *int32
	SetBody(v *DescribeMvRecommendTasksResponseBody) *DescribeMvRecommendTasksResponse
	GetBody() *DescribeMvRecommendTasksResponseBody
}

type DescribeMvRecommendTasksResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeMvRecommendTasksResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeMvRecommendTasksResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeMvRecommendTasksResponse) GoString() string {
	return s.String()
}

func (s *DescribeMvRecommendTasksResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeMvRecommendTasksResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeMvRecommendTasksResponse) GetBody() *DescribeMvRecommendTasksResponseBody {
	return s.Body
}

func (s *DescribeMvRecommendTasksResponse) SetHeaders(v map[string]*string) *DescribeMvRecommendTasksResponse {
	s.Headers = v
	return s
}

func (s *DescribeMvRecommendTasksResponse) SetStatusCode(v int32) *DescribeMvRecommendTasksResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeMvRecommendTasksResponse) SetBody(v *DescribeMvRecommendTasksResponseBody) *DescribeMvRecommendTasksResponse {
	s.Body = v
	return s
}

func (s *DescribeMvRecommendTasksResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
