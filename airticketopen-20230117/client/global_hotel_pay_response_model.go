// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGlobalHotelPayResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GlobalHotelPayResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GlobalHotelPayResponse
	GetStatusCode() *int32
	SetBody(v *GlobalHotelPayResponseBody) *GlobalHotelPayResponse
	GetBody() *GlobalHotelPayResponseBody
}

type GlobalHotelPayResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GlobalHotelPayResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GlobalHotelPayResponse) String() string {
	return dara.Prettify(s)
}

func (s GlobalHotelPayResponse) GoString() string {
	return s.String()
}

func (s *GlobalHotelPayResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GlobalHotelPayResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GlobalHotelPayResponse) GetBody() *GlobalHotelPayResponseBody {
	return s.Body
}

func (s *GlobalHotelPayResponse) SetHeaders(v map[string]*string) *GlobalHotelPayResponse {
	s.Headers = v
	return s
}

func (s *GlobalHotelPayResponse) SetStatusCode(v int32) *GlobalHotelPayResponse {
	s.StatusCode = &v
	return s
}

func (s *GlobalHotelPayResponse) SetBody(v *GlobalHotelPayResponseBody) *GlobalHotelPayResponse {
	s.Body = v
	return s
}

func (s *GlobalHotelPayResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
