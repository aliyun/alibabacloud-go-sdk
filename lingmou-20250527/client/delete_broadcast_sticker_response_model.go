// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteBroadcastStickerResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteBroadcastStickerResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteBroadcastStickerResponse
	GetStatusCode() *int32
	SetBody(v *DeleteBroadcastStickerResponseBody) *DeleteBroadcastStickerResponse
	GetBody() *DeleteBroadcastStickerResponseBody
}

type DeleteBroadcastStickerResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteBroadcastStickerResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteBroadcastStickerResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteBroadcastStickerResponse) GoString() string {
	return s.String()
}

func (s *DeleteBroadcastStickerResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteBroadcastStickerResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteBroadcastStickerResponse) GetBody() *DeleteBroadcastStickerResponseBody {
	return s.Body
}

func (s *DeleteBroadcastStickerResponse) SetHeaders(v map[string]*string) *DeleteBroadcastStickerResponse {
	s.Headers = v
	return s
}

func (s *DeleteBroadcastStickerResponse) SetStatusCode(v int32) *DeleteBroadcastStickerResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteBroadcastStickerResponse) SetBody(v *DeleteBroadcastStickerResponseBody) *DeleteBroadcastStickerResponse {
	s.Body = v
	return s
}

func (s *DeleteBroadcastStickerResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
