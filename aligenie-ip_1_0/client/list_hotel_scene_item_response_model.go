// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListHotelSceneItemResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListHotelSceneItemResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListHotelSceneItemResponse
	GetStatusCode() *int32
	SetBody(v *ListHotelSceneItemResponseBody) *ListHotelSceneItemResponse
	GetBody() *ListHotelSceneItemResponseBody
}

type ListHotelSceneItemResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListHotelSceneItemResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListHotelSceneItemResponse) String() string {
	return dara.Prettify(s)
}

func (s ListHotelSceneItemResponse) GoString() string {
	return s.String()
}

func (s *ListHotelSceneItemResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListHotelSceneItemResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListHotelSceneItemResponse) GetBody() *ListHotelSceneItemResponseBody {
	return s.Body
}

func (s *ListHotelSceneItemResponse) SetHeaders(v map[string]*string) *ListHotelSceneItemResponse {
	s.Headers = v
	return s
}

func (s *ListHotelSceneItemResponse) SetStatusCode(v int32) *ListHotelSceneItemResponse {
	s.StatusCode = &v
	return s
}

func (s *ListHotelSceneItemResponse) SetBody(v *ListHotelSceneItemResponseBody) *ListHotelSceneItemResponse {
	s.Body = v
	return s
}

func (s *ListHotelSceneItemResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
