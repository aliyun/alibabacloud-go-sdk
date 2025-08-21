// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListRecommendContentResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListRecommendContentResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListRecommendContentResponse
	GetStatusCode() *int32
	SetBody(v *ListRecommendContentResponseBody) *ListRecommendContentResponse
	GetBody() *ListRecommendContentResponseBody
}

type ListRecommendContentResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListRecommendContentResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListRecommendContentResponse) String() string {
	return dara.Prettify(s)
}

func (s ListRecommendContentResponse) GoString() string {
	return s.String()
}

func (s *ListRecommendContentResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListRecommendContentResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListRecommendContentResponse) GetBody() *ListRecommendContentResponseBody {
	return s.Body
}

func (s *ListRecommendContentResponse) SetHeaders(v map[string]*string) *ListRecommendContentResponse {
	s.Headers = v
	return s
}

func (s *ListRecommendContentResponse) SetStatusCode(v int32) *ListRecommendContentResponse {
	s.StatusCode = &v
	return s
}

func (s *ListRecommendContentResponse) SetBody(v *ListRecommendContentResponseBody) *ListRecommendContentResponse {
	s.Body = v
	return s
}

func (s *ListRecommendContentResponse) Validate() error {
	return dara.Validate(s)
}
