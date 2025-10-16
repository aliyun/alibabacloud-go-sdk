// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteEduRoomResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteEduRoomResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteEduRoomResponse
	GetStatusCode() *int32
	SetBody(v *DeleteEduRoomResponseBody) *DeleteEduRoomResponse
	GetBody() *DeleteEduRoomResponseBody
}

type DeleteEduRoomResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteEduRoomResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteEduRoomResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteEduRoomResponse) GoString() string {
	return s.String()
}

func (s *DeleteEduRoomResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteEduRoomResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteEduRoomResponse) GetBody() *DeleteEduRoomResponseBody {
	return s.Body
}

func (s *DeleteEduRoomResponse) SetHeaders(v map[string]*string) *DeleteEduRoomResponse {
	s.Headers = v
	return s
}

func (s *DeleteEduRoomResponse) SetStatusCode(v int32) *DeleteEduRoomResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteEduRoomResponse) SetBody(v *DeleteEduRoomResponseBody) *DeleteEduRoomResponse {
	s.Body = v
	return s
}

func (s *DeleteEduRoomResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
