// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMobileRecommendResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *MobileRecommendResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *MobileRecommendResponse
	GetStatusCode() *int32
	SetBody(v *MobileRecommendResponseBody) *MobileRecommendResponse
	GetBody() *MobileRecommendResponseBody
}

type MobileRecommendResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *MobileRecommendResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s MobileRecommendResponse) String() string {
	return dara.Prettify(s)
}

func (s MobileRecommendResponse) GoString() string {
	return s.String()
}

func (s *MobileRecommendResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *MobileRecommendResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *MobileRecommendResponse) GetBody() *MobileRecommendResponseBody {
	return s.Body
}

func (s *MobileRecommendResponse) SetHeaders(v map[string]*string) *MobileRecommendResponse {
	s.Headers = v
	return s
}

func (s *MobileRecommendResponse) SetStatusCode(v int32) *MobileRecommendResponse {
	s.StatusCode = &v
	return s
}

func (s *MobileRecommendResponse) SetBody(v *MobileRecommendResponseBody) *MobileRecommendResponse {
	s.Body = v
	return s
}

func (s *MobileRecommendResponse) Validate() error {
	return dara.Validate(s)
}
