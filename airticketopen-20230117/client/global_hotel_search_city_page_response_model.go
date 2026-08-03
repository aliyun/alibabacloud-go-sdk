// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGlobalHotelSearchCityPageResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GlobalHotelSearchCityPageResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GlobalHotelSearchCityPageResponse
	GetStatusCode() *int32
	SetBody(v *GlobalHotelSearchCityPageResponseBody) *GlobalHotelSearchCityPageResponse
	GetBody() *GlobalHotelSearchCityPageResponseBody
}

type GlobalHotelSearchCityPageResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GlobalHotelSearchCityPageResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GlobalHotelSearchCityPageResponse) String() string {
	return dara.Prettify(s)
}

func (s GlobalHotelSearchCityPageResponse) GoString() string {
	return s.String()
}

func (s *GlobalHotelSearchCityPageResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GlobalHotelSearchCityPageResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GlobalHotelSearchCityPageResponse) GetBody() *GlobalHotelSearchCityPageResponseBody {
	return s.Body
}

func (s *GlobalHotelSearchCityPageResponse) SetHeaders(v map[string]*string) *GlobalHotelSearchCityPageResponse {
	s.Headers = v
	return s
}

func (s *GlobalHotelSearchCityPageResponse) SetStatusCode(v int32) *GlobalHotelSearchCityPageResponse {
	s.StatusCode = &v
	return s
}

func (s *GlobalHotelSearchCityPageResponse) SetBody(v *GlobalHotelSearchCityPageResponseBody) *GlobalHotelSearchCityPageResponse {
	s.Body = v
	return s
}

func (s *GlobalHotelSearchCityPageResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
