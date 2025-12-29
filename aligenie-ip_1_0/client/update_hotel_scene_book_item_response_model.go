// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateHotelSceneBookItemResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateHotelSceneBookItemResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateHotelSceneBookItemResponse
	GetStatusCode() *int32
	SetBody(v *UpdateHotelSceneBookItemResponseBody) *UpdateHotelSceneBookItemResponse
	GetBody() *UpdateHotelSceneBookItemResponseBody
}

type UpdateHotelSceneBookItemResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateHotelSceneBookItemResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateHotelSceneBookItemResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateHotelSceneBookItemResponse) GoString() string {
	return s.String()
}

func (s *UpdateHotelSceneBookItemResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateHotelSceneBookItemResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateHotelSceneBookItemResponse) GetBody() *UpdateHotelSceneBookItemResponseBody {
	return s.Body
}

func (s *UpdateHotelSceneBookItemResponse) SetHeaders(v map[string]*string) *UpdateHotelSceneBookItemResponse {
	s.Headers = v
	return s
}

func (s *UpdateHotelSceneBookItemResponse) SetStatusCode(v int32) *UpdateHotelSceneBookItemResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateHotelSceneBookItemResponse) SetBody(v *UpdateHotelSceneBookItemResponseBody) *UpdateHotelSceneBookItemResponse {
	s.Body = v
	return s
}

func (s *UpdateHotelSceneBookItemResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
