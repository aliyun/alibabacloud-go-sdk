// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListHotelServiceCategoryResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListHotelServiceCategoryResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListHotelServiceCategoryResponse
	GetStatusCode() *int32
	SetBody(v *ListHotelServiceCategoryResponseBody) *ListHotelServiceCategoryResponse
	GetBody() *ListHotelServiceCategoryResponseBody
}

type ListHotelServiceCategoryResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListHotelServiceCategoryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListHotelServiceCategoryResponse) String() string {
	return dara.Prettify(s)
}

func (s ListHotelServiceCategoryResponse) GoString() string {
	return s.String()
}

func (s *ListHotelServiceCategoryResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListHotelServiceCategoryResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListHotelServiceCategoryResponse) GetBody() *ListHotelServiceCategoryResponseBody {
	return s.Body
}

func (s *ListHotelServiceCategoryResponse) SetHeaders(v map[string]*string) *ListHotelServiceCategoryResponse {
	s.Headers = v
	return s
}

func (s *ListHotelServiceCategoryResponse) SetStatusCode(v int32) *ListHotelServiceCategoryResponse {
	s.StatusCode = &v
	return s
}

func (s *ListHotelServiceCategoryResponse) SetBody(v *ListHotelServiceCategoryResponseBody) *ListHotelServiceCategoryResponse {
	s.Body = v
	return s
}

func (s *ListHotelServiceCategoryResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
