// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListHotelSceneBookItemsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListHotelSceneBookItemsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListHotelSceneBookItemsResponse
	GetStatusCode() *int32
	SetBody(v *ListHotelSceneBookItemsResponseBody) *ListHotelSceneBookItemsResponse
	GetBody() *ListHotelSceneBookItemsResponseBody
}

type ListHotelSceneBookItemsResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListHotelSceneBookItemsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListHotelSceneBookItemsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListHotelSceneBookItemsResponse) GoString() string {
	return s.String()
}

func (s *ListHotelSceneBookItemsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListHotelSceneBookItemsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListHotelSceneBookItemsResponse) GetBody() *ListHotelSceneBookItemsResponseBody {
	return s.Body
}

func (s *ListHotelSceneBookItemsResponse) SetHeaders(v map[string]*string) *ListHotelSceneBookItemsResponse {
	s.Headers = v
	return s
}

func (s *ListHotelSceneBookItemsResponse) SetStatusCode(v int32) *ListHotelSceneBookItemsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListHotelSceneBookItemsResponse) SetBody(v *ListHotelSceneBookItemsResponseBody) *ListHotelSceneBookItemsResponse {
	s.Body = v
	return s
}

func (s *ListHotelSceneBookItemsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
