// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUploadRecommendationDataResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UploadRecommendationDataResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UploadRecommendationDataResponse
	GetStatusCode() *int32
	SetBody(v *UploadRecommendationDataResponseBody) *UploadRecommendationDataResponse
	GetBody() *UploadRecommendationDataResponseBody
}

type UploadRecommendationDataResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UploadRecommendationDataResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UploadRecommendationDataResponse) String() string {
	return dara.Prettify(s)
}

func (s UploadRecommendationDataResponse) GoString() string {
	return s.String()
}

func (s *UploadRecommendationDataResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UploadRecommendationDataResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UploadRecommendationDataResponse) GetBody() *UploadRecommendationDataResponseBody {
	return s.Body
}

func (s *UploadRecommendationDataResponse) SetHeaders(v map[string]*string) *UploadRecommendationDataResponse {
	s.Headers = v
	return s
}

func (s *UploadRecommendationDataResponse) SetStatusCode(v int32) *UploadRecommendationDataResponse {
	s.StatusCode = &v
	return s
}

func (s *UploadRecommendationDataResponse) SetBody(v *UploadRecommendationDataResponseBody) *UploadRecommendationDataResponse {
	s.Body = v
	return s
}

func (s *UploadRecommendationDataResponse) Validate() error {
	return dara.Validate(s)
}
