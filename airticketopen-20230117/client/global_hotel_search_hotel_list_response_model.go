// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGlobalHotelSearchHotelListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GlobalHotelSearchHotelListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GlobalHotelSearchHotelListResponse
	GetStatusCode() *int32
	SetBody(v *GlobalHotelSearchHotelListResponseBody) *GlobalHotelSearchHotelListResponse
	GetBody() *GlobalHotelSearchHotelListResponseBody
}

type GlobalHotelSearchHotelListResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GlobalHotelSearchHotelListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GlobalHotelSearchHotelListResponse) String() string {
	return dara.Prettify(s)
}

func (s GlobalHotelSearchHotelListResponse) GoString() string {
	return s.String()
}

func (s *GlobalHotelSearchHotelListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GlobalHotelSearchHotelListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GlobalHotelSearchHotelListResponse) GetBody() *GlobalHotelSearchHotelListResponseBody {
	return s.Body
}

func (s *GlobalHotelSearchHotelListResponse) SetHeaders(v map[string]*string) *GlobalHotelSearchHotelListResponse {
	s.Headers = v
	return s
}

func (s *GlobalHotelSearchHotelListResponse) SetStatusCode(v int32) *GlobalHotelSearchHotelListResponse {
	s.StatusCode = &v
	return s
}

func (s *GlobalHotelSearchHotelListResponse) SetBody(v *GlobalHotelSearchHotelListResponseBody) *GlobalHotelSearchHotelListResponse {
	s.Body = v
	return s
}

func (s *GlobalHotelSearchHotelListResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
