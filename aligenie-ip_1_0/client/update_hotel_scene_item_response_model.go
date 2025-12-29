// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateHotelSceneItemResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateHotelSceneItemResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateHotelSceneItemResponse
	GetStatusCode() *int32
	SetBody(v *UpdateHotelSceneItemResponseBody) *UpdateHotelSceneItemResponse
	GetBody() *UpdateHotelSceneItemResponseBody
}

type UpdateHotelSceneItemResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateHotelSceneItemResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateHotelSceneItemResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateHotelSceneItemResponse) GoString() string {
	return s.String()
}

func (s *UpdateHotelSceneItemResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateHotelSceneItemResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateHotelSceneItemResponse) GetBody() *UpdateHotelSceneItemResponseBody {
	return s.Body
}

func (s *UpdateHotelSceneItemResponse) SetHeaders(v map[string]*string) *UpdateHotelSceneItemResponse {
	s.Headers = v
	return s
}

func (s *UpdateHotelSceneItemResponse) SetStatusCode(v int32) *UpdateHotelSceneItemResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateHotelSceneItemResponse) SetBody(v *UpdateHotelSceneItemResponseBody) *UpdateHotelSceneItemResponse {
	s.Body = v
	return s
}

func (s *UpdateHotelSceneItemResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
