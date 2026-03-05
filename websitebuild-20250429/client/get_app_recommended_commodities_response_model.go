// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAppRecommendedCommoditiesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetAppRecommendedCommoditiesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetAppRecommendedCommoditiesResponse
	GetStatusCode() *int32
	SetBody(v *GetAppRecommendedCommoditiesResponseBody) *GetAppRecommendedCommoditiesResponse
	GetBody() *GetAppRecommendedCommoditiesResponseBody
}

type GetAppRecommendedCommoditiesResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetAppRecommendedCommoditiesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetAppRecommendedCommoditiesResponse) String() string {
	return dara.Prettify(s)
}

func (s GetAppRecommendedCommoditiesResponse) GoString() string {
	return s.String()
}

func (s *GetAppRecommendedCommoditiesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetAppRecommendedCommoditiesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetAppRecommendedCommoditiesResponse) GetBody() *GetAppRecommendedCommoditiesResponseBody {
	return s.Body
}

func (s *GetAppRecommendedCommoditiesResponse) SetHeaders(v map[string]*string) *GetAppRecommendedCommoditiesResponse {
	s.Headers = v
	return s
}

func (s *GetAppRecommendedCommoditiesResponse) SetStatusCode(v int32) *GetAppRecommendedCommoditiesResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAppRecommendedCommoditiesResponse) SetBody(v *GetAppRecommendedCommoditiesResponseBody) *GetAppRecommendedCommoditiesResponse {
	s.Body = v
	return s
}

func (s *GetAppRecommendedCommoditiesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
