// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListHotelSceneItemsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListHotelSceneItemsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListHotelSceneItemsResponse
	GetStatusCode() *int32
	SetBody(v *ListHotelSceneItemsResponseBody) *ListHotelSceneItemsResponse
	GetBody() *ListHotelSceneItemsResponseBody
}

type ListHotelSceneItemsResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListHotelSceneItemsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListHotelSceneItemsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListHotelSceneItemsResponse) GoString() string {
	return s.String()
}

func (s *ListHotelSceneItemsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListHotelSceneItemsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListHotelSceneItemsResponse) GetBody() *ListHotelSceneItemsResponseBody {
	return s.Body
}

func (s *ListHotelSceneItemsResponse) SetHeaders(v map[string]*string) *ListHotelSceneItemsResponse {
	s.Headers = v
	return s
}

func (s *ListHotelSceneItemsResponse) SetStatusCode(v int32) *ListHotelSceneItemsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListHotelSceneItemsResponse) SetBody(v *ListHotelSceneItemsResponseBody) *ListHotelSceneItemsResponse {
	s.Body = v
	return s
}

func (s *ListHotelSceneItemsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
