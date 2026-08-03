// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSearchHotelListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SearchHotelListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SearchHotelListResponse
	GetStatusCode() *int32
	SetBody(v *SearchHotelListResponseBody) *SearchHotelListResponse
	GetBody() *SearchHotelListResponseBody
}

type SearchHotelListResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SearchHotelListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SearchHotelListResponse) String() string {
	return dara.Prettify(s)
}

func (s SearchHotelListResponse) GoString() string {
	return s.String()
}

func (s *SearchHotelListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SearchHotelListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SearchHotelListResponse) GetBody() *SearchHotelListResponseBody {
	return s.Body
}

func (s *SearchHotelListResponse) SetHeaders(v map[string]*string) *SearchHotelListResponse {
	s.Headers = v
	return s
}

func (s *SearchHotelListResponse) SetStatusCode(v int32) *SearchHotelListResponse {
	s.StatusCode = &v
	return s
}

func (s *SearchHotelListResponse) SetBody(v *SearchHotelListResponseBody) *SearchHotelListResponse {
	s.Body = v
	return s
}

func (s *SearchHotelListResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
