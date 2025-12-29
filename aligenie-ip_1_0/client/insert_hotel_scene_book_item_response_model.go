// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInsertHotelSceneBookItemResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *InsertHotelSceneBookItemResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *InsertHotelSceneBookItemResponse
	GetStatusCode() *int32
	SetBody(v *InsertHotelSceneBookItemResponseBody) *InsertHotelSceneBookItemResponse
	GetBody() *InsertHotelSceneBookItemResponseBody
}

type InsertHotelSceneBookItemResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *InsertHotelSceneBookItemResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s InsertHotelSceneBookItemResponse) String() string {
	return dara.Prettify(s)
}

func (s InsertHotelSceneBookItemResponse) GoString() string {
	return s.String()
}

func (s *InsertHotelSceneBookItemResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *InsertHotelSceneBookItemResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *InsertHotelSceneBookItemResponse) GetBody() *InsertHotelSceneBookItemResponseBody {
	return s.Body
}

func (s *InsertHotelSceneBookItemResponse) SetHeaders(v map[string]*string) *InsertHotelSceneBookItemResponse {
	s.Headers = v
	return s
}

func (s *InsertHotelSceneBookItemResponse) SetStatusCode(v int32) *InsertHotelSceneBookItemResponse {
	s.StatusCode = &v
	return s
}

func (s *InsertHotelSceneBookItemResponse) SetBody(v *InsertHotelSceneBookItemResponseBody) *InsertHotelSceneBookItemResponse {
	s.Body = v
	return s
}

func (s *InsertHotelSceneBookItemResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
