// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteHotelSceneBookItemResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteHotelSceneBookItemResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteHotelSceneBookItemResponse
	GetStatusCode() *int32
	SetBody(v *DeleteHotelSceneBookItemResponseBody) *DeleteHotelSceneBookItemResponse
	GetBody() *DeleteHotelSceneBookItemResponseBody
}

type DeleteHotelSceneBookItemResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteHotelSceneBookItemResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteHotelSceneBookItemResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteHotelSceneBookItemResponse) GoString() string {
	return s.String()
}

func (s *DeleteHotelSceneBookItemResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteHotelSceneBookItemResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteHotelSceneBookItemResponse) GetBody() *DeleteHotelSceneBookItemResponseBody {
	return s.Body
}

func (s *DeleteHotelSceneBookItemResponse) SetHeaders(v map[string]*string) *DeleteHotelSceneBookItemResponse {
	s.Headers = v
	return s
}

func (s *DeleteHotelSceneBookItemResponse) SetStatusCode(v int32) *DeleteHotelSceneBookItemResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteHotelSceneBookItemResponse) SetBody(v *DeleteHotelSceneBookItemResponseBody) *DeleteHotelSceneBookItemResponse {
	s.Body = v
	return s
}

func (s *DeleteHotelSceneBookItemResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
