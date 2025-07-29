// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iHotNewsRecommendResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *HotNewsRecommendResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *HotNewsRecommendResponse
	GetStatusCode() *int32
	SetBody(v *HotNewsRecommendResponseBody) *HotNewsRecommendResponse
	GetBody() *HotNewsRecommendResponseBody
}

type HotNewsRecommendResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *HotNewsRecommendResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s HotNewsRecommendResponse) String() string {
	return dara.Prettify(s)
}

func (s HotNewsRecommendResponse) GoString() string {
	return s.String()
}

func (s *HotNewsRecommendResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *HotNewsRecommendResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *HotNewsRecommendResponse) GetBody() *HotNewsRecommendResponseBody {
	return s.Body
}

func (s *HotNewsRecommendResponse) SetHeaders(v map[string]*string) *HotNewsRecommendResponse {
	s.Headers = v
	return s
}

func (s *HotNewsRecommendResponse) SetStatusCode(v int32) *HotNewsRecommendResponse {
	s.StatusCode = &v
	return s
}

func (s *HotNewsRecommendResponse) SetBody(v *HotNewsRecommendResponseBody) *HotNewsRecommendResponse {
	s.Body = v
	return s
}

func (s *HotNewsRecommendResponse) Validate() error {
	return dara.Validate(s)
}
